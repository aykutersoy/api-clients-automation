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

namespace Algolia.Search.Models.Recommend;

/// <summary>
/// RecommendedForYouQuery
/// </summary>
[DataContract(Name = "recommendedForYouQuery")]
[JsonObject(MemberSerialization.OptOut)]
public partial class RecommendedForYouQuery
{

  /// <summary>
  /// Gets or Sets Model
  /// </summary>
  [DataMember(Name = "model", IsRequired = true, EmitDefaultValue = false)]
  public RecommendedForYouModel Model { get; set; }
  /// <summary>
  /// Initializes a new instance of the RecommendedForYouQuery class.
  /// </summary>
  [JsonConstructor]
  public RecommendedForYouQuery() { }
  /// <summary>
  /// Initializes a new instance of the RecommendedForYouQuery class.
  /// </summary>
  /// <param name="indexName">Algolia index name. (required).</param>
  /// <param name="model">model (required).</param>
  public RecommendedForYouQuery(string indexName, RecommendedForYouModel model)
  {
    IndexName = indexName ?? throw new ArgumentNullException(nameof(indexName));
    Model = model;
  }

  /// <summary>
  /// Algolia index name.
  /// </summary>
  /// <value>Algolia index name.</value>
  [DataMember(Name = "indexName", IsRequired = true, EmitDefaultValue = false)]
  public string IndexName { get; set; }

  /// <summary>
  /// Recommendations with a confidence score lower than `threshold` won't appear in results. > **Note**: Each recommendation has a confidence score of 0 to 100. The closer the score is to 100, the more relevant the recommendations are. 
  /// </summary>
  /// <value>Recommendations with a confidence score lower than `threshold` won't appear in results. > **Note**: Each recommendation has a confidence score of 0 to 100. The closer the score is to 100, the more relevant the recommendations are. </value>
  [DataMember(Name = "threshold", EmitDefaultValue = false)]
  public int? Threshold { get; set; }

  /// <summary>
  /// Maximum number of recommendations to retrieve. If 0, all recommendations will be returned.
  /// </summary>
  /// <value>Maximum number of recommendations to retrieve. If 0, all recommendations will be returned.</value>
  [DataMember(Name = "maxRecommendations", EmitDefaultValue = false)]
  public int? MaxRecommendations { get; set; }

  /// <summary>
  /// Gets or Sets QueryParameters
  /// </summary>
  [DataMember(Name = "queryParameters", EmitDefaultValue = false)]
  public RecommendedForYouQueryParameters QueryParameters { get; set; }

  /// <summary>
  /// Gets or Sets FallbackParameters
  /// </summary>
  [DataMember(Name = "fallbackParameters", EmitDefaultValue = false)]
  public RecommendedForYouQueryParameters FallbackParameters { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class RecommendedForYouQuery {\n");
    sb.Append("  IndexName: ").Append(IndexName).Append("\n");
    sb.Append("  Threshold: ").Append(Threshold).Append("\n");
    sb.Append("  MaxRecommendations: ").Append(MaxRecommendations).Append("\n");
    sb.Append("  Model: ").Append(Model).Append("\n");
    sb.Append("  QueryParameters: ").Append(QueryParameters).Append("\n");
    sb.Append("  FallbackParameters: ").Append(FallbackParameters).Append("\n");
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

