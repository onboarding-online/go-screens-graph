package models

// CalendarBlock - Styles for calendar
type CalendarBlock struct {

	// Background of calendar
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Background of calendar header
	HeaderBackgroundColor string `json:"headerBackgroundColor,omitempty"`

	HeaderLabel LabelBlock `json:"headerLabel"`

	HeaderWeekendLabel LabelBlock `json:"headerWeekendLabel"`

	TodayLabel LabelBlock `json:"todayLabel"`

	MonthLabel LabelBlock `json:"monthLabel"`

	CurrentMonthLabel LabelBlock `json:"currentMonthLabel"`
}
