<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Api;

use Algolia\AlgoliaSearch\Algolia;
use Algolia\AlgoliaSearch\Configuration\MonitoringConfig;
use Algolia\AlgoliaSearch\ObjectSerializer;
use Algolia\AlgoliaSearch\RetryStrategy\ApiWrapper;
use Algolia\AlgoliaSearch\RetryStrategy\ApiWrapperInterface;
use Algolia\AlgoliaSearch\RetryStrategy\ClusterHosts;

/**
 * MonitoringClient Class Doc Comment
 *
 * @category Class
 * @package  Algolia\AlgoliaSearch
 */
class MonitoringClient
{
    /**
     * @var ApiWrapperInterface
     */
    protected $api;

    /**
     * @var MonitoringConfig
     */
    protected $config;

    /**
     * @param MonitoringConfig $config
     * @param ApiWrapperInterface $apiWrapper
     */
    public function __construct(ApiWrapperInterface $apiWrapper, MonitoringConfig $config)
    {
        $this->config = $config;
        $this->api = $apiWrapper;
    }

    /**
     * Instantiate the client with basic credentials
     *
     * @param string $appId  Application ID
     * @param string $apiKey Algolia API Key
     */
    public static function create($appId = null, $apiKey = null)
    {
        return static::createWithConfig(MonitoringConfig::create($appId, $apiKey));
    }

    /**
     * Instantiate the client with configuration
     *
     * @param MonitoringConfig $config Configuration
     */
    public static function createWithConfig(MonitoringConfig $config)
    {
        $config = clone $config;

        $apiWrapper = new ApiWrapper(
            Algolia::getHttpClient(),
            $config,
            self::getClusterHosts($config)
        );

        return new static($apiWrapper, $config);
    }

    /**
     * Gets the cluster hosts depending on the config
     *
     * @param MonitoringConfig $config
     *
     * @return ClusterHosts
     */
    public static function getClusterHosts(MonitoringConfig $config)
    {

        if ($hosts = $config->getHosts()) {
            // If a list of hosts was passed, we ignore the cache
            $clusterHosts = ClusterHosts::create($hosts);
        } else {
            $url = $config->getRegion() !== null && $config->getRegion() !== '' ?
                str_replace('{region}', $config->getRegion(), '') :
                '';
            $clusterHosts = ClusterHosts::create($url);
        }

        return $clusterHosts;
    }

