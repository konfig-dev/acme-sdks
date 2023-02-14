# \PetApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](PetApi.md#Add) | **Post** /pet | Add a new pet to the store
[**Delete**](PetApi.md#Delete) | **Delete** /pet/{petId} | Deletes a pet
[**FindByStatus**](PetApi.md#FindByStatus) | **Get** /pet/findByStatus | Finds Pets by status
[**FindByTags**](PetApi.md#FindByTags) | **Get** /pet/findByTags | Finds Pets by tags
[**GetById**](PetApi.md#GetById) | **Get** /pet/{petId} | Find pet by ID
[**Update**](PetApi.md#Update) | **Put** /pet | Update an existing pet
[**UpdateWithForm**](PetApi.md#UpdateWithForm) | **Post** /pet/{petId} | Updates a pet in the store with form data
[**UploadImage**](PetApi.md#UploadImage) | **Post** /pet/{petId}/uploadImage | uploads an image



## Add

> Pet Add(ctx).Pet(pet).Execute()

Add a new pet to the store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    pet := *acme.NewPet("doggie", []string{"PhotoUrls_example"}) // Pet | Pet object that needs to be added to the store

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.Add(context.Background()).Pet(pet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.Add``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Add`: Pet
    fmt.Fprintf(os.Stdout, "Response from `PetApi.Add`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, petId).Execute()

Deletes a pet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    petId := int64(789) // int64 | Pet id to delete

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.Delete(context.Background(), petId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**petId** | **int64** | Pet id to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindByStatus

> []Pet FindByStatus(ctx).Status(status).Execute()

Finds Pets by status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    status := []string{"Status_example"} // []string | Status values that need to be considered for filter

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.FindByStatus(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.FindByStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindByStatus`: []Pet
    fmt.Fprintf(os.Stdout, "Response from `PetApi.FindByStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | Status values that need to be considered for filter | 

### Return type

[**[]Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindByTags

> []Pet FindByTags(ctx).Tags(tags).Execute()

Finds Pets by tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    tags := []string{"Inner_example"} // []string | Tags to filter by

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.FindByTags(context.Background()).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.FindByTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindByTags`: []Pet
    fmt.Fprintf(os.Stdout, "Response from `PetApi.FindByTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindByTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **[]string** | Tags to filter by | 

### Return type

[**[]Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetById

> Pet GetById(ctx, petId).Execute()

Find pet by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    petId := int64(789) // int64 | ID of pet to return

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.GetById(context.Background(), petId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.GetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetById`: Pet
    fmt.Fprintf(os.Stdout, "Response from `PetApi.GetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**petId** | **int64** | ID of pet to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Pet**](Pet.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Pet Update(ctx).Pet(pet).Execute()

Update an existing pet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    pet := *acme.NewPet("doggie", []string{"PhotoUrls_example"}) // Pet | Pet object that needs to be added to the store

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.Update(context.Background()).Pet(pet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Pet
    fmt.Fprintf(os.Stdout, "Response from `PetApi.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWithForm

> UpdateWithForm(ctx, petId).Name(name).Status(status).Execute()

Updates a pet in the store with form data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    petId := int64(789) // int64 | ID of pet that needs to be updated
    name := "name_example" // string | Updated name of the pet (optional)
    status := "status_example" // string | Updated status of the pet (optional)

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.UpdateWithForm(context.Background(), petId).Name(name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.UpdateWithForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**petId** | **int64** | ID of pet that needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWithFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Updated name of the pet | 
 **status** | **string** | Updated status of the pet | 

### Return type

 (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImage

> ApiResponse UploadImage(ctx, petId).AdditionalMetadata(additionalMetadata).File(file).Execute()

uploads an image



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func main() {
    petId := int64(789) // int64 | ID of pet to update
    additionalMetadata := "additionalMetadata_example" // string | Additional data to pass to server (optional)
    file := os.NewFile(1234, "some_file") // *os.File | file to upload (optional)

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.PetApi.UploadImage(context.Background(), petId).AdditionalMetadata(additionalMetadata).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetApi.UploadImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadImage`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PetApi.UploadImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**petId** | **int64** | ID of pet to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **additionalMetadata** | **string** | Additional data to pass to server | 
 **file** | ***os.File** | file to upload | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

