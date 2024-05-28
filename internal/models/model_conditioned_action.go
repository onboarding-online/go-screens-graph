package models

// ConditionedAction - Conditioned action
type ConditionedAction struct {

	// List of conditions
	Rule []Condition `json:"rule"`

	// Next screen id
	NextScreenId string `json:"nextScreenId"`

	TransitionKind ScreenTransitionKind `json:"transitionKind"`
}
