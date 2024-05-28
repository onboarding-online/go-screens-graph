package models

// SliderItem - Information about one slide of slider
type SliderItem struct {

	// Section display start
	ValueFrom int32 `json:"valueFrom"`

	// Section display end
	ValueTo int32 `json:"valueTo"`

	Content SlideContent `json:"content"`
}
