package models

// Field - Field element
type Field struct {

	Box Box `json:"box"`

	Type FieldType `json:"type"`

	Subtype FieldSubtype `json:"subtype,omitempty"`

	ValidationRegex string `json:"validationRegex,omitempty"`

	ErrorMessage BaseText `json:"errorMessage,omitempty"`

	Placeholder BaseText `json:"placeholder"`

	// Field value
	Value string `json:"value"`

	Styles FieldBlock `json:"styles"`
}
