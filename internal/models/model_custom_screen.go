package models

type CustomScreen struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	CustomScreenDescription string `json:"customScreenDescription"`

	// Dictionary of output labels for custom screen
	Labels map[string]CustomScreenOutputLabel `json:"labels"`

	// Dictionary of input values from custom screen
	Values map[string]CustomScreenInputValue `json:"values"`

	Callback Callback `json:"callback"`
}
