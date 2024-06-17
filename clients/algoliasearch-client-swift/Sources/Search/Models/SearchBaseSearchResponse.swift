// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

public struct SearchBaseSearchResponse: Codable, JSONEncodable {
    /// A/B test ID. This is only included in the response for indices that are part of an A/B test.
    public var abTestID: Int?
    /// Variant ID. This is only included in the response for indices that are part of an A/B test.
    public var abTestVariantID: Int?
    /// Computed geographical location.
    public var aroundLatLng: String?
    /// Distance from a central coordinate provided by `aroundLatLng`.
    public var automaticRadius: String?
    public var exhaustive: SearchExhaustive?
    /// See the `facetsCount` field of the `exhaustive` object in the response.
    @available(*, deprecated, message: "This property is deprecated.")
    public var exhaustiveFacetsCount: Bool?
    /// See the `nbHits` field of the `exhaustive` object in the response.
    @available(*, deprecated, message: "This property is deprecated.")
    public var exhaustiveNbHits: Bool?
    /// See the `typo` field of the `exhaustive` object in the response.
    @available(*, deprecated, message: "This property is deprecated.")
    public var exhaustiveTypo: Bool?
    /// Facet counts.
    public var facets: [String: [String: Int]]?
    /// Statistics for numerical facets.
    public var facetsStats: [String: SearchFacetsStats]?
    /// Number of hits per page.
    public var hitsPerPage: Int
    /// Index name used for the query.
    public var index: String?
    /// Index name used for the query. During A/B testing, the targeted index isn't always the index used by the query.
    public var indexUsed: String?
    /// Warnings about the query.
    public var message: String?
    /// Number of results (hits).
    public var nbHits: Int
    /// Number of pages of results.
    public var nbPages: Int
    /// Number of hits selected and sorted by the relevant sort algorithm.
    public var nbSortedHits: Int?
    /// Page of search results to retrieve.
    public var page: Int
    /// Post-[normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/#what-does-normalization-mean)
    /// query string that will be searched.
    public var parsedQuery: String?
    /// Time the server took to process the request, in milliseconds.
    public var processingTimeMS: Int
    /// Experimental. List of processing steps and their times, in milliseconds. You can use this list to investigate
    /// performance issues.
    public var processingTimingsMS: AnyCodable?
    /// Markup text indicating which parts of the original query have been removed to retrieve a non-empty result set.
    public var queryAfterRemoval: String?
    public var redirect: SearchRedirect?
    public var renderingContent: SearchRenderingContent?
    /// Time the server took to process the request, in milliseconds.
    public var serverTimeMS: Int?
    /// Host name of the server that processed the request.
    public var serverUsed: String?
    /// An object with custom data.  You can store up to 32kB as custom data.
    public var userData: AnyCodable?
    /// Unique identifier for the query. This is used for [click
    /// analytics](https://www.algolia.com/doc/guides/analytics/click-analytics/).
    public var queryID: String?

