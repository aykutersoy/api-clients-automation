<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Recommend;

/**
 * QueryType Class Doc Comment.
 *
 * @category Class
 * @description Determines how query words are [interpreted as prefixes](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/in-depth/prefix-searching/).
 */
class QueryType
{
    /**
     * Possible values of this enum.
     */
    public const PREFIX_LAST = 'prefixLast';

    public const PREFIX_ALL = 'prefixAll';

    public const PREFIX_NONE = 'prefixNone';

    /**
     * Gets allowable values of the enum.
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::PREFIX_LAST,
            self::PREFIX_ALL,
            self::PREFIX_NONE,
        ];
    }
}
