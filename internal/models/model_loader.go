package models

// Loader - Loader parameters
type Loader struct {

	Box Box `json:"box"`

	Kind LoaderKind `json:"kind"`

	Styles LoaderBlock `json:"styles"`
}
