# acme-php

This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.


## Installation & Usage

### Requirements

PHP 7.4 and later.
Should also work with PHP 8.0.

### Composer

To install the bindings via [Composer](https://getcomposer.org/), add the following to `composer.json`:

```json
{
  "repositories": [
    {
      "type": "vcs",
      "url": "https://github.com/konfig-dev/acme-sdks/tree/main/php.git"
    }
  ],
  "require": {
    "konfig-dev/acme-sdks/tree/main/php": "*@dev"
  }
}
```

Then run `composer install`

### Manual Installation

Download the files and include `autoload.php`:

```php
<?php
require_once('/path/to/acme-php/vendor/autoload.php');
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

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

## API Endpoints

All URIs are relative to *http://petstore.swagger.io/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MiscellaneousApi* | [**paginate**](docs/Api/MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox
*PetApi* | [**add**](docs/Api/PetApi.md#add) | **POST** /pet | Add a new pet to the store
*PetApi* | [**delete**](docs/Api/PetApi.md#delete) | **DELETE** /pet/{petId} | Deletes a pet
*PetApi* | [**findByStatus**](docs/Api/PetApi.md#findbystatus) | **GET** /pet/findByStatus | Finds Pets by status
*PetApi* | [**findByTags**](docs/Api/PetApi.md#findbytags) | **GET** /pet/findByTags | Finds Pets by tags
*PetApi* | [**getById**](docs/Api/PetApi.md#getbyid) | **GET** /pet/{petId} | Find pet by ID
*PetApi* | [**update**](docs/Api/PetApi.md#update) | **PUT** /pet | Update an existing pet
*PetApi* | [**updateWithForm**](docs/Api/PetApi.md#updatewithform) | **POST** /pet/{petId} | Updates a pet in the store with form data
*PetApi* | [**uploadImage**](docs/Api/PetApi.md#uploadimage) | **POST** /pet/{petId}/uploadImage | uploads an image
*StoreApi* | [**deleteOrder**](docs/Api/StoreApi.md#deleteorder) | **DELETE** /store/order/{orderId} | Delete purchase order by ID
*StoreApi* | [**getInventory**](docs/Api/StoreApi.md#getinventory) | **GET** /store/inventory | Returns pet inventories by status
*StoreApi* | [**getOrderById**](docs/Api/StoreApi.md#getorderbyid) | **GET** /store/order/{orderId} | Find purchase order by ID
*StoreApi* | [**placeOrder**](docs/Api/StoreApi.md#placeorder) | **POST** /store/order | Place an order for a pet
*UserApi* | [**create**](docs/Api/UserApi.md#create) | **POST** /user | Create user
*UserApi* | [**createWithArray**](docs/Api/UserApi.md#createwitharray) | **POST** /user/createWithArray | Creates list of users with given input array
*UserApi* | [**createWithList**](docs/Api/UserApi.md#createwithlist) | **POST** /user/createWithList | Creates list of users with given input array
*UserApi* | [**delete**](docs/Api/UserApi.md#delete) | **DELETE** /user/{username} | Delete user
*UserApi* | [**getByName**](docs/Api/UserApi.md#getbyname) | **GET** /user/{username} | Get user by user name
*UserApi* | [**login**](docs/Api/UserApi.md#login) | **GET** /user/login | Logs user into the system
*UserApi* | [**logout**](docs/Api/UserApi.md#logout) | **GET** /user/logout | Logs out current logged in user session
*UserApi* | [**update**](docs/Api/UserApi.md#update) | **PUT** /user/{username} | Updated user

## Models

- [ApiResponse](docs/Model/ApiResponse.md)
- [Category](docs/Model/Category.md)
- [Order](docs/Model/Order.md)
- [PaginateRequest](docs/Model/PaginateRequest.md)
- [PaginateResponse](docs/Model/PaginateResponse.md)
- [PaginateResponseEdges](docs/Model/PaginateResponseEdges.md)
- [PaginateResponseEdgesNode](docs/Model/PaginateResponseEdgesNode.md)
- [PaginateResponsePageInfo](docs/Model/PaginateResponsePageInfo.md)
- [Pet](docs/Model/Pet.md)
- [Tag](docs/Model/Tag.md)
- [User](docs/Model/User.md)

## Authorization

### api_key

- **Type**: API key
- **API key parameter name**: api_key
- **Location**: HTTP header



### petstore_auth

- **Type**: `OAuth`
- **Flow**: `implicit`
- **Authorization URL**: `http://petstore.swagger.io/api/oauth/dialog`
- **Scopes**: 
    - **write:pets**: modify pets in your account
    - **read:pets**: read your pets

## Tests

To run the tests, use:

```bash
composer install
vendor/bin/phpunit
```

## Author



## About this package

This PHP package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: `1.0.0`
    - Package version: `1.0.0`
- Build package: `org.openapitools.codegen.languages.PhpClientCodegen`
