# PaginateResponsePageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNextPage** | Pointer to **bool** |  | [optional] 
**HasPreviousPage** | Pointer to **bool** |  | [optional] 
**StartCursor** | Pointer to **string** |  | [optional] 
**EndCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginateResponsePageInfo

`func NewPaginateResponsePageInfo() *PaginateResponsePageInfo`

NewPaginateResponsePageInfo instantiates a new PaginateResponsePageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginateResponsePageInfoWithDefaults

`func NewPaginateResponsePageInfoWithDefaults() *PaginateResponsePageInfo`

NewPaginateResponsePageInfoWithDefaults instantiates a new PaginateResponsePageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNextPage

`func (o *PaginateResponsePageInfo) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *PaginateResponsePageInfo) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *PaginateResponsePageInfo) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.

### HasHasNextPage

`func (o *PaginateResponsePageInfo) HasHasNextPage() bool`

HasHasNextPage returns a boolean if a field has been set.

### GetHasPreviousPage

`func (o *PaginateResponsePageInfo) GetHasPreviousPage() bool`

GetHasPreviousPage returns the HasPreviousPage field if non-nil, zero value otherwise.

### GetHasPreviousPageOk

`func (o *PaginateResponsePageInfo) GetHasPreviousPageOk() (*bool, bool)`

GetHasPreviousPageOk returns a tuple with the HasPreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreviousPage

`func (o *PaginateResponsePageInfo) SetHasPreviousPage(v bool)`

SetHasPreviousPage sets HasPreviousPage field to given value.

### HasHasPreviousPage

`func (o *PaginateResponsePageInfo) HasHasPreviousPage() bool`

HasHasPreviousPage returns a boolean if a field has been set.

### GetStartCursor

`func (o *PaginateResponsePageInfo) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *PaginateResponsePageInfo) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *PaginateResponsePageInfo) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *PaginateResponsePageInfo) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *PaginateResponsePageInfo) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *PaginateResponsePageInfo) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *PaginateResponsePageInfo) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *PaginateResponsePageInfo) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

