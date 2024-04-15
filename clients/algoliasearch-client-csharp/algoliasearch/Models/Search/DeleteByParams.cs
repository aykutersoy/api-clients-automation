//
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
//
using System;
using System.Text;
using System.Linq;
using System.Text.Json.Serialization;
using System.Collections.Generic;
using Algolia.Search.Serializer;
using System.Text.Json;

namespace Algolia.Search.Models.Search;

/// <summary>
/// DeleteByParams
/// </summary>
public partial class DeleteByParams
{
  /// <summary>
  /// Initializes a new instance of the DeleteByParams class.
  /// </summary>
  public DeleteByParams()
  {
  }

  /// <summary>
  /// Gets or Sets FacetFilters
  /// </summary>
  [JsonPropertyName("facetFilters")]
  public FacetFilters FacetFilters { get; set; }

  /// <summary>
  /// Filter expression to only include items that match the filter criteria in the response.  You can use these filter expressions:  - **Numeric filters.** `<facet> <op> <number>`, where `<op>` is one of `<`, `<=`, `=`, `!=`, `>`, `>=`. - **Ranges.** `<facet>:<lower> TO <upper>` where `<lower>` and `<upper>` are the lower and upper limits of the range (inclusive). - **Facet filters.** `<facet>:<value>` where `<facet>` is a facet attribute (case-sensitive) and `<value>` a facet value. - **Tag filters.** `_tags:<value>` or just `<value>` (case-sensitive). - **Boolean filters.** `<facet>: true | false`.  You can combine filters with `AND`, `OR`, and `NOT` operators with the following restrictions:  - You can only combine filters of the same type with `OR`.   **Not supported:** `facet:value OR num > 3`. - You can't use `NOT` with combinations of filters.   **Not supported:** `NOT(facet:value OR facet:value)` - You can't combine conjunctions (`AND`) with `OR`.   **Not supported:** `facet:value OR (facet:value AND facet:value)`  Use quotes around your filters, if the facet attribute name or facet value has spaces, keywords (`OR`, `AND`, `NOT`), or quotes. If a facet attribute is an array, the filter matches if it matches at least one element of the array.  For more information, see [Filters](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/). 
  /// </summary>
  /// <value>Filter expression to only include items that match the filter criteria in the response.  You can use these filter expressions:  - **Numeric filters.** `<facet> <op> <number>`, where `<op>` is one of `<`, `<=`, `=`, `!=`, `>`, `>=`. - **Ranges.** `<facet>:<lower> TO <upper>` where `<lower>` and `<upper>` are the lower and upper limits of the range (inclusive). - **Facet filters.** `<facet>:<value>` where `<facet>` is a facet attribute (case-sensitive) and `<value>` a facet value. - **Tag filters.** `_tags:<value>` or just `<value>` (case-sensitive). - **Boolean filters.** `<facet>: true | false`.  You can combine filters with `AND`, `OR`, and `NOT` operators with the following restrictions:  - You can only combine filters of the same type with `OR`.   **Not supported:** `facet:value OR num > 3`. - You can't use `NOT` with combinations of filters.   **Not supported:** `NOT(facet:value OR facet:value)` - You can't combine conjunctions (`AND`) with `OR`.   **Not supported:** `facet:value OR (facet:value AND facet:value)`  Use quotes around your filters, if the facet attribute name or facet value has spaces, keywords (`OR`, `AND`, `NOT`), or quotes. If a facet attribute is an array, the filter matches if it matches at least one element of the array.  For more information, see [Filters](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/). </value>
  [JsonPropertyName("filters")]
  public string Filters { get; set; }

  /// <summary>
  /// Gets or Sets NumericFilters
  /// </summary>
  [JsonPropertyName("numericFilters")]
  public NumericFilters NumericFilters { get; set; }

  /// <summary>
  /// Gets or Sets TagFilters
  /// </summary>
  [JsonPropertyName("tagFilters")]
  public TagFilters TagFilters { get; set; }

  /// <summary>
  /// Coordinates for the center of a circle, expressed as a comma-separated string of latitude and longitude.  Only records included within circle around this central location are included in the results. The radius of the circle is determined by the `aroundRadius` and `minimumAroundRadius` settings. This parameter is ignored if you also specify `insidePolygon` or `insideBoundingBox`. 
  /// </summary>
  /// <value>Coordinates for the center of a circle, expressed as a comma-separated string of latitude and longitude.  Only records included within circle around this central location are included in the results. The radius of the circle is determined by the `aroundRadius` and `minimumAroundRadius` settings. This parameter is ignored if you also specify `insidePolygon` or `insideBoundingBox`. </value>
  [JsonPropertyName("aroundLatLng")]
  public string AroundLatLng { get; set; }

  /// <summary>
  /// Gets or Sets AroundRadius
  /// </summary>
  [JsonPropertyName("aroundRadius")]
  public AroundRadius AroundRadius { get; set; }

