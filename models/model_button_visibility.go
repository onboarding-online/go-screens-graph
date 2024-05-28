package models
// ButtonVisibility : Button visibility
type ButtonVisibility string

// List of ButtonVisibility
const (
	BUTTONVISIBILITY_DEFAULT ButtonVisibility = "default"
	BUTTONVISIBILITY_DISABLED_UNTIL_VALUE_ENTERED ButtonVisibility = "disabledUntilValueEntered"
	BUTTONVISIBILITY_HIDDEN_UNTIL_VALUE_ENTERED ButtonVisibility = "hiddenUntilValueEntered"
)
