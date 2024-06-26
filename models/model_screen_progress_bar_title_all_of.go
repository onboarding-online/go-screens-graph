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

// ScreenProgressBarTitleAllOf Screen - Progress bar/Title
type ScreenProgressBarTitleAllOf struct {
	ProgressBarTitleDescription string `json:"progressBarTitleDescription"`
	ProgressBar ProgressBar `json:"progressBar"`
	Title Text `json:"title"`
	Subtitle *Text `json:"subtitle,omitempty"`
}

// NewScreenProgressBarTitleAllOf instantiates a new ScreenProgressBarTitleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreenProgressBarTitleAllOf(progressBarTitleDescription string, progressBar ProgressBar, title Text) *ScreenProgressBarTitleAllOf {
	this := ScreenProgressBarTitleAllOf{}
	this.ProgressBarTitleDescription = progressBarTitleDescription
	this.ProgressBar = progressBar
	this.Title = title
	return &this
}

// NewScreenProgressBarTitleAllOfWithDefaults instantiates a new ScreenProgressBarTitleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenProgressBarTitleAllOfWithDefaults() *ScreenProgressBarTitleAllOf {
	this := ScreenProgressBarTitleAllOf{}
	return &this
}

// GetProgressBarTitleDescription returns the ProgressBarTitleDescription field value
func (o *ScreenProgressBarTitleAllOf) GetProgressBarTitleDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProgressBarTitleDescription
}

// GetProgressBarTitleDescriptionOk returns a tuple with the ProgressBarTitleDescription field value
// and a boolean to check if the value has been set.
func (o *ScreenProgressBarTitleAllOf) GetProgressBarTitleDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProgressBarTitleDescription, true
}

// SetProgressBarTitleDescription sets field value
func (o *ScreenProgressBarTitleAllOf) SetProgressBarTitleDescription(v string) {
	o.ProgressBarTitleDescription = v
}

// GetProgressBar returns the ProgressBar field value
func (o *ScreenProgressBarTitleAllOf) GetProgressBar() ProgressBar {
	if o == nil {
		var ret ProgressBar
		return ret
	}

	return o.ProgressBar
}

// GetProgressBarOk returns a tuple with the ProgressBar field value
// and a boolean to check if the value has been set.
func (o *ScreenProgressBarTitleAllOf) GetProgressBarOk() (*ProgressBar, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProgressBar, true
}

// SetProgressBar sets field value
func (o *ScreenProgressBarTitleAllOf) SetProgressBar(v ProgressBar) {
	o.ProgressBar = v
}

// GetTitle returns the Title field value
func (o *ScreenProgressBarTitleAllOf) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScreenProgressBarTitleAllOf) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScreenProgressBarTitleAllOf) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *ScreenProgressBarTitleAllOf) GetSubtitle() Text {
	if o == nil || isNil(o.Subtitle) {
		var ret Text
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreenProgressBarTitleAllOf) GetSubtitleOk() (*Text, bool) {
	if o == nil || isNil(o.Subtitle) {
    return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *ScreenProgressBarTitleAllOf) HasSubtitle() bool {
	if o != nil && !isNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given Text and assigns it to the Subtitle field.
func (o *ScreenProgressBarTitleAllOf) SetSubtitle(v Text) {
	o.Subtitle = &v
}

func (o ScreenProgressBarTitleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["progressBarTitleDescription"] = o.ProgressBarTitleDescription
	}
	if true {
		toSerialize["progressBar"] = o.ProgressBar
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	return json.Marshal(toSerialize)
}

type NullableScreenProgressBarTitleAllOf struct {
	value *ScreenProgressBarTitleAllOf
	isSet bool
}

func (v NullableScreenProgressBarTitleAllOf) Get() *ScreenProgressBarTitleAllOf {
	return v.value
}

func (v *NullableScreenProgressBarTitleAllOf) Set(val *ScreenProgressBarTitleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScreenProgressBarTitleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScreenProgressBarTitleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreenProgressBarTitleAllOf(val *ScreenProgressBarTitleAllOf) *NullableScreenProgressBarTitleAllOf {
	return &NullableScreenProgressBarTitleAllOf{value: val, isSet: true}
}

func (v NullableScreenProgressBarTitleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreenProgressBarTitleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


