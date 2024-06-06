// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SaveObjectResponse struct for SaveObjectResponse.
type SaveObjectResponse struct {
	// Date and time when the object was created, in RFC 3339 format.
	CreatedAt string `json:"createdAt"`
	// Unique identifier of a task.  A successful API response means that a task was added to a queue. It might not run immediately. You can check the task's progress with the [`task` operation](#tag/Indices/operation/getTask) and this `taskID`.
	TaskID int64 `json:"taskID"`
	// Unique record identifier.
	ObjectID *string `json:"objectID,omitempty"`
}

type SaveObjectResponseOption func(f *SaveObjectResponse)

func WithSaveObjectResponseObjectID(val string) SaveObjectResponseOption {
	return func(f *SaveObjectResponse) {
		f.ObjectID = &val
	}
}

// NewSaveObjectResponse instantiates a new SaveObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSaveObjectResponse(createdAt string, taskID int64, opts ...SaveObjectResponseOption) *SaveObjectResponse {
	this := &SaveObjectResponse{}
	this.CreatedAt = createdAt
	this.TaskID = taskID
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySaveObjectResponse return a pointer to an empty SaveObjectResponse object.
func NewEmptySaveObjectResponse() *SaveObjectResponse {
	return &SaveObjectResponse{}
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SaveObjectResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SaveObjectResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SaveObjectResponse) SetCreatedAt(v string) *SaveObjectResponse {
	o.CreatedAt = v
	return o
}

// GetTaskID returns the TaskID field value.
func (o *SaveObjectResponse) GetTaskID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TaskID
}

// GetTaskIDOk returns a tuple with the TaskID field value
// and a boolean to check if the value has been set.
func (o *SaveObjectResponse) GetTaskIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskID, true
}

// SetTaskID sets field value.
func (o *SaveObjectResponse) SetTaskID(v int64) *SaveObjectResponse {
	o.TaskID = v
	return o
}

// GetObjectID returns the ObjectID field value if set, zero value otherwise.
func (o *SaveObjectResponse) GetObjectID() string {
	if o == nil || o.ObjectID == nil {
		var ret string
		return ret
	}
	return *o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveObjectResponse) GetObjectIDOk() (*string, bool) {
	if o == nil || o.ObjectID == nil {
		return nil, false
	}
	return o.ObjectID, true
}

// HasObjectID returns a boolean if a field has been set.
func (o *SaveObjectResponse) HasObjectID() bool {
	if o != nil && o.ObjectID != nil {
		return true
	}

	return false
}

// SetObjectID gets a reference to the given string and assigns it to the ObjectID field.
func (o *SaveObjectResponse) SetObjectID(v string) *SaveObjectResponse {
	o.ObjectID = &v
	return o
}

func (o SaveObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["taskID"] = o.TaskID
	}
	if o.ObjectID != nil {
		toSerialize["objectID"] = o.ObjectID
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SaveObjectResponse: %w", err)
	}

	return serialized, nil
}

func (o SaveObjectResponse) String() string {
	out := ""
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	out += fmt.Sprintf("  taskID=%v\n", o.TaskID)
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	return fmt.Sprintf("SaveObjectResponse {\n%s}", out)
}

type NullableSaveObjectResponse struct {
	value *SaveObjectResponse
	isSet bool
}

func (v NullableSaveObjectResponse) Get() *SaveObjectResponse {
	return v.value
}

func (v *NullableSaveObjectResponse) Set(val *SaveObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveObjectResponse(val *SaveObjectResponse) *NullableSaveObjectResponse {
	return &NullableSaveObjectResponse{value: val, isSet: true}
}

func (v NullableSaveObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSaveObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
