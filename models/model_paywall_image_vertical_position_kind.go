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

// PaywallImageVerticalPositionKind Paywall image vertical position kind
type PaywallImageVerticalPositionKind string

// List of PaywallImageVerticalPositionKind
const (
	PAYWALLIMAGEVERTICALPOSITIONKIND_HEADER_TITLE PaywallImageVerticalPositionKind = "HeaderTitle"
	PAYWALLIMAGEVERTICALPOSITIONKIND_HEADER_SUBTITLE PaywallImageVerticalPositionKind = "HeaderSubtitle"
	PAYWALLIMAGEVERTICALPOSITIONKIND_HEADER_LIST_TOP PaywallImageVerticalPositionKind = "HeaderListTop"
	PAYWALLIMAGEVERTICALPOSITIONKIND_HEADER_LIST_BOTTOM PaywallImageVerticalPositionKind = "HeaderListBottom"
	PAYWALLIMAGEVERTICALPOSITIONKIND_HEADER_BOTTOM PaywallImageVerticalPositionKind = "HeaderBottom"
	PAYWALLIMAGEVERTICALPOSITIONKIND_HEADER_CENTER PaywallImageVerticalPositionKind = "HeaderCenter"
	PAYWALLIMAGEVERTICALPOSITIONKIND_HEADER_TOP PaywallImageVerticalPositionKind = "HeaderTop"
	PAYWALLIMAGEVERTICALPOSITIONKIND_BODY_BOTTOM PaywallImageVerticalPositionKind = "BodyBottom"
	PAYWALLIMAGEVERTICALPOSITIONKIND_FOOTER_TOP PaywallImageVerticalPositionKind = "FooterTop"
	PAYWALLIMAGEVERTICALPOSITIONKIND_PAYWALL_IMAGE_VERTICAL_POSITION_KIND1 PaywallImageVerticalPositionKind = "PaywallImageVerticalPositionKind1"
	PAYWALLIMAGEVERTICALPOSITIONKIND_PAYWALL_IMAGE_VERTICAL_POSITION_KIND2 PaywallImageVerticalPositionKind = "PaywallImageVerticalPositionKind2"
	PAYWALLIMAGEVERTICALPOSITIONKIND_PAYWALL_IMAGE_VERTICAL_POSITION_KIND3 PaywallImageVerticalPositionKind = "PaywallImageVerticalPositionKind3"
	PAYWALLIMAGEVERTICALPOSITIONKIND_PAYWALL_IMAGE_VERTICAL_POSITION_KIND4 PaywallImageVerticalPositionKind = "PaywallImageVerticalPositionKind4"
	PAYWALLIMAGEVERTICALPOSITIONKIND_PAYWALL_IMAGE_VERTICAL_POSITION_KIND5 PaywallImageVerticalPositionKind = "PaywallImageVerticalPositionKind5"
)

// All allowed values of PaywallImageVerticalPositionKind enum
var AllowedPaywallImageVerticalPositionKindEnumValues = []PaywallImageVerticalPositionKind{
	"HeaderTitle",
	"HeaderSubtitle",
	"HeaderListTop",
	"HeaderListBottom",
	"HeaderBottom",
	"HeaderCenter",
	"HeaderTop",
	"BodyBottom",
	"FooterTop",
	"PaywallImageVerticalPositionKind1",
	"PaywallImageVerticalPositionKind2",
	"PaywallImageVerticalPositionKind3",
	"PaywallImageVerticalPositionKind4",
	"PaywallImageVerticalPositionKind5",
}

func (v *PaywallImageVerticalPositionKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaywallImageVerticalPositionKind(value)
	for _, existing := range AllowedPaywallImageVerticalPositionKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaywallImageVerticalPositionKind", value)
}

// NewPaywallImageVerticalPositionKindFromValue returns a pointer to a valid PaywallImageVerticalPositionKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaywallImageVerticalPositionKindFromValue(v string) (*PaywallImageVerticalPositionKind, error) {
	ev := PaywallImageVerticalPositionKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaywallImageVerticalPositionKind: valid values are %v", v, AllowedPaywallImageVerticalPositionKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaywallImageVerticalPositionKind) IsValid() bool {
	for _, existing := range AllowedPaywallImageVerticalPositionKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaywallImageVerticalPositionKind value
func (v PaywallImageVerticalPositionKind) Ptr() *PaywallImageVerticalPositionKind {
	return &v
}

type NullablePaywallImageVerticalPositionKind struct {
	value *PaywallImageVerticalPositionKind
	isSet bool
}

func (v NullablePaywallImageVerticalPositionKind) Get() *PaywallImageVerticalPositionKind {
	return v.value
}

func (v *NullablePaywallImageVerticalPositionKind) Set(val *PaywallImageVerticalPositionKind) {
	v.value = val
	v.isSet = true
}

func (v NullablePaywallImageVerticalPositionKind) IsSet() bool {
	return v.isSet
}

func (v *NullablePaywallImageVerticalPositionKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaywallImageVerticalPositionKind(val *PaywallImageVerticalPositionKind) *NullablePaywallImageVerticalPositionKind {
	return &NullablePaywallImageVerticalPositionKind{value: val, isSet: true}
}

func (v NullablePaywallImageVerticalPositionKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaywallImageVerticalPositionKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