    public init(
        abTestID: Int? = nil,
        abTestVariantID: Int? = nil,
        aroundLatLng: String? = nil,
        automaticRadius: String? = nil,
        exhaustive: SearchExhaustive? = nil,
        exhaustiveFacetsCount: Bool? = nil,
        exhaustiveNbHits: Bool? = nil,
        exhaustiveTypo: Bool? = nil,
        facets: [String: [String: Int]]? = nil,
        facetsStats: [String: SearchFacetsStats]? = nil,
        hitsPerPage: Int,
        index: String? = nil,
        indexUsed: String? = nil,
        message: String? = nil,
        nbHits: Int,
        nbPages: Int,
        nbSortedHits: Int? = nil,
        page: Int,
        parsedQuery: String? = nil,
        processingTimeMS: Int,
        processingTimingsMS: AnyCodable? = nil,
        queryAfterRemoval: String? = nil,
        redirect: SearchRedirect? = nil,
        renderingContent: SearchRenderingContent? = nil,
        serverTimeMS: Int? = nil,
        serverUsed: String? = nil,
        userData: AnyCodable? = nil,
        queryID: String? = nil
    ) {
        self.abTestID = abTestID
        self.abTestVariantID = abTestVariantID
        self.aroundLatLng = aroundLatLng
        self.automaticRadius = automaticRadius
        self.exhaustive = exhaustive
        self.exhaustiveFacetsCount = exhaustiveFacetsCount
        self.exhaustiveNbHits = exhaustiveNbHits
        self.exhaustiveTypo = exhaustiveTypo
        self.facets = facets
        self.facetsStats = facetsStats
        self.hitsPerPage = hitsPerPage
        self.index = index
        self.indexUsed = indexUsed
        self.message = message
        self.nbHits = nbHits
        self.nbPages = nbPages
        self.nbSortedHits = nbSortedHits
        self.page = page
        self.parsedQuery = parsedQuery
        self.processingTimeMS = processingTimeMS
        self.processingTimingsMS = processingTimingsMS
        self.queryAfterRemoval = queryAfterRemoval
        self.redirect = redirect
        self.renderingContent = renderingContent
        self.serverTimeMS = serverTimeMS
        self.serverUsed = serverUsed
        self.userData = userData
        self.queryID = queryID
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case abTestID
        case abTestVariantID
        case aroundLatLng
        case automaticRadius
        case exhaustive
        case exhaustiveFacetsCount
        case exhaustiveNbHits
        case exhaustiveTypo
        case facets
        case facetsStats = "facets_stats"
        case hitsPerPage
        case index
        case indexUsed
        case message
        case nbHits
        case nbPages
        case nbSortedHits
        case page
        case parsedQuery
        case processingTimeMS
        case processingTimingsMS
        case queryAfterRemoval
        case redirect
        case renderingContent
        case serverTimeMS
        case serverUsed
        case userData
        case queryID
    }

    public var additionalProperties: [String: AnyCodable] = [:]

    public subscript(key: String) -> AnyCodable? {
        get {
            if let value = additionalProperties[key] {
                return value
            }
            return nil
        }

        set {
            self.additionalProperties[key] = newValue
        }
    }

