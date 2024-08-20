<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Ingestion;

use Algolia\AlgoliaSearch\Model\AbstractModel;

/**
 * ListSourcesResponse Class Doc Comment.
 *
 * @category Class
 */
class ListSourcesResponse extends AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'sources' => '\Algolia\AlgoliaSearch\Model\Ingestion\Source[]',
        'pagination' => '\Algolia\AlgoliaSearch\Model\Ingestion\Pagination',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'sources' => null,
        'pagination' => null,
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'sources' => 'sources',
        'pagination' => 'pagination',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'sources' => 'setSources',
        'pagination' => 'setPagination',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'sources' => 'getSources',
        'pagination' => 'getPagination',
    ];

    /**
     * Associative array for storing property values.
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor.
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(?array $data = null)
    {
        if (isset($data['sources'])) {
            $this->container['sources'] = $data['sources'];
        }
        if (isset($data['pagination'])) {
            $this->container['pagination'] = $data['pagination'];
        }
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if (!isset($this->container['sources']) || null === $this->container['sources']) {
            $invalidProperties[] = "'sources' can't be null";
        }
        if (!isset($this->container['pagination']) || null === $this->container['pagination']) {
            $invalidProperties[] = "'pagination' can't be null";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed.
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return 0 === count($this->listInvalidProperties());
    }

    /**
     * Gets sources.
     *
     * @return \Algolia\AlgoliaSearch\Model\Ingestion\Source[]
     */
    public function getSources()
    {
        return $this->container['sources'] ?? null;
    }

    /**
     * Sets sources.
     *
     * @param \Algolia\AlgoliaSearch\Model\Ingestion\Source[] $sources sources
     *
     * @return self
     */
    public function setSources($sources)
    {
        $this->container['sources'] = $sources;

        return $this;
    }

    /**
     * Gets pagination.
     *
     * @return Pagination
     */
    public function getPagination()
    {
        return $this->container['pagination'] ?? null;
    }

    /**
     * Sets pagination.
     *
     * @param Pagination $pagination pagination
     *
     * @return self
     */
    public function setPagination($pagination)
    {
        $this->container['pagination'] = $pagination;

        return $this;
    }

    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     */
    public function offsetExists($offset): bool
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return null|mixed
     */
    public function offsetGet($offset): mixed
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param null|int $offset Offset
     * @param mixed    $value  Value to be set
     */
    public function offsetSet($offset, $value): void
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     */
    public function offsetUnset($offset): void
    {
        unset($this->container[$offset]);
    }
}
