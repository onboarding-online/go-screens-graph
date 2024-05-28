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

// BadgeBlockAllOf Styles for badge block
type BadgeBlockAllOf struct {
	Position *BadgePosition `json:"position,omitempty"`
	// Radius of border
	BorderRadius *float32 `json:"borderRadius,omitempty"`
	// Width of border
	BorderWidth *float32 `json:"borderWidth,omitempty"`
	// Color of border
	BorderColor *string `json:"borderColor,omitempty"`
}

// NewBadgeBlockAllOf instantiates a new BadgeBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadgeBlockAllOf() *BadgeBlockAllOf {
	this := BadgeBlockAllOf{}
	return &this
}

// NewBadgeBlockAllOfWithDefaults instantiates a new BadgeBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadgeBlockAllOfWithDefaults() *BadgeBlockAllOf {
	this := BadgeBlockAllOf{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *BadgeBlockAllOf) GetPosition() BadgePosition {
	if o == nil || isNil(o.Position) {
		var ret BadgePosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadgeBlockAllOf) GetPositionOk() (*BadgePosition, bool) {
	if o == nil || isNil(o.Position) {
    return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *BadgeBlockAllOf) HasPosition() bool {
	if o != nil && !isNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given BadgePosition and assigns it to the Position field.
func (o *BadgeBlockAllOf) SetPosition(v BadgePosition) {
	o.Position = &v
}

// GetBorderRadius returns the BorderRadius field value if set, zero value otherwise.
func (o *BadgeBlockAllOf) GetBorderRadius() float32 {
	if o == nil || isNil(o.BorderRadius) {
		var ret float32
		return ret
	}
	return *o.BorderRadius
}

// GetBorderRadiusOk returns a tuple with the BorderRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadgeBlockAllOf) GetBorderRadiusOk() (*float32, bool) {
	if o == nil || isNil(o.BorderRadius) {
    return nil, false
	}
	return o.BorderRadius, true
}

// HasBorderRadius returns a boolean if a field has been set.
func (o *BadgeBlockAllOf) HasBorderRadius() bool {
	if o != nil && !isNil(o.BorderRadius) {
		return true
	}

	return false
}

// SetBorderRadius gets a reference to the given float32 and assigns it to the BorderRadius field.
func (o *BadgeBlockAllOf) SetBorderRadius(v float32) {
	o.BorderRadius = &v
}

// GetBorderWidth returns the BorderWidth field value if set, zero value otherwise.
func (o *BadgeBlockAllOf) GetBorderWidth() float32 {
	if o == nil || isNil(o.BorderWidth) {
		var ret float32
		return ret
	}
	return *o.BorderWidth
}

// GetBorderWidthOk returns a tuple with the BorderWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadgeBlockAllOf) GetBorderWidthOk() (*float32, bool) {
	if o == nil || isNil(o.BorderWidth) {
    return nil, false
	}
	return o.BorderWidth, true
}

// HasBorderWidth returns a boolean if a field has been set.
func (o *BadgeBlockAllOf) HasBorderWidth() bool {
	if o != nil && !isNil(o.BorderWidth) {
		return true
	}

	return false
}

// SetBorderWidth gets a reference to the given float32 and assigns it to the BorderWidth field.
func (o *BadgeBlockAllOf) SetBorderWidth(v float32) {
	o.BorderWidth = &v
}

// GetBorderColor returns the BorderColor field value if set, zero value otherwise.
func (o *BadgeBlockAllOf) GetBorderColor() string {
	if o == nil || isNil(o.BorderColor) {
		var ret string
		return ret
	}
	return *o.BorderColor
}

// GetBorderColorOk returns a tuple with the BorderColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadgeBlockAllOf) GetBorderColorOk() (*string, bool) {
	if o == nil || isNil(o.BorderColor) {
    return nil, false
	}
	return o.BorderColor, true
}

// HasBorderColor returns a boolean if a field has been set.
func (o *BadgeBlockAllOf) HasBorderColor() bool {
	if o != nil && !isNil(o.BorderColor) {
		return true
	}

	return false
}

// SetBorderColor gets a reference to the given string and assigns it to the BorderColor field.
func (o *BadgeBlockAllOf) SetBorderColor(v string) {
	o.BorderColor = &v
}

func (o BadgeBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !isNil(o.BorderRadius) {
		toSerialize["borderRadius"] = o.BorderRadius
	}
	if !isNil(o.BorderWidth) {
		toSerialize["borderWidth"] = o.BorderWidth
	}
	if !isNil(o.BorderColor) {
		toSerialize["borderColor"] = o.BorderColor
	}
	return json.Marshal(toSerialize)
}

type NullableBadgeBlockAllOf struct {
	value *BadgeBlockAllOf
	isSet bool
}

func (v NullableBadgeBlockAllOf) Get() *BadgeBlockAllOf {
	return v.value
}

func (v *NullableBadgeBlockAllOf) Set(val *BadgeBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBadgeBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBadgeBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadgeBlockAllOf(val *BadgeBlockAllOf) *NullableBadgeBlockAllOf {
	return &NullableBadgeBlockAllOf{value: val, isSet: true}
}

func (v NullableBadgeBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadgeBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


