# Acme.Net.Api.PetApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**Add**](PetApi.md#add) | **POST** /pet | Add a new pet to the store |
| [**Delete**](PetApi.md#delete) | **DELETE** /pet/{petId} | Deletes a pet |
| [**FindByStatus**](PetApi.md#findbystatus) | **GET** /pet/findByStatus | Finds Pets by status |
| [**FindByTags**](PetApi.md#findbytags) | **GET** /pet/findByTags | Finds Pets by tags |
| [**GetById**](PetApi.md#getbyid) | **GET** /pet/{petId} | Find pet by ID |
| [**Update**](PetApi.md#update) | **PUT** /pet | Update an existing pet |
| [**UpdateWithForm**](PetApi.md#updatewithform) | **POST** /pet/{petId} | Updates a pet in the store with form data |
| [**UploadImage**](PetApi.md#uploadimage) | **POST** /pet/{petId}/uploadImage | uploads an image |

<a name="add"></a>
# **Add**
> Pet Add (Pet pet)

Add a new pet to the store

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class AddExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure OAuth2 access token for authorization: petstore_auth
            config.AccessToken = "YOUR_ACCESS_TOKEN";

            var apiInstance = new PetApi(config);
            var pet = new Pet(); // Pet | Pet object that needs to be added to the store

            try
            {
                // Add a new pet to the store
                Pet result = apiInstance.Add(pet);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.Add: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the AddWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Add a new pet to the store
    ApiResponse<Pet> response = apiInstance.AddWithHttpInfo(pet);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.AddWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **pet** | [**Pet**](Pet.md) | Pet object that needs to be added to the store |  |

### Return type

[**Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |
| **405** | Invalid input |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="delete"></a>
# **Delete**
> void Delete (long petId)

Deletes a pet

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class DeleteExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure API key authorization: api_key
            config.AddApiKey("api_key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("api_key", "Bearer");
            // Configure OAuth2 access token for authorization: petstore_auth
            config.AccessToken = "YOUR_ACCESS_TOKEN";

            var apiInstance = new PetApi(config);
            var petId = 789L;  // long | Pet id to delete

            try
            {
                // Deletes a pet
                apiInstance.Delete(petId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.Delete: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
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
    // Deletes a pet
    apiInstance.DeleteWithHttpInfo(petId);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.DeleteWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **petId** | **long** | Pet id to delete |  |

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key), [petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | Invalid pet value |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findbystatus"></a>
# **FindByStatus**
> List&lt;Pet&gt; FindByStatus (List<string> status)

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class FindByStatusExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure OAuth2 access token for authorization: petstore_auth
            config.AccessToken = "YOUR_ACCESS_TOKEN";

            var apiInstance = new PetApi(config);
            var status = new List<string>(); // List<string> | Status values that need to be considered for filter

            try
            {
                // Finds Pets by status
                List<Pet> result = apiInstance.FindByStatus(status);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.FindByStatus: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the FindByStatusWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Finds Pets by status
    ApiResponse<List<Pet>> response = apiInstance.FindByStatusWithHttpInfo(status);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.FindByStatusWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **status** | [**List&lt;string&gt;**](string.md) | Status values that need to be considered for filter |  |

### Return type

[**List&lt;Pet&gt;**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |
| **400** | Invalid status value |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findbytags"></a>
# **FindByTags**
> List&lt;Pet&gt; FindByTags (List<string> tags)

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class FindByTagsExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure OAuth2 access token for authorization: petstore_auth
            config.AccessToken = "YOUR_ACCESS_TOKEN";

            var apiInstance = new PetApi(config);
            var tags = new List<string>(); // List<string> | Tags to filter by

            try
            {
                // Finds Pets by tags
                List<Pet> result = apiInstance.FindByTags(tags);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.FindByTags: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the FindByTagsWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Finds Pets by tags
    ApiResponse<List<Pet>> response = apiInstance.FindByTagsWithHttpInfo(tags);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.FindByTagsWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **tags** | [**List&lt;string&gt;**](string.md) | Tags to filter by |  |

### Return type

[**List&lt;Pet&gt;**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |
| **400** | Invalid tag value |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getbyid"></a>
# **GetById**
> Pet GetById (long petId)

Find pet by ID

Returns a single pet

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class GetByIdExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure API key authorization: api_key
            config.AddApiKey("api_key", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("api_key", "Bearer");

            var apiInstance = new PetApi(config);
            var petId = 789L;  // long | ID of pet to return

            try
            {
                // Find pet by ID
                Pet result = apiInstance.GetById(petId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.GetById: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the GetByIdWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Find pet by ID
    ApiResponse<Pet> response = apiInstance.GetByIdWithHttpInfo(petId);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.GetByIdWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **petId** | **long** | ID of pet to return |  |

### Return type

[**Pet**](Pet.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |
| **400** | Invalid ID supplied |  -  |
| **404** | Pet not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="update"></a>
# **Update**
> Pet Update (Pet pet)

Update an existing pet

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class UpdateExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure OAuth2 access token for authorization: petstore_auth
            config.AccessToken = "YOUR_ACCESS_TOKEN";

            var apiInstance = new PetApi(config);
            var pet = new Pet(); // Pet | Pet object that needs to be added to the store

            try
            {
                // Update an existing pet
                Pet result = apiInstance.Update(pet);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.Update: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
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
    // Update an existing pet
    ApiResponse<Pet> response = apiInstance.UpdateWithHttpInfo(pet);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.UpdateWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **pet** | [**Pet**](Pet.md) | Pet object that needs to be added to the store |  |

### Return type

[**Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |
| **400** | Invalid ID supplied |  -  |
| **404** | Pet not found |  -  |
| **405** | Validation exception |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="updatewithform"></a>
# **UpdateWithForm**
> void UpdateWithForm (long petId, string name = null, string status = null)

Updates a pet in the store with form data

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class UpdateWithFormExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure OAuth2 access token for authorization: petstore_auth
            config.AccessToken = "YOUR_ACCESS_TOKEN";

            var apiInstance = new PetApi(config);
            var petId = 789L;  // long | ID of pet that needs to be updated
            var name = "name_example";  // string | Updated name of the pet (optional) 
            var status = "status_example";  // string | Updated status of the pet (optional) 

            try
            {
                // Updates a pet in the store with form data
                apiInstance.UpdateWithForm(petId, name, status);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.UpdateWithForm: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the UpdateWithFormWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Updates a pet in the store with form data
    apiInstance.UpdateWithFormWithHttpInfo(petId, name, status);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.UpdateWithFormWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **petId** | **long** | ID of pet that needs to be updated |  |
| **name** | **string** | Updated name of the pet | [optional]  |
| **status** | **string** | Updated status of the pet | [optional]  |

### Return type

void (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | Invalid input |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="uploadimage"></a>
# **UploadImage**
> ApiResponse UploadImage (long petId, string additionalMetadata = null, System.IO.Stream file = null)

uploads an image

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class UploadImageExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            // Configure OAuth2 access token for authorization: petstore_auth
            config.AccessToken = "YOUR_ACCESS_TOKEN";

            var apiInstance = new PetApi(config);
            var petId = 789L;  // long | ID of pet to update
            var additionalMetadata = "additionalMetadata_example";  // string | Additional data to pass to server (optional) 
            var file = new System.IO.MemoryStream(System.IO.File.ReadAllBytes("/path/to/file.txt"));  // System.IO.Stream | file to upload (optional) 

            try
            {
                // uploads an image
                ApiResponse result = apiInstance.UploadImage(petId, additionalMetadata, file);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling PetApi.UploadImage: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the UploadImageWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // uploads an image
    ApiResponse<ApiResponse> response = apiInstance.UploadImageWithHttpInfo(petId, additionalMetadata, file);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling PetApi.UploadImageWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **petId** | **long** | ID of pet to update |  |
| **additionalMetadata** | **string** | Additional data to pass to server | [optional]  |
| **file** | **System.IO.Stream****System.IO.Stream** | file to upload | [optional]  |

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

