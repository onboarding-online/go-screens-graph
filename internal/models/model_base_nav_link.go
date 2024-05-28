package models

// BaseNavLink - Base navigation link parameters
type BaseNavLink struct {

	Kind NavLinkKind `json:"kind"`

	Content BaseNavLinkContent `json:"content"`

	Styles NavLinkBlock `json:"styles"`

	// The uri Link
	Uri string `json:"uri"`
}
