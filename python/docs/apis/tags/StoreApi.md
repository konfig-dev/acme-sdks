<a name="__pageTop"></a>
# acme_client.apis.tags.store_api.StoreApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_order**](#delete_order) | **delete** /store/order/{orderId} | Delete purchase order by ID
[**get_inventory**](#get_inventory) | **get** /store/inventory | Returns pet inventories by status
[**get_order_by_id**](#get_order_by_id) | **get** /store/order/{orderId} | Find purchase order by ID
[**place_order**](#place_order) | **post** /store/order | Place an order for a pet

# **delete_order**

Delete purchase order by ID

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
)

try:
    # Delete purchase order by ID
    acme.store.delete_order(
        order_id="orderId_example",  # required
    )
    pprint(delete_order_response.headers)
    pprint(delete_order_response.status)
    pprint(delete_order_response.round_trip_time)
except ApiException as e:
    print("Exception when calling StoreApi.delete_order: %s\n" % e)
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
orderId | OrderIdSchema | | 

# OrderIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#delete_order.ApiResponseFor200) | OK
400 | [ApiResponseFor400](#delete_order.ApiResponseFor400) | Invalid ID supplied
404 | [ApiResponseFor404](#delete_order.ApiResponseFor404) | Order not found

#### delete_order.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### delete_order.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### delete_order.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **get_inventory**

Returns pet inventories by status

Returns a map of status codes to quantities

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
    # Returns pet inventories by status
    get_inventory_response = acme.store.get_inventory()
    pprint(get_inventory_response.body)
    pprint(get_inventory_response.headers)
    pprint(get_inventory_response.status)
    pprint(get_inventory_response.round_trip_time)
except ApiException as e:
    print("Exception when calling StoreApi.get_inventory: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#get_inventory.ApiResponseFor200) | successful operation

#### get_inventory.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**GetInventoryResponse**](../../models/GetInventoryResponse.md) |  | 


### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **get_order_by_id**

Find purchase order by ID

For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
)

try:
    # Find purchase order by ID
    get_order_by_id_response = acme.store.get_order_by_id(
        order_id=1,  # required
    )
    pprint(get_order_by_id_response.body)
    pprint(get_order_by_id_response.body["id"])
    pprint(get_order_by_id_response.body["pet_id"])
    pprint(get_order_by_id_response.body["quantity"])
    pprint(get_order_by_id_response.body["ship_date"])
    pprint(get_order_by_id_response.body["status"])
    pprint(get_order_by_id_response.body["complete"])
    pprint(get_order_by_id_response.headers)
    pprint(get_order_by_id_response.status)
    pprint(get_order_by_id_response.round_trip_time)
except ApiException as e:
    print("Exception when calling StoreApi.get_order_by_id: %s\n" % e)
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
orderId | OrderIdSchema | | 

# OrderIdSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  |  | value must be a 64 bit integer

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#get_order_by_id.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#get_order_by_id.ApiResponseFor400) | Invalid ID supplied
404 | [ApiResponseFor404](#get_order_by_id.ApiResponseFor404) | Order not found

#### get_order_by_id.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**Order**](../../models/Order.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Order**](../../models/Order.md) |  | 


#### get_order_by_id.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### get_order_by_id.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **place_order**

Place an order for a pet

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
)

try:
    # Place an order for a pet
    place_order_response = acme.store.place_order(
        id=1,  # optional
        pet_id=1,  # optional
        quantity=1,  # optional
        ship_date="1970-01-01T00:00:00.00Z",  # optional
        status="placed",  # optional
        complete=False,  # optional
    )
    pprint(place_order_response.body)
    pprint(place_order_response.body["id"])
    pprint(place_order_response.body["pet_id"])
    pprint(place_order_response.body["quantity"])
    pprint(place_order_response.body["ship_date"])
    pprint(place_order_response.body["status"])
    pprint(place_order_response.body["complete"])
    pprint(place_order_response.headers)
    pprint(place_order_response.status)
    pprint(place_order_response.round_trip_time)
except ApiException as e:
    print("Exception when calling StoreApi.place_order: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/xml', 'application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Order**](../../models/Order.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#place_order.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#place_order.ApiResponseFor400) | Invalid Order

#### place_order.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**Order**](../../models/Order.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Order**](../../models/Order.md) |  | 


#### place_order.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

