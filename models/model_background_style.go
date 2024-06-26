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

// BackgroundStyle Background styles configuration
type BackgroundStyle struct {
	Kind BackgroundStyleKind `json:"kind"`
	Styles BackgroundStyleStyles `json:"styles"`
}

// NewBackgroundStyle instantiates a new BackgroundStyle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundStyle(kind BackgroundStyleKind, styles BackgroundStyleStyles) *BackgroundStyle {
	this := BackgroundStyle{}
	this.Kind = kind
	this.Styles = styles
	return &this
}

// NewBackgroundStyleWithDefaults instantiates a new BackgroundStyle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundStyleWithDefaults() *BackgroundStyle {
	this := BackgroundStyle{}
	return &this
}

// GetKind returns the Kind field value
func (o *BackgroundStyle) GetKind() BackgroundStyleKind {
	if o == nil {
		var ret BackgroundStyleKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BackgroundStyle) GetKindOk() (*BackgroundStyleKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BackgroundStyle) SetKind(v BackgroundStyleKind) {
	o.Kind = v
}

// GetStyles returns the Styles field value
func (o *BackgroundStyle) GetStyles() BackgroundStyleStyles {
	if o == nil {
		var ret BackgroundStyleStyles
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *BackgroundStyle) GetStylesOk() (*BackgroundStyleStyles, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *BackgroundStyle) SetStyles(v BackgroundStyleStyles) {
	o.Styles = v
}

func (o BackgroundStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableBackgroundStyle struct {
	value *BackgroundStyle
	isSet bool
}

func (v NullableBackgroundStyle) Get() *BackgroundStyle {
	return v.value
}

func (v *NullableBackgroundStyle) Set(val *BackgroundStyle) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundStyle) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundStyle(val *BackgroundStyle) *NullableBackgroundStyle {
	return &NullableBackgroundStyle{value: val, isSet: true}
}

func (v NullableBackgroundStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


