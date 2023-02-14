# APetApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add**](APetApi.md#add) | **POST** /pet | Add a new pet to the store
[**delete**](APetApi.md#delete) | **DELETE** /pet/{petId} | Deletes a pet
[**findByStatus**](APetApi.md#findbystatus) | **GET** /pet/findByStatus | Finds Pets by status
[**findByTags**](APetApi.md#findbytags) | **GET** /pet/findByTags | Finds Pets by tags
[**getById**](APetApi.md#getbyid) | **GET** /pet/{petId} | Find pet by ID
[**update**](APetApi.md#update) | **PUT** /pet | Update an existing pet
[**updateWithForm**](APetApi.md#updatewithform) | **POST** /pet/{petId} | Updates a pet in the store with form data
[**uploadImage**](APetApi.md#uploadimage) | **POST** /pet/{petId}/uploadImage | uploads an image


# **add**
```objc
-(NSURLSessionTask*) addWithPet: (APet*) pet
        completionHandler: (void (^)(APet* output, NSError* error)) handler;
```

Add a new pet to the store



### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure OAuth2 access token for authorization: (authentication scheme: petstore_auth)
[apiConfig setAccessToken:@"YOUR_ACCESS_TOKEN"];


APet* pet = [[APet alloc] init]; // Pet object that needs to be added to the store

APetApi*apiInstance = [[APetApi alloc] init];

// Add a new pet to the store
[apiInstance addWithPet:pet
          completionHandler: ^(APet* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling APetApi->add: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**APet***](APet.md)| Pet object that needs to be added to the store | 

### Return type

[**APet***](APet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete**
```objc
-(NSURLSessionTask*) deleteWithPetId: (NSNumber*) petId
        completionHandler: (void (^)(NSError* error)) handler;
```

Deletes a pet



### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: api_key)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"api_key"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"api_key"];

// Configure OAuth2 access token for authorization: (authentication scheme: petstore_auth)
[apiConfig setAccessToken:@"YOUR_ACCESS_TOKEN"];


NSNumber* petId = @56; // Pet id to delete

APetApi*apiInstance = [[APetApi alloc] init];

// Deletes a pet
[apiInstance deleteWithPetId:petId
          completionHandler: ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling APetApi->delete: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **NSNumber***| Pet id to delete | 

### Return type

void (empty response body)

### Authorization

[api_key](../README.md#api_key), [petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **findByStatus**
```objc
-(NSURLSessionTask*) findByStatusWithStatus: (NSArray<NSString*>*) status
        completionHandler: (void (^)(NSArray<APet>* output, NSError* error)) handler;
```

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure OAuth2 access token for authorization: (authentication scheme: petstore_auth)
[apiConfig setAccessToken:@"YOUR_ACCESS_TOKEN"];


NSArray<NSString*>* status = @[@"status_example"]; // Status values that need to be considered for filter

APetApi*apiInstance = [[APetApi alloc] init];

// Finds Pets by status
[apiInstance findByStatusWithStatus:status
          completionHandler: ^(NSArray<APet>* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling APetApi->findByStatus: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Status values that need to be considered for filter | 

### Return type

[**NSArray<APet>***](APet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **findByTags**
```objc
-(NSURLSessionTask*) findByTagsWithTags: (NSArray<NSString*>*) tags
        completionHandler: (void (^)(NSArray<APet>* output, NSError* error)) handler;
```

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure OAuth2 access token for authorization: (authentication scheme: petstore_auth)
[apiConfig setAccessToken:@"YOUR_ACCESS_TOKEN"];


NSArray<NSString*>* tags = @[@"tags_example"]; // Tags to filter by

APetApi*apiInstance = [[APetApi alloc] init];

// Finds Pets by tags
[apiInstance findByTagsWithTags:tags
          completionHandler: ^(NSArray<APet>* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling APetApi->findByTags: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Tags to filter by | 

### Return type

[**NSArray<APet>***](APet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getById**
```objc
-(NSURLSessionTask*) getByIdWithPetId: (NSNumber*) petId
        completionHandler: (void (^)(APet* output, NSError* error)) handler;
```

Find pet by ID

Returns a single pet

### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: api_key)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"api_key"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"api_key"];


NSNumber* petId = @56; // ID of pet to return

APetApi*apiInstance = [[APetApi alloc] init];

// Find pet by ID
[apiInstance getByIdWithPetId:petId
          completionHandler: ^(APet* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling APetApi->getById: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **NSNumber***| ID of pet to return | 

### Return type

[**APet***](APet.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update**
```objc
-(NSURLSessionTask*) updateWithPet: (APet*) pet
        completionHandler: (void (^)(APet* output, NSError* error)) handler;
```

Update an existing pet



### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure OAuth2 access token for authorization: (authentication scheme: petstore_auth)
[apiConfig setAccessToken:@"YOUR_ACCESS_TOKEN"];


APet* pet = [[APet alloc] init]; // Pet object that needs to be added to the store

APetApi*apiInstance = [[APetApi alloc] init];

// Update an existing pet
[apiInstance updateWithPet:pet
          completionHandler: ^(APet* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling APetApi->update: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**APet***](APet.md)| Pet object that needs to be added to the store | 

### Return type

[**APet***](APet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateWithForm**
```objc
-(NSURLSessionTask*) updateWithFormWithPetId: (NSNumber*) petId
    name: (NSString*) name
    status: (NSString*) status
        completionHandler: (void (^)(NSError* error)) handler;
```

Updates a pet in the store with form data



### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure OAuth2 access token for authorization: (authentication scheme: petstore_auth)
[apiConfig setAccessToken:@"YOUR_ACCESS_TOKEN"];


NSNumber* petId = @56; // ID of pet that needs to be updated
NSString* name = @"name_example"; // Updated name of the pet (optional)
NSString* status = @"status_example"; // Updated status of the pet (optional)

APetApi*apiInstance = [[APetApi alloc] init];

// Updates a pet in the store with form data
[apiInstance updateWithFormWithPetId:petId
              name:name
              status:status
          completionHandler: ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling APetApi->updateWithForm: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **NSNumber***| ID of pet that needs to be updated | 
 **name** | **NSString***| Updated name of the pet | [optional] 
 **status** | **NSString***| Updated status of the pet | [optional] 

### Return type

void (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **uploadImage**
```objc
-(NSURLSessionTask*) uploadImageWithPetId: (NSNumber*) petId
    additionalMetadata: (NSString*) additionalMetadata
    file: (NSURL*) file
        completionHandler: (void (^)(AApiResponse* output, NSError* error)) handler;
```

uploads an image



### Example
```objc
ADefaultConfiguration *apiConfig = [ADefaultConfiguration sharedConfig];

// Configure OAuth2 access token for authorization: (authentication scheme: petstore_auth)
[apiConfig setAccessToken:@"YOUR_ACCESS_TOKEN"];


NSNumber* petId = @56; // ID of pet to update
NSString* additionalMetadata = @"additionalMetadata_example"; // Additional data to pass to server (optional)
NSURL* file = [NSURL fileURLWithPath:@"/path/to/file"]; // file to upload (optional)

APetApi*apiInstance = [[APetApi alloc] init];

// uploads an image
[apiInstance uploadImageWithPetId:petId
              additionalMetadata:additionalMetadata
              file:file
          completionHandler: ^(AApiResponse* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling APetApi->uploadImage: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **petId** | **NSNumber***| ID of pet to update | 
 **additionalMetadata** | **NSString***| Additional data to pass to server | [optional] 
 **file** | **NSURL*****NSURL***| file to upload | [optional] 

### Return type

[**AApiResponse***](AApiResponse.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

