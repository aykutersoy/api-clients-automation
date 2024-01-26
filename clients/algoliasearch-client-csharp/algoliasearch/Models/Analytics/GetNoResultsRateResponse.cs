//
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
//
using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using Algolia.Search.Models;
using Algolia.Search.Models.Common;
using Algolia.Search.Serializer;

namespace Algolia.Search.Models.Analytics;

/// <summary>
/// GetNoResultsRateResponse
/// </summary>
[DataContract(Name = "getNoResultsRateResponse")]
[JsonObject(MemberSerialization.OptOut)]
public partial class GetNoResultsRateResponse
{
  /// <summary>
  /// Initializes a new instance of the GetNoResultsRateResponse class.
  /// </summary>
  [JsonConstructor]
  public GetNoResultsRateResponse() { }
  /// <summary>
  /// Initializes a new instance of the GetNoResultsRateResponse class.
  /// </summary>
  /// <param name="rate">[Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).  (required).</param>
  /// <param name="count">Number of occurrences. (required).</param>
  /// <param name="noResultCount">Number of occurrences. (required).</param>
  /// <param name="dates">Overall count of searches without results plus a daily breakdown. (required).</param>
  public GetNoResultsRateResponse(double rate, int? count, int? noResultCount, List<NoResultsRateEvent> dates)
  {
    Rate = rate;
    Count = count;
    NoResultCount = noResultCount;
    Dates = dates ?? throw new ArgumentNullException(nameof(dates));
  }

  /// <summary>
  /// [Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate). 
  /// </summary>
  /// <value>[Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate). </value>
  [DataMember(Name = "rate", IsRequired = true, EmitDefaultValue = false)]
  public double Rate { get; set; }

  /// <summary>
  /// Number of occurrences.
  /// </summary>
  /// <value>Number of occurrences.</value>
  [DataMember(Name = "count", IsRequired = true, EmitDefaultValue = false)]
  public int? Count { get; set; }

  /// <summary>
  /// Number of occurrences.
  /// </summary>
  /// <value>Number of occurrences.</value>
  [DataMember(Name = "noResultCount", IsRequired = true, EmitDefaultValue = false)]
  public int? NoResultCount { get; set; }

  /// <summary>
  /// Overall count of searches without results plus a daily breakdown.
  /// </summary>
  /// <value>Overall count of searches without results plus a daily breakdown.</value>
  [DataMember(Name = "dates", IsRequired = true, EmitDefaultValue = false)]
  public List<NoResultsRateEvent> Dates { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class GetNoResultsRateResponse {\n");
    sb.Append("  Rate: ").Append(Rate).Append("\n");
    sb.Append("  Count: ").Append(Count).Append("\n");
    sb.Append("  NoResultCount: ").Append(NoResultCount).Append("\n");
    sb.Append("  Dates: ").Append(Dates).Append("\n");
    sb.Append("}\n");
    return sb.ToString();
  }

  /// <summary>
  /// Returns the JSON string presentation of the object
  /// </summary>
  /// <returns>JSON string presentation of the object</returns>
  public virtual string ToJson()
  {
    return JsonConvert.SerializeObject(this, Formatting.Indented);
  }

}

