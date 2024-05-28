package models

// BackgroundStyle - Background styles configuration
type BackgroundStyle struct {

	Kind BackgroundStyleKind `json:"kind"`

	Styles BackgroundStyleStyles `json:"styles"`
}
