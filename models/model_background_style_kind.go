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

// BackgroundStyleKind Background style kind
type BackgroundStyleKind string

// List of BackgroundStyleKind
const (
	BACKGROUNDSTYLEKIND_IMAGE BackgroundStyleKind = "image"
	BACKGROUNDSTYLEKIND_VIDEO BackgroundStyleKind = "video"
	BACKGROUNDSTYLEKIND_COLOR BackgroundStyleKind = "color"
)

// All allowed values of BackgroundStyleKind enum
var AllowedBackgroundStyleKindEnumValues = []BackgroundStyleKind{
	"image",
	"video",
	"color",
}

func (v *BackgroundStyleKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackgroundStyleKind(value)
	for _, existing := range AllowedBackgroundStyleKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackgroundStyleKind", value)
}

// NewBackgroundStyleKindFromValue returns a pointer to a valid BackgroundStyleKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackgroundStyleKindFromValue(v string) (*BackgroundStyleKind, error) {
	ev := BackgroundStyleKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackgroundStyleKind: valid values are %v", v, AllowedBackgroundStyleKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackgroundStyleKind) IsValid() bool {
	for _, existing := range AllowedBackgroundStyleKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackgroundStyleKind value
func (v BackgroundStyleKind) Ptr() *BackgroundStyleKind {
	return &v
}

type NullableBackgroundStyleKind struct {
	value *BackgroundStyleKind
	isSet bool
}

func (v NullableBackgroundStyleKind) Get() *BackgroundStyleKind {
	return v.value
}

func (v *NullableBackgroundStyleKind) Set(val *BackgroundStyleKind) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundStyleKind) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundStyleKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundStyleKind(val *BackgroundStyleKind) *NullableBackgroundStyleKind {
	return &NullableBackgroundStyleKind{value: val, isSet: true}
}

func (v NullableBackgroundStyleKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundStyleKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

