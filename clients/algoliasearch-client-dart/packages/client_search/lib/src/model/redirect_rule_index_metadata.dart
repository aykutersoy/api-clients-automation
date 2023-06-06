// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_client_search/src/model/redirect_rule_index_metadata_data.dart';

import 'package:json_annotation/json_annotation.dart';

part 'redirect_rule_index_metadata.g.dart';

@JsonSerializable()
final class RedirectRuleIndexMetadata {
  /// Returns a new [RedirectRuleIndexMetadata] instance.
  const RedirectRuleIndexMetadata({
    required this.source_,
    required this.dest,
    required this.reason,
    required this.succeed,
    required this.data,
  });

  /// Source index for the redirect rule.
  @JsonKey(name: r'source')
  final String source_;

  /// Destination index for the redirect rule.
  @JsonKey(name: r'dest')
  final String dest;

  /// Reason for the redirect rule.
  @JsonKey(name: r'reason')
  final String reason;

  /// Status for the redirect rule.
  @JsonKey(name: r'succeed')
  final bool succeed;

  @JsonKey(name: r'data')
  final RedirectRuleIndexMetadataData data;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is RedirectRuleIndexMetadata &&
          other.source_ == source_ &&
          other.dest == dest &&
          other.reason == reason &&
          other.succeed == succeed &&
          other.data == data;

  @override
  int get hashCode =>
      source_.hashCode +
      dest.hashCode +
      reason.hashCode +
      succeed.hashCode +
      data.hashCode;

  factory RedirectRuleIndexMetadata.fromJson(Map<String, dynamic> json) =>
      _$RedirectRuleIndexMetadataFromJson(json);

  Map<String, dynamic> toJson() => _$RedirectRuleIndexMetadataToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
