package framework

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/constructors"
	"github.com/ozontech/allure-go/pkg/framework/internal/structs"
	"github.com/ozontech/allure-go/pkg/provider"
)

var matchMethod = flag.String("allure-testify-1.m", "", "regular expression to select tests of the testify suite to run")

type TestSuite interface {
	SetT(t *provider.T)
	T() *provider.T
	GetName() string
	GetParent() string
	GetPackage() string
}

type Test struct {
	Name string
	F    func(t *provider.T)
}

type InternalSuite struct {
	runName       string
	packageName   string
	parentName    string
	languageName  string
	frameworkName string
	hostName      string
	suite         TestSuite
	suiteName     string
	tests         []*InternalTest
}

func NewInternalSuite(realT *testing.T, packageName string, suite TestSuite) *InternalSuite {
	parts := strings.Split(realT.Name(), "/")
	runName := parts[0]
	hostName, _ := os.Hostname()
	internalSuite := &InternalSuite{
		runName:       runName,
		packageName:   packageName,
		languageName:  runtime.Version(),
		frameworkName: allure.DefaultVersion,
		hostName:      hostName,
		suite:         suite,
	}
	methodFinder := reflect.TypeOf(suite)
	if suite.GetName() == "" {
		internalSuite.suiteName = methodFinder.Elem().Name()
	} else {
		internalSuite.suiteName = suite.GetName()
	}

	if suite.GetParent() != "" {
		internalSuite.parentName = suite.GetParent()
	}

	t := provider.NewT(realT, internalSuite.runName, internalSuite.suiteName)

	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)

		ok, err := methodFilter(method.Name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "testify: invalid regexp for -m: %s\n", err)
			os.Exit(1)
		}

		if !ok {
			continue
		}

		internalSuite.addTest(t, method)
	}

	return internalSuite
}

func (is *InternalSuite) Run(t *testing.T) {
	newT := provider.NewT(t, is.runName, is.suiteName)
	is.suite.SetT(newT)
	defer func() {
		container := newT.GetContainer()
		for _, test := range is.tests {
			container.AddChild(test.result.UUID)
		}
		_ = container.Print()
	}()
	defer func() {
		if afterAll, ok := is.suite.(AllureAfterSuite); ok {
			newT.WithAfterSuite(afterAll.AfterAll)
		}
		newT.GetContainer().Finish()
	}()
	newT.GetContainer().Begin()
	if beforeAll, ok := is.suite.(AllureBeforeSuite); ok {
		newT.WithBeforeSuite(beforeAll.BeforeAll)
	}

	if len(is.tests) == 0 {
		newT.Errorf("[WARN]: no tests to run!")
		return
	}

	for _, test := range is.tests {
		t.Run(test.testName, func(realT *testing.T) {
			tForTest := provider.NewT(realT, is.runName, is.suiteName)
			test.testFunction(tForTest)
		})
	}
}

func (is *InternalSuite) addTest(t *provider.T, method reflect.Method) {

	if parts := strings.Split(t.Name(), "/"); len(parts) > 2 && is.parentName != "" {
		is.parentName = parts[len(parts)-2]
	}

	opts := make(map[structs.OptionName]string)
	opts[structs.TestName] = method.Name
	opts[structs.FullName] = fmt.Sprintf("%s/%s", t.Name(), opts[structs.TestName])
	opts[structs.Host] = is.hostName
	opts[structs.Language] = is.languageName
	opts[structs.Framework] = is.frameworkName
	opts[structs.PackageName] = is.packageName
	opts[structs.SuiteName] = is.suiteName
	opts[structs.ParentName] = is.parentName
	opts[structs.ThreadName] = opts[structs.FullName]

	_opts := structs.NewOptions(opts)

	result := constructors.NewResultHelper().GetNewResultFromOpts(_opts)

	is.tests = append(is.tests, NewTest(func(t *provider.T) {

		t.GetContainer().AddChild(result.UUID)
		oldT := is.suite.T()

		newT := provider.NewTForTest(t, result, t.GetContainer())

		defer func() {
			result.Done()
		}()

		defer func() {
			r := recover()
			is.suite.SetT(oldT)
			if r != nil {
				errMsg := fmt.Sprintf("test panicked: %v\n%s", r, debug.Stack())
				newT.BreakResult(errMsg)
				newT.Errorf(errMsg)
				newT.FailNow()
			}
		}()
		is.suite.SetT(newT)
		defer func() {
			result.Finish()
			if tearDownTest, ok := is.suite.(AllureAfterTest); ok {
				provider.ExecuteAfter(newT, tearDownTest.AfterEach)
			}
			result.Container.Finish()
		}()

		result.Container.Begin()
		if setupTest, ok := is.suite.(AllureBeforeTest); ok {
			provider.ExecuteBefore(newT, setupTest.BeforeEach)
		}

		result.Begin()
		method.Func.Call([]reflect.Value{reflect.ValueOf(is.suite)})
	}, result))
}

// Filtering method according to set regular expression
// specified command-line argument -m
func methodFilter(name string) (bool, error) {
	if ok, _ := regexp.MatchString("^Test", name); !ok {
		return false, nil
	}
	return regexp.MatchString(*matchMethod, name)
}
