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

// FontFamily Font family
type FontFamily string

// List of FontFamily
const (
	FONTFAMILY_SF_PRO FontFamily = "SF Pro"
	FONTFAMILY_SF_PRO_ROUNDED FontFamily = "SF ProRounded"
	FONTFAMILY_SF_MONO FontFamily = "SF Mono"
	FONTFAMILY_NEW_YORK FontFamily = "New York"
)

// All allowed values of FontFamily enum
var AllowedFontFamilyEnumValues = []FontFamily{
	"SF Pro",
	"SF ProRounded",
	"SF Mono",
	"New York",
}

func (v *FontFamily) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FontFamily(value)
	for _, existing := range AllowedFontFamilyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FontFamily", value)
}

// NewFontFamilyFromValue returns a pointer to a valid FontFamily
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFontFamilyFromValue(v string) (*FontFamily, error) {
	ev := FontFamily(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FontFamily: valid values are %v", v, AllowedFontFamilyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FontFamily) IsValid() bool {
	for _, existing := range AllowedFontFamilyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FontFamily value
func (v FontFamily) Ptr() *FontFamily {
	return &v
}

type NullableFontFamily struct {
	value *FontFamily
	isSet bool
}

func (v NullableFontFamily) Get() *FontFamily {
	return v.value
}

func (v *NullableFontFamily) Set(val *FontFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableFontFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableFontFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFontFamily(val *FontFamily) *NullableFontFamily {
	return &NullableFontFamily{value: val, isSet: true}
}

func (v NullableFontFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFontFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

