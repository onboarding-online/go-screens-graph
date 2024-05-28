package models

// Divider - Divider element
type Divider struct {

	Box Box `json:"box"`

	Kind BaseDividerKind `json:"kind"`

	Styles BaseDividerBlock `json:"styles"`
}