    /**
     * @return MonitoringConfig
     */
    public function getClientConfig()
    {
        return $this->config;
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function del($path, $parameters = null, $requestOptions = [])
    {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `del`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace(
                '{path}',
                $path,
                $resourcePath
            );
        }

        return $this->sendRequest('DELETE', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function get($path, $parameters = null, $requestOptions = [])
    {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `get`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace(
                '{path}',
                $path,
                $resourcePath
            );
        }

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * List incidents for selected clusters.
     *
     * @param string $clusters Subset of clusters, separated by comma. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\IncidentsResponse
     */
    public function getClusterIncidents($clusters, $requestOptions = [])
    {
        // verify the required parameter 'clusters' is set
        if (!isset($clusters)) {
            throw new \InvalidArgumentException(
                'Parameter `clusters` is required when calling `getClusterIncidents`.'
            );
        }

        $resourcePath = '/1/incidents/{clusters}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($clusters !== null) {
            $resourcePath = str_replace(
                '{clusters}',
                ObjectSerializer::toPathValue($clusters),
                $resourcePath
            );
        }

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * List statuses of selected clusters.
     *
     * @param string $clusters Subset of clusters, separated by comma. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\StatusResponse
     */
    public function getClusterStatus($clusters, $requestOptions = [])
    {
        // verify the required parameter 'clusters' is set
        if (!isset($clusters)) {
            throw new \InvalidArgumentException(
                'Parameter `clusters` is required when calling `getClusterStatus`.'
            );
        }

        $resourcePath = '/1/status/{clusters}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($clusters !== null) {
            $resourcePath = str_replace(
                '{clusters}',
                ObjectSerializer::toPathValue($clusters),
                $resourcePath
            );
        }

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * List incidents.
     *
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\IncidentsResponse
     */
    public function getIncidents($requestOptions = [])
    {

        $resourcePath = '/1/incidents';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * Get indexing times.
     *
     * @param string $clusters Subset of clusters, separated by comma. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\IndexingTimeResponse
     */
    public function getIndexingTime($clusters, $requestOptions = [])
    {
        // verify the required parameter 'clusters' is set
        if (!isset($clusters)) {
            throw new \InvalidArgumentException(
                'Parameter `clusters` is required when calling `getIndexingTime`.'
            );
        }

        $resourcePath = '/1/indexing/{clusters}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($clusters !== null) {
            $resourcePath = str_replace(
                '{clusters}',
                ObjectSerializer::toPathValue($clusters),
                $resourcePath
            );
        }

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * List servers.
     *
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\InventoryResponse
     */
    public function getInventory($requestOptions = [])
    {

        $resourcePath = '/1/inventory/servers';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * Get search latency times.
     *
     * @param string $clusters Subset of clusters, separated by comma. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\LatencyResponse
     */
    public function getLatency($clusters, $requestOptions = [])
    {
        // verify the required parameter 'clusters' is set
        if (!isset($clusters)) {
            throw new \InvalidArgumentException(
                'Parameter `clusters` is required when calling `getLatency`.'
            );
        }

        $resourcePath = '/1/latency/{clusters}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($clusters !== null) {
            $resourcePath = str_replace(
                '{clusters}',
                ObjectSerializer::toPathValue($clusters),
                $resourcePath
            );
        }

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * Get metrics for a given period.
     *
     * @param array $metric Metric to report.  For more information about the individual metrics, see the response. To include all metrics, use &#x60;*&#x60; as the parameter. (required)
     * @param array $period Period over which to aggregate the metrics:  - &#x60;minute&#x60;. Aggregate the last minute. 1 data point per 10 seconds. - &#x60;hour&#x60;. Aggregate the last hour. 1 data point per minute. - &#x60;day&#x60;. Aggregate the last day. 1 data point per 10 minutes. - &#x60;week&#x60;. Aggregate the last week. 1 data point per hour. - &#x60;month&#x60;. Aggregate the last month. 1 data point per day. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\InfrastructureResponse
     */
    public function getMetrics($metric, $period, $requestOptions = [])
    {
        // verify the required parameter 'metric' is set
        if (!isset($metric)) {
            throw new \InvalidArgumentException(
                'Parameter `metric` is required when calling `getMetrics`.'
            );
        }
        // verify the required parameter 'period' is set
        if (!isset($period)) {
            throw new \InvalidArgumentException(
                'Parameter `period` is required when calling `getMetrics`.'
            );
        }

        $resourcePath = '/1/infrastructure/{metric}/period/{period}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($metric !== null) {
            $resourcePath = str_replace(
                '{metric}',
                ObjectSerializer::toPathValue($metric),
                $resourcePath
            );
        }

        // path params
        if ($period !== null) {
            $resourcePath = str_replace(
                '{period}',
                ObjectSerializer::toPathValue($period),
                $resourcePath
            );
        }

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * Test the reachability of clusters.
     *
     * @param string $clusters Subset of clusters, separated by comma. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|array<string,array>
     */
    public function getReachability($clusters, $requestOptions = [])
    {
        // verify the required parameter 'clusters' is set
        if (!isset($clusters)) {
            throw new \InvalidArgumentException(
                'Parameter `clusters` is required when calling `getReachability`.'
            );
        }

        $resourcePath = '/1/reachability/{clusters}/probes';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($clusters !== null) {
            $resourcePath = str_replace(
                '{clusters}',
                ObjectSerializer::toPathValue($clusters),
                $resourcePath
            );
        }

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * List cluster statuses.
     *
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Monitoring\StatusResponse
     */
    public function getStatus($requestOptions = [])
    {

        $resourcePath = '/1/status';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        return $this->sendRequest('GET', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $body The parameters to send with the custom request. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function post($path, $parameters = null, $body = null, $requestOptions = [])
    {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `post`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody =  isset($body) ? $body : [];

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace(
                '{path}',
                $path,
                $resourcePath
            );
        }

        return $this->sendRequest('POST', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $body The parameters to send with the custom request. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function put($path, $parameters = null, $body = null, $requestOptions = [])
    {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `put`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody =  isset($body) ? $body : [];

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace(
                '{path}',
                $path,
                $resourcePath
            );
        }

        return $this->sendRequest('PUT', $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, );
    }

    private function sendRequest($method, $resourcePath, $headers, $queryParameters, $httpBody, $requestOptions, $useReadTransporter = false)
    {
        if (!isset($requestOptions['headers'])) {
            $requestOptions['headers'] = [];
        }
        if (!isset($requestOptions['queryParameters'])) {
            $requestOptions['queryParameters'] = [];
        }

        $requestOptions['headers'] = array_merge($headers, $requestOptions['headers']);
        $requestOptions['queryParameters'] = array_merge($queryParameters, $requestOptions['queryParameters']);
        $query = \GuzzleHttp\Psr7\Query::build($requestOptions['queryParameters']);

        return $this->api->sendRequest(
            $method,
            $resourcePath . ($query ? "?{$query}" : ''),
            $httpBody,
            $requestOptions,
            $useReadTransporter
        );
    }
}
