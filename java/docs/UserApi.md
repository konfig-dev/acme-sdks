# UserApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**create**](UserApi.md#create) | **POST** /user | Create user |
| [**createWithArray**](UserApi.md#createWithArray) | **POST** /user/createWithArray | Creates list of users with given input array |
| [**createWithList**](UserApi.md#createWithList) | **POST** /user/createWithList | Creates list of users with given input array |
| [**delete**](UserApi.md#delete) | **DELETE** /user/{username} | Delete user |
| [**getByName**](UserApi.md#getByName) | **GET** /user/{username} | Get user by user name |
| [**login**](UserApi.md#login) | **GET** /user/login | Logs user into the system |
| [**logout**](UserApi.md#logout) | **GET** /user/logout | Logs out current logged in user session |
| [**update**](UserApi.md#update) | **PUT** /user/{username} | Updated user |


<a name="create"></a>
# **create**
> create(user)

Create user

This can only be done by the logged in user.

### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    UserApi apiInstance = new UserApi(defaultClient);
    User user = new User(); // User | Created user object
    try {
      apiInstance.create(user);
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#create");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | [**User**](User.md)| Created user object | |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | successful operation |  -  |

<a name="createWithArray"></a>
# **createWithArray**
> createWithArray(user)

Creates list of users with given input array



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    UserApi apiInstance = new UserApi(defaultClient);
    List<User> user = Arrays.asList(); // List<User> | List of user object
    try {
      apiInstance.createWithArray(user);
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#createWithArray");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | [**List&lt;User&gt;**](User.md)| List of user object | |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | successful operation |  -  |

<a name="createWithList"></a>
# **createWithList**
> createWithList(user)

Creates list of users with given input array



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    UserApi apiInstance = new UserApi(defaultClient);
    List<User> user = Arrays.asList(); // List<User> | List of user object
    try {
      apiInstance.createWithList(user);
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#createWithList");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | [**List&lt;User&gt;**](User.md)| List of user object | |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | successful operation |  -  |

<a name="delete"></a>
# **delete**
> delete(username)

Delete user

This can only be done by the logged in user.

### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    UserApi apiInstance = new UserApi(defaultClient);
    String username = "username_example"; // String | The name that needs to be deleted
    try {
      apiInstance.delete(username);
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#delete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| The name that needs to be deleted | |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid username supplied |  -  |
| **404** | User not found |  -  |

<a name="getByName"></a>
# **getByName**
> User getByName(username)

Get user by user name



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");

    UserApi apiInstance = new UserApi(defaultClient);
    String username = "username_example"; // String | The name that needs to be fetched. Use user1 for testing.
    try {
      User result = apiInstance.getByName(username);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#getByName");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| The name that needs to be fetched. Use user1 for testing. | |

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
| **200** | successful operation |  -  |
| **400** | Invalid username supplied |  -  |
| **404** | User not found |  -  |

<a name="login"></a>
# **login**
> String login(username, password)

Logs user into the system



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");

    UserApi apiInstance = new UserApi(defaultClient);
    String username = "username_example"; // String | The user name for login
    String password = "password_example"; // String | The password for login in clear text
    try {
      String result = apiInstance.login(username, password);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#login");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| The user name for login | |
| **password** | **String**| The password for login in clear text | |

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  * Set-Cookie - Cookie authentication key for use with the &#x60;api_key&#x60; apiKey authentication. <br>  * X-Rate-Limit - calls per hour allowed by the user <br>  * X-Expires-After - date in UTC when token expires <br>  |
| **400** | Invalid username/password supplied |  -  |

<a name="logout"></a>
# **logout**
> logout()

Logs out current logged in user session



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    UserApi apiInstance = new UserApi(defaultClient);
    try {
      apiInstance.logout();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#logout");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | successful operation |  -  |

<a name="update"></a>
# **update**
> update(username, user)

Updated user

This can only be done by the logged in user.

### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    UserApi apiInstance = new UserApi(defaultClient);
    String username = "username_example"; // String | name that need to be deleted
    User user = new User(); // User | Updated user object
    try {
      apiInstance.update(username, user);
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#update");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| name that need to be deleted | |
| **user** | [**User**](User.md)| Updated user object | |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid user supplied |  -  |
| **404** | User not found |  -  |

