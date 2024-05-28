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

// CalendarDays Calendar days config
type CalendarDays struct {
	Past CalendarDay `json:"past"`
	Today CalendarDay `json:"today"`
	Future CalendarDay `json:"future"`
	Weekend CalendarDay `json:"weekend"`
}

// NewCalendarDays instantiates a new CalendarDays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarDays(past CalendarDay, today CalendarDay, future CalendarDay, weekend CalendarDay) *CalendarDays {
	this := CalendarDays{}
	this.Past = past
	this.Today = today
	this.Future = future
	this.Weekend = weekend
	return &this
}

// NewCalendarDaysWithDefaults instantiates a new CalendarDays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarDaysWithDefaults() *CalendarDays {
	this := CalendarDays{}
	return &this
}

// GetPast returns the Past field value
func (o *CalendarDays) GetPast() CalendarDay {
	if o == nil {
		var ret CalendarDay
		return ret
	}

	return o.Past
}

// GetPastOk returns a tuple with the Past field value
// and a boolean to check if the value has been set.
func (o *CalendarDays) GetPastOk() (*CalendarDay, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Past, true
}

// SetPast sets field value
func (o *CalendarDays) SetPast(v CalendarDay) {
	o.Past = v
}

// GetToday returns the Today field value
func (o *CalendarDays) GetToday() CalendarDay {
	if o == nil {
		var ret CalendarDay
		return ret
	}

	return o.Today
}

// GetTodayOk returns a tuple with the Today field value
// and a boolean to check if the value has been set.
func (o *CalendarDays) GetTodayOk() (*CalendarDay, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Today, true
}

// SetToday sets field value
func (o *CalendarDays) SetToday(v CalendarDay) {
	o.Today = v
}

// GetFuture returns the Future field value
func (o *CalendarDays) GetFuture() CalendarDay {
	if o == nil {
		var ret CalendarDay
		return ret
	}

	return o.Future
}

// GetFutureOk returns a tuple with the Future field value
// and a boolean to check if the value has been set.
func (o *CalendarDays) GetFutureOk() (*CalendarDay, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Future, true
}

// SetFuture sets field value
func (o *CalendarDays) SetFuture(v CalendarDay) {
	o.Future = v
}

// GetWeekend returns the Weekend field value
func (o *CalendarDays) GetWeekend() CalendarDay {
	if o == nil {
		var ret CalendarDay
		return ret
	}

	return o.Weekend
}

// GetWeekendOk returns a tuple with the Weekend field value
// and a boolean to check if the value has been set.
func (o *CalendarDays) GetWeekendOk() (*CalendarDay, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Weekend, true
}

// SetWeekend sets field value
func (o *CalendarDays) SetWeekend(v CalendarDay) {
	o.Weekend = v
}

func (o CalendarDays) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["past"] = o.Past
	}
	if true {
		toSerialize["today"] = o.Today
	}
	if true {
		toSerialize["future"] = o.Future
	}
	if true {
		toSerialize["weekend"] = o.Weekend
	}
	return json.Marshal(toSerialize)
}

type NullableCalendarDays struct {
	value *CalendarDays
	isSet bool
}

func (v NullableCalendarDays) Get() *CalendarDays {
	return v.value
}

func (v *NullableCalendarDays) Set(val *CalendarDays) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarDays) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarDays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarDays(val *CalendarDays) *NullableCalendarDays {
	return &NullableCalendarDays{value: val, isSet: true}
}

func (v NullableCalendarDays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarDays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


