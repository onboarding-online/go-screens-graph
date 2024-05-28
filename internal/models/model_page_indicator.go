package models

// PageIndicator - The line which shows where we are in range from 0 to 100%
type PageIndicator struct {

	// Progress value in percentage
	Value float32 `json:"value"`

	Styles PageIndicatorBlock `json:"styles"`

	Box Box `json:"box"`
}
