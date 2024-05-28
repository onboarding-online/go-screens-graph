package models

// ScreenTableSingleSelectionAllOf - Screen table single selection - Title/Subtitle/Selection list
type ScreenTableSingleSelectionAllOf struct {

	SingleSelectionDescription string `json:"singleSelectionDescription"`

	Media Media `json:"media,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List SingleSelectionList `json:"list"`
}
