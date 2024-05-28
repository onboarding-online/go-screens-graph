package models

// BaseText - Base text parameters
type BaseText struct {

	Kind TextKind `json:"kind"`

	// Localized string
	L10n map[string]string `json:"l10n"`

	Parameters TemplateParameters `json:"parameters"`

	Styles LabelBlock `json:"styles"`
}
