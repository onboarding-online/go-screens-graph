package models

// CustomScreenInputValue - Custom screen input value, allows user to send data to onboarding flow
type CustomScreenInputValue struct {

	Name string `json:"name"`

	Type ScreenInputType `json:"type"`

	Value string `json:"value"`
}
