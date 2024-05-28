package models

// BaseLoader - Base loader parameters
type BaseLoader struct {

	Kind LoaderKind `json:"kind"`

	Styles LoaderBlock `json:"styles"`
}
