package models

// Condition - The condition
type Condition struct {

	// Condition key
	Key string `json:"key"`

	Operator Operator `json:"operator"`

	// Comparison value
	Value string `json:"value"`
}
