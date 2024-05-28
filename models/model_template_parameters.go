package models

// TemplateParameters - Template parameters
type TemplateParameters struct {

	// Dictionary of labels styles
	Labels map[string]LabelBlock `json:"labels"`

	// Template link dictionary
	Links map[string]string `json:"links"`
}
