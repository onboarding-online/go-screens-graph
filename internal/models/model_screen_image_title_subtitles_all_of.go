package models

// ScreenImageTitleSubtitlesAllOf - Screen - Image/Title/Subtitles list
type ScreenImageTitleSubtitlesAllOf struct {

	ImageTitleSubtitlesDescription string `json:"imageTitleSubtitlesDescription"`

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle1 Text `json:"subtitle1"`

	Subtitle2 Text `json:"subtitle2"`
}
