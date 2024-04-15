// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// TrendingItemsQuery struct for TrendingItemsQuery.
type TrendingItemsQuery struct {
	// Index name (case-sensitive).
	IndexName string `json:"indexName"`
	// Minimum score a recommendation must have to be included in the response.
	Threshold float64 `json:"threshold"`
	// Maximum number of recommendations to retrieve. By default, all recommendations are returned and no fallback request is made. Depending on the available recommendations and the other request parameters, the actual number of recommendations may be lower than this value.
	MaxRecommendations *int32        `json:"maxRecommendations,omitempty"`
	QueryParameters    *SearchParams `json:"queryParameters,omitempty"`
	// Facet attribute. To be used in combination with `facetValue`. If specified, only recommendations matching the facet filter will be returned.
	FacetName string `json:"facetName"`
	// Facet value. To be used in combination with `facetName`. If specified, only recommendations matching the facet filter will be returned.
	FacetValue         string              `json:"facetValue"`
	Model              TrendingItemsModel  `json:"model"`
	FallbackParameters *SearchParamsObject `json:"fallbackParameters,omitempty"`
}

type TrendingItemsQueryOption func(f *TrendingItemsQuery)

func WithTrendingItemsQueryMaxRecommendations(val int32) TrendingItemsQueryOption {
	return func(f *TrendingItemsQuery) {
		f.MaxRecommendations = &val
	}
}

func WithTrendingItemsQueryQueryParameters(val SearchParams) TrendingItemsQueryOption {
	return func(f *TrendingItemsQuery) {
		f.QueryParameters = &val
	}
}

func WithTrendingItemsQueryFallbackParameters(val SearchParamsObject) TrendingItemsQueryOption {
	return func(f *TrendingItemsQuery) {
		f.FallbackParameters = &val
	}
}

// NewTrendingItemsQuery instantiates a new TrendingItemsQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTrendingItemsQuery(indexName string, threshold float64, facetName string, facetValue string, model TrendingItemsModel, opts ...TrendingItemsQueryOption) *TrendingItemsQuery {
	this := &TrendingItemsQuery{}
	this.IndexName = indexName
	this.Threshold = threshold
	this.FacetName = facetName
	this.FacetValue = facetValue
	this.Model = model
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyTrendingItemsQuery return a pointer to an empty TrendingItemsQuery object.
func NewEmptyTrendingItemsQuery() *TrendingItemsQuery {
	return &TrendingItemsQuery{}
}

// GetIndexName returns the IndexName field value.
func (o *TrendingItemsQuery) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *TrendingItemsQuery) SetIndexName(v string) *TrendingItemsQuery {
	o.IndexName = v
	return o
}

