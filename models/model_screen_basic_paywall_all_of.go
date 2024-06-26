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

// ScreenBasicPaywallAllOf Screen basic paywall
type ScreenBasicPaywallAllOf struct {
	ScreenBasicPaywall string `json:"screenBasicPaywall"`
	NavigationBar *PaywallNavigationBar `json:"navigationBar,omitempty"`
	Footer *PaywallFooter `json:"footer,omitempty"`
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
	Styles ScreenBasicPaywallBlock `json:"styles"`
}

// NewScreenBasicPaywallAllOf instantiates a new ScreenBasicPaywallAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenBasicPaywallAllOf(screenBasicPaywall string, title Text, subtitle Text, list RegularList, flags []PurchaseFlag, subscriptions SubscriptionList, currencyFormat CurrencyFormatKind, styles ScreenBasicPaywallBlock) *ScreenBasicPaywallAllOf {
	this := ScreenBasicPaywallAllOf{}
	this.ScreenBasicPaywall = screenBasicPaywall
	this.Title = title
	this.Subtitle = subtitle
	this.List = list
	this.Flags = flags
	this.Subscriptions = subscriptions
	this.CurrencyFormat = currencyFormat
	this.Styles = styles
	return &this
}

// NewScreenBasicPaywallAllOfWithDefaults instantiates a new ScreenBasicPaywallAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenBasicPaywallAllOfWithDefaults() *ScreenBasicPaywallAllOf {
	this := ScreenBasicPaywallAllOf{}
	return &this
}

// GetScreenBasicPaywall returns the ScreenBasicPaywall field value
func (o *ScreenBasicPaywallAllOf) GetScreenBasicPaywall() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScreenBasicPaywall
}

// GetScreenBasicPaywallOk returns a tuple with the ScreenBasicPaywall field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetScreenBasicPaywallOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScreenBasicPaywall, true
}

// SetScreenBasicPaywall sets field value
func (o *ScreenBasicPaywallAllOf) SetScreenBasicPaywall(v string) {
	o.ScreenBasicPaywall = v
}

// GetNavigationBar returns the NavigationBar field value if set, zero value otherwise.
func (o *ScreenBasicPaywallAllOf) GetNavigationBar() PaywallNavigationBar {
	if o == nil || isNil(o.NavigationBar) {
		var ret PaywallNavigationBar
		return ret
	}
	return *o.NavigationBar
}

// GetNavigationBarOk returns a tuple with the NavigationBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetNavigationBarOk() (*PaywallNavigationBar, bool) {
	if o == nil || isNil(o.NavigationBar) {
    return nil, false
	}
	return o.NavigationBar, true
}

// HasNavigationBar returns a boolean if a field has been set.
func (o *ScreenBasicPaywallAllOf) HasNavigationBar() bool {
	if o != nil && !isNil(o.NavigationBar) {
		return true
	}

	return false
}

// SetNavigationBar gets a reference to the given PaywallNavigationBar and assigns it to the NavigationBar field.
func (o *ScreenBasicPaywallAllOf) SetNavigationBar(v PaywallNavigationBar) {
	o.NavigationBar = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *ScreenBasicPaywallAllOf) GetFooter() PaywallFooter {
	if o == nil || isNil(o.Footer) {
		var ret PaywallFooter
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetFooterOk() (*PaywallFooter, bool) {
	if o == nil || isNil(o.Footer) {
    return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *ScreenBasicPaywallAllOf) HasFooter() bool {
	if o != nil && !isNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given PaywallFooter and assigns it to the Footer field.
func (o *ScreenBasicPaywallAllOf) SetFooter(v PaywallFooter) {
	o.Footer = &v
}

// GetTitle returns the Title field value
func (o *ScreenBasicPaywallAllOf) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScreenBasicPaywallAllOf) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *ScreenBasicPaywallAllOf) GetSubtitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetSubtitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *ScreenBasicPaywallAllOf) SetSubtitle(v Text) {
	o.Subtitle = v
}

// GetDivider returns the Divider field value if set, zero value otherwise.
func (o *ScreenBasicPaywallAllOf) GetDivider() Divider {
	if o == nil || isNil(o.Divider) {
		var ret Divider
		return ret
	}
	return *o.Divider
}

// GetDividerOk returns a tuple with the Divider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetDividerOk() (*Divider, bool) {
	if o == nil || isNil(o.Divider) {
    return nil, false
	}
	return o.Divider, true
}

// HasDivider returns a boolean if a field has been set.
func (o *ScreenBasicPaywallAllOf) HasDivider() bool {
	if o != nil && !isNil(o.Divider) {
		return true
	}

	return false
}

