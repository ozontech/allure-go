package common

import "github.com/ozontech/allure-go/pkg/allure"

/*
Labels
*/

// Label provides possibility to add any Label to test result
func (c *common) Label(label allure.Label) {
	c.safely(func(result *allure.Result) {
		result.Labels = append(result.Labels, label)
	})
}

// Labels provides possibility to add few Labels to test result
func (c *common) Labels(labels ...allure.Label) {
	c.safely(func(result *allure.Result) {
		result.Labels = append(result.Labels, labels...)
	})
}

func (c *common) ReplaceLabel(label allure.Label) {
	c.safely(func(result *allure.Result) {
		for idx := range result.Labels {
			if result.Labels[idx].Name == label.Name {
				result.Labels[idx].Value = label.Value
				return
			}
		}
		c.Label(label)
	})
}

// Epic adds Epic label to test result
func (c *common) Epic(value string) {
	c.Label(allure.EpicLabel(value))
}

// Feature adds Feature label to test result
func (c *common) Feature(value string) {
	c.Label(allure.FeatureLabel(value))
}

// Story adds Story label to test result
func (c *common) Story(value string) {
	c.Label(allure.StoryLabel(value))
}

// FrameWork adds FrameWork label to test result
func (c *common) FrameWork(value string) {
	c.Label(allure.FrameWorkLabel(value))
}

// Host adds Host label to test result
func (c *common) Host(value string) {
	c.Label(allure.HostLabel(value))
}

// Thread adds Thread label to test result
// Seems like there is no way to access an identifier for the current goroutine in Go.
func (c *common) Thread(value string) {
	c.Label(allure.ThreadLabel(value))
}

// ID adds ID label to test result
func (c *common) ID(value string) {
	c.Label(allure.IDLabel(value))
}

// Language adds Language label to test result
func (c *common) Language(value string) {
	c.Label(allure.LanguageLabel(value))
}

// AddSuiteLabel adds suite label to test result
func (c *common) AddSuiteLabel(value string) {
	c.Label(allure.SuiteLabel(value))
}

// AddSubSuite adds AddSubSuite label to test result
func (c *common) AddSubSuite(value string) {
	c.Label(allure.SubSuiteLabel(value))
}

// AddParentSuite adds AddParentSuite label to test result
func (c *common) AddParentSuite(value string) {
	c.Label(allure.ParentSuiteLabel(value))
}

// Severity adds Severity label to test result
func (c *common) Severity(value allure.SeverityType) {
	c.Label(allure.SeverityLabel(value))
}

// Tag adds Tag label to test result
func (c *common) Tag(value string) {
	c.Label(allure.TagLabel(value))
}

// Tags adds a multiple Tag label to test result
func (c *common) Tags(values ...string) {
	c.Labels(allure.TagLabels(values...)...)
}

// Package adds Package label to test result
func (c *common) Package(value string) {
	c.Label(allure.PackageLabel(value))
}

// Owner adds Owner label to test result
func (c *common) Owner(value string) {
	c.Label(allure.OwnerLabel(value))
}

// Lead adds Lead label to test result
func (c *common) Lead(value string) {
	c.Label(allure.LeadLabel(value))
}

func (c *common) AllureID(value string) {
	c.Label(allure.IDAllureLabel(value))
}
