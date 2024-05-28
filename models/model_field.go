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

// Field Field element
type Field struct {
	Box Box `json:"box"`
	Type FieldType `json:"type"`
	Subtype *FieldSubtype `json:"subtype,omitempty"`
	ValidationRegex *string `json:"validationRegex,omitempty"`
	ErrorMessage *BaseText `json:"errorMessage,omitempty"`
	Placeholder BaseText `json:"placeholder"`
	// Field value
	Value string `json:"value"`
	Styles FieldBlock `json:"styles"`
}

// NewField instantiates a new Field object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewField(box Box, type_ FieldType, placeholder BaseText, value string, styles FieldBlock) *Field {
	this := Field{}
	this.Type = type_
	this.Placeholder = placeholder
	this.Value = value
	this.Styles = styles
	return &this
}

// NewFieldWithDefaults instantiates a new Field object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldWithDefaults() *Field {
	this := Field{}
	return &this
}

// GetBox returns the Box field value
func (o *Field) GetBox() Box {
	if o == nil {
		var ret Box
		return ret
	}

	return o.Box
}

// GetBoxOk returns a tuple with the Box field value
// and a boolean to check if the value has been set.
func (o *Field) GetBoxOk() (*Box, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Box, true
}

// SetBox sets field value
func (o *Field) SetBox(v Box) {
	o.Box = v
}

// GetType returns the Type field value
func (o *Field) GetType() FieldType {
	if o == nil {
		var ret FieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Field) GetTypeOk() (*FieldType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Field) SetType(v FieldType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *Field) GetSubtype() FieldSubtype {
	if o == nil || isNil(o.Subtype) {
		var ret FieldSubtype
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetSubtypeOk() (*FieldSubtype, bool) {
	if o == nil || isNil(o.Subtype) {
    return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *Field) HasSubtype() bool {
	if o != nil && !isNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given FieldSubtype and assigns it to the Subtype field.
func (o *Field) SetSubtype(v FieldSubtype) {
	o.Subtype = &v
}

// GetValidationRegex returns the ValidationRegex field value if set, zero value otherwise.
func (o *Field) GetValidationRegex() string {
	if o == nil || isNil(o.ValidationRegex) {
		var ret string
		return ret
	}
	return *o.ValidationRegex
}

// GetValidationRegexOk returns a tuple with the ValidationRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetValidationRegexOk() (*string, bool) {
	if o == nil || isNil(o.ValidationRegex) {
    return nil, false
	}
	return o.ValidationRegex, true
}

// HasValidationRegex returns a boolean if a field has been set.
func (o *Field) HasValidationRegex() bool {
	if o != nil && !isNil(o.ValidationRegex) {
		return true
	}

	return false
}

// SetValidationRegex gets a reference to the given string and assigns it to the ValidationRegex field.
func (o *Field) SetValidationRegex(v string) {
	o.ValidationRegex = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *Field) GetErrorMessage() BaseText {
	if o == nil || isNil(o.ErrorMessage) {
		var ret BaseText
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetErrorMessageOk() (*BaseText, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Field) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given BaseText and assigns it to the ErrorMessage field.
func (o *Field) SetErrorMessage(v BaseText) {
	o.ErrorMessage = &v
}

// GetPlaceholder returns the Placeholder field value
func (o *Field) GetPlaceholder() BaseText {
	if o == nil {
		var ret BaseText
		return ret
	}

	return o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value
// and a boolean to check if the value has been set.
func (o *Field) GetPlaceholderOk() (*BaseText, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Placeholder, true
}

// SetPlaceholder sets field value
func (o *Field) SetPlaceholder(v BaseText) {
	o.Placeholder = v
}

// GetValue returns the Value field value
func (o *Field) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Field) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Field) SetValue(v string) {
	o.Value = v
}

// GetStyles returns the Styles field value
func (o *Field) GetStyles() FieldBlock {
	if o == nil {
		var ret FieldBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *Field) GetStylesOk() (*FieldBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *Field) SetStyles(v FieldBlock) {
	o.Styles = v
}

func (o Field) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["box"] = o.Box
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !isNil(o.ValidationRegex) {
		toSerialize["validationRegex"] = o.ValidationRegex
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if true {
		toSerialize["placeholder"] = o.Placeholder
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	return json.Marshal(toSerialize)
}

type NullableField struct {
	value *Field
	isSet bool
}

func (v NullableField) Get() *Field {
	return v.value
}

func (v *NullableField) Set(val *Field) {
	v.value = val
	v.isSet = true
}

func (v NullableField) IsSet() bool {
	return v.isSet
}

func (v *NullableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableField(val *Field) *NullableField {
	return &NullableField{value: val, isSet: true}
}

func (v NullableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


