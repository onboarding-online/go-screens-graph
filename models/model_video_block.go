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

// VideoBlock Styles for video
type VideoBlock struct {
	ScaleMode *VideoScaleMode `json:"scaleMode,omitempty"`
	// width of video
	Width *float32 `json:"width,omitempty"`
	// height of video
	Height *float32 `json:"height,omitempty"`
	// Height in percentage
	HeightPercentage *float32 `json:"heightPercentage,omitempty"`
	CornerRadiusLeftTop *float32 `json:"cornerRadiusLeftTop,omitempty"`
	CornerRadiusLeftBottom *float32 `json:"cornerRadiusLeftBottom,omitempty"`
	CornerRadiusRightTop *float32 `json:"cornerRadiusRightTop,omitempty"`
	CornerRadiusRightBottom *float32 `json:"cornerRadiusRightBottom,omitempty"`
	// apply corner radius for all corners
	MainCornerRadius *float32 `json:"mainCornerRadius,omitempty"`
	// repeat video after finish
	Repeat *bool `json:"repeat,omitempty"`
}

// NewVideoBlock instantiates a new VideoBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoBlock() *VideoBlock {
	this := VideoBlock{}
	return &this
}

// NewVideoBlockWithDefaults instantiates a new VideoBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoBlockWithDefaults() *VideoBlock {
	this := VideoBlock{}
	return &this
}

// GetScaleMode returns the ScaleMode field value if set, zero value otherwise.
func (o *VideoBlock) GetScaleMode() VideoScaleMode {
	if o == nil || isNil(o.ScaleMode) {
		var ret VideoScaleMode
		return ret
	}
	return *o.ScaleMode
}

// GetScaleModeOk returns a tuple with the ScaleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetScaleModeOk() (*VideoScaleMode, bool) {
	if o == nil || isNil(o.ScaleMode) {
    return nil, false
	}
	return o.ScaleMode, true
}

// HasScaleMode returns a boolean if a field has been set.
func (o *VideoBlock) HasScaleMode() bool {
	if o != nil && !isNil(o.ScaleMode) {
		return true
	}

	return false
}

// SetScaleMode gets a reference to the given VideoScaleMode and assigns it to the ScaleMode field.
func (o *VideoBlock) SetScaleMode(v VideoScaleMode) {
	o.ScaleMode = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *VideoBlock) GetWidth() float32 {
	if o == nil || isNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetWidthOk() (*float32, bool) {
	if o == nil || isNil(o.Width) {
    return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *VideoBlock) HasWidth() bool {
	if o != nil && !isNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *VideoBlock) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *VideoBlock) GetHeight() float32 {
	if o == nil || isNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetHeightOk() (*float32, bool) {
	if o == nil || isNil(o.Height) {
    return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *VideoBlock) HasHeight() bool {
	if o != nil && !isNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *VideoBlock) SetHeight(v float32) {
	o.Height = &v
}

// GetHeightPercentage returns the HeightPercentage field value if set, zero value otherwise.
func (o *VideoBlock) GetHeightPercentage() float32 {
	if o == nil || isNil(o.HeightPercentage) {
		var ret float32
		return ret
	}
	return *o.HeightPercentage
}

// GetHeightPercentageOk returns a tuple with the HeightPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetHeightPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.HeightPercentage) {
    return nil, false
	}
	return o.HeightPercentage, true
}

// HasHeightPercentage returns a boolean if a field has been set.
func (o *VideoBlock) HasHeightPercentage() bool {
	if o != nil && !isNil(o.HeightPercentage) {
		return true
	}

	return false
}

// SetHeightPercentage gets a reference to the given float32 and assigns it to the HeightPercentage field.
func (o *VideoBlock) SetHeightPercentage(v float32) {
	o.HeightPercentage = &v
}

// GetCornerRadiusLeftTop returns the CornerRadiusLeftTop field value if set, zero value otherwise.
func (o *VideoBlock) GetCornerRadiusLeftTop() float32 {
	if o == nil || isNil(o.CornerRadiusLeftTop) {
		var ret float32
		return ret
	}
	return *o.CornerRadiusLeftTop
}

// GetCornerRadiusLeftTopOk returns a tuple with the CornerRadiusLeftTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetCornerRadiusLeftTopOk() (*float32, bool) {
	if o == nil || isNil(o.CornerRadiusLeftTop) {
    return nil, false
	}
	return o.CornerRadiusLeftTop, true
}

