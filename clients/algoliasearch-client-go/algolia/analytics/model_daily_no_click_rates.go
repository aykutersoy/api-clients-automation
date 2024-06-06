// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// DailyNoClickRates struct for DailyNoClickRates.
type DailyNoClickRates struct {
	// No click rate, calculated as number of tracked searches without any click divided by the number of tracked searches.
	Rate float64 `json:"rate"`
	// Number of tracked searches. Tracked searches are search requests where the `clickAnalytics` parameter is true.
	Count int32 `json:"count"`
	// Number of times this search was returned as a result without any click.
	NoClickCount int32 `json:"noClickCount"`
	// Date in the format YYYY-MM-DD.
	Date string `json:"date"`
}

// NewDailyNoClickRates instantiates a new DailyNoClickRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDailyNoClickRates(rate float64, count int32, noClickCount int32, date string) *DailyNoClickRates {
	this := &DailyNoClickRates{}
	this.Rate = rate
	this.Count = count
	this.NoClickCount = noClickCount
	this.Date = date
	return this
}

// NewEmptyDailyNoClickRates return a pointer to an empty DailyNoClickRates object.
func NewEmptyDailyNoClickRates() *DailyNoClickRates {
	return &DailyNoClickRates{}
}

// GetRate returns the Rate field value.
func (o *DailyNoClickRates) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *DailyNoClickRates) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value.
func (o *DailyNoClickRates) SetRate(v float64) *DailyNoClickRates {
	o.Rate = v
	return o
}

// GetCount returns the Count field value.
func (o *DailyNoClickRates) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DailyNoClickRates) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *DailyNoClickRates) SetCount(v int32) *DailyNoClickRates {
	o.Count = v
	return o
}

// GetNoClickCount returns the NoClickCount field value.
func (o *DailyNoClickRates) GetNoClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NoClickCount
}

// GetNoClickCountOk returns a tuple with the NoClickCount field value
// and a boolean to check if the value has been set.
func (o *DailyNoClickRates) GetNoClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoClickCount, true
}

// SetNoClickCount sets field value.
func (o *DailyNoClickRates) SetNoClickCount(v int32) *DailyNoClickRates {
	o.NoClickCount = v
	return o
}

// GetDate returns the Date field value.
func (o *DailyNoClickRates) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DailyNoClickRates) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *DailyNoClickRates) SetDate(v string) *DailyNoClickRates {
	o.Date = v
	return o
}

func (o DailyNoClickRates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["rate"] = o.Rate
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["noClickCount"] = o.NoClickCount
	}
	if true {
		toSerialize["date"] = o.Date
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DailyNoClickRates: %w", err)
	}

	return serialized, nil
}

func (o DailyNoClickRates) String() string {
	out := ""
	out += fmt.Sprintf("  rate=%v\n", o.Rate)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	out += fmt.Sprintf("  noClickCount=%v\n", o.NoClickCount)
	out += fmt.Sprintf("  date=%v\n", o.Date)
	return fmt.Sprintf("DailyNoClickRates {\n%s}", out)
}

type NullableDailyNoClickRates struct {
	value *DailyNoClickRates
	isSet bool
}

func (v NullableDailyNoClickRates) Get() *DailyNoClickRates {
	return v.value
}

func (v *NullableDailyNoClickRates) Set(val *DailyNoClickRates) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyNoClickRates) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyNoClickRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyNoClickRates(val *DailyNoClickRates) *NullableDailyNoClickRates {
	return &NullableDailyNoClickRates{value: val, isSet: true}
}

func (v NullableDailyNoClickRates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDailyNoClickRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
