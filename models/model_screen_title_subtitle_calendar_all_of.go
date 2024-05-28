package models

// ScreenTitleSubtitleCalendarAllOf - Screen - Title/Subtitle/Calendar
type ScreenTitleSubtitleCalendarAllOf struct {

	TitleSubtitleCalendarDescription string `json:"titleSubtitleCalendarDescription"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Calendar Calendar `json:"calendar"`
}
