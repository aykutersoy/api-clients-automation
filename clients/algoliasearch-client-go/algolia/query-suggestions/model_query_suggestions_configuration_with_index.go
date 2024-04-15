// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// QuerySuggestionsConfigurationWithIndex Query Suggestions configuration.
type QuerySuggestionsConfigurationWithIndex struct {
	// Algolia indices from which to get the popular searches for query suggestions.
	SourceIndices []SourceIndex `json:"sourceIndices"`
	Languages     *Languages    `json:"languages,omitempty"`
	Exclude       []string      `json:"exclude,omitempty"`
	// Whether to turn on personalized query suggestions.
	EnablePersonalization *bool `json:"enablePersonalization,omitempty"`
	// Whether to include suggestions with special characters.
	AllowSpecialCharacters *bool `json:"allowSpecialCharacters,omitempty"`
	// Name of the Query Suggestions index (case-sensitive).
	IndexName string `json:"indexName"`
}

type QuerySuggestionsConfigurationWithIndexOption func(f *QuerySuggestionsConfigurationWithIndex)

func WithQuerySuggestionsConfigurationWithIndexLanguages(val Languages) QuerySuggestionsConfigurationWithIndexOption {
	return func(f *QuerySuggestionsConfigurationWithIndex) {
		f.Languages = &val
	}
}

func WithQuerySuggestionsConfigurationWithIndexExclude(val []string) QuerySuggestionsConfigurationWithIndexOption {
	return func(f *QuerySuggestionsConfigurationWithIndex) {
		f.Exclude = val
	}
}

func WithQuerySuggestionsConfigurationWithIndexEnablePersonalization(val bool) QuerySuggestionsConfigurationWithIndexOption {
	return func(f *QuerySuggestionsConfigurationWithIndex) {
		f.EnablePersonalization = &val
	}
}

func WithQuerySuggestionsConfigurationWithIndexAllowSpecialCharacters(val bool) QuerySuggestionsConfigurationWithIndexOption {
	return func(f *QuerySuggestionsConfigurationWithIndex) {
		f.AllowSpecialCharacters = &val
	}
}

// NewQuerySuggestionsConfigurationWithIndex instantiates a new QuerySuggestionsConfigurationWithIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQuerySuggestionsConfigurationWithIndex(sourceIndices []SourceIndex, indexName string, opts ...QuerySuggestionsConfigurationWithIndexOption) *QuerySuggestionsConfigurationWithIndex {
	this := &QuerySuggestionsConfigurationWithIndex{}
	this.SourceIndices = sourceIndices
	this.IndexName = indexName
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyQuerySuggestionsConfigurationWithIndex return a pointer to an empty QuerySuggestionsConfigurationWithIndex object.
func NewEmptyQuerySuggestionsConfigurationWithIndex() *QuerySuggestionsConfigurationWithIndex {
	return &QuerySuggestionsConfigurationWithIndex{}
}

// GetSourceIndices returns the SourceIndices field value.
func (o *QuerySuggestionsConfigurationWithIndex) GetSourceIndices() []SourceIndex {
	if o == nil {
		var ret []SourceIndex
		return ret
	}

	return o.SourceIndices
}

// GetSourceIndicesOk returns a tuple with the SourceIndices field value
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfigurationWithIndex) GetSourceIndicesOk() ([]SourceIndex, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIndices, true
}

// SetSourceIndices sets field value.
func (o *QuerySuggestionsConfigurationWithIndex) SetSourceIndices(v []SourceIndex) *QuerySuggestionsConfigurationWithIndex {
	o.SourceIndices = v
	return o
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *QuerySuggestionsConfigurationWithIndex) GetLanguages() Languages {
	if o == nil || o.Languages == nil {
		var ret Languages
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfigurationWithIndex) GetLanguagesOk() (*Languages, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *QuerySuggestionsConfigurationWithIndex) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given Languages and assigns it to the Languages field.
func (o *QuerySuggestionsConfigurationWithIndex) SetLanguages(v *Languages) *QuerySuggestionsConfigurationWithIndex {
	o.Languages = v
	return o
}

// GetExclude returns the Exclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuerySuggestionsConfigurationWithIndex) GetExclude() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *QuerySuggestionsConfigurationWithIndex) GetExcludeOk() ([]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *QuerySuggestionsConfigurationWithIndex) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *QuerySuggestionsConfigurationWithIndex) SetExclude(v []string) *QuerySuggestionsConfigurationWithIndex {
	o.Exclude = v
	return o
}

