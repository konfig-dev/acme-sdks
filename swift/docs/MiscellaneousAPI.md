# MiscellaneousAPI

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**paginate**](MiscellaneousAPI.md#paginate) | **GET** /pagination | Pagination sandbox


# **paginate**
```swift
    open class func paginate(paginateRequest: PaginateRequest? = nil, completion: @escaping (_ data: PaginateResponse?, _ error: Error?) -> Void)
```

Pagination sandbox

Iterate through a bunch of items

### Example
```swift
import Acme

let paginateRequest = PaginateRequest(first: 123, after: "after_example") // PaginateRequest |  (optional)

// Pagination sandbox
MiscellaneousAPI.paginate(paginateRequest: paginateRequest) { (response, error) in
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
 **paginateRequest** | [**PaginateRequest**](PaginateRequest.md) |  | [optional] 

### Return type

[**PaginateResponse**](PaginateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

