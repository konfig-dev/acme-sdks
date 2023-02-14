# Acme.Net.Api.MiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**Paginate**](MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox |

<a name="paginate"></a>
# **Paginate**
> PaginateResponse Paginate (PaginateRequest paginateRequest = null)

Pagination sandbox

Iterate through a bunch of items

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Acme.Net.Api;
using Acme.Net.Client;
using Acme.Net.Model;

namespace Example
{
    public class PaginateExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://petstore.swagger.io/v2";
            var apiInstance = new MiscellaneousApi(config);
            var paginateRequest = new PaginateRequest(); // PaginateRequest |  (optional) 

            try
            {
                // Pagination sandbox
                PaginateResponse result = apiInstance.Paginate(paginateRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling MiscellaneousApi.Paginate: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the PaginateWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Pagination sandbox
    ApiResponse<PaginateResponse> response = apiInstance.PaginateWithHttpInfo(paginateRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling MiscellaneousApi.PaginateWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **paginateRequest** | [**PaginateRequest**](PaginateRequest.md) |  | [optional]  |

### Return type

[**PaginateResponse**](PaginateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

