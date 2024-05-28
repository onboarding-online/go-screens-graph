package models

type ScreenBasicPaywall struct {

	NavigationBar PaywallNavigationBar `json:"navigationBar"`

	Footer PaywallFooter `json:"footer"`

	Styles ScreenBasicPaywallBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	ScreenBasicPaywall string `json:"screenBasicPaywall"`

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
}
