package models

// Image - Image parameters
type Image struct {

	Box Box `json:"box"`

	// Dictionary of localized Asset
	L10n map[string]Asset `json:"l10n"`

	Styles ImageBlock `json:"styles"`
}
