// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package monitoring

import (
	"encoding/json"
	"fmt"
)

// IncidentsResponse struct for IncidentsResponse.
type IncidentsResponse struct {
	Incidents *map[string][]IncidentsInner `json:"incidents,omitempty"`
}

type IncidentsResponseOption func(f *IncidentsResponse)

func WithIncidentsResponseIncidents(val map[string][]IncidentsInner) IncidentsResponseOption {
	return func(f *IncidentsResponse) {
		f.Incidents = &val
	}
}

// NewIncidentsResponse instantiates a new IncidentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentsResponse(opts ...IncidentsResponseOption) *IncidentsResponse {
	this := &IncidentsResponse{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyIncidentsResponse return a pointer to an empty IncidentsResponse object.
func NewEmptyIncidentsResponse() *IncidentsResponse {
	return &IncidentsResponse{}
}

// GetIncidents returns the Incidents field value if set, zero value otherwise.
func (o *IncidentsResponse) GetIncidents() map[string][]IncidentsInner {
	if o == nil || o.Incidents == nil {
		var ret map[string][]IncidentsInner
		return ret
	}
	return *o.Incidents
}

// GetIncidentsOk returns a tuple with the Incidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsResponse) GetIncidentsOk() (*map[string][]IncidentsInner, bool) {
	if o == nil || o.Incidents == nil {
		return nil, false
	}
	return o.Incidents, true
}

// HasIncidents returns a boolean if a field has been set.
func (o *IncidentsResponse) HasIncidents() bool {
	if o != nil && o.Incidents != nil {
		return true
	}

	return false
}

// SetIncidents gets a reference to the given map[string][]IncidentsInner and assigns it to the Incidents field.
func (o *IncidentsResponse) SetIncidents(v map[string][]IncidentsInner) *IncidentsResponse {
	o.Incidents = &v
	return o
}

func (o IncidentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Incidents != nil {
		toSerialize["incidents"] = o.Incidents
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal IncidentsResponse: %w", err)
	}

	return serialized, nil
}

func (o IncidentsResponse) String() string {
	out := ""
	out += fmt.Sprintf("  incidents=%v\n", o.Incidents)
	return fmt.Sprintf("IncidentsResponse {\n%s}", out)
}

type NullableIncidentsResponse struct {
	value *IncidentsResponse
	isSet bool
}

func (v NullableIncidentsResponse) Get() *IncidentsResponse {
	return v.value
}

func (v *NullableIncidentsResponse) Set(val *IncidentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentsResponse(val *IncidentsResponse) *NullableIncidentsResponse {
	return &NullableIncidentsResponse{value: val, isSet: true}
}

func (v NullableIncidentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableIncidentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
