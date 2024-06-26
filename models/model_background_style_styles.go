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

// BackgroundStyleStyles - struct for BackgroundStyleStyles
type BackgroundStyleStyles struct {
	BackgroundStyleColor *BackgroundStyleColor
	BackgroundStyleImage *BackgroundStyleImage
	BackgroundStyleVideo *BackgroundStyleVideo
}

// BackgroundStyleColorAsBackgroundStyleStyles is a convenience function that returns BackgroundStyleColor wrapped in BackgroundStyleStyles
func BackgroundStyleColorAsBackgroundStyleStyles(v *BackgroundStyleColor) BackgroundStyleStyles {
	return BackgroundStyleStyles{
		BackgroundStyleColor: v,
	}
}

// BackgroundStyleImageAsBackgroundStyleStyles is a convenience function that returns BackgroundStyleImage wrapped in BackgroundStyleStyles
func BackgroundStyleImageAsBackgroundStyleStyles(v *BackgroundStyleImage) BackgroundStyleStyles {
	return BackgroundStyleStyles{
		BackgroundStyleImage: v,
	}
}

// BackgroundStyleVideoAsBackgroundStyleStyles is a convenience function that returns BackgroundStyleVideo wrapped in BackgroundStyleStyles
func BackgroundStyleVideoAsBackgroundStyleStyles(v *BackgroundStyleVideo) BackgroundStyleStyles {
	return BackgroundStyleStyles{
		BackgroundStyleVideo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BackgroundStyleStyles) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BackgroundStyleColor
	err = newStrictDecoder(data).Decode(&dst.BackgroundStyleColor)
	if err == nil {
		jsonBackgroundStyleColor, _ := json.Marshal(dst.BackgroundStyleColor)
		if string(jsonBackgroundStyleColor) == "{}" { // empty struct
			dst.BackgroundStyleColor = nil
		} else {
			match++
		}
	} else {
		dst.BackgroundStyleColor = nil
	}

	// try to unmarshal data into BackgroundStyleImage
	err = newStrictDecoder(data).Decode(&dst.BackgroundStyleImage)
	if err == nil {
		jsonBackgroundStyleImage, _ := json.Marshal(dst.BackgroundStyleImage)
		if string(jsonBackgroundStyleImage) == "{}" { // empty struct
			dst.BackgroundStyleImage = nil
		} else {
			match++
		}
	} else {
		dst.BackgroundStyleImage = nil
	}

	// try to unmarshal data into BackgroundStyleVideo
	err = newStrictDecoder(data).Decode(&dst.BackgroundStyleVideo)
	if err == nil {
		jsonBackgroundStyleVideo, _ := json.Marshal(dst.BackgroundStyleVideo)
		if string(jsonBackgroundStyleVideo) == "{}" { // empty struct
			dst.BackgroundStyleVideo = nil
		} else {
			match++
		}
	} else {
		dst.BackgroundStyleVideo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BackgroundStyleColor = nil
		dst.BackgroundStyleImage = nil
		dst.BackgroundStyleVideo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BackgroundStyleStyles)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BackgroundStyleStyles)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BackgroundStyleStyles) MarshalJSON() ([]byte, error) {
	if src.BackgroundStyleColor != nil {
		return json.Marshal(&src.BackgroundStyleColor)
	}

	if src.BackgroundStyleImage != nil {
		return json.Marshal(&src.BackgroundStyleImage)
	}

	if src.BackgroundStyleVideo != nil {
		return json.Marshal(&src.BackgroundStyleVideo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BackgroundStyleStyles) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BackgroundStyleColor != nil {
		return obj.BackgroundStyleColor
	}

	if obj.BackgroundStyleImage != nil {
		return obj.BackgroundStyleImage
	}

	if obj.BackgroundStyleVideo != nil {
		return obj.BackgroundStyleVideo
	}

	// all schemas are nil
	return nil
}

type NullableBackgroundStyleStyles struct {
	value *BackgroundStyleStyles
	isSet bool
}

func (v NullableBackgroundStyleStyles) Get() *BackgroundStyleStyles {
	return v.value
}

func (v *NullableBackgroundStyleStyles) Set(val *BackgroundStyleStyles) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundStyleStyles) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundStyleStyles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundStyleStyles(val *BackgroundStyleStyles) *NullableBackgroundStyleStyles {
	return &NullableBackgroundStyleStyles{value: val, isSet: true}
}

func (v NullableBackgroundStyleStyles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundStyleStyles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


