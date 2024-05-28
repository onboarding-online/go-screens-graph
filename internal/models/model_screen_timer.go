package models

// ScreenTimer - Timer to show the screen
type ScreenTimer struct {

	// Duration to show some screen in seconds
	Duration int32 `json:"duration"`

	Action Action `json:"action"`

	Unit TimerUnit `json:"unit,omitempty"`
}
