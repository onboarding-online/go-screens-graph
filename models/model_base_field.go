package models

// BaseField - Base field element
type BaseField struct {

	Type FieldType `json:"type"`

	Subtype FieldSubtype `json:"subtype,omitempty"`

	ValidationRegex string `json:"validationRegex,omitempty"`

	ErrorMessage BaseText `json:"errorMessage,omitempty"`

	Placeholder BaseText `json:"placeholder"`

	// Field value
	Value string `json:"value"`

	Styles FieldBlock `json:"styles"`
}
