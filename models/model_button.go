package models

// Button - Button parameters
type Button struct {

	Box Box `json:"box"`

	Kind ButtonKind `json:"kind"`

	Content BaseButtonContent `json:"content"`

	Styles ButtonBlock `json:"styles"`

	Action Action `json:"action"`
}
