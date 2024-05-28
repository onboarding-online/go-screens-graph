package models

type ScreenProgressBarTitle struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	ProgressBarTitleDescription string `json:"progressBarTitleDescription"`

	ProgressBar ProgressBar `json:"progressBar"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle,omitempty"`
}