    public init(from dictionary: [String: AnyCodable]) throws {
        self.abTestID = dictionary["abTestID"]?.value as? Int

        self.abTestVariantID = dictionary["abTestVariantID"]?.value as? Int

        self.aroundLatLng = dictionary["aroundLatLng"]?.value as? String

        self.automaticRadius = dictionary["automaticRadius"]?.value as? String

        self.exhaustive = dictionary["exhaustive"]?.value as? SearchExhaustive

        self.exhaustiveFacetsCount = dictionary["exhaustiveFacetsCount"]?.value as? Bool

        self.exhaustiveNbHits = dictionary["exhaustiveNbHits"]?.value as? Bool

        self.exhaustiveTypo = dictionary["exhaustiveTypo"]?.value as? Bool

        self.facets = dictionary["facets"]?.value as? [String: [String: Int]]

        self.facetsStats = dictionary["facetsStats"]?.value as? [String: SearchFacetsStats]

        guard let hitsPerPage = dictionary["hitsPerPage"]?.value as? Int else {
            throw GenericError(description: "Failed to cast")
        }
        self.hitsPerPage = hitsPerPage
        self.index = dictionary["index"]?.value as? String

        self.indexUsed = dictionary["indexUsed"]?.value as? String

        self.message = dictionary["message"]?.value as? String

        guard let nbHits = dictionary["nbHits"]?.value as? Int else {
            throw GenericError(description: "Failed to cast")
        }
        self.nbHits = nbHits
        guard let nbPages = dictionary["nbPages"]?.value as? Int else {
            throw GenericError(description: "Failed to cast")
        }
        self.nbPages = nbPages
        self.nbSortedHits = dictionary["nbSortedHits"]?.value as? Int

        guard let page = dictionary["page"]?.value as? Int else {
            throw GenericError(description: "Failed to cast")
        }
        self.page = page
        self.parsedQuery = dictionary["parsedQuery"]?.value as? String

        guard let processingTimeMS = dictionary["processingTimeMS"]?.value as? Int else {
            throw GenericError(description: "Failed to cast")
        }
        self.processingTimeMS = processingTimeMS
        self.processingTimingsMS = dictionary["processingTimingsMS"]?.value as? AnyCodable

        self.queryAfterRemoval = dictionary["queryAfterRemoval"]?.value as? String

        self.redirect = dictionary["redirect"]?.value as? SearchRedirect

        self.renderingContent = dictionary["renderingContent"]?.value as? SearchRenderingContent

        self.serverTimeMS = dictionary["serverTimeMS"]?.value as? Int

        self.serverUsed = dictionary["serverUsed"]?.value as? String

        self.userData = dictionary["userData"]?.value as? AnyCodable

        self.queryID = dictionary["queryID"]?.value as? String

        for (key, value) in dictionary {
            switch key {
            case "abTestID", "abTestVariantID", "aroundLatLng", "automaticRadius", "exhaustive",
                 "exhaustiveFacetsCount", "exhaustiveNbHits", "exhaustiveTypo", "facets", "facetsStats", "hitsPerPage",
                 "index", "indexUsed", "message", "nbHits", "nbPages", "nbSortedHits", "page", "parsedQuery",
                 "processingTimeMS", "processingTimingsMS", "queryAfterRemoval", "redirect", "renderingContent",
                 "serverTimeMS", "serverUsed", "userData", "queryID":
                continue
            default:
                self.additionalProperties[key] = value
            }
        }
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(self.abTestID, forKey: .abTestID)
        try container.encodeIfPresent(self.abTestVariantID, forKey: .abTestVariantID)
        try container.encodeIfPresent(self.aroundLatLng, forKey: .aroundLatLng)
        try container.encodeIfPresent(self.automaticRadius, forKey: .automaticRadius)
        try container.encodeIfPresent(self.exhaustive, forKey: .exhaustive)
        try container.encodeIfPresent(self.exhaustiveFacetsCount, forKey: .exhaustiveFacetsCount)
        try container.encodeIfPresent(self.exhaustiveNbHits, forKey: .exhaustiveNbHits)
        try container.encodeIfPresent(self.exhaustiveTypo, forKey: .exhaustiveTypo)
        try container.encodeIfPresent(self.facets, forKey: .facets)
        try container.encodeIfPresent(self.facetsStats, forKey: .facetsStats)
        try container.encode(self.hitsPerPage, forKey: .hitsPerPage)
        try container.encodeIfPresent(self.index, forKey: .index)
        try container.encodeIfPresent(self.indexUsed, forKey: .indexUsed)
        try container.encodeIfPresent(self.message, forKey: .message)
        try container.encode(self.nbHits, forKey: .nbHits)
        try container.encode(self.nbPages, forKey: .nbPages)
        try container.encodeIfPresent(self.nbSortedHits, forKey: .nbSortedHits)
        try container.encode(self.page, forKey: .page)
        try container.encodeIfPresent(self.parsedQuery, forKey: .parsedQuery)
        try container.encode(self.processingTimeMS, forKey: .processingTimeMS)
        try container.encodeIfPresent(self.processingTimingsMS, forKey: .processingTimingsMS)
        try container.encodeIfPresent(self.queryAfterRemoval, forKey: .queryAfterRemoval)
        try container.encodeIfPresent(self.redirect, forKey: .redirect)
        try container.encodeIfPresent(self.renderingContent, forKey: .renderingContent)
        try container.encodeIfPresent(self.serverTimeMS, forKey: .serverTimeMS)
        try container.encodeIfPresent(self.serverUsed, forKey: .serverUsed)
        try container.encodeIfPresent(self.userData, forKey: .userData)
        try container.encodeIfPresent(self.queryID, forKey: .queryID)
        var additionalPropertiesContainer = encoder.container(keyedBy: String.self)
        try additionalPropertiesContainer.encodeMap(self.additionalProperties)
    }

    // Decodable protocol methods

