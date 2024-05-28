package models

// Picker - Picker element
type Picker struct {

	Box Box `json:"box"`

	DataType FieldType `json:"dataType"`

	CurrentValue string `json:"currentValue"`

	// Localized string
	DefaultValue map[string]string `json:"defaultValue"`

	Separator string `json:"separator"`

	// Picker wheels list
	Wheels []PickerWheel `json:"wheels"`

	Styles PickerBlock `json:"styles"`

	LabelStyles LabelBlock `json:"labelStyles"`
}
