package models

// RequestPermission - Request permissions object
type RequestPermission struct {

	Type ScreenPermissionType `json:"type"`

	Value bool `json:"value"`

	Action Action `json:"action"`
}
