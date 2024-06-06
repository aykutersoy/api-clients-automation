// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// DeleteByParams struct for DeleteByParams.
type DeleteByParams struct {
	FacetFilters *FacetFilters `json:"facetFilters,omitempty"`
	// Filter expression to only include items that match the filter criteria in the response.  You can use these filter expressions:  - **Numeric filters.** `<facet> <op> <number>`, where `<op>` is one of `<`, `<=`, `=`, `!=`, `>`, `>=`. - **Ranges.** `<facet>:<lower> TO <upper>` where `<lower>` and `<upper>` are the lower and upper limits of the range (inclusive). - **Facet filters.** `<facet>:<value>` where `<facet>` is a facet attribute (case-sensitive) and `<value>` a facet value. - **Tag filters.** `_tags:<value>` or just `<value>` (case-sensitive). - **Boolean filters.** `<facet>: true | false`.  You can combine filters with `AND`, `OR`, and `NOT` operators with the following restrictions:  - You can only combine filters of the same type with `OR`.   **Not supported:** `facet:value OR num > 3`. - You can't use `NOT` with combinations of filters.   **Not supported:** `NOT(facet:value OR facet:value)` - You can't combine conjunctions (`AND`) with `OR`.   **Not supported:** `facet:value OR (facet:value AND facet:value)`  Use quotes around your filters, if the facet attribute name or facet value has spaces, keywords (`OR`, `AND`, `NOT`), or quotes. If a facet attribute is an array, the filter matches if it matches at least one element of the array.  For more information, see [Filters](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/).
	Filters        *string         `json:"filters,omitempty"`
	NumericFilters *NumericFilters `json:"numericFilters,omitempty"`
	TagFilters     *TagFilters     `json:"tagFilters,omitempty"`
	// Coordinates for the center of a circle, expressed as a comma-separated string of latitude and longitude.  Only records included within circle around this central location are included in the results. The radius of the circle is determined by the `aroundRadius` and `minimumAroundRadius` settings. This parameter is ignored if you also specify `insidePolygon` or `insideBoundingBox`.
	AroundLatLng *string       `json:"aroundLatLng,omitempty"`
	AroundRadius *AroundRadius `json:"aroundRadius,omitempty"`
	// Coordinates for a rectangular area in which to search.  Each bounding box is defined by the two opposite points of its diagonal, and expressed as latitude and longitude pair: `[p1 lat, p1 long, p2 lat, p2 long]`. Provide multiple bounding boxes as nested arrays. For more information, see [rectangular area](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas).
	InsideBoundingBox [][]float64 `json:"insideBoundingBox,omitempty"`
	// Coordinates of a polygon in which to search.  Polygons are defined by 3 to 10,000 points. Each point is represented by its latitude and longitude. Provide multiple polygons as nested arrays. For more information, see [filtering inside polygons](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). This parameter is ignored if you also specify `insideBoundingBox`.
	InsidePolygon [][]float64 `json:"insidePolygon,omitempty"`
}

type DeleteByParamsOption func(f *DeleteByParams)

func WithDeleteByParamsFacetFilters(val FacetFilters) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.FacetFilters = &val
	}
}

func WithDeleteByParamsFilters(val string) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.Filters = &val
	}
}

func WithDeleteByParamsNumericFilters(val NumericFilters) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.NumericFilters = &val
	}
}

func WithDeleteByParamsTagFilters(val TagFilters) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.TagFilters = &val
	}
}

func WithDeleteByParamsAroundLatLng(val string) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.AroundLatLng = &val
	}
}

func WithDeleteByParamsAroundRadius(val AroundRadius) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.AroundRadius = &val
	}
}

func WithDeleteByParamsInsideBoundingBox(val [][]float64) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.InsideBoundingBox = val
	}
}

func WithDeleteByParamsInsidePolygon(val [][]float64) DeleteByParamsOption {
	return func(f *DeleteByParams) {
		f.InsidePolygon = val
	}
}

// NewDeleteByParams instantiates a new DeleteByParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteByParams(opts ...DeleteByParamsOption) *DeleteByParams {
	this := &DeleteByParams{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyDeleteByParams return a pointer to an empty DeleteByParams object.
func NewEmptyDeleteByParams() *DeleteByParams {
	return &DeleteByParams{}
}

// GetFacetFilters returns the FacetFilters field value if set, zero value otherwise.
func (o *DeleteByParams) GetFacetFilters() FacetFilters {
	if o == nil || o.FacetFilters == nil {
		var ret FacetFilters
		return ret
	}
	return *o.FacetFilters
}

// GetFacetFiltersOk returns a tuple with the FacetFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetFacetFiltersOk() (*FacetFilters, bool) {
	if o == nil || o.FacetFilters == nil {
		return nil, false
	}
	return o.FacetFilters, true
}

