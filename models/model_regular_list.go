package models

// RegularList - Regular list element
type RegularList struct {

	ItemType RegularListItemType `json:"itemType"`

	// Item rows
	Items []ItemTypeRegular `json:"items"`

	Styles ListBlock `json:"styles"`

	Box Box `json:"box"`
}
