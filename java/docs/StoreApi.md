# StoreApi

All URIs are relative to *http://petstore.swagger.io/v2*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteOrder**](StoreApi.md#deleteOrder) | **DELETE** /store/order/{orderId} | Delete purchase order by ID |
| [**getInventory**](StoreApi.md#getInventory) | **GET** /store/inventory | Returns pet inventories by status |
| [**getOrderById**](StoreApi.md#getOrderById) | **GET** /store/order/{orderId} | Find purchase order by ID |
| [**placeOrder**](StoreApi.md#placeOrder) | **POST** /store/order | Place an order for a pet |


<a name="deleteOrder"></a>
# **deleteOrder**
> deleteOrder(orderId).execute();

Delete purchase order by ID

For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.StoreApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";

    Acme client = new Acme(configuration);
    String orderId = "orderId_example"; // ID of the order that needs to be deleted
    try {
      client
              .store
              .deleteOrder(orderId)
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#deleteOrder");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      client
              .store
              .deleteOrder(orderId)
              .executeWithHttpInfo();
    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#deleteOrder");
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
| **orderId** | **String**| ID of the order that needs to be deleted | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid ID supplied |  -  |
| **404** | Order not found |  -  |

<a name="getInventory"></a>
# **getInventory**
> Map&lt;String, Integer&gt; getInventory().execute();

Returns pet inventories by status

Returns a map of status codes to quantities

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.auth.*;
import com.acme.client.model.*;
import com.acme.client.api.StoreApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";
    
    // Configure API key authorization: api_key
    configuration.api_key  = "YOUR API KEY";

    Acme client = new Acme(configuration);
    try {
      Map<String, Integer> result = client
              .store
              .getInventory()
              .execute();
    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#getInventory");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Map<String, Integer>> response = client
              .store
              .getInventory()
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#getInventory");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters
This endpoint does not need any parameter.

### Return type

**Map&lt;String, Integer&gt;**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |

<a name="getOrderById"></a>
# **getOrderById**
> Order getOrderById(orderId).execute();

Find purchase order by ID

For valid response try integer IDs with value &lt;&#x3D; 5 or &gt; 10. Other values will generated exceptions

### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.StoreApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";

    Acme client = new Acme(configuration);
    Long orderId = 56L; // ID of pet that needs to be fetched
    try {
      Order result = client
              .store
              .getOrderById(orderId)
              .execute();
      System.out.println(result);

      System.out.println(result.getId());

      System.out.println(result.getPetId());

      System.out.println(result.getQuantity());

      System.out.println(result.getShipDate());

      System.out.println(result.getStatus());

      System.out.println(result.getComplete());

    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#getOrderById");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Order> response = client
              .store
              .getOrderById(orderId)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#getOrderById");
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
| **orderId** | **Long**| ID of pet that needs to be fetched | |

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |
| **400** | Invalid ID supplied |  -  |
| **404** | Order not found |  -  |

<a name="placeOrder"></a>
# **placeOrder**
> Order placeOrder(order).execute();

Place an order for a pet



### Example
```java
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.ApiResponse;
import com.acme.client.Acme;
import com.acme.client.Configuration;
import com.acme.client.model.*;
import com.acme.client.api.StoreApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "http://petstore.swagger.io/v2";

    Acme client = new Acme(configuration);
    Long id = 56L;
    Long petId = 56L;
    Integer quantity = 56;
    OffsetDateTime shipDate = OffsetDateTime.now();
    String status = "placed"; // Order Status
    Boolean complete = false;
    try {
      Order result = client
              .store
              .placeOrder()
              .id(id)
              .petId(petId)
              .quantity(quantity)
              .shipDate(shipDate)
              .status(status)
              .complete(complete)
              .execute();
      System.out.println(result);

      System.out.println(result.getId());

      System.out.println(result.getPetId());

      System.out.println(result.getQuantity());

      System.out.println(result.getShipDate());

      System.out.println(result.getStatus());

      System.out.println(result.getComplete());

    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#placeOrder");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Order> response = client
              .store
              .placeOrder()
              .id(id)
              .petId(petId)
              .quantity(quantity)
              .shipDate(shipDate)
              .status(status)
              .complete(complete)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling StoreApi#placeOrder");
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
| **order** | [**Order**](Order.md)| order placed for purchasing the pet | |

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/xml, application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | successful operation |  -  |
| **400** | Invalid Order |  -  |

