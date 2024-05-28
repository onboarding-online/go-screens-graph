package models

// ItemTypeRegular - Structure of full collection item type for regular list
type ItemTypeRegular struct {

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle,omitempty"`

	Box Box `json:"box"`
}
