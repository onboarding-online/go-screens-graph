package models

// ScreenTwoColumnSingleSelectionAllOf - Screen two column table single selection - Title/Subtitle/Selection list
type ScreenTwoColumnSingleSelectionAllOf struct {

	TwoColumnSingleSelectionDescription string `json:"twoColumnSingleSelectionDescription"`

	Media Media `json:"media,omitempty"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	List TwoColumnSingleSelectionList `json:"list"`
}
