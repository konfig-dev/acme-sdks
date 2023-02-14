# acme_client.UserApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](UserApi.md#create) | **POST** /user | Create user
[**create_with_array**](UserApi.md#create_with_array) | **POST** /user/createWithArray | Creates list of users with given input array
[**create_with_list**](UserApi.md#create_with_list) | **POST** /user/createWithList | Creates list of users with given input array
[**delete**](UserApi.md#delete) | **DELETE** /user/{username} | Delete user
[**get_by_name**](UserApi.md#get_by_name) | **GET** /user/{username} | Get user by user name
[**login**](UserApi.md#login) | **GET** /user/login | Logs user into the system
[**logout**](UserApi.md#logout) | **GET** /user/logout | Logs out current logged in user session
[**update**](UserApi.md#update) | **PUT** /user/{username} | Updated user


# **create**
> create(user)

Create user

This can only be done by the logged in user.

### Example

* Api Key Authentication (api_key):

```python
import time
import acme_client
from acme_client.api import user_api
from acme_client.model.user import User
from pprint import pprint
# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: api_key
configuration = acme_client.Configuration(
    api_key={'api_key': 'YOUR_API_KEY'},
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host = 'http://petstore.swagger.io/v2'
)

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api_key'] = 'Bearer'

# Enter a context with an instance of the API client
with acme_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    user = User(
        id=1,
        username="username_example",
        first_name="first_name_example",
        last_name="last_name_example",
        email="email_example",
        password="password_example",
        phone="phone_example",
        user_status=1,
    ) # User | Created user object

    # example passing only required values which don't have defaults set
    try:
        # Create user
        api_instance.create(user)
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->create: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md)| Created user object |

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_with_array**
> create_with_array(create_with_array_request)

Creates list of users with given input array



### Example

* Api Key Authentication (api_key):

```python
import time
import acme_client
from acme_client.api import user_api
from acme_client.model.create_with_array_request import CreateWithArrayRequest
from pprint import pprint
# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: api_key
configuration = acme_client.Configuration(
    api_key={'api_key': 'YOUR_API_KEY'},
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host = 'http://petstore.swagger.io/v2'
)

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api_key'] = 'Bearer'

# Enter a context with an instance of the API client
with acme_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    create_with_array_request = CreateWithArrayRequest([
        User(
            id=1,
            username="username_example",
            first_name="first_name_example",
            last_name="last_name_example",
            email="email_example",
            password="password_example",
            phone="phone_example",
            user_status=1,
        ),
    ]) # CreateWithArrayRequest | List of user object

    # example passing only required values which don't have defaults set
    try:
        # Creates list of users with given input array
        api_instance.create_with_array(create_with_array_request)
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->create_with_array: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_with_array_request** | [**CreateWithArrayRequest**](CreateWithArrayRequest.md)| List of user object |

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_with_list**
> create_with_list(create_with_array_request)

Creates list of users with given input array



### Example

* Api Key Authentication (api_key):

```python
import time
import acme_client
from acme_client.api import user_api
from acme_client.model.create_with_array_request import CreateWithArrayRequest
from pprint import pprint
# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: api_key
configuration = acme_client.Configuration(
    api_key={'api_key': 'YOUR_API_KEY'},
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host = 'http://petstore.swagger.io/v2'
)

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api_key'] = 'Bearer'

# Enter a context with an instance of the API client
with acme_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    create_with_array_request = CreateWithArrayRequest([
        User(
            id=1,
            username="username_example",
            first_name="first_name_example",
            last_name="last_name_example",
            email="email_example",
            password="password_example",
            phone="phone_example",
            user_status=1,
        ),
    ]) # CreateWithArrayRequest | List of user object

    # example passing only required values which don't have defaults set
    try:
        # Creates list of users with given input array
        api_instance.create_with_list(create_with_array_request)
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->create_with_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_with_array_request** | [**CreateWithArrayRequest**](CreateWithArrayRequest.md)| List of user object |

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete**
> delete(username)

Delete user

This can only be done by the logged in user.

### Example

* Api Key Authentication (api_key):

```python
import time
import acme_client
from acme_client.api import user_api
from pprint import pprint
# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: api_key
configuration = acme_client.Configuration(
    api_key={'api_key': 'YOUR_API_KEY'},
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host = 'http://petstore.swagger.io/v2'
)

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api_key'] = 'Bearer'

# Enter a context with an instance of the API client
with acme_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    username = "username_example" # str | The name that needs to be deleted

    # example passing only required values which don't have defaults set
    try:
        # Delete user
        api_instance.delete(username)
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->delete: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **str**| The name that needs to be deleted |

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid username supplied |  -  |
**404** | User not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_by_name**
> User get_by_name(username)

Get user by user name



### Example


```python
import time
import acme_client
from acme_client.api import user_api
from acme_client.model.user import User
from pprint import pprint

# Enter a context with an instance of the API client
with acme_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    username = "username_example" # str | The name that needs to be fetched. Use user1 for testing.

    # example passing only required values which don't have defaults set
    try:
        # Get user by user name
        api_response = api_instance.get_by_name(username)
        pprint(api_response)
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->get_by_name: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **str**| The name that needs to be fetched. Use user1 for testing. |

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**400** | Invalid username supplied |  -  |
**404** | User not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **login**
> str login(username, password)

Logs user into the system



### Example


```python
import time
import acme_client
from acme_client.api import user_api
from pprint import pprint

# Enter a context with an instance of the API client
with acme_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    username = "CbUUGjjNSwg0_bs9ZayIMrKdgNvb6gvxmPb9GcsM61ate1RA89q3w1l4eH4XxEz.5awLMdeXylwK0lMGUSM4jsrh4dstlnQUN5vVdMLPA" # str | The user name for login
    password = "password_example" # str | The password for login in clear text

    # example passing only required values which don't have defaults set
    try:
        # Logs user into the system
        api_response = api_instance.login(username, password)
        pprint(api_response)
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->login: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **str**| The user name for login |
 **password** | **str**| The password for login in clear text |

### Return type

**str**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  * Set-Cookie - Cookie authentication key for use with the &#x60;api_key&#x60; apiKey authentication. <br>  * X-Rate-Limit - calls per hour allowed by the user <br>  * X-Expires-After - date in UTC when token expires <br>  |
**400** | Invalid username/password supplied |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **logout**
> logout()

Logs out current logged in user session



### Example

* Api Key Authentication (api_key):

```python
import time
import acme_client
from acme_client.api import user_api
from pprint import pprint
# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: api_key
configuration = acme_client.Configuration(
    api_key={'api_key': 'YOUR_API_KEY'},
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host = 'http://petstore.swagger.io/v2'
)

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api_key'] = 'Bearer'

# Enter a context with an instance of the API client
with acme_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Logs out current logged in user session
        api_instance.logout()
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->logout: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update**
> update(username, user)

Updated user

This can only be done by the logged in user.

### Example

* Api Key Authentication (api_key):

```python
import time
import acme_client
from acme_client.api import user_api
from acme_client.model.user import User
from pprint import pprint
# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: api_key
configuration = acme_client.Configuration(
    api_key={'api_key': 'YOUR_API_KEY'},
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host = 'http://petstore.swagger.io/v2'
)

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api_key'] = 'Bearer'

# Enter a context with an instance of the API client
with acme_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    username = "username_example" # str | name that need to be deleted
    user = User(
        id=1,
        username="username_example",
        first_name="first_name_example",
        last_name="last_name_example",
        email="email_example",
        password="password_example",
        phone="phone_example",
        user_status=1,
    ) # User | Updated user object

    # example passing only required values which don't have defaults set
    try:
        # Updated user
        api_instance.update(username, user)
    except acme_client.ApiException as e:
        print("Exception when calling UserApi->update: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **str**| name that need to be deleted |
 **user** | [**User**](User.md)| Updated user object |

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Invalid user supplied |  -  |
**404** | User not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

