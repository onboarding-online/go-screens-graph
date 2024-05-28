package models

// ScreenBasicPaywallAllOf - Screen basic paywall
type ScreenBasicPaywallAllOf struct {

	ScreenBasicPaywall string `json:"screenBasicPaywall"`

	NavigationBar PaywallNavigationBar `json:"navigationBar,omitempty"`

	Footer PaywallFooter `json:"footer,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Divider Divider `json:"divider,omitempty"`

	Media Media `json:"media,omitempty"`

	List RegularList `json:"list"`

	Loader Loader `json:"loader,omitempty"`

	// Purchase flags
	Flags []PurchaseFlag `json:"flags"`

	Subscriptions SubscriptionList `json:"subscriptions"`

	CurrencyFormat CurrencyFormatKind `json:"currencyFormat"`

	Styles ScreenBasicPaywallBlock `json:"styles"`
}
