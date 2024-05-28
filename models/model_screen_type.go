package models
// ScreenType : The screen type
type ScreenType string

// List of ScreenType
const (
	SCREENTYPE_CUSTOM_SCREEN ScreenType = "CustomScreen"
	SCREENTYPE_SCREEN_IMAGE_TITLE_SUBTITLES ScreenType = "ScreenImageTitleSubtitles"
	SCREENTYPE_SCREEN_SLIDER ScreenType = "ScreenSlider"
	SCREENTYPE_SCREEN_TABLE_MULTIPLE_SELECTION ScreenType = "ScreenTableMultipleSelection"
	SCREENTYPE_SCREEN_TABLE_SINGLE_SELECTION ScreenType = "ScreenTableSingleSelection"
	SCREENTYPE_SCREEN_TITLE_SUBTITLE_FIELD ScreenType = "ScreenTitleSubtitleField"
	SCREENTYPE_SCREEN_PROGRESS_BAR_TITLE ScreenType = "ScreenProgressBarTitle"
	SCREENTYPE_SCREEN_TWO_COLUMN_MULTIPLE_SELECTION ScreenType = "ScreenTwoColumnMultipleSelection"
	SCREENTYPE_SCREEN_TWO_COLUMN_SINGLE_SELECTION ScreenType = "ScreenTwoColumnSingleSelection"
	SCREENTYPE_SCREEN_IMAGE_TITLE_SUBTITLE_LIST ScreenType = "ScreenImageTitleSubtitleList"
	SCREENTYPE_SCREEN_TITLE_SUBTITLE_CALENDAR ScreenType = "ScreenTitleSubtitleCalendar"
	SCREENTYPE_SCREEN_TITLE_SUBTITLE_PICKER ScreenType = "ScreenTitleSubtitlePicker"
	SCREENTYPE_SCREEN_IMAGE_TITLE_SUBTITLE_PICKER ScreenType = "ScreenImageTitleSubtitlePicker"
	SCREENTYPE_SCREEN_IMAGE_TITLE_SUBTITLE_MULTIPLE_SELECTION_LIST ScreenType = "ScreenImageTitleSubtitleMultipleSelectionList"
	SCREENTYPE_SCREEN_TOOLTIP_PERMISSIONS ScreenType = "ScreenTooltipPermissions"
	SCREENTYPE_SCREEN_BASIC_PAYWALL ScreenType = "ScreenBasicPaywall"
)
