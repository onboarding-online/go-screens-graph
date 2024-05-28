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

// ScreenTwoColumnMultipleSelection struct for ScreenTwoColumnMultipleSelection
type ScreenTwoColumnMultipleSelection struct {
	NavigationBar NavigationBar `json:"navigationBar"`
	Footer Footer `json:"footer"`
	Styles BasicScreenBlock `json:"styles"`
	Permission NullableRequestPermission `json:"permission"`
	Timer NullableScreenTimer `json:"timer"`
	AnimationEnabled bool `json:"animationEnabled"`
	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`
	TwoColumnMultipleSelectionDescription string `json:"twoColumnMultipleSelectionDescription"`
	Media *Media `json:"media,omitempty"`
	Title Text `json:"title"`
	Subtitle Text `json:"subtitle"`
	List TwoColumnMultipleSelectionList `json:"list"`
}

// NewScreenTwoColumnMultipleSelection instantiates a new ScreenTwoColumnMultipleSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenTwoColumnMultipleSelection(navigationBar NavigationBar, footer Footer, styles BasicScreenBlock, permission NullableRequestPermission, timer NullableScreenTimer, animationEnabled bool, useLocalAssetsIfAvailable bool, twoColumnMultipleSelectionDescription string, title Text, subtitle Text, list TwoColumnMultipleSelectionList) *ScreenTwoColumnMultipleSelection {
	this := ScreenTwoColumnMultipleSelection{}
	this.NavigationBar = navigationBar
	this.Footer = footer
	this.Styles = styles
	this.Permission = permission
	this.Timer = timer
	this.AnimationEnabled = animationEnabled
	this.UseLocalAssetsIfAvailable = useLocalAssetsIfAvailable
	this.TwoColumnMultipleSelectionDescription = twoColumnMultipleSelectionDescription
	this.Title = title
	this.Subtitle = subtitle
	this.List = list
	return &this
}

// NewScreenTwoColumnMultipleSelectionWithDefaults instantiates a new ScreenTwoColumnMultipleSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenTwoColumnMultipleSelectionWithDefaults() *ScreenTwoColumnMultipleSelection {
	this := ScreenTwoColumnMultipleSelection{}
	return &this
}

// GetNavigationBar returns the NavigationBar field value
func (o *ScreenTwoColumnMultipleSelection) GetNavigationBar() NavigationBar {
	if o == nil {
		var ret NavigationBar
		return ret
	}

	return o.NavigationBar
}

// GetNavigationBarOk returns a tuple with the NavigationBar field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetNavigationBarOk() (*NavigationBar, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NavigationBar, true
}

// SetNavigationBar sets field value
func (o *ScreenTwoColumnMultipleSelection) SetNavigationBar(v NavigationBar) {
	o.NavigationBar = v
}

// GetFooter returns the Footer field value
func (o *ScreenTwoColumnMultipleSelection) GetFooter() Footer {
	if o == nil {
		var ret Footer
		return ret
	}

	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetFooterOk() (*Footer, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Footer, true
}

// SetFooter sets field value
func (o *ScreenTwoColumnMultipleSelection) SetFooter(v Footer) {
	o.Footer = v
}

// GetStyles returns the Styles field value
func (o *ScreenTwoColumnMultipleSelection) GetStyles() BasicScreenBlock {
	if o == nil {
		var ret BasicScreenBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetStylesOk() (*BasicScreenBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *ScreenTwoColumnMultipleSelection) SetStyles(v BasicScreenBlock) {
	o.Styles = v
}

// GetPermission returns the Permission field value
// If the value is explicit nil, the zero value for RequestPermission will be returned
func (o *ScreenTwoColumnMultipleSelection) GetPermission() RequestPermission {
	if o == nil || o.Permission.Get() == nil {
		var ret RequestPermission
		return ret
	}

	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenTwoColumnMultipleSelection) GetPermissionOk() (*RequestPermission, bool) {
	if o == nil {
    return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// SetPermission sets field value
func (o *ScreenTwoColumnMultipleSelection) SetPermission(v RequestPermission) {
	o.Permission.Set(&v)
}

// GetTimer returns the Timer field value
// If the value is explicit nil, the zero value for ScreenTimer will be returned
func (o *ScreenTwoColumnMultipleSelection) GetTimer() ScreenTimer {
	if o == nil || o.Timer.Get() == nil {
		var ret ScreenTimer
		return ret
	}

	return *o.Timer.Get()
}

// GetTimerOk returns a tuple with the Timer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenTwoColumnMultipleSelection) GetTimerOk() (*ScreenTimer, bool) {
	if o == nil {
    return nil, false
	}
	return o.Timer.Get(), o.Timer.IsSet()
}

// SetTimer sets field value
func (o *ScreenTwoColumnMultipleSelection) SetTimer(v ScreenTimer) {
	o.Timer.Set(&v)
}

// GetAnimationEnabled returns the AnimationEnabled field value
func (o *ScreenTwoColumnMultipleSelection) GetAnimationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AnimationEnabled
}

// GetAnimationEnabledOk returns a tuple with the AnimationEnabled field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetAnimationEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AnimationEnabled, true
}

// SetAnimationEnabled sets field value
func (o *ScreenTwoColumnMultipleSelection) SetAnimationEnabled(v bool) {
	o.AnimationEnabled = v
}

// GetUseLocalAssetsIfAvailable returns the UseLocalAssetsIfAvailable field value
func (o *ScreenTwoColumnMultipleSelection) GetUseLocalAssetsIfAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseLocalAssetsIfAvailable
}

// GetUseLocalAssetsIfAvailableOk returns a tuple with the UseLocalAssetsIfAvailable field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetUseLocalAssetsIfAvailableOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UseLocalAssetsIfAvailable, true
}

// SetUseLocalAssetsIfAvailable sets field value
func (o *ScreenTwoColumnMultipleSelection) SetUseLocalAssetsIfAvailable(v bool) {
	o.UseLocalAssetsIfAvailable = v
}

// GetTwoColumnMultipleSelectionDescription returns the TwoColumnMultipleSelectionDescription field value
func (o *ScreenTwoColumnMultipleSelection) GetTwoColumnMultipleSelectionDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TwoColumnMultipleSelectionDescription
}

// GetTwoColumnMultipleSelectionDescriptionOk returns a tuple with the TwoColumnMultipleSelectionDescription field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetTwoColumnMultipleSelectionDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TwoColumnMultipleSelectionDescription, true
}

// SetTwoColumnMultipleSelectionDescription sets field value
func (o *ScreenTwoColumnMultipleSelection) SetTwoColumnMultipleSelectionDescription(v string) {
	o.TwoColumnMultipleSelectionDescription = v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *ScreenTwoColumnMultipleSelection) GetMedia() Media {
	if o == nil || isNil(o.Media) {
		var ret Media
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetMediaOk() (*Media, bool) {
	if o == nil || isNil(o.Media) {
    return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *ScreenTwoColumnMultipleSelection) HasMedia() bool {
	if o != nil && !isNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given Media and assigns it to the Media field.
func (o *ScreenTwoColumnMultipleSelection) SetMedia(v Media) {
	o.Media = &v
}

// GetTitle returns the Title field value
func (o *ScreenTwoColumnMultipleSelection) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScreenTwoColumnMultipleSelection) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *ScreenTwoColumnMultipleSelection) GetSubtitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetSubtitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *ScreenTwoColumnMultipleSelection) SetSubtitle(v Text) {
	o.Subtitle = v
}

// GetList returns the List field value
func (o *ScreenTwoColumnMultipleSelection) GetList() TwoColumnMultipleSelectionList {
	if o == nil {
		var ret TwoColumnMultipleSelectionList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ScreenTwoColumnMultipleSelection) GetListOk() (*TwoColumnMultipleSelectionList, bool) {
	if o == nil {
    return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *ScreenTwoColumnMultipleSelection) SetList(v TwoColumnMultipleSelectionList) {
	o.List = v
}

func (o ScreenTwoColumnMultipleSelection) MarshalJSON() ([]byte, error) {
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
		toSerialize["twoColumnMultipleSelectionDescription"] = o.TwoColumnMultipleSelectionDescription
	}
	if !isNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["subtitle"] = o.Subtitle
	}
	if true {
		toSerialize["list"] = o.List
	}
	return json.Marshal(toSerialize)
}

type NullableScreenTwoColumnMultipleSelection struct {
	value *ScreenTwoColumnMultipleSelection
	isSet bool
}

func (v NullableScreenTwoColumnMultipleSelection) Get() *ScreenTwoColumnMultipleSelection {
	return v.value
}

func (v *NullableScreenTwoColumnMultipleSelection) Set(val *ScreenTwoColumnMultipleSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenTwoColumnMultipleSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenTwoColumnMultipleSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenTwoColumnMultipleSelection(val *ScreenTwoColumnMultipleSelection) *NullableScreenTwoColumnMultipleSelection {
	return &NullableScreenTwoColumnMultipleSelection{value: val, isSet: true}
}

func (v NullableScreenTwoColumnMultipleSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenTwoColumnMultipleSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


