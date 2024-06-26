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

// ItemTypeSelection Structure of full collection item type for selection list
type ItemTypeSelection struct {
	Image Image `json:"image"`
	Title Text `json:"title"`
	Subtitle Text `json:"subtitle"`
	CheckBox CheckBox `json:"checkBox"`
	Action Action `json:"action"`
	Box Box `json:"box"`
	Weight float32 `json:"weight"`
}

// NewItemTypeSelection instantiates a new ItemTypeSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemTypeSelection(image Image, title Text, subtitle Text, checkBox CheckBox, action Action, box Box, weight float32) *ItemTypeSelection {
	this := ItemTypeSelection{}
	this.Image = image
	this.Title = title
	this.Subtitle = subtitle
	this.CheckBox = checkBox
	this.Action = action
	this.Box = box
	this.Weight = weight
	return &this
}

// NewItemTypeSelectionWithDefaults instantiates a new ItemTypeSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemTypeSelectionWithDefaults() *ItemTypeSelection {
	this := ItemTypeSelection{}
	return &this
}

// GetImage returns the Image field value
func (o *ItemTypeSelection) GetImage() Image {
	if o == nil {
		var ret Image
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ItemTypeSelection) GetImageOk() (*Image, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ItemTypeSelection) SetImage(v Image) {
	o.Image = v
}

// GetTitle returns the Title field value
func (o *ItemTypeSelection) GetTitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ItemTypeSelection) GetTitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ItemTypeSelection) SetTitle(v Text) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *ItemTypeSelection) GetSubtitle() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *ItemTypeSelection) GetSubtitleOk() (*Text, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *ItemTypeSelection) SetSubtitle(v Text) {
	o.Subtitle = v
}

// GetCheckBox returns the CheckBox field value
func (o *ItemTypeSelection) GetCheckBox() CheckBox {
	if o == nil {
		var ret CheckBox
		return ret
	}

	return o.CheckBox
}

// GetCheckBoxOk returns a tuple with the CheckBox field value
// and a boolean to check if the value has been set.
func (o *ItemTypeSelection) GetCheckBoxOk() (*CheckBox, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CheckBox, true
}

// SetCheckBox sets field value
func (o *ItemTypeSelection) SetCheckBox(v CheckBox) {
	o.CheckBox = v
}

// GetAction returns the Action field value
func (o *ItemTypeSelection) GetAction() Action {
	if o == nil {
		var ret Action
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ItemTypeSelection) GetActionOk() (*Action, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ItemTypeSelection) SetAction(v Action) {
	o.Action = v
}

// GetBox returns the Box field value
func (o *ItemTypeSelection) GetBox() Box {
	if o == nil {
		var ret Box
		return ret
	}

	return o.Box
}

// GetBoxOk returns a tuple with the Box field value
// and a boolean to check if the value has been set.
func (o *ItemTypeSelection) GetBoxOk() (*Box, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Box, true
}

// SetBox sets field value
func (o *ItemTypeSelection) SetBox(v Box) {
	o.Box = v
}

// GetWeight returns the Weight field value
func (o *ItemTypeSelection) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ItemTypeSelection) GetWeightOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ItemTypeSelection) SetWeight(v float32) {
	o.Weight = v
}

func (o ItemTypeSelection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["image"] = o.Image
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["subtitle"] = o.Subtitle
	}
	if true {
		toSerialize["checkBox"] = o.CheckBox
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["box"] = o.Box
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableItemTypeSelection struct {
	value *ItemTypeSelection
	isSet bool
}

func (v NullableItemTypeSelection) Get() *ItemTypeSelection {
	return v.value
}

func (v *NullableItemTypeSelection) Set(val *ItemTypeSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableItemTypeSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableItemTypeSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemTypeSelection(val *ItemTypeSelection) *NullableItemTypeSelection {
	return &NullableItemTypeSelection{value: val, isSet: true}
}

func (v NullableItemTypeSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemTypeSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


