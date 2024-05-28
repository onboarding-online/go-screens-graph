package models

// BasicScreenBlock - Basic screen styles
type BasicScreenBlock struct {

	Background BackgroundStyle `json:"background"`

	// Padding left for body
	PaddingLeft float32 `json:"paddingLeft,omitempty"`

	// Padding right for body
	PaddingRight float32 `json:"paddingRight,omitempty"`

	// Padding top for body
	PaddingTop float32 `json:"paddingTop,omitempty"`

	// Padding bottom for body
	PaddingBottom float32 `json:"paddingBottom,omitempty"`
}
