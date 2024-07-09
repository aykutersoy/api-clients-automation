<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Recommend;

use Algolia\AlgoliaSearch\Model\AbstractModel;

/**
 * RecommendRule Class Doc Comment.
 *
 * @category Class
 *
 * @description Recommend rule.
 */
class RecommendRule extends AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'metadata' => '\Algolia\AlgoliaSearch\Model\Recommend\RecommendRuleMetadata',
        'objectID' => 'string',
        'condition' => '\Algolia\AlgoliaSearch\Model\Recommend\Condition',
        'consequence' => '\Algolia\AlgoliaSearch\Model\Recommend\Consequence',
        'description' => 'string',
        'enabled' => 'bool',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'metadata' => null,
        'objectID' => null,
        'condition' => null,
        'consequence' => null,
        'description' => null,
        'enabled' => null,
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'metadata' => '_metadata',
        'objectID' => 'objectID',
        'condition' => 'condition',
        'consequence' => 'consequence',
        'description' => 'description',
        'enabled' => 'enabled',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'metadata' => 'setMetadata',
        'objectID' => 'setObjectID',
        'condition' => 'setCondition',
        'consequence' => 'setConsequence',
        'description' => 'setDescription',
        'enabled' => 'setEnabled',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'metadata' => 'getMetadata',
        'objectID' => 'getObjectID',
        'condition' => 'getCondition',
        'consequence' => 'getConsequence',
        'description' => 'getDescription',
        'enabled' => 'getEnabled',
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
        if (isset($data['metadata'])) {
            $this->container['metadata'] = $data['metadata'];
        }
        if (isset($data['objectID'])) {
            $this->container['objectID'] = $data['objectID'];
        }
        if (isset($data['condition'])) {
            $this->container['condition'] = $data['condition'];
        }
        if (isset($data['consequence'])) {
            $this->container['consequence'] = $data['consequence'];
        }
        if (isset($data['description'])) {
            $this->container['description'] = $data['description'];
        }
        if (isset($data['enabled'])) {
            $this->container['enabled'] = $data['enabled'];
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
        return [];
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
     * Gets metadata.
     *
     * @return null|RecommendRuleMetadata
     */
    public function getMetadata()
    {
        return $this->container['metadata'] ?? null;
    }

    /**
     * Sets metadata.
     *
     * @param null|RecommendRuleMetadata $metadata metadata
     *
     * @return self
     */
    public function setMetadata($metadata)
    {
        $this->container['metadata'] = $metadata;

        return $this;
    }

    /**
     * Gets objectID.
     *
     * @return null|string
     */
    public function getObjectID()
    {
        return $this->container['objectID'] ?? null;
    }

    /**
     * Sets objectID.
     *
     * @param null|string $objectID unique identifier of a rule object
     *
     * @return self
     */
    public function setObjectID($objectID)
    {
        $this->container['objectID'] = $objectID;

        return $this;
    }

    /**
     * Gets condition.
     *
     * @return null|Condition
     */
    public function getCondition()
    {
        return $this->container['condition'] ?? null;
    }

    /**
     * Sets condition.
     *
     * @param null|Condition $condition condition
     *
     * @return self
     */
    public function setCondition($condition)
    {
        $this->container['condition'] = $condition;

        return $this;
    }

    /**
     * Gets consequence.
     *
     * @return null|Consequence
     */
    public function getConsequence()
    {
        return $this->container['consequence'] ?? null;
    }

    /**
     * Sets consequence.
     *
     * @param null|Consequence $consequence consequence
     *
     * @return self
     */
    public function setConsequence($consequence)
    {
        $this->container['consequence'] = $consequence;

        return $this;
    }

    /**
     * Gets description.
     *
     * @return null|string
     */
    public function getDescription()
    {
        return $this->container['description'] ?? null;
    }

    /**
     * Sets description.
     *
     * @param null|string $description Description of the rule's purpose. This can be helpful for display in the Algolia dashboard.
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets enabled.
     *
     * @return null|bool
     */
    public function getEnabled()
    {
        return $this->container['enabled'] ?? null;
    }

    /**
     * Sets enabled.
     *
     * @param null|bool $enabled Indicates whether to enable the rule. If it isn't enabled, it isn't applied at query time.
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

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
     * @return null|mixed
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param null|int $offset Offset
     * @param mixed    $value  Value to be set
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
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}
