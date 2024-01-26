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

namespace Algolia.Search.Models.Search;

/// <summary>
/// UserHighlightResult
/// </summary>
[DataContract(Name = "userHighlightResult")]
[JsonObject(MemberSerialization.OptOut)]
public partial class UserHighlightResult
{
  /// <summary>
  /// Initializes a new instance of the UserHighlightResult class.
  /// </summary>
  [JsonConstructor]
  public UserHighlightResult() { }
  /// <summary>
  /// Initializes a new instance of the UserHighlightResult class.
  /// </summary>
  /// <param name="userID">userID (required).</param>
  /// <param name="clusterName">clusterName (required).</param>
  public UserHighlightResult(HighlightResult userID, HighlightResult clusterName)
  {
    UserID = userID ?? throw new ArgumentNullException(nameof(userID));
    ClusterName = clusterName ?? throw new ArgumentNullException(nameof(clusterName));
  }

  /// <summary>
  /// Gets or Sets UserID
  /// </summary>
  [DataMember(Name = "userID", IsRequired = true, EmitDefaultValue = false)]
  public HighlightResult UserID { get; set; }

  /// <summary>
  /// Gets or Sets ClusterName
  /// </summary>
  [DataMember(Name = "clusterName", IsRequired = true, EmitDefaultValue = false)]
  public HighlightResult ClusterName { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class UserHighlightResult {\n");
    sb.Append("  UserID: ").Append(UserID).Append("\n");
    sb.Append("  ClusterName: ").Append(ClusterName).Append("\n");
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

