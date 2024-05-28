package models

// FieldBlock - Styles for field
type FieldBlock struct {

	// Background color for field
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Border color for field
	BorderColor string `json:"borderColor,omitempty"`

	// Entered text color for field
	EnteredTextColor string `json:"enteredTextColor,omitempty"`
}
