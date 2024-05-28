package models

// BaseVideo - Base video parameters
type BaseVideo struct {

	// Dictionary of localized Asset
	L10n map[string]Asset `json:"l10n"`

	Styles VideoBlock `json:"styles"`
}
