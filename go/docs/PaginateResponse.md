# PaginateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edges** | Pointer to [**PaginateResponseEdges**](PaginateResponseEdges.md) |  | [optional] 
**PageInfo** | Pointer to [**PaginateResponsePageInfo**](PaginateResponsePageInfo.md) |  | [optional] 

## Methods

### NewPaginateResponse

`func NewPaginateResponse() *PaginateResponse`

NewPaginateResponse instantiates a new PaginateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginateResponseWithDefaults

`func NewPaginateResponseWithDefaults() *PaginateResponse`

NewPaginateResponseWithDefaults instantiates a new PaginateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdges

`func (o *PaginateResponse) GetEdges() PaginateResponseEdges`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *PaginateResponse) GetEdgesOk() (*PaginateResponseEdges, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *PaginateResponse) SetEdges(v PaginateResponseEdges)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *PaginateResponse) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetPageInfo

`func (o *PaginateResponse) GetPageInfo() PaginateResponsePageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *PaginateResponse) GetPageInfoOk() (*PaginateResponsePageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *PaginateResponse) SetPageInfo(v PaginateResponsePageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *PaginateResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


