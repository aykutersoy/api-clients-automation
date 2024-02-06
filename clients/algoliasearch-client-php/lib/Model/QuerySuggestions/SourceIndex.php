<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\QuerySuggestions;

/**
 * SourceIndex Class Doc Comment.
 *
 * @category Class
 * @description Configuration of an Algolia index for Query Suggestions.
 */
class SourceIndex extends \Algolia\AlgoliaSearch\Model\AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'indexName' => 'string',
        'replicas' => 'bool',
        'analyticsTags' => 'string[]',
        'facets' => '\Algolia\AlgoliaSearch\Model\QuerySuggestions\Facet[]',
        'minHits' => 'int',
        'minLetters' => 'int',
        'generate' => 'string[][]',
        'external' => 'string[]',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'indexName' => null,
        'replicas' => null,
        'analyticsTags' => null,
        'facets' => null,
        'minHits' => null,
        'minLetters' => null,
        'generate' => null,
        'external' => null,
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'indexName' => 'indexName',
        'replicas' => 'replicas',
        'analyticsTags' => 'analyticsTags',
        'facets' => 'facets',
        'minHits' => 'minHits',
        'minLetters' => 'minLetters',
        'generate' => 'generate',
        'external' => 'external',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'indexName' => 'setIndexName',
        'replicas' => 'setReplicas',
        'analyticsTags' => 'setAnalyticsTags',
        'facets' => 'setFacets',
        'minHits' => 'setMinHits',
        'minLetters' => 'setMinLetters',
        'generate' => 'setGenerate',
        'external' => 'setExternal',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'indexName' => 'getIndexName',
        'replicas' => 'getReplicas',
        'analyticsTags' => 'getAnalyticsTags',
        'facets' => 'getFacets',
        'minHits' => 'getMinHits',
        'minLetters' => 'getMinLetters',
        'generate' => 'getGenerate',
        'external' => 'getExternal',
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
    public function __construct(array $data = null)
    {
        if (isset($data['indexName'])) {
            $this->container['indexName'] = $data['indexName'];
        }
        if (isset($data['replicas'])) {
            $this->container['replicas'] = $data['replicas'];
        }
        if (isset($data['analyticsTags'])) {
            $this->container['analyticsTags'] = $data['analyticsTags'];
        }
        if (isset($data['facets'])) {
            $this->container['facets'] = $data['facets'];
        }
        if (isset($data['minHits'])) {
            $this->container['minHits'] = $data['minHits'];
        }
        if (isset($data['minLetters'])) {
            $this->container['minLetters'] = $data['minLetters'];
        }
        if (isset($data['generate'])) {
            $this->container['generate'] = $data['generate'];
        }
        if (isset($data['external'])) {
            $this->container['external'] = $data['external'];
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

        if (!isset($this->container['indexName']) || null === $this->container['indexName']) {
            $invalidProperties[] = "'indexName' can't be null";
        }
        if (isset($this->container['minHits']) && ($this->container['minHits'] < 0)) {
            $invalidProperties[] = "invalid value for 'minHits', must be bigger than or equal to 0.";
        }

        if (isset($this->container['minLetters']) && ($this->container['minLetters'] < 0)) {
            $invalidProperties[] = "invalid value for 'minLetters', must be bigger than or equal to 0.";
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
     * Gets indexName.
     *
     * @return string
     */
    public function getIndexName()
    {
        return $this->container['indexName'] ?? null;
    }

    /**
     * Sets indexName.
     *
     * @param string $indexName name of the Algolia index to use as source for query suggestions
     *
     * @return self
     */
    public function setIndexName($indexName)
    {
        $this->container['indexName'] = $indexName;

        return $this;
    }

    /**
     * Gets replicas.
     *
     * @return null|bool
     */
    public function getReplicas()
    {
        return $this->container['replicas'] ?? null;
    }

    /**
     * Sets replicas.
     *
     * @param null|bool $replicas If true, Query Suggestions uses all replicas of the primary index to find popular searches. If false, only the primary index is used.
     *
     * @return self
     */
    public function setReplicas($replicas)
    {
        $this->container['replicas'] = $replicas;

        return $this;
    }

    /**
     * Gets analyticsTags.
     *
     * @return null|string[]
     */
    public function getAnalyticsTags()
    {
        return $this->container['analyticsTags'] ?? null;
    }

    /**
     * Sets analyticsTags.
     *
     * @param null|string[] $analyticsTags [Analytics tags](https://www.algolia.com/doc/api-reference/api-parameters/analyticsTags/) for filtering the popular searches.
     *
     * @return self
     */
    public function setAnalyticsTags($analyticsTags)
    {
        $this->container['analyticsTags'] = $analyticsTags;

        return $this;
    }

    /**
     * Gets facets.
     *
     * @return null|\Algolia\AlgoliaSearch\Model\QuerySuggestions\Facet[]
     */
    public function getFacets()
    {
        return $this->container['facets'] ?? null;
    }

    /**
     * Sets facets.
     *
     * @param null|\Algolia\AlgoliaSearch\Model\QuerySuggestions\Facet[] $facets Facets to use as top categories with your suggestions.  If provided, Query Suggestions adds the top facet values to each suggestion.
     *
     * @return self
     */
    public function setFacets($facets)
    {
        $this->container['facets'] = $facets;

        return $this;
    }

    /**
     * Gets minHits.
     *
     * @return null|int
     */
    public function getMinHits()
    {
        return $this->container['minHits'] ?? null;
    }

    /**
     * Sets minHits.
     *
     * @param null|int $minHits Minimum number of hits required to be included as a suggestion.  A search query must at least generate `minHits` hits to be included in the Query Suggestions index.
     *
     * @return self
     */
    public function setMinHits($minHits)
    {
        if (!is_null($minHits) && ($minHits < 0)) {
            throw new \InvalidArgumentException('invalid value for $minHits when calling SourceIndex., must be bigger than or equal to 0.');
        }

        $this->container['minHits'] = $minHits;

        return $this;
    }

    /**
     * Gets minLetters.
     *
     * @return null|int
     */
    public function getMinLetters()
    {
        return $this->container['minLetters'] ?? null;
    }

    /**
     * Sets minLetters.
     *
     * @param null|int $minLetters Minimum letters required to be included as a suggestion.  A search query must be at least `minLetters` long to be included in the Query Suggestions index.
     *
     * @return self
     */
    public function setMinLetters($minLetters)
    {
        if (!is_null($minLetters) && ($minLetters < 0)) {
            throw new \InvalidArgumentException('invalid value for $minLetters when calling SourceIndex., must be bigger than or equal to 0.');
        }

        $this->container['minLetters'] = $minLetters;

        return $this;
    }

    /**
     * Gets generate.
     *
     * @return null|string[][]
     */
    public function getGenerate()
    {
        return $this->container['generate'] ?? null;
    }

    /**
     * Sets generate.
     *
     * @param null|string[][] $generate generate
     *
     * @return self
     */
    public function setGenerate($generate)
    {
        $this->container['generate'] = $generate;

        return $this;
    }

    /**
     * Gets external.
     *
     * @return null|string[]
     */
    public function getExternal()
    {
        return $this->container['external'] ?? null;
    }

    /**
     * Sets external.
     *
     * @param null|string[] $external Algolia indices with popular searches to use as query suggestions.  Records of these indices must have these attributes:    - `query`: search query which will be added as a suggestion   - `count`: measure of popularity of that search query  For example, you can export popular searches from an external analytics tool, such as Google Analytics or Adobe Analytics, and feed this data into an external Algolia index. You can use this external index to generate query suggestions until your Algolia analytics has collected enough data.
     *
     * @return self
     */
    public function setExternal($external)
    {
        $this->container['external'] = $external;

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
