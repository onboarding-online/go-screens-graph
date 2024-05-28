package models

// SingleSelectionList - Single selection list element
type SingleSelectionList struct {

	ItemType SingleSelectionListItemType `json:"itemType"`

	// Item rows
	Items []ItemTypeSelection `json:"items"`

	Styles ListBlock `json:"styles"`

	SelectedBlock SelectedListItemBlock `json:"selectedBlock"`

	Box Box `json:"box"`
}
