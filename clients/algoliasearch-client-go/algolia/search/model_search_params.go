// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchParams - struct for SearchParams.
type SearchParams struct {
	SearchParamsObject *SearchParamsObject
	SearchParamsString *SearchParamsString
}

// SearchParamsStringAsSearchParams is a convenience function that returns SearchParamsString wrapped in SearchParams.
func SearchParamsStringAsSearchParams(v *SearchParamsString) *SearchParams {
	return &SearchParams{
		SearchParamsString: v,
	}
}

// SearchParamsObjectAsSearchParams is a convenience function that returns SearchParamsObject wrapped in SearchParams.
func SearchParamsObjectAsSearchParams(v *SearchParamsObject) *SearchParams {
	return &SearchParams{
		SearchParamsObject: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *SearchParams) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]any
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup (SearchParamsString).")
	}

	// Hold the schema validity between checks
	validSchemaForModel := true

	// If the model wasn't discriminated yet, continue checking for other discriminating properties
	if validSchemaForModel {
		// Check if the model holds a property 'params'
		if _, ok := jsonDict["params"]; !ok {
			validSchemaForModel = false
		}
	}

	if validSchemaForModel {
		// try to unmarshal data into SearchParamsString
		err = newStrictDecoder(data).Decode(&dst.SearchParamsString)
		if err == nil && validateStruct(dst.SearchParamsString) == nil {
			jsonSearchParamsString, _ := json.Marshal(dst.SearchParamsString)
			if string(jsonSearchParamsString) == "{}" { // empty struct
				dst.SearchParamsString = nil
			} else {
				return nil
			}
		} else {
			dst.SearchParamsString = nil
		}
	}

	// Reset the schema validity for the next class check
	validSchemaForModel = true
	// try to unmarshal data into SearchParamsObject
	err = newStrictDecoder(data).Decode(&dst.SearchParamsObject)
	if err == nil && validateStruct(dst.SearchParamsObject) == nil {
		jsonSearchParamsObject, _ := json.Marshal(dst.SearchParamsObject)
		if string(jsonSearchParamsObject) == "{}" { // empty struct
			dst.SearchParamsObject = nil
		} else {
			return nil
		}
	} else {
		dst.SearchParamsObject = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(SearchParams)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src SearchParams) MarshalJSON() ([]byte, error) {
	if src.SearchParamsObject != nil {
		serialized, err := json.Marshal(&src.SearchParamsObject)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SearchParamsObject of SearchParams: %w", err)
		}

		return serialized, nil
	}

	if src.SearchParamsString != nil {
		serialized, err := json.Marshal(&src.SearchParamsString)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SearchParamsString of SearchParams: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj SearchParams) GetActualInstance() any {
	if obj.SearchParamsObject != nil {
		return *obj.SearchParamsObject
	}

	if obj.SearchParamsString != nil {
		return *obj.SearchParamsString
	}

	// all schemas are nil
	return nil
}

type NullableSearchParams struct {
	value *SearchParams
	isSet bool
}

func (v NullableSearchParams) Get() *SearchParams {
	return v.value
}

func (v *NullableSearchParams) Set(val *SearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchParams(val *SearchParams) *NullableSearchParams {
	return &NullableSearchParams{value: val, isSet: true}
}

func (v NullableSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
