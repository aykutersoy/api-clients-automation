// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algoliasearch/src/model/base_search_response_redirect.dart';
import 'package:algoliasearch/src/model/hit.dart';
import 'package:algoliasearch/src/model/facets_stats.dart';
import 'package:algoliasearch/src/model/rendering_content.dart';

import 'package:json_annotation/json_annotation.dart';

part 'browse_response.g.dart';

@JsonSerializable()
final class BrowseResponse {
  /// Returns a new [BrowseResponse] instance.
  const BrowseResponse({
    this.abTestID,
    this.abTestVariantID,
    this.aroundLatLng,
    this.automaticRadius,
    this.exhaustiveFacetsCount,
    required this.exhaustiveNbHits,
    this.exhaustiveTypo,
    this.facets,
    this.facetsStats,
    required this.hitsPerPage,
    this.index,
    this.indexUsed,
    this.message,
    required this.nbHits,
    required this.nbPages,
    this.nbSortedHits,
    required this.page,
    required this.params,
    this.redirect,
    this.parsedQuery,
    required this.processingTimeMS,
    required this.query,
    this.queryAfterRemoval,
    this.serverUsed,
    this.userData,
    this.renderingContent,
    required this.hits,
    this.cursor,
  });

  /// If a search encounters an index that is being A/B tested, abTestID reports the ongoing A/B test ID.
  @JsonKey(name: r'abTestID')
  final int? abTestID;

  /// If a search encounters an index that is being A/B tested, abTestVariantID reports the variant ID of the index used (starting at 1).
  @JsonKey(name: r'abTestVariantID')
  final int? abTestVariantID;

  /// The computed geo location.
  @JsonKey(name: r'aroundLatLng')
  final String? aroundLatLng;

  /// The automatically computed radius. For legacy reasons, this parameter is a string and not an integer.
  @JsonKey(name: r'automaticRadius')
  final String? automaticRadius;

  /// Whether the facet count is exhaustive or approximate.
  @JsonKey(name: r'exhaustiveFacetsCount')
  final bool? exhaustiveFacetsCount;

  /// Indicate if the nbHits count was exhaustive or approximate.
  @JsonKey(name: r'exhaustiveNbHits')
  final bool exhaustiveNbHits;

  /// Indicate if the typo-tolerance search was exhaustive or approximate (only included when typo-tolerance is enabled).
  @JsonKey(name: r'exhaustiveTypo')
  final bool? exhaustiveTypo;

  /// A mapping of each facet name to the corresponding facet counts.
  @JsonKey(name: r'facets')
  final Map<String, Map<String, int>>? facets;

  /// Statistics for numerical facets.
  @JsonKey(name: r'facets_stats')
  final Map<String, FacetsStats>? facetsStats;

  /// Set the number of hits per page.
  @JsonKey(name: r'hitsPerPage')
  final int hitsPerPage;

  /// Index name used for the query.
  @JsonKey(name: r'index')
  final String? index;

  /// Index name used for the query. In the case of an A/B test, the targeted index isn't always the index used by the query.
  @JsonKey(name: r'indexUsed')
  final String? indexUsed;

  /// Used to return warnings about the query.
  @JsonKey(name: r'message')
  final String? message;

  /// Number of hits that the search query matched.
  @JsonKey(name: r'nbHits')
  final int nbHits;

  /// Number of pages available for the current query.
  @JsonKey(name: r'nbPages')
  final int nbPages;

  /// The number of hits selected and sorted by the relevant sort algorithm.
  @JsonKey(name: r'nbSortedHits')
  final int? nbSortedHits;

  /// Specify the page to retrieve.
  @JsonKey(name: r'page')
  final int page;

  /// A url-encoded string of all search parameters.
  @JsonKey(name: r'params')
  final String params;

  @JsonKey(name: r'redirect')
  final BaseSearchResponseRedirect? redirect;

  /// The query string that will be searched, after normalization.
  @JsonKey(name: r'parsedQuery')
  final String? parsedQuery;

  /// Time the server took to process the request, in milliseconds.
  @JsonKey(name: r'processingTimeMS')
  final int processingTimeMS;

  /// The text to search in the index.
  @JsonKey(name: r'query')
  final String query;

  /// A markup text indicating which parts of the original query have been removed in order to retrieve a non-empty result set.
  @JsonKey(name: r'queryAfterRemoval')
  final String? queryAfterRemoval;

  /// Actual host name of the server that processed the request.
  @JsonKey(name: r'serverUsed')
  final String? serverUsed;

  /// Lets you store custom data in your indices.
  @JsonKey(name: r'userData')
  final Object? userData;

  @JsonKey(name: r'renderingContent')
  final RenderingContent? renderingContent;

  @JsonKey(name: r'hits')
  final List<Hit> hits;

  /// Cursor indicating the location to resume browsing from. Must match the value returned by the previous call.
  @JsonKey(name: r'cursor')
  final String? cursor;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is BrowseResponse &&
          other.abTestID == abTestID &&
          other.abTestVariantID == abTestVariantID &&
          other.aroundLatLng == aroundLatLng &&
          other.automaticRadius == automaticRadius &&
          other.exhaustiveFacetsCount == exhaustiveFacetsCount &&
          other.exhaustiveNbHits == exhaustiveNbHits &&
          other.exhaustiveTypo == exhaustiveTypo &&
          other.facets == facets &&
          other.facetsStats == facetsStats &&
          other.hitsPerPage == hitsPerPage &&
          other.index == index &&
          other.indexUsed == indexUsed &&
          other.message == message &&
          other.nbHits == nbHits &&
          other.nbPages == nbPages &&
          other.nbSortedHits == nbSortedHits &&
          other.page == page &&
          other.params == params &&
          other.redirect == redirect &&
          other.parsedQuery == parsedQuery &&
          other.processingTimeMS == processingTimeMS &&
          other.query == query &&
          other.queryAfterRemoval == queryAfterRemoval &&
          other.serverUsed == serverUsed &&
          other.userData == userData &&
          other.renderingContent == renderingContent &&
          other.hits == hits &&
          other.cursor == cursor;

  @override
  int get hashCode =>
      abTestID.hashCode +
      abTestVariantID.hashCode +
      aroundLatLng.hashCode +
      automaticRadius.hashCode +
      exhaustiveFacetsCount.hashCode +
      exhaustiveNbHits.hashCode +
      exhaustiveTypo.hashCode +
      facets.hashCode +
      facetsStats.hashCode +
      hitsPerPage.hashCode +
      index.hashCode +
      indexUsed.hashCode +
      message.hashCode +
      nbHits.hashCode +
      nbPages.hashCode +
      nbSortedHits.hashCode +
      page.hashCode +
      params.hashCode +
      redirect.hashCode +
      parsedQuery.hashCode +
      processingTimeMS.hashCode +
      query.hashCode +
      queryAfterRemoval.hashCode +
      serverUsed.hashCode +
      userData.hashCode +
      renderingContent.hashCode +
      hits.hashCode +
      cursor.hashCode;

  factory BrowseResponse.fromJson(Map<String, dynamic> json) =>
      _$BrowseResponseFromJson(json);

  Map<String, dynamic> toJson() => _$BrowseResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
