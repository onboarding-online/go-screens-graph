package models

// CheckBoxBlock - Styles for checkbox
type CheckBoxBlock struct {

	// Color of checkbox
	Color string `json:"color,omitempty"`

	IsBackgroundFilled bool `json:"isBackgroundFilled,omitempty"`

	Width float32 `json:"width,omitempty"`

	Height float32 `json:"height,omitempty"`
}