// HasFacetFilters returns a boolean if a field has been set.
func (o *DeleteByParams) HasFacetFilters() bool {
	if o != nil && o.FacetFilters != nil {
		return true
	}

	return false
}

// SetFacetFilters gets a reference to the given FacetFilters and assigns it to the FacetFilters field.
func (o *DeleteByParams) SetFacetFilters(v *FacetFilters) *DeleteByParams {
	o.FacetFilters = v
	return o
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DeleteByParams) GetFilters() string {
	if o == nil || o.Filters == nil {
		var ret string
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetFiltersOk() (*string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DeleteByParams) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given string and assigns it to the Filters field.
func (o *DeleteByParams) SetFilters(v string) *DeleteByParams {
	o.Filters = &v
	return o
}

// GetNumericFilters returns the NumericFilters field value if set, zero value otherwise.
func (o *DeleteByParams) GetNumericFilters() NumericFilters {
	if o == nil || o.NumericFilters == nil {
		var ret NumericFilters
		return ret
	}
	return *o.NumericFilters
}

// GetNumericFiltersOk returns a tuple with the NumericFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetNumericFiltersOk() (*NumericFilters, bool) {
	if o == nil || o.NumericFilters == nil {
		return nil, false
	}
	return o.NumericFilters, true
}

// HasNumericFilters returns a boolean if a field has been set.
func (o *DeleteByParams) HasNumericFilters() bool {
	if o != nil && o.NumericFilters != nil {
		return true
	}

	return false
}

// SetNumericFilters gets a reference to the given NumericFilters and assigns it to the NumericFilters field.
func (o *DeleteByParams) SetNumericFilters(v *NumericFilters) *DeleteByParams {
	o.NumericFilters = v
	return o
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *DeleteByParams) GetTagFilters() TagFilters {
	if o == nil || o.TagFilters == nil {
		var ret TagFilters
		return ret
	}
	return *o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetTagFiltersOk() (*TagFilters, bool) {
	if o == nil || o.TagFilters == nil {
		return nil, false
	}
	return o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *DeleteByParams) HasTagFilters() bool {
	if o != nil && o.TagFilters != nil {
		return true
	}

	return false
}

// SetTagFilters gets a reference to the given TagFilters and assigns it to the TagFilters field.
func (o *DeleteByParams) SetTagFilters(v *TagFilters) *DeleteByParams {
	o.TagFilters = v
	return o
}

// GetAroundLatLng returns the AroundLatLng field value if set, zero value otherwise.
func (o *DeleteByParams) GetAroundLatLng() string {
	if o == nil || o.AroundLatLng == nil {
		var ret string
		return ret
	}
	return *o.AroundLatLng
}

// GetAroundLatLngOk returns a tuple with the AroundLatLng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetAroundLatLngOk() (*string, bool) {
	if o == nil || o.AroundLatLng == nil {
		return nil, false
	}
	return o.AroundLatLng, true
}

// HasAroundLatLng returns a boolean if a field has been set.
func (o *DeleteByParams) HasAroundLatLng() bool {
	if o != nil && o.AroundLatLng != nil {
		return true
	}

	return false
}

// SetAroundLatLng gets a reference to the given string and assigns it to the AroundLatLng field.
func (o *DeleteByParams) SetAroundLatLng(v string) *DeleteByParams {
	o.AroundLatLng = &v
	return o
}

// GetAroundRadius returns the AroundRadius field value if set, zero value otherwise.
func (o *DeleteByParams) GetAroundRadius() AroundRadius {
	if o == nil || o.AroundRadius == nil {
		var ret AroundRadius
		return ret
	}
	return *o.AroundRadius
}

// GetAroundRadiusOk returns a tuple with the AroundRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetAroundRadiusOk() (*AroundRadius, bool) {
	if o == nil || o.AroundRadius == nil {
		return nil, false
	}
	return o.AroundRadius, true
}

// HasAroundRadius returns a boolean if a field has been set.
func (o *DeleteByParams) HasAroundRadius() bool {
	if o != nil && o.AroundRadius != nil {
		return true
	}

	return false
}