// GetEnablePersonalization returns the EnablePersonalization field value if set, zero value otherwise.
func (o *QuerySuggestionsConfigurationWithIndex) GetEnablePersonalization() bool {
	if o == nil || o.EnablePersonalization == nil {
		var ret bool
		return ret
	}
	return *o.EnablePersonalization
}

// GetEnablePersonalizationOk returns a tuple with the EnablePersonalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfigurationWithIndex) GetEnablePersonalizationOk() (*bool, bool) {
	if o == nil || o.EnablePersonalization == nil {
		return nil, false
	}
	return o.EnablePersonalization, true
}

// HasEnablePersonalization returns a boolean if a field has been set.
func (o *QuerySuggestionsConfigurationWithIndex) HasEnablePersonalization() bool {
	if o != nil && o.EnablePersonalization != nil {
		return true
	}

	return false
}

// SetEnablePersonalization gets a reference to the given bool and assigns it to the EnablePersonalization field.
func (o *QuerySuggestionsConfigurationWithIndex) SetEnablePersonalization(v bool) *QuerySuggestionsConfigurationWithIndex {
	o.EnablePersonalization = &v
	return o
}

// GetAllowSpecialCharacters returns the AllowSpecialCharacters field value if set, zero value otherwise.
func (o *QuerySuggestionsConfigurationWithIndex) GetAllowSpecialCharacters() bool {
	if o == nil || o.AllowSpecialCharacters == nil {
		var ret bool
		return ret
	}
	return *o.AllowSpecialCharacters
}

// GetAllowSpecialCharactersOk returns a tuple with the AllowSpecialCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfigurationWithIndex) GetAllowSpecialCharactersOk() (*bool, bool) {
	if o == nil || o.AllowSpecialCharacters == nil {
		return nil, false
	}
	return o.AllowSpecialCharacters, true
}

// HasAllowSpecialCharacters returns a boolean if a field has been set.
func (o *QuerySuggestionsConfigurationWithIndex) HasAllowSpecialCharacters() bool {
	if o != nil && o.AllowSpecialCharacters != nil {
		return true
	}

	return false
}

// SetAllowSpecialCharacters gets a reference to the given bool and assigns it to the AllowSpecialCharacters field.
func (o *QuerySuggestionsConfigurationWithIndex) SetAllowSpecialCharacters(v bool) *QuerySuggestionsConfigurationWithIndex {
	o.AllowSpecialCharacters = &v
	return o
}

// GetIndexName returns the IndexName field value.
func (o *QuerySuggestionsConfigurationWithIndex) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfigurationWithIndex) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *QuerySuggestionsConfigurationWithIndex) SetIndexName(v string) *QuerySuggestionsConfigurationWithIndex {
	o.IndexName = v
	return o
}

func (o QuerySuggestionsConfigurationWithIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["sourceIndices"] = o.SourceIndices
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.EnablePersonalization != nil {
		toSerialize["enablePersonalization"] = o.EnablePersonalization
	}
	if o.AllowSpecialCharacters != nil {
		toSerialize["allowSpecialCharacters"] = o.AllowSpecialCharacters
	}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal QuerySuggestionsConfigurationWithIndex: %w", err)
	}

	return serialized, nil
}

func (o QuerySuggestionsConfigurationWithIndex) String() string {
	out := ""
	out += fmt.Sprintf("  sourceIndices=%v\n", o.SourceIndices)
	out += fmt.Sprintf("  languages=%v\n", o.Languages)
	out += fmt.Sprintf("  exclude=%v\n", o.Exclude)
	out += fmt.Sprintf("  enablePersonalization=%v\n", o.EnablePersonalization)
	out += fmt.Sprintf("  allowSpecialCharacters=%v\n", o.AllowSpecialCharacters)
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	return fmt.Sprintf("QuerySuggestionsConfigurationWithIndex {\n%s}", out)
}

type NullableQuerySuggestionsConfigurationWithIndex struct {
	value *QuerySuggestionsConfigurationWithIndex
	isSet bool
}

func (v NullableQuerySuggestionsConfigurationWithIndex) Get() *QuerySuggestionsConfigurationWithIndex {
	return v.value
}

func (v *NullableQuerySuggestionsConfigurationWithIndex) Set(val *QuerySuggestionsConfigurationWithIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySuggestionsConfigurationWithIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySuggestionsConfigurationWithIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySuggestionsConfigurationWithIndex(val *QuerySuggestionsConfigurationWithIndex) *NullableQuerySuggestionsConfigurationWithIndex {
	return &NullableQuerySuggestionsConfigurationWithIndex{value: val, isSet: true}
}

func (v NullableQuerySuggestionsConfigurationWithIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableQuerySuggestionsConfigurationWithIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
