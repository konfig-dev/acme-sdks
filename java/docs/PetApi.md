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
> Pet add(pet).execute();

Add a new pet to the store



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure OAuth2 access token for authorization: petstore_auth
    configuration.accessToken = "YOUR ACCESS TOKEN";

    Acme client = new Acme(configuration);
    String name = "name_example";
    List<String> photoUrls = Arrays.asList();
    List<Tag> tags = Arrays.asList();
    Long id = 56L;
    Category category = new Category();
    String status = "available"; // pet status in the store
    try {
      Pet result = client
              .pet
              .add(name, photoUrls)
              .tags(tags)
              .id(id)
              .category(category)
              .status(status)
              .execute();
      System.out.println(result);

      System.out.println(result.getTags());

      System.out.println(result.getId());

      System.out.println(result.getCategory());

      System.out.println(result.getName());

      System.out.println(result.getPhotoUrls());

      System.out.println(result.getStatus());

    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#add");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Pet> response = client
              .pet
              .add(name, photoUrls)
              .tags(tags)
              .id(id)
              .category(category)
              .status(status)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#add");
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
> delete(petId).execute();

Deletes a pet



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    // Configure OAuth2 access token for authorization: petstore_auth
    configuration.accessToken = "YOUR ACCESS TOKEN";

    Acme client = new Acme(configuration);
    Long petId = 56L; // Pet id to delete
    try {
      client
              .pet
              .delete(petId)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#delete");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .pet
              .delete(petId)
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#delete");
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
> List&lt;Pet&gt; findByStatus(status).execute();

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure OAuth2 access token for authorization: petstore_auth
    configuration.accessToken = "YOUR ACCESS TOKEN";

    Acme client = new Acme(configuration);
    List<String> status = Arrays.asList("available"); // Status values that need to be considered for filter
    try {
      List<Pet> result = client
              .pet
              .findByStatus(status)
              .execute();
      System.out.println(result);

    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#findByStatus");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<List<Pet>> response = client
              .pet
              .findByStatus(status)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#findByStatus");
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
> List&lt;Pet&gt; findByTags(tags).execute();

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure OAuth2 access token for authorization: petstore_auth
    configuration.accessToken = "YOUR ACCESS TOKEN";

    Acme client = new Acme(configuration);
    List<String> tags = Arrays.asList(); // Tags to filter by
    try {
      List<Pet> result = client
              .pet
              .findByTags(tags)
              .execute();
      System.out.println(result);

    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#findByTags");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<List<Pet>> response = client
              .pet
              .findByTags(tags)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#findByTags");
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
> Pet getById(petId).execute();

Find pet by ID

Returns a single pet

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
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
    Long petId = 56L; // ID of pet to return
    try {
      Pet result = client
              .pet
              .getById(petId)
              .execute();
      System.out.println(result);

      System.out.println(result.getTags());

      System.out.println(result.getId());

      System.out.println(result.getCategory());

      System.out.println(result.getName());

      System.out.println(result.getPhotoUrls());

      System.out.println(result.getStatus());

    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#getById");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Pet> response = client
              .pet
              .getById(petId)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#getById");
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
> Pet update(pet).execute();

Update an existing pet



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure OAuth2 access token for authorization: petstore_auth
    configuration.accessToken = "YOUR ACCESS TOKEN";

    Acme client = new Acme(configuration);
    String name = "name_example";
    List<String> photoUrls = Arrays.asList();
    List<Tag> tags = Arrays.asList();
    Long id = 56L;
    Category category = new Category();
    String status = "available"; // pet status in the store
    try {
      Pet result = client
              .pet
              .update(name, photoUrls)
              .tags(tags)
              .id(id)
              .category(category)
              .status(status)
              .execute();
      System.out.println(result);

      System.out.println(result.getTags());

      System.out.println(result.getId());

      System.out.println(result.getCategory());

      System.out.println(result.getName());

      System.out.println(result.getPhotoUrls());

      System.out.println(result.getStatus());

    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#update");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Pet> response = client
              .pet
              .update(name, photoUrls)
              .tags(tags)
              .id(id)
              .category(category)
              .status(status)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#update");
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
> updateWithForm(petId).name(name).status(status).execute();

Updates a pet in the store with form data



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure OAuth2 access token for authorization: petstore_auth
    configuration.accessToken = "YOUR ACCESS TOKEN";

    Acme client = new Acme(configuration);
    Long petId = 56L; // ID of pet that needs to be updated
    String name = "name_example"; // Updated name of the pet
    String status = "status_example"; // Updated status of the pet
    try {
      client
              .pet
              .updateWithForm(petId)
              .name(name)
              .status(status)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#updateWithForm");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .pet
              .updateWithForm(petId)
              .name(name)
              .status(status)
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#updateWithForm");
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
> ModelApiResponse uploadImage(petId).additionalMetadata(additionalMetadata)._file(_file).execute();

uploads an image



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.PetApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure OAuth2 access token for authorization: petstore_auth
    configuration.accessToken = "YOUR ACCESS TOKEN";

    Acme client = new Acme(configuration);
    Long petId = 56L; // ID of pet to update
    String additionalMetadata = "additionalMetadata_example"; // Additional data to pass to server
    File _file = new File("/path/to/file"); // file to upload
    try {
      ModelApiResponse result = client
              .pet
              .uploadImage(petId)
              .additionalMetadata(additionalMetadata)
              ._file(_file)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#uploadImage");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<ModelApiResponse> response = client
              .pet
              .uploadImage(petId)
              .additionalMetadata(additionalMetadata)
              ._file(_file)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling PetApi#uploadImage");
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

