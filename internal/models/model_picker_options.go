package models

// PickerOptions - Picker options
type PickerOptions struct {

	From float32 `json:"from"`

	To float32 `json:"to"`

	// Localized string
	LocalizedOptions map[string][]string `json:"localizedOptions"`
}
