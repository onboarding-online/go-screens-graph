package models

// BaseButtonContent - Base button content
type BaseButtonContent struct {

	Kind TextKind `json:"kind"`

	// Dictionary of localized Asset
	L10n map[string]Asset `json:"l10n"`

	Parameters TemplateParameters `json:"parameters"`

	Styles ImageBlock `json:"styles"`
}
