package provider

import (
	"github.com/ozontech/allure-go/pkg/allure"
)

type AllureLabels interface {
	ID(value string)
	Epic(value string)
	AddSuiteLabel(value string)
	AddSubSuite(value string)
	AddParentSuite(value string)
	Feature(value string)
	Story(value string)
	Tag(value string)
	Tags(values ...string)
	Package(value string)
	Severity(value allure.SeverityType)
	FrameWork(value string)
	Host(value string)
	Thread(value string)
	Language(value string)
	Owner(value string)
	Lead(value string)
	Label(label allure.Label)
	Labels(labels ...allure.Label)
}

/*
Labels
*/

// Label provides possibility to add any Label to test result
func (t *T) Label(label allure.Label) {
	t.safely(func(result *allure.Result) {
		result.Labels = append(result.Labels, label)
	})
}

// Labels provides possibility to add few Labels to test result
func (t *T) Labels(labels ...allure.Label) {
	t.safely(func(result *allure.Result) {
		result.Labels = append(result.Labels, labels...)
	})
}

func (t *T) ReplaceLabel(label allure.Label) {
	t.safely(func(result *allure.Result) {
		for idx := range result.Labels {
			if result.Labels[idx].Name == label.Name {
				result.Labels[idx].Value = label.Value
				return
			}
		}
		t.Label(label)
	})
}

// Feature adds Feature label to test result
func (t *T) Feature(value string) {
	t.Label(allure.FeatureLabel(value))
}

// Story adds Story label to test result
func (t *T) Story(value string) {
	t.Label(allure.StoryLabel(value))
}

// Epic adds Epic label to test result
func (t *T) Epic(value string) {
	t.Label(allure.EpicLabel(value))
}

// FrameWork adds FrameWork label to test result
func (t *T) FrameWork(value string) {
	t.Label(allure.FrameWorkLabel(value))
}

// Host adds Host label to test result
func (t *T) Host(value string) {
	t.Label(allure.HostLabel(value))
}

// Thread adds Thread label to test result
// Seems like there is no way to access an identifier for the current goroutine in Go.
func (t *T) Thread(value string) {
	t.Label(allure.ThreadLabel(value))
}

// ID adds ID label to test result
func (t *T) ID(value string) {
	t.Label(allure.IDLabel(value))
}

// Language adds Language label to test result
func (t *T) Language(value string) {
	t.Label(allure.LanguageLabel(value))
}

// AddSuiteLabel adds suite label to test result
func (t *T) AddSuiteLabel(value string) {
	t.Label(allure.SuiteLabel(value))
}

// AddSubSuite adds AddSubSuite label to test result
func (t *T) AddSubSuite(value string) {
	t.Label(allure.SubSuiteLabel(value))
}

// AddParentSuite adds AddParentSuite label to test result
func (t *T) AddParentSuite(value string) {
	t.Label(allure.ParentSuiteLabel(value))
}

// Severity adds Severity label to test result
func (t *T) Severity(value allure.SeverityType) {
	t.Label(allure.SeverityLabel(value))
}

// Tag adds Tag label to test result
func (t *T) Tag(value string) {
	t.Label(allure.TagLabel(value))
}

// Tags adds a multiple Tag label to test result
func (t *T) Tags(values ...string) {
	t.Labels(allure.TagLabels(values...)...)
}

// Package adds Package label to test result
func (t *T) Package(value string) {
	t.Label(allure.PackageLabel(value))
}

// Owner adds Owner label to test result
func (t *T) Owner(value string) {
	t.Label(allure.OwnerLabel(value))
}

// Lead adds Lead label to test result
func (t *T) Lead(value string) {
	t.Label(allure.LeadLabel(value))
}
