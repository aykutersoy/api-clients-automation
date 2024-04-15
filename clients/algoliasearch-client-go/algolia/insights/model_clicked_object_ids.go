// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// ClickedObjectIDs Use this event to track when users click items unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track click events related to Algolia requests, use the \"Clicked object IDs after search\" event.
type ClickedObjectIDs struct {
	// Event name, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
	EventName string     `json:"eventName"`
	EventType ClickEvent `json:"eventType"`
	// Index name (case-sensitive) to which the event's items belong.
	Index string `json:"index"`
	// Object IDs of the records that are part of the event.
	ObjectIDs []string `json:"objectIDs"`
	// Anonymous or pseudonymous user identifier.  Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
	UserToken string `json:"userToken"`
	// Identifier for authenticated users.  When the user signs in, you can get an identifier from your system and send it as `authenticatedUserToken`. This lets you keep using the `userToken` from before the user signed in, while providing a reliable way to identify users across sessions. Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
	AuthenticatedUserToken *string `json:"authenticatedUserToken,omitempty"`
	// Timestamp of the event, measured in milliseconds since the Unix epoch. By default, the Insights API uses the time it receives an event as its timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type ClickedObjectIDsOption func(f *ClickedObjectIDs)

func WithClickedObjectIDsAuthenticatedUserToken(val string) ClickedObjectIDsOption {
	return func(f *ClickedObjectIDs) {
		f.AuthenticatedUserToken = &val
	}
}

func WithClickedObjectIDsTimestamp(val int64) ClickedObjectIDsOption {
	return func(f *ClickedObjectIDs) {
		f.Timestamp = &val
	}
}

// NewClickedObjectIDs instantiates a new ClickedObjectIDs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickedObjectIDs(eventName string, eventType ClickEvent, index string, objectIDs []string, userToken string, opts ...ClickedObjectIDsOption) *ClickedObjectIDs {
	this := &ClickedObjectIDs{}
	this.EventName = eventName
	this.EventType = eventType
	this.Index = index
	this.ObjectIDs = objectIDs
	this.UserToken = userToken
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyClickedObjectIDs return a pointer to an empty ClickedObjectIDs object.
func NewEmptyClickedObjectIDs() *ClickedObjectIDs {
	return &ClickedObjectIDs{}
}

// GetEventName returns the EventName field value.
func (o *ClickedObjectIDs) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value.
func (o *ClickedObjectIDs) SetEventName(v string) *ClickedObjectIDs {
	o.EventName = v
	return o
}

// GetEventType returns the EventType field value.
func (o *ClickedObjectIDs) GetEventType() ClickEvent {
	if o == nil {
		var ret ClickEvent
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetEventTypeOk() (*ClickEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *ClickedObjectIDs) SetEventType(v ClickEvent) *ClickedObjectIDs {
	o.EventType = v
	return o
}

// GetIndex returns the Index field value.
func (o *ClickedObjectIDs) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *ClickedObjectIDs) SetIndex(v string) *ClickedObjectIDs {
	o.Index = v
	return o
}

// GetObjectIDs returns the ObjectIDs field value.
func (o *ClickedObjectIDs) GetObjectIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectIDs
}

// GetObjectIDsOk returns a tuple with the ObjectIDs field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetObjectIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectIDs, true
}

// SetObjectIDs sets field value.
func (o *ClickedObjectIDs) SetObjectIDs(v []string) *ClickedObjectIDs {
	o.ObjectIDs = v
	return o
}

// GetUserToken returns the UserToken field value.
func (o *ClickedObjectIDs) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value.
func (o *ClickedObjectIDs) SetUserToken(v string) *ClickedObjectIDs {
	o.UserToken = v
	return o
}

// GetAuthenticatedUserToken returns the AuthenticatedUserToken field value if set, zero value otherwise.
func (o *ClickedObjectIDs) GetAuthenticatedUserToken() string {
	if o == nil || o.AuthenticatedUserToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatedUserToken
}

// GetAuthenticatedUserTokenOk returns a tuple with the AuthenticatedUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetAuthenticatedUserTokenOk() (*string, bool) {
	if o == nil || o.AuthenticatedUserToken == nil {
		return nil, false
	}
	return o.AuthenticatedUserToken, true
}

// HasAuthenticatedUserToken returns a boolean if a field has been set.
func (o *ClickedObjectIDs) HasAuthenticatedUserToken() bool {
	if o != nil && o.AuthenticatedUserToken != nil {
		return true
	}

	return false
}

// SetAuthenticatedUserToken gets a reference to the given string and assigns it to the AuthenticatedUserToken field.
func (o *ClickedObjectIDs) SetAuthenticatedUserToken(v string) *ClickedObjectIDs {
	o.AuthenticatedUserToken = &v
	return o
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClickedObjectIDs) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClickedObjectIDs) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ClickedObjectIDs) SetTimestamp(v int64) *ClickedObjectIDs {
	o.Timestamp = &v
	return o
}

func (o ClickedObjectIDs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["objectIDs"] = o.ObjectIDs
	}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if o.AuthenticatedUserToken != nil {
		toSerialize["authenticatedUserToken"] = o.AuthenticatedUserToken
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ClickedObjectIDs: %w", err)
	}

	return serialized, nil
}

func (o ClickedObjectIDs) String() string {
	out := ""
	out += fmt.Sprintf("  eventName=%v\n", o.EventName)
	out += fmt.Sprintf("  eventType=%v\n", o.EventType)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  objectIDs=%v\n", o.ObjectIDs)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  authenticatedUserToken=%v\n", o.AuthenticatedUserToken)
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	return fmt.Sprintf("ClickedObjectIDs {\n%s}", out)
}

type NullableClickedObjectIDs struct {
	value *ClickedObjectIDs
	isSet bool
}

func (v NullableClickedObjectIDs) Get() *ClickedObjectIDs {
	return v.value
}

func (v *NullableClickedObjectIDs) Set(val *ClickedObjectIDs) {
	v.value = val
	v.isSet = true
}

func (v NullableClickedObjectIDs) IsSet() bool {
	return v.isSet
}

func (v *NullableClickedObjectIDs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickedObjectIDs(val *ClickedObjectIDs) *NullableClickedObjectIDs {
	return &NullableClickedObjectIDs{value: val, isSet: true}
}

func (v NullableClickedObjectIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableClickedObjectIDs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
