package models

// ProgressBarBlock - Styles for progress bar
type ProgressBarBlock struct {

	// width of progress bar
	Width float32 `json:"width,omitempty"`

	// height of progress bar
	Height float32 `json:"height,omitempty"`

	// Height in percentage
	HeightPercentage float32 `json:"heightPercentage,omitempty"`

	// Color for progress bar
	Color string `json:"color,omitempty"`

	// Track color
	TrackColor string `json:"trackColor,omitempty"`

	// Fill color
	FillColor string `json:"fillColor,omitempty"`

	// Thickness for progress bar
	Thickness float32 `json:"thickness,omitempty"`

	// Track thickness
	TrackThickness float32 `json:"trackThickness,omitempty"`

	VerticalAlignment VerticalAlignment `json:"verticalAlignment,omitempty"`
}
