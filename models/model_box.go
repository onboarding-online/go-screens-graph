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

// Box Container for elements
type Box struct {
	Styles BoxBlock `json:"styles"`
}

// NewBox instantiates a new Box object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBox(styles BoxBlock) *Box {
	this := Box{}
	this.Styles = styles
	return &this
}

// NewBoxWithDefaults instantiates a new Box object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoxWithDefaults() *Box {
	this := Box{}
	return &this
}

// GetStyles returns the Styles field value
func (o *Box) GetStyles() BoxBlock {
	if o == nil {
		var ret BoxBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *Box) GetStylesOk() (*BoxBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *Box) SetStyles(v BoxBlock) {
	o.Styles = v
}

func (o Box) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableBox struct {
	value *Box
	isSet bool
}

func (v NullableBox) Get() *Box {
	return v.value
}

func (v *NullableBox) Set(val *Box) {
	v.value = val
	v.isSet = true
}

func (v NullableBox) IsSet() bool {
	return v.isSet
}

func (v *NullableBox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBox(val *Box) *NullableBox {
	return &NullableBox{value: val, isSet: true}
}

func (v NullableBox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


