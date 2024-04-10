// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package personalization

import (
	"encoding/json"
	"fmt"
)

// FacetScoring struct for FacetScoring.
type FacetScoring struct {
	// Event score.
	Score int32 `json:"score"`
	// Facet attribute name.
	FacetName string `json:"facetName"`
}

// NewFacetScoring instantiates a new FacetScoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetScoring(score int32, facetName string) *FacetScoring {
	this := &FacetScoring{}
	this.Score = score
	this.FacetName = facetName
	return this
}

// NewEmptyFacetScoring return a pointer to an empty FacetScoring object.
func NewEmptyFacetScoring() *FacetScoring {
	return &FacetScoring{}
}

// GetScore returns the Score field value.
func (o *FacetScoring) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *FacetScoring) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value.
func (o *FacetScoring) SetScore(v int32) *FacetScoring {
	o.Score = v
	return o
}

// GetFacetName returns the FacetName field value.
func (o *FacetScoring) GetFacetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value
// and a boolean to check if the value has been set.
func (o *FacetScoring) GetFacetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetName, true
}

// SetFacetName sets field value.
func (o *FacetScoring) SetFacetName(v string) *FacetScoring {
	o.FacetName = v
	return o
}

func (o FacetScoring) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["facetName"] = o.FacetName
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal FacetScoring: %w", err)
	}

	return serialized, nil
}

func (o FacetScoring) String() string {
	out := ""
	out += fmt.Sprintf("  score=%v\n", o.Score)
	out += fmt.Sprintf("  facetName=%v\n", o.FacetName)
	return fmt.Sprintf("FacetScoring {\n%s}", out)
}

type NullableFacetScoring struct {
	value *FacetScoring
	isSet bool
}

func (v NullableFacetScoring) Get() *FacetScoring {
	return v.value
}

func (v *NullableFacetScoring) Set(val *FacetScoring) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetScoring) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetScoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetScoring(val *FacetScoring) *NullableFacetScoring {
	return &NullableFacetScoring{value: val, isSet: true}
}

func (v NullableFacetScoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableFacetScoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