    public init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)

        self.abTestID = try container.decodeIfPresent(Int.self, forKey: .abTestID)
        self.abTestVariantID = try container.decodeIfPresent(Int.self, forKey: .abTestVariantID)
        self.aroundLatLng = try container.decodeIfPresent(String.self, forKey: .aroundLatLng)
        self.automaticRadius = try container.decodeIfPresent(String.self, forKey: .automaticRadius)
        self.exhaustive = try container.decodeIfPresent(SearchExhaustive.self, forKey: .exhaustive)
        self.exhaustiveFacetsCount = try container.decodeIfPresent(Bool.self, forKey: .exhaustiveFacetsCount)
        self.exhaustiveNbHits = try container.decodeIfPresent(Bool.self, forKey: .exhaustiveNbHits)
        self.exhaustiveTypo = try container.decodeIfPresent(Bool.self, forKey: .exhaustiveTypo)
        self.facets = try container.decodeIfPresent([String: [String: Int]].self, forKey: .facets)
        self.facetsStats = try container.decodeIfPresent([String: SearchFacetsStats].self, forKey: .facetsStats)
        self.hitsPerPage = try container.decode(Int.self, forKey: .hitsPerPage)
        self.index = try container.decodeIfPresent(String.self, forKey: .index)
        self.indexUsed = try container.decodeIfPresent(String.self, forKey: .indexUsed)
        self.message = try container.decodeIfPresent(String.self, forKey: .message)
        self.nbHits = try container.decode(Int.self, forKey: .nbHits)
        self.nbPages = try container.decode(Int.self, forKey: .nbPages)
        self.nbSortedHits = try container.decodeIfPresent(Int.self, forKey: .nbSortedHits)
        self.page = try container.decode(Int.self, forKey: .page)
        self.parsedQuery = try container.decodeIfPresent(String.self, forKey: .parsedQuery)
        self.processingTimeMS = try container.decode(Int.self, forKey: .processingTimeMS)
        self.processingTimingsMS = try container.decodeIfPresent(AnyCodable.self, forKey: .processingTimingsMS)
        self.queryAfterRemoval = try container.decodeIfPresent(String.self, forKey: .queryAfterRemoval)
        self.redirect = try container.decodeIfPresent(SearchRedirect.self, forKey: .redirect)
        self.renderingContent = try container.decodeIfPresent(SearchRenderingContent.self, forKey: .renderingContent)
        self.serverTimeMS = try container.decodeIfPresent(Int.self, forKey: .serverTimeMS)
        self.serverUsed = try container.decodeIfPresent(String.self, forKey: .serverUsed)
        self.userData = try container.decodeIfPresent(AnyCodable.self, forKey: .userData)
        self.queryID = try container.decodeIfPresent(String.self, forKey: .queryID)
        var nonAdditionalPropertyKeys = Set<String>()
        nonAdditionalPropertyKeys.insert("abTestID")
        nonAdditionalPropertyKeys.insert("abTestVariantID")
        nonAdditionalPropertyKeys.insert("aroundLatLng")
        nonAdditionalPropertyKeys.insert("automaticRadius")
        nonAdditionalPropertyKeys.insert("exhaustive")
        nonAdditionalPropertyKeys.insert("exhaustiveFacetsCount")
        nonAdditionalPropertyKeys.insert("exhaustiveNbHits")
        nonAdditionalPropertyKeys.insert("exhaustiveTypo")
        nonAdditionalPropertyKeys.insert("facets")
        nonAdditionalPropertyKeys.insert("facets_stats")
        nonAdditionalPropertyKeys.insert("hitsPerPage")
        nonAdditionalPropertyKeys.insert("index")
        nonAdditionalPropertyKeys.insert("indexUsed")
        nonAdditionalPropertyKeys.insert("message")
        nonAdditionalPropertyKeys.insert("nbHits")
        nonAdditionalPropertyKeys.insert("nbPages")
        nonAdditionalPropertyKeys.insert("nbSortedHits")
        nonAdditionalPropertyKeys.insert("page")
        nonAdditionalPropertyKeys.insert("parsedQuery")
        nonAdditionalPropertyKeys.insert("processingTimeMS")
        nonAdditionalPropertyKeys.insert("processingTimingsMS")
        nonAdditionalPropertyKeys.insert("queryAfterRemoval")
        nonAdditionalPropertyKeys.insert("redirect")
        nonAdditionalPropertyKeys.insert("renderingContent")
        nonAdditionalPropertyKeys.insert("serverTimeMS")
        nonAdditionalPropertyKeys.insert("serverUsed")
        nonAdditionalPropertyKeys.insert("userData")
        nonAdditionalPropertyKeys.insert("queryID")
        let additionalPropertiesContainer = try decoder.container(keyedBy: String.self)
        self.additionalProperties = try additionalPropertiesContainer.decodeMap(
            AnyCodable.self,
            excludedKeys: nonAdditionalPropertyKeys
        )
    }
}

