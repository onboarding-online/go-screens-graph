package models

// Calendar - Calendar
type Calendar struct {

	Box Box `json:"box"`

	Days CalendarDays `json:"days"`

	Styles CalendarBlock `json:"styles"`

	Values []string `json:"values"`
}
