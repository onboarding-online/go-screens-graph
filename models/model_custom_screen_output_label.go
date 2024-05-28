package models

// CustomScreenOutputLabel - Custom screen output label, allows user to show this data on mobile
type CustomScreenOutputLabel struct {

	Name string `json:"name"`

	Value BaseText `json:"value"`
}
