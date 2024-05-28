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

// Screen struct for Screen
type Screen struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// Screen tags
	Tags []string `json:"tags"`
	ScreenType ScreenType `json:"screenType"`
	Struct ScreenStruct `json:"struct"`
	Position PositionOnConfigurator `json:"position"`
}

// NewScreen instantiates a new Screen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreen(id string, name string, tags []string, screenType ScreenType, struct_ ScreenStruct, position PositionOnConfigurator) *Screen {
	this := Screen{}
	this.Id = id
	this.Name = name
	this.Tags = tags
	this.ScreenType = screenType
	this.Struct = struct_
	this.Position = position
	return &this
}

// NewScreenWithDefaults instantiates a new Screen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreenWithDefaults() *Screen {
	this := Screen{}
	return &this
}

// GetId returns the Id field value
func (o *Screen) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Screen) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Screen) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Screen) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Screen) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Screen) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value
func (o *Screen) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Screen) GetTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Screen) SetTags(v []string) {
	o.Tags = v
}

// GetScreenType returns the ScreenType field value
func (o *Screen) GetScreenType() ScreenType {
	if o == nil {
		var ret ScreenType
		return ret
	}

	return o.ScreenType
}

// GetScreenTypeOk returns a tuple with the ScreenType field value
// and a boolean to check if the value has been set.
func (o *Screen) GetScreenTypeOk() (*ScreenType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScreenType, true
}

// SetScreenType sets field value
func (o *Screen) SetScreenType(v ScreenType) {
	o.ScreenType = v
}

// GetStruct returns the Struct field value
func (o *Screen) GetStruct() ScreenStruct {
	if o == nil {
		var ret ScreenStruct
		return ret
	}

	return o.Struct
}

// GetStructOk returns a tuple with the Struct field value
// and a boolean to check if the value has been set.
func (o *Screen) GetStructOk() (*ScreenStruct, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Struct, true
}

// SetStruct sets field value
func (o *Screen) SetStruct(v ScreenStruct) {
	o.Struct = v
}

// GetPosition returns the Position field value
func (o *Screen) GetPosition() PositionOnConfigurator {
	if o == nil {
		var ret PositionOnConfigurator
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *Screen) GetPositionOk() (*PositionOnConfigurator, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *Screen) SetPosition(v PositionOnConfigurator) {
	o.Position = v
}

func (o Screen) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["screenType"] = o.ScreenType
	}
	if true {
		toSerialize["struct"] = o.Struct
	}
	if true {
		toSerialize["position"] = o.Position
	}
	return json.Marshal(toSerialize)
}

type NullableScreen struct {
	value *Screen
	isSet bool
}

func (v NullableScreen) Get() *Screen {
	return v.value
}

func (v *NullableScreen) Set(val *Screen) {
	v.value = val
	v.isSet = true
}

func (v NullableScreen) IsSet() bool {
	return v.isSet
}

func (v *NullableScreen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreen(val *Screen) *NullableScreen {
	return &NullableScreen{value: val, isSet: true}
}

func (v NullableScreen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


