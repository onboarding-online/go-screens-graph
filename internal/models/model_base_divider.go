package models

// BaseDivider - Base divider parameters
type BaseDivider struct {

	Kind BaseDividerKind `json:"kind"`

	Styles BaseDividerBlock `json:"styles"`
}
