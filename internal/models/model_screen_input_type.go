package models
// ScreenInputType : Possible data types for input on screen
type ScreenInputType string

// List of ScreenInputType
const (
	SCREENINPUTTYPE_STRING ScreenInputType = "string"
	SCREENINPUTTYPE_INT ScreenInputType = "int"
	SCREENINPUTTYPE_DOUBLE ScreenInputType = "double"
	SCREENINPUTTYPE_DATE ScreenInputType = "date"
	SCREENINPUTTYPE_LIST_STRING ScreenInputType = "listString"
	SCREENINPUTTYPE_LIST_INT ScreenInputType = "listInt"
	SCREENINPUTTYPE_LIST_DOUBLE ScreenInputType = "listDouble"
	SCREENINPUTTYPE_LIST_DATE ScreenInputType = "listDate"
)
