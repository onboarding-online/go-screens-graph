package models

// BaseNavLinkContent - Base navigation link content
type BaseNavLinkContent struct {

	Kind TextKind `json:"kind"`

	// Dictionary of localized Asset
	L10n map[string]Asset `json:"l10n"`

	Parameters TemplateParameters `json:"parameters"`

	Styles ImageBlock `json:"styles"`
}
