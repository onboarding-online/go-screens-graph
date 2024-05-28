package models

// VideoBlock - Styles for video
type VideoBlock struct {

	ScaleMode VideoScaleMode `json:"scaleMode,omitempty"`

	// width of video
	Width float32 `json:"width,omitempty"`

	// height of video
	Height float32 `json:"height,omitempty"`

	// Height in percentage
	HeightPercentage float32 `json:"heightPercentage,omitempty"`

	CornerRadiusLeftTop float32 `json:"cornerRadiusLeftTop,omitempty"`

	CornerRadiusLeftBottom float32 `json:"cornerRadiusLeftBottom,omitempty"`

	CornerRadiusRightTop float32 `json:"cornerRadiusRightTop,omitempty"`

	CornerRadiusRightBottom float32 `json:"cornerRadiusRightBottom,omitempty"`

	// apply corner radius for all corners
	MainCornerRadius float32 `json:"mainCornerRadius,omitempty"`

	// repeat video after finish
	Repeat bool `json:"repeat,omitempty"`
}
