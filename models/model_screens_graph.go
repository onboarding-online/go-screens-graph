package models

// ScreensGraph - The graph of screens for onboarding
type ScreensGraph struct {

	DefaultLanguage Language `json:"defaultLanguage"`

	// The list of languages
	Languages []Language `json:"languages"`

	// Schema version in semver format
	SchemaVersion string `json:"schemaVersion,omitempty"`

	// First screen id
	LaunchScreenId string `json:"launchScreenId"`

	// All screens nodes dictionary
	Screens map[string]Screen `json:"screens"`

	// Metadata dictionary
	Metadata map[string]string `json:"metadata"`
}
