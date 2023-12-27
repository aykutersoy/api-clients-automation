# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Any, ClassVar, Dict, List, Optional, Self

from pydantic import BaseModel, Field

from algoliasearch.recommend.models.facet_ordering import FacetOrdering


class RenderingContent(BaseModel):
    """
    Extra content for the search UI, for example, to control the [ordering and display of facets](https://www.algolia.com/doc/guides/managing-results/rules/merchandising-and-promoting/how-to/merchandising-facets/#merchandise-facets-and-their-values-in-the-manual-editor). You can set a default value and dynamically override it with [Rules](https://www.algolia.com/doc/guides/managing-results/rules/rules-overview/).
    """

    facet_ordering: Optional[FacetOrdering] = Field(default=None, alias="facetOrdering")
    __properties: ClassVar[List[str]] = ["facetOrdering"]

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of RenderingContent from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of
        # facet_ordering
        if self.facet_ordering:
            _dict["facetOrdering"] = self.facet_ordering.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of RenderingContent from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "facetOrdering": FacetOrdering.from_dict(obj.get("facetOrdering"))
                if obj.get("facetOrdering") is not None
                else None
            }
        )
        return _obj
