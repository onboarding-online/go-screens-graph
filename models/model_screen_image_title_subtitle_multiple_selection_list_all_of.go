package models

// ScreenImageTitleSubtitleMultipleSelectionListAllOf - Screen - Image/Title/Subtitle/Multiple selection list
type ScreenImageTitleSubtitleMultipleSelectionListAllOf struct {

	ImageTitleSubtitleMultipleSelectionListDescription string `json:"imageTitleSubtitleMultipleSelectionListDescription"`

	Image Image `json:"image"`

	// Defines the scale of the image in percentage
	ImageScale float32 `json:"imageScale,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List MultipleSelectionList `json:"list"`
}
