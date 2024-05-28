package models

// ButtonBlock - Styles for button
type ButtonBlock struct {

	Visibility ButtonVisibility `json:"visibility,omitempty"`

	MovingAnimation ButtonMovingAnimation `json:"movingAnimation,omitempty"`

	// If true, the input will take up the full width of its container
	FullWidth bool `json:"fullWidth,omitempty"`

	// Background of button
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Padding left for container
	PaddingLeft float32 `json:"paddingLeft,omitempty"`

	// Padding right for container
	PaddingRight float32 `json:"paddingRight,omitempty"`

	// Padding top for container
	PaddingTop float32 `json:"paddingTop,omitempty"`

	// Padding bottom for container
	PaddingBottom float32 `json:"paddingBottom,omitempty"`

	// Radius of border
	BorderRadius float32 `json:"borderRadius,omitempty"`

	// Color of border
	BorderColor string `json:"borderColor,omitempty"`

	// Width of border
	BorderWidth float32 `json:"borderWidth,omitempty"`

	// Width of button
	Width float32 `json:"width,omitempty"`

	// Height of button
	Height float32 `json:"height,omitempty"`
}
