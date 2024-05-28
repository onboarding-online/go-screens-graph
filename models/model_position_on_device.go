package models

// PositionOnDevice - Information about position on device in percentage
type PositionOnDevice struct {

	// x offset in percentage 0-100
	X float32 `json:"x"`

	// y offset in percentage 0-100
	Y float32 `json:"y"`
}
