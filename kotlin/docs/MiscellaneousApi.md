# MiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**paginate**](MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox


<a name="paginate"></a>
# **paginate**
> PaginateResponse paginate(paginateRequest)

Pagination sandbox

Iterate through a bunch of items

### Example
```kotlin
// Import classes:
//import com.acme.client.infrastructure.*
//import com.acme.client.models.*

val apiInstance = MiscellaneousApi()
val paginateRequest : PaginateRequest =  // PaginateRequest | 
try {
    val result : PaginateResponse = apiInstance.paginate(paginateRequest)
    println(result)
} catch (e: ClientException) {
    println("4xx response calling MiscellaneousApi#paginate")
    e.printStackTrace()
} catch (e: ServerException) {
    println("5xx response calling MiscellaneousApi#paginate")
    e.printStackTrace()
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginateRequest** | [**PaginateRequest**](PaginateRequest.md)|  | [optional]

### Return type

[**PaginateResponse**](PaginateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

