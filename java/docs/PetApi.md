# PetApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**add**](PetApi.md#add) | **POST** /pet | Add a new pet to the store |
| [**delete**](PetApi.md#delete) | **DELETE** /pet/{petId} | Deletes a pet |
| [**findByStatus**](PetApi.md#findByStatus) | **GET** /pet/findByStatus | Finds Pets by status |
| [**findByTags**](PetApi.md#findByTags) | **GET** /pet/findByTags | Finds Pets by tags |
| [**getById**](PetApi.md#getById) | **GET** /pet/{petId} | Find pet by ID |
| [**update**](PetApi.md#update) | **PUT** /pet | Update an existing pet |
| [**updateWithForm**](PetApi.md#updateWithForm) | **POST** /pet/{petId} | Updates a pet in the store with form data |
| [**uploadImage**](PetApi.md#uploadImage) | **POST** /pet/{petId}/uploadImage | uploads an image |


<a name="add"></a>
# **add**
> Pet add(pet)

Add a new pet to the store



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure OAuth2 access token for authorization: petstore_auth
    OAuth petstore_auth = (OAuth) defaultClient.getAuthentication("petstore_auth");
    petstore_auth.setAccessToken("YOUR ACCESS TOKEN");

    PetApi apiInstance = new PetApi(defaultClient);
    Pet pet = new Pet(); // Pet | Pet object that needs to be added to the store
    try {
      Pet result = apiInstance.add(pet);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#add");
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
| **pet** | [**Pet**](Pet.md)| Pet object that needs to be added to the store | |

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

<a name="delete"></a>
# **delete**
> delete(petId)

Deletes a pet



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    // Configure OAuth2 access token for authorization: petstore_auth
    OAuth petstore_auth = (OAuth) defaultClient.getAuthentication("petstore_auth");
    petstore_auth.setAccessToken("YOUR ACCESS TOKEN");

    PetApi apiInstance = new PetApi(defaultClient);
    Long petId = 56L; // Long | Pet id to delete
    try {
      apiInstance.delete(petId);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#delete");
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
| **petId** | **Long**| Pet id to delete | |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key), [petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | Invalid pet value |  -  |

<a name="findByStatus"></a>
# **findByStatus**
> List&lt;Pet&gt; findByStatus(status)

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure OAuth2 access token for authorization: petstore_auth
    OAuth petstore_auth = (OAuth) defaultClient.getAuthentication("petstore_auth");
    petstore_auth.setAccessToken("YOUR ACCESS TOKEN");

    PetApi apiInstance = new PetApi(defaultClient);
    List<String> status = Arrays.asList("available"); // List<String> | Status values that need to be considered for filter
    try {
      List<Pet> result = apiInstance.findByStatus(status);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#findByStatus");
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
| **status** | [**List&lt;String&gt;**](String.md)| Status values that need to be considered for filter | [enum: available, pending, sold] |

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

<a name="findByTags"></a>
# **findByTags**
> List&lt;Pet&gt; findByTags(tags)

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure OAuth2 access token for authorization: petstore_auth
    OAuth petstore_auth = (OAuth) defaultClient.getAuthentication("petstore_auth");
    petstore_auth.setAccessToken("YOUR ACCESS TOKEN");

    PetApi apiInstance = new PetApi(defaultClient);
    List<String> tags = Arrays.asList(); // List<String> | Tags to filter by
    try {
      List<Pet> result = apiInstance.findByTags(tags);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#findByTags");
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
| **tags** | [**List&lt;String&gt;**](String.md)| Tags to filter by | |

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

<a name="getById"></a>
# **getById**
> Pet getById(petId)

Find pet by ID

Returns a single pet

### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure API key authorization: api_key
    ApiKeyAuth api_key = (ApiKeyAuth) defaultClient.getAuthentication("api_key");
    api_key.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //api_key.setApiKeyPrefix("Token");

    PetApi apiInstance = new PetApi(defaultClient);
    Long petId = 56L; // Long | ID of pet to return
    try {
      Pet result = apiInstance.getById(petId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#getById");
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
| **petId** | **Long**| ID of pet to return | |

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

<a name="update"></a>
# **update**
> Pet update(pet)

Update an existing pet



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure OAuth2 access token for authorization: petstore_auth
    OAuth petstore_auth = (OAuth) defaultClient.getAuthentication("petstore_auth");
    petstore_auth.setAccessToken("YOUR ACCESS TOKEN");

    PetApi apiInstance = new PetApi(defaultClient);
    Pet pet = new Pet(); // Pet | Pet object that needs to be added to the store
    try {
      Pet result = apiInstance.update(pet);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#update");
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
| **pet** | [**Pet**](Pet.md)| Pet object that needs to be added to the store | |

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

<a name="updateWithForm"></a>
# **updateWithForm**
> updateWithForm(petId, name, status)

Updates a pet in the store with form data



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure OAuth2 access token for authorization: petstore_auth
    OAuth petstore_auth = (OAuth) defaultClient.getAuthentication("petstore_auth");
    petstore_auth.setAccessToken("YOUR ACCESS TOKEN");

    PetApi apiInstance = new PetApi(defaultClient);
    Long petId = 56L; // Long | ID of pet that needs to be updated
    String name = "name_example"; // String | Updated name of the pet
    String status = "status_example"; // String | Updated status of the pet
    try {
      apiInstance.updateWithForm(petId, name, status);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#updateWithForm");
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
| **petId** | **Long**| ID of pet that needs to be updated | |
| **name** | **String**| Updated name of the pet | [optional] |
| **status** | **String**| Updated status of the pet | [optional] |

### Return type

null (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | Invalid input |  -  |

<a name="uploadImage"></a>
# **uploadImage**
> ModelApiResponse uploadImage(petId, additionalMetadata, _file)

uploads an image



### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");
    
    // Configure OAuth2 access token for authorization: petstore_auth
    OAuth petstore_auth = (OAuth) defaultClient.getAuthentication("petstore_auth");
    petstore_auth.setAccessToken("YOUR ACCESS TOKEN");

    PetApi apiInstance = new PetApi(defaultClient);
    Long petId = 56L; // Long | ID of pet to update
    String additionalMetadata = "additionalMetadata_example"; // String | Additional data to pass to server
    File _file = new File("/path/to/file"); // File | file to upload
    try {
      ModelApiResponse result = apiInstance.uploadImage(petId, additionalMetadata, _file);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#uploadImage");
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
| **petId** | **Long**| ID of pet to update | |
| **additionalMetadata** | **String**| Additional data to pass to server | [optional] |
| **_file** | **File**| file to upload | [optional] |

### Return type

[**ModelApiResponse**](ModelApiResponse.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |

