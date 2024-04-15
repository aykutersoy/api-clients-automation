// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// TrendingFacetsQuery struct for TrendingFacetsQuery.
type TrendingFacetsQuery struct {
	// Index name (case-sensitive).
	IndexName string `json:"indexName"`
	// Minimum score a recommendation must have to be included in the response.
	Threshold float64 `json:"threshold"`
	// Maximum number of recommendations to retrieve. By default, all recommendations are returned and no fallback request is made. Depending on the available recommendations and the other request parameters, the actual number of recommendations may be lower than this value.
	MaxRecommendations *int32        `json:"maxRecommendations,omitempty"`
	QueryParameters    *SearchParams `json:"queryParameters,omitempty"`
	// Facet attribute for which to retrieve trending facet values.
	FacetName          map[string]interface{} `json:"facetName"`
	Model              TrendingFacetsModel    `json:"model"`
	FallbackParameters *FallbackParams        `json:"fallbackParameters,omitempty"`
}

type TrendingFacetsQueryOption func(f *TrendingFacetsQuery)

func WithTrendingFacetsQueryMaxRecommendations(val int32) TrendingFacetsQueryOption {
	return func(f *TrendingFacetsQuery) {
		f.MaxRecommendations = &val
	}
}

func WithTrendingFacetsQueryQueryParameters(val SearchParams) TrendingFacetsQueryOption {
	return func(f *TrendingFacetsQuery) {
		f.QueryParameters = &val
	}
}

func WithTrendingFacetsQueryFallbackParameters(val FallbackParams) TrendingFacetsQueryOption {
	return func(f *TrendingFacetsQuery) {
		f.FallbackParameters = &val
	}
}

// NewTrendingFacetsQuery instantiates a new TrendingFacetsQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTrendingFacetsQuery(indexName string, threshold float64, facetName map[string]interface{}, model TrendingFacetsModel, opts ...TrendingFacetsQueryOption) *TrendingFacetsQuery {
	this := &TrendingFacetsQuery{}
	this.IndexName = indexName
	this.Threshold = threshold
	this.FacetName = facetName
	this.Model = model
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyTrendingFacetsQuery return a pointer to an empty TrendingFacetsQuery object.
func NewEmptyTrendingFacetsQuery() *TrendingFacetsQuery {
	return &TrendingFacetsQuery{}
}

// GetIndexName returns the IndexName field value.
func (o *TrendingFacetsQuery) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *TrendingFacetsQuery) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *TrendingFacetsQuery) SetIndexName(v string) *TrendingFacetsQuery {
	o.IndexName = v
	return o
}

// GetThreshold returns the Threshold field value.
func (o *TrendingFacetsQuery) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *TrendingFacetsQuery) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value.
func (o *TrendingFacetsQuery) SetThreshold(v float64) *TrendingFacetsQuery {
	o.Threshold = v
	return o
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *TrendingFacetsQuery) GetMaxRecommendations() int32 {
	if o == nil || o.MaxRecommendations == nil {
		var ret int32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingFacetsQuery) GetMaxRecommendationsOk() (*int32, bool) {
	if o == nil || o.MaxRecommendations == nil {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *TrendingFacetsQuery) HasMaxRecommendations() bool {
	if o != nil && o.MaxRecommendations != nil {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given int32 and assigns it to the MaxRecommendations field.
func (o *TrendingFacetsQuery) SetMaxRecommendations(v int32) *TrendingFacetsQuery {
	o.MaxRecommendations = &v
	return o
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *TrendingFacetsQuery) GetQueryParameters() SearchParams {
	if o == nil || o.QueryParameters == nil {
		var ret SearchParams
		return ret
	}
	return *o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingFacetsQuery) GetQueryParametersOk() (*SearchParams, bool) {
	if o == nil || o.QueryParameters == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *TrendingFacetsQuery) HasQueryParameters() bool {
	if o != nil && o.QueryParameters != nil {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given SearchParams and assigns it to the QueryParameters field.
func (o *TrendingFacetsQuery) SetQueryParameters(v *SearchParams) *TrendingFacetsQuery {
	o.QueryParameters = v
	return o
}

// GetFacetName returns the FacetName field value.
func (o *TrendingFacetsQuery) GetFacetName() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value
// and a boolean to check if the value has been set.
func (o *TrendingFacetsQuery) GetFacetNameOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacetName, true
}

// SetFacetName sets field value.
func (o *TrendingFacetsQuery) SetFacetName(v map[string]interface{}) *TrendingFacetsQuery {
	o.FacetName = v
	return o
}

// GetModel returns the Model field value.
func (o *TrendingFacetsQuery) GetModel() TrendingFacetsModel {
	if o == nil {
		var ret TrendingFacetsModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *TrendingFacetsQuery) GetModelOk() (*TrendingFacetsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value.
func (o *TrendingFacetsQuery) SetModel(v TrendingFacetsModel) *TrendingFacetsQuery {
	o.Model = v
	return o
}

// GetFallbackParameters returns the FallbackParameters field value if set, zero value otherwise.
func (o *TrendingFacetsQuery) GetFallbackParameters() FallbackParams {
	if o == nil || o.FallbackParameters == nil {
		var ret FallbackParams
		return ret
	}
	return *o.FallbackParameters
}

// GetFallbackParametersOk returns a tuple with the FallbackParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingFacetsQuery) GetFallbackParametersOk() (*FallbackParams, bool) {
	if o == nil || o.FallbackParameters == nil {
		return nil, false
	}
	return o.FallbackParameters, true
}

// HasFallbackParameters returns a boolean if a field has been set.
func (o *TrendingFacetsQuery) HasFallbackParameters() bool {
	if o != nil && o.FallbackParameters != nil {
		return true
	}

	return false
}

// SetFallbackParameters gets a reference to the given FallbackParams and assigns it to the FallbackParameters field.
func (o *TrendingFacetsQuery) SetFallbackParameters(v *FallbackParams) *TrendingFacetsQuery {
	o.FallbackParameters = v
	return o
}

func (o TrendingFacetsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if o.MaxRecommendations != nil {
		toSerialize["maxRecommendations"] = o.MaxRecommendations
	}
	if o.QueryParameters != nil {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if true {
		toSerialize["facetName"] = o.FacetName
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.FallbackParameters != nil {
		toSerialize["fallbackParameters"] = o.FallbackParameters
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TrendingFacetsQuery: %w", err)
	}

	return serialized, nil
}

func (o TrendingFacetsQuery) String() string {
	out := ""
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  threshold=%v\n", o.Threshold)
	out += fmt.Sprintf("  maxRecommendations=%v\n", o.MaxRecommendations)
	out += fmt.Sprintf("  queryParameters=%v\n", o.QueryParameters)
	out += fmt.Sprintf("  facetName=%v\n", o.FacetName)
	out += fmt.Sprintf("  model=%v\n", o.Model)
	out += fmt.Sprintf("  fallbackParameters=%v\n", o.FallbackParameters)
	return fmt.Sprintf("TrendingFacetsQuery {\n%s}", out)
}

type NullableTrendingFacetsQuery struct {
	value *TrendingFacetsQuery
	isSet bool
}

func (v NullableTrendingFacetsQuery) Get() *TrendingFacetsQuery {
	return v.value
}

func (v *NullableTrendingFacetsQuery) Set(val *TrendingFacetsQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTrendingFacetsQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTrendingFacetsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrendingFacetsQuery(val *TrendingFacetsQuery) *NullableTrendingFacetsQuery {
	return &NullableTrendingFacetsQuery{value: val, isSet: true}
}

func (v NullableTrendingFacetsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTrendingFacetsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
