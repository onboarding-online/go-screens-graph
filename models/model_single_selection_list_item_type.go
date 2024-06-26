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

// SingleSelectionListItemType Combination of fields for single selection list item
type SingleSelectionListItemType string

// List of SingleSelectionListItemType
const (
	SINGLESELECTIONLISTITEMTYPE_IMAGE_TITLE_SUBTITLE SingleSelectionListItemType = "ImageTitleSubtitle"
	SINGLESELECTIONLISTITEMTYPE_TITLE_SUBTITLE_IMAGE SingleSelectionListItemType = "TitleSubtitleImage"
	SINGLESELECTIONLISTITEMTYPE_TITLE_SUBTITLE SingleSelectionListItemType = "TitleSubtitle"
	SINGLESELECTIONLISTITEMTYPE_TITLE_IMAGE SingleSelectionListItemType = "TitleImage"
	SINGLESELECTIONLISTITEMTYPE_IMAGE_TITLE SingleSelectionListItemType = "ImageTitle"
	SINGLESELECTIONLISTITEMTYPE_TITLE SingleSelectionListItemType = "Title"
)

// All allowed values of SingleSelectionListItemType enum
var AllowedSingleSelectionListItemTypeEnumValues = []SingleSelectionListItemType{
	"ImageTitleSubtitle",
	"TitleSubtitleImage",
	"TitleSubtitle",
	"TitleImage",
	"ImageTitle",
	"Title",
}

func (v *SingleSelectionListItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SingleSelectionListItemType(value)
	for _, existing := range AllowedSingleSelectionListItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SingleSelectionListItemType", value)
}

// NewSingleSelectionListItemTypeFromValue returns a pointer to a valid SingleSelectionListItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSingleSelectionListItemTypeFromValue(v string) (*SingleSelectionListItemType, error) {
	ev := SingleSelectionListItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SingleSelectionListItemType: valid values are %v", v, AllowedSingleSelectionListItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SingleSelectionListItemType) IsValid() bool {
	for _, existing := range AllowedSingleSelectionListItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SingleSelectionListItemType value
func (v SingleSelectionListItemType) Ptr() *SingleSelectionListItemType {
	return &v
}

type NullableSingleSelectionListItemType struct {
	value *SingleSelectionListItemType
	isSet bool
}

func (v NullableSingleSelectionListItemType) Get() *SingleSelectionListItemType {
	return v.value
}

func (v *NullableSingleSelectionListItemType) Set(val *SingleSelectionListItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleSelectionListItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleSelectionListItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleSelectionListItemType(val *SingleSelectionListItemType) *NullableSingleSelectionListItemType {
	return &NullableSingleSelectionListItemType{value: val, isSet: true}
}

func (v NullableSingleSelectionListItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleSelectionListItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

