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

// PaywallFooter Paywall footer
type PaywallFooter struct {
	PaywallFooter string `json:"paywallFooter"`
	Purchase *Button `json:"purchase,omitempty"`
	AutoRenewLabel *Text `json:"autoRenewLabel,omitempty"`
	BottomContainer PaywallFooterBottomContainer `json:"bottomContainer"`
	Styles PaywallFooterBlock `json:"styles"`
	Button1 *Button `json:"button1,omitempty"`
	Button2 *Button `json:"button2,omitempty"`
	Kind *BasicFooterKind `json:"kind,omitempty"`
}

// NewPaywallFooter instantiates a new PaywallFooter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaywallFooter(paywallFooter string, bottomContainer PaywallFooterBottomContainer, styles PaywallFooterBlock) *PaywallFooter {
	this := PaywallFooter{}
	this.Styles = styles
	return &this
}

// NewPaywallFooterWithDefaults instantiates a new PaywallFooter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaywallFooterWithDefaults() *PaywallFooter {
	this := PaywallFooter{}
	return &this
}

// GetPaywallFooter returns the PaywallFooter field value
func (o *PaywallFooter) GetPaywallFooter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaywallFooter
}

// GetPaywallFooterOk returns a tuple with the PaywallFooter field value
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetPaywallFooterOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PaywallFooter, true
}

// SetPaywallFooter sets field value
func (o *PaywallFooter) SetPaywallFooter(v string) {
	o.PaywallFooter = v
}

// GetPurchase returns the Purchase field value if set, zero value otherwise.
func (o *PaywallFooter) GetPurchase() Button {
	if o == nil || isNil(o.Purchase) {
		var ret Button
		return ret
	}
	return *o.Purchase
}

// GetPurchaseOk returns a tuple with the Purchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetPurchaseOk() (*Button, bool) {
	if o == nil || isNil(o.Purchase) {
    return nil, false
	}
	return o.Purchase, true
}

// HasPurchase returns a boolean if a field has been set.
func (o *PaywallFooter) HasPurchase() bool {
	if o != nil && !isNil(o.Purchase) {
		return true
	}

	return false
}

// SetPurchase gets a reference to the given Button and assigns it to the Purchase field.
func (o *PaywallFooter) SetPurchase(v Button) {
	o.Purchase = &v
}

// GetAutoRenewLabel returns the AutoRenewLabel field value if set, zero value otherwise.
func (o *PaywallFooter) GetAutoRenewLabel() Text {
	if o == nil || isNil(o.AutoRenewLabel) {
		var ret Text
		return ret
	}
	return *o.AutoRenewLabel
}

// GetAutoRenewLabelOk returns a tuple with the AutoRenewLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetAutoRenewLabelOk() (*Text, bool) {
	if o == nil || isNil(o.AutoRenewLabel) {
    return nil, false
	}
	return o.AutoRenewLabel, true
}

// HasAutoRenewLabel returns a boolean if a field has been set.
func (o *PaywallFooter) HasAutoRenewLabel() bool {
	if o != nil && !isNil(o.AutoRenewLabel) {
		return true
	}

	return false
}

// SetAutoRenewLabel gets a reference to the given Text and assigns it to the AutoRenewLabel field.
func (o *PaywallFooter) SetAutoRenewLabel(v Text) {
	o.AutoRenewLabel = &v
}

// GetBottomContainer returns the BottomContainer field value
func (o *PaywallFooter) GetBottomContainer() PaywallFooterBottomContainer {
	if o == nil {
		var ret PaywallFooterBottomContainer
		return ret
	}

	return o.BottomContainer
}

// GetBottomContainerOk returns a tuple with the BottomContainer field value
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetBottomContainerOk() (*PaywallFooterBottomContainer, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BottomContainer, true
}

// SetBottomContainer sets field value
func (o *PaywallFooter) SetBottomContainer(v PaywallFooterBottomContainer) {
	o.BottomContainer = v
}

// GetStyles returns the Styles field value
func (o *PaywallFooter) GetStyles() PaywallFooterBlock {
	if o == nil {
		var ret PaywallFooterBlock
		return ret
	}

	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetStylesOk() (*PaywallFooterBlock, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Styles, true
}

// SetStyles sets field value
func (o *PaywallFooter) SetStyles(v PaywallFooterBlock) {
	o.Styles = v
}

// GetButton1 returns the Button1 field value if set, zero value otherwise.
func (o *PaywallFooter) GetButton1() Button {
	if o == nil || isNil(o.Button1) {
		var ret Button
		return ret
	}
	return *o.Button1
}

// GetButton1Ok returns a tuple with the Button1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetButton1Ok() (*Button, bool) {
	if o == nil || isNil(o.Button1) {
    return nil, false
	}
	return o.Button1, true
}

// HasButton1 returns a boolean if a field has been set.
func (o *PaywallFooter) HasButton1() bool {
	if o != nil && !isNil(o.Button1) {
		return true
	}

	return false
}

// SetButton1 gets a reference to the given Button and assigns it to the Button1 field.
func (o *PaywallFooter) SetButton1(v Button) {
	o.Button1 = &v
}

// GetButton2 returns the Button2 field value if set, zero value otherwise.
func (o *PaywallFooter) GetButton2() Button {
	if o == nil || isNil(o.Button2) {
		var ret Button
		return ret
	}
	return *o.Button2
}

// GetButton2Ok returns a tuple with the Button2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetButton2Ok() (*Button, bool) {
	if o == nil || isNil(o.Button2) {
    return nil, false
	}
	return o.Button2, true
}

// HasButton2 returns a boolean if a field has been set.
func (o *PaywallFooter) HasButton2() bool {
	if o != nil && !isNil(o.Button2) {
		return true
	}

	return false
}

// SetButton2 gets a reference to the given Button and assigns it to the Button2 field.
func (o *PaywallFooter) SetButton2(v Button) {
	o.Button2 = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *PaywallFooter) GetKind() BasicFooterKind {
	if o == nil || isNil(o.Kind) {
		var ret BasicFooterKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaywallFooter) GetKindOk() (*BasicFooterKind, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *PaywallFooter) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given BasicFooterKind and assigns it to the Kind field.
func (o *PaywallFooter) SetKind(v BasicFooterKind) {
	o.Kind = &v
}

func (o PaywallFooter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paywallFooter"] = o.PaywallFooter
	}
	if !isNil(o.Purchase) {
		toSerialize["purchase"] = o.Purchase
	}
	if !isNil(o.AutoRenewLabel) {
		toSerialize["autoRenewLabel"] = o.AutoRenewLabel
	}
	if true {
		toSerialize["bottomContainer"] = o.BottomContainer
	}
	if true {
		toSerialize["styles"] = o.Styles
	}
	if !isNil(o.Button1) {
		toSerialize["button1"] = o.Button1
	}
	if !isNil(o.Button2) {
		toSerialize["button2"] = o.Button2
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullablePaywallFooter struct {
	value *PaywallFooter
	isSet bool
}

func (v NullablePaywallFooter) Get() *PaywallFooter {
	return v.value
}

func (v *NullablePaywallFooter) Set(val *PaywallFooter) {
	v.value = val
	v.isSet = true
}

func (v NullablePaywallFooter) IsSet() bool {
	return v.isSet
}

func (v *NullablePaywallFooter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaywallFooter(val *PaywallFooter) *NullablePaywallFooter {
	return &NullablePaywallFooter{value: val, isSet: true}
}

func (v NullablePaywallFooter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaywallFooter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


