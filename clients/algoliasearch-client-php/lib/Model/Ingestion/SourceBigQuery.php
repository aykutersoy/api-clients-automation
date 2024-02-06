<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Ingestion;

/**
 * SourceBigQuery Class Doc Comment.
 *
 * @category Class
 */
class SourceBigQuery extends \Algolia\AlgoliaSearch\Model\AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'projectID' => 'string',
        'datasetID' => 'string',
        'dataType' => '\Algolia\AlgoliaSearch\Model\Ingestion\BigQueryDataType',
        'table' => 'string',
        'tablePrefix' => 'string',
        'customSQLRequest' => 'string',
        'uniqueIDColumn' => 'string',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'projectID' => null,
        'datasetID' => null,
        'dataType' => null,
        'table' => null,
        'tablePrefix' => null,
        'customSQLRequest' => null,
        'uniqueIDColumn' => null,
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'projectID' => 'projectID',
        'datasetID' => 'datasetID',
        'dataType' => 'dataType',
        'table' => 'table',
        'tablePrefix' => 'tablePrefix',
        'customSQLRequest' => 'customSQLRequest',
        'uniqueIDColumn' => 'uniqueIDColumn',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'projectID' => 'setProjectID',
        'datasetID' => 'setDatasetID',
        'dataType' => 'setDataType',
        'table' => 'setTable',
        'tablePrefix' => 'setTablePrefix',
        'customSQLRequest' => 'setCustomSQLRequest',
        'uniqueIDColumn' => 'setUniqueIDColumn',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'projectID' => 'getProjectID',
        'datasetID' => 'getDatasetID',
        'dataType' => 'getDataType',
        'table' => 'getTable',
        'tablePrefix' => 'getTablePrefix',
        'customSQLRequest' => 'getCustomSQLRequest',
        'uniqueIDColumn' => 'getUniqueIDColumn',
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
        if (isset($data['projectID'])) {
            $this->container['projectID'] = $data['projectID'];
        }
        if (isset($data['datasetID'])) {
            $this->container['datasetID'] = $data['datasetID'];
        }
        if (isset($data['dataType'])) {
            $this->container['dataType'] = $data['dataType'];
        }
        if (isset($data['table'])) {
            $this->container['table'] = $data['table'];
        }
        if (isset($data['tablePrefix'])) {
            $this->container['tablePrefix'] = $data['tablePrefix'];
        }
        if (isset($data['customSQLRequest'])) {
            $this->container['customSQLRequest'] = $data['customSQLRequest'];
        }
        if (isset($data['uniqueIDColumn'])) {
            $this->container['uniqueIDColumn'] = $data['uniqueIDColumn'];
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

        if (!isset($this->container['projectID']) || null === $this->container['projectID']) {
            $invalidProperties[] = "'projectID' can't be null";
        }
        if (!isset($this->container['datasetID']) || null === $this->container['datasetID']) {
            $invalidProperties[] = "'datasetID' can't be null";
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
     * Gets projectID.
     *
     * @return string
     */
    public function getProjectID()
    {
        return $this->container['projectID'] ?? null;
    }

    /**
     * Sets projectID.
     *
     * @param string $projectID project ID of the BigQuery Source
     *
     * @return self
     */
    public function setProjectID($projectID)
    {
        $this->container['projectID'] = $projectID;

        return $this;
    }

    /**
     * Gets datasetID.
     *
     * @return string
     */
    public function getDatasetID()
    {
        return $this->container['datasetID'] ?? null;
    }

    /**
     * Sets datasetID.
     *
     * @param string $datasetID dataset ID of the BigQuery Source
     *
     * @return self
     */
    public function setDatasetID($datasetID)
    {
        $this->container['datasetID'] = $datasetID;

        return $this;
    }

    /**
     * Gets dataType.
     *
     * @return null|\Algolia\AlgoliaSearch\Model\Ingestion\BigQueryDataType
     */
    public function getDataType()
    {
        return $this->container['dataType'] ?? null;
    }

    /**
     * Sets dataType.
     *
     * @param null|\Algolia\AlgoliaSearch\Model\Ingestion\BigQueryDataType $dataType dataType
     *
     * @return self
     */
    public function setDataType($dataType)
    {
        $this->container['dataType'] = $dataType;

        return $this;
    }

    /**
     * Gets table.
     *
     * @return null|string
     */
    public function getTable()
    {
        return $this->container['table'] ?? null;
    }

    /**
     * Sets table.
     *
     * @param null|string $table table name (for default BQ)
     *
     * @return self
     */
    public function setTable($table)
    {
        $this->container['table'] = $table;

        return $this;
    }

    /**
     * Gets tablePrefix.
     *
     * @return null|string
     */
    public function getTablePrefix()
    {
        return $this->container['tablePrefix'] ?? null;
    }

    /**
     * Sets tablePrefix.
     *
     * @param null|string $tablePrefix table prefix (for Google Analytics)
     *
     * @return self
     */
    public function setTablePrefix($tablePrefix)
    {
        $this->container['tablePrefix'] = $tablePrefix;

        return $this;
    }

    /**
     * Gets customSQLRequest.
     *
     * @return null|string
     */
    public function getCustomSQLRequest()
    {
        return $this->container['customSQLRequest'] ?? null;
    }

    /**
     * Sets customSQLRequest.
     *
     * @param null|string $customSQLRequest custom SQL request to extract data from the BigQuery table
     *
     * @return self
     */
    public function setCustomSQLRequest($customSQLRequest)
    {
        $this->container['customSQLRequest'] = $customSQLRequest;

        return $this;
    }

    /**
     * Gets uniqueIDColumn.
     *
     * @return null|string
     */
    public function getUniqueIDColumn()
    {
        return $this->container['uniqueIDColumn'] ?? null;
    }

    /**
     * Sets uniqueIDColumn.
     *
     * @param null|string $uniqueIDColumn the name of the column that contains the unique ID, used as `objectID` in Algolia
     *
     * @return self
     */
    public function setUniqueIDColumn($uniqueIDColumn)
    {
        $this->container['uniqueIDColumn'] = $uniqueIDColumn;

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
