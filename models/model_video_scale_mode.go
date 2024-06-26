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

// VideoScaleMode Video scale mode
type VideoScaleMode string

// List of VideoScaleMode
const (
	VIDEOSCALEMODE_SCALE_TO_FILL VideoScaleMode = "scaleToFill"
	VIDEOSCALEMODE_SCALE_ASPECT_FIT VideoScaleMode = "scaleAspectFit"
	VIDEOSCALEMODE_SCALE_ASPECT_FILL VideoScaleMode = "scaleAspectFill"
	VIDEOSCALEMODE_CENTER VideoScaleMode = "center"
	VIDEOSCALEMODE_TOP VideoScaleMode = "top"
	VIDEOSCALEMODE_BOTTOM VideoScaleMode = "bottom"
	VIDEOSCALEMODE_LEFT VideoScaleMode = "left"
	VIDEOSCALEMODE_RIGHT VideoScaleMode = "right"
)

// All allowed values of VideoScaleMode enum
var AllowedVideoScaleModeEnumValues = []VideoScaleMode{
	"scaleToFill",
	"scaleAspectFit",
	"scaleAspectFill",
	"center",
	"top",
	"bottom",
	"left",
	"right",
}

func (v *VideoScaleMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VideoScaleMode(value)
	for _, existing := range AllowedVideoScaleModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VideoScaleMode", value)
}

// NewVideoScaleModeFromValue returns a pointer to a valid VideoScaleMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVideoScaleModeFromValue(v string) (*VideoScaleMode, error) {
	ev := VideoScaleMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VideoScaleMode: valid values are %v", v, AllowedVideoScaleModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VideoScaleMode) IsValid() bool {
	for _, existing := range AllowedVideoScaleModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VideoScaleMode value
func (v VideoScaleMode) Ptr() *VideoScaleMode {
	return &v
}

type NullableVideoScaleMode struct {
	value *VideoScaleMode
	isSet bool
}

func (v NullableVideoScaleMode) Get() *VideoScaleMode {
	return v.value
}

func (v *NullableVideoScaleMode) Set(val *VideoScaleMode) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoScaleMode) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoScaleMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoScaleMode(val *VideoScaleMode) *NullableVideoScaleMode {
	return &NullableVideoScaleMode{value: val, isSet: true}
}

func (v NullableVideoScaleMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoScaleMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

