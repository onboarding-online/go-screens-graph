package models

// CalendarDays - Calendar days config
type CalendarDays struct {

	Past CalendarDay `json:"past"`

	Today CalendarDay `json:"today"`

	Future CalendarDay `json:"future"`

	Weekend CalendarDay `json:"weekend"`
}
