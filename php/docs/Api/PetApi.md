# OpenAPI\Client\PetApi

All URIs are relative to http://petstore.swagger.io/v2, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**add()**](PetApi.md#add) | **POST** /pet | Add a new pet to the store |
| [**delete()**](PetApi.md#delete) | **DELETE** /pet/{petId} | Deletes a pet |
| [**findByStatus()**](PetApi.md#findByStatus) | **GET** /pet/findByStatus | Finds Pets by status |
| [**findByTags()**](PetApi.md#findByTags) | **GET** /pet/findByTags | Finds Pets by tags |
| [**getById()**](PetApi.md#getById) | **GET** /pet/{petId} | Find pet by ID |
| [**update()**](PetApi.md#update) | **PUT** /pet | Update an existing pet |
| [**updateWithForm()**](PetApi.md#updateWithForm) | **POST** /pet/{petId} | Updates a pet in the store with form data |
| [**uploadImage()**](PetApi.md#uploadImage) | **POST** /pet/{petId}/uploadImage | uploads an image |


## `add()`

```php
add($pet): \OpenAPI\Client\Model\Pet
```

Add a new pet to the store



### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure OAuth2 access token for authorization: petstore_auth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$pet = new \OpenAPI\Client\Model\Pet(); // \OpenAPI\Client\Model\Pet | Pet object that needs to be added to the store

try {
    $result = $apiInstance->add($pet);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->add: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **pet** | [**\OpenAPI\Client\Model\Pet**](../Model/Pet.md)| Pet object that needs to be added to the store | |

### Return type

[**\OpenAPI\Client\Model\Pet**](../Model/Pet.md)

### Authorization

[petstore_auth](../../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `application/json`, `application/xml`
- **Accept**: `application/xml`, `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `delete()`

```php
delete($pet_id)
```

Deletes a pet



### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: api_key
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('api_key', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('api_key', 'Bearer');

// Configure OAuth2 access token for authorization: petstore_auth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$pet_id = 56; // int | Pet id to delete

try {
    $apiInstance->delete($pet_id);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->delete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **pet_id** | **int**| Pet id to delete | |

### Return type

void (empty response body)

### Authorization

[api_key](../../README.md#api_key), [petstore_auth](../../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `findByStatus()`

```php
findByStatus($status): \OpenAPI\Client\Model\Pet[]
```

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure OAuth2 access token for authorization: petstore_auth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$status = array('status_example'); // string[] | Status values that need to be considered for filter

try {
    $result = $apiInstance->findByStatus($status);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->findByStatus: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **status** | [**string[]**](../Model/string.md)| Status values that need to be considered for filter | |

### Return type

[**\OpenAPI\Client\Model\Pet[]**](../Model/Pet.md)

### Authorization

[petstore_auth](../../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/xml`, `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `findByTags()`

```php
findByTags($tags): \OpenAPI\Client\Model\Pet[]
```

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure OAuth2 access token for authorization: petstore_auth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$tags = array('tags_example'); // string[] | Tags to filter by

try {
    $result = $apiInstance->findByTags($tags);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->findByTags: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **tags** | [**string[]**](../Model/string.md)| Tags to filter by | |

### Return type

[**\OpenAPI\Client\Model\Pet[]**](../Model/Pet.md)

### Authorization

[petstore_auth](../../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/xml`, `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getById()`

```php
getById($pet_id): \OpenAPI\Client\Model\Pet
```

Find pet by ID

Returns a single pet

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: api_key
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('api_key', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('api_key', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$pet_id = 56; // int | ID of pet to return

try {
    $result = $apiInstance->getById($pet_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->getById: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **pet_id** | **int**| ID of pet to return | |

### Return type

[**\OpenAPI\Client\Model\Pet**](../Model/Pet.md)

### Authorization

[api_key](../../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/xml`, `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `update()`

```php
update($pet): \OpenAPI\Client\Model\Pet
```

Update an existing pet



### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure OAuth2 access token for authorization: petstore_auth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$pet = new \OpenAPI\Client\Model\Pet(); // \OpenAPI\Client\Model\Pet | Pet object that needs to be added to the store

try {
    $result = $apiInstance->update($pet);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->update: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **pet** | [**\OpenAPI\Client\Model\Pet**](../Model/Pet.md)| Pet object that needs to be added to the store | |

### Return type

[**\OpenAPI\Client\Model\Pet**](../Model/Pet.md)

### Authorization

[petstore_auth](../../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `application/json`, `application/xml`
- **Accept**: `application/xml`, `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateWithForm()`

```php
updateWithForm($pet_id, $name, $status)
```

Updates a pet in the store with form data



### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure OAuth2 access token for authorization: petstore_auth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$pet_id = 56; // int | ID of pet that needs to be updated
$name = 'name_example'; // string | Updated name of the pet
$status = 'status_example'; // string | Updated status of the pet

try {
    $apiInstance->updateWithForm($pet_id, $name, $status);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->updateWithForm: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **pet_id** | **int**| ID of pet that needs to be updated | |
| **name** | **string**| Updated name of the pet | [optional] |
| **status** | **string**| Updated status of the pet | [optional] |

### Return type

void (empty response body)

### Authorization

[petstore_auth](../../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `application/x-www-form-urlencoded`
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `uploadImage()`

```php
uploadImage($pet_id, $additional_metadata, $file): \OpenAPI\Client\Model\ApiResponse
```

uploads an image



### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure OAuth2 access token for authorization: petstore_auth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PetApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$pet_id = 56; // int | ID of pet to update
$additional_metadata = 'additional_metadata_example'; // string | Additional data to pass to server
$file = "/path/to/file.txt"; // \SplFileObject | file to upload

try {
    $result = $apiInstance->uploadImage($pet_id, $additional_metadata, $file);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PetApi->uploadImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **pet_id** | **int**| ID of pet to update | |
| **additional_metadata** | **string**| Additional data to pass to server | [optional] |
| **file** | **\SplFileObject****\SplFileObject**| file to upload | [optional] |

### Return type

[**\OpenAPI\Client\Model\ApiResponse**](../Model/ApiResponse.md)

### Authorization

[petstore_auth](../../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
