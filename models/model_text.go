package models

// Text - Text parameters
type Text struct {

	Box Box `json:"box"`

	Kind TextKind `json:"kind"`

	// Localized string
	L10n map[string]string `json:"l10n"`

	Parameters TemplateParameters `json:"parameters"`

	Styles LabelBlock `json:"styles"`
}
