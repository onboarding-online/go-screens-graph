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

// ProgressBarLabel Progress bar label element
type ProgressBarLabel struct {
	Styles LabelBlock `json:"styles"`
}

// NewProgressBarLabel instantiates a new ProgressBarLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgressBarLabel(styles LabelBlock) *ProgressBarLabel {
	this := ProgressBarLabel{}
	this.Styles = styles
	return &this
}

// NewProgressBarLabelWithDefaults instantiates a new ProgressBarLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgressBarLabelWithDefaults() *ProgressBarLabel {
	this := ProgressBarLabel{}
	return &this
}

// GetStyles returns the Styles field value
func (o *ProgressBarLabel) GetStyles() LabelBlock {
	if o == nil {
		var ret LabelBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *ProgressBarLabel) GetStylesOk() (*LabelBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *ProgressBarLabel) SetStyles(v LabelBlock) {
	o.Styles = v
}

func (o ProgressBarLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableProgressBarLabel struct {
	value *ProgressBarLabel
	isSet bool
}

func (v NullableProgressBarLabel) Get() *ProgressBarLabel {
	return v.value
}

func (v *NullableProgressBarLabel) Set(val *ProgressBarLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableProgressBarLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableProgressBarLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgressBarLabel(val *ProgressBarLabel) *NullableProgressBarLabel {
	return &NullableProgressBarLabel{value: val, isSet: true}
}

func (v NullableProgressBarLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgressBarLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


