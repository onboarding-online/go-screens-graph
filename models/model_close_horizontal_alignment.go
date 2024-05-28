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

// CloseHorizontalAlignment Close button horizontal alignment
type CloseHorizontalAlignment string

// List of CloseHorizontalAlignment
const (
	CLOSEHORIZONTALALIGNMENT_LEFT CloseHorizontalAlignment = "left"
	CLOSEHORIZONTALALIGNMENT_RIGHT CloseHorizontalAlignment = "right"
)

// All allowed values of CloseHorizontalAlignment enum
var AllowedCloseHorizontalAlignmentEnumValues = []CloseHorizontalAlignment{
	"left",
	"right",
}

func (v *CloseHorizontalAlignment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloseHorizontalAlignment(value)
	for _, existing := range AllowedCloseHorizontalAlignmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloseHorizontalAlignment", value)
}

// NewCloseHorizontalAlignmentFromValue returns a pointer to a valid CloseHorizontalAlignment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloseHorizontalAlignmentFromValue(v string) (*CloseHorizontalAlignment, error) {
	ev := CloseHorizontalAlignment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloseHorizontalAlignment: valid values are %v", v, AllowedCloseHorizontalAlignmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloseHorizontalAlignment) IsValid() bool {
	for _, existing := range AllowedCloseHorizontalAlignmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloseHorizontalAlignment value
func (v CloseHorizontalAlignment) Ptr() *CloseHorizontalAlignment {
	return &v
}

type NullableCloseHorizontalAlignment struct {
	value *CloseHorizontalAlignment
	isSet bool
}

func (v NullableCloseHorizontalAlignment) Get() *CloseHorizontalAlignment {
	return v.value
}

func (v *NullableCloseHorizontalAlignment) Set(val *CloseHorizontalAlignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseHorizontalAlignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseHorizontalAlignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseHorizontalAlignment(val *CloseHorizontalAlignment) *NullableCloseHorizontalAlignment {
	return &NullableCloseHorizontalAlignment{value: val, isSet: true}
}

func (v NullableCloseHorizontalAlignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseHorizontalAlignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

