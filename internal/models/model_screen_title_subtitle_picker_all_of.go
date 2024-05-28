package models

// ScreenTitleSubtitlePickerAllOf - Screen - Title/Subtitle/Picker
type ScreenTitleSubtitlePickerAllOf struct {

	TitleSubtitlePickerDescription string `json:"titleSubtitlePickerDescription"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Picker Picker `json:"picker"`
}
