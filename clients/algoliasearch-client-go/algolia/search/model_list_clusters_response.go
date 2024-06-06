// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// ListClustersResponse Clusters.
type ListClustersResponse struct {
	// Key-value pairs with cluster names as keys and lists of users with the highest number of records per cluster as values.
	TopUsers []string `json:"topUsers"`
}

// NewListClustersResponse instantiates a new ListClustersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListClustersResponse(topUsers []string) *ListClustersResponse {
	this := &ListClustersResponse{}
	this.TopUsers = topUsers
	return this
}

// NewEmptyListClustersResponse return a pointer to an empty ListClustersResponse object.
func NewEmptyListClustersResponse() *ListClustersResponse {
	return &ListClustersResponse{}
}

// GetTopUsers returns the TopUsers field value.
func (o *ListClustersResponse) GetTopUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TopUsers
}

// GetTopUsersOk returns a tuple with the TopUsers field value
// and a boolean to check if the value has been set.
func (o *ListClustersResponse) GetTopUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopUsers, true
}

// SetTopUsers sets field value.
func (o *ListClustersResponse) SetTopUsers(v []string) *ListClustersResponse {
	o.TopUsers = v
	return o
}

func (o ListClustersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["topUsers"] = o.TopUsers
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ListClustersResponse: %w", err)
	}

	return serialized, nil
}

func (o ListClustersResponse) String() string {
	out := ""
	out += fmt.Sprintf("  topUsers=%v\n", o.TopUsers)
	return fmt.Sprintf("ListClustersResponse {\n%s}", out)
}

type NullableListClustersResponse struct {
	value *ListClustersResponse
	isSet bool
}

func (v NullableListClustersResponse) Get() *ListClustersResponse {
	return v.value
}

func (v *NullableListClustersResponse) Set(val *ListClustersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListClustersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListClustersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClustersResponse(val *ListClustersResponse) *NullableListClustersResponse {
	return &NullableListClustersResponse{value: val, isSet: true}
}

func (v NullableListClustersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableListClustersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