  /// <summary>
  /// Coordinates for a rectangular area in which to search.  Each bounding box is defined by the two opposite points of its diagonal, and expressed as latitude and longitude pair: `[p1 lat, p1 long, p2 lat, p2 long]`. Provide multiple bounding boxes as nested arrays. For more information, see [rectangular area](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). 
  /// </summary>
  /// <value>Coordinates for a rectangular area in which to search.  Each bounding box is defined by the two opposite points of its diagonal, and expressed as latitude and longitude pair: `[p1 lat, p1 long, p2 lat, p2 long]`. Provide multiple bounding boxes as nested arrays. For more information, see [rectangular area](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). </value>
  [JsonPropertyName("insideBoundingBox")]
  public List<List<double>> InsideBoundingBox { get; set; }

  /// <summary>
  /// Coordinates of a polygon in which to search.  Polygons are defined by 3 to 10,000 points. Each point is represented by its latitude and longitude. Provide multiple polygons as nested arrays. For more information, see [filtering inside polygons](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). This parameter is ignored if you also specify `insideBoundingBox`. 
  /// </summary>
  /// <value>Coordinates of a polygon in which to search.  Polygons are defined by 3 to 10,000 points. Each point is represented by its latitude and longitude. Provide multiple polygons as nested arrays. For more information, see [filtering inside polygons](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). This parameter is ignored if you also specify `insideBoundingBox`. </value>
  [JsonPropertyName("insidePolygon")]
  public List<List<double>> InsidePolygon { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class DeleteByParams {\n");
    sb.Append("  FacetFilters: ").Append(FacetFilters).Append("\n");
    sb.Append("  Filters: ").Append(Filters).Append("\n");
    sb.Append("  NumericFilters: ").Append(NumericFilters).Append("\n");
    sb.Append("  TagFilters: ").Append(TagFilters).Append("\n");
    sb.Append("  AroundLatLng: ").Append(AroundLatLng).Append("\n");
    sb.Append("  AroundRadius: ").Append(AroundRadius).Append("\n");
    sb.Append("  InsideBoundingBox: ").Append(InsideBoundingBox).Append("\n");
    sb.Append("  InsidePolygon: ").Append(InsidePolygon).Append("\n");
    sb.Append("}\n");
    return sb.ToString();
  }

  /// <summary>
  /// Returns the JSON string presentation of the object
  /// </summary>
  /// <returns>JSON string presentation of the object</returns>
  public virtual string ToJson()
  {
    return JsonSerializer.Serialize(this, JsonConfig.Options);
  }

  /// <summary>
  /// Returns true if objects are equal
  /// </summary>
  /// <param name="obj">Object to be compared</param>
  /// <returns>Boolean</returns>
  public override bool Equals(object obj)
  {
    if (obj is not DeleteByParams input)
    {
      return false;
    }

    return
        (FacetFilters == input.FacetFilters || (FacetFilters != null && FacetFilters.Equals(input.FacetFilters))) &&
        (Filters == input.Filters || (Filters != null && Filters.Equals(input.Filters))) &&
        (NumericFilters == input.NumericFilters || (NumericFilters != null && NumericFilters.Equals(input.NumericFilters))) &&
        (TagFilters == input.TagFilters || (TagFilters != null && TagFilters.Equals(input.TagFilters))) &&
        (AroundLatLng == input.AroundLatLng || (AroundLatLng != null && AroundLatLng.Equals(input.AroundLatLng))) &&
        (AroundRadius == input.AroundRadius || (AroundRadius != null && AroundRadius.Equals(input.AroundRadius))) &&
        (InsideBoundingBox == input.InsideBoundingBox || InsideBoundingBox != null && input.InsideBoundingBox != null && InsideBoundingBox.SequenceEqual(input.InsideBoundingBox)) &&
        (InsidePolygon == input.InsidePolygon || InsidePolygon != null && input.InsidePolygon != null && InsidePolygon.SequenceEqual(input.InsidePolygon));
  }

  /// <summary>
  /// Gets the hash code
  /// </summary>
  /// <returns>Hash code</returns>
  public override int GetHashCode()
  {
    unchecked // Overflow is fine, just wrap
    {
      int hashCode = 41;
      if (FacetFilters != null)
      {
        hashCode = (hashCode * 59) + FacetFilters.GetHashCode();
      }
      if (Filters != null)
      {
        hashCode = (hashCode * 59) + Filters.GetHashCode();
      }
      if (NumericFilters != null)
      {
        hashCode = (hashCode * 59) + NumericFilters.GetHashCode();
      }
      if (TagFilters != null)
      {
        hashCode = (hashCode * 59) + TagFilters.GetHashCode();
      }
      if (AroundLatLng != null)
      {
        hashCode = (hashCode * 59) + AroundLatLng.GetHashCode();
      }
      if (AroundRadius != null)
      {
        hashCode = (hashCode * 59) + AroundRadius.GetHashCode();
      }
      if (InsideBoundingBox != null)
      {
        hashCode = (hashCode * 59) + InsideBoundingBox.GetHashCode();
      }
      if (InsidePolygon != null)
      {
        hashCode = (hashCode * 59) + InsidePolygon.GetHashCode();
      }
      return hashCode;
    }
  }

}

