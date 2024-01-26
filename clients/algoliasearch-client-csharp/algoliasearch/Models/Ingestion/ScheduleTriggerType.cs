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

namespace Algolia.Search.Models.Ingestion;

/// <summary>
/// A task which is triggered by a schedule (cron expression).
/// </summary>
/// <value>A task which is triggered by a schedule (cron expression).</value>
[JsonConverter(typeof(StringEnumConverter))]
public enum ScheduleTriggerType
{
  /// <summary>
  /// Enum Schedule for value: schedule
  /// </summary>
  [EnumMember(Value = "schedule")]
  Schedule = 1
}

