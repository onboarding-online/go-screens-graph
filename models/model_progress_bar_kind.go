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

// ProgressBarKind Kind of progress bar
type ProgressBarKind string

// List of ProgressBarKind
const (
	PROGRESSBARKIND_CIRCLE ProgressBarKind = "circle"
	PROGRESSBARKIND_LINE ProgressBarKind = "line"
	PROGRESSBARKIND_MULTIPLE_LINE ProgressBarKind = "multipleLine"
	PROGRESSBARKIND_PROGRESS_BAR_KIND1 ProgressBarKind = "ProgressBarKind1"
	PROGRESSBARKIND_PROGRESS_BAR_KIND2 ProgressBarKind = "ProgressBarKind2"
	PROGRESSBARKIND_PROGRESS_BAR_KIND3 ProgressBarKind = "ProgressBarKind3"
	PROGRESSBARKIND_PROGRESS_BAR_KIND4 ProgressBarKind = "ProgressBarKind4"
	PROGRESSBARKIND_PROGRESS_BAR_KIND5 ProgressBarKind = "ProgressBarKind5"
)

// All allowed values of ProgressBarKind enum
var AllowedProgressBarKindEnumValues = []ProgressBarKind{
	"circle",
	"line",
	"multipleLine",
	"ProgressBarKind1",
	"ProgressBarKind2",
	"ProgressBarKind3",
	"ProgressBarKind4",
	"ProgressBarKind5",
}

func (v *ProgressBarKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProgressBarKind(value)
	for _, existing := range AllowedProgressBarKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProgressBarKind", value)
}

// NewProgressBarKindFromValue returns a pointer to a valid ProgressBarKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProgressBarKindFromValue(v string) (*ProgressBarKind, error) {
	ev := ProgressBarKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProgressBarKind: valid values are %v", v, AllowedProgressBarKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProgressBarKind) IsValid() bool {
	for _, existing := range AllowedProgressBarKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProgressBarKind value
func (v ProgressBarKind) Ptr() *ProgressBarKind {
	return &v
}

type NullableProgressBarKind struct {
	value *ProgressBarKind
	isSet bool
}

func (v NullableProgressBarKind) Get() *ProgressBarKind {
	return v.value
}

func (v *NullableProgressBarKind) Set(val *ProgressBarKind) {
	v.value = val
	v.isSet = true
}

func (v NullableProgressBarKind) IsSet() bool {
	return v.isSet
}

func (v *NullableProgressBarKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgressBarKind(val *ProgressBarKind) *NullableProgressBarKind {
	return &NullableProgressBarKind{value: val, isSet: true}
}

func (v NullableProgressBarKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgressBarKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

