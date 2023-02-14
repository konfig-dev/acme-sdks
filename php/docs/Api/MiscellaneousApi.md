# OpenAPI\Client\MiscellaneousApi

All URIs are relative to http://petstore.swagger.io/v2, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**paginate()**](MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox |


## `paginate()`

```php
paginate($paginate_request): \OpenAPI\Client\Model\PaginateResponse
```

Pagination sandbox

Iterate through a bunch of items

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\MiscellaneousApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$paginate_request = new \OpenAPI\Client\Model\PaginateRequest(); // \OpenAPI\Client\Model\PaginateRequest

try {
    $result = $apiInstance->paginate($paginate_request);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MiscellaneousApi->paginate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **paginate_request** | [**\OpenAPI\Client\Model\PaginateRequest**](../Model/PaginateRequest.md)|  | [optional] |

### Return type

[**\OpenAPI\Client\Model\PaginateResponse**](../Model/PaginateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
