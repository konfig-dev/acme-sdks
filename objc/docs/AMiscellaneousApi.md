# AMiscellaneousApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**paginate**](AMiscellaneousApi.md#paginate) | **GET** /pagination | Pagination sandbox


# **paginate**
```objc
-(NSURLSessionTask*) paginateWithPaginateRequest: (APaginateRequest*) paginateRequest
        completionHandler: (void (^)(APaginateResponse* output, NSError* error)) handler;
```

Pagination sandbox

Iterate through a bunch of items

### Example
```objc

APaginateRequest* paginateRequest = [[APaginateRequest alloc] init]; //  (optional)

AMiscellaneousApi*apiInstance = [[AMiscellaneousApi alloc] init];

// Pagination sandbox
[apiInstance paginateWithPaginateRequest:paginateRequest
          completionHandler: ^(APaginateResponse* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling AMiscellaneousApi->paginate: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginateRequest** | [**APaginateRequest***](APaginateRequest.md)|  | [optional] 

### Return type

[**APaginateResponse***](APaginateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

