// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// Personalization struct for Personalization.
type Personalization struct {
	// The score of the filters.
	FiltersScore *int32 `json:"filtersScore,omitempty"`
	// The score of the ranking.
	RankingScore *int32 `json:"rankingScore,omitempty"`
	// The score of the event.
	Score *int32 `json:"score,omitempty"`
}

type PersonalizationOption func(f *Personalization)

func WithPersonalizationFiltersScore(val int32) PersonalizationOption {
	return func(f *Personalization) {
		f.FiltersScore = &val
	}
}

func WithPersonalizationRankingScore(val int32) PersonalizationOption {
	return func(f *Personalization) {
		f.RankingScore = &val
	}
}

func WithPersonalizationScore(val int32) PersonalizationOption {
	return func(f *Personalization) {
		f.Score = &val
	}
}

// NewPersonalization instantiates a new Personalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersonalization(opts ...PersonalizationOption) *Personalization {
	this := &Personalization{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyPersonalization return a pointer to an empty Personalization object.
func NewEmptyPersonalization() *Personalization {
	return &Personalization{}
}

// GetFiltersScore returns the FiltersScore field value if set, zero value otherwise.
func (o *Personalization) GetFiltersScore() int32 {
	if o == nil || o.FiltersScore == nil {
		var ret int32
		return ret
	}
	return *o.FiltersScore
}

// GetFiltersScoreOk returns a tuple with the FiltersScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Personalization) GetFiltersScoreOk() (*int32, bool) {
	if o == nil || o.FiltersScore == nil {
		return nil, false
	}
	return o.FiltersScore, true
}

// HasFiltersScore returns a boolean if a field has been set.
func (o *Personalization) HasFiltersScore() bool {
	if o != nil && o.FiltersScore != nil {
		return true
	}

	return false
}

// SetFiltersScore gets a reference to the given int32 and assigns it to the FiltersScore field.
func (o *Personalization) SetFiltersScore(v int32) *Personalization {
	o.FiltersScore = &v
	return o
}

// GetRankingScore returns the RankingScore field value if set, zero value otherwise.
func (o *Personalization) GetRankingScore() int32 {
	if o == nil || o.RankingScore == nil {
		var ret int32
		return ret
	}
	return *o.RankingScore
}

// GetRankingScoreOk returns a tuple with the RankingScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Personalization) GetRankingScoreOk() (*int32, bool) {
	if o == nil || o.RankingScore == nil {
		return nil, false
	}
	return o.RankingScore, true
}

// HasRankingScore returns a boolean if a field has been set.
func (o *Personalization) HasRankingScore() bool {
	if o != nil && o.RankingScore != nil {
		return true
	}

	return false
}

// SetRankingScore gets a reference to the given int32 and assigns it to the RankingScore field.
func (o *Personalization) SetRankingScore(v int32) *Personalization {
	o.RankingScore = &v
	return o
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *Personalization) GetScore() int32 {
	if o == nil || o.Score == nil {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Personalization) GetScoreOk() (*int32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *Personalization) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *Personalization) SetScore(v int32) *Personalization {
	o.Score = &v
	return o
}

func (o Personalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.FiltersScore != nil {
		toSerialize["filtersScore"] = o.FiltersScore
	}
	if o.RankingScore != nil {
		toSerialize["rankingScore"] = o.RankingScore
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Personalization: %w", err)
	}

	return serialized, nil
}

func (o Personalization) String() string {
	out := ""
	out += fmt.Sprintf("  filtersScore=%v\n", o.FiltersScore)
	out += fmt.Sprintf("  rankingScore=%v\n", o.RankingScore)
	out += fmt.Sprintf("  score=%v\n", o.Score)
	return fmt.Sprintf("Personalization {\n%s}", out)
}

type NullablePersonalization struct {
	value *Personalization
	isSet bool
}

func (v NullablePersonalization) Get() *Personalization {
	return v.value
}

func (v *NullablePersonalization) Set(val *Personalization) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalization) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalization(val *Personalization) *NullablePersonalization {
	return &NullablePersonalization{value: val, isSet: true}
}

func (v NullablePersonalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullablePersonalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
