<a name="__pageTop"></a>
# acme_client.apis.tags.pet_api.PetApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add**](#add) | **post** /pet | Add a new pet to the store
[**delete**](#delete) | **delete** /pet/{petId} | Deletes a pet
[**find_by_status**](#find_by_status) | **get** /pet/findByStatus | Finds Pets by status
[**find_by_tags**](#find_by_tags) | **get** /pet/findByTags | Finds Pets by tags
[**get_by_id**](#get_by_id) | **get** /pet/{petId} | Find pet by ID
[**update**](#update) | **put** /pet | Update an existing pet
[**update_with_form**](#update_with_form) | **post** /pet/{petId} | Updates a pet in the store with form data
[**upload_image**](#upload_image) | **post** /pet/{petId}/uploadImage | uploads an image

# **add**

Add a new pet to the store

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'petstore_auth': 'Bearer'},
    # Configure OAuth2 access token for authorization: petstore_auth
    access_token="YOUR_ACCESS_TOKEN",
)

try:
    # Add a new pet to the store
    add_response = acme.pet.add(
        name="doggie",  # required
        photo_urls=["string_example"],  # required
        tags=[{}],  # optional
        id=1,  # optional
        category={},  # optional
        status="available",  # optional
    )
    pprint(add_response.body)
    pprint(add_response.body["name"])
    pprint(add_response.body["photo_urls"])
    pprint(add_response.body["tags"])
    pprint(add_response.body["id"])
    pprint(add_response.body["category"])
    pprint(add_response.body["status"])
    pprint(add_response.headers)
    pprint(add_response.status)
    pprint(add_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.add: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson, SchemaForRequestBodyApplicationXml] | required |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/xml', 'application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


# SchemaForRequestBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#add.ApiResponseFor200) | successful operation
405 | [ApiResponseFor405](#add.ApiResponseFor405) | Invalid input

#### add.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


#### add.ApiResponseFor405
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[petstore_auth](../../../README.md#petstore_auth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **delete**

Deletes a pet

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Configure API key authorization: api_key
    api_key="YOUR_API_KEY",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'api_key': 'Bearer'},
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'petstore_auth': 'Bearer'},
    # Configure OAuth2 access token for authorization: petstore_auth
    access_token="YOUR_ACCESS_TOKEN",
)

try:
    # Deletes a pet
    acme.pet.delete(
        pet_id=1,  # required
    )
    pprint(delete_response.headers)
    pprint(delete_response.status)
    pprint(delete_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.delete: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
path_params | RequestPathParams | |
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
petId | PetIdSchema | | 

# PetIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | value must be a 64 bit integer

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
default | [ApiResponseForDefault](#delete.ApiResponseForDefault) | Invalid pet value

#### delete.ApiResponseForDefault
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[api_key](../../../README.md#api_key), [petstore_auth](../../../README.md#petstore_auth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **find_by_status**

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'petstore_auth': 'Bearer'},
    # Configure OAuth2 access token for authorization: petstore_auth
    access_token="YOUR_ACCESS_TOKEN",
)

try:
    # Finds Pets by status
    find_by_status_response = acme.pet.find_by_status(
        status=["available"],  # required
    )
    pprint(find_by_status_response.body)
    pprint(find_by_status_response.headers)
    pprint(find_by_status_response.status)
    pprint(find_by_status_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.find_by_status: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
accept_content_types | typing.Tuple[str] | default is ('application/xml', 'application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
status | StatusSchema | | 


# StatusSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
items | str,  | str,  |  | must be one of ["available", "pending", "sold", ] if omitted the server will use the default value of "available"

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#find_by_status.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#find_by_status.ApiResponseFor400) | Invalid status value

#### find_by_status.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**FindByStatusResponse**](../../models/FindByStatusResponse.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**FindByStatus200Response**](../../models/FindByStatus200Response.md) |  | 


#### find_by_status.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[petstore_auth](../../../README.md#petstore_auth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **find_by_tags**

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'petstore_auth': 'Bearer'},
    # Configure OAuth2 access token for authorization: petstore_auth
    access_token="YOUR_ACCESS_TOKEN",
)

try:
    # Finds Pets by tags
    find_by_tags_response = acme.pet.find_by_tags(
        tags=["tags_example"],  # required
    )
    pprint(find_by_tags_response.body)
    pprint(find_by_tags_response.headers)
    pprint(find_by_tags_response.status)
    pprint(find_by_tags_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.find_by_tags: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
query_params | RequestQueryParams | |
accept_content_types | typing.Tuple[str] | default is ('application/xml', 'application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### query_params
#### RequestQueryParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
tags | TagsSchema | | 


# TagsSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
items | str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#find_by_tags.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#find_by_tags.ApiResponseFor400) | Invalid tag value

#### find_by_tags.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**FindByTagsResponse**](../../models/FindByTagsResponse.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**FindByTags200Response**](../../models/FindByTags200Response.md) |  | 


#### find_by_tags.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[petstore_auth](../../../README.md#petstore_auth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **get_by_id**

Find pet by ID

Returns a single pet

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Configure API key authorization: api_key
    api_key="YOUR_API_KEY",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'api_key': 'Bearer'},
)

try:
    # Find pet by ID
    get_by_id_response = acme.pet.get_by_id(
        pet_id=1,  # required
    )
    pprint(get_by_id_response.body)
    pprint(get_by_id_response.body["name"])
    pprint(get_by_id_response.body["photo_urls"])
    pprint(get_by_id_response.body["tags"])
    pprint(get_by_id_response.body["id"])
    pprint(get_by_id_response.body["category"])
    pprint(get_by_id_response.body["status"])
    pprint(get_by_id_response.headers)
    pprint(get_by_id_response.status)
    pprint(get_by_id_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.get_by_id: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
path_params | RequestPathParams | |
accept_content_types | typing.Tuple[str] | default is ('application/xml', 'application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
petId | PetIdSchema | | 

# PetIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | value must be a 64 bit integer

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#get_by_id.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#get_by_id.ApiResponseFor400) | Invalid ID supplied
404 | [ApiResponseFor404](#get_by_id.ApiResponseFor404) | Pet not found

#### get_by_id.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


#### get_by_id.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### get_by_id.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **update**

Update an existing pet

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'petstore_auth': 'Bearer'},
    # Configure OAuth2 access token for authorization: petstore_auth
    access_token="YOUR_ACCESS_TOKEN",
)

try:
    # Update an existing pet
    update_response = acme.pet.update(
        name="doggie",  # required
        photo_urls=["string_example"],  # required
        tags=[{}],  # optional
        id=1,  # optional
        category={},  # optional
        status="available",  # optional
    )
    pprint(update_response.body)
    pprint(update_response.body["name"])
    pprint(update_response.body["photo_urls"])
    pprint(update_response.body["tags"])
    pprint(update_response.body["id"])
    pprint(update_response.body["category"])
    pprint(update_response.body["status"])
    pprint(update_response.headers)
    pprint(update_response.status)
    pprint(update_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.update: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson, SchemaForRequestBodyApplicationXml] | required |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/xml', 'application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


# SchemaForRequestBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#update.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#update.ApiResponseFor400) | Invalid ID supplied
404 | [ApiResponseFor404](#update.ApiResponseFor404) | Pet not found
405 | [ApiResponseFor405](#update.ApiResponseFor405) | Validation exception

#### update.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Pet**](../../models/Pet.md) |  | 


#### update.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### update.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### update.ApiResponseFor405
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[petstore_auth](../../../README.md#petstore_auth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **update_with_form**

Updates a pet in the store with form data

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'petstore_auth': 'Bearer'},
    # Configure OAuth2 access token for authorization: petstore_auth
    access_token="YOUR_ACCESS_TOKEN",
)

try:
    # Updates a pet in the store with form data
    acme.pet.update_with_form(
        pet_id=1,  # required
        name="string_example",  # optional
        status="string_example",  # optional
    )
    pprint(update_with_form_response.headers)
    pprint(update_with_form_response.status)
    pprint(update_with_form_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.update_with_form: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationXWwwFormUrlencoded, Unset] | optional, default is unset |
path_params | RequestPathParams | |
content_type | str | optional, default is 'application/x-www-form-urlencoded' | Selects the schema and serialization of the request body
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationXWwwFormUrlencoded
Type | Description  | Notes
------------- | ------------- | -------------
[**UpdateWithFormRequest**](../../models/UpdateWithFormRequest.md) |  | 


### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
petId | PetIdSchema | | 

# PetIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | value must be a 64 bit integer

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
default | [ApiResponseForDefault](#update_with_form.ApiResponseForDefault) | Invalid input

#### update_with_form.ApiResponseForDefault
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[petstore_auth](../../../README.md#petstore_auth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **upload_image**

uploads an image

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
    # Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
    # api_key_prefix = {'petstore_auth': 'Bearer'},
    # Configure OAuth2 access token for authorization: petstore_auth
    access_token="YOUR_ACCESS_TOKEN",
)

try:
    # uploads an image
    upload_image_response = acme.pet.upload_image(
        pet_id=1,  # required
        additional_metadata="string_example",  # optional
        file=open("/path/to/file", "rb"),  # optional
    )
    pprint(upload_image_response.body)
    pprint(upload_image_response.body["code"])
    pprint(upload_image_response.body["type"])
    pprint(upload_image_response.body["message"])
    pprint(upload_image_response.headers)
    pprint(upload_image_response.status)
    pprint(upload_image_response.round_trip_time)
except ApiException as e:
    print("Exception when calling PetApi.upload_image: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyMultipartFormData, Unset] | optional, default is unset |
path_params | RequestPathParams | |
content_type | str | optional, default is 'multipart/form-data' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyMultipartFormData
Type | Description  | Notes
------------- | ------------- | -------------
[**UploadImageRequest**](../../models/UploadImageRequest.md) |  | 


### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
petId | PetIdSchema | | 

# PetIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | value must be a 64 bit integer

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#upload_image.ApiResponseFor200) | successful operation

#### upload_image.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ApiResponse**](../../models/ApiResponse.md) |  | 


### Authorization

[petstore_auth](../../../README.md#petstore_auth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

