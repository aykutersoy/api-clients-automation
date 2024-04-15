# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from re import match
from typing import Annotated, Any, Dict, List, Optional, Self

from pydantic import BaseModel, Field, StrictInt, StrictStr, field_validator

from algoliasearch.insights.models.click_event import ClickEvent


class ClickedFilters(BaseModel):
    """
    Use this event to track when users click facet filters in your user interface.
    """

    event_name: Annotated[str, Field(min_length=1, strict=True, max_length=64)] = Field(
        description="Event name, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework. ",
        alias="eventName",
    )
    event_type: ClickEvent = Field(alias="eventType")
    index: StrictStr = Field(
        description="Index name (case-sensitive) to which the event's items belong."
    )
    filters: Annotated[List[StrictStr], Field(min_length=1, max_length=20)] = Field(
        description="Applied facet filters.  Facet filters are `facet:value` pairs. Facet values must be URL-encoded, such as, `discount:10%25`. "
    )
    user_token: Annotated[str, Field(min_length=1, strict=True, max_length=129)] = (
        Field(
            description="Anonymous or pseudonymous user identifier.  Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/). ",
            alias="userToken",
        )
    )
    authenticated_user_token: Optional[
        Annotated[str, Field(min_length=1, strict=True, max_length=129)]
    ] = Field(
        default=None,
        description="Identifier for authenticated users.  When the user signs in, you can get an identifier from your system and send it as `authenticatedUserToken`. This lets you keep using the `userToken` from before the user signed in, while providing a reliable way to identify users across sessions. Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/). ",
        alias="authenticatedUserToken",
    )
    timestamp: Optional[StrictInt] = Field(
        default=None,
        description="Timestamp of the event, measured in milliseconds since the Unix epoch. By default, the Insights API uses the time it receives an event as its timestamp. ",
    )

    @field_validator("event_name")
    def event_name_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if not match(r"[\x20-\x7E]{1,64}", value):
            raise ValueError(
                r"must validate the regular expression /[\x20-\x7E]{1,64}/"
            )
        return value

    @field_validator("user_token")
    def user_token_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if not match(r"[a-zA-Z0-9_=\/+-]{1,129}", value):
            raise ValueError(
                r"must validate the regular expression /[a-zA-Z0-9_=\/+-]{1,129}/"
            )
        return value

    @field_validator("authenticated_user_token")
    def authenticated_user_token_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not match(r"[a-zA-Z0-9_=\/+-]{1,129}", value):
            raise ValueError(
                r"must validate the regular expression /[a-zA-Z0-9_=\/+-]{1,129}/"
            )
        return value

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of ClickedFilters from a JSON string"""
        return cls.from_dict(loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={},
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ClickedFilters from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "eventName": obj.get("eventName"),
                "eventType": obj.get("eventType"),
                "index": obj.get("index"),
                "filters": obj.get("filters"),
                "userToken": obj.get("userToken"),
                "authenticatedUserToken": obj.get("authenticatedUserToken"),
                "timestamp": obj.get("timestamp"),
            }
        )
        return _obj
