package models

type ScreenImageTitleSubtitles struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	ImageTitleSubtitlesDescription string `json:"imageTitleSubtitlesDescription"`

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle1 Text `json:"subtitle1"`

	Subtitle2 Text `json:"subtitle2"`
}
