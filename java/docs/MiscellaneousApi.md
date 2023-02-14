# MiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**paginate**](MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox |


<a name="paginate"></a>
# **paginate**
> PaginateResponse paginate(paginateRequest)

Pagination sandbox

Iterate through a bunch of items

### Example
```java
// Import classes:
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.MiscellaneousApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://petstore.swagger.io/v2");

    MiscellaneousApi apiInstance = new MiscellaneousApi(defaultClient);
    PaginateRequest paginateRequest = new PaginateRequest(); // PaginateRequest | 
    try {
      PaginateResponse result = apiInstance.paginate(paginateRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling MiscellaneousApi#paginate");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **paginateRequest** | [**PaginateRequest**](PaginateRequest.md)|  | [optional] |

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

