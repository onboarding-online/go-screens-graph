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

// ScreenBasicPaywall struct for ScreenBasicPaywall
type ScreenBasicPaywall struct {
	NavigationBar PaywallNavigationBar `json:"navigationBar"`
	Footer PaywallFooter `json:"footer"`
	Styles ScreenBasicPaywallBlock `json:"styles"`
	Permission NullableRequestPermission `json:"permission"`
	Timer NullableScreenTimer `json:"timer"`
	AnimationEnabled bool `json:"animationEnabled"`
	UseLocalAssetsIfAvailable bool `json:"useLocalAssetsIfAvailable"`
	ScreenBasicPaywall string `json:"screenBasicPaywall"`
	Title Text `json:"title"`
	Subtitle Text `json:"subtitle"`
	Divider *Divider `json:"divider,omitempty"`
	Media *Media `json:"media,omitempty"`
	List RegularList `json:"list"`
	Loader *Loader `json:"loader,omitempty"`
	// Purchase flags
	Flags []PurchaseFlag `json:"flags"`
	Subscriptions SubscriptionList `json:"subscriptions"`
	CurrencyFormat CurrencyFormatKind `json:"currencyFormat"`
}

// NewScreenBasicPaywall instantiates a new ScreenBasicPaywall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenBasicPaywall(navigationBar PaywallNavigationBar, footer PaywallFooter, styles ScreenBasicPaywallBlock, permission NullableRequestPermission, timer NullableScreenTimer, animationEnabled bool, useLocalAssetsIfAvailable bool, screenBasicPaywall string, title Text, subtitle Text, list RegularList, flags []PurchaseFlag, subscriptions SubscriptionList, currencyFormat CurrencyFormatKind) *ScreenBasicPaywall {
	this := ScreenBasicPaywall{}
	this.NavigationBar = navigationBar
	this.Footer = footer
	this.Styles = styles
	this.Permission = permission
	this.Timer = timer
	this.AnimationEnabled = animationEnabled
	this.UseLocalAssetsIfAvailable = useLocalAssetsIfAvailable
	this.ScreenBasicPaywall = screenBasicPaywall
	this.Title = title
	this.Subtitle = subtitle
	this.List = list
	this.Flags = flags
	this.Subscriptions = subscriptions
	this.CurrencyFormat = currencyFormat
	return &this
}

// NewScreenBasicPaywallWithDefaults instantiates a new ScreenBasicPaywall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenBasicPaywallWithDefaults() *ScreenBasicPaywall {
	this := ScreenBasicPaywall{}
	return &this
}

// GetNavigationBar returns the NavigationBar field value
func (o *ScreenBasicPaywall) GetNavigationBar() PaywallNavigationBar {
	if o == nil {
		var ret PaywallNavigationBar
		return ret
	}

	return o.NavigationBar
}

// GetNavigationBarOk returns a tuple with the NavigationBar field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetNavigationBarOk() (*PaywallNavigationBar, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NavigationBar, true
}

// SetNavigationBar sets field value
func (o *ScreenBasicPaywall) SetNavigationBar(v PaywallNavigationBar) {
	o.NavigationBar = v
}

