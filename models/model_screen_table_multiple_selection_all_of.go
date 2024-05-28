package models

// ScreenTableMultipleSelectionAllOf - Screen table multiple selection - Title/Subtitle/Selection list
type ScreenTableMultipleSelectionAllOf struct {

	MultipleSelectionDescription string `json:"multipleSelectionDescription"`

	Media Media `json:"media,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List MultipleSelectionList `json:"list"`
}
