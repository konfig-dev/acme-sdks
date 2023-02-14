# UserApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](UserApi.md#create) | **POST** /user | Create user
[**createWithArray**](UserApi.md#createWithArray) | **POST** /user/createWithArray | Creates list of users with given input array
[**createWithList**](UserApi.md#createWithList) | **POST** /user/createWithList | Creates list of users with given input array
[**delete**](UserApi.md#delete) | **DELETE** /user/{username} | Delete user
[**getByName**](UserApi.md#getByName) | **GET** /user/{username} | Get user by user name
[**login**](UserApi.md#login) | **GET** /user/login | Logs user into the system
[**logout**](UserApi.md#logout) | **GET** /user/logout | Logs out current logged in user session
[**update**](UserApi.md#update) | **PUT** /user/{username} | Updated user


<a name="create"></a>
# **create**
> create(user)

Create user

This can only be done by the logged in user.

### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
val user : User =  // User | Created user object
try {
    apiInstance.create(user)
} catch (e: ClientException) {
    println("4xx response calling UserApi#create")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#create")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md)| Created user object |

### Return type

null (empty response body)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

<a name="createWithArray"></a>
# **createWithArray**
> createWithArray(user)

Creates list of users with given input array



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
val user : kotlin.collections.List<User> =  // kotlin.collections.List<User> | List of user object
try {
    apiInstance.createWithArray(user)
} catch (e: ClientException) {
    println("4xx response calling UserApi#createWithArray")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#createWithArray")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**kotlin.collections.List&lt;User&gt;**](User.md)| List of user object |

### Return type

null (empty response body)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

<a name="createWithList"></a>
# **createWithList**
> createWithList(user)

Creates list of users with given input array



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
val user : kotlin.collections.List<User> =  // kotlin.collections.List<User> | List of user object
try {
    apiInstance.createWithList(user)
} catch (e: ClientException) {
    println("4xx response calling UserApi#createWithList")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#createWithList")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**kotlin.collections.List&lt;User&gt;**](User.md)| List of user object |

### Return type

null (empty response body)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

<a name="delete"></a>
# **delete**
> delete(username)

Delete user

This can only be done by the logged in user.

### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
val username : kotlin.String = username_example // kotlin.String | The name that needs to be deleted
try {
    apiInstance.delete(username)
} catch (e: ClientException) {
    println("4xx response calling UserApi#delete")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#delete")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **kotlin.String**| The name that needs to be deleted |

### Return type

null (empty response body)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="getByName"></a>
# **getByName**
> User getByName(username)

Get user by user name



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
val username : kotlin.String = username_example // kotlin.String | The name that needs to be fetched. Use user1 for testing.
try {
    val result : User = apiInstance.getByName(username)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling UserApi#getByName")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#getByName")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **kotlin.String**| The name that needs to be fetched. Use user1 for testing. |

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="login"></a>
# **login**
> kotlin.String login(username, password)

Logs user into the system



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
val username : kotlin.String = username_example // kotlin.String | The user name for login
val password : kotlin.String = password_example // kotlin.String | The password for login in clear text
try {
    val result : kotlin.String = apiInstance.login(username, password)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling UserApi#login")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#login")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **kotlin.String**| The user name for login |
 **password** | **kotlin.String**| The password for login in clear text |

### Return type

**kotlin.String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="logout"></a>
# **logout**
> logout()

Logs out current logged in user session



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
try {
    apiInstance.logout()
} catch (e: ClientException) {
    println("4xx response calling UserApi#logout")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#logout")
    e.printStackTrace()
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="update"></a>
# **update**
> update(username, user)

Updated user

This can only be done by the logged in user.

### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = UserApi()
val username : kotlin.String = username_example // kotlin.String | name that need to be deleted
val user : User =  // User | Updated user object
try {
    apiInstance.update(username, user)
} catch (e: ClientException) {
    println("4xx response calling UserApi#update")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling UserApi#update")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **kotlin.String**| name that need to be deleted |
 **user** | [**User**](User.md)| Updated user object |

### Return type

null (empty response body)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

