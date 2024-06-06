// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

// DailyConversionRates struct for DailyConversionRates.
type DailyConversionRates struct {
	// Conversion rate, calculated as number of tracked searches with at least one conversion event divided by the number of tracked searches. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true.
	Rate utils.NullableFloat64 `json:"rate"`
	// Number of tracked searches. Tracked searches are search requests where the `clickAnalytics` parameter is true.
	TrackedSearchCount int32 `json:"trackedSearchCount"`
	// Number of conversions from this search.
	ConversionCount int32 `json:"conversionCount"`
	// Date in the format YYYY-MM-DD.
	Date string `json:"date"`
}

// NewDailyConversionRates instantiates a new DailyConversionRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDailyConversionRates(rate utils.NullableFloat64, trackedSearchCount int32, conversionCount int32, date string) *DailyConversionRates {
	this := &DailyConversionRates{}
	this.Rate = rate
	this.TrackedSearchCount = trackedSearchCount
	this.ConversionCount = conversionCount
	this.Date = date
	return this
}

// NewEmptyDailyConversionRates return a pointer to an empty DailyConversionRates object.
func NewEmptyDailyConversionRates() *DailyConversionRates {
	return &DailyConversionRates{}
}

// GetRate returns the Rate field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *DailyConversionRates) GetRate() float64 {
	if o == nil || o.Rate.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Rate.Get()
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DailyConversionRates) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rate.Get(), o.Rate.IsSet()
}

// SetRate sets field value.
func (o *DailyConversionRates) SetRate(v float64) *DailyConversionRates {
	o.Rate.Set(&v)
	return o
}

// GetTrackedSearchCount returns the TrackedSearchCount field value.
func (o *DailyConversionRates) GetTrackedSearchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrackedSearchCount
}

// GetTrackedSearchCountOk returns a tuple with the TrackedSearchCount field value
// and a boolean to check if the value has been set.
func (o *DailyConversionRates) GetTrackedSearchCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackedSearchCount, true
}

// SetTrackedSearchCount sets field value.
func (o *DailyConversionRates) SetTrackedSearchCount(v int32) *DailyConversionRates {
	o.TrackedSearchCount = v
	return o
}

// GetConversionCount returns the ConversionCount field value.
func (o *DailyConversionRates) GetConversionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConversionCount
}

// GetConversionCountOk returns a tuple with the ConversionCount field value
// and a boolean to check if the value has been set.
func (o *DailyConversionRates) GetConversionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionCount, true
}

// SetConversionCount sets field value.
func (o *DailyConversionRates) SetConversionCount(v int32) *DailyConversionRates {
	o.ConversionCount = v
	return o
}

// GetDate returns the Date field value.
func (o *DailyConversionRates) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DailyConversionRates) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *DailyConversionRates) SetDate(v string) *DailyConversionRates {
	o.Date = v
	return o
}

func (o DailyConversionRates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["rate"] = o.Rate.Get()
	}
	if true {
		toSerialize["trackedSearchCount"] = o.TrackedSearchCount
	}
	if true {
		toSerialize["conversionCount"] = o.ConversionCount
	}
	if true {
		toSerialize["date"] = o.Date
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DailyConversionRates: %w", err)
	}

	return serialized, nil
}

func (o DailyConversionRates) String() string {
	out := ""
	out += fmt.Sprintf("  rate=%v\n", o.Rate)
	out += fmt.Sprintf("  trackedSearchCount=%v\n", o.TrackedSearchCount)
	out += fmt.Sprintf("  conversionCount=%v\n", o.ConversionCount)
	out += fmt.Sprintf("  date=%v\n", o.Date)
	return fmt.Sprintf("DailyConversionRates {\n%s}", out)
}

type NullableDailyConversionRates struct {
	value *DailyConversionRates
	isSet bool
}

func (v NullableDailyConversionRates) Get() *DailyConversionRates {
	return v.value
}

func (v *NullableDailyConversionRates) Set(val *DailyConversionRates) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyConversionRates) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyConversionRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyConversionRates(val *DailyConversionRates) *NullableDailyConversionRates {
	return &NullableDailyConversionRates{value: val, isSet: true}
}

func (v NullableDailyConversionRates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDailyConversionRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
