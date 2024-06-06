// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// SortRemainingBy Order of facet values that aren't explicitly positioned with the `order` setting.  - `count`.   Order remaining facet values by decreasing count.   The count is the number of matching records containing this facet value.  - `alpha`.   Sort facet values alphabetically.  - `hidden`.   Don't show facet values that aren't explicitly positioned.
type SortRemainingBy string

// List of sortRemainingBy.
const (
	SORTREMAININGBY_COUNT  SortRemainingBy = "count"
	SORTREMAININGBY_ALPHA  SortRemainingBy = "alpha"
	SORTREMAININGBY_HIDDEN SortRemainingBy = "hidden"
)

// All allowed values of SortRemainingBy enum.
var AllowedSortRemainingByEnumValues = []SortRemainingBy{
	"count",
	"alpha",
	"hidden",
}

func (v *SortRemainingBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'SortRemainingBy': %w", string(src), err)
	}
	enumTypeValue := SortRemainingBy(value)
	for _, existing := range AllowedSortRemainingByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortRemainingBy", value)
}

// NewSortRemainingByFromValue returns a pointer to a valid SortRemainingBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSortRemainingByFromValue(v string) (*SortRemainingBy, error) {
	ev := SortRemainingBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortRemainingBy: valid values are %v", v, AllowedSortRemainingByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SortRemainingBy) IsValid() bool {
	for _, existing := range AllowedSortRemainingByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sortRemainingBy value.
func (v SortRemainingBy) Ptr() *SortRemainingBy {
	return &v
}

type NullableSortRemainingBy struct {
	value *SortRemainingBy
	isSet bool
}

func (v NullableSortRemainingBy) Get() *SortRemainingBy {
	return v.value
}

func (v *NullableSortRemainingBy) Set(val *SortRemainingBy) {
	v.value = val
	v.isSet = true
}

func (v NullableSortRemainingBy) IsSet() bool {
	return v.isSet
}

func (v *NullableSortRemainingBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortRemainingBy(val *SortRemainingBy) *NullableSortRemainingBy {
	return &NullableSortRemainingBy{value: val, isSet: true}
}

func (v NullableSortRemainingBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSortRemainingBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