// SetDivider gets a reference to the given Divider and assigns it to the Divider field.
func (o *ScreenBasicPaywallAllOf) SetDivider(v Divider) {
	o.Divider = &v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *ScreenBasicPaywallAllOf) GetMedia() Media {
	if o == nil || isNil(o.Media) {
		var ret Media
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetMediaOk() (*Media, bool) {
	if o == nil || isNil(o.Media) {
    return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *ScreenBasicPaywallAllOf) HasMedia() bool {
	if o != nil && !isNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given Media and assigns it to the Media field.
func (o *ScreenBasicPaywallAllOf) SetMedia(v Media) {
	o.Media = &v
}

// GetList returns the List field value
func (o *ScreenBasicPaywallAllOf) GetList() RegularList {
	if o == nil {
		var ret RegularList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetListOk() (*RegularList, bool) {
	if o == nil {
    return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *ScreenBasicPaywallAllOf) SetList(v RegularList) {
	o.List = v
}

// GetLoader returns the Loader field value if set, zero value otherwise.
func (o *ScreenBasicPaywallAllOf) GetLoader() Loader {
	if o == nil || isNil(o.Loader) {
		var ret Loader
		return ret
	}
	return *o.Loader
}

// GetLoaderOk returns a tuple with the Loader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetLoaderOk() (*Loader, bool) {
	if o == nil || isNil(o.Loader) {
    return nil, false
	}
	return o.Loader, true
}

// HasLoader returns a boolean if a field has been set.
func (o *ScreenBasicPaywallAllOf) HasLoader() bool {
	if o != nil && !isNil(o.Loader) {
		return true
	}

	return false
}

// SetLoader gets a reference to the given Loader and assigns it to the Loader field.
func (o *ScreenBasicPaywallAllOf) SetLoader(v Loader) {
	o.Loader = &v
}

// GetFlags returns the Flags field value
func (o *ScreenBasicPaywallAllOf) GetFlags() []PurchaseFlag {
	if o == nil {
		var ret []PurchaseFlag
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetFlagsOk() ([]PurchaseFlag, bool) {
	if o == nil {
    return nil, false
	}
	return o.Flags, true
}

// SetFlags sets field value
func (o *ScreenBasicPaywallAllOf) SetFlags(v []PurchaseFlag) {
	o.Flags = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *ScreenBasicPaywallAllOf) GetSubscriptions() SubscriptionList {
	if o == nil {
		var ret SubscriptionList
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetSubscriptionsOk() (*SubscriptionList, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ScreenBasicPaywallAllOf) SetSubscriptions(v SubscriptionList) {
	o.Subscriptions = v
}

// GetCurrencyFormat returns the CurrencyFormat field value
func (o *ScreenBasicPaywallAllOf) GetCurrencyFormat() CurrencyFormatKind {
	if o == nil {
		var ret CurrencyFormatKind
		return ret
	}

	return o.CurrencyFormat
}

// GetCurrencyFormatOk returns a tuple with the CurrencyFormat field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetCurrencyFormatOk() (*CurrencyFormatKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CurrencyFormat, true
}

// SetCurrencyFormat sets field value
func (o *ScreenBasicPaywallAllOf) SetCurrencyFormat(v CurrencyFormatKind) {
	o.CurrencyFormat = v
}

// GetStyles returns the Styles field value
func (o *ScreenBasicPaywallAllOf) GetStyles() ScreenBasicPaywallBlock {
	if o == nil {
		var ret ScreenBasicPaywallBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *ScreenBasicPaywallAllOf) GetStylesOk() (*ScreenBasicPaywallBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *ScreenBasicPaywallAllOf) SetStyles(v ScreenBasicPaywallBlock) {
	o.Styles = v
}

func (o ScreenBasicPaywallAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["screenBasicPaywall"] = o.ScreenBasicPaywall
	}
	if !isNil(o.NavigationBar) {
		toSerialize["navigationBar"] = o.NavigationBar
	}
	if !isNil(o.Footer) {
		toSerialize["footer"] = o.Footer
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
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableScreenBasicPaywallAllOf struct {
	value *ScreenBasicPaywallAllOf
	isSet bool
}

func (v NullableScreenBasicPaywallAllOf) Get() *ScreenBasicPaywallAllOf {
	return v.value
}

func (v *NullableScreenBasicPaywallAllOf) Set(val *ScreenBasicPaywallAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenBasicPaywallAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenBasicPaywallAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenBasicPaywallAllOf(val *ScreenBasicPaywallAllOf) *NullableScreenBasicPaywallAllOf {
	return &NullableScreenBasicPaywallAllOf{value: val, isSet: true}
}

func (v NullableScreenBasicPaywallAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenBasicPaywallAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