// GetFooter returns the Footer field value
func (o *ScreenBasicPaywall) GetFooter() PaywallFooter {
	if o == nil {
		var ret PaywallFooter
		return ret
	}

	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetFooterOk() (*PaywallFooter, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Footer, true
}

// SetFooter sets field value
func (o *ScreenBasicPaywall) SetFooter(v PaywallFooter) {
	o.Footer = v
}

// GetStyles returns the Styles field value
func (o *ScreenBasicPaywall) GetStyles() ScreenBasicPaywallBlock {
	if o == nil {
		var ret ScreenBasicPaywallBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetStylesOk() (*ScreenBasicPaywallBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *ScreenBasicPaywall) SetStyles(v ScreenBasicPaywallBlock) {
	o.Styles = v
}

// GetPermission returns the Permission field value
// If the value is explicit nil, the zero value for RequestPermission will be returned
func (o *ScreenBasicPaywall) GetPermission() RequestPermission {
	if o == nil || o.Permission.Get() == nil {
		var ret RequestPermission
		return ret
	}

	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenBasicPaywall) GetPermissionOk() (*RequestPermission, bool) {
	if o == nil {
    return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// SetPermission sets field value
func (o *ScreenBasicPaywall) SetPermission(v RequestPermission) {
	o.Permission.Set(&v)
}

// GetTimer returns the Timer field value
// If the value is explicit nil, the zero value for ScreenTimer will be returned
func (o *ScreenBasicPaywall) GetTimer() ScreenTimer {
	if o == nil || o.Timer.Get() == nil {
		var ret ScreenTimer
		return ret
	}

	return *o.Timer.Get()
}

// GetTimerOk returns a tuple with the Timer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScreenBasicPaywall) GetTimerOk() (*ScreenTimer, bool) {
	if o == nil {
    return nil, false
	}
	return o.Timer.Get(), o.Timer.IsSet()
}

// SetTimer sets field value
func (o *ScreenBasicPaywall) SetTimer(v ScreenTimer) {
	o.Timer.Set(&v)
}

// GetAnimationEnabled returns the AnimationEnabled field value
func (o *ScreenBasicPaywall) GetAnimationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AnimationEnabled
}

// GetAnimationEnabledOk returns a tuple with the AnimationEnabled field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetAnimationEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AnimationEnabled, true
}

// SetAnimationEnabled sets field value
func (o *ScreenBasicPaywall) SetAnimationEnabled(v bool) {
	o.AnimationEnabled = v
}

// GetUseLocalAssetsIfAvailable returns the UseLocalAssetsIfAvailable field value
func (o *ScreenBasicPaywall) GetUseLocalAssetsIfAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseLocalAssetsIfAvailable
}

// GetUseLocalAssetsIfAvailableOk returns a tuple with the UseLocalAssetsIfAvailable field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetUseLocalAssetsIfAvailableOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UseLocalAssetsIfAvailable, true
}

// SetUseLocalAssetsIfAvailable sets field value
func (o *ScreenBasicPaywall) SetUseLocalAssetsIfAvailable(v bool) {
	o.UseLocalAssetsIfAvailable = v
}

// GetScreenBasicPaywall returns the ScreenBasicPaywall field value
func (o *ScreenBasicPaywall) GetScreenBasicPaywall() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScreenBasicPaywall
}

// GetScreenBasicPaywallOk returns a tuple with the ScreenBasicPaywall field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetScreenBasicPaywallOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScreenBasicPaywall, true
}

// SetScreenBasicPaywall sets field value
func (o *ScreenBasicPaywall) SetScreenBasicPaywall(v string) {
	o.ScreenBasicPaywall = v
}

// GetTitle returns the Title field value
func (o *ScreenBasicPaywall) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScreenBasicPaywall) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *ScreenBasicPaywall) GetSubtitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetSubtitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *ScreenBasicPaywall) SetSubtitle(v Text) {
	o.Subtitle = v
}

// GetDivider returns the Divider field value if set, zero value otherwise.
func (o *ScreenBasicPaywall) GetDivider() Divider {
	if o == nil || isNil(o.Divider) {
		var ret Divider
		return ret
	}
	return *o.Divider
}

// GetDividerOk returns a tuple with the Divider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetDividerOk() (*Divider, bool) {
	if o == nil || isNil(o.Divider) {
    return nil, false
	}
	return o.Divider, true
}

// HasDivider returns a boolean if a field has been set.
func (o *ScreenBasicPaywall) HasDivider() bool {
	if o != nil && !isNil(o.Divider) {
		return true
	}

	return false
}

