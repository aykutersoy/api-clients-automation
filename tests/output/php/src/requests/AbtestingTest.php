<?php

namespace Algolia\AlgoliaSearch\Test\Api;

use Algolia\AlgoliaSearch\Api\AbtestingClient;
use Algolia\AlgoliaSearch\Configuration\AbtestingConfig;
use Algolia\AlgoliaSearch\Http\HttpClientInterface;
use Algolia\AlgoliaSearch\Http\Psr7\Response;
use Algolia\AlgoliaSearch\RetryStrategy\ApiWrapper;
use Algolia\AlgoliaSearch\RetryStrategy\ClusterHosts;
use Dotenv\Dotenv;
use GuzzleHttp\Psr7\Query;
use PHPUnit\Framework\TestCase;
use Psr\Http\Message\RequestInterface;

// we only read .env file if we run locally
if (isset($_ENV['DOCKER']) && 'true' === $_ENV['DOCKER']) {
    $dotenv = Dotenv::createImmutable('tests');
    $dotenv->load();
} else {
    $_ENV = getenv();
}

/**
 * AbtestingTest.
 *
 * @category Class
 *
 * @internal
 * @coversNothing
 */
class AbtestingTest extends TestCase implements HttpClientInterface
{
    /**
     * @var RequestInterface[]
     */
    private $recordedRequests = [];

    public function sendRequest(RequestInterface $request, $timeout, $connectTimeout)
    {
        $this->recordedRequests[] = $request;

        return new Response(200, [], '{}');
    }

    /**
     * Test case for AddABTests
     * addABTests with minimal parameters.
     */
    public function testAddABTests0()
    {
        $client = $this->getClient();
        $client->addABTests(
            ['endAt' => '2022-12-31T00:00:00.000Z',
                'name' => 'myABTest',
                'variants' => [
                    ['index' => 'AB_TEST_1',
                        'trafficPercentage' => 30,
                    ],

                    ['index' => 'AB_TEST_2',
                        'trafficPercentage' => 50,
                    ],
                ],
            ],
        );

        $this->assertRequests([
            [
                'path' => '/2/abtests',
                'method' => 'POST',
                'body' => json_decode('{"endAt":"2022-12-31T00:00:00.000Z","name":"myABTest","variants":[{"index":"AB_TEST_1","trafficPercentage":30},{"index":"AB_TEST_2","trafficPercentage":50}]}'),
            ],
        ]);
    }

    /**
     * Test case for CustomDelete
     * allow del method for a custom path with minimal parameters.
     */
    public function testCustomDelete0()
    {
        $client = $this->getClient();
        $client->customDelete(
            '/test/minimal',
        );

        $this->assertRequests([
            [
                'path' => '/1/test/minimal',
                'method' => 'DELETE',
                'body' => null,
            ],
        ]);
    }

