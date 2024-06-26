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

// MediaContent - Media content
type MediaContent struct {
	MediaImage *MediaImage
	MediaVideo *MediaVideo
}

// MediaImageAsMediaContent is a convenience function that returns MediaImage wrapped in MediaContent
func MediaImageAsMediaContent(v *MediaImage) MediaContent {
	return MediaContent{
		MediaImage: v,
	}
}

// MediaVideoAsMediaContent is a convenience function that returns MediaVideo wrapped in MediaContent
func MediaVideoAsMediaContent(v *MediaVideo) MediaContent {
	return MediaContent{
		MediaVideo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MediaContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MediaImage
	err = newStrictDecoder(data).Decode(&dst.MediaImage)
	if err == nil {
		jsonMediaImage, _ := json.Marshal(dst.MediaImage)
		if string(jsonMediaImage) == "{}" { // empty struct
			dst.MediaImage = nil
		} else {
			match++
		}
	} else {
		dst.MediaImage = nil
	}

	// try to unmarshal data into MediaVideo
	err = newStrictDecoder(data).Decode(&dst.MediaVideo)
	if err == nil {
		jsonMediaVideo, _ := json.Marshal(dst.MediaVideo)
		if string(jsonMediaVideo) == "{}" { // empty struct
			dst.MediaVideo = nil
		} else {
			match++
		}
	} else {
		dst.MediaVideo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MediaImage = nil
		dst.MediaVideo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MediaContent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MediaContent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MediaContent) MarshalJSON() ([]byte, error) {
	if src.MediaImage != nil {
		return json.Marshal(&src.MediaImage)
	}

	if src.MediaVideo != nil {
		return json.Marshal(&src.MediaVideo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MediaContent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MediaImage != nil {
		return obj.MediaImage
	}

	if obj.MediaVideo != nil {
		return obj.MediaVideo
	}

	// all schemas are nil
	return nil
}

type NullableMediaContent struct {
	value *MediaContent
	isSet bool
}

func (v NullableMediaContent) Get() *MediaContent {
	return v.value
}

func (v *NullableMediaContent) Set(val *MediaContent) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaContent) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaContent(val *MediaContent) *NullableMediaContent {
	return &NullableMediaContent{value: val, isSet: true}
}

func (v NullableMediaContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


