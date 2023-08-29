# MiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**paginate**](MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox |


<a name="paginate"></a>
# **paginate**
> PaginateResponse paginate().paginateRequest(paginateRequest).execute();

Pagination sandbox

Iterate through a bunch of items

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.MiscellaneousApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";

    Acme client = new Acme(configuration);
    Integer first = 56;
    String after = "after_example";
    try {
      PaginateResponse result = client
              .miscellaneous
              .paginate()
              .first(first)
              .after(after)
              .execute();
      System.out.println(result);

      System.out.println(result.getEdges());

      System.out.println(result.getPageInfo());

    } catch (ApiException e) {
      System.err.println("Exception when calling MiscellaneousApi#paginate");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<PaginateResponse> response = client
              .miscellaneous
              .paginate()
              .first(first)
              .after(after)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling MiscellaneousApi#paginate");
      System.err.println("Status code: " + e.getStatusCode());
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

