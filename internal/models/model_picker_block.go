package models

// PickerBlock - Styles for picker item
type PickerBlock struct {

	VerticalAlignment VerticalAlignment `json:"verticalAlignment,omitempty"`

	// Background color for picker
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
}
