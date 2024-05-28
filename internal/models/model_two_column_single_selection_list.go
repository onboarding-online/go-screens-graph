package models

// TwoColumnSingleSelectionList - Two column single selection list element
type TwoColumnSingleSelectionList struct {

	ItemType TwoColumnSingleSelectionListItemType `json:"itemType"`

	// Item rows
	Items []ItemTypeSelection `json:"items"`

	Styles ListBlock `json:"styles"`

	SelectedBlock SelectedListItemBlock `json:"selectedBlock"`

	Box Box `json:"box"`
}
