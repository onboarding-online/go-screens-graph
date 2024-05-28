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

// BackgroundStyleColor Background color for screen
type BackgroundStyleColor struct {
	Color string `json:"color"`
}

// NewBackgroundStyleColor instantiates a new BackgroundStyleColor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundStyleColor(color string) *BackgroundStyleColor {
	this := BackgroundStyleColor{}
	this.Color = color
	return &this
}

// NewBackgroundStyleColorWithDefaults instantiates a new BackgroundStyleColor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundStyleColorWithDefaults() *BackgroundStyleColor {
	this := BackgroundStyleColor{}
	return &this
}

// GetColor returns the Color field value
func (o *BackgroundStyleColor) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *BackgroundStyleColor) GetColorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *BackgroundStyleColor) SetColor(v string) {
	o.Color = v
}

func (o BackgroundStyleColor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["color"] = o.Color
	}
	return json.Marshal(toSerialize)
}

type NullableBackgroundStyleColor struct {
	value *BackgroundStyleColor
	isSet bool
}

func (v NullableBackgroundStyleColor) Get() *BackgroundStyleColor {
	return v.value
}

func (v *NullableBackgroundStyleColor) Set(val *BackgroundStyleColor) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundStyleColor) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundStyleColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundStyleColor(val *BackgroundStyleColor) *NullableBackgroundStyleColor {
	return &NullableBackgroundStyleColor{value: val, isSet: true}
}

func (v NullableBackgroundStyleColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundStyleColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


