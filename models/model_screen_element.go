package models
// ScreenElement : Enum of all main screen elements
type ScreenElement string

// List of ScreenElement
const (
	SCREENELEMENT_NAVIGATION_BAR ScreenElement = "NavigationBar"
	SCREENELEMENT_IMAGE ScreenElement = "Image"
	SCREENELEMENT_VIDEO ScreenElement = "Video"
	SCREENELEMENT_TITLE ScreenElement = "Title"
	SCREENELEMENT_SUBTITLE ScreenElement = "Subtitle"
	SCREENELEMENT_SUBTITLE1 ScreenElement = "Subtitle1"
	SCREENELEMENT_SUBTITLE2 ScreenElement = "Subtitle2"
	SCREENELEMENT_PICKER ScreenElement = "Picker"
	SCREENELEMENT_CALENDAR ScreenElement = "Calendar"
	SCREENELEMENT_FIELD ScreenElement = "Field"
	SCREENELEMENT_SINGLE_SELECTION_LIST ScreenElement = "SingleSelectionList"
	SCREENELEMENT_MULTIPLE_SELECTION_LIST ScreenElement = "MultipleSelectionList"
	SCREENELEMENT_TWO_COLUMN_SINGLE_SELECTION_LIST ScreenElement = "TwoColumnSingleSelectionList"
	SCREENELEMENT_TWO_COLUMN_MULTIPLE_SELECTION_LIST ScreenElement = "TwoColumnMultipleSelectionList"
	SCREENELEMENT_REGULAR_LIST ScreenElement = "RegularList"
	SCREENELEMENT_PROGRESS_BAR ScreenElement = "ProgressBar"
	SCREENELEMENT_SLIDER ScreenElement = "Slider"
	SCREENELEMENT_FOOTER ScreenElement = "Footer"
	SCREENELEMENT_TOOLTIP ScreenElement = "Tooltip"
)
