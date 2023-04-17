#import <Foundation/Foundation.h>
#import "APaginateRequest.h"
#import "APaginateResponse.h"
#import "AApi.h"

/**
* OpenAPI Petstore
* This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
*
* The version of the OpenAPI document: 1.0.0
* 
*
* NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
* https://openapi-generator.tech
* Do not edit the class manually.
*/



@interface AMiscellaneousApi: NSObject <AApi>

extern NSString* kAMiscellaneousApiErrorDomain;
extern NSInteger kAMiscellaneousApiMissingParamErrorCode;

-(instancetype) initWithApiClient:(AApiClient *)apiClient NS_DESIGNATED_INITIALIZER;

/// Pagination sandbox
/// Iterate through a bunch of items
///
/// @param paginateRequest  (optional)
/// 
///  code:200 message:"OK"
///
/// @return APaginateResponse*
-(NSURLSessionTask*) paginateWithPaginateRequest: (APaginateRequest*) paginateRequest
    completionHandler: (void (^)(APaginateResponse* output, NSError* error)) handler;



@end