// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// RunReasonCode A code for the task run's outcome. A readable description of the code is included in the `reason` response property.
type RunReasonCode string

// List of RunReasonCode.
const (
	RUNREASONCODE_INTERNAL        RunReasonCode = "internal"
	RUNREASONCODE_CRITICAL        RunReasonCode = "critical"
	RUNREASONCODE_NO_EVENTS       RunReasonCode = "no_events"
	RUNREASONCODE_TOO_MANY_ERRORS RunReasonCode = "too_many_errors"
	RUNREASONCODE_OK              RunReasonCode = "ok"
	RUNREASONCODE_DISCARDED       RunReasonCode = "discarded"
	RUNREASONCODE_BLOCKING        RunReasonCode = "blocking"
)

// All allowed values of RunReasonCode enum.
var AllowedRunReasonCodeEnumValues = []RunReasonCode{
	"internal",
	"critical",
	"no_events",
	"too_many_errors",
	"ok",
	"discarded",
	"blocking",
}

func (v *RunReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'RunReasonCode': %w", string(src), err)
	}
	enumTypeValue := RunReasonCode(value)
	for _, existing := range AllowedRunReasonCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RunReasonCode", value)
}

// NewRunReasonCodeFromValue returns a pointer to a valid RunReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunReasonCodeFromValue(v string) (*RunReasonCode, error) {
	ev := RunReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RunReasonCode: valid values are %v", v, AllowedRunReasonCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunReasonCode) IsValid() bool {
	for _, existing := range AllowedRunReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunReasonCode value.
func (v RunReasonCode) Ptr() *RunReasonCode {
	return &v
}

type NullableRunReasonCode struct {
	value *RunReasonCode
	isSet bool
}

func (v NullableRunReasonCode) Get() *RunReasonCode {
	return v.value
}

func (v *NullableRunReasonCode) Set(val *RunReasonCode) {
	v.value = val
	v.isSet = true
}

func (v NullableRunReasonCode) IsSet() bool {
	return v.isSet
}

func (v *NullableRunReasonCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunReasonCode(val *RunReasonCode) *NullableRunReasonCode {
	return &NullableRunReasonCode{value: val, isSet: true}
}

func (v NullableRunReasonCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableRunReasonCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
