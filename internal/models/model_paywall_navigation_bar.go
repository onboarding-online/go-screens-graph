package models

// PaywallNavigationBar - Paywall navigation bar
type PaywallNavigationBar struct {

	PaywallNavigationBar string `json:"paywallNavigationBar"`

	Close Button `json:"close,omitempty"`

	Restore Text `json:"restore,omitempty"`

	Styles PaywallNavigationBarBlock `json:"styles"`

	Back Button `json:"back,omitempty"`

	Skip Button `json:"skip,omitempty"`

	PageIndicator PageIndicator `json:"pageIndicator,omitempty"`

	PageIndicatorKind PageIndicatorKind `json:"pageIndicatorKind,omitempty"`

	DashesPageIndicator DashesPageIndicator `json:"dashesPageIndicator,omitempty"`
}
