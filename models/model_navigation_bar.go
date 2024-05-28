package models

// NavigationBar - Basic navigation bar
type NavigationBar struct {

	Back Button `json:"back,omitempty"`

	Skip Button `json:"skip,omitempty"`

	PageIndicator PageIndicator `json:"pageIndicator,omitempty"`

	PageIndicatorKind PageIndicatorKind `json:"pageIndicatorKind,omitempty"`

	DashesPageIndicator DashesPageIndicator `json:"dashesPageIndicator,omitempty"`
}
