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
> create(user).execute();

Create user

This can only be done by the logged in user.

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    Acme client = new Acme(configuration);
    Long id = 56L;
    String username = "username_example";
    String firstName = "firstName_example";
    String lastName = "lastName_example";
    String email = "email_example";
    String password = "password_example";
    String phone = "phone_example";
    Integer userStatus = 56; // User Status
    try {
      client
              .user
              .create()
              .id(id)
              .username(username)
              .firstName(firstName)
              .lastName(lastName)
              .email(email)
              .password(password)
              .phone(phone)
              .userStatus(userStatus)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#create");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .user
              .create()
              .id(id)
              .username(username)
              .firstName(firstName)
              .lastName(lastName)
              .email(email)
              .password(password)
              .phone(phone)
              .userStatus(userStatus)
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#create");
      System.err.println("Status code: " + e.getStatusCode());
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
> createWithArray(user).execute();

Creates list of users with given input array



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    Acme client = new Acme(configuration);
    try {
      client
              .user
              .createWithArray()
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#createWithArray");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .user
              .createWithArray()
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#createWithArray");
      System.err.println("Status code: " + e.getStatusCode());
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
> createWithList(user).execute();

Creates list of users with given input array



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    Acme client = new Acme(configuration);
    try {
      client
              .user
              .createWithList()
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#createWithList");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .user
              .createWithList()
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#createWithList");
      System.err.println("Status code: " + e.getStatusCode());
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
> delete(username).execute();

Delete user

This can only be done by the logged in user.

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    Acme client = new Acme(configuration);
    String username = "username_example"; // The name that needs to be deleted
    try {
      client
              .user
              .delete(username)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#delete");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .user
              .delete(username)
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#delete");
      System.err.println("Status code: " + e.getStatusCode());
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
> User getByName(username).execute();

Get user by user name



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";

    Acme client = new Acme(configuration);
    String username = "username_example"; // The name that needs to be fetched. Use user1 for testing.
    try {
      User result = client
              .user
              .getByName(username)
              .execute();
      System.out.println(result);

      System.out.println(result.getId());

      System.out.println(result.getUsername());

      System.out.println(result.getFirstName());

      System.out.println(result.getLastName());

      System.out.println(result.getEmail());

      System.out.println(result.getPassword());

      System.out.println(result.getPhone());

      System.out.println(result.getUserStatus());

    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#getByName");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<User> response = client
              .user
              .getByName(username)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#getByName");
      System.err.println("Status code: " + e.getStatusCode());
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
> String login(username, password).execute();

Logs user into the system



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";

    Acme client = new Acme(configuration);
    String username = "username_example"; // The user name for login
    String password = "password_example"; // The password for login in clear text
    try {
      String result = client
              .user
              .login(username, password)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#login");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<String> response = client
              .user
              .login(username, password)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#login");
      System.err.println("Status code: " + e.getStatusCode());
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
> logout().execute();

Logs out current logged in user session



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    Acme client = new Acme(configuration);
    try {
      client
              .user
              .logout()
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#logout");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .user
              .logout()
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#logout");
      System.err.println("Status code: " + e.getStatusCode());
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
> update(username, user).execute();

Updated user

This can only be done by the logged in user.

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.UserApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    Acme client = new Acme(configuration);
    String username = "username_example"; // name that need to be deleted
    Long id = 56L;
    String username = "username_example";
    String firstName = "firstName_example";
    String lastName = "lastName_example";
    String email = "email_example";
    String password = "password_example";
    String phone = "phone_example";
    Integer userStatus = 56; // User Status
    try {
      client
              .user
              .update(username)
              .id(id)
              .username(username)
              .firstName(firstName)
              .lastName(lastName)
              .email(email)
              .password(password)
              .phone(phone)
              .userStatus(userStatus)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#update");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .user
              .update(username)
              .id(id)
              .username(username)
              .firstName(firstName)
              .lastName(lastName)
              .email(email)
              .password(password)
              .phone(phone)
              .userStatus(userStatus)
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling UserApi#update");
      System.err.println("Status code: " + e.getStatusCode());
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

