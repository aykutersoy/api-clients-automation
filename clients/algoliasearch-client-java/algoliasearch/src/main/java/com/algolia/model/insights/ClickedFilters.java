// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.insights;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** Use this event to track when users click facet filters in your user interface. */
@JsonDeserialize(as = ClickedFilters.class)
public class ClickedFilters implements EventsItems {

  @JsonProperty("eventName")
  private String eventName;

  @JsonProperty("eventType")
  private ClickEvent eventType;

  @JsonProperty("index")
  private String index;

  @JsonProperty("filters")
  private List<String> filters = new ArrayList<>();

  @JsonProperty("userToken")
  private String userToken;

  @JsonProperty("authenticatedUserToken")
  private String authenticatedUserToken;

  @JsonProperty("timestamp")
  private Long timestamp;

  public ClickedFilters setEventName(String eventName) {
    this.eventName = eventName;
    return this;
  }

  /**
   * Event name, up to 64 ASCII characters. Consider naming events consistently—for example, by
   * adopting Segment's
   * [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework)
   * framework.
   */
  @javax.annotation.Nonnull
  public String getEventName() {
    return eventName;
  }

  public ClickedFilters setEventType(ClickEvent eventType) {
    this.eventType = eventType;
    return this;
  }

  /** Get eventType */
  @javax.annotation.Nonnull
  public ClickEvent getEventType() {
    return eventType;
  }

  public ClickedFilters setIndex(String index) {
    this.index = index;
    return this;
  }

  /** Index name (case-sensitive) to which the event's items belong. */
  @javax.annotation.Nonnull
  public String getIndex() {
    return index;
  }

  public ClickedFilters setFilters(List<String> filters) {
    this.filters = filters;
    return this;
  }

  public ClickedFilters addFilters(String filtersItem) {
    this.filters.add(filtersItem);
    return this;
  }

  /**
   * Applied facet filters. Facet filters are `facet:value` pairs. Facet values must be URL-encoded,
   * such as, `discount:10%25`.
   */
  @javax.annotation.Nonnull
  public List<String> getFilters() {
    return filters;
  }

  public ClickedFilters setUserToken(String userToken) {
    this.userToken = userToken;
    return this;
  }

  /**
   * Anonymous or pseudonymous user identifier. Don't use personally identifiable information in
   * user tokens. For more information, see [User
   * token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
   */
  @javax.annotation.Nonnull
  public String getUserToken() {
    return userToken;
  }

  public ClickedFilters setAuthenticatedUserToken(String authenticatedUserToken) {
    this.authenticatedUserToken = authenticatedUserToken;
    return this;
  }

  /**
   * Identifier for authenticated users. When the user signs in, you can get an identifier from your
   * system and send it as `authenticatedUserToken`. This lets you keep using the `userToken` from
   * before the user signed in, while providing a reliable way to identify users across sessions.
   * Don't use personally identifiable information in user tokens. For more information, see [User
   * token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
   */
  @javax.annotation.Nullable
  public String getAuthenticatedUserToken() {
    return authenticatedUserToken;
  }

  public ClickedFilters setTimestamp(Long timestamp) {
    this.timestamp = timestamp;
    return this;
  }

  /**
   * Timestamp of the event, measured in milliseconds since the Unix epoch. By default, the Insights
   * API uses the time it receives an event as its timestamp.
   */
  @javax.annotation.Nullable
  public Long getTimestamp() {
    return timestamp;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClickedFilters clickedFilters = (ClickedFilters) o;
    return (
      Objects.equals(this.eventName, clickedFilters.eventName) &&
      Objects.equals(this.eventType, clickedFilters.eventType) &&
      Objects.equals(this.index, clickedFilters.index) &&
      Objects.equals(this.filters, clickedFilters.filters) &&
      Objects.equals(this.userToken, clickedFilters.userToken) &&
      Objects.equals(this.authenticatedUserToken, clickedFilters.authenticatedUserToken) &&
      Objects.equals(this.timestamp, clickedFilters.timestamp)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(eventName, eventType, index, filters, userToken, authenticatedUserToken, timestamp);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClickedFilters {\n");
    sb.append("    eventName: ").append(toIndentedString(eventName)).append("\n");
    sb.append("    eventType: ").append(toIndentedString(eventType)).append("\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    filters: ").append(toIndentedString(filters)).append("\n");
    sb.append("    userToken: ").append(toIndentedString(userToken)).append("\n");
    sb.append("    authenticatedUserToken: ").append(toIndentedString(authenticatedUserToken)).append("\n");
    sb.append("    timestamp: ").append(toIndentedString(timestamp)).append("\n");
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
