package suite_demo

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type SomeInterface interface {
	TestMe()
}

type someSubInterface struct {
	idx int
}

func (i *someSubInterface) TestMe() {
	return
}

type SomeSubStruct struct {
	name string
}

type SomeStruct struct {
	test int
	sub  *SomeSubStruct
	int  SomeInterface
}

func (str *SomeStruct) String() string {
	return str.sub.name
}

type ParametrizedSuite struct {
	suite.Suite
	ParamStrings []*SomeStruct
}

func (s *ParametrizedSuite) BeforeAll(t provider.T) {
	for i := 0; i < 10; i++ {
		s.ParamStrings = append(s.ParamStrings, &SomeStruct{
			test: i,
			sub:  &SomeSubStruct{name: fmt.Sprintf("TestName - %d", i)},
			int:  &someSubInterface{idx: i + 1},
		})
	}
}

func (s *ParametrizedSuite) TableTestStrings(t provider.T, stringParam *SomeStruct) {
	t.Parallel()
	t.Require().NotEqual(2, stringParam.test)
	t.Require().NotEmpty(stringParam)
}

func (s *ParametrizedSuite) TestJustTest(t provider.T) {
	t.Parallel()
	t.Require().NotNil(t)
}

func TestNewParametrizedDemo(t *testing.T) {
	suite.RunSuite(t, new(ParametrizedSuite))
}
