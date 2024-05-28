package models

// BasicPaywallBodyStyle - Base button content
type BasicPaywallBodyStyle struct {

	MaskKind PaywallBodyMaskKind `json:"maskKind,omitempty"`

	// List of colors
	Colors []string `json:"colors"`

	// List of locations
	Locations []float32 `json:"locations,omitempty"`

	StartPoint XyPoint `json:"startPoint"`

	EndPoint XyPoint `json:"endPoint"`

	GradientKind GradientKind `json:"gradientKind,omitempty"`

	// Height in percentage
	HeightPercentage float32 `json:"heightPercentage,omitempty"`
}
