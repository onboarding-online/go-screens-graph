package models

// CustomScreenAllOf - Custom screen
type CustomScreenAllOf struct {

	CustomScreenDescription string `json:"customScreenDescription"`

	// Dictionary of output labels for custom screen
	Labels map[string]CustomScreenOutputLabel `json:"labels"`

	// Dictionary of input values from custom screen
	Values map[string]CustomScreenInputValue `json:"values"`

	Callback Callback `json:"callback"`
}
