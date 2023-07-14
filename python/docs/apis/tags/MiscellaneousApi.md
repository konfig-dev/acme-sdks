<a name="__pageTop"></a>
# acme_client.apis.tags.miscellaneous_api.MiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**paginate**](#paginate) | **get** /pagination | Pagination sandbox

# **paginate**

Pagination sandbox

Iterate through a bunch of items

### Example

```python
from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
)

try:
    # Pagination sandbox
    paginate_response = acme.miscellaneous.paginate(
        first=1,  # optional
        after="string_example",  # optional
    )
    pprint(paginate_response.body)
    pprint(paginate_response.body["edges"])
    pprint(paginate_response.body["page_info"])
    pprint(paginate_response.headers)
    pprint(paginate_response.status)
    pprint(paginate_response.round_trip_time)
except ApiException as e:
    print("Exception when calling MiscellaneousApi.paginate: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson, Unset] | optional, default is unset |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**PaginateRequest**](../../models/PaginateRequest.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#paginate.ApiResponseFor200) | OK

#### paginate.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**PaginateResponse**](../../models/PaginateResponse.md) |  | 


### Authorization

No authorization required

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

