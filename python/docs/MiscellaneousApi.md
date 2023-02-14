# acme_client.MiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**paginate**](MiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox


# **paginate**
> PaginateResponse paginate()

Pagination sandbox

Iterate through a bunch of items

### Example


```python
import time
import acme_client
from acme_client.api import miscellaneous_api
from acme_client.model.paginate_request import PaginateRequest
from acme_client.model.paginate_response import PaginateResponse
from pprint import pprint

# Enter a context with an instance of the API client
with acme_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = miscellaneous_api.MiscellaneousApi(api_client)
    paginate_request = PaginateRequest(
        first=1,
        after="after_example",
    ) # PaginateRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Pagination sandbox
        api_response = api_instance.paginate(paginate_request=paginate_request)
        pprint(api_response)
    except acme_client.ApiException as e:
        print("Exception when calling MiscellaneousApi->paginate: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginate_request** | [**PaginateRequest**](PaginateRequest.md)|  | [optional]

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

