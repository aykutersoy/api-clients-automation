// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

public struct GetUserTokenResponse: Codable, JSONEncodable {
    /// Unique pseudonymous or anonymous user identifier.  This helps with analytics and click and conversion events.
    /// For more information, see [user token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
    public var userToken: String
    /// Date and time of the last event from this user, in RFC 3339 format.
    public var lastEventAt: String
    /// Scores for different facet values.  Scores represent the user affinity for a user profile towards specific facet
    /// values, given the personalization strategy and past events.
    public var scores: AnyCodable

    public init(userToken: String, lastEventAt: String, scores: AnyCodable) {
        self.userToken = userToken
        self.lastEventAt = lastEventAt
        self.scores = scores
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case userToken
        case lastEventAt
        case scores
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.userToken, forKey: .userToken)
        try container.encode(self.lastEventAt, forKey: .lastEventAt)
        try container.encode(self.scores, forKey: .scores)
    }
}

extension GetUserTokenResponse: Equatable {
    public static func ==(lhs: GetUserTokenResponse, rhs: GetUserTokenResponse) -> Bool {
        lhs.userToken == rhs.userToken &&
            lhs.lastEventAt == rhs.lastEventAt &&
            lhs.scores == rhs.scores
    }
}

extension GetUserTokenResponse: Hashable {
    public func hash(into hasher: inout Hasher) {
        hasher.combine(self.userToken.hashValue)
        hasher.combine(self.lastEventAt.hashValue)
        hasher.combine(self.scores.hashValue)
    }
}
