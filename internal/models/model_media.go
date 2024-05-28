package models

// Media - Media parameters
type Media struct {

	Kind MediaKind `json:"kind"`

	Content MediaContent `json:"content"`

	Styles MediaBlock `json:"styles"`

	Box Box `json:"box"`
}
