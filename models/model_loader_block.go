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

// LoaderBlock Styles for loader block
type LoaderBlock struct {
	// Color for loader
	Color *string `json:"color,omitempty"`
	// Padding left for container
	PaddingLeft *float32 `json:"paddingLeft,omitempty"`
	// Padding right for container
	PaddingRight *float32 `json:"paddingRight,omitempty"`
	// Padding top for container
	PaddingTop *float32 `json:"paddingTop,omitempty"`
	// Padding bottom for container
	PaddingBottom *float32 `json:"paddingBottom,omitempty"`
}

// NewLoaderBlock instantiates a new LoaderBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoaderBlock() *LoaderBlock {
	this := LoaderBlock{}
	return &this
}

// NewLoaderBlockWithDefaults instantiates a new LoaderBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoaderBlockWithDefaults() *LoaderBlock {
	this := LoaderBlock{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *LoaderBlock) GetColor() string {
	if o == nil || isNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoaderBlock) GetColorOk() (*string, bool) {
	if o == nil || isNil(o.Color) {
    return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *LoaderBlock) HasColor() bool {
	if o != nil && !isNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *LoaderBlock) SetColor(v string) {
	o.Color = &v
}

// GetPaddingLeft returns the PaddingLeft field value if set, zero value otherwise.
func (o *LoaderBlock) GetPaddingLeft() float32 {
	if o == nil || isNil(o.PaddingLeft) {
		var ret float32
		return ret
	}
	return *o.PaddingLeft
}

// GetPaddingLeftOk returns a tuple with the PaddingLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoaderBlock) GetPaddingLeftOk() (*float32, bool) {
	if o == nil || isNil(o.PaddingLeft) {
    return nil, false
	}
	return o.PaddingLeft, true
}

// HasPaddingLeft returns a boolean if a field has been set.
func (o *LoaderBlock) HasPaddingLeft() bool {
	if o != nil && !isNil(o.PaddingLeft) {
		return true
	}

	return false
}

// SetPaddingLeft gets a reference to the given float32 and assigns it to the PaddingLeft field.
func (o *LoaderBlock) SetPaddingLeft(v float32) {
	o.PaddingLeft = &v
}

// GetPaddingRight returns the PaddingRight field value if set, zero value otherwise.
func (o *LoaderBlock) GetPaddingRight() float32 {
	if o == nil || isNil(o.PaddingRight) {
		var ret float32
		return ret
	}
	return *o.PaddingRight
}

// GetPaddingRightOk returns a tuple with the PaddingRight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoaderBlock) GetPaddingRightOk() (*float32, bool) {
	if o == nil || isNil(o.PaddingRight) {
    return nil, false
	}
	return o.PaddingRight, true
}

// HasPaddingRight returns a boolean if a field has been set.
func (o *LoaderBlock) HasPaddingRight() bool {
	if o != nil && !isNil(o.PaddingRight) {
		return true
	}

	return false
}

// SetPaddingRight gets a reference to the given float32 and assigns it to the PaddingRight field.
func (o *LoaderBlock) SetPaddingRight(v float32) {
	o.PaddingRight = &v
}

// GetPaddingTop returns the PaddingTop field value if set, zero value otherwise.
func (o *LoaderBlock) GetPaddingTop() float32 {
	if o == nil || isNil(o.PaddingTop) {
		var ret float32
		return ret
	}
	return *o.PaddingTop
}

// GetPaddingTopOk returns a tuple with the PaddingTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoaderBlock) GetPaddingTopOk() (*float32, bool) {
	if o == nil || isNil(o.PaddingTop) {
    return nil, false
	}
	return o.PaddingTop, true
}

// HasPaddingTop returns a boolean if a field has been set.
func (o *LoaderBlock) HasPaddingTop() bool {
	if o != nil && !isNil(o.PaddingTop) {
		return true
	}

	return false
}

// SetPaddingTop gets a reference to the given float32 and assigns it to the PaddingTop field.
func (o *LoaderBlock) SetPaddingTop(v float32) {
	o.PaddingTop = &v
}

// GetPaddingBottom returns the PaddingBottom field value if set, zero value otherwise.
func (o *LoaderBlock) GetPaddingBottom() float32 {
	if o == nil || isNil(o.PaddingBottom) {
		var ret float32
		return ret
	}
	return *o.PaddingBottom
}

// GetPaddingBottomOk returns a tuple with the PaddingBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoaderBlock) GetPaddingBottomOk() (*float32, bool) {
	if o == nil || isNil(o.PaddingBottom) {
    return nil, false
	}
	return o.PaddingBottom, true
}

// HasPaddingBottom returns a boolean if a field has been set.
func (o *LoaderBlock) HasPaddingBottom() bool {
	if o != nil && !isNil(o.PaddingBottom) {
		return true
	}

	return false
}

// SetPaddingBottom gets a reference to the given float32 and assigns it to the PaddingBottom field.
func (o *LoaderBlock) SetPaddingBottom(v float32) {
	o.PaddingBottom = &v
}

func (o LoaderBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !isNil(o.PaddingLeft) {
		toSerialize["paddingLeft"] = o.PaddingLeft
	}
	if !isNil(o.PaddingRight) {
		toSerialize["paddingRight"] = o.PaddingRight
	}
	if !isNil(o.PaddingTop) {
		toSerialize["paddingTop"] = o.PaddingTop
	}
	if !isNil(o.PaddingBottom) {
		toSerialize["paddingBottom"] = o.PaddingBottom
	}
	return json.Marshal(toSerialize)
}

type NullableLoaderBlock struct {
	value *LoaderBlock
	isSet bool
}

func (v NullableLoaderBlock) Get() *LoaderBlock {
	return v.value
}

func (v *NullableLoaderBlock) Set(val *LoaderBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableLoaderBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableLoaderBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoaderBlock(val *LoaderBlock) *NullableLoaderBlock {
	return &NullableLoaderBlock{value: val, isSet: true}
}

func (v NullableLoaderBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoaderBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


