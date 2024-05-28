package models

// LoaderBlock - Styles for loader block
type LoaderBlock struct {

	// Color for loader
	Color string `json:"color,omitempty"`

	// Padding left for container
	PaddingLeft float32 `json:"paddingLeft,omitempty"`

	// Padding right for container
	PaddingRight float32 `json:"paddingRight,omitempty"`

	// Padding top for container
	PaddingTop float32 `json:"paddingTop,omitempty"`

	// Padding bottom for container
	PaddingBottom float32 `json:"paddingBottom,omitempty"`
}
