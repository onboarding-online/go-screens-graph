package models

// ScreenBasicPaywallBlockAllOf - Styles for ScreenBasicPaywall
type ScreenBasicPaywallBlockAllOf struct {

	BodyColor string `json:"bodyColor,omitempty"`

	BodyOpacity string `json:"bodyOpacity,omitempty"`

	ImageVerticalPosition PaywallImageVerticalPositionKind `json:"imageVerticalPosition,omitempty"`

	BodyBackgroundColor string `json:"bodyBackgroundColor,omitempty"`

	BodyBackgroundOpacity float32 `json:"bodyBackgroundOpacity,omitempty"`

	BodyKind BasicPaywallBodyKind `json:"bodyKind,omitempty"`

	BodyStyle BasicPaywallBodyStyle `json:"bodyStyle,omitempty"`

	Blur BlurKind `json:"blur,omitempty"`
}
