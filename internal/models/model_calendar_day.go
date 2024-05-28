package models

// CalendarDay - Calendar day config
type CalendarDay struct {

	LabelBlock LabelBlock `json:"labelBlock"`

	SelectedLabelBlock LabelBlock `json:"selectedLabelBlock"`

	CheckBox BaseCheckBox `json:"checkBox"`
}
