package models

// Badge - Badge parameters
type Badge struct {

	Box Box `json:"box"`

	Kind TextKind `json:"kind"`

	// Localized string
	L10n map[string]string `json:"l10n"`

	Parameters TemplateParameters `json:"parameters"`

	Styles BadgeBlock `json:"styles"`
}
