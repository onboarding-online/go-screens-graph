package models

// BadgeBlockAllOf - Styles for badge block
type BadgeBlockAllOf struct {

	Position BadgePosition `json:"position,omitempty"`

	// Radius of border
	BorderRadius float32 `json:"borderRadius,omitempty"`

	// Width of border
	BorderWidth float32 `json:"borderWidth,omitempty"`

	// Color of border
	BorderColor string `json:"borderColor,omitempty"`
}
