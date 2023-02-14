# \UserApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserApi.md#Create) | **Post** /user | Create user
[**CreateWithArray**](UserApi.md#CreateWithArray) | **Post** /user/createWithArray | Creates list of users with given input array
[**CreateWithList**](UserApi.md#CreateWithList) | **Post** /user/createWithList | Creates list of users with given input array
[**Delete**](UserApi.md#Delete) | **Delete** /user/{username} | Delete user
[**GetByName**](UserApi.md#GetByName) | **Get** /user/{username} | Get user by user name
[**Login**](UserApi.md#Login) | **Get** /user/login | Logs user into the system
[**Logout**](UserApi.md#Logout) | **Get** /user/logout | Logs out current logged in user session
[**Update**](UserApi.md#Update) | **Put** /user/{username} | Updated user



## Create

> Create(ctx).User(user).Execute()

Create user



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
    user := *acme.NewUser() // User | Created user object

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.Create(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) | Created user object | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWithArray

> CreateWithArray(ctx).User(user).Execute()

Creates list of users with given input array



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
    user := []acme.User{*acme.NewUser()} // []User | List of user object

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.CreateWithArray(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateWithArray``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWithArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**[]User**](User.md) | List of user object | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWithList

> CreateWithList(ctx).User(user).Execute()

Creates list of users with given input array



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
    user := []acme.User{*acme.NewUser()} // []User | List of user object

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.CreateWithList(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateWithList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWithListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**[]User**](User.md) | List of user object | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, username).Execute()

Delete user



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
    username := "username_example" // string | The name that needs to be deleted

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.Delete(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The name that needs to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByName

> User GetByName(ctx, username).Execute()

Get user by user name



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
    username := "username_example" // string | The name that needs to be fetched. Use user1 for testing.

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetByName(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByName`: User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The name that needs to be fetched. Use user1 for testing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> string Login(ctx).Username(username).Password(password).Execute()

Logs user into the system



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
    username := "username_example" // string | The user name for login
    password := "password_example" // string | The password for login in clear text

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.Login(context.Background()).Username(username).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | The user name for login | 
 **password** | **string** | The password for login in clear text | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx).Execute()

Logs out current logged in user session



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

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, username).User(user).Execute()

Updated user



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
    username := "username_example" // string | name that need to be deleted
    user := *acme.NewUser() // User | Updated user object

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.Update(context.Background(), username).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | name that need to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | Updated user object | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

