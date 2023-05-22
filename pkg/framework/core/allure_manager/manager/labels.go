package manager

import (
	"github.com/ozontech/allure-go/pkg/allure"
)

/*
Labels
*/

// Label provides possibility to add any Label to test result
func (a *allureManager) Label(label *allure.Label) {
	a.safely(func(result *allure.Result) {
		result.Labels = append(result.Labels, label)
	})
}

// Labels provides possibility to add few Labels to test result
func (a *allureManager) Labels(labels ...*allure.Label) {
	a.safely(func(result *allure.Result) {
		result.Labels = append(result.Labels, labels...)
	})
}

func (a *allureManager) ReplaceLabel(label *allure.Label) {
	a.safely(func(result *allure.Result) {
		for idx := range result.Labels {
			if result.Labels[idx].Name == label.Name {
				result.Labels[idx].Value = label.Value
				return
			}
		}
		a.Label(label)
	})
}

// Epic adds Epic label to test result
func (a *allureManager) Epic(value string) {
	a.Label(allure.EpicLabel(value))
}

// Layer adds Layer label to test result
func (a *allureManager) Layer(value string) {
	a.Label(allure.LayerLabel(value))
}

// Feature adds Feature label to test result
func (a *allureManager) Feature(value string) {
	a.Label(allure.FeatureLabel(value))
}

// Story adds Story label to test result
func (a *allureManager) Story(value string) {
	a.Label(allure.StoryLabel(value))
}

// FrameWork adds FrameWork label to test result
func (a *allureManager) FrameWork(value string) {
	a.ReplaceLabel(allure.FrameWorkLabel(value))
}

// Host adds Host label to test result
func (a *allureManager) Host(value string) {
	a.ReplaceLabel(allure.HostLabel(value))
}

// Thread adds Thread label to test result
// Seems like there is no way to access an identifier for the current goroutine in Go.
func (a *allureManager) Thread(value string) {
	a.ReplaceLabel(allure.ThreadLabel(value))
}

// ID adds ID label to test result
func (a *allureManager) ID(value string) {
	a.ReplaceLabel(allure.IDLabel(value))
}

// Language adds Language label to test result
func (a *allureManager) Language(value string) {
	a.ReplaceLabel(allure.LanguageLabel(value))
}

// AddSuiteLabel adds suite label to test result
func (a *allureManager) AddSuiteLabel(value string) {
	a.Label(allure.SuiteLabel(value))
}

// AddSubSuite adds AddSubSuite label to test result
func (a *allureManager) AddSubSuite(value string) {
	a.Label(allure.SubSuiteLabel(value))
}

// AddParentSuite adds AddParentSuite label to test result
func (a *allureManager) AddParentSuite(value string) {
	a.Label(allure.ParentSuiteLabel(value))
}

// Severity adds Severity label to test result
func (a *allureManager) Severity(value allure.SeverityType) {
	a.ReplaceLabel(allure.SeverityLabel(value))
}

// Tag adds Tag label to test result
func (a *allureManager) Tag(value string) {
	a.Label(allure.TagLabel(value))
}

// Tags adds a multiple Tag label to test result
func (a *allureManager) Tags(values ...string) {
	a.Labels(allure.TagLabels(values...)...)
}

// Package adds Package label to test result
func (a *allureManager) Package(value string) {
	a.ReplaceLabel(allure.PackageLabel(value))
}

// Owner adds Owner label to test result
func (a *allureManager) Owner(value string) {
	a.Label(allure.OwnerLabel(value))
}

// Lead adds Lead label to test result
func (a *allureManager) Lead(value string) {
	a.Label(allure.LeadLabel(value))
}

func (a *allureManager) AllureID(value string) {
	a.ReplaceLabel(allure.IDAllureLabel(value))
}
