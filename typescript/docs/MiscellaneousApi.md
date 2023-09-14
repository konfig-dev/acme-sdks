# MiscellaneousApi

All URIs are relative to *https://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**paginate**](MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox


# **paginate**

#### **GET** /pagination

### Description
Iterate through a bunch of items

### Example


```typescript
import { Acme } from "acme-typescript-sdk";

const acme = new Acme({
  // Defining the base path is optional and defaults to https://petstore.swagger.io/v2
  // basePath: "https://petstore.swagger.io/v2",
});

const paginateResponse = await acme.miscellaneous.paginate({});

console.log(paginateResponse);
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginateRequest** | **PaginateRequest**|  |


### Return type

**PaginateResponse**

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


