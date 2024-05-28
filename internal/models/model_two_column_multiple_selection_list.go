package models

// TwoColumnMultipleSelectionList - Two column multiple selection list element
type TwoColumnMultipleSelectionList struct {

	ItemType TwoColumnMultipleSelectionListItemType `json:"itemType"`

	// Item rows
	Items []ItemTypeSelection `json:"items"`

	Styles ListBlock `json:"styles"`

	SelectedBlock SelectedListItemBlock `json:"selectedBlock"`

	Box Box `json:"box"`
}
