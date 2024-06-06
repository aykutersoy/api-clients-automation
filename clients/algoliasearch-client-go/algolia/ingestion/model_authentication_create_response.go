// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthenticationCreateResponse API response for the successful creation of an authentication resource.
type AuthenticationCreateResponse struct {
	// Universally unique identifier (UUID) of an authentication resource.
	AuthenticationID string `json:"authenticationID"`
	// Descriptive name for the resource.
	Name string `json:"name"`
	// Date of creation in RFC3339 format.
	CreatedAt string `json:"createdAt"`
}

// NewAuthenticationCreateResponse instantiates a new AuthenticationCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthenticationCreateResponse(authenticationID string, name string, createdAt string) *AuthenticationCreateResponse {
	this := &AuthenticationCreateResponse{}
	this.AuthenticationID = authenticationID
	this.Name = name
	this.CreatedAt = createdAt
	return this
}

// NewEmptyAuthenticationCreateResponse return a pointer to an empty AuthenticationCreateResponse object.
func NewEmptyAuthenticationCreateResponse() *AuthenticationCreateResponse {
	return &AuthenticationCreateResponse{}
}

// GetAuthenticationID returns the AuthenticationID field value.
func (o *AuthenticationCreateResponse) GetAuthenticationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationID
}

// GetAuthenticationIDOk returns a tuple with the AuthenticationID field value
// and a boolean to check if the value has been set.
func (o *AuthenticationCreateResponse) GetAuthenticationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationID, true
}

// SetAuthenticationID sets field value.
func (o *AuthenticationCreateResponse) SetAuthenticationID(v string) *AuthenticationCreateResponse {
	o.AuthenticationID = v
	return o
}

// GetName returns the Name field value.
func (o *AuthenticationCreateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticationCreateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AuthenticationCreateResponse) SetName(v string) *AuthenticationCreateResponse {
	o.Name = v
	return o
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AuthenticationCreateResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AuthenticationCreateResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AuthenticationCreateResponse) SetCreatedAt(v string) *AuthenticationCreateResponse {
	o.CreatedAt = v
	return o
}

func (o AuthenticationCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["authenticationID"] = o.AuthenticationID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthenticationCreateResponse: %w", err)
	}

	return serialized, nil
}

func (o AuthenticationCreateResponse) String() string {
	out := ""
	out += fmt.Sprintf("  authenticationID=%v\n", o.AuthenticationID)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	return fmt.Sprintf("AuthenticationCreateResponse {\n%s}", out)
}

type NullableAuthenticationCreateResponse struct {
	value *AuthenticationCreateResponse
	isSet bool
}

func (v NullableAuthenticationCreateResponse) Get() *AuthenticationCreateResponse {
	return v.value
}

func (v *NullableAuthenticationCreateResponse) Set(val *AuthenticationCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationCreateResponse(val *AuthenticationCreateResponse) *NullableAuthenticationCreateResponse {
	return &NullableAuthenticationCreateResponse{value: val, isSet: true}
}

func (v NullableAuthenticationCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAuthenticationCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
