package suite_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type AllureIdSuite struct {
	suite.Suite
}

func (s *AllureIdSuite) BeforeAll(t provider.T) {
	// code that can fail here
}

func (s *AllureIdSuite) TestMyTestWithAllureID(t provider.T) {
	// code of your test here
}

func TestNewDemo(t *testing.T) {
	ais := new(AllureIdSuite)
	ais.AddAllureIDMapping("TestMyTestWithAllureID", "12345")
	suite.RunSuite(t, ais)
}
