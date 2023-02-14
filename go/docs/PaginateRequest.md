# PaginateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to **int32** |  | [optional] 
**After** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginateRequest

`func NewPaginateRequest() *PaginateRequest`

NewPaginateRequest instantiates a new PaginateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginateRequestWithDefaults

`func NewPaginateRequestWithDefaults() *PaginateRequest`

NewPaginateRequestWithDefaults instantiates a new PaginateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *PaginateRequest) GetFirst() int32`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginateRequest) GetFirstOk() (*int32, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginateRequest) SetFirst(v int32)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PaginateRequest) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetAfter

`func (o *PaginateRequest) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *PaginateRequest) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *PaginateRequest) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *PaginateRequest) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


