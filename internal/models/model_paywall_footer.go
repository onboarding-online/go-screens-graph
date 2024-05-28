package models

// PaywallFooter - Paywall footer
type PaywallFooter struct {

	PaywallFooter string `json:"paywallFooter"`

	Purchase Button `json:"purchase,omitempty"`

	AutoRenewLabel Text `json:"autoRenewLabel,omitempty"`

	BottomContainer PaywallFooterBottomContainer `json:"bottomContainer"`

	Styles PaywallFooterBlock `json:"styles"`

	Button1 Button `json:"button1,omitempty"`

	Button2 Button `json:"button2,omitempty"`

	Kind BasicFooterKind `json:"kind,omitempty"`
}
