package models

// PaywallFooterBlock - Styles for paywall footer
type PaywallFooterBlock struct {

	ElementsOrder PaywallFooterElementsOrder `json:"elementsOrder,omitempty"`

	// Background color for footer
	BackgroundColor string `json:"backgroundColor,omitempty"`

	Opacity float32 `json:"opacity,omitempty"`

	// Padding left for container
	PaddingLeft float32 `json:"paddingLeft,omitempty"`

	// Padding right for container
	PaddingRight float32 `json:"paddingRight,omitempty"`

	// Padding top for container
	PaddingTop float32 `json:"paddingTop,omitempty"`

	// Padding bottom for container
	PaddingBottom float32 `json:"paddingBottom,omitempty"`
}
