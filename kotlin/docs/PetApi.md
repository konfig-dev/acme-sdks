# PetApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add**](PetApi.md#add) | **POST** /pet | Add a new pet to the store
[**delete**](PetApi.md#delete) | **DELETE** /pet/{petId} | Deletes a pet
[**findByStatus**](PetApi.md#findByStatus) | **GET** /pet/findByStatus | Finds Pets by status
[**findByTags**](PetApi.md#findByTags) | **GET** /pet/findByTags | Finds Pets by tags
[**getById**](PetApi.md#getById) | **GET** /pet/{petId} | Find pet by ID
[**update**](PetApi.md#update) | **PUT** /pet | Update an existing pet
[**updateWithForm**](PetApi.md#updateWithForm) | **POST** /pet/{petId} | Updates a pet in the store with form data
[**uploadImage**](PetApi.md#uploadImage) | **POST** /pet/{petId}/uploadImage | uploads an image


<a name="add"></a>
# **add**
> Pet add(pet)

Add a new pet to the store



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val pet : Pet =  // Pet | Pet object that needs to be added to the store
try {
    val result : Pet = apiInstance.add(pet)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling PetApi#add")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#add")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**Pet**](Pet.md)| Pet object that needs to be added to the store |

### Return type

[**Pet**](Pet.md)

### Authorization


Configure petstore_auth:
    ApiClient.accessToken = ""

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="delete"></a>
# **delete**
> delete(petId)

Deletes a pet



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val petId : kotlin.Long = 789 // kotlin.Long | Pet id to delete
try {
    apiInstance.delete(petId)
} catch (e: ClientException) {
    println("4xx response calling PetApi#delete")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#delete")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **kotlin.Long**| Pet id to delete |

### Return type

null (empty response body)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""
Configure petstore_auth:
    ApiClient.accessToken = ""

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="findByStatus"></a>
# **findByStatus**
> kotlin.collections.List&lt;Pet&gt; findByStatus(status)

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val status : kotlin.collections.List<kotlin.String> =  // kotlin.collections.List<kotlin.String> | Status values that need to be considered for filter
try {
    val result : kotlin.collections.List<Pet> = apiInstance.findByStatus(status)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling PetApi#findByStatus")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#findByStatus")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**kotlin.collections.List&lt;kotlin.String&gt;**](kotlin.String.md)| Status values that need to be considered for filter | [enum: available, pending, sold]

### Return type

[**kotlin.collections.List&lt;Pet&gt;**](Pet.md)

### Authorization


Configure petstore_auth:
    ApiClient.accessToken = ""

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="findByTags"></a>
# **findByTags**
> kotlin.collections.List&lt;Pet&gt; findByTags(tags)

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val tags : kotlin.collections.List<kotlin.String> =  // kotlin.collections.List<kotlin.String> | Tags to filter by
try {
    val result : kotlin.collections.List<Pet> = apiInstance.findByTags(tags)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling PetApi#findByTags")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#findByTags")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**kotlin.collections.List&lt;kotlin.String&gt;**](kotlin.String.md)| Tags to filter by |

### Return type

[**kotlin.collections.List&lt;Pet&gt;**](Pet.md)

### Authorization


Configure petstore_auth:
    ApiClient.accessToken = ""

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getById"></a>
# **getById**
> Pet getById(petId)

Find pet by ID

Returns a single pet

### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val petId : kotlin.Long = 789 // kotlin.Long | ID of pet to return
try {
    val result : Pet = apiInstance.getById(petId)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling PetApi#getById")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#getById")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **kotlin.Long**| ID of pet to return |

### Return type

[**Pet**](Pet.md)

### Authorization


Configure api_key:
    ApiClient.apiKey["api_key"] = ""
    ApiClient.apiKeyPrefix["api_key"] = ""

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="update"></a>
# **update**
> Pet update(pet)

Update an existing pet



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val pet : Pet =  // Pet | Pet object that needs to be added to the store
try {
    val result : Pet = apiInstance.update(pet)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling PetApi#update")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#update")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**Pet**](Pet.md)| Pet object that needs to be added to the store |

### Return type

[**Pet**](Pet.md)

### Authorization


Configure petstore_auth:
    ApiClient.accessToken = ""

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateWithForm"></a>
# **updateWithForm**
> updateWithForm(petId, name, status)

Updates a pet in the store with form data



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val petId : kotlin.Long = 789 // kotlin.Long | ID of pet that needs to be updated
val name : kotlin.String = name_example // kotlin.String | Updated name of the pet
val status : kotlin.String = status_example // kotlin.String | Updated status of the pet
try {
    apiInstance.updateWithForm(petId, name, status)
} catch (e: ClientException) {
    println("4xx response calling PetApi#updateWithForm")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#updateWithForm")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **kotlin.Long**| ID of pet that needs to be updated |
 **name** | **kotlin.String**| Updated name of the pet | [optional]
 **status** | **kotlin.String**| Updated status of the pet | [optional]

### Return type

null (empty response body)

### Authorization


Configure petstore_auth:
    ApiClient.accessToken = ""

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

<a name="uploadImage"></a>
# **uploadImage**
> ModelApiResponse uploadImage(petId, additionalMetadata, file)

uploads an image



### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = PetApi()
val petId : kotlin.Long = 789 // kotlin.Long | ID of pet to update
val additionalMetadata : kotlin.String = additionalMetadata_example // kotlin.String | Additional data to pass to server
val file : java.io.File = BINARY_DATA_HERE // java.io.File | file to upload
try {
    val result : ModelApiResponse = apiInstance.uploadImage(petId, additionalMetadata, file)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling PetApi#uploadImage")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling PetApi#uploadImage")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **kotlin.Long**| ID of pet to update |
 **additionalMetadata** | **kotlin.String**| Additional data to pass to server | [optional]
 **file** | **java.io.File**| file to upload | [optional]

### Return type

[**ModelApiResponse**](ModelApiResponse.md)

### Authorization


Configure petstore_auth:
    ApiClient.accessToken = ""

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

