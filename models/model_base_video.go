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

// BaseVideo Base video parameters
type BaseVideo struct {
	// Dictionary of localized Asset
	L10n map[string]Asset `json:"l10n"`
	Styles VideoBlock `json:"styles"`
}

// NewBaseVideo instantiates a new BaseVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseVideo(l10n map[string]Asset, styles VideoBlock) *BaseVideo {
	this := BaseVideo{}
	this.L10n = l10n
	this.Styles = styles
	return &this
}

// NewBaseVideoWithDefaults instantiates a new BaseVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseVideoWithDefaults() *BaseVideo {
	this := BaseVideo{}
	return &this
}

// GetL10n returns the L10n field value
func (o *BaseVideo) GetL10n() map[string]Asset {
	if o == nil {
		var ret map[string]Asset
		return ret
	}

	return o.L10n
}

// GetL10nOk returns a tuple with the L10n field value
// and a boolean to check if the value has been set.
func (o *BaseVideo) GetL10nOk() (*map[string]Asset, bool) {
	if o == nil {
    return nil, false
	}
	return &o.L10n, true
}

// SetL10n sets field value
func (o *BaseVideo) SetL10n(v map[string]Asset) {
	o.L10n = v
}

// GetStyles returns the Styles field value
func (o *BaseVideo) GetStyles() VideoBlock {
	if o == nil {
		var ret VideoBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *BaseVideo) GetStylesOk() (*VideoBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *BaseVideo) SetStyles(v VideoBlock) {
	o.Styles = v
}

func (o BaseVideo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["l10n"] = o.L10n
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableBaseVideo struct {
	value *BaseVideo
	isSet bool
}

func (v NullableBaseVideo) Get() *BaseVideo {
	return v.value
}

func (v *NullableBaseVideo) Set(val *BaseVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseVideo(val *BaseVideo) *NullableBaseVideo {
	return &NullableBaseVideo{value: val, isSet: true}
}

func (v NullableBaseVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


