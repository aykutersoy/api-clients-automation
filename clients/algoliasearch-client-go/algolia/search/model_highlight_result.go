// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// HighlightResult - struct for HighlightResult.
type HighlightResult struct {
	HighlightResultOption               *HighlightResultOption
	ArrayOfHighlightResultOption        *[]HighlightResultOption
	MapmapOfStringHighlightResultOption *map[string]HighlightResultOption
}

// HighlightResultOptionAsHighlightResult is a convenience function that returns HighlightResultOption wrapped in HighlightResult.
func HighlightResultOptionAsHighlightResult(v *HighlightResultOption) *HighlightResult {
	return &HighlightResult{
		HighlightResultOption: v,
	}
}

// map[string]HighlightResultOptionAsHighlightResult is a convenience function that returns map[string]HighlightResultOption wrapped in HighlightResult.
func MapmapOfStringHighlightResultOptionAsHighlightResult(v map[string]HighlightResultOption) *HighlightResult {
	return &HighlightResult{
		MapmapOfStringHighlightResultOption: &v,
	}
}

// []HighlightResultOptionAsHighlightResult is a convenience function that returns []HighlightResultOption wrapped in HighlightResult.
func ArrayOfHighlightResultOptionAsHighlightResult(v []HighlightResultOption) *HighlightResult {
	return &HighlightResult{
		ArrayOfHighlightResultOption: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *HighlightResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into HighlightResultOption
	err = newStrictDecoder(data).Decode(&dst.HighlightResultOption)
	if err == nil && validateStruct(dst.HighlightResultOption) == nil {
		jsonHighlightResultOption, _ := json.Marshal(dst.HighlightResultOption)
		if string(jsonHighlightResultOption) == "{}" { // empty struct
			dst.HighlightResultOption = nil
		} else {
			return nil
		}
	} else {
		dst.HighlightResultOption = nil
	}

	// try to unmarshal data into ArrayOfHighlightResultOption
	err = newStrictDecoder(data).Decode(&dst.ArrayOfHighlightResultOption)
	if err == nil && validateStruct(dst.ArrayOfHighlightResultOption) == nil {
		jsonArrayOfHighlightResultOption, _ := json.Marshal(dst.ArrayOfHighlightResultOption)
		if string(jsonArrayOfHighlightResultOption) == "{}" { // empty struct
			dst.ArrayOfHighlightResultOption = nil
		} else {
			return nil
		}
	} else {
		dst.ArrayOfHighlightResultOption = nil
	}

	// try to unmarshal data into MapmapOfStringHighlightResultOption
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringHighlightResultOption)
	if err == nil && validateStruct(dst.MapmapOfStringHighlightResultOption) == nil {
		jsonMapmapOfStringHighlightResultOption, _ := json.Marshal(dst.MapmapOfStringHighlightResultOption)
		if string(jsonMapmapOfStringHighlightResultOption) == "{}" { // empty struct
			dst.MapmapOfStringHighlightResultOption = nil
		} else {
			return nil
		}
	} else {
		dst.MapmapOfStringHighlightResultOption = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(HighlightResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src HighlightResult) MarshalJSON() ([]byte, error) {
	if src.HighlightResultOption != nil {
		serialized, err := json.Marshal(&src.HighlightResultOption)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of HighlightResultOption of HighlightResult: %w", err)
		}

		return serialized, nil
	}

	if src.ArrayOfHighlightResultOption != nil {
		serialized, err := json.Marshal(&src.ArrayOfHighlightResultOption)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of ArrayOfHighlightResultOption of HighlightResult: %w", err)
		}

		return serialized, nil
	}

	if src.MapmapOfStringHighlightResultOption != nil {
		serialized, err := json.Marshal(&src.MapmapOfStringHighlightResultOption)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of MapmapOfStringHighlightResultOption of HighlightResult: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj HighlightResult) GetActualInstance() any {
	if obj.HighlightResultOption != nil {
		return *obj.HighlightResultOption
	}

	if obj.ArrayOfHighlightResultOption != nil {
		return *obj.ArrayOfHighlightResultOption
	}

	if obj.MapmapOfStringHighlightResultOption != nil {
		return *obj.MapmapOfStringHighlightResultOption
	}

	// all schemas are nil
	return nil
}

type NullableHighlightResult struct {
	value *HighlightResult
	isSet bool
}

func (v NullableHighlightResult) Get() *HighlightResult {
	return v.value
}

func (v *NullableHighlightResult) Set(val *HighlightResult) {
	v.value = val
	v.isSet = true
}

func (v NullableHighlightResult) IsSet() bool {
	return v.isSet
}

func (v *NullableHighlightResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighlightResult(val *HighlightResult) *NullableHighlightResult {
	return &NullableHighlightResult{value: val, isSet: true}
}

func (v NullableHighlightResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableHighlightResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
