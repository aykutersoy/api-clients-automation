// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// DailyUsers struct for DailyUsers.
type DailyUsers struct {
	// Date in the format YYYY-MM-DD.
	Date string `json:"date"`
	// Number of unique users.
	Count int32 `json:"count"`
}

// NewDailyUsers instantiates a new DailyUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDailyUsers(date string, count int32) *DailyUsers {
	this := &DailyUsers{}
	this.Date = date
	this.Count = count
	return this
}

// NewEmptyDailyUsers return a pointer to an empty DailyUsers object.
func NewEmptyDailyUsers() *DailyUsers {
	return &DailyUsers{}
}

// GetDate returns the Date field value.
func (o *DailyUsers) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DailyUsers) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *DailyUsers) SetDate(v string) *DailyUsers {
	o.Date = v
	return o
}

// GetCount returns the Count field value.
func (o *DailyUsers) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DailyUsers) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *DailyUsers) SetCount(v int32) *DailyUsers {
	o.Count = v
	return o
}

func (o DailyUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["count"] = o.Count
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DailyUsers: %w", err)
	}

	return serialized, nil
}

func (o DailyUsers) String() string {
	out := ""
	out += fmt.Sprintf("  date=%v\n", o.Date)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	return fmt.Sprintf("DailyUsers {\n%s}", out)
}

type NullableDailyUsers struct {
	value *DailyUsers
	isSet bool
}

func (v NullableDailyUsers) Get() *DailyUsers {
	return v.value
}

func (v *NullableDailyUsers) Set(val *DailyUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyUsers(val *DailyUsers) *NullableDailyUsers {
	return &NullableDailyUsers{value: val, isSet: true}
}

func (v NullableDailyUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDailyUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