// SetAroundRadius gets a reference to the given AroundRadius and assigns it to the AroundRadius field.
func (o *DeleteByParams) SetAroundRadius(v *AroundRadius) *DeleteByParams {
	o.AroundRadius = v
	return o
}

// GetInsideBoundingBox returns the InsideBoundingBox field value if set, zero value otherwise.
func (o *DeleteByParams) GetInsideBoundingBox() [][]float64 {
	if o == nil || o.InsideBoundingBox == nil {
		var ret [][]float64
		return ret
	}
	return o.InsideBoundingBox
}

// GetInsideBoundingBoxOk returns a tuple with the InsideBoundingBox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetInsideBoundingBoxOk() ([][]float64, bool) {
	if o == nil || o.InsideBoundingBox == nil {
		return nil, false
	}
	return o.InsideBoundingBox, true
}

// HasInsideBoundingBox returns a boolean if a field has been set.
func (o *DeleteByParams) HasInsideBoundingBox() bool {
	if o != nil && o.InsideBoundingBox != nil {
		return true
	}

	return false
}

// SetInsideBoundingBox gets a reference to the given [][]float64 and assigns it to the InsideBoundingBox field.
func (o *DeleteByParams) SetInsideBoundingBox(v [][]float64) *DeleteByParams {
	o.InsideBoundingBox = v
	return o
}

// GetInsidePolygon returns the InsidePolygon field value if set, zero value otherwise.
func (o *DeleteByParams) GetInsidePolygon() [][]float64 {
	if o == nil || o.InsidePolygon == nil {
		var ret [][]float64
		return ret
	}
	return o.InsidePolygon
}

// GetInsidePolygonOk returns a tuple with the InsidePolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteByParams) GetInsidePolygonOk() ([][]float64, bool) {
	if o == nil || o.InsidePolygon == nil {
		return nil, false
	}
	return o.InsidePolygon, true
}

// HasInsidePolygon returns a boolean if a field has been set.
func (o *DeleteByParams) HasInsidePolygon() bool {
	if o != nil && o.InsidePolygon != nil {
		return true
	}

	return false
}

// SetInsidePolygon gets a reference to the given [][]float64 and assigns it to the InsidePolygon field.
func (o *DeleteByParams) SetInsidePolygon(v [][]float64) *DeleteByParams {
	o.InsidePolygon = v
	return o
}

func (o DeleteByParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.FacetFilters != nil {
		toSerialize["facetFilters"] = o.FacetFilters
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.NumericFilters != nil {
		toSerialize["numericFilters"] = o.NumericFilters
	}
	if o.TagFilters != nil {
		toSerialize["tagFilters"] = o.TagFilters
	}
	if o.AroundLatLng != nil {
		toSerialize["aroundLatLng"] = o.AroundLatLng
	}
	if o.AroundRadius != nil {
		toSerialize["aroundRadius"] = o.AroundRadius
	}
	if o.InsideBoundingBox != nil {
		toSerialize["insideBoundingBox"] = o.InsideBoundingBox
	}
	if o.InsidePolygon != nil {
		toSerialize["insidePolygon"] = o.InsidePolygon
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DeleteByParams: %w", err)
	}

	return serialized, nil
}

func (o DeleteByParams) String() string {
	out := ""
	out += fmt.Sprintf("  facetFilters=%v\n", o.FacetFilters)
	out += fmt.Sprintf("  filters=%v\n", o.Filters)
	out += fmt.Sprintf("  numericFilters=%v\n", o.NumericFilters)
	out += fmt.Sprintf("  tagFilters=%v\n", o.TagFilters)
	out += fmt.Sprintf("  aroundLatLng=%v\n", o.AroundLatLng)
	out += fmt.Sprintf("  aroundRadius=%v\n", o.AroundRadius)
	out += fmt.Sprintf("  insideBoundingBox=%v\n", o.InsideBoundingBox)
	out += fmt.Sprintf("  insidePolygon=%v\n", o.InsidePolygon)
	return fmt.Sprintf("DeleteByParams {\n%s}", out)
}

type NullableDeleteByParams struct {
	value *DeleteByParams
	isSet bool
}

func (v NullableDeleteByParams) Get() *DeleteByParams {
	return v.value
}

func (v *NullableDeleteByParams) Set(val *DeleteByParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteByParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteByParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteByParams(val *DeleteByParams) *NullableDeleteByParams {
	return &NullableDeleteByParams{value: val, isSet: true}
}

func (v NullableDeleteByParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDeleteByParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
