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

// Language language codes in ISO 639-1 and ISO 3166-1 standards
type Language string

// List of Language
const (
	LANGUAGE_AR Language = "ar"
	LANGUAGE_CA Language = "ca"
	LANGUAGE_ZH_HANS Language = "zh-Hans"
	LANGUAGE_ZH_HANT Language = "zh-Hant"
	LANGUAGE_HR Language = "hr"
	LANGUAGE_CS Language = "cs"
	LANGUAGE_DA Language = "da"
	LANGUAGE_NL Language = "nl"
	LANGUAGE_EN_AU Language = "en-AU"
	LANGUAGE_EN_CA Language = "en-CA"
	LANGUAGE_EN_GB Language = "en-GB"
	LANGUAGE_EN_US Language = "en-US"
	LANGUAGE_FI Language = "fi"
	LANGUAGE_FR Language = "fr"
	LANGUAGE_FR_CA Language = "fr-CA"
	LANGUAGE_DE Language = "de"
	LANGUAGE_EL Language = "el"
	LANGUAGE_HE Language = "he"
	LANGUAGE_HI Language = "hi"
	LANGUAGE_HU Language = "hu"
	LANGUAGE_ID Language = "id"
	LANGUAGE_IT Language = "it"
	LANGUAGE_JA Language = "ja"
	LANGUAGE_KO Language = "ko"
	LANGUAGE_MS Language = "ms"
	LANGUAGE_NO Language = "no"
	LANGUAGE_PL Language = "pl"
	LANGUAGE_PT_BR Language = "pt-BR"
	LANGUAGE_PT_PT Language = "pt-PT"
	LANGUAGE_RO Language = "ro"
	LANGUAGE_RU Language = "ru"
	LANGUAGE_SK Language = "sk"
	LANGUAGE_ES_MX Language = "es-MX"
	LANGUAGE_ES_ES Language = "es-ES"
	LANGUAGE_SV Language = "sv"
	LANGUAGE_TH Language = "th"
	LANGUAGE_TR Language = "tr"
	LANGUAGE_UK Language = "uk"
	LANGUAGE_VI Language = "vi"
)

// All allowed values of Language enum
var AllowedLanguageEnumValues = []Language{
	"ar",
	"ca",
	"zh-Hans",
	"zh-Hant",
	"hr",
	"cs",
	"da",
	"nl",
	"en-AU",
	"en-CA",
	"en-GB",
	"en-US",
	"fi",
	"fr",
	"fr-CA",
	"de",
	"el",
	"he",
	"hi",
	"hu",
	"id",
	"it",
	"ja",
	"ko",
	"ms",
	"no",
	"pl",
	"pt-BR",
	"pt-PT",
	"ro",
	"ru",
	"sk",
	"es-MX",
	"es-ES",
	"sv",
	"th",
	"tr",
	"uk",
	"vi",
}

func (v *Language) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Language(value)
	for _, existing := range AllowedLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Language", value)
}

// NewLanguageFromValue returns a pointer to a valid Language
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLanguageFromValue(v string) (*Language, error) {
	ev := Language(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Language: valid values are %v", v, AllowedLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Language) IsValid() bool {
	for _, existing := range AllowedLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Language value
func (v Language) Ptr() *Language {
	return &v
}

type NullableLanguage struct {
	value *Language
	isSet bool
}

func (v NullableLanguage) Get() *Language {
	return v.value
}

func (v *NullableLanguage) Set(val *Language) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguage(val *Language) *NullableLanguage {
	return &NullableLanguage{value: val, isSet: true}
}

func (v NullableLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

