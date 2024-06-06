// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SupportedLanguage ISO code for a supported language.
type SupportedLanguage string

// List of supportedLanguage.
const (
	SUPPORTEDLANGUAGE_AF    SupportedLanguage = "af"
	SUPPORTEDLANGUAGE_AR    SupportedLanguage = "ar"
	SUPPORTEDLANGUAGE_AZ    SupportedLanguage = "az"
	SUPPORTEDLANGUAGE_BG    SupportedLanguage = "bg"
	SUPPORTEDLANGUAGE_BN    SupportedLanguage = "bn"
	SUPPORTEDLANGUAGE_CA    SupportedLanguage = "ca"
	SUPPORTEDLANGUAGE_CS    SupportedLanguage = "cs"
	SUPPORTEDLANGUAGE_CY    SupportedLanguage = "cy"
	SUPPORTEDLANGUAGE_DA    SupportedLanguage = "da"
	SUPPORTEDLANGUAGE_DE    SupportedLanguage = "de"
	SUPPORTEDLANGUAGE_EL    SupportedLanguage = "el"
	SUPPORTEDLANGUAGE_EN    SupportedLanguage = "en"
	SUPPORTEDLANGUAGE_EO    SupportedLanguage = "eo"
	SUPPORTEDLANGUAGE_ES    SupportedLanguage = "es"
	SUPPORTEDLANGUAGE_ET    SupportedLanguage = "et"
	SUPPORTEDLANGUAGE_EU    SupportedLanguage = "eu"
	SUPPORTEDLANGUAGE_FA    SupportedLanguage = "fa"
	SUPPORTEDLANGUAGE_FI    SupportedLanguage = "fi"
	SUPPORTEDLANGUAGE_FO    SupportedLanguage = "fo"
	SUPPORTEDLANGUAGE_FR    SupportedLanguage = "fr"
	SUPPORTEDLANGUAGE_GA    SupportedLanguage = "ga"
	SUPPORTEDLANGUAGE_GL    SupportedLanguage = "gl"
	SUPPORTEDLANGUAGE_HE    SupportedLanguage = "he"
	SUPPORTEDLANGUAGE_HI    SupportedLanguage = "hi"
	SUPPORTEDLANGUAGE_HU    SupportedLanguage = "hu"
	SUPPORTEDLANGUAGE_HY    SupportedLanguage = "hy"
	SUPPORTEDLANGUAGE_ID    SupportedLanguage = "id"
	SUPPORTEDLANGUAGE_IS    SupportedLanguage = "is"
	SUPPORTEDLANGUAGE_IT    SupportedLanguage = "it"
	SUPPORTEDLANGUAGE_JA    SupportedLanguage = "ja"
	SUPPORTEDLANGUAGE_KA    SupportedLanguage = "ka"
	SUPPORTEDLANGUAGE_KK    SupportedLanguage = "kk"
	SUPPORTEDLANGUAGE_KO    SupportedLanguage = "ko"
	SUPPORTEDLANGUAGE_KU    SupportedLanguage = "ku"
	SUPPORTEDLANGUAGE_KY    SupportedLanguage = "ky"
	SUPPORTEDLANGUAGE_LT    SupportedLanguage = "lt"
	SUPPORTEDLANGUAGE_LV    SupportedLanguage = "lv"
	SUPPORTEDLANGUAGE_MI    SupportedLanguage = "mi"
	SUPPORTEDLANGUAGE_MN    SupportedLanguage = "mn"
	SUPPORTEDLANGUAGE_MR    SupportedLanguage = "mr"
	SUPPORTEDLANGUAGE_MS    SupportedLanguage = "ms"
	SUPPORTEDLANGUAGE_MT    SupportedLanguage = "mt"
	SUPPORTEDLANGUAGE_NB    SupportedLanguage = "nb"
	SUPPORTEDLANGUAGE_NL    SupportedLanguage = "nl"
	SUPPORTEDLANGUAGE_NO    SupportedLanguage = "no"
	SUPPORTEDLANGUAGE_NS    SupportedLanguage = "ns"
	SUPPORTEDLANGUAGE_PL    SupportedLanguage = "pl"
	SUPPORTEDLANGUAGE_PS    SupportedLanguage = "ps"
	SUPPORTEDLANGUAGE_PT    SupportedLanguage = "pt"
	SUPPORTEDLANGUAGE_PT_BR SupportedLanguage = "pt-br"
	SUPPORTEDLANGUAGE_QU    SupportedLanguage = "qu"
	SUPPORTEDLANGUAGE_RO    SupportedLanguage = "ro"
	SUPPORTEDLANGUAGE_RU    SupportedLanguage = "ru"
	SUPPORTEDLANGUAGE_SK    SupportedLanguage = "sk"
	SUPPORTEDLANGUAGE_SQ    SupportedLanguage = "sq"
	SUPPORTEDLANGUAGE_SV    SupportedLanguage = "sv"
	SUPPORTEDLANGUAGE_SW    SupportedLanguage = "sw"
	SUPPORTEDLANGUAGE_TA    SupportedLanguage = "ta"
	SUPPORTEDLANGUAGE_TE    SupportedLanguage = "te"
	SUPPORTEDLANGUAGE_TH    SupportedLanguage = "th"
	SUPPORTEDLANGUAGE_TL    SupportedLanguage = "tl"
	SUPPORTEDLANGUAGE_TN    SupportedLanguage = "tn"
	SUPPORTEDLANGUAGE_TR    SupportedLanguage = "tr"
	SUPPORTEDLANGUAGE_TT    SupportedLanguage = "tt"
	SUPPORTEDLANGUAGE_UK    SupportedLanguage = "uk"
	SUPPORTEDLANGUAGE_UR    SupportedLanguage = "ur"
	SUPPORTEDLANGUAGE_UZ    SupportedLanguage = "uz"
	SUPPORTEDLANGUAGE_ZH    SupportedLanguage = "zh"
)

