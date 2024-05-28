package models
// VideoScaleMode : Video scale mode
type VideoScaleMode string

// List of VideoScaleMode
const (
	VIDEOSCALEMODE_SCALE_TO_FILL VideoScaleMode = "scaleToFill"
	VIDEOSCALEMODE_SCALE_ASPECT_FIT VideoScaleMode = "scaleAspectFit"
	VIDEOSCALEMODE_SCALE_ASPECT_FILL VideoScaleMode = "scaleAspectFill"
	VIDEOSCALEMODE_CENTER VideoScaleMode = "center"
	VIDEOSCALEMODE_TOP VideoScaleMode = "top"
	VIDEOSCALEMODE_BOTTOM VideoScaleMode = "bottom"
	VIDEOSCALEMODE_LEFT VideoScaleMode = "left"
	VIDEOSCALEMODE_RIGHT VideoScaleMode = "right"
)
