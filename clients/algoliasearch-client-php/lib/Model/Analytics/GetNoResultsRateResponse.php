<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Analytics;

use Algolia\AlgoliaSearch\Model\AbstractModel;

/**
 * GetNoResultsRateResponse Class Doc Comment.
 *
 * @category Class
 */
class GetNoResultsRateResponse extends AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'rate' => 'float',
        'count' => 'int',
        'noResultCount' => 'int',
        'dates' => '\Algolia\AlgoliaSearch\Model\Analytics\DailyNoResultsRates[]',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'rate' => 'double',
        'count' => null,
        'noResultCount' => null,
        'dates' => null,
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'rate' => 'rate',
        'count' => 'count',
        'noResultCount' => 'noResultCount',
        'dates' => 'dates',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'rate' => 'setRate',
        'count' => 'setCount',
        'noResultCount' => 'setNoResultCount',
        'dates' => 'setDates',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'rate' => 'getRate',
        'count' => 'getCount',
        'noResultCount' => 'getNoResultCount',
        'dates' => 'getDates',
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
        if (isset($data['rate'])) {
            $this->container['rate'] = $data['rate'];
        }
        if (isset($data['count'])) {
            $this->container['count'] = $data['count'];
        }
        if (isset($data['noResultCount'])) {
            $this->container['noResultCount'] = $data['noResultCount'];
        }
        if (isset($data['dates'])) {
            $this->container['dates'] = $data['dates'];
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

        if (!isset($this->container['rate']) || null === $this->container['rate']) {
            $invalidProperties[] = "'rate' can't be null";
        }
        if ($this->container['rate'] > 1) {
            $invalidProperties[] = "invalid value for 'rate', must be smaller than or equal to 1.";
        }

        if ($this->container['rate'] < 0) {
            $invalidProperties[] = "invalid value for 'rate', must be bigger than or equal to 0.";
        }

        if (!isset($this->container['count']) || null === $this->container['count']) {
            $invalidProperties[] = "'count' can't be null";
        }
        if (!isset($this->container['noResultCount']) || null === $this->container['noResultCount']) {
            $invalidProperties[] = "'noResultCount' can't be null";
        }
        if (!isset($this->container['dates']) || null === $this->container['dates']) {
            $invalidProperties[] = "'dates' can't be null";
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
     * Gets rate.
     *
     * @return float
     */
    public function getRate()
    {
        return $this->container['rate'] ?? null;
    }

    /**
     * Sets rate.
     *
     * @param float $rate no results rate, calculated as number of searches with zero results divided by the total number of searches
     *
     * @return self
     */
    public function setRate($rate)
    {
        if ($rate > 1) {
            throw new \InvalidArgumentException('invalid value for $rate when calling GetNoResultsRateResponse., must be smaller than or equal to 1.');
        }
        if ($rate < 0) {
            throw new \InvalidArgumentException('invalid value for $rate when calling GetNoResultsRateResponse., must be bigger than or equal to 0.');
        }

        $this->container['rate'] = $rate;

        return $this;
    }

    /**
     * Gets count.
     *
     * @return int
     */
    public function getCount()
    {
        return $this->container['count'] ?? null;
    }

    /**
     * Sets count.
     *
     * @param int $count number of searches
     *
     * @return self
     */
    public function setCount($count)
    {
        $this->container['count'] = $count;

        return $this;
    }

    /**
     * Gets noResultCount.
     *
     * @return int
     */
    public function getNoResultCount()
    {
        return $this->container['noResultCount'] ?? null;
    }

    /**
     * Sets noResultCount.
     *
     * @param int $noResultCount number of searches without any results
     *
     * @return self
     */
    public function setNoResultCount($noResultCount)
    {
        $this->container['noResultCount'] = $noResultCount;

        return $this;
    }

    /**
     * Gets dates.
     *
     * @return \Algolia\AlgoliaSearch\Model\Analytics\DailyNoResultsRates[]
     */
    public function getDates()
    {
        return $this->container['dates'] ?? null;
    }

    /**
     * Sets dates.
     *
     * @param \Algolia\AlgoliaSearch\Model\Analytics\DailyNoResultsRates[] $dates daily no results rates
     *
     * @return self
     */
    public function setDates($dates)
    {
        $this->container['dates'] = $dates;

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
