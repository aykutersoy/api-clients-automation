# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Annotated, Any, ClassVar, Dict, List, Optional, Self, Union

from pydantic import BaseModel, Field, StrictBool, StrictFloat, StrictInt, StrictStr

from algoliasearch.search.models.advanced_syntax_features import AdvancedSyntaxFeatures
from algoliasearch.search.models.alternatives_as_exact import AlternativesAsExact
from algoliasearch.search.models.around_precision import AroundPrecision
from algoliasearch.search.models.around_radius import AroundRadius
from algoliasearch.search.models.distinct import Distinct
from algoliasearch.search.models.exact_on_single_word_query import (
    ExactOnSingleWordQuery,
)
from algoliasearch.search.models.facet_filters import FacetFilters
from algoliasearch.search.models.ignore_plurals import IgnorePlurals
from algoliasearch.search.models.mode import Mode
from algoliasearch.search.models.numeric_filters import NumericFilters
from algoliasearch.search.models.optional_filters import OptionalFilters
from algoliasearch.search.models.query_type import QueryType
from algoliasearch.search.models.re_ranking_apply_filter import ReRankingApplyFilter
from algoliasearch.search.models.remove_stop_words import RemoveStopWords
from algoliasearch.search.models.remove_words_if_no_results import (
    RemoveWordsIfNoResults,
)
from algoliasearch.search.models.rendering_content import RenderingContent
from algoliasearch.search.models.search_type_facet import SearchTypeFacet
from algoliasearch.search.models.semantic_search import SemanticSearch
from algoliasearch.search.models.tag_filters import TagFilters
from algoliasearch.search.models.typo_tolerance import TypoTolerance


