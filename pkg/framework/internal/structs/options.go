package structs

import (
	"os"
	"runtime"

	"github.com/ozontech/allure-go/pkg/allure"
)

// Options struct provide easy way to configure tests in suites
type Options struct {
	TestName    string
	SuiteName   string
	ParentName  string
	FullName    string
	PackageName string
	ThreadName  string
	Framework   string
	Language    string
	Host        string
}

// OptionName ...
type OptionName string

// OptionName constants. Describes supported option parameters
const (
	TestName    OptionName = "TestName"
	SuiteName   OptionName = "SuiteName"
	ParentName  OptionName = "ParentName"
	FullName    OptionName = "FullName"
	PackageName OptionName = "PackageName"
	Framework   OptionName = "Framework"
	ThreadName  OptionName = "FullName"
	Language    OptionName = "Language"
	Host        OptionName = "Host"
)

// NewOptions returns pointer to the new Options structure
func NewOptions(opts map[OptionName]string) *Options {
	var testName string
	var suiteName string
	var parentName string
	var fullName string
	var packageName string
	var threadName string
	var framework string
	var language string
	var host string

	for name, value := range opts {
		if name == TestName && opts[name] != "" {
			testName = value
		}
		if name == SuiteName && opts[name] != "" {
			suiteName = value
		}
		if name == ParentName && opts[name] != "" {
			parentName = value
		}
		if name == PackageName && opts[name] != "" {
			packageName = value
		}
		if name == FullName && opts[name] != "" {
			fullName = value
		}
		if name == Framework && opts[name] != "" {
			framework = value
		}
		if name == Language && opts[name] != "" {
			language = value
		}
		if name == Host && opts[name] != "" {
			host = value
		}
		if name == ThreadName && opts[name] != "" {
			threadName = value
		}
	}
	if framework == "" {
		framework = allure.DefaultVersion
	}
	if host == "" {
		hostName, _ := os.Hostname()
		host = hostName
	}
	if language == "" {
		language = runtime.Version()
	}

	return &Options{
		TestName:    testName,
		SuiteName:   suiteName,
		ParentName:  parentName,
		FullName:    fullName,
		PackageName: packageName,
		ThreadName:  threadName,
		Framework:   framework,
		Language:    language,
		Host:        host,
	}
}
