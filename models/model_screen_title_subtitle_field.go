package models

type ScreenTitleSubtitleField struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	TitleSubtitleFieldDescription string `json:"titleSubtitleFieldDescription"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Field Field `json:"field"`
}
