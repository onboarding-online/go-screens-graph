package models
// MediaScaleMode : Media scale mode
type MediaScaleMode string

// List of MediaScaleMode
const (
	MEDIASCALEMODE_SCALE_TO_FILL MediaScaleMode = "scaleToFill"
	MEDIASCALEMODE_SCALE_ASPECT_FIT MediaScaleMode = "scaleAspectFit"
	MEDIASCALEMODE_SCALE_ASPECT_FILL MediaScaleMode = "scaleAspectFill"
	MEDIASCALEMODE_CENTER MediaScaleMode = "center"
	MEDIASCALEMODE_TOP MediaScaleMode = "top"
	MEDIASCALEMODE_BOTTOM MediaScaleMode = "bottom"
	MEDIASCALEMODE_LEFT MediaScaleMode = "left"
	MEDIASCALEMODE_RIGHT MediaScaleMode = "right"
)
