package models

// ScreenTitleSubtitleFieldAllOf - Screen - Title/Subtitle/Field
type ScreenTitleSubtitleFieldAllOf struct {

	TitleSubtitleFieldDescription string `json:"titleSubtitleFieldDescription"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Field Field `json:"field"`
}
