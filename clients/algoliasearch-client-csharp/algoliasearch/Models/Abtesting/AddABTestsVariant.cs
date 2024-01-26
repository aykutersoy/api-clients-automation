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
using System.Reflection;
using Algolia.Search.Models;
using Algolia.Search.Models.Common;
using Algolia.Search.Serializer;

namespace Algolia.Search.Models.Abtesting;

/// <summary>
/// AddABTestsVariant
/// </summary>
[JsonConverter(typeof(AddABTestsVariantJsonConverter))]
[DataContract(Name = "AddABTestsVariant")]
public partial class AddABTestsVariant : AbstractSchema
{
  /// <summary>
  /// Initializes a new instance of the AddABTestsVariant class
  /// with a AbTestsVariant
  /// </summary>
  /// <param name="actualInstance">An instance of AbTestsVariant.</param>
  public AddABTestsVariant(AbTestsVariant actualInstance)
  {
    IsNullable = false;
    SchemaType = "oneOf";
    ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
  }

  /// <summary>
  /// Initializes a new instance of the AddABTestsVariant class
  /// with a AbTestsVariantSearchParams
  /// </summary>
  /// <param name="actualInstance">An instance of AbTestsVariantSearchParams.</param>
  public AddABTestsVariant(AbTestsVariantSearchParams actualInstance)
  {
    IsNullable = false;
    SchemaType = "oneOf";
    ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
  }


  /// <summary>
  /// Gets or Sets ActualInstance
  /// </summary>
  public sealed override object ActualInstance { get; set; }

  /// <summary>
  /// Get the actual instance of `AbTestsVariant`. If the actual instance is not `AbTestsVariant`,
  /// the InvalidClassException will be thrown
  /// </summary>
  /// <returns>An instance of AbTestsVariant</returns>
  public AbTestsVariant AsAbTestsVariant()
  {
    return (AbTestsVariant)ActualInstance;
  }

  /// <summary>
  /// Get the actual instance of `AbTestsVariantSearchParams`. If the actual instance is not `AbTestsVariantSearchParams`,
  /// the InvalidClassException will be thrown
  /// </summary>
  /// <returns>An instance of AbTestsVariantSearchParams</returns>
  public AbTestsVariantSearchParams AsAbTestsVariantSearchParams()
  {
    return (AbTestsVariantSearchParams)ActualInstance;
  }


  /// <summary>
  /// Check if the actual instance is of `AbTestsVariant` type.
  /// </summary>
  /// <returns>Whether or not the instance is the type</returns>
  public bool IsAbTestsVariant()
  {
    return ActualInstance.GetType() == typeof(AbTestsVariant);
  }

  /// <summary>
  /// Check if the actual instance is of `AbTestsVariantSearchParams` type.
  /// </summary>
  /// <returns>Whether or not the instance is the type</returns>
  public bool IsAbTestsVariantSearchParams()
  {
    return ActualInstance.GetType() == typeof(AbTestsVariantSearchParams);
  }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    var sb = new StringBuilder();
    sb.Append("class AddABTestsVariant {\n");
    sb.Append("  ActualInstance: ").Append(ActualInstance).Append("\n");
    sb.Append("}\n");
    return sb.ToString();
  }

  /// <summary>
  /// Returns the JSON string presentation of the object
  /// </summary>
  /// <returns>JSON string presentation of the object</returns>
  public override string ToJson()
  {
    return JsonConvert.SerializeObject(ActualInstance, JsonConfig.AlgoliaJsonSerializerSettings);
  }

  /// <summary>
  /// Converts the JSON string into an instance of AddABTestsVariant
  /// </summary>
  /// <param name="jsonString">JSON string</param>
  /// <returns>An instance of AddABTestsVariant</returns>
  public static AddABTestsVariant FromJson(string jsonString)
  {
    AddABTestsVariant newAddABTestsVariant = null;

    if (string.IsNullOrEmpty(jsonString))
    {
      return newAddABTestsVariant;
    }
    try
    {
      return new AddABTestsVariant(JsonConvert.DeserializeObject<AbTestsVariant>(jsonString, AdditionalPropertiesSerializerSettings));
    }
    catch (Exception exception)
    {
      // deserialization failed, try the next one
      System.Diagnostics.Debug.WriteLine($"Failed to deserialize `{jsonString}` into AbTestsVariant: {exception}");
    }
    try
    {
      return new AddABTestsVariant(JsonConvert.DeserializeObject<AbTestsVariantSearchParams>(jsonString, AdditionalPropertiesSerializerSettings));
    }
    catch (Exception exception)
    {
      // deserialization failed, try the next one
      System.Diagnostics.Debug.WriteLine($"Failed to deserialize `{jsonString}` into AbTestsVariantSearchParams: {exception}");
    }

    throw new InvalidDataException($"The JSON string `{jsonString}` cannot be deserialized into any schema defined.");
  }

}

/// <summary>
/// Custom JSON converter for AddABTestsVariant
/// </summary>
public class AddABTestsVariantJsonConverter : JsonConverter
{
  /// <summary>
  /// To write the JSON string
  /// </summary>
  /// <param name="writer">JSON writer</param>
  /// <param name="value">Object to be converted into a JSON string</param>
  /// <param name="serializer">JSON Serializer</param>
  public override void WriteJson(JsonWriter writer, object value, JsonSerializer serializer)
  {
    writer.WriteRawValue((string)(typeof(AddABTestsVariant).GetMethod("ToJson")?.Invoke(value, null)));
  }

  /// <summary>
  /// To convert a JSON string into an object
  /// </summary>
  /// <param name="reader">JSON reader</param>
  /// <param name="objectType">Object type</param>
  /// <param name="existingValue">Existing value</param>
  /// <param name="serializer">JSON Serializer</param>
  /// <returns>The object converted from the JSON string</returns>
  public override object ReadJson(JsonReader reader, Type objectType, object existingValue, JsonSerializer serializer)
  {
    if (reader.TokenType != JsonToken.Null)
    {
      return objectType.GetMethod("FromJson")?.Invoke(null, new object[] { JObject.Load(reader).ToString(Formatting.None) });
    }
    return null;
  }

  /// <summary>
  /// Check if the object can be converted
  /// </summary>
  /// <param name="objectType">Object type</param>
  /// <returns>True if the object can be converted</returns>
  public override bool CanConvert(Type objectType)
  {
    return false;
  }
}

