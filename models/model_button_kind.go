/*
Onboarding online screens graph models

Onboarding online screens graph models and interfaces

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// ButtonKind Button kind
type ButtonKind string

// List of ButtonKind
const (
	BUTTONKIND_TEXT ButtonKind = "text"
	BUTTONKIND_IMAGE ButtonKind = "image"
)

// All allowed values of ButtonKind enum
var AllowedButtonKindEnumValues = []ButtonKind{
	"text",
	"image",
}

func (v *ButtonKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ButtonKind(value)
	for _, existing := range AllowedButtonKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ButtonKind", value)
}

// NewButtonKindFromValue returns a pointer to a valid ButtonKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewButtonKindFromValue(v string) (*ButtonKind, error) {
	ev := ButtonKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ButtonKind: valid values are %v", v, AllowedButtonKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ButtonKind) IsValid() bool {
	for _, existing := range AllowedButtonKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ButtonKind value
func (v ButtonKind) Ptr() *ButtonKind {
	return &v
}

type NullableButtonKind struct {
	value *ButtonKind
	isSet bool
}

func (v NullableButtonKind) Get() *ButtonKind {
	return v.value
}

func (v *NullableButtonKind) Set(val *ButtonKind) {
	v.value = val
	v.isSet = true
}

func (v NullableButtonKind) IsSet() bool {
	return v.isSet
}

func (v *NullableButtonKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButtonKind(val *ButtonKind) *NullableButtonKind {
	return &NullableButtonKind{value: val, isSet: true}
}

func (v NullableButtonKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButtonKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

