package models

// SlideContent - Slide content
type SlideContent struct {

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle1 Text `json:"subtitle1"`

	Subtitle2 Text `json:"subtitle2"`
}
