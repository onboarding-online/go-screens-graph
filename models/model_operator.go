package models
// Operator : Possible comparison operations
type Operator string

// List of Operator
const (
	OPERATOR_LT Operator = "lt"
	OPERATOR_GT Operator = "gt"
	OPERATOR_LTE Operator = "lte"
	OPERATOR_GTE Operator = "gte"
	OPERATOR_EQ Operator = "eq"
	OPERATOR_NEQ Operator = "neq"
	OPERATOR_IN Operator = "in"
	OPERATOR_NOT_IN Operator = "notIn"
)
