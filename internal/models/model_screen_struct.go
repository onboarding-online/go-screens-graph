package models

// ScreenStruct - Screen struct
type ScreenStruct struct {

	NavigationBar PaywallNavigationBar `json:"navigationBar"`

	Footer PaywallFooter `json:"footer"`

	Styles ScreenBasicPaywallBlock `json:"styles"`

	Permission *RequestPermission `json:"permission"`

	Timer *ScreenTimer `json:"timer"`

	AnimationEnabled bool `json:"animationEnabled"`

	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`

	CustomScreenDescription string `json:"customScreenDescription"`

	// Dictionary of output labels for custom screen
	Labels map[string]CustomScreenOutputLabel `json:"labels"`

	// Dictionary of input values from custom screen
	Values map[string]CustomScreenInputValue `json:"values"`

	Callback Callback `json:"callback"`

	ImageTitleSubtitlesDescription string `json:"imageTitleSubtitlesDescription"`

	Image Image `json:"image"`

	Title Text `json:"title"`

	Subtitle1 Text `json:"subtitle1"`

	Subtitle2 Text `json:"subtitle2"`

	SliderDescription string `json:"sliderDescription"`

	Slider Slider `json:"slider"`

	MultipleSelectionDescription string `json:"multipleSelectionDescription"`

	Media Media `json:"media,omitempty"`

	Subtitle Text `json:"subtitle"`

	List RegularList `json:"list"`

	SingleSelectionDescription string `json:"singleSelectionDescription"`

	TitleSubtitleFieldDescription string `json:"titleSubtitleFieldDescription"`

	Field Field `json:"field"`

	ProgressBarTitleDescription string `json:"progressBarTitleDescription"`

	ProgressBar ProgressBar `json:"progressBar"`

	TwoColumnMultipleSelectionDescription string `json:"twoColumnMultipleSelectionDescription"`

	TwoColumnSingleSelectionDescription string `json:"twoColumnSingleSelectionDescription"`

	ImageTitleSubtitleListDescription string `json:"imageTitleSubtitleListDescription"`

	TitleSubtitleCalendarDescription string `json:"titleSubtitleCalendarDescription"`

	Calendar Calendar `json:"calendar"`

	TitleSubtitlePickerDescription string `json:"titleSubtitlePickerDescription"`

	Picker Picker `json:"picker"`

	ImageTitleSubtitlePickerDescription string `json:"imageTitleSubtitlePickerDescription"`

	ImageTitleSubtitleMultipleSelectionListDescription string `json:"imageTitleSubtitleMultipleSelectionListDescription"`

	// Defines the scale of the image in percentage
	ImageScale float32 `json:"imageScale,omitempty"`

	TooltipPermissionsDescription string `json:"tooltipPermissionsDescription"`

	Tooltip Tooltip `json:"tooltip"`

	ScreenBasicPaywall string `json:"screenBasicPaywall"`

	Divider Divider `json:"divider,omitempty"`

	Loader Loader `json:"loader,omitempty"`

	// Purchase flags
	Flags []PurchaseFlag `json:"flags"`

	Subscriptions SubscriptionList `json:"subscriptions"`

	CurrencyFormat CurrencyFormatKind `json:"currencyFormat"`
}
