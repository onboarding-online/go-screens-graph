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

// SelectedListItemBlock Selected list item
type SelectedListItemBlock struct {
	Styles ListBlock `json:"styles"`
}

// NewSelectedListItemBlock instantiates a new SelectedListItemBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectedListItemBlock(styles ListBlock) *SelectedListItemBlock {
	this := SelectedListItemBlock{}
	this.Styles = styles
	return &this
}

// NewSelectedListItemBlockWithDefaults instantiates a new SelectedListItemBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectedListItemBlockWithDefaults() *SelectedListItemBlock {
	this := SelectedListItemBlock{}
	return &this
}

// GetStyles returns the Styles field value
func (o *SelectedListItemBlock) GetStyles() ListBlock {
	if o == nil {
		var ret ListBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *SelectedListItemBlock) GetStylesOk() (*ListBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *SelectedListItemBlock) SetStyles(v ListBlock) {
	o.Styles = v
}

func (o SelectedListItemBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableSelectedListItemBlock struct {
	value *SelectedListItemBlock
	isSet bool
}

func (v NullableSelectedListItemBlock) Get() *SelectedListItemBlock {
	return v.value
}

func (v *NullableSelectedListItemBlock) Set(val *SelectedListItemBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedListItemBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedListItemBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedListItemBlock(val *SelectedListItemBlock) *NullableSelectedListItemBlock {
	return &NullableSelectedListItemBlock{value: val, isSet: true}
}

func (v NullableSelectedListItemBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedListItemBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


