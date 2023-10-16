package allure

import (
	"fmt"
	"strings"
)

// Label is the implementation of the label.
// A label is an entity used by Allure to make metrics and grouping of tests.
type Label struct {
	Name  string      `json:"name"`  // Label's name
	Value interface{} `json:"value"` // Label's value
}

// GetValue returns label value as string
func (l *Label) GetValue() string {
	return strings.Trim(fmt.Sprintf("%s", l.Value), "\"")
}

// NewLabel - builds and returns a new allure.Label. The label key depends on the passed LabelType.
func NewLabel(labelType LabelType, value string) *Label {
	return &Label{
		Name:  labelType.ToString(),
		Value: value,
	}
}

type LabelType string

// LabelType constants
const (
	Epic        LabelType = "epic"
	Layer       LabelType = "layer"
	Feature     LabelType = "feature"
	Story       LabelType = "story"
	ID          LabelType = "as_id"
	Severity    LabelType = "severity"
	ParentSuite LabelType = "parentSuite"
	Suite       LabelType = "suite"
	SubSuite    LabelType = "subSuite"
	Package     LabelType = "package"
	Thread      LabelType = "thread"
	Host        LabelType = "host"
	Tag         LabelType = "tag"
	Framework   LabelType = "framework"
	Language    LabelType = "language"
	Owner       LabelType = "owner"
	Lead        LabelType = "lead"
	AllureID    LabelType = "ALLURE_ID"
)

func (l LabelType) ToString() string {
	return string(l)
}

type SeverityType string

// SeverityType constants
const (
	BLOCKER  SeverityType = "blocker"
	CRITICAL SeverityType = "critical"
	NORMAL   SeverityType = "normal"
	MINOR    SeverityType = "minor"
	TRIVIAL  SeverityType = "trivial"
)

// ToString casts SeverityType to string
func (s SeverityType) ToString() string {
	return string(s)
}

// LanguageLabel returns Language Label
func LanguageLabel(language string) *Label {
	return NewLabel(Language, language)
}

// FrameWorkLabel returns Framework Label
func FrameWorkLabel(framework string) *Label {
	return NewLabel(Framework, framework)
}

// IDLabel returns ID Label
func IDLabel(testID string) *Label {
	return NewLabel(ID, testID)
}

// TagLabel returns Tag Label
func TagLabel(tag string) *Label {
	return NewLabel(Tag, tag)
}

// TagLabels returns array of Tag Label
func TagLabels(tags ...string) []*Label {
	var result []*Label
	for _, tag := range tags {
		result = append(result, TagLabel(tag))
	}
	return result
}

// HostLabel returns Host Label
func HostLabel(host string) *Label {
	return NewLabel(Host, host)
}

// ThreadLabel returns Thread Label
func ThreadLabel(thread string) *Label {
	return NewLabel(Thread, thread)
}

// SeverityLabel returns Severity Label
func SeverityLabel(severity SeverityType) *Label {
	return NewLabel(Severity, severity.ToString())
}

// SubSuiteLabel returns SubSuite Label
func SubSuiteLabel(subSuite string) *Label {
	return NewLabel(SubSuite, subSuite)
}

// EpicLabel returns Epic Label
func EpicLabel(epic string) *Label {
	return NewLabel(Epic, epic)
}

// LayerLabel returns Layer Label
func LayerLabel(layer string) *Label {
	return NewLabel(Layer, layer)
}

// StoryLabel returns Story Label
func StoryLabel(story string) *Label {
	return NewLabel(Story, story)
}

// FeatureLabel returns Feature Label
func FeatureLabel(feature string) *Label {
	return NewLabel(Feature, feature)
}

// ParentSuiteLabel returns ParentSuite Label
func ParentSuiteLabel(parent string) *Label {
	return NewLabel(ParentSuite, parent)
}

// SuiteLabel returns Suite Label
func SuiteLabel(suite string) *Label {
	return NewLabel(Suite, suite)
}

// PackageLabel returns Package Label
func PackageLabel(packageName string) *Label {
	return NewLabel(Package, packageName)
}

// OwnerLabel returns Owner Label
func OwnerLabel(ownerName string) *Label {
	return NewLabel(Owner, ownerName)
}

// LeadLabel returns Lead Label
func LeadLabel(leadName string) *Label {
	return NewLabel(Lead, leadName)
}

// IDAllureLabel returns AllureID Label
func IDAllureLabel(allureID string) *Label {
	return NewLabel(AllureID, allureID)
}
