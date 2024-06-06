// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package abtesting

import (
	"encoding/json"
	"fmt"
)

// Currency struct for Currency.
type Currency struct {
	// Currency code.
	Currency *string `json:"currency,omitempty"`
	// Revenue for this currency.
	Revenue *float64 `json:"revenue,omitempty"`
	// Mean for this currency.
	Mean *float64 `json:"mean,omitempty"`
	// Standard deviation for this currency.
	StandardDeviation *float64 `json:"standardDeviation,omitempty"`
}

type CurrencyOption func(f *Currency)

func WithCurrencyCurrency(val string) CurrencyOption {
	return func(f *Currency) {
		f.Currency = &val
	}
}

func WithCurrencyRevenue(val float64) CurrencyOption {
	return func(f *Currency) {
		f.Revenue = &val
	}
}

func WithCurrencyMean(val float64) CurrencyOption {
	return func(f *Currency) {
		f.Mean = &val
	}
}

func WithCurrencyStandardDeviation(val float64) CurrencyOption {
	return func(f *Currency) {
		f.StandardDeviation = &val
	}
}

// NewCurrency instantiates a new Currency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCurrency(opts ...CurrencyOption) *Currency {
	this := &Currency{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyCurrency return a pointer to an empty Currency object.
func NewEmptyCurrency() *Currency {
	return &Currency{}
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Currency) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Currency) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Currency) SetCurrency(v string) *Currency {
	o.Currency = &v
	return o
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *Currency) GetRevenue() float64 {
	if o == nil || o.Revenue == nil {
		var ret float64
		return ret
	}
	return *o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetRevenueOk() (*float64, bool) {
	if o == nil || o.Revenue == nil {
		return nil, false
	}
	return o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *Currency) HasRevenue() bool {
	if o != nil && o.Revenue != nil {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given float64 and assigns it to the Revenue field.
func (o *Currency) SetRevenue(v float64) *Currency {
	o.Revenue = &v
	return o
}

// GetMean returns the Mean field value if set, zero value otherwise.
func (o *Currency) GetMean() float64 {
	if o == nil || o.Mean == nil {
		var ret float64
		return ret
	}
	return *o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetMeanOk() (*float64, bool) {
	if o == nil || o.Mean == nil {
		return nil, false
	}
	return o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *Currency) HasMean() bool {
	if o != nil && o.Mean != nil {
		return true
	}

	return false
}

// SetMean gets a reference to the given float64 and assigns it to the Mean field.
func (o *Currency) SetMean(v float64) *Currency {
	o.Mean = &v
	return o
}

// GetStandardDeviation returns the StandardDeviation field value if set, zero value otherwise.
func (o *Currency) GetStandardDeviation() float64 {
	if o == nil || o.StandardDeviation == nil {
		var ret float64
		return ret
	}
	return *o.StandardDeviation
}

// GetStandardDeviationOk returns a tuple with the StandardDeviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetStandardDeviationOk() (*float64, bool) {
	if o == nil || o.StandardDeviation == nil {
		return nil, false
	}
	return o.StandardDeviation, true
}

// HasStandardDeviation returns a boolean if a field has been set.
func (o *Currency) HasStandardDeviation() bool {
	if o != nil && o.StandardDeviation != nil {
		return true
	}

	return false
}

// SetStandardDeviation gets a reference to the given float64 and assigns it to the StandardDeviation field.
func (o *Currency) SetStandardDeviation(v float64) *Currency {
	o.StandardDeviation = &v
	return o
}

func (o Currency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Revenue != nil {
		toSerialize["revenue"] = o.Revenue
	}
	if o.Mean != nil {
		toSerialize["mean"] = o.Mean
	}
	if o.StandardDeviation != nil {
		toSerialize["standardDeviation"] = o.StandardDeviation
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Currency: %w", err)
	}

	return serialized, nil
}

func (o Currency) String() string {
	out := ""
	out += fmt.Sprintf("  currency=%v\n", o.Currency)
	out += fmt.Sprintf("  revenue=%v\n", o.Revenue)
	out += fmt.Sprintf("  mean=%v\n", o.Mean)
	out += fmt.Sprintf("  standardDeviation=%v\n", o.StandardDeviation)
	return fmt.Sprintf("Currency {\n%s}", out)
}

type NullableCurrency struct {
	value *Currency
	isSet bool
}

func (v NullableCurrency) Get() *Currency {
	return v.value
}

func (v *NullableCurrency) Set(val *Currency) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrency(val *Currency) *NullableCurrency {
	return &NullableCurrency{value: val, isSet: true}
}

func (v NullableCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