// SetDivider gets a reference to the given Divider and assigns it to the Divider field.
func (o *ScreenBasicPaywall) SetDivider(v Divider) {
	o.Divider = &v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *ScreenBasicPaywall) GetMedia() Media {
	if o == nil || isNil(o.Media) {
		var ret Media
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetMediaOk() (*Media, bool) {
	if o == nil || isNil(o.Media) {
    return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *ScreenBasicPaywall) HasMedia() bool {
	if o != nil && !isNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given Media and assigns it to the Media field.
func (o *ScreenBasicPaywall) SetMedia(v Media) {
	o.Media = &v
}

// GetList returns the List field value
func (o *ScreenBasicPaywall) GetList() RegularList {
	if o == nil {
		var ret RegularList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetListOk() (*RegularList, bool) {
	if o == nil {
    return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *ScreenBasicPaywall) SetList(v RegularList) {
	o.List = v
}

// GetLoader returns the Loader field value if set, zero value otherwise.
func (o *ScreenBasicPaywall) GetLoader() Loader {
	if o == nil || isNil(o.Loader) {
		var ret Loader
		return ret
	}
	return *o.Loader
}

// GetLoaderOk returns a tuple with the Loader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetLoaderOk() (*Loader, bool) {
	if o == nil || isNil(o.Loader) {
    return nil, false
	}
	return o.Loader, true
}

// HasLoader returns a boolean if a field has been set.
func (o *ScreenBasicPaywall) HasLoader() bool {
	if o != nil && !isNil(o.Loader) {
		return true
	}

	return false
}

// SetLoader gets a reference to the given Loader and assigns it to the Loader field.
func (o *ScreenBasicPaywall) SetLoader(v Loader) {
	o.Loader = &v
}

// GetFlags returns the Flags field value
func (o *ScreenBasicPaywall) GetFlags() []PurchaseFlag {
	if o == nil {
		var ret []PurchaseFlag
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetFlagsOk() ([]PurchaseFlag, bool) {
	if o == nil {
    return nil, false
	}
	return o.Flags, true
}

// SetFlags sets field value
func (o *ScreenBasicPaywall) SetFlags(v []PurchaseFlag) {
	o.Flags = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *ScreenBasicPaywall) GetSubscriptions() SubscriptionList {
	if o == nil {
		var ret SubscriptionList
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetSubscriptionsOk() (*SubscriptionList, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ScreenBasicPaywall) SetSubscriptions(v SubscriptionList) {
	o.Subscriptions = v
}

// GetCurrencyFormat returns the CurrencyFormat field value
func (o *ScreenBasicPaywall) GetCurrencyFormat() CurrencyFormatKind {
	if o == nil {
		var ret CurrencyFormatKind
		return ret
	}

	return o.CurrencyFormat
}

// GetCurrencyFormatOk returns a tuple with the CurrencyFormat field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywall) GetCurrencyFormatOk() (*CurrencyFormatKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CurrencyFormat, true
}

// SetCurrencyFormat sets field value
func (o *ScreenBasicPaywall) SetCurrencyFormat(v CurrencyFormatKind) {
	o.CurrencyFormat = v
}

func (o ScreenBasicPaywall) MarshalJSON() ([]byte, error) {
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
		toSerialize["screenBasicPaywall"] = o.ScreenBasicPaywall
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !isNil(o.Divider) {
		toSerialize["divider"] = o.Divider
	}
	if !isNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if true {
		toSerialize["list"] = o.List
	}
	if !isNil(o.Loader) {
		toSerialize["loader"] = o.Loader
	}
	if true {
		toSerialize["flags"] = o.Flags
	}
	if true {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if true {
		toSerialize["currencyFormat"] = o.CurrencyFormat
	}
	return json.Marshal(toSerialize)
}

type NullableScreenBasicPaywall struct {
	value *ScreenBasicPaywall
	isSet bool
}

func (v NullableScreenBasicPaywall) Get() *ScreenBasicPaywall {
	return v.value
}

func (v *NullableScreenBasicPaywall) Set(val *ScreenBasicPaywall) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenBasicPaywall) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenBasicPaywall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenBasicPaywall(val *ScreenBasicPaywall) *NullableScreenBasicPaywall {
	return &NullableScreenBasicPaywall{value: val, isSet: true}
}

func (v NullableScreenBasicPaywall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenBasicPaywall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