// GetThreshold returns the Threshold field value.
func (o *TrendingItemsQuery) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value.
func (o *TrendingItemsQuery) SetThreshold(v float64) *TrendingItemsQuery {
	o.Threshold = v
	return o
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *TrendingItemsQuery) GetMaxRecommendations() int32 {
	if o == nil || o.MaxRecommendations == nil {
		var ret int32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetMaxRecommendationsOk() (*int32, bool) {
	if o == nil || o.MaxRecommendations == nil {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *TrendingItemsQuery) HasMaxRecommendations() bool {
	if o != nil && o.MaxRecommendations != nil {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given int32 and assigns it to the MaxRecommendations field.
func (o *TrendingItemsQuery) SetMaxRecommendations(v int32) *TrendingItemsQuery {
	o.MaxRecommendations = &v
	return o
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *TrendingItemsQuery) GetQueryParameters() SearchParams {
	if o == nil || o.QueryParameters == nil {
		var ret SearchParams
		return ret
	}
	return *o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetQueryParametersOk() (*SearchParams, bool) {
	if o == nil || o.QueryParameters == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *TrendingItemsQuery) HasQueryParameters() bool {
	if o != nil && o.QueryParameters != nil {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given SearchParams and assigns it to the QueryParameters field.
func (o *TrendingItemsQuery) SetQueryParameters(v *SearchParams) *TrendingItemsQuery {
	o.QueryParameters = v
	return o
}

// GetFacetName returns the FacetName field value.
func (o *TrendingItemsQuery) GetFacetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetFacetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetName, true
}

// SetFacetName sets field value.
func (o *TrendingItemsQuery) SetFacetName(v string) *TrendingItemsQuery {
	o.FacetName = v
	return o
}

// GetFacetValue returns the FacetValue field value.
func (o *TrendingItemsQuery) GetFacetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacetValue
}

// GetFacetValueOk returns a tuple with the FacetValue field value
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetFacetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetValue, true
}

// SetFacetValue sets field value.
func (o *TrendingItemsQuery) SetFacetValue(v string) *TrendingItemsQuery {
	o.FacetValue = v
	return o
}

// GetModel returns the Model field value.
func (o *TrendingItemsQuery) GetModel() TrendingItemsModel {
	if o == nil {
		var ret TrendingItemsModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetModelOk() (*TrendingItemsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value.
func (o *TrendingItemsQuery) SetModel(v TrendingItemsModel) *TrendingItemsQuery {
	o.Model = v
	return o
}

// GetFallbackParameters returns the FallbackParameters field value if set, zero value otherwise.
func (o *TrendingItemsQuery) GetFallbackParameters() SearchParamsObject {
	if o == nil || o.FallbackParameters == nil {
		var ret SearchParamsObject
		return ret
	}
	return *o.FallbackParameters
}

// GetFallbackParametersOk returns a tuple with the FallbackParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingItemsQuery) GetFallbackParametersOk() (*SearchParamsObject, bool) {
	if o == nil || o.FallbackParameters == nil {
		return nil, false
	}
	return o.FallbackParameters, true
}

// HasFallbackParameters returns a boolean if a field has been set.
func (o *TrendingItemsQuery) HasFallbackParameters() bool {
	if o != nil && o.FallbackParameters != nil {
		return true
	}

	return false
}

// SetFallbackParameters gets a reference to the given SearchParamsObject and assigns it to the FallbackParameters field.
func (o *TrendingItemsQuery) SetFallbackParameters(v *SearchParamsObject) *TrendingItemsQuery {
	o.FallbackParameters = v
	return o
}

func (o TrendingItemsQuery) MarshalJSON() ([]byte, error) {
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
		toSerialize["facetValue"] = o.FacetValue
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.FallbackParameters != nil {
		toSerialize["fallbackParameters"] = o.FallbackParameters
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TrendingItemsQuery: %w", err)
	}

	return serialized, nil
}

func (o TrendingItemsQuery) String() string {
	out := ""
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  threshold=%v\n", o.Threshold)
	out += fmt.Sprintf("  maxRecommendations=%v\n", o.MaxRecommendations)
	out += fmt.Sprintf("  queryParameters=%v\n", o.QueryParameters)
	out += fmt.Sprintf("  facetName=%v\n", o.FacetName)
	out += fmt.Sprintf("  facetValue=%v\n", o.FacetValue)
	out += fmt.Sprintf("  model=%v\n", o.Model)
	out += fmt.Sprintf("  fallbackParameters=%v\n", o.FallbackParameters)
	return fmt.Sprintf("TrendingItemsQuery {\n%s}", out)
}

type NullableTrendingItemsQuery struct {
	value *TrendingItemsQuery
	isSet bool
}

func (v NullableTrendingItemsQuery) Get() *TrendingItemsQuery {
	return v.value
}

func (v *NullableTrendingItemsQuery) Set(val *TrendingItemsQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTrendingItemsQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTrendingItemsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrendingItemsQuery(val *TrendingItemsQuery) *NullableTrendingItemsQuery {
	return &NullableTrendingItemsQuery{value: val, isSet: true}
}

func (v NullableTrendingItemsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTrendingItemsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
