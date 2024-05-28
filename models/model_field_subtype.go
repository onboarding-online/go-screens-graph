package models
// FieldSubtype : Possible field subtype
type FieldSubtype string

// List of FieldSubtype
const (
	FIELDSUBTYPE_EMAIL FieldSubtype = "Email"
	FIELDSUBTYPE_BIRTH_DATE FieldSubtype = "BirthDate"
	FIELDSUBTYPE_HEIGHT FieldSubtype = "Height"
	FIELDSUBTYPE_WEIGHT FieldSubtype = "Weight"
)
