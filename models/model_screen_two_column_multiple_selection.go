package models

type ScreenTwoColumnMultipleSelection struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	TwoColumnMultipleSelectionDescription string `json:"twoColumnMultipleSelectionDescription"`

	Media Media `json:"media,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List TwoColumnMultipleSelectionList `json:"list"`
}
