package models

type ScreenTwoColumnSingleSelection struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	TwoColumnSingleSelectionDescription string `json:"twoColumnSingleSelectionDescription"`

	Media Media `json:"media,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List TwoColumnSingleSelectionList `json:"list"`
}
