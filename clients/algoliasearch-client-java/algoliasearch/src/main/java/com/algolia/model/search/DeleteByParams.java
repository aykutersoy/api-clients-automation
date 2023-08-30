// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** DeleteByParams */
public class DeleteByParams {

  @JsonProperty("facetFilters")
  private FacetFilters facetFilters;

  @JsonProperty("filters")
  private String filters;

  @JsonProperty("numericFilters")
  private NumericFilters numericFilters;

  @JsonProperty("tagFilters")
  private TagFilters tagFilters;

  @JsonProperty("aroundLatLng")
  private String aroundLatLng;

  @JsonProperty("aroundRadius")
  private AroundRadius aroundRadius;

  @JsonProperty("insideBoundingBox")
  private List<Double> insideBoundingBox;

  @JsonProperty("insidePolygon")
  private List<Double> insidePolygon;

  public DeleteByParams setFacetFilters(FacetFilters facetFilters) {
    this.facetFilters = facetFilters;
    return this;
  }

  /** Get facetFilters */
  @javax.annotation.Nullable
  public FacetFilters getFacetFilters() {
    return facetFilters;
  }

  public DeleteByParams setFilters(String filters) {
    this.filters = filters;
    return this;
  }

  /**
   * [Filter](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/) the
   * query with numeric, facet, or tag filters.
   */
  @javax.annotation.Nullable
  public String getFilters() {
    return filters;
  }

  public DeleteByParams setNumericFilters(NumericFilters numericFilters) {
    this.numericFilters = numericFilters;
    return this;
  }

  /** Get numericFilters */
  @javax.annotation.Nullable
  public NumericFilters getNumericFilters() {
    return numericFilters;
  }

  public DeleteByParams setTagFilters(TagFilters tagFilters) {
    this.tagFilters = tagFilters;
    return this;
  }

  /** Get tagFilters */
  @javax.annotation.Nullable
  public TagFilters getTagFilters() {
    return tagFilters;
  }

  public DeleteByParams setAroundLatLng(String aroundLatLng) {
    this.aroundLatLng = aroundLatLng;
    return this;
  }

  /**
   * Search for entries [around a central
   * location](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filter-around-a-central-point),
   * enabling a geographical search within a circular area.
   */
  @javax.annotation.Nullable
  public String getAroundLatLng() {
    return aroundLatLng;
  }

  public DeleteByParams setAroundRadius(AroundRadius aroundRadius) {
    this.aroundRadius = aroundRadius;
    return this;
  }

  /** Get aroundRadius */
  @javax.annotation.Nullable
  public AroundRadius getAroundRadius() {
    return aroundRadius;
  }

  public DeleteByParams setInsideBoundingBox(List<Double> insideBoundingBox) {
    this.insideBoundingBox = insideBoundingBox;
    return this;
  }

  public DeleteByParams addInsideBoundingBox(Double insideBoundingBoxItem) {
    if (this.insideBoundingBox == null) {
      this.insideBoundingBox = new ArrayList<>();
    }
    this.insideBoundingBox.add(insideBoundingBoxItem);
    return this;
  }

  /**
   * Search inside a [rectangular
   * area](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas)
   * (in geographical coordinates).
   */
  @javax.annotation.Nullable
  public List<Double> getInsideBoundingBox() {
    return insideBoundingBox;
  }

  public DeleteByParams setInsidePolygon(List<Double> insidePolygon) {
    this.insidePolygon = insidePolygon;
    return this;
  }

  public DeleteByParams addInsidePolygon(Double insidePolygonItem) {
    if (this.insidePolygon == null) {
      this.insidePolygon = new ArrayList<>();
    }
    this.insidePolygon.add(insidePolygonItem);
    return this;
  }

  /**
   * Search inside a
   * [polygon](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas)
   * (in geographical coordinates).
   */
  @javax.annotation.Nullable
  public List<Double> getInsidePolygon() {
    return insidePolygon;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeleteByParams deleteByParams = (DeleteByParams) o;
    return (
      Objects.equals(this.facetFilters, deleteByParams.facetFilters) &&
      Objects.equals(this.filters, deleteByParams.filters) &&
      Objects.equals(this.numericFilters, deleteByParams.numericFilters) &&
      Objects.equals(this.tagFilters, deleteByParams.tagFilters) &&
      Objects.equals(this.aroundLatLng, deleteByParams.aroundLatLng) &&
      Objects.equals(this.aroundRadius, deleteByParams.aroundRadius) &&
      Objects.equals(this.insideBoundingBox, deleteByParams.insideBoundingBox) &&
      Objects.equals(this.insidePolygon, deleteByParams.insidePolygon)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(facetFilters, filters, numericFilters, tagFilters, aroundLatLng, aroundRadius, insideBoundingBox, insidePolygon);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeleteByParams {\n");
    sb.append("    facetFilters: ").append(toIndentedString(facetFilters)).append("\n");
    sb.append("    filters: ").append(toIndentedString(filters)).append("\n");
    sb.append("    numericFilters: ").append(toIndentedString(numericFilters)).append("\n");
    sb.append("    tagFilters: ").append(toIndentedString(tagFilters)).append("\n");
    sb.append("    aroundLatLng: ").append(toIndentedString(aroundLatLng)).append("\n");
    sb.append("    aroundRadius: ").append(toIndentedString(aroundRadius)).append("\n");
    sb.append("    insideBoundingBox: ").append(toIndentedString(insideBoundingBox)).append("\n");
    sb.append("    insidePolygon: ").append(toIndentedString(insidePolygon)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
