// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// RenderingContent Extra data that can be used in the search UI.  You can use this to control aspects of your search UI, such as, the order of facet names and values without changing your frontend code.
type RenderingContent struct {
	FacetOrdering *FacetOrdering `json:"facetOrdering,omitempty"`
}

type RenderingContentOption func(f *RenderingContent)

func WithRenderingContentFacetOrdering(val FacetOrdering) RenderingContentOption {
	return func(f *RenderingContent) {
		f.FacetOrdering = &val
	}
}

// NewRenderingContent instantiates a new RenderingContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRenderingContent(opts ...RenderingContentOption) *RenderingContent {
	this := &RenderingContent{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyRenderingContent return a pointer to an empty RenderingContent object.
func NewEmptyRenderingContent() *RenderingContent {
	return &RenderingContent{}
}

// GetFacetOrdering returns the FacetOrdering field value if set, zero value otherwise.
func (o *RenderingContent) GetFacetOrdering() FacetOrdering {
	if o == nil || o.FacetOrdering == nil {
		var ret FacetOrdering
		return ret
	}
	return *o.FacetOrdering
}

// GetFacetOrderingOk returns a tuple with the FacetOrdering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderingContent) GetFacetOrderingOk() (*FacetOrdering, bool) {
	if o == nil || o.FacetOrdering == nil {
		return nil, false
	}
	return o.FacetOrdering, true
}

// HasFacetOrdering returns a boolean if a field has been set.
func (o *RenderingContent) HasFacetOrdering() bool {
	if o != nil && o.FacetOrdering != nil {
		return true
	}

	return false
}

// SetFacetOrdering gets a reference to the given FacetOrdering and assigns it to the FacetOrdering field.
func (o *RenderingContent) SetFacetOrdering(v *FacetOrdering) *RenderingContent {
	o.FacetOrdering = v
	return o
}

func (o RenderingContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.FacetOrdering != nil {
		toSerialize["facetOrdering"] = o.FacetOrdering
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal RenderingContent: %w", err)
	}

	return serialized, nil
}

func (o RenderingContent) String() string {
	out := ""
	out += fmt.Sprintf("  facetOrdering=%v\n", o.FacetOrdering)
	return fmt.Sprintf("RenderingContent {\n%s}", out)
}

type NullableRenderingContent struct {
	value *RenderingContent
	isSet bool
}

func (v NullableRenderingContent) Get() *RenderingContent {
	return v.value
}

func (v *NullableRenderingContent) Set(val *RenderingContent) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderingContent) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderingContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderingContent(val *RenderingContent) *NullableRenderingContent {
	return &NullableRenderingContent{value: val, isSet: true}
}

func (v NullableRenderingContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableRenderingContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
