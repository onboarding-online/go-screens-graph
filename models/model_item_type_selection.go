package models

// ItemTypeSelection - Structure of full collection item type for selection list
type ItemTypeSelection struct {

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle Text `json:"subtitle"`

	CheckBox CheckBox `json:"checkBox"`

	Action Action `json:"action"`

	Box Box `json:"box"`

	Weight float32 `json:"weight"`
}
