package models

// ScreenTwoColumnMultipleSelectionAllOf - Screen two column table multiple selection - Title/Subtitle/Selection list
type ScreenTwoColumnMultipleSelectionAllOf struct {

	TwoColumnMultipleSelectionDescription string `json:"twoColumnMultipleSelectionDescription"`

	Media Media `json:"media,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List TwoColumnMultipleSelectionList `json:"list"`
}
