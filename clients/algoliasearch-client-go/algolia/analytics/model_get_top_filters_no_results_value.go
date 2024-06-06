// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetTopFiltersNoResultsValue struct for GetTopFiltersNoResultsValue.
type GetTopFiltersNoResultsValue struct {
	// Attribute name.
	Attribute string   `json:"attribute"`
	Operator  Operator `json:"operator"`
	// Attribute value.
	Value string `json:"value"`
}

// NewGetTopFiltersNoResultsValue instantiates a new GetTopFiltersNoResultsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetTopFiltersNoResultsValue(attribute string, operator Operator, value string) *GetTopFiltersNoResultsValue {
	this := &GetTopFiltersNoResultsValue{}
	this.Attribute = attribute
	this.Operator = operator
	this.Value = value
	return this
}

// NewEmptyGetTopFiltersNoResultsValue return a pointer to an empty GetTopFiltersNoResultsValue object.
func NewEmptyGetTopFiltersNoResultsValue() *GetTopFiltersNoResultsValue {
	return &GetTopFiltersNoResultsValue{}
}

// GetAttribute returns the Attribute field value.
func (o *GetTopFiltersNoResultsValue) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *GetTopFiltersNoResultsValue) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value.
func (o *GetTopFiltersNoResultsValue) SetAttribute(v string) *GetTopFiltersNoResultsValue {
	o.Attribute = v
	return o
}

// GetOperator returns the Operator field value.
func (o *GetTopFiltersNoResultsValue) GetOperator() Operator {
	if o == nil {
		var ret Operator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *GetTopFiltersNoResultsValue) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *GetTopFiltersNoResultsValue) SetOperator(v Operator) *GetTopFiltersNoResultsValue {
	o.Operator = v
	return o
}

// GetValue returns the Value field value.
func (o *GetTopFiltersNoResultsValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTopFiltersNoResultsValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *GetTopFiltersNoResultsValue) SetValue(v string) *GetTopFiltersNoResultsValue {
	o.Value = v
	return o
}

func (o GetTopFiltersNoResultsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["attribute"] = o.Attribute
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if true {
		toSerialize["value"] = o.Value
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetTopFiltersNoResultsValue: %w", err)
	}

	return serialized, nil
}

func (o GetTopFiltersNoResultsValue) String() string {
	out := ""
	out += fmt.Sprintf("  attribute=%v\n", o.Attribute)
	out += fmt.Sprintf("  operator=%v\n", o.Operator)
	out += fmt.Sprintf("  value=%v\n", o.Value)
	return fmt.Sprintf("GetTopFiltersNoResultsValue {\n%s}", out)
}

type NullableGetTopFiltersNoResultsValue struct {
	value *GetTopFiltersNoResultsValue
	isSet bool
}

func (v NullableGetTopFiltersNoResultsValue) Get() *GetTopFiltersNoResultsValue {
	return v.value
}

func (v *NullableGetTopFiltersNoResultsValue) Set(val *GetTopFiltersNoResultsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTopFiltersNoResultsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTopFiltersNoResultsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTopFiltersNoResultsValue(val *GetTopFiltersNoResultsValue) *NullableGetTopFiltersNoResultsValue {
	return &NullableGetTopFiltersNoResultsValue{value: val, isSet: true}
}

func (v NullableGetTopFiltersNoResultsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableGetTopFiltersNoResultsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
