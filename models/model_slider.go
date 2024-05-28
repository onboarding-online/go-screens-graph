package models

// Slider - Slider
type Slider struct {

	Kind SliderKind `json:"kind"`

	Timer ElementTimer `json:"timer"`

	// Sections for slider
	Items []SliderItem `json:"items"`

	Styles SliderBlock `json:"styles"`

	Box Box `json:"box"`
}
