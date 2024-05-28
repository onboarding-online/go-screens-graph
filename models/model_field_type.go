package models
// FieldType : Possible field types
type FieldType string

// List of FieldType
const (
	FIELDTYPE_STRING FieldType = "string"
	FIELDTYPE_INT FieldType = "int"
	FIELDTYPE_DOUBLE FieldType = "double"
	FIELDTYPE_DATE FieldType = "date"
)
