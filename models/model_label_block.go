package models

// LabelBlock - Styles for label block
type LabelBlock struct {

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
}
