# UserAPI

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](UserAPI.md#create) | **POST** /user | Create user
[**createWithArray**](UserAPI.md#createwitharray) | **POST** /user/createWithArray | Creates list of users with given input array
[**createWithList**](UserAPI.md#createwithlist) | **POST** /user/createWithList | Creates list of users with given input array
[**delete**](UserAPI.md#delete) | **DELETE** /user/{username} | Delete user
[**getByName**](UserAPI.md#getbyname) | **GET** /user/{username} | Get user by user name
[**login**](UserAPI.md#login) | **GET** /user/login | Logs user into the system
[**logout**](UserAPI.md#logout) | **GET** /user/logout | Logs out current logged in user session
[**update**](UserAPI.md#update) | **PUT** /user/{username} | Updated user


# **create**
```swift
    open class func create(user: User, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Create user

This can only be done by the logged in user.

### Example
```swift
import Acme

let user = User(id: 123, username: "username_example", firstName: "firstName_example", lastName: "lastName_example", email: "email_example", password: "password_example", phone: "phone_example", userStatus: 123) // User | Created user object

// Create user
UserAPI.create(user: user) { (response, error) in
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
 **user** | [**User**](User.md) | Created user object | 

### Return type

Void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **createWithArray**
```swift
    open class func createWithArray(user: [User], completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Creates list of users with given input array



### Example
```swift
import Acme

let user = [User(id: 123, username: "username_example", firstName: "firstName_example", lastName: "lastName_example", email: "email_example", password: "password_example", phone: "phone_example", userStatus: 123)] // [User] | List of user object

// Creates list of users with given input array
UserAPI.createWithArray(user: user) { (response, error) in
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
 **user** | [**[User]**](User.md) | List of user object | 

### Return type

Void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **createWithList**
```swift
    open class func createWithList(user: [User], completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Creates list of users with given input array



### Example
```swift
import Acme

let user = [User(id: 123, username: "username_example", firstName: "firstName_example", lastName: "lastName_example", email: "email_example", password: "password_example", phone: "phone_example", userStatus: 123)] // [User] | List of user object

// Creates list of users with given input array
UserAPI.createWithList(user: user) { (response, error) in
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
 **user** | [**[User]**](User.md) | List of user object | 

### Return type

Void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **delete**
```swift
    open class func delete(username: String, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Delete user

This can only be done by the logged in user.

### Example
```swift
import Acme

let username = "username_example" // String | The name that needs to be deleted

// Delete user
UserAPI.delete(username: username) { (response, error) in
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
 **username** | **String** | The name that needs to be deleted | 

### Return type

Void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **getByName**
```swift
    open class func getByName(username: String, completion: @escaping (_ data: User?, _ error: Error?) -> Void)
```

Get user by user name



### Example
```swift
import Acme

let username = "username_example" // String | The name that needs to be fetched. Use user1 for testing.

// Get user by user name
UserAPI.getByName(username: username) { (response, error) in
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
 **username** | **String** | The name that needs to be fetched. Use user1 for testing. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **login**
```swift
    open class func login(username: String, password: String, completion: @escaping (_ data: String?, _ error: Error?) -> Void)
```

Logs user into the system



### Example
```swift
import Acme

let username = "username_example" // String | The user name for login
let password = "password_example" // String | The password for login in clear text

// Logs user into the system
UserAPI.login(username: username, password: password) { (response, error) in
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
 **username** | **String** | The user name for login | 
 **password** | **String** | The password for login in clear text | 

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **logout**
```swift
    open class func logout(completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Logs out current logged in user session



### Example
```swift
import Acme


// Logs out current logged in user session
UserAPI.logout() { (response, error) in
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
This endpoint does not need any parameter.

### Return type

Void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

# **update**
```swift
    open class func update(username: String, user: User, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

Updated user

This can only be done by the logged in user.

### Example
```swift
import Acme

let username = "username_example" // String | name that need to be deleted
let user = User(id: 123, username: "username_example", firstName: "firstName_example", lastName: "lastName_example", email: "email_example", password: "password_example", phone: "phone_example", userStatus: 123) // User | Updated user object

// Updated user
UserAPI.update(username: username, user: user) { (response, error) in
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
 **username** | **String** | name that need to be deleted | 
 **user** | [**User**](User.md) | Updated user object | 

### Return type

Void (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

