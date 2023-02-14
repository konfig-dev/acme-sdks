# PetAPI

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add**](PetAPI.md#add) | **POST** /pet | Add a new pet to the store
[**delete**](PetAPI.md#delete) | **DELETE** /pet/{petId} | Deletes a pet
[**findByStatus**](PetAPI.md#findbystatus) | **GET** /pet/findByStatus | Finds Pets by status
[**findByTags**](PetAPI.md#findbytags) | **GET** /pet/findByTags | Finds Pets by tags
[**getById**](PetAPI.md#getbyid) | **GET** /pet/{petId} | Find pet by ID
[**update**](PetAPI.md#update) | **PUT** /pet | Update an existing pet
[**updateWithForm**](PetAPI.md#updatewithform) | **POST** /pet/{petId} | Updates a pet in the store with form data
[**uploadImage**](PetAPI.md#uploadimage) | **POST** /pet/{petId}/uploadImage | uploads an image


# **add**
```swift
    open class func add(pet: Pet, completion: @escaping (_ data: Pet?, _ error: Error?) -> Void)
```

Add a new pet to the store



### Example
```swift
import Acme

let pet = Pet(tags: [Tag(id: 123, name: "name_example")], id: 123, category: Category(id: 123, name: "name_example"), name: "name_example", photoUrls: ["photoUrls_example"], status: "status_example") // Pet | Pet object that needs to be added to the store

// Add a new pet to the store
PetAPI.add(pet: pet) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**Pet**](Pet.md) | Pet object that needs to be added to the store | 

### Return type

[**Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **delete**
```swift
    open class func delete(petId: Int64, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Deletes a pet



### Example
```swift
import Acme

let petId = 987 // Int64 | Pet id to delete

// Deletes a pet
PetAPI.delete(petId: petId) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **Int64** | Pet id to delete | 

### Return type

Void (empty response body)

### Authorization

[api_key](../README.md#api_key), [petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **findByStatus**
```swift
    open class func findByStatus(status: [Status_findByStatus], completion: @escaping (_ data: [Pet]?, _ error: Error?) -> Void)
```

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example
```swift
import Acme

let status = ["status_example"] // [String] | Status values that need to be considered for filter

// Finds Pets by status
PetAPI.findByStatus(status: status) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**[String]**](String.md) | Status values that need to be considered for filter | 

### Return type

[**[Pet]**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **findByTags**
```swift
    open class func findByTags(tags: [String], completion: @escaping (_ data: [Pet]?, _ error: Error?) -> Void)
```

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example
```swift
import Acme

let tags = ["inner_example"] // [String] | Tags to filter by

// Finds Pets by tags
PetAPI.findByTags(tags: tags) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**[String]**](String.md) | Tags to filter by | 

### Return type

[**[Pet]**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **getById**
```swift
    open class func getById(petId: Int64, completion: @escaping (_ data: Pet?, _ error: Error?) -> Void)
```

Find pet by ID

Returns a single pet

### Example
```swift
import Acme

let petId = 987 // Int64 | ID of pet to return

// Find pet by ID
PetAPI.getById(petId: petId) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **Int64** | ID of pet to return | 

### Return type

[**Pet**](Pet.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **update**
```swift
    open class func update(pet: Pet, completion: @escaping (_ data: Pet?, _ error: Error?) -> Void)
```

Update an existing pet



### Example
```swift
import Acme

let pet = Pet(tags: [Tag(id: 123, name: "name_example")], id: 123, category: Category(id: 123, name: "name_example"), name: "name_example", photoUrls: ["photoUrls_example"], status: "status_example") // Pet | Pet object that needs to be added to the store

// Update an existing pet
PetAPI.update(pet: pet) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**Pet**](Pet.md) | Pet object that needs to be added to the store | 

### Return type

[**Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **updateWithForm**
```swift
    open class func updateWithForm(petId: Int64, name: String? = nil, status: String? = nil, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Updates a pet in the store with form data



### Example
```swift
import Acme

let petId = 987 // Int64 | ID of pet that needs to be updated
let name = "name_example" // String | Updated name of the pet (optional)
let status = "status_example" // String | Updated status of the pet (optional)

// Updates a pet in the store with form data
PetAPI.updateWithForm(petId: petId, name: name, status: status) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **Int64** | ID of pet that needs to be updated | 
 **name** | **String** | Updated name of the pet | [optional] 
 **status** | **String** | Updated status of the pet | [optional] 

### Return type

Void (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **uploadImage**
```swift
    open class func uploadImage(petId: Int64, additionalMetadata: String? = nil, file: URL? = nil, completion: @escaping (_ data: ApiResponse?, _ error: Error?) -> Void)
```

uploads an image



### Example
```swift
import Acme

let petId = 987 // Int64 | ID of pet to update
let additionalMetadata = "additionalMetadata_example" // String | Additional data to pass to server (optional)
let file = URL(string: "https://example.com")! // URL | file to upload (optional)

// uploads an image
PetAPI.uploadImage(petId: petId, additionalMetadata: additionalMetadata, file: file) { (response, error) in
    guard error == nil else {
        print(error!)
        return
    }

    if response != nil {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **Int64** | ID of pet to update | 
 **additionalMetadata** | **String** | Additional data to pass to server | [optional] 
 **file** | **URL** | file to upload | [optional] 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

