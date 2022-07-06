package async

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type MixedAsyncSuite struct {
	suite.Suite
}

func (s *MixedAsyncSuite) BeforeEach(t provider.T) {
	t.Epic("Async")
	t.Feature("Mixed Suite")
	t.Tags("async", "suite", "steps")
}

func (s *MixedAsyncSuite) TestSelectionProductsLists(t provider.T) {
	t.SkipOnPrint()
	testCases := []struct {
		testName string
	}{{"test1"}, {"test2"}, {"test3"}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t provider.T) {
			name := tc.testName
			t.Parallel()
			t.NewStep(name)
		})
	}
}

func (s *MixedAsyncSuite) TestSelectionProductsLists2(t provider.T) {
	t.SkipOnPrint()
	t.Parallel()
	testCases := []struct {
		testName string
	}{{"test1"}, {"test2"}, {"test3"}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t provider.T) {
			name := tc.testName
			t.Parallel()
			t.NewStep(name)
		})
	}
}

func (s *MixedAsyncSuite) TestSelectionProductsLists3(t provider.T) {
	t.SkipOnPrint()
	testCases := []struct {
		testName string
	}{{"test1"}, {"test2"}, {"test3"}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t provider.T) {
			name := tc.testName
			t.Parallel()
			t.NewStep(name)
			t.Fatalf("WHOOPS")
		})
	}
}

func (s *MixedAsyncSuite) TestSelectionProductsLists4(t provider.T) {
	t.SkipOnPrint()
	testCases := []struct {
		testName string
	}{{"test1"}, {"test2"}, {"test3"}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t provider.T) {
			name := tc.testName
			t.Parallel()
			t.NewStep(name)
			panic("WHOOPS")
		})
	}
}

func TestMixedAsyncSuite(t *testing.T) {
	suite.RunSuite(t, new(MixedAsyncSuite))
}
