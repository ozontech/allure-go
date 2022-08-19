package suite_demo

import (
	"testing"

	"github.com/jackc/fake"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ParametrizedSuite struct {
	suite.Suite
	ParamCities []string
}

func (s *ParametrizedSuite) BeforeAll(t provider.T) {
	for i := 0; i < 10; i++ {
		s.ParamCities = append(s.ParamCities, fake.City())
	}
}

func (s *ParametrizedSuite) TableTestCities(t provider.T, city string) {
	t.Parallel()
	t.Require().NotEmpty(city)
}

func TestNewParametrizedDemo(t *testing.T) {
	suite.RunSuite(t, new(ParametrizedSuite))
}