extension SearchBaseSearchResponse: Equatable {
    public static func ==(lhs: SearchBaseSearchResponse, rhs: SearchBaseSearchResponse) -> Bool {
        lhs.abTestID == rhs.abTestID &&
            lhs.abTestVariantID == rhs.abTestVariantID &&
            lhs.aroundLatLng == rhs.aroundLatLng &&
            lhs.automaticRadius == rhs.automaticRadius &&
            lhs.exhaustive == rhs.exhaustive &&
            lhs.exhaustiveFacetsCount == rhs.exhaustiveFacetsCount &&
            lhs.exhaustiveNbHits == rhs.exhaustiveNbHits &&
            lhs.exhaustiveTypo == rhs.exhaustiveTypo &&
            lhs.facets == rhs.facets &&
            lhs.facetsStats == rhs.facetsStats &&
            lhs.hitsPerPage == rhs.hitsPerPage &&
            lhs.index == rhs.index &&
            lhs.indexUsed == rhs.indexUsed &&
            lhs.message == rhs.message &&
            lhs.nbHits == rhs.nbHits &&
            lhs.nbPages == rhs.nbPages &&
            lhs.nbSortedHits == rhs.nbSortedHits &&
            lhs.page == rhs.page &&
            lhs.parsedQuery == rhs.parsedQuery &&
            lhs.processingTimeMS == rhs.processingTimeMS &&
            lhs.processingTimingsMS == rhs.processingTimingsMS &&
            lhs.queryAfterRemoval == rhs.queryAfterRemoval &&
            lhs.redirect == rhs.redirect &&
            lhs.renderingContent == rhs.renderingContent &&
            lhs.serverTimeMS == rhs.serverTimeMS &&
            lhs.serverUsed == rhs.serverUsed &&
            lhs.userData == rhs.userData &&
            lhs.queryID == rhs.queryID
            && lhs.additionalProperties == rhs.additionalProperties
    }
}

extension SearchBaseSearchResponse: Hashable {
    public func hash(into hasher: inout Hasher) {
        hasher.combine(self.abTestID?.hashValue)
        hasher.combine(self.abTestVariantID?.hashValue)
        hasher.combine(self.aroundLatLng?.hashValue)
        hasher.combine(self.automaticRadius?.hashValue)
        hasher.combine(self.exhaustive?.hashValue)
        hasher.combine(self.exhaustiveFacetsCount?.hashValue)
        hasher.combine(self.exhaustiveNbHits?.hashValue)
        hasher.combine(self.exhaustiveTypo?.hashValue)
        hasher.combine(self.facets?.hashValue)
        hasher.combine(self.facetsStats?.hashValue)
        hasher.combine(self.hitsPerPage.hashValue)
        hasher.combine(self.index?.hashValue)
        hasher.combine(self.indexUsed?.hashValue)
        hasher.combine(self.message?.hashValue)
        hasher.combine(self.nbHits.hashValue)
        hasher.combine(self.nbPages.hashValue)
        hasher.combine(self.nbSortedHits?.hashValue)
        hasher.combine(self.page.hashValue)
        hasher.combine(self.parsedQuery?.hashValue)
        hasher.combine(self.processingTimeMS.hashValue)
        hasher.combine(self.processingTimingsMS?.hashValue)
        hasher.combine(self.queryAfterRemoval?.hashValue)
        hasher.combine(self.redirect?.hashValue)
        hasher.combine(self.renderingContent?.hashValue)
        hasher.combine(self.serverTimeMS?.hashValue)
        hasher.combine(self.serverUsed?.hashValue)
        hasher.combine(self.userData?.hashValue)
        hasher.combine(self.queryID?.hashValue)
        hasher.combine(self.additionalProperties.hashValue)
    }
}
