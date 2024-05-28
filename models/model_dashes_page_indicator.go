package models

// DashesPageIndicator - Dashes page indicator
type DashesPageIndicator struct {

	NumberOfDashes float32 `json:"numberOfDashes"`

	Image Image `json:"image,omitempty"`

	Title Text `json:"title,omitempty"`

	NumOption1 float32 `json:"numOption1,omitempty"`

	NumOption2 float32 `json:"numOption2,omitempty"`

	StrOption1 string `json:"strOption1,omitempty"`

	StrOption2 string `json:"strOption2,omitempty"`

	Styles DashesPageIndicatorBlock `json:"styles"`
}
