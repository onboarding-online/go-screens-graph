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

// Button Button parameters
type Button struct {
	Box Box `json:"box"`
	Kind ButtonKind `json:"kind"`
	Content BaseButtonContent `json:"content"`
	Styles ButtonBlock `json:"styles"`
	Action Action `json:"action"`
}

// NewButton instantiates a new Button object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewButton(box Box, kind ButtonKind, content BaseButtonContent, styles ButtonBlock, action Action) *Button {
	this := Button{}
	this.Kind = kind
	this.Content = content
	this.Styles = styles
	this.Action = action
	return &this
}

// NewButtonWithDefaults instantiates a new Button object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewButtonWithDefaults() *Button {
	this := Button{}
	return &this
}

// GetBox returns the Box field value
func (o *Button) GetBox() Box {
	if o == nil {
		var ret Box
		return ret
	}

	return o.Box
}

// GetBoxOk returns a tuple with the Box field value
// and a boolean to check if the value has been set.
func (o *Button) GetBoxOk() (*Box, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Box, true
}

// SetBox sets field value
func (o *Button) SetBox(v Box) {
	o.Box = v
}

// GetKind returns the Kind field value
func (o *Button) GetKind() ButtonKind {
	if o == nil {
		var ret ButtonKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Button) GetKindOk() (*ButtonKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Button) SetKind(v ButtonKind) {
	o.Kind = v
}

// GetContent returns the Content field value
func (o *Button) GetContent() BaseButtonContent {
	if o == nil {
		var ret BaseButtonContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Button) GetContentOk() (*BaseButtonContent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Button) SetContent(v BaseButtonContent) {
	o.Content = v
}

// GetStyles returns the Styles field value
func (o *Button) GetStyles() ButtonBlock {
	if o == nil {
		var ret ButtonBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *Button) GetStylesOk() (*ButtonBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *Button) SetStyles(v ButtonBlock) {
	o.Styles = v
}

// GetAction returns the Action field value
func (o *Button) GetAction() Action {
	if o == nil {
		var ret Action
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *Button) GetActionOk() (*Action, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *Button) SetAction(v Action) {
	o.Action = v
}

func (o Button) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["box"] = o.Box
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	if true {
		toSerialize["action"] = o.Action
	}
	return json.Marshal(toSerialize)
}

type NullableButton struct {
	value *Button
	isSet bool
}

func (v NullableButton) Get() *Button {
	return v.value
}

func (v *NullableButton) Set(val *Button) {
	v.value = val
	v.isSet = true
}

func (v NullableButton) IsSet() bool {
	return v.isSet
}

func (v *NullableButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButton(val *Button) *NullableButton {
	return &NullableButton{value: val, isSet: true}
}

func (v NullableButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


