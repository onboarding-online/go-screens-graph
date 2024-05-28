package models

// Action - Parameters for action
type Action struct {

	// List of conditioned actions
	Edges []ConditionedAction `json:"edges"`

	Kind ActionKind `json:"kind"`
}
