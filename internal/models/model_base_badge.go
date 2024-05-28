package models

type BaseBadge struct {

	Kind TextKind `json:"kind"`

	// Localized string
	L10n map[string]string `json:"l10n"`

	Parameters TemplateParameters `json:"parameters"`

	Styles BadgeBlock `json:"styles"`
}
