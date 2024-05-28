package models

// ProgressBarItemContent - Progress bar item content
type ProgressBarItemContent struct {

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle,omitempty"`

	Description Text `json:"description,omitempty"`

	Image Image `json:"image,omitempty"`
}
