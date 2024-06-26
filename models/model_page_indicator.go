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

// PageIndicator The line which shows where we are in range from 0 to 100%
type PageIndicator struct {
	// Progress value in percentage
	Value float32 `json:"value"`
	Styles PageIndicatorBlock `json:"styles"`
	Box Box `json:"box"`
}

// NewPageIndicator instantiates a new PageIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageIndicator(value float32, styles PageIndicatorBlock, box Box) *PageIndicator {
	this := PageIndicator{}
	this.Value = value
	this.Styles = styles
	this.Box = box
	return &this
}

// NewPageIndicatorWithDefaults instantiates a new PageIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageIndicatorWithDefaults() *PageIndicator {
	this := PageIndicator{}
	return &this
}

// GetValue returns the Value field value
func (o *PageIndicator) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PageIndicator) GetValueOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PageIndicator) SetValue(v float32) {
	o.Value = v
}

// GetStyles returns the Styles field value
func (o *PageIndicator) GetStyles() PageIndicatorBlock {
	if o == nil {
		var ret PageIndicatorBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *PageIndicator) GetStylesOk() (*PageIndicatorBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *PageIndicator) SetStyles(v PageIndicatorBlock) {
	o.Styles = v
}

// GetBox returns the Box field value
func (o *PageIndicator) GetBox() Box {
	if o == nil {
		var ret Box
		return ret
	}

	return o.Box
}

// GetBoxOk returns a tuple with the Box field value
// and a boolean to check if the value has been set.
func (o *PageIndicator) GetBoxOk() (*Box, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Box, true
}

// SetBox sets field value
func (o *PageIndicator) SetBox(v Box) {
	o.Box = v
}

func (o PageIndicator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	if true {
		toSerialize["box"] = o.Box
	}
	return json.Marshal(toSerialize)
}

type NullablePageIndicator struct {
	value *PageIndicator
	isSet bool
}

func (v NullablePageIndicator) Get() *PageIndicator {
	return v.value
}

func (v *NullablePageIndicator) Set(val *PageIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullablePageIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullablePageIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageIndicator(val *PageIndicator) *NullablePageIndicator {
	return &NullablePageIndicator{value: val, isSet: true}
}

func (v NullablePageIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


