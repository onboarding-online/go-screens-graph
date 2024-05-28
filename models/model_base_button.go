package models

// BaseButton - Base button parameters
type BaseButton struct {

	Kind ButtonKind `json:"kind"`

	Content BaseButtonContent `json:"content"`

	Styles ButtonBlock `json:"styles"`

	Action Action `json:"action"`
}
