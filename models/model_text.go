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

// Text Text parameters
type Text struct {
	Box Box `json:"box"`
	Kind TextKind `json:"kind"`
	// Localized string
	L10n map[string]string `json:"l10n"`
	Parameters TemplateParameters `json:"parameters"`
	Styles LabelBlock `json:"styles"`
}

// NewText instantiates a new Text object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewText(box Box, kind TextKind, l10n map[string]string, parameters TemplateParameters, styles LabelBlock) *Text {
	this := Text{}
	this.Kind = kind
	this.L10n = l10n
	this.Parameters = parameters
	this.Styles = styles
	return &this
}

// NewTextWithDefaults instantiates a new Text object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextWithDefaults() *Text {
	this := Text{}
	return &this
}

// GetBox returns the Box field value
func (o *Text) GetBox() Box {
	if o == nil {
		var ret Box
		return ret
	}

	return o.Box
}

// GetBoxOk returns a tuple with the Box field value
// and a boolean to check if the value has been set.
func (o *Text) GetBoxOk() (*Box, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Box, true
}

// SetBox sets field value
func (o *Text) SetBox(v Box) {
	o.Box = v
}

// GetKind returns the Kind field value
func (o *Text) GetKind() TextKind {
	if o == nil {
		var ret TextKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Text) GetKindOk() (*TextKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Text) SetKind(v TextKind) {
	o.Kind = v
}

// GetL10n returns the L10n field value
func (o *Text) GetL10n() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.L10n
}

// GetL10nOk returns a tuple with the L10n field value
// and a boolean to check if the value has been set.
func (o *Text) GetL10nOk() (*map[string]string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.L10n, true
}

// SetL10n sets field value
func (o *Text) SetL10n(v map[string]string) {
	o.L10n = v
}

// GetParameters returns the Parameters field value
func (o *Text) GetParameters() TemplateParameters {
	if o == nil {
		var ret TemplateParameters
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *Text) GetParametersOk() (*TemplateParameters, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *Text) SetParameters(v TemplateParameters) {
	o.Parameters = v
}

// GetStyles returns the Styles field value
func (o *Text) GetStyles() LabelBlock {
	if o == nil {
		var ret LabelBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *Text) GetStylesOk() (*LabelBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *Text) SetStyles(v LabelBlock) {
	o.Styles = v
}

func (o Text) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["box"] = o.Box
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["l10n"] = o.L10n
	}
	if true {
		toSerialize["parameters"] = o.Parameters
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableText struct {
	value *Text
	isSet bool
}

func (v NullableText) Get() *Text {
	return v.value
}

func (v *NullableText) Set(val *Text) {
	v.value = val
	v.isSet = true
}

func (v NullableText) IsSet() bool {
	return v.isSet
}

func (v *NullableText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableText(val *Text) *NullableText {
	return &NullableText{value: val, isSet: true}
}

func (v NullableText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


