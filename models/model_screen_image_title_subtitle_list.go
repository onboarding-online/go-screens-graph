package models

type ScreenImageTitleSubtitleList struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	ImageTitleSubtitleListDescription string `json:"imageTitleSubtitleListDescription"`

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List RegularList `json:"list"`
}
