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

// BaseLoader Base loader parameters
type BaseLoader struct {
	Kind LoaderKind `json:"kind"`
	Styles LoaderBlock `json:"styles"`
}

// NewBaseLoader instantiates a new BaseLoader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseLoader(kind LoaderKind, styles LoaderBlock) *BaseLoader {
	this := BaseLoader{}
	this.Kind = kind
	this.Styles = styles
	return &this
}

// NewBaseLoaderWithDefaults instantiates a new BaseLoader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseLoaderWithDefaults() *BaseLoader {
	this := BaseLoader{}
	return &this
}

// GetKind returns the Kind field value
func (o *BaseLoader) GetKind() LoaderKind {
	if o == nil {
		var ret LoaderKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BaseLoader) GetKindOk() (*LoaderKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BaseLoader) SetKind(v LoaderKind) {
	o.Kind = v
}

// GetStyles returns the Styles field value
func (o *BaseLoader) GetStyles() LoaderBlock {
	if o == nil {
		var ret LoaderBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *BaseLoader) GetStylesOk() (*LoaderBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *BaseLoader) SetStyles(v LoaderBlock) {
	o.Styles = v
}

func (o BaseLoader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableBaseLoader struct {
	value *BaseLoader
	isSet bool
}

func (v NullableBaseLoader) Get() *BaseLoader {
	return v.value
}

func (v *NullableBaseLoader) Set(val *BaseLoader) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseLoader) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseLoader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseLoader(val *BaseLoader) *NullableBaseLoader {
	return &NullableBaseLoader{value: val, isSet: true}
}

func (v NullableBaseLoader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseLoader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


