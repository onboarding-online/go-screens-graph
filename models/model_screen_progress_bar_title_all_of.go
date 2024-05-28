package models

// ScreenProgressBarTitleAllOf - Screen - Progress bar/Title
type ScreenProgressBarTitleAllOf struct {

	ProgressBarTitleDescription string `json:"progressBarTitleDescription"`

	ProgressBar ProgressBar `json:"progressBar"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle,omitempty"`
}
