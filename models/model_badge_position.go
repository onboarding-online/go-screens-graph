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

// BadgePosition Badge position
type BadgePosition string

// List of BadgePosition
const (
	BADGEPOSITION_LEFT_START BadgePosition = "leftStart"
	BADGEPOSITION_LEFT_END BadgePosition = "leftEnd"
	BADGEPOSITION_LEFT_CENTER BadgePosition = "leftCenter"
	BADGEPOSITION_RIGHT_START BadgePosition = "rightStart"
	BADGEPOSITION_RIGHT_END BadgePosition = "rightEnd"
	BADGEPOSITION_RIGHT_CENTER BadgePosition = "rightCenter"
	BADGEPOSITION_TOP_LEFT BadgePosition = "topLeft"
	BADGEPOSITION_TOP_RIGHT BadgePosition = "topRight"
	BADGEPOSITION_TOP_CENTER BadgePosition = "topCenter"
	BADGEPOSITION_BOTTOM_LEFT BadgePosition = "bottomLeft"
	BADGEPOSITION_BOTTOM_RIGHT BadgePosition = "bottomRight"
	BADGEPOSITION_BOTTOM_CENTER BadgePosition = "bottomCenter"
)

// All allowed values of BadgePosition enum
var AllowedBadgePositionEnumValues = []BadgePosition{
	"leftStart",
	"leftEnd",
	"leftCenter",
	"rightStart",
	"rightEnd",
	"rightCenter",
	"topLeft",
	"topRight",
	"topCenter",
	"bottomLeft",
	"bottomRight",
	"bottomCenter",
}

func (v *BadgePosition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BadgePosition(value)
	for _, existing := range AllowedBadgePositionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BadgePosition", value)
}

// NewBadgePositionFromValue returns a pointer to a valid BadgePosition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBadgePositionFromValue(v string) (*BadgePosition, error) {
	ev := BadgePosition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BadgePosition: valid values are %v", v, AllowedBadgePositionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BadgePosition) IsValid() bool {
	for _, existing := range AllowedBadgePositionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BadgePosition value
func (v BadgePosition) Ptr() *BadgePosition {
	return &v
}

type NullableBadgePosition struct {
	value *BadgePosition
	isSet bool
}

func (v NullableBadgePosition) Get() *BadgePosition {
	return v.value
}

func (v *NullableBadgePosition) Set(val *BadgePosition) {
	v.value = val
	v.isSet = true
}

func (v NullableBadgePosition) IsSet() bool {
	return v.isSet
}

func (v *NullableBadgePosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadgePosition(val *BadgePosition) *NullableBadgePosition {
	return &NullableBadgePosition{value: val, isSet: true}
}

func (v NullableBadgePosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadgePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