// All allowed values of SupportedLanguage enum.
var AllowedSupportedLanguageEnumValues = []SupportedLanguage{
	"af",
	"ar",
	"az",
	"bg",
	"bn",
	"ca",
	"cs",
	"cy",
	"da",
	"de",
	"el",
	"en",
	"eo",
	"es",
	"et",
	"eu",
	"fa",
	"fi",
	"fo",
	"fr",
	"ga",
	"gl",
	"he",
	"hi",
	"hu",
	"hy",
	"id",
	"is",
	"it",
	"ja",
	"ka",
	"kk",
	"ko",
	"ku",
	"ky",
	"lt",
	"lv",
	"mi",
	"mn",
	"mr",
	"ms",
	"mt",
	"nb",
	"nl",
	"no",
	"ns",
	"pl",
	"ps",
	"pt",
	"pt-br",
	"qu",
	"ro",
	"ru",
	"sk",
	"sq",
	"sv",
	"sw",
	"ta",
	"te",
	"th",
	"tl",
	"tn",
	"tr",
	"tt",
	"uk",
	"ur",
	"uz",
	"zh",
}

func (v *SupportedLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'SupportedLanguage': %w", string(src), err)
	}
	enumTypeValue := SupportedLanguage(value)
	for _, existing := range AllowedSupportedLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportedLanguage", value)
}

// NewSupportedLanguageFromValue returns a pointer to a valid SupportedLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSupportedLanguageFromValue(v string) (*SupportedLanguage, error) {
	ev := SupportedLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportedLanguage: valid values are %v", v, AllowedSupportedLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SupportedLanguage) IsValid() bool {
	for _, existing := range AllowedSupportedLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to supportedLanguage value.
func (v SupportedLanguage) Ptr() *SupportedLanguage {
	return &v
}

type NullableSupportedLanguage struct {
	value *SupportedLanguage
	isSet bool
}

func (v NullableSupportedLanguage) Get() *SupportedLanguage {
	return v.value
}

func (v *NullableSupportedLanguage) Set(val *SupportedLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedLanguage(val *SupportedLanguage) *NullableSupportedLanguage {
	return &NullableSupportedLanguage{value: val, isSet: true}
}

func (v NullableSupportedLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSupportedLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
