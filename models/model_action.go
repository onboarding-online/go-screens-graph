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

// Action Parameters for action
type Action struct {
	// List of conditioned actions
	Edges []ConditionedAction `json:"edges"`
	Kind ActionKind `json:"kind"`
}

// NewAction instantiates a new Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAction(edges []ConditionedAction, kind ActionKind) *Action {
	this := Action{}
	this.Edges = edges
	this.Kind = kind
	return &this
}

// NewActionWithDefaults instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefaults() *Action {
	this := Action{}
	return &this
}

// GetEdges returns the Edges field value
func (o *Action) GetEdges() []ConditionedAction {
	if o == nil {
		var ret []ConditionedAction
		return ret
	}

	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value
// and a boolean to check if the value has been set.
func (o *Action) GetEdgesOk() ([]ConditionedAction, bool) {
	if o == nil {
    return nil, false
	}
	return o.Edges, true
}

// SetEdges sets field value
func (o *Action) SetEdges(v []ConditionedAction) {
	o.Edges = v
}

// GetKind returns the Kind field value
func (o *Action) GetKind() ActionKind {
	if o == nil {
		var ret ActionKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Action) GetKindOk() (*ActionKind, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Action) SetKind(v ActionKind) {
	o.Kind = v
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["edges"] = o.Edges
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


