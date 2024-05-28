package models

// Footer - Footer
type Footer struct {

	Button1 Button `json:"button1,omitempty"`

	Button2 Button `json:"button2,omitempty"`

	Kind BasicFooterKind `json:"kind,omitempty"`

	Styles BasicFooterBlock `json:"styles,omitempty"`
}
