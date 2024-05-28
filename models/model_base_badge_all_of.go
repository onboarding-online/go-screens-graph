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

// BaseBadgeAllOf Base badge parameters
type BaseBadgeAllOf struct {
	Styles BadgeBlock `json:"styles"`
}

// NewBaseBadgeAllOf instantiates a new BaseBadgeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseBadgeAllOf(styles BadgeBlock) *BaseBadgeAllOf {
	this := BaseBadgeAllOf{}
	this.Styles = styles
	return &this
}

// NewBaseBadgeAllOfWithDefaults instantiates a new BaseBadgeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseBadgeAllOfWithDefaults() *BaseBadgeAllOf {
	this := BaseBadgeAllOf{}
	return &this
}

// GetStyles returns the Styles field value
func (o *BaseBadgeAllOf) GetStyles() BadgeBlock {
	if o == nil {
		var ret BadgeBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *BaseBadgeAllOf) GetStylesOk() (*BadgeBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *BaseBadgeAllOf) SetStyles(v BadgeBlock) {
	o.Styles = v
}

func (o BaseBadgeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableBaseBadgeAllOf struct {
	value *BaseBadgeAllOf
	isSet bool
}

func (v NullableBaseBadgeAllOf) Get() *BaseBadgeAllOf {
	return v.value
}

func (v *NullableBaseBadgeAllOf) Set(val *BaseBadgeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseBadgeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseBadgeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseBadgeAllOf(val *BaseBadgeAllOf) *NullableBaseBadgeAllOf {
	return &NullableBaseBadgeAllOf{value: val, isSet: true}
}

func (v NullableBaseBadgeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseBadgeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


