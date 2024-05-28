package models

// ScreenImageTitleSubtitlePickerAllOf - Screen - Image/Title/Subtitle/Picker
type ScreenImageTitleSubtitlePickerAllOf struct {

	ImageTitleSubtitlePickerDescription string `json:"imageTitleSubtitlePickerDescription"`

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	Picker Picker `json:"picker"`
}
