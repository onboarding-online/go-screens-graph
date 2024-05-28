package models

// DashesPageIndicatorBlock - Styles for dashes page indicator block
type DashesPageIndicatorBlock struct {

	NotFilledColor string `json:"notFilledColor,omitempty"`

	FilledColor string `json:"filledColor,omitempty"`

	DashHeight float32 `json:"dashHeight,omitempty"`

	DashesSpacing float32 `json:"dashesSpacing,omitempty"`
}
