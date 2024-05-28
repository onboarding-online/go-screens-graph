package models

// CheckBox - Checkbox parameters
type CheckBox struct {

	Box Box `json:"box"`

	Kind CheckBoxKind `json:"kind"`

	// Flag if checkbox checked
	Checked bool `json:"checked"`

	Styles CheckBoxBlock `json:"styles"`

	SelectedBlock SelectedCheckBoxBlock `json:"selectedBlock"`
}
