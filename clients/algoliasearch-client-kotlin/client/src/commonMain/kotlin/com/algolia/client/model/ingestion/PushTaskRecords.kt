/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * PushTaskRecords
 *
 * @param objectID Unique record identifier.
 */
@Serializable(PushTaskRecordsSerializer::class)
public data class PushTaskRecords(

  /** Unique record identifier. */
  val objectID: String,

  val additionalProperties: Map<String, JsonElement>? = null,
)

internal object PushTaskRecordsSerializer : KSerializer<PushTaskRecords> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("PushTaskRecords") {
    element<String>("objectID")
  }

  override fun deserialize(decoder: Decoder): PushTaskRecords {
    val input = decoder.asJsonDecoder()
    val tree = input.decodeJsonObject()
    return PushTaskRecords(
      objectID = tree.getValue("objectID").let { input.json.decodeFromJsonElement(it) },
      additionalProperties = tree.filterKeys { it !in descriptor.elementNames },
    )
  }

  override fun serialize(encoder: Encoder, value: PushTaskRecords) {
    val output = encoder.asJsonEncoder()
    val json = buildJsonObject {
      put("objectID", output.json.encodeToJsonElement(value.objectID))
      value.additionalProperties?.onEach { (key, element) -> put(key, element) }
    }
    (encoder as JsonEncoder).encodeJsonElement(json)
  }
}
