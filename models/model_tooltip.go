package models

// Tooltip - Tooltip parameters
type Tooltip struct {

	Image Image `json:"image"`

	Title Text `json:"title"`

	Position PositionOnDevice `json:"position"`

	Styles TooltipBlock `json:"styles"`

	Box Box `json:"box"`
}
