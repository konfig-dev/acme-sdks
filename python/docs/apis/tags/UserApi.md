<a name="__pageTop"></a>
# acme_client.apis.tags.user_api.UserApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](#create) | **post** /user | Create user
[**create_with_array**](#create_with_array) | **post** /user/createWithArray | Creates list of users with given input array
[**create_with_list**](#create_with_list) | **post** /user/createWithList | Creates list of users with given input array
[**delete**](#delete) | **delete** /user/{username} | Delete user
[**get_by_name**](#get_by_name) | **get** /user/{username} | Get user by user name
[**login**](#login) | **get** /user/login | Logs user into the system
[**logout**](#logout) | **get** /user/logout | Logs out current logged in user session
[**update**](#update) | **put** /user/{username} | Updated user

# **create**

Create user

This can only be done by the logged in user.

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
    # Create user
    acme.user.create(
        id=1,  # optional
        username="string_example",  # optional
        first_name="string_example",  # optional
        last_name="string_example",  # optional
        email="string_example",  # optional
        password="string_example",  # optional
        phone="string_example",  # optional
        user_status=1,  # optional
    )
    pprint(create_response.headers)
    pprint(create_response.status)
    pprint(create_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.create: %s\n" % e)
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
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**User**](../../models/User.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
default | [ApiResponseForDefault](#create.ApiResponseForDefault) | successful operation

#### create.ApiResponseForDefault
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **create_with_array**

Creates list of users with given input array

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
    # Creates list of users with given input array
    acme.user.create_with_array(
        body=[{}],  # required
    )
    pprint(create_with_array_response.headers)
    pprint(create_with_array_response.status)
    pprint(create_with_array_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.create_with_array: %s\n" % e)
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
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**CreateWithArrayRequest**](../../models/CreateWithArrayRequest.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
default | [ApiResponseForDefault](#create_with_array.ApiResponseForDefault) | successful operation

#### create_with_array.ApiResponseForDefault
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **create_with_list**

Creates list of users with given input array

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
    # Creates list of users with given input array
    acme.user.create_with_list(
        body=[{}],  # required
    )
    pprint(create_with_list_response.headers)
    pprint(create_with_list_response.status)
    pprint(create_with_list_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.create_with_list: %s\n" % e)
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
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**CreateWithArrayRequest**](../../models/CreateWithArrayRequest.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
default | [ApiResponseForDefault](#create_with_list.ApiResponseForDefault) | successful operation

#### create_with_list.ApiResponseForDefault
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **delete**

Delete user

This can only be done by the logged in user.

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
    # Delete user
    acme.user.delete(
        username="username_example",  # required
    )
    pprint(delete_response.headers)
    pprint(delete_response.status)
    pprint(delete_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.delete: %s\n" % e)
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
username | UsernameSchema | | 

# UsernameSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#delete.ApiResponseFor200) | OK
400 | [ApiResponseFor400](#delete.ApiResponseFor400) | Invalid username supplied
404 | [ApiResponseFor404](#delete.ApiResponseFor404) | User not found

#### delete.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### delete.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### delete.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **get_by_name**

Get user by user name

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
    # Get user by user name
    get_by_name_response = acme.user.get_by_name(
        username="username_example",  # required
    )
    pprint(get_by_name_response.body)
    pprint(get_by_name_response.body["id"])
    pprint(get_by_name_response.body["username"])
    pprint(get_by_name_response.body["first_name"])
    pprint(get_by_name_response.body["last_name"])
    pprint(get_by_name_response.body["email"])
    pprint(get_by_name_response.body["password"])
    pprint(get_by_name_response.body["phone"])
    pprint(get_by_name_response.body["user_status"])
    pprint(get_by_name_response.headers)
    pprint(get_by_name_response.status)
    pprint(get_by_name_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.get_by_name: %s\n" % e)
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
username | UsernameSchema | | 

# UsernameSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#get_by_name.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#get_by_name.ApiResponseFor400) | Invalid username supplied
404 | [ApiResponseFor404](#get_by_name.ApiResponseFor404) | User not found

#### get_by_name.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationXml
Type | Description  | Notes
------------- | ------------- | -------------
[**User**](../../models/User.md) |  | 


# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**User**](../../models/User.md) |  | 


#### get_by_name.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

#### get_by_name.ApiResponseFor404
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **login**

Logs user into the system

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
    # Logs user into the system
    login_response = acme.user.login(
        username="CbUUGjjNSwg0_bs9ZayIMrKdgNvb6gvxmPb9GcsM61ate1RA89q3w1l4eH4XxEz.5awLMdeXylwK0lMGUSM4jsrh4dstlnQUN5vVdMLPA",  # required
        password="password_example",  # required
    )
    pprint(login_response.headers)
    pprint(login_response.status)
    pprint(login_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.login: %s\n" % e)
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
username | UsernameSchema | | 
password | PasswordSchema | | 


# UsernameSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# PasswordSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#login.ApiResponseFor200) | successful operation
400 | [ApiResponseFor400](#login.ApiResponseFor400) | Invalid username/password supplied

#### login.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationXml, SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | ResponseHeadersFor200 |  |

# SchemaFor200ResponseBodyApplicationXml

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

# SchemaFor200ResponseBodyApplicationJson

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 
#### ResponseHeadersFor200

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
Set-Cookie | SetCookieSchema | | optional
X-Rate-Limit | XRateLimitSchema | | optional
X-Expires-After | XExpiresAfterSchema | | optional

# SetCookieSchema

Cookie authentication key for use with the `api_key` apiKey authentication.

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  | Cookie authentication key for use with the &#x60;api_key&#x60; apiKey authentication. | 

# XRateLimitSchema

calls per hour allowed by the user

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
decimal.Decimal, int,  | decimal.Decimal,  | calls per hour allowed by the user | value must be a 32 bit integer

# XExpiresAfterSchema

date in UTC when token expires

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str, datetime,  | str,  | date in UTC when token expires | value must conform to RFC-3339 date-time


#### login.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **logout**

Logs out current logged in user session

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
    # Logs out current logged in user session
    acme.user.logout()
    pprint(logout_response.headers)
    pprint(logout_response.status)
    pprint(logout_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.logout: %s\n" % e)
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
default | [ApiResponseForDefault](#logout.ApiResponseForDefault) | successful operation

#### logout.ApiResponseForDefault
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **update**

Updated user

This can only be done by the logged in user.

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
    # Updated user
    acme.user.update(
        username="username_example",  # required
        id=1,  # optional
        first_name="string_example",  # optional
        last_name="string_example",  # optional
        email="string_example",  # optional
        password="string_example",  # optional
        phone="string_example",  # optional
        user_status=1,  # optional
    )
    pprint(update_response.headers)
    pprint(update_response.status)
    pprint(update_response.round_trip_time)
except ApiException as e:
    print("Exception when calling UserApi.update: %s\n" % e)
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
path_params | RequestPathParams | |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**UserCreateRequest**](../../models/UserCreateRequest.md) |  | 


### path_params
#### RequestPathParams

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
username | UsernameSchema | | 

# UsernameSchema

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
str,  | str,  |  | 

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#update.ApiResponseFor200) | OK
400 | [ApiResponseFor400](#update.ApiResponseFor400) | Invalid user supplied
404 | [ApiResponseFor404](#update.ApiResponseFor404) | User not found

#### update.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | Unset | body was not defined |
headers | Unset | headers were not defined |

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

### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

