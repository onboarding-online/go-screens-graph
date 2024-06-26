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

// ScreenImageTitleSubtitleList struct for ScreenImageTitleSubtitleList
type ScreenImageTitleSubtitleList struct {
	NavigationBar NavigationBar `json:"navigationBar"`
	Footer Footer `json:"footer"`
	Styles BasicScreenBlock `json:"styles"`
	Permission NullableRequestPermission `json:"permission"`
	Timer NullableScreenTimer `json:"timer"`
	AnimationEnabled bool `json:"animationEnabled"`
	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`
	ImageTitleSubtitleListDescription string `json:"imageTitleSubtitleListDescription"`
	Image Image `json:"image"`
	Title Text `json:"title"`
	Subtitle Text `json:"subtitle"`
	List RegularList `json:"list"`
}

// NewScreenImageTitleSubtitleList instantiates a new ScreenImageTitleSubtitleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenImageTitleSubtitleList(navigationBar NavigationBar, footer Footer, styles BasicScreenBlock, permission NullableRequestPermission, timer NullableScreenTimer, animationEnabled bool, useLocalAssetsIfAvailable bool, imageTitleSubtitleListDescription string, image Image, title Text, subtitle Text, list RegularList) *ScreenImageTitleSubtitleList {
	this := ScreenImageTitleSubtitleList{}
	this.NavigationBar = navigationBar
	this.Footer = footer
	this.Styles = styles
	this.Permission = permission
	this.Timer = timer
	this.AnimationEnabled = animationEnabled
	this.UseLocalAssetsIfAvailable = useLocalAssetsIfAvailable
	this.ImageTitleSubtitleListDescription = imageTitleSubtitleListDescription
	this.Image = image
	this.Title = title
	this.Subtitle = subtitle
	this.List = list
	return &this
}

// NewScreenImageTitleSubtitleListWithDefaults instantiates a new ScreenImageTitleSubtitleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenImageTitleSubtitleListWithDefaults() *ScreenImageTitleSubtitleList {
	this := ScreenImageTitleSubtitleList{}
	return &this
}

// GetNavigationBar returns the NavigationBar field value
func (o *ScreenImageTitleSubtitleList) GetNavigationBar() NavigationBar {
	if o == nil {
		var ret NavigationBar
		return ret
	}

	return o.NavigationBar
}

// GetNavigationBarOk returns a tuple with the NavigationBar field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetNavigationBarOk() (*NavigationBar, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NavigationBar, true
}

// SetNavigationBar sets field value
func (o *ScreenImageTitleSubtitleList) SetNavigationBar(v NavigationBar) {
	o.NavigationBar = v
}

// GetFooter returns the Footer field value
func (o *ScreenImageTitleSubtitleList) GetFooter() Footer {
	if o == nil {
		var ret Footer
		return ret
	}

	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetFooterOk() (*Footer, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Footer, true
}

// SetFooter sets field value
func (o *ScreenImageTitleSubtitleList) SetFooter(v Footer) {
	o.Footer = v
}

// GetStyles returns the Styles field value
func (o *ScreenImageTitleSubtitleList) GetStyles() BasicScreenBlock {
	if o == nil {
		var ret BasicScreenBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetStylesOk() (*BasicScreenBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *ScreenImageTitleSubtitleList) SetStyles(v BasicScreenBlock) {
	o.Styles = v
}

// GetPermission returns the Permission field value
// If the value is explicit nil, the zero value for RequestPermission will be returned
func (o *ScreenImageTitleSubtitleList) GetPermission() RequestPermission {
	if o == nil || o.Permission.Get() == nil {
		var ret RequestPermission
		return ret
	}

	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenImageTitleSubtitleList) GetPermissionOk() (*RequestPermission, bool) {
	if o == nil {
    return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// SetPermission sets field value
func (o *ScreenImageTitleSubtitleList) SetPermission(v RequestPermission) {
	o.Permission.Set(&v)
}

// GetTimer returns the Timer field value
// If the value is explicit nil, the zero value for ScreenTimer will be returned
func (o *ScreenImageTitleSubtitleList) GetTimer() ScreenTimer {
	if o == nil || o.Timer.Get() == nil {
		var ret ScreenTimer
		return ret
	}

	return *o.Timer.Get()
}

// GetTimerOk returns a tuple with the Timer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenImageTitleSubtitleList) GetTimerOk() (*ScreenTimer, bool) {
	if o == nil {
    return nil, false
	}
	return o.Timer.Get(), o.Timer.IsSet()
}

// SetTimer sets field value
func (o *ScreenImageTitleSubtitleList) SetTimer(v ScreenTimer) {
	o.Timer.Set(&v)
}

// GetAnimationEnabled returns the AnimationEnabled field value
func (o *ScreenImageTitleSubtitleList) GetAnimationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AnimationEnabled
}

// GetAnimationEnabledOk returns a tuple with the AnimationEnabled field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetAnimationEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AnimationEnabled, true
}

// SetAnimationEnabled sets field value
func (o *ScreenImageTitleSubtitleList) SetAnimationEnabled(v bool) {
	o.AnimationEnabled = v
}

// GetUseLocalAssetsIfAvailable returns the UseLocalAssetsIfAvailable field value
func (o *ScreenImageTitleSubtitleList) GetUseLocalAssetsIfAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseLocalAssetsIfAvailable
}

// GetUseLocalAssetsIfAvailableOk returns a tuple with the UseLocalAssetsIfAvailable field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetUseLocalAssetsIfAvailableOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UseLocalAssetsIfAvailable, true
}

// SetUseLocalAssetsIfAvailable sets field value
func (o *ScreenImageTitleSubtitleList) SetUseLocalAssetsIfAvailable(v bool) {
	o.UseLocalAssetsIfAvailable = v
}

// GetImageTitleSubtitleListDescription returns the ImageTitleSubtitleListDescription field value
func (o *ScreenImageTitleSubtitleList) GetImageTitleSubtitleListDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTitleSubtitleListDescription
}

// GetImageTitleSubtitleListDescriptionOk returns a tuple with the ImageTitleSubtitleListDescription field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetImageTitleSubtitleListDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ImageTitleSubtitleListDescription, true
}

// SetImageTitleSubtitleListDescription sets field value
func (o *ScreenImageTitleSubtitleList) SetImageTitleSubtitleListDescription(v string) {
	o.ImageTitleSubtitleListDescription = v
}

// GetImage returns the Image field value
func (o *ScreenImageTitleSubtitleList) GetImage() Image {
	if o == nil {
		var ret Image
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetImageOk() (*Image, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ScreenImageTitleSubtitleList) SetImage(v Image) {
	o.Image = v
}

// GetTitle returns the Title field value
func (o *ScreenImageTitleSubtitleList) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScreenImageTitleSubtitleList) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *ScreenImageTitleSubtitleList) GetSubtitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetSubtitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *ScreenImageTitleSubtitleList) SetSubtitle(v Text) {
	o.Subtitle = v
}

// GetList returns the List field value
func (o *ScreenImageTitleSubtitleList) GetList() RegularList {
	if o == nil {
		var ret RegularList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ScreenImageTitleSubtitleList) GetListOk() (*RegularList, bool) {
	if o == nil {
    return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *ScreenImageTitleSubtitleList) SetList(v RegularList) {
	o.List = v
}

func (o ScreenImageTitleSubtitleList) MarshalJSON() ([]byte, error) {
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
		toSerialize["imageTitleSubtitleListDescription"] = o.ImageTitleSubtitleListDescription
	}
	if true {
		toSerialize["image"] = o.Image
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

type NullableScreenImageTitleSubtitleList struct {
	value *ScreenImageTitleSubtitleList
	isSet bool
}

func (v NullableScreenImageTitleSubtitleList) Get() *ScreenImageTitleSubtitleList {
	return v.value
}

func (v *NullableScreenImageTitleSubtitleList) Set(val *ScreenImageTitleSubtitleList) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenImageTitleSubtitleList) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenImageTitleSubtitleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenImageTitleSubtitleList(val *ScreenImageTitleSubtitleList) *NullableScreenImageTitleSubtitleList {
	return &NullableScreenImageTitleSubtitleList{value: val, isSet: true}
}

func (v NullableScreenImageTitleSubtitleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenImageTitleSubtitleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


