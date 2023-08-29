# acme-python-sdk@1.0.0
This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.


## Requirements

Python >=3.7

## Installing

```sh
pip install acme-python-sdk==1.0.0
```

## Getting Started

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
)

try:
    # Pagination sandbox
    paginate_response = acme.miscellaneous.paginate(
        first=1,  # optional
        after="string_example",  # optional
    )
    pprint(paginate_response.body)
    pprint(paginate_response.body["edges"])
    pprint(paginate_response.body["page_info"])
    pprint(paginate_response.headers)
    pprint(paginate_response.status)
    pprint(paginate_response.round_trip_time)
except ApiException as e:
    print("Exception when calling MiscellaneousApi.paginate: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```

## Async

`async` support is available by prepending `a` to any method.

```python
import asyncio
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
)


async def main():
    try:
        # Pagination sandbox
        paginate_response = await acme.miscellaneous.apaginate(
            first=1,  # optional
            after="string_example",  # optional
        )
        pprint(paginate_response.body)
        pprint(paginate_response.body["edges"])
        pprint(paginate_response.body["page_info"])
        pprint(paginate_response.headers)
        pprint(paginate_response.status)
        pprint(paginate_response.round_trip_time)
    except ApiException as e:
        print("Exception when calling MiscellaneousApi.paginate: %s\n" % e)
        pprint(e.body)
        pprint(e.headers)
        pprint(e.status)
        pprint(e.reason)
        pprint(e.round_trip_time)


asyncio.run(main())
```


## Documentation for API Endpoints

All URIs are relative to *http://petstore.swagger.io/v2*

Tag | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*miscellaneous* | [**paginate**](docs/apis/tags/MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox
*pet* | [**add**](docs/apis/tags/PetApi.md#add) | **POST** /pet | Add a new pet to the store
*pet* | [**delete**](docs/apis/tags/PetApi.md#delete) | **DELETE** /pet/{petId} | Deletes a pet
*pet* | [**find_by_status**](docs/apis/tags/PetApi.md#find_by_status) | **GET** /pet/findByStatus | Finds Pets by status
*pet* | [**find_by_tags**](docs/apis/tags/PetApi.md#find_by_tags) | **GET** /pet/findByTags | Finds Pets by tags
*pet* | [**get_by_id**](docs/apis/tags/PetApi.md#get_by_id) | **GET** /pet/{petId} | Find pet by ID
*pet* | [**update**](docs/apis/tags/PetApi.md#update) | **PUT** /pet | Update an existing pet
*pet* | [**update_with_form**](docs/apis/tags/PetApi.md#update_with_form) | **POST** /pet/{petId} | Updates a pet in the store with form data
*pet* | [**upload_image**](docs/apis/tags/PetApi.md#upload_image) | **POST** /pet/{petId}/uploadImage | uploads an image
*store* | [**delete_order**](docs/apis/tags/StoreApi.md#delete_order) | **DELETE** /store/order/{orderId} | Delete purchase order by ID
*store* | [**get_inventory**](docs/apis/tags/StoreApi.md#get_inventory) | **GET** /store/inventory | Returns pet inventories by status
*store* | [**get_order_by_id**](docs/apis/tags/StoreApi.md#get_order_by_id) | **GET** /store/order/{orderId} | Find purchase order by ID
*store* | [**place_order**](docs/apis/tags/StoreApi.md#place_order) | **POST** /store/order | Place an order for a pet
*user* | [**create**](docs/apis/tags/UserApi.md#create) | **POST** /user | Create user
*user* | [**create_with_array**](docs/apis/tags/UserApi.md#create_with_array) | **POST** /user/createWithArray | Creates list of users with given input array
*user* | [**create_with_list**](docs/apis/tags/UserApi.md#create_with_list) | **POST** /user/createWithList | Creates list of users with given input array
*user* | [**delete**](docs/apis/tags/UserApi.md#delete) | **DELETE** /user/{username} | Delete user
*user* | [**get_by_name**](docs/apis/tags/UserApi.md#get_by_name) | **GET** /user/{username} | Get user by user name
*user* | [**login**](docs/apis/tags/UserApi.md#login) | **GET** /user/login | Logs user into the system
*user* | [**logout**](docs/apis/tags/UserApi.md#logout) | **GET** /user/logout | Logs out current logged in user session
*user* | [**update**](docs/apis/tags/UserApi.md#update) | **PUT** /user/{username} | Updated user

## Documentation For Models

 - [ApiResponse](docs/models/ApiResponse.md)
 - [Category](docs/models/Category.md)
 - [CreateWithArrayRequest](docs/models/CreateWithArrayRequest.md)
 - [FindByStatus200Response](docs/models/FindByStatus200Response.md)
 - [FindByStatusResponse](docs/models/FindByStatusResponse.md)
 - [FindByTags200Response](docs/models/FindByTags200Response.md)
 - [FindByTagsResponse](docs/models/FindByTagsResponse.md)
 - [GetInventoryResponse](docs/models/GetInventoryResponse.md)
 - [Login200Response](docs/models/Login200Response.md)
 - [LoginResponse](docs/models/LoginResponse.md)
 - [Order](docs/models/Order.md)
 - [PaginateRequest](docs/models/PaginateRequest.md)
 - [PaginateResponse](docs/models/PaginateResponse.md)
 - [PaginateResponseEdges](docs/models/PaginateResponseEdges.md)
 - [PaginateResponseEdgesNode](docs/models/PaginateResponseEdgesNode.md)
 - [PaginateResponsePageInfo](docs/models/PaginateResponsePageInfo.md)
 - [Pet](docs/models/Pet.md)
 - [PetPhotoUrls](docs/models/PetPhotoUrls.md)
 - [Tag](docs/models/Tag.md)
 - [UpdateWithFormRequest](docs/models/UpdateWithFormRequest.md)
 - [UploadImageRequest](docs/models/UploadImageRequest.md)
 - [User](docs/models/User.md)


## Author
This Python package is automatically generated by [Konfig](https://konfigthis.com)
