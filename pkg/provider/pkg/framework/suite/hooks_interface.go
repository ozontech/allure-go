package suite

import (
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

// AllureBeforeTest has a BeforeEach method, which will run before each
// test in the suite.
type AllureBeforeTest interface {
	BeforeEach(t provider.T)
}

// AllureAfterTest has a AfterEach method, which will run after
// each test in the suite.
type AllureAfterTest interface {
	AfterEach(t provider.T)
}

// AllureBeforeSuite has a BeforeAll method, which will run before the
// tests in the suite are run.
type AllureBeforeSuite interface {
	BeforeAll(t provider.T)
}

// AllureAfterSuite has a AfterAll method, which will run after
// all the tests in the suite have been run.
type AllureAfterSuite interface {
	AfterAll(t provider.T)
}