// HasCornerRadiusLeftTop returns a boolean if a field has been set.
func (o *VideoBlock) HasCornerRadiusLeftTop() bool {
	if o != nil && !isNil(o.CornerRadiusLeftTop) {
		return true
	}

	return false
}

// SetCornerRadiusLeftTop gets a reference to the given float32 and assigns it to the CornerRadiusLeftTop field.
func (o *VideoBlock) SetCornerRadiusLeftTop(v float32) {
	o.CornerRadiusLeftTop = &v
}

// GetCornerRadiusLeftBottom returns the CornerRadiusLeftBottom field value if set, zero value otherwise.
func (o *VideoBlock) GetCornerRadiusLeftBottom() float32 {
	if o == nil || isNil(o.CornerRadiusLeftBottom) {
		var ret float32
		return ret
	}
	return *o.CornerRadiusLeftBottom
}

// GetCornerRadiusLeftBottomOk returns a tuple with the CornerRadiusLeftBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetCornerRadiusLeftBottomOk() (*float32, bool) {
	if o == nil || isNil(o.CornerRadiusLeftBottom) {
    return nil, false
	}
	return o.CornerRadiusLeftBottom, true
}

// HasCornerRadiusLeftBottom returns a boolean if a field has been set.
func (o *VideoBlock) HasCornerRadiusLeftBottom() bool {
	if o != nil && !isNil(o.CornerRadiusLeftBottom) {
		return true
	}

	return false
}

// SetCornerRadiusLeftBottom gets a reference to the given float32 and assigns it to the CornerRadiusLeftBottom field.
func (o *VideoBlock) SetCornerRadiusLeftBottom(v float32) {
	o.CornerRadiusLeftBottom = &v
}

// GetCornerRadiusRightTop returns the CornerRadiusRightTop field value if set, zero value otherwise.
func (o *VideoBlock) GetCornerRadiusRightTop() float32 {
	if o == nil || isNil(o.CornerRadiusRightTop) {
		var ret float32
		return ret
	}
	return *o.CornerRadiusRightTop
}

// GetCornerRadiusRightTopOk returns a tuple with the CornerRadiusRightTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetCornerRadiusRightTopOk() (*float32, bool) {
	if o == nil || isNil(o.CornerRadiusRightTop) {
    return nil, false
	}
	return o.CornerRadiusRightTop, true
}

// HasCornerRadiusRightTop returns a boolean if a field has been set.
func (o *VideoBlock) HasCornerRadiusRightTop() bool {
	if o != nil && !isNil(o.CornerRadiusRightTop) {
		return true
	}

	return false
}

// SetCornerRadiusRightTop gets a reference to the given float32 and assigns it to the CornerRadiusRightTop field.
func (o *VideoBlock) SetCornerRadiusRightTop(v float32) {
	o.CornerRadiusRightTop = &v
}

// GetCornerRadiusRightBottom returns the CornerRadiusRightBottom field value if set, zero value otherwise.
func (o *VideoBlock) GetCornerRadiusRightBottom() float32 {
	if o == nil || isNil(o.CornerRadiusRightBottom) {
		var ret float32
		return ret
	}
	return *o.CornerRadiusRightBottom
}

// GetCornerRadiusRightBottomOk returns a tuple with the CornerRadiusRightBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetCornerRadiusRightBottomOk() (*float32, bool) {
	if o == nil || isNil(o.CornerRadiusRightBottom) {
    return nil, false
	}
	return o.CornerRadiusRightBottom, true
}

