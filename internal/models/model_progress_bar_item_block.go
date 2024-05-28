package models

// ProgressBarItemBlock - Styles for progress bar item
type ProgressBarItemBlock struct {

	// Color for progress bar
	Color string `json:"color,omitempty"`

	// Track color
	TrackColor string `json:"trackColor,omitempty"`

	// Fill color
	FillColor string `json:"fillColor,omitempty"`
}
