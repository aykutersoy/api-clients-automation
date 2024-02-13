// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

public struct Variant: Codable, JSONEncodable, Hashable {
    /** Number of add-to-cart events for this variant. */
    public var addToCartCount: Int
    /** Variant's [add-to-cart rate](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#add-to-cart-rate). */
    public var addToCartRate: Double?
    /** Variant's [average click position](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-position). */
    public var averageClickPosition: Int?
    /** Number of click events for this variant. */
    public var clickCount: Int
    /** Variant's [click-through rate](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate). */
    public var clickThroughRate: Double?
    /** Number of click events for this variant. */
    public var conversionCount: Int
    /** Variant's [conversion rate](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#conversion-rate). */
    public var conversionRate: Double?
    /** A/B test currencies. */
    public var currencies: [String: CurrenciesValue]?
    /** A/B test description. */
    public var description: String
    /** The estimated number of searches that will need to be run to achieve the desired confidence level and statistical power. A `minimumDetectableEffect` must be set in the `configuration` object for this to be used. */
    public var estimatedSampleSize: Int?
    public var filterEffects: FilterEffects?
    /** A/B test index. */
    public var index: String
    /** Number of [searches without results](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#searches-without-results) for that variant. */
    public var noResultCount: Int?
    /** Number of purchase events for this variant. */
    public var purchaseCount: Int
    /** Variant's [purchase rate](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#purchase-rate). */
    public var purchaseRate: Double?
    /** Number of searches carried out during the A/B test. */
    public var searchCount: Int?
    /** Number of tracked searches. This is the number of search requests where the `clickAnalytics` parameter is `true`. */
    public var trackedSearchCount: Int?
    /** A/B test traffic percentage. */
    public var trafficPercentage: Int
    /** Number of users during the A/B test. */
    public var userCount: Int?
    /** Number of users that performed a tracked search during the A/B test. */
    public var trackedUserCount: Int?

    public init(addToCartCount: Int, addToCartRate: Double?, averageClickPosition: Int?, clickCount: Int, clickThroughRate: Double?, conversionCount: Int, conversionRate: Double?, currencies: [String: CurrenciesValue]? = nil, description: String, estimatedSampleSize: Int? = nil, filterEffects: FilterEffects? = nil, index: String, noResultCount: Int?, purchaseCount: Int, purchaseRate: Double?, searchCount: Int?, trackedSearchCount: Int?, trafficPercentage: Int, userCount: Int?, trackedUserCount: Int?) {
        self.addToCartCount = addToCartCount
        self.addToCartRate = addToCartRate
        self.averageClickPosition = averageClickPosition
        self.clickCount = clickCount
        self.clickThroughRate = clickThroughRate
        self.conversionCount = conversionCount
        self.conversionRate = conversionRate
        self.currencies = currencies
        self.description = description
        self.estimatedSampleSize = estimatedSampleSize
        self.filterEffects = filterEffects
        self.index = index
        self.noResultCount = noResultCount
        self.purchaseCount = purchaseCount
        self.purchaseRate = purchaseRate
        self.searchCount = searchCount
        self.trackedSearchCount = trackedSearchCount
        self.trafficPercentage = trafficPercentage
        self.userCount = userCount
        self.trackedUserCount = trackedUserCount
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case addToCartCount
        case addToCartRate
        case averageClickPosition
        case clickCount
        case clickThroughRate
        case conversionCount
        case conversionRate
        case currencies
        case description
        case estimatedSampleSize
        case filterEffects
        case index
        case noResultCount
        case purchaseCount
        case purchaseRate
        case searchCount
        case trackedSearchCount
        case trafficPercentage
        case userCount
        case trackedUserCount
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(addToCartCount, forKey: .addToCartCount)
        try container.encode(addToCartRate, forKey: .addToCartRate)
        try container.encode(averageClickPosition, forKey: .averageClickPosition)
        try container.encode(clickCount, forKey: .clickCount)
        try container.encode(clickThroughRate, forKey: .clickThroughRate)
        try container.encode(conversionCount, forKey: .conversionCount)
        try container.encode(conversionRate, forKey: .conversionRate)
        try container.encodeIfPresent(currencies, forKey: .currencies)
        try container.encode(description, forKey: .description)
        try container.encodeIfPresent(estimatedSampleSize, forKey: .estimatedSampleSize)
        try container.encodeIfPresent(filterEffects, forKey: .filterEffects)
        try container.encode(index, forKey: .index)
        try container.encode(noResultCount, forKey: .noResultCount)
        try container.encode(purchaseCount, forKey: .purchaseCount)
        try container.encode(purchaseRate, forKey: .purchaseRate)
        try container.encode(searchCount, forKey: .searchCount)
        try container.encode(trackedSearchCount, forKey: .trackedSearchCount)
        try container.encode(trafficPercentage, forKey: .trafficPercentage)
        try container.encode(userCount, forKey: .userCount)
        try container.encode(trackedUserCount, forKey: .trackedUserCount)
    }
}
