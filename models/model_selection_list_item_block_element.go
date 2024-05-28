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

// SelectionListItemBlockElement Possible elements on selection list
type SelectionListItemBlockElement string

// List of SelectionListItemBlockElement
const (
	SELECTIONLISTITEMBLOCKELEMENT_IMAGE SelectionListItemBlockElement = "Image"
	SELECTIONLISTITEMBLOCKELEMENT_TITLE SelectionListItemBlockElement = "Title"
	SELECTIONLISTITEMBLOCKELEMENT_SUBTITLE SelectionListItemBlockElement = "Subtitle"
	SELECTIONLISTITEMBLOCKELEMENT_CHECK_BOX SelectionListItemBlockElement = "CheckBox"
)

// All allowed values of SelectionListItemBlockElement enum
var AllowedSelectionListItemBlockElementEnumValues = []SelectionListItemBlockElement{
	"Image",
	"Title",
	"Subtitle",
	"CheckBox",
}

func (v *SelectionListItemBlockElement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelectionListItemBlockElement(value)
	for _, existing := range AllowedSelectionListItemBlockElementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelectionListItemBlockElement", value)
}

// NewSelectionListItemBlockElementFromValue returns a pointer to a valid SelectionListItemBlockElement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelectionListItemBlockElementFromValue(v string) (*SelectionListItemBlockElement, error) {
	ev := SelectionListItemBlockElement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelectionListItemBlockElement: valid values are %v", v, AllowedSelectionListItemBlockElementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelectionListItemBlockElement) IsValid() bool {
	for _, existing := range AllowedSelectionListItemBlockElementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelectionListItemBlockElement value
func (v SelectionListItemBlockElement) Ptr() *SelectionListItemBlockElement {
	return &v
}

type NullableSelectionListItemBlockElement struct {
	value *SelectionListItemBlockElement
	isSet bool
}

func (v NullableSelectionListItemBlockElement) Get() *SelectionListItemBlockElement {
	return v.value
}

func (v *NullableSelectionListItemBlockElement) Set(val *SelectionListItemBlockElement) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectionListItemBlockElement) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectionListItemBlockElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectionListItemBlockElement(val *SelectionListItemBlockElement) *NullableSelectionListItemBlockElement {
	return &NullableSelectionListItemBlockElement{value: val, isSet: true}
}

func (v NullableSelectionListItemBlockElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectionListItemBlockElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

