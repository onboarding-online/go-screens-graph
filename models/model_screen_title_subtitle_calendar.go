package models

type ScreenTitleSubtitleCalendar struct {

	NavigationBar NavigationBar `json:"navigationBar"`

	Footer Footer `json:"footer"`

	Styles BasicScreenBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	TitleSubtitleCalendarDescription string `json:"titleSubtitleCalendarDescription"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Calendar Calendar `json:"calendar"`
}