class SearchForFacets(BaseModel):
    """
    SearchForFacets
    """

    params: Optional[StrictStr] = Field(
        default="", description="Search parameters as a URL-encoded query string."
    )
    query: Optional[StrictStr] = Field(
        default="", description="Text to search for in an index."
    )
    similar_query: Optional[StrictStr] = Field(
        default="",
        description="Overrides the query parameter and performs a more generic search.",
        alias="similarQuery",
    )
    filters: Optional[StrictStr] = Field(
        default="",
        description="[Filter](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/) the query with numeric, facet, or tag filters. ",
    )
    facet_filters: Optional[FacetFilters] = Field(default=None, alias="facetFilters")
    optional_filters: Optional[OptionalFilters] = Field(
        default=None, alias="optionalFilters"
    )
    numeric_filters: Optional[NumericFilters] = Field(
        default=None, alias="numericFilters"
    )
    tag_filters: Optional[TagFilters] = Field(default=None, alias="tagFilters")
    sum_or_filters_scores: Optional[StrictBool] = Field(
        default=False,
        description="Determines how to calculate [filter scores](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/in-depth/filter-scoring/#accumulating-scores-with-sumorfiltersscores). If `false`, maximum score is kept. If `true`, score is summed. ",
        alias="sumOrFiltersScores",
    )
    restrict_searchable_attributes: Optional[List[StrictStr]] = Field(
        default=None,
        description="Restricts a query to only look at a subset of your [searchable attributes](https://www.algolia.com/doc/guides/managing-results/must-do/searchable-attributes/).",
        alias="restrictSearchableAttributes",
    )
    facets: Optional[List[StrictStr]] = Field(
        default=None,
        description="Returns [facets](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/#contextual-facet-values-and-counts), their facet values, and the number of matching facet values.",
    )
    faceting_after_distinct: Optional[StrictBool] = Field(
        default=False,
        description="Forces faceting to be applied after [de-duplication](https://www.algolia.com/doc/guides/managing-results/refine-results/grouping/) (with the distinct feature). Alternatively, the `afterDistinct` [modifier](https://www.algolia.com/doc/api-reference/api-parameters/attributesForFaceting/#modifiers) of `attributesForFaceting` allows for more granular control. ",
        alias="facetingAfterDistinct",
    )
    page: Optional[StrictInt] = Field(
        default=0, description="Page to retrieve (the first page is `0`, not `1`)."
    )
    offset: Optional[StrictInt] = Field(
        default=None,
        description="Specifies the offset of the first hit to return. > **Note**: Using `page` and `hitsPerPage` is the recommended method for [paging results](https://www.algolia.com/doc/guides/building-search-ui/ui-and-ux-patterns/pagination/js/). However, you can use `offset` and `length` to implement [an alternative approach to paging](https://www.algolia.com/doc/guides/building-search-ui/ui-and-ux-patterns/pagination/js/#retrieving-a-subset-of-records-with-offset-and-length). ",
    )
    length: Optional[Annotated[int, Field(le=1000, strict=True, ge=1)]] = Field(
        default=None,
        description="Sets the number of hits to retrieve (for use with `offset`). > **Note**: Using `page` and `hitsPerPage` is the recommended method for [paging results](https://www.algolia.com/doc/guides/building-search-ui/ui-and-ux-patterns/pagination/js/). However, you can use `offset` and `length` to implement [an alternative approach to paging](https://www.algolia.com/doc/guides/building-search-ui/ui-and-ux-patterns/pagination/js/#retrieving-a-subset-of-records-with-offset-and-length). ",
    )
    around_lat_lng: Optional[StrictStr] = Field(
        default="",
        description="Search for entries [around a central location](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filter-around-a-central-point), enabling a geographical search within a circular area.",
        alias="aroundLatLng",
    )
    around_lat_lng_via_ip: Optional[StrictBool] = Field(
        default=False,
        description="Search for entries around a location. The location is automatically computed from the requester's IP address.",
        alias="aroundLatLngViaIP",
    )
    around_radius: Optional[AroundRadius] = Field(default=None, alias="aroundRadius")
    around_precision: Optional[AroundPrecision] = Field(
        default=None, alias="aroundPrecision"
    )
    minimum_around_radius: Optional[Annotated[int, Field(strict=True, ge=1)]] = Field(
        default=None,
        description="Minimum radius (in meters) used for a geographical search when `aroundRadius` isn't set.",
        alias="minimumAroundRadius",
    )
    inside_bounding_box: Optional[List[List[Union[StrictFloat, StrictInt]]]] = Field(
        default=None,
        description="Search inside a [rectangular area](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas) (in geographical coordinates).",
        alias="insideBoundingBox",
    )
    inside_polygon: Optional[List[List[Union[StrictFloat, StrictInt]]]] = Field(
        default=None,
        description="Search inside a [polygon](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas) (in geographical coordinates).",
        alias="insidePolygon",
    )
    natural_languages: Optional[List[StrictStr]] = Field(
        default=None,
        description="Changes the default values of parameters that work best for a natural language query, such as `ignorePlurals`, `removeStopWords`, `removeWordsIfNoResults`, `analyticsTags`, and `ruleContexts`. These parameters work well together when the query consists of fuller natural language strings instead of keywords, for example when processing voice search queries.",
        alias="naturalLanguages",
    )
    rule_contexts: Optional[List[StrictStr]] = Field(
        default=None,
        description="Assigns [rule contexts](https://www.algolia.com/doc/guides/managing-results/rules/rules-overview/how-to/customize-search-results-by-platform/#whats-a-context) to search queries.",
        alias="ruleContexts",
    )
    personalization_impact: Optional[StrictInt] = Field(
        default=100,
        description="Defines how much [Personalization affects results](https://www.algolia.com/doc/guides/personalization/personalizing-results/in-depth/configuring-personalization/#understanding-personalization-impact).",
        alias="personalizationImpact",
    )
    user_token: Optional[StrictStr] = Field(
        default=None,
        description="Associates a [user token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/) with the current search.",
        alias="userToken",
    )
    get_ranking_info: Optional[StrictBool] = Field(
        default=False,
        description="Incidates whether the search response includes [detailed ranking information](https://www.algolia.com/doc/guides/building-search-ui/going-further/backend-search/in-depth/understanding-the-api-response/#ranking-information).",
        alias="getRankingInfo",
    )
    explain: Optional[List[StrictStr]] = Field(
        default=None,
        description="Enriches the API's response with information about how the query was processed.",
    )
    synonyms: Optional[StrictBool] = Field(
        default=True,
        description="Whether to take into account an index's synonyms for a particular search.",
    )
    click_analytics: Optional[StrictBool] = Field(
        default=False,
        description="Indicates whether a query ID parameter is included in the search response. This is required for [tracking click and conversion events](https://www.algolia.com/doc/guides/sending-events/concepts/event-types/#events-related-to-algolia-requests).",
        alias="clickAnalytics",
    )
    analytics: Optional[StrictBool] = Field(
        default=True,
        description="Indicates whether this query will be included in [analytics](https://www.algolia.com/doc/guides/search-analytics/guides/exclude-queries/).",
    )
    analytics_tags: Optional[List[StrictStr]] = Field(
        default=None,
        description="Tags to apply to the query for [segmenting analytics data](https://www.algolia.com/doc/guides/search-analytics/guides/segments/).",
        alias="analyticsTags",
    )
    percentile_computation: Optional[StrictBool] = Field(
        default=True,
        description="Whether to include or exclude a query from the processing-time percentile computation.",
        alias="percentileComputation",
    )
    enable_ab_test: Optional[StrictBool] = Field(
        default=True,
        description="Incidates whether this search will be considered in A/B testing.",
        alias="enableABTest",
    )
    attributes_for_faceting: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes used for [faceting](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/) and the [modifiers](https://www.algolia.com/doc/api-reference/api-parameters/attributesForFaceting/#modifiers) that can be applied: `filterOnly`, `searchable`, and `afterDistinct`. ",
        alias="attributesForFaceting",
    )
    attributes_to_retrieve: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes to include in the API response. To reduce the size of your response, you can retrieve only some of the attributes. By default, the response includes all attributes.",
        alias="attributesToRetrieve",
    )
    ranking: Optional[List[StrictStr]] = Field(
        default=None,
        description="Determines the order in which Algolia [returns your results](https://www.algolia.com/doc/guides/managing-results/relevance-overview/in-depth/ranking-criteria/).",
    )
    custom_ranking: Optional[List[StrictStr]] = Field(
        default=None,
        description="Specifies the [Custom ranking criterion](https://www.algolia.com/doc/guides/managing-results/must-do/custom-ranking/). Use the `asc` and `desc` modifiers to specify the ranking order: ascending or descending. ",
        alias="customRanking",
    )
    relevancy_strictness: Optional[StrictInt] = Field(
        default=100,
        description="Relevancy threshold below which less relevant results aren't included in the results.",
        alias="relevancyStrictness",
    )
    attributes_to_highlight: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes to highlight. Strings that match the search query in the attributes are highlighted by surrounding them with HTML tags (`highlightPreTag` and `highlightPostTag`).",
        alias="attributesToHighlight",
    )
    attributes_to_snippet: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes to _snippet_. 'Snippeting' is shortening the attribute to a certain number of words. If not specified, the attribute is shortened to the 10 words around the matching string but you can specify the number. For example: `body:20`. ",
        alias="attributesToSnippet",
    )
    highlight_pre_tag: Optional[StrictStr] = Field(
        default="<em>",
        description="HTML string to insert before the highlighted parts in all highlight and snippet results.",
        alias="highlightPreTag",
    )
    highlight_post_tag: Optional[StrictStr] = Field(
        default="</em>",
        description="HTML string to insert after the highlighted parts in all highlight and snippet results.",
        alias="highlightPostTag",
    )
    snippet_ellipsis_text: Optional[StrictStr] = Field(
        default="…",
        description="String used as an ellipsis indicator when a snippet is truncated.",
        alias="snippetEllipsisText",
    )
    restrict_highlight_and_snippet_arrays: Optional[StrictBool] = Field(
        default=False,
        description="Restrict highlighting and snippeting to items that matched the query.",
        alias="restrictHighlightAndSnippetArrays",
    )
    hits_per_page: Optional[Annotated[int, Field(le=1000, strict=True, ge=1)]] = Field(
        default=20, description="Number of hits per page.", alias="hitsPerPage"
    )
    min_word_sizefor1_typo: Optional[StrictInt] = Field(
        default=4,
        description="Minimum number of characters a word in the query string must contain to accept matches with [one typo](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/in-depth/configuring-typo-tolerance/#configuring-word-length-for-typos).",
        alias="minWordSizefor1Typo",
    )
    min_word_sizefor2_typos: Optional[StrictInt] = Field(
        default=8,
        description="Minimum number of characters a word in the query string must contain to accept matches with [two typos](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/in-depth/configuring-typo-tolerance/#configuring-word-length-for-typos).",
        alias="minWordSizefor2Typos",
    )
    typo_tolerance: Optional[TypoTolerance] = Field(default=None, alias="typoTolerance")
    allow_typos_on_numeric_tokens: Optional[StrictBool] = Field(
        default=True,
        description='Whether to allow typos on numbers ("numeric tokens") in the query string.',
        alias="allowTyposOnNumericTokens",
    )
    disable_typo_tolerance_on_attributes: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes for which you want to turn off [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/).",
        alias="disableTypoToleranceOnAttributes",
    )
    ignore_plurals: Optional[IgnorePlurals] = Field(default=None, alias="ignorePlurals")
    remove_stop_words: Optional[RemoveStopWords] = Field(
        default=None, alias="removeStopWords"
    )
    keep_diacritics_on_characters: Optional[StrictStr] = Field(
        default="",
        description="Characters that the engine shouldn't automatically [normalize](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).",
        alias="keepDiacriticsOnCharacters",
    )
    query_languages: Optional[List[StrictStr]] = Field(
        default=None,
        description="Sets your user's search language. This adjusts language-specific settings and features such as `ignorePlurals`, `removeStopWords`, and [CJK](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/#normalization-for-logogram-based-languages-cjk) word detection.",
        alias="queryLanguages",
    )
    decompound_query: Optional[StrictBool] = Field(
        default=True,
        description="[Splits compound words](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/#splitting-compound-words) into their component word parts in the query. ",
        alias="decompoundQuery",
    )
    enable_rules: Optional[StrictBool] = Field(
        default=True,
        description="Incidates whether [Rules](https://www.algolia.com/doc/guides/managing-results/rules/rules-overview/) are enabled.",
        alias="enableRules",
    )
    enable_personalization: Optional[StrictBool] = Field(
        default=False,
        description="Incidates whether [Personalization](https://www.algolia.com/doc/guides/personalization/what-is-personalization/) is enabled.",
        alias="enablePersonalization",
    )
    query_type: Optional[QueryType] = Field(default=None, alias="queryType")
    remove_words_if_no_results: Optional[RemoveWordsIfNoResults] = Field(
        default=None, alias="removeWordsIfNoResults"
    )
    mode: Optional[Mode] = None
    semantic_search: Optional[SemanticSearch] = Field(
        default=None, alias="semanticSearch"
    )
    advanced_syntax: Optional[StrictBool] = Field(
        default=False,
        description="Enables the [advanced query syntax](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/#advanced-syntax).",
        alias="advancedSyntax",
    )
    optional_words: Optional[List[StrictStr]] = Field(
        default=None,
        description="Words which should be considered [optional](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/empty-or-insufficient-results/#creating-a-list-of-optional-words) when found in a query.",
        alias="optionalWords",
    )
    disable_exact_on_attributes: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes for which you want to [turn off the exact ranking criterion](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/in-depth/adjust-exact-settings/#turn-off-exact-for-some-attributes).",
        alias="disableExactOnAttributes",
    )
    exact_on_single_word_query: Optional[ExactOnSingleWordQuery] = Field(
        default=None, alias="exactOnSingleWordQuery"
    )
    alternatives_as_exact: Optional[List[AlternativesAsExact]] = Field(
        default=None,
        description="Alternatives that should be considered an exact match by [the exact ranking criterion](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/in-depth/adjust-exact-settings/#turn-off-exact-for-some-attributes).",
        alias="alternativesAsExact",
    )
    advanced_syntax_features: Optional[List[AdvancedSyntaxFeatures]] = Field(
        default=None,
        description="Allows you to specify which advanced syntax features are active when `advancedSyntax` is enabled.",
        alias="advancedSyntaxFeatures",
    )
    distinct: Optional[Distinct] = None
    replace_synonyms_in_highlight: Optional[StrictBool] = Field(
        default=False,
        description="Whether to highlight and snippet the original word that matches the synonym or the synonym itself.",
        alias="replaceSynonymsInHighlight",
    )
    min_proximity: Optional[Annotated[int, Field(le=7, strict=True, ge=1)]] = Field(
        default=1,
        description="Precision of the [proximity ranking criterion](https://www.algolia.com/doc/guides/managing-results/relevance-overview/in-depth/ranking-criteria/#proximity).",
        alias="minProximity",
    )
    response_fields: Optional[List[StrictStr]] = Field(
        default=None,
        description="Attributes to include in the API response for search and browse queries.",
        alias="responseFields",
    )
    max_facet_hits: Optional[Annotated[int, Field(le=100, strict=True)]] = Field(
        default=10,
        description="Maximum number of facet hits to return when [searching for facet values](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/#search-for-facet-values).",
        alias="maxFacetHits",
    )
    max_values_per_facet: Optional[StrictInt] = Field(
        default=100,
        description="Maximum number of facet values to return for each facet.",
        alias="maxValuesPerFacet",
    )
    sort_facet_values_by: Optional[StrictStr] = Field(
        default="count",
        description="Controls how facet values are fetched.",
        alias="sortFacetValuesBy",
    )
    attribute_criteria_computed_by_min_proximity: Optional[StrictBool] = Field(
        default=False,
        description="When the [Attribute criterion is ranked above Proximity](https://www.algolia.com/doc/guides/managing-results/relevance-overview/in-depth/ranking-criteria/#attribute-and-proximity-combinations) in your ranking formula, Proximity is used to select which searchable attribute is matched in the Attribute ranking stage.",
        alias="attributeCriteriaComputedByMinProximity",
    )
    rendering_content: Optional[RenderingContent] = Field(
        default=None, alias="renderingContent"
    )
    enable_re_ranking: Optional[StrictBool] = Field(
        default=True,
        description="Indicates whether this search will use [Dynamic Re-Ranking](https://www.algolia.com/doc/guides/algolia-ai/re-ranking/).",
        alias="enableReRanking",
    )
    re_ranking_apply_filter: Optional[ReRankingApplyFilter] = Field(
        default=None, alias="reRankingApplyFilter"
    )
    facet: StrictStr = Field(description="Facet name.")
    index_name: StrictStr = Field(description="Algolia index name.", alias="indexName")
    facet_query: Optional[StrictStr] = Field(
        default="",
        description="Text to search inside the facet's values.",
        alias="facetQuery",
    )
    type: SearchTypeFacet
    __properties: ClassVar[List[str]] = [
        "params",
        "query",
        "similarQuery",
        "filters",
        "facetFilters",
        "optionalFilters",
        "numericFilters",
        "tagFilters",
        "sumOrFiltersScores",
        "restrictSearchableAttributes",
        "facets",
        "facetingAfterDistinct",
        "page",
        "offset",
        "length",
        "aroundLatLng",
        "aroundLatLngViaIP",
        "aroundRadius",
        "aroundPrecision",
        "minimumAroundRadius",
        "insideBoundingBox",
        "insidePolygon",
        "naturalLanguages",
        "ruleContexts",
        "personalizationImpact",
        "userToken",
        "getRankingInfo",
        "explain",
        "synonyms",
        "clickAnalytics",
        "analytics",
        "analyticsTags",
        "percentileComputation",
        "enableABTest",
        "attributesForFaceting",
        "attributesToRetrieve",
        "ranking",
        "customRanking",
        "relevancyStrictness",
        "attributesToHighlight",
        "attributesToSnippet",
        "highlightPreTag",
        "highlightPostTag",
        "snippetEllipsisText",
        "restrictHighlightAndSnippetArrays",
        "hitsPerPage",
        "minWordSizefor1Typo",
        "minWordSizefor2Typos",
        "typoTolerance",
        "allowTyposOnNumericTokens",
        "disableTypoToleranceOnAttributes",
        "ignorePlurals",
        "removeStopWords",
        "keepDiacriticsOnCharacters",
        "queryLanguages",
        "decompoundQuery",
        "enableRules",
        "enablePersonalization",
        "queryType",
        "removeWordsIfNoResults",
        "mode",
        "semanticSearch",
        "advancedSyntax",
        "optionalWords",
        "disableExactOnAttributes",
        "exactOnSingleWordQuery",
        "alternativesAsExact",
        "advancedSyntaxFeatures",
        "distinct",
        "replaceSynonymsInHighlight",
        "minProximity",
        "responseFields",
        "maxFacetHits",
        "maxValuesPerFacet",
        "sortFacetValuesBy",
        "attributeCriteriaComputedByMinProximity",
        "renderingContent",
        "enableReRanking",
        "reRankingApplyFilter",
        "facet",
        "indexName",
        "facetQuery",
        "type",
    ]

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of SearchForFacets from a JSON string"""
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
        # facet_filters
        if self.facet_filters:
            _dict["facetFilters"] = self.facet_filters.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # optional_filters
        if self.optional_filters:
            _dict["optionalFilters"] = self.optional_filters.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # numeric_filters
        if self.numeric_filters:
            _dict["numericFilters"] = self.numeric_filters.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # tag_filters
        if self.tag_filters:
            _dict["tagFilters"] = self.tag_filters.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # around_radius
        if self.around_radius:
            _dict["aroundRadius"] = self.around_radius.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # around_precision
        if self.around_precision:
            _dict["aroundPrecision"] = self.around_precision.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # typo_tolerance
        if self.typo_tolerance:
            _dict["typoTolerance"] = self.typo_tolerance.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # ignore_plurals
        if self.ignore_plurals:
            _dict["ignorePlurals"] = self.ignore_plurals.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # remove_stop_words
        if self.remove_stop_words:
            _dict["removeStopWords"] = self.remove_stop_words.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # semantic_search
        if self.semantic_search:
            _dict["semanticSearch"] = self.semantic_search.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # distinct
        if self.distinct:
            _dict["distinct"] = self.distinct.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # rendering_content
        if self.rendering_content:
            _dict["renderingContent"] = self.rendering_content.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # re_ranking_apply_filter
        if self.re_ranking_apply_filter:
            _dict["reRankingApplyFilter"] = self.re_ranking_apply_filter.to_dict()
        # set to None if re_ranking_apply_filter (nullable) is None
        # and model_fields_set contains the field
        if (
            self.re_ranking_apply_filter is None
            and "re_ranking_apply_filter" in self.model_fields_set
        ):
            _dict["reRankingApplyFilter"] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of SearchForFacets from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "params": obj.get("params") if obj.get("params") is not None else "",
                "query": obj.get("query") if obj.get("query") is not None else "",
                "similarQuery": obj.get("similarQuery")
                if obj.get("similarQuery") is not None
                else "",
                "filters": obj.get("filters") if obj.get("filters") is not None else "",
                "facetFilters": FacetFilters.from_dict(obj.get("facetFilters"))
                if obj.get("facetFilters") is not None
                else None,
                "optionalFilters": OptionalFilters.from_dict(obj.get("optionalFilters"))
                if obj.get("optionalFilters") is not None
                else None,
                "numericFilters": NumericFilters.from_dict(obj.get("numericFilters"))
                if obj.get("numericFilters") is not None
                else None,
                "tagFilters": TagFilters.from_dict(obj.get("tagFilters"))
                if obj.get("tagFilters") is not None
                else None,
                "sumOrFiltersScores": obj.get("sumOrFiltersScores")
                if obj.get("sumOrFiltersScores") is not None
                else False,
                "restrictSearchableAttributes": obj.get("restrictSearchableAttributes"),
                "facets": obj.get("facets"),
                "facetingAfterDistinct": obj.get("facetingAfterDistinct")
                if obj.get("facetingAfterDistinct") is not None
                else False,
                "page": obj.get("page") if obj.get("page") is not None else 0,
                "offset": obj.get("offset"),
                "length": obj.get("length"),
                "aroundLatLng": obj.get("aroundLatLng")
                if obj.get("aroundLatLng") is not None
                else "",
                "aroundLatLngViaIP": obj.get("aroundLatLngViaIP")
                if obj.get("aroundLatLngViaIP") is not None
                else False,
                "aroundRadius": AroundRadius.from_dict(obj.get("aroundRadius"))
                if obj.get("aroundRadius") is not None
                else None,
                "aroundPrecision": AroundPrecision.from_dict(obj.get("aroundPrecision"))
                if obj.get("aroundPrecision") is not None
                else None,
                "minimumAroundRadius": obj.get("minimumAroundRadius"),
                "insideBoundingBox": obj.get("insideBoundingBox"),
                "insidePolygon": obj.get("insidePolygon"),
                "naturalLanguages": obj.get("naturalLanguages"),
                "ruleContexts": obj.get("ruleContexts"),
                "personalizationImpact": obj.get("personalizationImpact")
                if obj.get("personalizationImpact") is not None
                else 100,
                "userToken": obj.get("userToken"),
                "getRankingInfo": obj.get("getRankingInfo")
                if obj.get("getRankingInfo") is not None
                else False,
                "explain": obj.get("explain"),
                "synonyms": obj.get("synonyms")
                if obj.get("synonyms") is not None
                else True,
                "clickAnalytics": obj.get("clickAnalytics")
                if obj.get("clickAnalytics") is not None
                else False,
                "analytics": obj.get("analytics")
                if obj.get("analytics") is not None
                else True,
                "analyticsTags": obj.get("analyticsTags"),
                "percentileComputation": obj.get("percentileComputation")
                if obj.get("percentileComputation") is not None
                else True,
                "enableABTest": obj.get("enableABTest")
                if obj.get("enableABTest") is not None
                else True,
                "attributesForFaceting": obj.get("attributesForFaceting"),
                "attributesToRetrieve": obj.get("attributesToRetrieve"),
                "ranking": obj.get("ranking"),
                "customRanking": obj.get("customRanking"),
                "relevancyStrictness": obj.get("relevancyStrictness")
                if obj.get("relevancyStrictness") is not None
                else 100,
                "attributesToHighlight": obj.get("attributesToHighlight"),
                "attributesToSnippet": obj.get("attributesToSnippet"),
                "highlightPreTag": obj.get("highlightPreTag")
                if obj.get("highlightPreTag") is not None
                else "<em>",
                "highlightPostTag": obj.get("highlightPostTag")
                if obj.get("highlightPostTag") is not None
                else "</em>",
                "snippetEllipsisText": obj.get("snippetEllipsisText")
                if obj.get("snippetEllipsisText") is not None
                else "…",
                "restrictHighlightAndSnippetArrays": obj.get(
                    "restrictHighlightAndSnippetArrays"
                )
                if obj.get("restrictHighlightAndSnippetArrays") is not None
                else False,
                "hitsPerPage": obj.get("hitsPerPage")
                if obj.get("hitsPerPage") is not None
                else 20,
                "minWordSizefor1Typo": obj.get("minWordSizefor1Typo")
                if obj.get("minWordSizefor1Typo") is not None
                else 4,
                "minWordSizefor2Typos": obj.get("minWordSizefor2Typos")
                if obj.get("minWordSizefor2Typos") is not None
                else 8,
                "typoTolerance": TypoTolerance.from_dict(obj.get("typoTolerance"))
                if obj.get("typoTolerance") is not None
                else None,
                "allowTyposOnNumericTokens": obj.get("allowTyposOnNumericTokens")
                if obj.get("allowTyposOnNumericTokens") is not None
                else True,
                "disableTypoToleranceOnAttributes": obj.get(
                    "disableTypoToleranceOnAttributes"
                ),
                "ignorePlurals": IgnorePlurals.from_dict(obj.get("ignorePlurals"))
                if obj.get("ignorePlurals") is not None
                else None,
                "removeStopWords": RemoveStopWords.from_dict(obj.get("removeStopWords"))
                if obj.get("removeStopWords") is not None
                else None,
                "keepDiacriticsOnCharacters": obj.get("keepDiacriticsOnCharacters")
                if obj.get("keepDiacriticsOnCharacters") is not None
                else "",
                "queryLanguages": obj.get("queryLanguages"),
                "decompoundQuery": obj.get("decompoundQuery")
                if obj.get("decompoundQuery") is not None
                else True,
                "enableRules": obj.get("enableRules")
                if obj.get("enableRules") is not None
                else True,
                "enablePersonalization": obj.get("enablePersonalization")
                if obj.get("enablePersonalization") is not None
                else False,
                "queryType": obj.get("queryType"),
                "removeWordsIfNoResults": obj.get("removeWordsIfNoResults"),
                "mode": obj.get("mode"),
                "semanticSearch": SemanticSearch.from_dict(obj.get("semanticSearch"))
                if obj.get("semanticSearch") is not None
                else None,
                "advancedSyntax": obj.get("advancedSyntax")
                if obj.get("advancedSyntax") is not None
                else False,
                "optionalWords": obj.get("optionalWords"),
                "disableExactOnAttributes": obj.get("disableExactOnAttributes"),
                "exactOnSingleWordQuery": obj.get("exactOnSingleWordQuery"),
                "alternativesAsExact": obj.get("alternativesAsExact"),
                "advancedSyntaxFeatures": obj.get("advancedSyntaxFeatures"),
                "distinct": Distinct.from_dict(obj.get("distinct"))
                if obj.get("distinct") is not None
                else None,
                "replaceSynonymsInHighlight": obj.get("replaceSynonymsInHighlight")
                if obj.get("replaceSynonymsInHighlight") is not None
                else False,
                "minProximity": obj.get("minProximity")
                if obj.get("minProximity") is not None
                else 1,
                "responseFields": obj.get("responseFields"),
                "maxFacetHits": obj.get("maxFacetHits")
                if obj.get("maxFacetHits") is not None
                else 10,
                "maxValuesPerFacet": obj.get("maxValuesPerFacet")
                if obj.get("maxValuesPerFacet") is not None
                else 100,
                "sortFacetValuesBy": obj.get("sortFacetValuesBy")
                if obj.get("sortFacetValuesBy") is not None
                else "count",
                "attributeCriteriaComputedByMinProximity": obj.get(
                    "attributeCriteriaComputedByMinProximity"
                )
                if obj.get("attributeCriteriaComputedByMinProximity") is not None
                else False,
                "renderingContent": RenderingContent.from_dict(
                    obj.get("renderingContent")
                )
                if obj.get("renderingContent") is not None
                else None,
                "enableReRanking": obj.get("enableReRanking")
                if obj.get("enableReRanking") is not None
                else True,
                "reRankingApplyFilter": ReRankingApplyFilter.from_dict(
                    obj.get("reRankingApplyFilter")
                )
                if obj.get("reRankingApplyFilter") is not None
                else None,
                "facet": obj.get("facet"),
                "indexName": obj.get("indexName"),
                "facetQuery": obj.get("facetQuery")
                if obj.get("facetQuery") is not None
                else "",
                "type": obj.get("type"),
            }
        )
        return _obj
