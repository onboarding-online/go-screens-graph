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

// XYPoint x,y coordinates
type XYPoint struct {
	// x coordinate
	X float32 `json:"x"`
	// y coordinate
	Y float32 `json:"y"`
}

// NewXYPoint instantiates a new XYPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXYPoint(x float32, y float32) *XYPoint {
	this := XYPoint{}
	this.X = x
	this.Y = y
	return &this
}

// NewXYPointWithDefaults instantiates a new XYPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXYPointWithDefaults() *XYPoint {
	this := XYPoint{}
	return &this
}

// GetX returns the X field value
func (o *XYPoint) GetX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *XYPoint) GetXOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *XYPoint) SetX(v float32) {
	o.X = v
}

// GetY returns the Y field value
func (o *XYPoint) GetY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *XYPoint) GetYOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *XYPoint) SetY(v float32) {
	o.Y = v
}

func (o XYPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["x"] = o.X
	}
	if true {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableXYPoint struct {
	value *XYPoint
	isSet bool
}

func (v NullableXYPoint) Get() *XYPoint {
	return v.value
}

func (v *NullableXYPoint) Set(val *XYPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableXYPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableXYPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXYPoint(val *XYPoint) *NullableXYPoint {
	return &NullableXYPoint{value: val, isSet: true}
}

func (v NullableXYPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXYPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