    /**
     * Test case for CustomDelete
     * allow del method for a custom path with all parameters.
     */
    public function testCustomDelete1()
    {
        $client = $this->getClient();
        $client->customDelete(
            '/test/all',
            ['query' => 'parameters',
            ],
        );

        $this->assertRequests([
            [
                'path' => '/1/test/all',
                'method' => 'DELETE',
                'body' => null,
                'queryParameters' => json_decode('{"query":"parameters"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomGet
     * allow get method for a custom path with minimal parameters.
     */
    public function testCustomGet0()
    {
        $client = $this->getClient();
        $client->customGet(
            '/test/minimal',
        );

        $this->assertRequests([
            [
                'path' => '/1/test/minimal',
                'method' => 'GET',
                'body' => null,
            ],
        ]);
    }

    /**
     * Test case for CustomGet
     * allow get method for a custom path with all parameters.
     */
    public function testCustomGet1()
    {
        $client = $this->getClient();
        $client->customGet(
            '/test/all',
            ['query' => 'parameters with space',
            ],
        );

        $this->assertRequests([
            [
                'path' => '/1/test/all',
                'method' => 'GET',
                'body' => null,
                'queryParameters' => json_decode('{"query":"parameters%20with%20space"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * allow post method for a custom path with minimal parameters.
     */
    public function testCustomPost0()
    {
        $client = $this->getClient();
        $client->customPost(
            '/test/minimal',
        );

        $this->assertRequests([
            [
                'path' => '/1/test/minimal',
                'method' => 'POST',
                'body' => json_decode('{}'),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * allow post method for a custom path with all parameters.
     */
    public function testCustomPost1()
    {
        $client = $this->getClient();
        $client->customPost(
            '/test/all',
            ['query' => 'parameters',
            ],
            ['body' => 'parameters',
            ],
        );

        $this->assertRequests([
            [
                'path' => '/1/test/all',
                'method' => 'POST',
                'body' => json_decode('{"body":"parameters"}'),
                'queryParameters' => json_decode('{"query":"parameters"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions can override default query parameters.
     */
    public function testCustomPost2()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
                'query' => 'myQueryParameter',
            ],
            'headers' => [
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"myQueryParameter"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions merges query parameters with default ones.
     */
    public function testCustomPost3()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
                'query2' => 'myQueryParameter',
            ],
            'headers' => [
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters","query2":"myQueryParameter"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions can override default headers.
     */
    public function testCustomPost4()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
            ],
            'headers' => [
                'x-algolia-api-key' => 'myApiKey',
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters"}', true),
                'headers' => json_decode('{"x-algolia-api-key":"myApiKey"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions merges headers with default ones.
     */
    public function testCustomPost5()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
            ],
            'headers' => [
                'x-algolia-api-key' => 'myApiKey',
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters"}', true),
                'headers' => json_decode('{"x-algolia-api-key":"myApiKey"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions queryParameters accepts booleans.
     */
    public function testCustomPost6()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
                'isItWorking' => true,
            ],
            'headers' => [
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters","isItWorking":"true"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions queryParameters accepts integers.
     */
    public function testCustomPost7()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
                'myParam' => 2,
            ],
            'headers' => [
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters","myParam":"2"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions queryParameters accepts list of string.
     */
    public function testCustomPost8()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
                'myParam' => ['c',  'd',
                ],
            ],
            'headers' => [
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters","myParam":"c%2Cd"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions queryParameters accepts list of booleans.
     */
    public function testCustomPost9()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
                'myParam' => [true,  true,  false,
                ],
            ],
            'headers' => [
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters","myParam":"true%2Ctrue%2Cfalse"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPost
     * requestOptions queryParameters accepts list of integers.
     */
    public function testCustomPost10()
    {
        $client = $this->getClient();
        $requestOptions = [
            'queryParameters' => [
                'myParam' => [1,  2,
                ],
            ],
            'headers' => [
            ],
        ];
        $client->customPost(
            '/test/requestOptions',
            ['query' => 'parameters',
            ],
            ['facet' => 'filters',
            ],
            $requestOptions
        );

        $this->assertRequests([
            [
                'path' => '/1/test/requestOptions',
                'method' => 'POST',
                'body' => json_decode('{"facet":"filters"}'),
                'queryParameters' => json_decode('{"query":"parameters","myParam":"1%2C2"}', true),
            ],
        ]);
    }

    /**
     * Test case for CustomPut
     * allow put method for a custom path with minimal parameters.
     */
    public function testCustomPut0()
    {
        $client = $this->getClient();
        $client->customPut(
            '/test/minimal',
        );

        $this->assertRequests([
            [
                'path' => '/1/test/minimal',
                'method' => 'PUT',
                'body' => json_decode('{}'),
            ],
        ]);
    }

    /**
     * Test case for CustomPut
     * allow put method for a custom path with all parameters.
     */
    public function testCustomPut1()
    {
        $client = $this->getClient();
        $client->customPut(
            '/test/all',
            ['query' => 'parameters',
            ],
            ['body' => 'parameters',
            ],
        );

        $this->assertRequests([
            [
                'path' => '/1/test/all',
                'method' => 'PUT',
                'body' => json_decode('{"body":"parameters"}'),
                'queryParameters' => json_decode('{"query":"parameters"}', true),
            ],
        ]);
    }

    /**
     * Test case for DeleteABTest
     * deleteABTest.
     */
    public function testDeleteABTest0()
    {
        $client = $this->getClient();
        $client->deleteABTest(
            42,
        );

        $this->assertRequests([
            [
                'path' => '/2/abtests/42',
                'method' => 'DELETE',
                'body' => null,
            ],
        ]);
    }

    /**
     * Test case for GetABTest
     * getABTest.
     */
    public function testGetABTest0()
    {
        $client = $this->getClient();
        $client->getABTest(
            42,
        );

        $this->assertRequests([
            [
                'path' => '/2/abtests/42',
                'method' => 'GET',
                'body' => null,
            ],
        ]);
    }

    /**
     * Test case for ListABTests
     * listABTests with minimal parameters.
     */
    public function testListABTests0()
    {
        $client = $this->getClient();
        $client->listABTests();

        $this->assertRequests([
            [
                'path' => '/2/abtests',
                'method' => 'GET',
                'body' => null,
            ],
        ]);
    }

    /**
     * Test case for ListABTests
     * listABTests with parameters.
     */
    public function testListABTests1()
    {
        $client = $this->getClient();
        $client->listABTests(
            0,
            21,
            'cts_e2e ab',
            't',
        );

        $this->assertRequests([
            [
                'path' => '/2/abtests',
                'method' => 'GET',
                'body' => null,
                'queryParameters' => json_decode('{"offset":"0","limit":"21","indexPrefix":"cts_e2e%20ab","indexSuffix":"t"}', true),
            ],
        ]);

        $e2eClient = $this->getE2EClient();
        $resp = $e2eClient->listABTests(
            0,
            21,
            'cts_e2e ab',
            't',
        );

        $expected = json_decode('{"abtests":[{"abTestID":84617,"createdAt":"2024-02-06T10:04:30.209477Z","endAt":"2024-05-06T09:04:26.469Z","name":"cts_e2e_abtest","status":"active","variants":[{"addToCartCount":0,"clickCount":0,"conversionCount":0,"description":"","index":"cts_e2e_search_facet","purchaseCount":0,"trafficPercentage":25},{"addToCartCount":0,"clickCount":0,"conversionCount":0,"description":"","index":"cts_e2e abtest","purchaseCount":0,"trafficPercentage":75}]}],"count":1,"total":1}', true);

        $this->assertEquals($this->union($expected, $resp), $expected);
    }

    /**
     * Test case for StopABTest
     * stopABTest.
     */
    public function testStopABTest0()
    {
        $client = $this->getClient();
        $client->stopABTest(
            42,
        );

        $this->assertRequests([
            [
                'path' => '/2/abtests/42/stop',
                'method' => 'POST',
                'body' => json_decode(''),
            ],
        ]);
    }

    protected function union($expected, $received)
    {
        $res = [];

        foreach ($expected as $k => $v) {
            if (isset($received[$k])) {
                if (is_array($v)) {
                    $res[$k] = $this->union($v, $received[$k]);
                } elseif (is_array($v)) {
                    if (!isset($res[$k])) {
                        $res[$k] = [];
                    }

                    foreach ($v as $iv => $v) {
                        $res[$k][] = $this->union($v, $received[$k][$iv]);
                    }
                } else {
                    $res[$k] = $received[$k];
                }
            }
        }

        return $res;
    }

    protected function assertRequests(array $requests)
    {
        $this->assertGreaterThan(0, count($requests));
        $this->assertEquals(count($requests), count($this->recordedRequests));

        foreach ($requests as $i => $request) {
            $recordedRequest = $this->recordedRequests[$i];

            $this->assertEquals($request['method'], $recordedRequest->getMethod());

            $this->assertEquals($request['path'], $recordedRequest->getUri()->getPath());

            if (isset($request['body'])) {
                $this->assertEquals(
                    json_encode($request['body']),
                    $recordedRequest->getBody()->getContents()
                );
            }

            if (isset($request['queryParameters'])) {
                $this->assertEquals(
                    Query::build($request['queryParameters'], false),
                    $recordedRequest->getUri()->getQuery()
                );
            }

            if (isset($request['headers'])) {
                foreach ($request['headers'] as $key => $value) {
                    $this->assertArrayHasKey(
                        $key,
                        $recordedRequest->getHeaders()
                    );
                    $this->assertEquals(
                        $recordedRequest->getHeaderLine($key),
                        $value
                    );
                }
            }
        }
    }

    protected function getE2EClient()
    {
        return AbtestingClient::create($_ENV['ALGOLIA_APPLICATION_ID'], $_ENV['ALGOLIA_ADMIN_KEY']);
    }

    protected function getClient()
    {
        $api = new ApiWrapper($this, AbtestingConfig::create(getenv('ALGOLIA_APP_ID'), getenv('ALGOLIA_API_KEY')), ClusterHosts::create('127.0.0.1'));
        $config = AbtestingConfig::create('foo', 'bar');

        return new AbtestingClient($api, $config);
    }
}
