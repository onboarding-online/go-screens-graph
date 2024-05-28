package models

type ScreenBasicPaywallBlock struct {

	Background BackgroundStyle `json:"background"`

	// Padding left for body
	PaddingLeft float32 `json:"paddingLeft,omitempty"`

	// Padding right for body
	PaddingRight float32 `json:"paddingRight,omitempty"`

	// Padding top for body
	PaddingTop float32 `json:"paddingTop,omitempty"`

	// Padding bottom for body
	PaddingBottom float32 `json:"paddingBottom,omitempty"`

	BodyColor string `json:"bodyColor,omitempty"`

	BodyOpacity string `json:"bodyOpacity,omitempty"`

	ImageVerticalPosition PaywallImageVerticalPositionKind `json:"imageVerticalPosition,omitempty"`

	BodyBackgroundColor string `json:"bodyBackgroundColor,omitempty"`

	BodyBackgroundOpacity float32 `json:"bodyBackgroundOpacity,omitempty"`

	BodyKind BasicPaywallBodyKind `json:"bodyKind,omitempty"`

	BodyStyle BasicPaywallBodyStyle `json:"bodyStyle,omitempty"`

	Blur BlurKind `json:"blur,omitempty"`
}
