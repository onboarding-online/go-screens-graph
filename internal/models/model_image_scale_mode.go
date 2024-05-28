package models
// ImageScaleMode : Image scale mode
type ImageScaleMode string

// List of ImageScaleMode
const (
	IMAGESCALEMODE_SCALE_TO_FILL ImageScaleMode = "scaleToFill"
	IMAGESCALEMODE_SCALE_ASPECT_FIT ImageScaleMode = "scaleAspectFit"
	IMAGESCALEMODE_SCALE_ASPECT_FILL ImageScaleMode = "scaleAspectFill"
	IMAGESCALEMODE_CENTER ImageScaleMode = "center"
	IMAGESCALEMODE_TOP ImageScaleMode = "top"
	IMAGESCALEMODE_BOTTOM ImageScaleMode = "bottom"
	IMAGESCALEMODE_LEFT ImageScaleMode = "left"
	IMAGESCALEMODE_RIGHT ImageScaleMode = "right"
)
