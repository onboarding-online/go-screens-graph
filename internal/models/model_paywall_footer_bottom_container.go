package models

// PaywallFooterBottomContainer - Paywall footer bottom container
type PaywallFooterBottomContainer struct {

	PrivacyPolicy NavLink `json:"privacyPolicy,omitempty"`

	Restore Text `json:"restore,omitempty"`

	TermsOfUse NavLink `json:"termsOfUse,omitempty"`

	Styles PaywallFooterBottomContainerBlock `json:"styles"`
}
