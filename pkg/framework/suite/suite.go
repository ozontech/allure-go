package suite

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/internal/file_manager"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/provider"
)

// AllureSuite is an interface that describes Suite behaviour
type AllureSuite interface {
	SetT(t *provider.T)
	T() *provider.T
	GetName() string
	setName(string)
	GetPackage() string
	setPackage(string)
	GetParent() string
	setParent(string)
}

// Suite is test-class like object, that allows group tests in test suites.
type Suite struct {
	name        string
	parent      string
	packageName string

	t *provider.T
}

// SetT setting t context to the suite
func (suite *Suite) SetT(t *provider.T) {
	suite.t = t
}

func (suite *Suite) setName(name string) {
	suite.name = name
}

func (suite *Suite) setPackage(name string) {
	suite.packageName = name
}

func (suite *Suite) setParent(name string) {
	suite.parent = name
}

// GetName returns suite's name
func (suite *Suite) GetName() string {
	return suite.name
}

// GetPackage returns suite's package
func (suite *Suite) GetPackage() string {
	return suite.packageName
}

// GetParent returns suite's parent
func (suite *Suite) GetParent() string {
	return suite.parent
}

// SkipOnPrint makes allure-testify skip current test results printing
func (suite *Suite) SkipOnPrint() {
	suite.T().GetResult().SkipOnPrint()
}

// T returns suite's *provider.T
func (suite *Suite) T() *provider.T {
	return suite.t
}

//RunTest works with parallel parametrizing, but you cannot use s as provider
func (suite *Suite) RunTest(testName string, test func(t *provider.T), tags ...string) bool {
	return suite.T().Run(testName, test, tags...)
}

//Run doesn't work with parallel parametrizing
func (suite *Suite) Run(testName string, test func(), tags ...string) bool {
	oldT := suite.T()
	realT := oldT.RealT()

	result := allure.NewResultHelper().GetNewResult(oldT, testName, oldT.GetPackage(), tags...)

	res := realT.Run(testName, func(t *testing.T) {
		newT := provider.NewTForTest(oldT, result, allure.NewContainer())
		// dirty magic
		newT.T = t
		suite.SetT(newT)
		defer func() {
			newT.GetResult().Finish()
			newT.GetResult().Done()
		}()
		defer func() {
			r := recover()
			if r != nil {
				errMsg := fmt.Sprintf("test panicked: %v\n%s", r, debug.Stack())
				newT.BreakResult(errMsg)
				newT.Errorf(errMsg)
				newT.FailNow()
			}
		}()
		test()
	})
	defer suite.SetT(oldT)
	return res
}

// RunSuite runs child suite of current suite
func (suite *Suite) RunSuite(t *provider.T, newSuite AllureSuite) {

	newSuite.setName(getSuiteName(newSuite))
	newSuite.setParent(suite.T().RealT().Name())
	newSuite.setPackage(file_manager.GetPackage(2))
	t.GetResult().SkipOnPrint()
	kek := testing.InternalTest{Name: newSuite.GetName(), F: func(t *testing.T) {
		runner.RunSuite(t, newSuite)
	}}

	realT := t.RealT()
	realT.Run(kek.Name, kek.F)
}

func getSuiteName(suite interface{}) string {
	t := reflect.TypeOf(suite)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}
