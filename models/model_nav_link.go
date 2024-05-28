package models

// NavLink - Navigation link parameters
type NavLink struct {

	Box Box `json:"box"`

	Kind NavLinkKind `json:"kind"`

	Content BaseNavLinkContent `json:"content"`

	Styles NavLinkBlock `json:"styles"`

	// The uri Link
	Uri string `json:"uri"`
}
