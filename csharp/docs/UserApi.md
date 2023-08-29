# Acme.Net.Api.UserApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**Create**](UserApi.md#create) | **POST** /user | Create user |
| [**CreateWithArray**](UserApi.md#createwitharray) | **POST** /user/createWithArray | Creates list of users with given input array |
| [**CreateWithList**](UserApi.md#createwithlist) | **POST** /user/createWithList | Creates list of users with given input array |
| [**Delete**](UserApi.md#delete) | **DELETE** /user/{username} | Delete user |
| [**GetByName**](UserApi.md#getbyname) | **GET** /user/{username} | Get user by user name |
| [**Login**](UserApi.md#login) | **GET** /user/login | Logs user into the system |
| [**Logout**](UserApi.md#logout) | **GET** /user/logout | Logs out current logged in user session |
| [**Update**](UserApi.md#update) | **PUT** /user/{username} | Updated user |

<a name="create"></a>
# **Create**
> void Create (User user)

Create user

This can only be done by the logged in user.

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class CreateExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");
            // Configure API key authorization: api_key
            client.SetApiKey("YOUR_API_KEY");

            var user = new User(); // User | Created user object

            try
            {
                // Create user
                client.User.Create(user);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.Create: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the CreateWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Create user
    apiInstance.CreateWithHttpInfo(user);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.CreateWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **user** | [**User**](User.md) | Created user object |  |

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
| **0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="createwitharray"></a>
# **CreateWithArray**
> void CreateWithArray (List<User> user)

Creates list of users with given input array

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class CreateWithArrayExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");
            // Configure API key authorization: api_key
            client.SetApiKey("YOUR_API_KEY");

            var user = new List<User>(); // List<User> | List of user object

            try
            {
                // Creates list of users with given input array
                client.User.CreateWithArray(user);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.CreateWithArray: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the CreateWithArrayWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Creates list of users with given input array
    apiInstance.CreateWithArrayWithHttpInfo(user);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.CreateWithArrayWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **user** | [**List&lt;User&gt;**](User.md) | List of user object |  |

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
| **0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="createwithlist"></a>
# **CreateWithList**
> void CreateWithList (List<User> user)

Creates list of users with given input array

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class CreateWithListExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");
            // Configure API key authorization: api_key
            client.SetApiKey("YOUR_API_KEY");

            var user = new List<User>(); // List<User> | List of user object

            try
            {
                // Creates list of users with given input array
                client.User.CreateWithList(user);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.CreateWithList: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the CreateWithListWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Creates list of users with given input array
    apiInstance.CreateWithListWithHttpInfo(user);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.CreateWithListWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **user** | [**List&lt;User&gt;**](User.md) | List of user object |  |

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
| **0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="delete"></a>
# **Delete**
> void Delete (string username)

Delete user

This can only be done by the logged in user.

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class DeleteExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");
            // Configure API key authorization: api_key
            client.SetApiKey("YOUR_API_KEY");

            var username = "username_example";  // string | The name that needs to be deleted

            try
            {
                // Delete user
                client.User.Delete(username);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.Delete: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the DeleteWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Delete user
    apiInstance.DeleteWithHttpInfo(username);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.DeleteWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **username** | **string** | The name that needs to be deleted |  |

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
| **200** | OK |  -  |
| **400** | Invalid username supplied |  -  |
| **404** | User not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getbyname"></a>
# **GetByName**
> User GetByName (string username)

Get user by user name

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class GetByNameExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");

            var username = "username_example";  // string | The name that needs to be fetched. Use user1 for testing.

            try
            {
                // Get user by user name
                User result = client.User.GetByName(username);
                Console.WriteLine(result);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.GetByName: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the GetByNameWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Get user by user name
    ApiResponse<User> response = apiInstance.GetByNameWithHttpInfo(username);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.GetByNameWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **username** | **string** | The name that needs to be fetched. Use user1 for testing. |  |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="login"></a>
# **Login**
> string Login (string username, string password)

Logs user into the system

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class LoginExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");

            var username = "username_example";  // string | The user name for login
            var password = "password_example";  // string | The password for login in clear text

            try
            {
                // Logs user into the system
                string result = client.User.Login(username, password);
                Console.WriteLine(result);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.Login: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the LoginWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Logs user into the system
    ApiResponse<string> response = apiInstance.LoginWithHttpInfo(username, password);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.LoginWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **username** | **string** | The user name for login |  |
| **password** | **string** | The password for login in clear text |  |

### Return type

**string**

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="logout"></a>
# **Logout**
> void Logout ()

Logs out current logged in user session

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class LogoutExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");
            // Configure API key authorization: api_key
            client.SetApiKey("YOUR_API_KEY");


            try
            {
                // Logs out current logged in user session
                client.User.Logout();
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.Logout: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the LogoutWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Logs out current logged in user session
    apiInstance.LogoutWithHttpInfo();
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.LogoutWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
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
| **0** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="update"></a>
# **Update**
> void Update (string username, User user)

Updated user

This can only be done by the logged in user.

### Example
```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class UpdateExample
    {
        public static void Main()
        {

            AcmeClient client = new AcmeClient();
            // Configure custom BasePath if desired
            client.SetBasePath("http://petstore.swagger.io/v2");
            // Configure API key authorization: api_key
            client.SetApiKey("YOUR_API_KEY");

            var username = "username_example";  // string | name that need to be deleted
            var user = new User(); // User | Updated user object

            try
            {
                // Updated user
                client.User.Update(username, user);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling UserApi.Update: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

#### Using the UpdateWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Updated user
    apiInstance.UpdateWithHttpInfo(username, user);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling UserApi.UpdateWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **username** | **string** | name that need to be deleted |  |
| **user** | [**User**](User.md) | Updated user object |  |

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
| **200** | OK |  -  |
| **400** | Invalid user supplied |  -  |
| **404** | User not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

