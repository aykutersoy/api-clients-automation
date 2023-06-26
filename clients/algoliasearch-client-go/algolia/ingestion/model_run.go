// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// Run struct for Run
type Run struct {
	// The run UUID.
	RunID string `json:"runID" validate:"required"`
	AppID string `json:"appID" validate:"required"`
	// The task UUID.
	TaskID   string       `json:"taskID" validate:"required"`
	Status   RunStatus    `json:"status" validate:"required"`
	Progress *RunProgress `json:"progress,omitempty"`
	Outcome  *RunOutcome  `json:"outcome,omitempty"`
	// Explains the result of outcome.
	Reason     *string        `json:"reason,omitempty"`
	ReasonCode *RunReasonCode `json:"reasonCode,omitempty"`
	Type       RunType        `json:"type" validate:"required"`
	// Date of creation (RFC3339 format).
	CreatedAt string `json:"createdAt" validate:"required"`
	// Date of start (RFC3339 format).
	StartedAt *string `json:"startedAt,omitempty"`
	// Date of finish (RFC3339 format).
	FinishedAt *string `json:"finishedAt,omitempty"`
}

type RunOption func(f *Run)

func WithRunProgress(val RunProgress) RunOption {
	return func(f *Run) {
		f.Progress = &val
	}
}

func WithRunOutcome(val RunOutcome) RunOption {
	return func(f *Run) {
		f.Outcome = &val
	}
}

func WithRunReason(val string) RunOption {
	return func(f *Run) {
		f.Reason = &val
	}
}

func WithRunReasonCode(val RunReasonCode) RunOption {
	return func(f *Run) {
		f.ReasonCode = &val
	}
}

func WithRunStartedAt(val string) RunOption {
	return func(f *Run) {
		f.StartedAt = &val
	}
}

func WithRunFinishedAt(val string) RunOption {
	return func(f *Run) {
		f.FinishedAt = &val
	}
}

// NewRun instantiates a new Run object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRun(runID string, appID string, taskID string, status RunStatus, type_ RunType, createdAt string, opts ...RunOption) *Run {
	this := &Run{}
	this.RunID = runID
	this.AppID = appID
	this.TaskID = taskID
	this.Status = status
	this.Type = type_
	this.CreatedAt = createdAt
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewRunWithDefaults instantiates a new Run object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunWithDefaults() *Run {
	this := &Run{}
	return this
}

// GetRunID returns the RunID field value
func (o *Run) GetRunID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunID
}

// GetRunIDOk returns a tuple with the RunID field value
// and a boolean to check if the value has been set.
func (o *Run) GetRunIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunID, true
}

// SetRunID sets field value
func (o *Run) SetRunID(v string) {
	o.RunID = v
}

// GetAppID returns the AppID field value
func (o *Run) GetAppID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppID
}

// GetAppIDOk returns a tuple with the AppID field value
// and a boolean to check if the value has been set.
func (o *Run) GetAppIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppID, true
}

// SetAppID sets field value
func (o *Run) SetAppID(v string) {
	o.AppID = v
}

// GetTaskID returns the TaskID field value
func (o *Run) GetTaskID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskID
}

// GetTaskIDOk returns a tuple with the TaskID field value
// and a boolean to check if the value has been set.
func (o *Run) GetTaskIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskID, true
}

// SetTaskID sets field value
func (o *Run) SetTaskID(v string) {
	o.TaskID = v
}

// GetStatus returns the Status field value
func (o *Run) GetStatus() RunStatus {
	if o == nil {
		var ret RunStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Run) GetStatusOk() (*RunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Run) SetStatus(v RunStatus) {
	o.Status = v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *Run) GetProgress() RunProgress {
	if o == nil || o.Progress == nil {
		var ret RunProgress
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetProgressOk() (*RunProgress, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *Run) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given RunProgress and assigns it to the Progress field.
func (o *Run) SetProgress(v RunProgress) {
	o.Progress = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *Run) GetOutcome() RunOutcome {
	if o == nil || o.Outcome == nil {
		var ret RunOutcome
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetOutcomeOk() (*RunOutcome, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *Run) HasOutcome() bool {
	if o != nil && o.Outcome != nil {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given RunOutcome and assigns it to the Outcome field.
func (o *Run) SetOutcome(v RunOutcome) {
	o.Outcome = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Run) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Run) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Run) SetReason(v string) {
	o.Reason = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *Run) GetReasonCode() RunReasonCode {
	if o == nil || o.ReasonCode == nil {
		var ret RunReasonCode
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetReasonCodeOk() (*RunReasonCode, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *Run) HasReasonCode() bool {
	if o != nil && o.ReasonCode != nil {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given RunReasonCode and assigns it to the ReasonCode field.
func (o *Run) SetReasonCode(v RunReasonCode) {
	o.ReasonCode = &v
}

// GetType returns the Type field value
func (o *Run) GetType() RunType {
	if o == nil {
		var ret RunType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Run) GetTypeOk() (*RunType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Run) SetType(v RunType) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Run) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Run) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Run) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Run) GetStartedAt() string {
	if o == nil || o.StartedAt == nil {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetStartedAtOk() (*string, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Run) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *Run) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *Run) GetFinishedAt() string {
	if o == nil || o.FinishedAt == nil {
		var ret string
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetFinishedAtOk() (*string, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *Run) HasFinishedAt() bool {
	if o != nil && o.FinishedAt != nil {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given string and assigns it to the FinishedAt field.
func (o *Run) SetFinishedAt(v string) {
	o.FinishedAt = &v
}

func (o Run) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["runID"] = o.RunID
	}
	if true {
		toSerialize["appID"] = o.AppID
	}
	if true {
		toSerialize["taskID"] = o.TaskID
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.StartedAt != nil {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.FinishedAt != nil {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	return json.Marshal(toSerialize)
}

func (o Run) String() string {
	out := ""
	out += fmt.Sprintf("  runID=%v\n", o.RunID)
	out += fmt.Sprintf("  appID=%v\n", o.AppID)
	out += fmt.Sprintf("  taskID=%v\n", o.TaskID)
	out += fmt.Sprintf("  status=%v\n", o.Status)
	out += fmt.Sprintf("  progress=%v\n", o.Progress)
	out += fmt.Sprintf("  outcome=%v\n", o.Outcome)
	out += fmt.Sprintf("  reason=%v\n", o.Reason)
	out += fmt.Sprintf("  reasonCode=%v\n", o.ReasonCode)
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	out += fmt.Sprintf("  startedAt=%v\n", o.StartedAt)
	out += fmt.Sprintf("  finishedAt=%v\n", o.FinishedAt)
	return fmt.Sprintf("Run {\n%s}", out)
}

type NullableRun struct {
	value *Run
	isSet bool
}

func (v NullableRun) Get() *Run {
	return v.value
}

func (v *NullableRun) Set(val *Run) {
	v.value = val
	v.isSet = true
}

func (v NullableRun) IsSet() bool {
	return v.isSet
}

func (v *NullableRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRun(val *Run) *NullableRun {
	return &NullableRun{value: val, isSet: true}
}

func (v NullableRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
