# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Any, ClassVar, Dict, List, Optional, Self

from pydantic import BaseModel, Field, StrictInt, StrictStr

from algoliasearch.search.models.highlight_result import HighlightResult
from algoliasearch.search.models.ranking_info import RankingInfo
from algoliasearch.search.models.snippet_result import SnippetResult


class Hit(BaseModel):
    """
    A single hit.
    """

    object_id: StrictStr = Field(
        description="Unique object identifier.", alias="objectID"
    )
    highlight_result: Optional[Dict[str, HighlightResult]] = Field(
        default=None,
        description="Show highlighted section and words matched on a query.",
        alias="_highlightResult",
    )
    snippet_result: Optional[Dict[str, SnippetResult]] = Field(
        default=None,
        description="Snippeted attributes show parts of the matched attributes. Only returned when attributesToSnippet is non-empty.",
        alias="_snippetResult",
    )
    ranking_info: Optional[RankingInfo] = Field(default=None, alias="_rankingInfo")
    distinct_seq_id: Optional[StrictInt] = Field(default=None, alias="_distinctSeqID")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = [
        "objectID",
        "_highlightResult",
        "_snippetResult",
        "_rankingInfo",
        "_distinctSeqID",
    ]

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of Hit from a JSON string"""
        return cls.from_dict(loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        * Fields in `self.additional_properties` are added to the output dict.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={
                "additional_properties",
            },
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of
        # each value in highlight_result (dict)
        _field_dict = {}
        if self.highlight_result:
            for _key in self.highlight_result:
                if self.highlight_result[_key]:
                    _field_dict[_key] = self.highlight_result[_key].to_dict()
            _dict["_highlightResult"] = _field_dict
        # override the default output from pydantic by calling `to_dict()` of
        # each value in snippet_result (dict)
        _field_dict = {}
        if self.snippet_result:
            for _key in self.snippet_result:
                if self.snippet_result[_key]:
                    _field_dict[_key] = self.snippet_result[_key].to_dict()
            _dict["_snippetResult"] = _field_dict
        # override the default output from pydantic by calling `to_dict()` of
        # ranking_info
        if self.ranking_info:
            _dict["_rankingInfo"] = self.ranking_info.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of Hit from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "objectID": obj.get("objectID"),
                "_highlightResult": dict(
                    (_k, HighlightResult.from_dict(_v))
                    for _k, _v in obj.get("_highlightResult").items()
                )
                if obj.get("_highlightResult") is not None
                else None,
                "_snippetResult": dict(
                    (_k, SnippetResult.from_dict(_v))
                    for _k, _v in obj.get("_snippetResult").items()
                )
                if obj.get("_snippetResult") is not None
                else None,
                "_rankingInfo": RankingInfo.from_dict(obj.get("_rankingInfo"))
                if obj.get("_rankingInfo") is not None
                else None,
                "_distinctSeqID": obj.get("_distinctSeqID"),
            }
        )
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj
