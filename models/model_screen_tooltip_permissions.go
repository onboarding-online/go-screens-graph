package models

type ScreenTooltipPermissions struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	TooltipPermissionsDescription string `json:"tooltipPermissionsDescription"`

	Tooltip Tooltip `json:"tooltip"`
}
