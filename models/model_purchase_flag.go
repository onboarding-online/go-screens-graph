/*
Onboarding online screens graph models

Onboarding online screens graph models and interfaces

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// PurchaseFlag Purchase flag
type PurchaseFlag string

// List of PurchaseFlag
const (
	PURCHASEFLAG_SHOULD_CLOSE_AFTER_RESTORE PurchaseFlag = "ShouldCloseAfterRestore"
	PURCHASEFLAG_SHOULD_CLOSE_AFTER_SUBSCRIPTION_ERROR PurchaseFlag = "ShouldCloseAfterSubscriptionError"
)

// All allowed values of PurchaseFlag enum
var AllowedPurchaseFlagEnumValues = []PurchaseFlag{
	"ShouldCloseAfterRestore",
	"ShouldCloseAfterSubscriptionError",
}

func (v *PurchaseFlag) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PurchaseFlag(value)
	for _, existing := range AllowedPurchaseFlagEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PurchaseFlag", value)
}

// NewPurchaseFlagFromValue returns a pointer to a valid PurchaseFlag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPurchaseFlagFromValue(v string) (*PurchaseFlag, error) {
	ev := PurchaseFlag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PurchaseFlag: valid values are %v", v, AllowedPurchaseFlagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PurchaseFlag) IsValid() bool {
	for _, existing := range AllowedPurchaseFlagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PurchaseFlag value
func (v PurchaseFlag) Ptr() *PurchaseFlag {
	return &v
}

type NullablePurchaseFlag struct {
	value *PurchaseFlag
	isSet bool
}

func (v NullablePurchaseFlag) Get() *PurchaseFlag {
	return v.value
}

func (v *NullablePurchaseFlag) Set(val *PurchaseFlag) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseFlag) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseFlag(val *PurchaseFlag) *NullablePurchaseFlag {
	return &NullablePurchaseFlag{value: val, isSet: true}
}

func (v NullablePurchaseFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

