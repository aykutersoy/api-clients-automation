// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// TopHit struct for TopHit.
type TopHit struct {
	// Object ID of a record that's returned as a search result.
	Hit string `json:"hit"`
	// Number of occurrences.
	Count int32 `json:"count"`
}

// NewTopHit instantiates a new TopHit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopHit(hit string, count int32) *TopHit {
	this := &TopHit{}
	this.Hit = hit
	this.Count = count
	return this
}

// NewEmptyTopHit return a pointer to an empty TopHit object.
func NewEmptyTopHit() *TopHit {
	return &TopHit{}
}

// GetHit returns the Hit field value.
func (o *TopHit) GetHit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hit
}

// GetHitOk returns a tuple with the Hit field value
// and a boolean to check if the value has been set.
func (o *TopHit) GetHitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hit, true
}

// SetHit sets field value.
func (o *TopHit) SetHit(v string) *TopHit {
	o.Hit = v
	return o
}

// GetCount returns the Count field value.
func (o *TopHit) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TopHit) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *TopHit) SetCount(v int32) *TopHit {
	o.Count = v
	return o
}

func (o TopHit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["hit"] = o.Hit
	}
	if true {
		toSerialize["count"] = o.Count
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TopHit: %w", err)
	}

	return serialized, nil
}

func (o TopHit) String() string {
	out := ""
	out += fmt.Sprintf("  hit=%v\n", o.Hit)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	return fmt.Sprintf("TopHit {\n%s}", out)
}

type NullableTopHit struct {
	value *TopHit
	isSet bool
}

func (v NullableTopHit) Get() *TopHit {
	return v.value
}

func (v *NullableTopHit) Set(val *TopHit) {
	v.value = val
	v.isSet = true
}

func (v NullableTopHit) IsSet() bool {
	return v.isSet
}

func (v *NullableTopHit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopHit(val *TopHit) *NullableTopHit {
	return &NullableTopHit{value: val, isSet: true}
}

func (v NullableTopHit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTopHit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
