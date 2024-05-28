package models

// ElementTimer - Timer to show an element
type ElementTimer struct {

	// Duration to show an element in seconds
	Duration int32 `json:"duration"`

	Action Action `json:"action"`

	Unit TimerUnit `json:"unit,omitempty"`
}
