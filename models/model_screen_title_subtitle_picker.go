/*
Onboarding online screens graph models

Onboarding online screens graph models and interfaces

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// ScreenTitleSubtitlePicker struct for ScreenTitleSubtitlePicker
type ScreenTitleSubtitlePicker struct {
	NavigationBar NavigationBar `json:"navigationBar"`
	Footer Footer `json:"footer"`
	Styles BasicScreenBlock `json:"styles"`
	Permission NullableRequestPermission `json:"permission"`
	Timer NullableScreenTimer `json:"timer"`
	AnimationEnabled bool `json:"animationEnabled"`
	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`
	TitleSubtitlePickerDescription string `json:"titleSubtitlePickerDescription"`
	Title Text `json:"title"`
	Subtitle Text `json:"subtitle"`
	Picker Picker `json:"picker"`
}

// NewScreenTitleSubtitlePicker instantiates a new ScreenTitleSubtitlePicker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenTitleSubtitlePicker(navigationBar NavigationBar, footer Footer, styles BasicScreenBlock, permission NullableRequestPermission, timer NullableScreenTimer, animationEnabled bool, useLocalAssetsIfAvailable bool, titleSubtitlePickerDescription string, title Text, subtitle Text, picker Picker) *ScreenTitleSubtitlePicker {
	this := ScreenTitleSubtitlePicker{}
	this.NavigationBar = navigationBar
	this.Footer = footer
	this.Styles = styles
	this.Permission = permission
	this.Timer = timer
	this.AnimationEnabled = animationEnabled
	this.UseLocalAssetsIfAvailable = useLocalAssetsIfAvailable
	this.TitleSubtitlePickerDescription = titleSubtitlePickerDescription
	this.Title = title
	this.Subtitle = subtitle
	this.Picker = picker
	return &this
}

// NewScreenTitleSubtitlePickerWithDefaults instantiates a new ScreenTitleSubtitlePicker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenTitleSubtitlePickerWithDefaults() *ScreenTitleSubtitlePicker {
	this := ScreenTitleSubtitlePicker{}
	return &this
}

// GetNavigationBar returns the NavigationBar field value
func (o *ScreenTitleSubtitlePicker) GetNavigationBar() NavigationBar {
	if o == nil {
		var ret NavigationBar
		return ret
	}

	return o.NavigationBar
}

// GetNavigationBarOk returns a tuple with the NavigationBar field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetNavigationBarOk() (*NavigationBar, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NavigationBar, true
}

// SetNavigationBar sets field value
func (o *ScreenTitleSubtitlePicker) SetNavigationBar(v NavigationBar) {
	o.NavigationBar = v
}

// GetFooter returns the Footer field value
func (o *ScreenTitleSubtitlePicker) GetFooter() Footer {
	if o == nil {
		var ret Footer
		return ret
	}

	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetFooterOk() (*Footer, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Footer, true
}

// SetFooter sets field value
func (o *ScreenTitleSubtitlePicker) SetFooter(v Footer) {
	o.Footer = v
}

// GetStyles returns the Styles field value
func (o *ScreenTitleSubtitlePicker) GetStyles() BasicScreenBlock {
	if o == nil {
		var ret BasicScreenBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetStylesOk() (*BasicScreenBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *ScreenTitleSubtitlePicker) SetStyles(v BasicScreenBlock) {
	o.Styles = v
}

// GetPermission returns the Permission field value
// If the value is explicit nil, the zero value for RequestPermission will be returned
func (o *ScreenTitleSubtitlePicker) GetPermission() RequestPermission {
	if o == nil || o.Permission.Get() == nil {
		var ret RequestPermission
		return ret
	}

	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenTitleSubtitlePicker) GetPermissionOk() (*RequestPermission, bool) {
	if o == nil {
    return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// SetPermission sets field value
func (o *ScreenTitleSubtitlePicker) SetPermission(v RequestPermission) {
	o.Permission.Set(&v)
}

// GetTimer returns the Timer field value
// If the value is explicit nil, the zero value for ScreenTimer will be returned
func (o *ScreenTitleSubtitlePicker) GetTimer() ScreenTimer {
	if o == nil || o.Timer.Get() == nil {
		var ret ScreenTimer
		return ret
	}

	return *o.Timer.Get()
}

// GetTimerOk returns a tuple with the Timer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenTitleSubtitlePicker) GetTimerOk() (*ScreenTimer, bool) {
	if o == nil {
    return nil, false
	}
	return o.Timer.Get(), o.Timer.IsSet()
}

// SetTimer sets field value
func (o *ScreenTitleSubtitlePicker) SetTimer(v ScreenTimer) {
	o.Timer.Set(&v)
}

// GetAnimationEnabled returns the AnimationEnabled field value
func (o *ScreenTitleSubtitlePicker) GetAnimationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AnimationEnabled
}

// GetAnimationEnabledOk returns a tuple with the AnimationEnabled field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetAnimationEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AnimationEnabled, true
}

// SetAnimationEnabled sets field value
func (o *ScreenTitleSubtitlePicker) SetAnimationEnabled(v bool) {
	o.AnimationEnabled = v
}

// GetUseLocalAssetsIfAvailable returns the UseLocalAssetsIfAvailable field value
func (o *ScreenTitleSubtitlePicker) GetUseLocalAssetsIfAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseLocalAssetsIfAvailable
}

// GetUseLocalAssetsIfAvailableOk returns a tuple with the UseLocalAssetsIfAvailable field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetUseLocalAssetsIfAvailableOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UseLocalAssetsIfAvailable, true
}

// SetUseLocalAssetsIfAvailable sets field value
func (o *ScreenTitleSubtitlePicker) SetUseLocalAssetsIfAvailable(v bool) {
	o.UseLocalAssetsIfAvailable = v
}

// GetTitleSubtitlePickerDescription returns the TitleSubtitlePickerDescription field value
func (o *ScreenTitleSubtitlePicker) GetTitleSubtitlePickerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TitleSubtitlePickerDescription
}

// GetTitleSubtitlePickerDescriptionOk returns a tuple with the TitleSubtitlePickerDescription field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetTitleSubtitlePickerDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TitleSubtitlePickerDescription, true
}

// SetTitleSubtitlePickerDescription sets field value
func (o *ScreenTitleSubtitlePicker) SetTitleSubtitlePickerDescription(v string) {
	o.TitleSubtitlePickerDescription = v
}

// GetTitle returns the Title field value
func (o *ScreenTitleSubtitlePicker) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScreenTitleSubtitlePicker) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *ScreenTitleSubtitlePicker) GetSubtitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetSubtitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *ScreenTitleSubtitlePicker) SetSubtitle(v Text) {
	o.Subtitle = v
}

// GetPicker returns the Picker field value
func (o *ScreenTitleSubtitlePicker) GetPicker() Picker {
	if o == nil {
		var ret Picker
		return ret
	}

	return o.Picker
}

// GetPickerOk returns a tuple with the Picker field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePicker) GetPickerOk() (*Picker, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Picker, true
}

// SetPicker sets field value
func (o *ScreenTitleSubtitlePicker) SetPicker(v Picker) {
	o.Picker = v
}

func (o ScreenTitleSubtitlePicker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["navigationBar"] = o.NavigationBar
	}
	if true {
		toSerialize["footer"] = o.Footer
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	if true {
		toSerialize["permission"] = o.Permission.Get()
	}
	if true {
		toSerialize["timer"] = o.Timer.Get()
	}
	if true {
		toSerialize["animationEnabled"] = o.AnimationEnabled
	}
	if true {
		toSerialize["useLocalAssetsIfAvailable"] = o.UseLocalAssetsIfAvailable
	}
	if true {
		toSerialize["titleSubtitlePickerDescription"] = o.TitleSubtitlePickerDescription
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["subtitle"] = o.Subtitle
	}
	if true {
		toSerialize["picker"] = o.Picker
	}
	return json.Marshal(toSerialize)
}

type NullableScreenTitleSubtitlePicker struct {
	value *ScreenTitleSubtitlePicker
	isSet bool
}

func (v NullableScreenTitleSubtitlePicker) Get() *ScreenTitleSubtitlePicker {
	return v.value
}

func (v *NullableScreenTitleSubtitlePicker) Set(val *ScreenTitleSubtitlePicker) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenTitleSubtitlePicker) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenTitleSubtitlePicker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenTitleSubtitlePicker(val *ScreenTitleSubtitlePicker) *NullableScreenTitleSubtitlePicker {
	return &NullableScreenTitleSubtitlePicker{value: val, isSet: true}
}

func (v NullableScreenTitleSubtitlePicker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenTitleSubtitlePicker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


