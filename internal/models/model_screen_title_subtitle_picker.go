package models

type ScreenTitleSubtitlePicker struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	TitleSubtitlePickerDescription string `json:"titleSubtitlePickerDescription"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Picker Picker `json:"picker"`
}
