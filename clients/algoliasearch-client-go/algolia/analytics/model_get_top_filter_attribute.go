// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetTopFilterAttribute struct for GetTopFilterAttribute.
type GetTopFilterAttribute struct {
	// Attribute name.
	Attribute string `json:"attribute"`
	// Number of occurrences.
	Count int32 `json:"count"`
}

// NewGetTopFilterAttribute instantiates a new GetTopFilterAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetTopFilterAttribute(attribute string, count int32) *GetTopFilterAttribute {
	this := &GetTopFilterAttribute{}
	this.Attribute = attribute
	this.Count = count
	return this
}

// NewEmptyGetTopFilterAttribute return a pointer to an empty GetTopFilterAttribute object.
func NewEmptyGetTopFilterAttribute() *GetTopFilterAttribute {
	return &GetTopFilterAttribute{}
}

// GetAttribute returns the Attribute field value.
func (o *GetTopFilterAttribute) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *GetTopFilterAttribute) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value.
func (o *GetTopFilterAttribute) SetAttribute(v string) *GetTopFilterAttribute {
	o.Attribute = v
	return o
}

// GetCount returns the Count field value.
func (o *GetTopFilterAttribute) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetTopFilterAttribute) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *GetTopFilterAttribute) SetCount(v int32) *GetTopFilterAttribute {
	o.Count = v
	return o
}

func (o GetTopFilterAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["attribute"] = o.Attribute
	}
	if true {
		toSerialize["count"] = o.Count
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetTopFilterAttribute: %w", err)
	}

	return serialized, nil
}

func (o GetTopFilterAttribute) String() string {
	out := ""
	out += fmt.Sprintf("  attribute=%v\n", o.Attribute)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	return fmt.Sprintf("GetTopFilterAttribute {\n%s}", out)
}

type NullableGetTopFilterAttribute struct {
	value *GetTopFilterAttribute
	isSet bool
}

func (v NullableGetTopFilterAttribute) Get() *GetTopFilterAttribute {
	return v.value
}

func (v *NullableGetTopFilterAttribute) Set(val *GetTopFilterAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTopFilterAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTopFilterAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTopFilterAttribute(val *GetTopFilterAttribute) *NullableGetTopFilterAttribute {
	return &NullableGetTopFilterAttribute{value: val, isSet: true}
}

func (v NullableGetTopFilterAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableGetTopFilterAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
