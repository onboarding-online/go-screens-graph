package models

// ProgressBarItem - Information about one section of progress bar
type ProgressBarItem struct {

	// Section display start
	ValueFrom int32 `json:"valueFrom"`

	// Section display end
	ValueTo int32 `json:"valueTo"`

	Timer ElementTimer `json:"timer,omitempty"`

	Content ProgressBarItemContent `json:"content"`

	Styles ProgressBarItemBlock `json:"styles,omitempty"`
}
