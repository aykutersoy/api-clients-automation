// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// OperationType Operation to perform on the index.
type OperationType string

// List of operationType.
const (
	OPERATIONTYPE_MOVE OperationType = "move"
	OPERATIONTYPE_COPY OperationType = "copy"
)

// All allowed values of OperationType enum.
var AllowedOperationTypeEnumValues = []OperationType{
	"move",
	"copy",
}

func (v *OperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'OperationType': %w", string(src), err)
	}
	enumTypeValue := OperationType(value)
	for _, existing := range AllowedOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationType", value)
}

// NewOperationTypeFromValue returns a pointer to a valid OperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOperationTypeFromValue(v string) (*OperationType, error) {
	ev := OperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationType: valid values are %v", v, AllowedOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OperationType) IsValid() bool {
	for _, existing := range AllowedOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to operationType value.
func (v OperationType) Ptr() *OperationType {
	return &v
}

type NullableOperationType struct {
	value *OperationType
	isSet bool
}

func (v NullableOperationType) Get() *OperationType {
	return v.value
}

func (v *NullableOperationType) Set(val *OperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationType(val *OperationType) *NullableOperationType {
	return &NullableOperationType{value: val, isSet: true}
}

func (v NullableOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
