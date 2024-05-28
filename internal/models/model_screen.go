package models

type Screen struct {

	Id string `json:"id"`

	Name string `json:"name"`

	// Screen tags
	Tags []string `json:"tags"`

	ScreenType ScreenType `json:"screenType"`

	Struct ScreenStruct `json:"struct"`

	Position PositionOnConfigurator `json:"position"`
}
