package models

// PaywallNavigationBarBlock - Styles for paywall navigation bar
type PaywallNavigationBarBlock struct {

	RestoreHorizontalAlignment RestoreHorizontalAlignment `json:"restoreHorizontalAlignment,omitempty"`

	CloseHorizontalAlignment CloseHorizontalAlignment `json:"closeHorizontalAlignment,omitempty"`

	CloseAppearance CloseAppearance `json:"closeAppearance,omitempty"`

	// Close button appearance visible after timer value in seconds
	CloseVisibleAfterTimerValue float32 `json:"closeVisibleAfterTimerValue,omitempty"`
}
