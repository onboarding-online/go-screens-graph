package models

// ScreenImageTitleSubtitleListAllOf - Screen - Image/Title/Subtitle/Regular list
type ScreenImageTitleSubtitleListAllOf struct {

	ImageTitleSubtitleListDescription string `json:"imageTitleSubtitleListDescription"`

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List RegularList `json:"list"`
}