// HasCornerRadiusRightBottom returns a boolean if a field has been set.
func (o *VideoBlock) HasCornerRadiusRightBottom() bool {
	if o != nil && !isNil(o.CornerRadiusRightBottom) {
		return true
	}

	return false
}

// SetCornerRadiusRightBottom gets a reference to the given float32 and assigns it to the CornerRadiusRightBottom field.
func (o *VideoBlock) SetCornerRadiusRightBottom(v float32) {
	o.CornerRadiusRightBottom = &v
}

// GetMainCornerRadius returns the MainCornerRadius field value if set, zero value otherwise.
func (o *VideoBlock) GetMainCornerRadius() float32 {
	if o == nil || isNil(o.MainCornerRadius) {
		var ret float32
		return ret
	}
	return *o.MainCornerRadius
}

// GetMainCornerRadiusOk returns a tuple with the MainCornerRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetMainCornerRadiusOk() (*float32, bool) {
	if o == nil || isNil(o.MainCornerRadius) {
    return nil, false
	}
	return o.MainCornerRadius, true
}

// HasMainCornerRadius returns a boolean if a field has been set.
func (o *VideoBlock) HasMainCornerRadius() bool {
	if o != nil && !isNil(o.MainCornerRadius) {
		return true
	}

	return false
}

// SetMainCornerRadius gets a reference to the given float32 and assigns it to the MainCornerRadius field.
func (o *VideoBlock) SetMainCornerRadius(v float32) {
	o.MainCornerRadius = &v
}

// GetRepeat returns the Repeat field value if set, zero value otherwise.
func (o *VideoBlock) GetRepeat() bool {
	if o == nil || isNil(o.Repeat) {
		var ret bool
		return ret
	}
	return *o.Repeat
}

// GetRepeatOk returns a tuple with the Repeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoBlock) GetRepeatOk() (*bool, bool) {
	if o == nil || isNil(o.Repeat) {
    return nil, false
	}
	return o.Repeat, true
}

// HasRepeat returns a boolean if a field has been set.
func (o *VideoBlock) HasRepeat() bool {
	if o != nil && !isNil(o.Repeat) {
		return true
	}

	return false
}

// SetRepeat gets a reference to the given bool and assigns it to the Repeat field.
func (o *VideoBlock) SetRepeat(v bool) {
	o.Repeat = &v
}

func (o VideoBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ScaleMode) {
		toSerialize["scaleMode"] = o.ScaleMode
	}
	if !isNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !isNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !isNil(o.HeightPercentage) {
		toSerialize["heightPercentage"] = o.HeightPercentage
	}
	if !isNil(o.CornerRadiusLeftTop) {
		toSerialize["cornerRadiusLeftTop"] = o.CornerRadiusLeftTop
	}
	if !isNil(o.CornerRadiusLeftBottom) {
		toSerialize["cornerRadiusLeftBottom"] = o.CornerRadiusLeftBottom
	}
	if !isNil(o.CornerRadiusRightTop) {
		toSerialize["cornerRadiusRightTop"] = o.CornerRadiusRightTop
	}
	if !isNil(o.CornerRadiusRightBottom) {
		toSerialize["cornerRadiusRightBottom"] = o.CornerRadiusRightBottom
	}
	if !isNil(o.MainCornerRadius) {
		toSerialize["mainCornerRadius"] = o.MainCornerRadius
	}
	if !isNil(o.Repeat) {
		toSerialize["repeat"] = o.Repeat
	}
	return json.Marshal(toSerialize)
}

type NullableVideoBlock struct {
	value *VideoBlock
	isSet bool
}

func (v NullableVideoBlock) Get() *VideoBlock {
	return v.value
}

func (v *NullableVideoBlock) Set(val *VideoBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoBlock(val *VideoBlock) *NullableVideoBlock {
	return &NullableVideoBlock{value: val, isSet: true}
}

func (v NullableVideoBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


