/*
Onboarding online screens graph models

Onboarding online screens graph models and interfaces

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// ScreenTitleSubtitlePickerAllOf Screen - Title/Subtitle/Picker
type ScreenTitleSubtitlePickerAllOf struct {
	TitleSubtitlePickerDescription string `json:"titleSubtitlePickerDescription"`
	Title Text `json:"title"`
	Subtitle Text `json:"subtitle"`
	Picker Picker `json:"picker"`
}

// NewScreenTitleSubtitlePickerAllOf instantiates a new ScreenTitleSubtitlePickerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenTitleSubtitlePickerAllOf(titleSubtitlePickerDescription string, title Text, subtitle Text, picker Picker) *ScreenTitleSubtitlePickerAllOf {
	this := ScreenTitleSubtitlePickerAllOf{}
	this.TitleSubtitlePickerDescription = titleSubtitlePickerDescription
	this.Title = title
	this.Subtitle = subtitle
	this.Picker = picker
	return &this
}

// NewScreenTitleSubtitlePickerAllOfWithDefaults instantiates a new ScreenTitleSubtitlePickerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenTitleSubtitlePickerAllOfWithDefaults() *ScreenTitleSubtitlePickerAllOf {
	this := ScreenTitleSubtitlePickerAllOf{}
	return &this
}

// GetTitleSubtitlePickerDescription returns the TitleSubtitlePickerDescription field value
func (o *ScreenTitleSubtitlePickerAllOf) GetTitleSubtitlePickerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TitleSubtitlePickerDescription
}

// GetTitleSubtitlePickerDescriptionOk returns a tuple with the TitleSubtitlePickerDescription field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePickerAllOf) GetTitleSubtitlePickerDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TitleSubtitlePickerDescription, true
}

// SetTitleSubtitlePickerDescription sets field value
func (o *ScreenTitleSubtitlePickerAllOf) SetTitleSubtitlePickerDescription(v string) {
	o.TitleSubtitlePickerDescription = v
}

// GetTitle returns the Title field value
func (o *ScreenTitleSubtitlePickerAllOf) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePickerAllOf) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScreenTitleSubtitlePickerAllOf) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *ScreenTitleSubtitlePickerAllOf) GetSubtitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePickerAllOf) GetSubtitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *ScreenTitleSubtitlePickerAllOf) SetSubtitle(v Text) {
	o.Subtitle = v
}

// GetPicker returns the Picker field value
func (o *ScreenTitleSubtitlePickerAllOf) GetPicker() Picker {
	if o == nil {
		var ret Picker
		return ret
	}

	return o.Picker
}

// GetPickerOk returns a tuple with the Picker field value
// and a boolean to check if the value has been set.
func (o *ScreenTitleSubtitlePickerAllOf) GetPickerOk() (*Picker, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Picker, true
}

// SetPicker sets field value
func (o *ScreenTitleSubtitlePickerAllOf) SetPicker(v Picker) {
	o.Picker = v
}

func (o ScreenTitleSubtitlePickerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["titleSubtitlePickerDescription"] = o.TitleSubtitlePickerDescription
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["subtitle"] = o.Subtitle
	}
	if true {
		toSerialize["picker"] = o.Picker
	}
	return json.Marshal(toSerialize)
}

type NullableScreenTitleSubtitlePickerAllOf struct {
	value *ScreenTitleSubtitlePickerAllOf
	isSet bool
}

func (v NullableScreenTitleSubtitlePickerAllOf) Get() *ScreenTitleSubtitlePickerAllOf {
	return v.value
}

func (v *NullableScreenTitleSubtitlePickerAllOf) Set(val *ScreenTitleSubtitlePickerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenTitleSubtitlePickerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenTitleSubtitlePickerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenTitleSubtitlePickerAllOf(val *ScreenTitleSubtitlePickerAllOf) *NullableScreenTitleSubtitlePickerAllOf {
	return &NullableScreenTitleSubtitlePickerAllOf{value: val, isSet: true}
}

func (v NullableScreenTitleSubtitlePickerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenTitleSubtitlePickerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


