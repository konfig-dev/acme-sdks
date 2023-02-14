# \MiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Paginate**](MiscellaneousApi.md#Paginate) | **Get** /pagination | Pagination sandbox



## Paginate

> PaginateResponse Paginate(ctx).PaginateRequest(paginateRequest).Execute()

Pagination sandbox



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
    paginateRequest := *acme.NewPaginateRequest() // PaginateRequest |  (optional)

    configuration := acme.NewConfiguration()
    apiClient := acme.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscellaneousApi.Paginate(context.Background()).PaginateRequest(paginateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.Paginate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Paginate`: PaginateResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.Paginate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaginateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginateRequest** | [**PaginateRequest**](PaginateRequest.md) |  | 

### Return type

[**PaginateResponse**](PaginateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

