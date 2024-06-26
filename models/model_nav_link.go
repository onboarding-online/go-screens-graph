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

// NavLink Navigation link parameters
type NavLink struct {
	Box Box `json:"box"`
	Kind NavLinkKind `json:"kind"`
	Content BaseNavLinkContent `json:"content"`
	Styles NavLinkBlock `json:"styles"`
	// The uri Link
	Uri string `json:"uri"`
}

// NewNavLink instantiates a new NavLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNavLink(box Box, kind NavLinkKind, content BaseNavLinkContent, styles NavLinkBlock, uri string) *NavLink {
	this := NavLink{}
	this.Kind = kind
	this.Content = content
	this.Styles = styles
	this.Uri = uri
	return &this
}

// NewNavLinkWithDefaults instantiates a new NavLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNavLinkWithDefaults() *NavLink {
	this := NavLink{}
	return &this
}

// GetBox returns the Box field value
func (o *NavLink) GetBox() Box {
	if o == nil {
		var ret Box
		return ret
	}

	return o.Box
}

// GetBoxOk returns a tuple with the Box field value
// and a boolean to check if the value has been set.
func (o *NavLink) GetBoxOk() (*Box, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Box, true
}

// SetBox sets field value
func (o *NavLink) SetBox(v Box) {
	o.Box = v
}

// GetKind returns the Kind field value
func (o *NavLink) GetKind() NavLinkKind {
	if o == nil {
		var ret NavLinkKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NavLink) GetKindOk() (*NavLinkKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NavLink) SetKind(v NavLinkKind) {
	o.Kind = v
}

// GetContent returns the Content field value
func (o *NavLink) GetContent() BaseNavLinkContent {
	if o == nil {
		var ret BaseNavLinkContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *NavLink) GetContentOk() (*BaseNavLinkContent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *NavLink) SetContent(v BaseNavLinkContent) {
	o.Content = v
}

// GetStyles returns the Styles field value
func (o *NavLink) GetStyles() NavLinkBlock {
	if o == nil {
		var ret NavLinkBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *NavLink) GetStylesOk() (*NavLinkBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *NavLink) SetStyles(v NavLinkBlock) {
	o.Styles = v
}

// GetUri returns the Uri field value
func (o *NavLink) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *NavLink) GetUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *NavLink) SetUri(v string) {
	o.Uri = v
}

func (o NavLink) MarshalJSON() ([]byte, error) {
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
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableNavLink struct {
	value *NavLink
	isSet bool
}

func (v NullableNavLink) Get() *NavLink {
	return v.value
}

func (v *NullableNavLink) Set(val *NavLink) {
	v.value = val
	v.isSet = true
}

func (v NullableNavLink) IsSet() bool {
	return v.isSet
}

func (v *NullableNavLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavLink(val *NavLink) *NullableNavLink {
	return &NullableNavLink{value: val, isSet: true}
}

func (v NullableNavLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


