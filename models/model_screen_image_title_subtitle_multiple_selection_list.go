package models

type ScreenImageTitleSubtitleMultipleSelectionList struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	ImageTitleSubtitleMultipleSelectionListDescription string `json:"imageTitleSubtitleMultipleSelectionListDescription"`

	Image Image `json:"image"`

	// Defines the scale of the image in percentage
	ImageScale float32 `json:"imageScale,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List MultipleSelectionList `json:"list"`
}
