<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Monitoring;

/**
 * IncidentsInner Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class IncidentsInner extends \Algolia\AlgoliaSearch\Model\AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $modelTypes = [
        't' => 'int',
        'v' => '\Algolia\AlgoliaSearch\Model\Monitoring\Incident',
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $modelFormats = [
        't' => 'int64',
        'v' => null,
    ];

    /**
      * Array of attributes where the key is the local name,
      * and the value is the original name
      *
      * @var string[]
    */
    protected static $attributeMap = [
        't' => 't',
        'v' => 'v',
    ];

    /**
      * Array of attributes where the key is the local name,
      * and the value is the original name
      *
      * @return array
      */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        't' => 'setT',
        'v' => 'setV',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        't' => 'getT',
        'v' => 'getV',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(array $data = null)
    {
        if (isset($data['t'])) {
            $this->container['t'] = $data['t'];
        }
        if (isset($data['v'])) {
            $this->container['v'] = $data['v'];
        }
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }

    /**
     * Gets t
     *
     * @return int|null
     */
    public function getT()
    {
        return $this->container['t'] ?? null;
    }

    /**
     * Sets t
     *
     * @param int|null $t Timestamp in [Unix epoch time](https://wikipedia.org/wiki/Unix_time) in milliseconds.
     *
     * @return self
     */
    public function setT($t)
    {
        $this->container['t'] = $t;

        return $this;
    }

    /**
     * Gets v
     *
     * @return \Algolia\AlgoliaSearch\Model\Monitoring\Incident|null
     */
    public function getV()
    {
        return $this->container['v'] ?? null;
    }

    /**
     * Sets v
     *
     * @param \Algolia\AlgoliaSearch\Model\Monitoring\Incident|null $v v
     *
     * @return self
     */
    public function setV($v)
    {
        $this->container['v'] = $v;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     *
     * @return bool
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
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
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}

