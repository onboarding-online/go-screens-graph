package models

// ProgressBar - Progress bar element
type ProgressBar struct {

	Kind ProgressBarKind `json:"kind"`

	Timer ElementTimer `json:"timer"`

	// Sections for progress bar
	Items []ProgressBarItem `json:"items"`

	Styles ProgressBarBlock `json:"styles"`

	Label ProgressBarLabel `json:"label"`

	Box Box `json:"box"`
}
