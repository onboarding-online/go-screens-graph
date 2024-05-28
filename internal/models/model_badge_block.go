package models

type BadgeBlock struct {

	TextAlign LabelPosition `json:"textAlign,omitempty"`

	FontFamily FontFamily `json:"fontFamily,omitempty"`

	// Height of line
	LineHeight float32 `json:"lineHeight,omitempty"`

	// Font size
	FontSize float32 `json:"fontSize,omitempty"`

	// Font weight
	FontWeight float32 `json:"fontWeight,omitempty"`

	// Color for text
	Color string `json:"color,omitempty"`

	// Alternative color for text
	AlternativeColor string `json:"alternativeColor,omitempty"`

	// Background color for text block
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Padding left for container
	PaddingLeft float32 `json:"paddingLeft,omitempty"`

	// Padding right for container
	PaddingRight float32 `json:"paddingRight,omitempty"`

	// Padding top for container
	PaddingTop float32 `json:"paddingTop,omitempty"`

	// Padding bottom for container
	PaddingBottom float32 `json:"paddingBottom,omitempty"`

	Position BadgePosition `json:"position,omitempty"`

	// Radius of border
	BorderRadius float32 `json:"borderRadius,omitempty"`

	// Width of border
	BorderWidth float32 `json:"borderWidth,omitempty"`

	// Color of border
	BorderColor string `json:"borderColor,omitempty"`
}
