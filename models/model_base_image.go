package models

// BaseImage - Base image parameters
type BaseImage struct {

	// Dictionary of localized Asset
	L10n map[string]Asset `json:"l10n"`

	Styles ImageBlock `json:"styles"`
}
