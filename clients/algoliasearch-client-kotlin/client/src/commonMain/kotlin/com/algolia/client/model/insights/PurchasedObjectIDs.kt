/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.insights

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * Use this event to track when users make a purchase unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track purchase events related to Algolia requests, use the \"Purchased object IDs after search\" event.
 *
 * @param eventName Event name, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
 * @param eventType
 * @param eventSubtype
 * @param index Index name (case-sensitive) to which the event's items belong.
 * @param objectIDs Object IDs of the records that are part of the event.
 * @param userToken Anonymous or pseudonymous user identifier.  Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
 * @param authenticatedUserToken Identifier for authenticated users.  When the user signs in, you can get an identifier from your system and send it as `authenticatedUserToken`. This lets you keep using the `userToken` from before the user signed in, while providing a reliable way to identify users across sessions. Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
 * @param currency Three-letter [currency code](https://www.iso.org/iso-4217-currency-codes.html).
 * @param objectData Extra information about the records involved in a purchase or add-to-cart event.  If specified, it must have the same length as `objectIDs`.
 * @param timestamp Timestamp of the event, measured in milliseconds since the Unix epoch. By default, the Insights API uses the time it receives an event as its timestamp.
 * @param `value`
 */
@Serializable
public data class PurchasedObjectIDs(

  /** Event name, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.  */
  @SerialName(value = "eventName") val eventName: String,

  @SerialName(value = "eventType") val eventType: ConversionEvent,

  @SerialName(value = "eventSubtype") val eventSubtype: PurchaseEvent,

  /** Index name (case-sensitive) to which the event's items belong. */
  @SerialName(value = "index") val index: String,

  /** Object IDs of the records that are part of the event. */
  @SerialName(value = "objectIDs") val objectIDs: List<String>,

  /** Anonymous or pseudonymous user identifier.  Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).  */
  @SerialName(value = "userToken") val userToken: String,

  /** Identifier for authenticated users.  When the user signs in, you can get an identifier from your system and send it as `authenticatedUserToken`. This lets you keep using the `userToken` from before the user signed in, while providing a reliable way to identify users across sessions. Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).  */
  @SerialName(value = "authenticatedUserToken") val authenticatedUserToken: String? = null,

  /** Three-letter [currency code](https://www.iso.org/iso-4217-currency-codes.html). */
  @SerialName(value = "currency") val currency: String? = null,

  /** Extra information about the records involved in a purchase or add-to-cart event.  If specified, it must have the same length as `objectIDs`.  */
  @SerialName(value = "objectData") val objectData: List<ObjectData>? = null,

  /** Timestamp of the event, measured in milliseconds since the Unix epoch. By default, the Insights API uses the time it receives an event as its timestamp.  */
  @SerialName(value = "timestamp") val timestamp: Long? = null,

  @SerialName(value = "value") val `value`: Value? = null,
) : EventsItems
