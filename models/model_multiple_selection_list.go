package models

// MultipleSelectionList - Multiple selection list element
type MultipleSelectionList struct {

	ItemType MultipleSelectionListItemType `json:"itemType"`

	// Item rows
	Items []ItemTypeSelection `json:"items"`

	Styles ListBlock `json:"styles"`

	SelectedBlock SelectedListItemBlock `json:"selectedBlock"`

	Box Box `json:"box"`
}
