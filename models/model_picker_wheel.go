package models

// PickerWheel - Picker wheel
type PickerWheel struct {

	Kind PickerWheelKind `json:"kind"`

	Options PickerOptions `json:"options"`

	// Localized string
	DefaultValue map[string]string `json:"defaultValue"`
}
