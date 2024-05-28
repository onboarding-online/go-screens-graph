package models

// BaseCheckBox - Base checkbox parameters
type BaseCheckBox struct {

	Kind CheckBoxKind `json:"kind"`

	// Flag if checkbox checked
	Checked bool `json:"checked"`

	Styles CheckBoxBlock `json:"styles"`

	SelectedBlock SelectedCheckBoxBlock `json:"selectedBlock"`
}
