/*
OpenAPI Petstore

This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package acme

import (
	"encoding/json"
)

// PaginateResponse struct for PaginateResponse
type PaginateResponse struct {
	Edges *PaginateResponseEdges `json:"edges,omitempty"`
	PageInfo *PaginateResponsePageInfo `json:"pageInfo,omitempty"`
}

// NewPaginateResponse instantiates a new PaginateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginateResponse() *PaginateResponse {
	this := PaginateResponse{}
	return &this
}

// NewPaginateResponseWithDefaults instantiates a new PaginateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginateResponseWithDefaults() *PaginateResponse {
	this := PaginateResponse{}
	return &this
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *PaginateResponse) GetEdges() PaginateResponseEdges {
	if o == nil || isNil(o.Edges) {
		var ret PaginateResponseEdges
		return ret
	}
	return *o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateResponse) GetEdgesOk() (*PaginateResponseEdges, bool) {
	if o == nil || isNil(o.Edges) {
    return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *PaginateResponse) HasEdges() bool {
	if o != nil && !isNil(o.Edges) {
		return true
	}

	return false
}

// SetEdges gets a reference to the given PaginateResponseEdges and assigns it to the Edges field.
func (o *PaginateResponse) SetEdges(v PaginateResponseEdges) {
	o.Edges = &v
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *PaginateResponse) GetPageInfo() PaginateResponsePageInfo {
	if o == nil || isNil(o.PageInfo) {
		var ret PaginateResponsePageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateResponse) GetPageInfoOk() (*PaginateResponsePageInfo, bool) {
	if o == nil || isNil(o.PageInfo) {
    return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *PaginateResponse) HasPageInfo() bool {
	if o != nil && !isNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PaginateResponsePageInfo and assigns it to the PageInfo field.
func (o *PaginateResponse) SetPageInfo(v PaginateResponsePageInfo) {
	o.PageInfo = &v
}

func (o PaginateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Edges) {
		toSerialize["edges"] = o.Edges
	}
	if !isNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	return json.Marshal(toSerialize)
}

type NullablePaginateResponse struct {
	value *PaginateResponse
	isSet bool
}

func (v NullablePaginateResponse) Get() *PaginateResponse {
	return v.value
}

func (v *NullablePaginateResponse) Set(val *PaginateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginateResponse(val *PaginateResponse) *NullablePaginateResponse {
	return &NullablePaginateResponse{value: val, isSet: true}
}

func (v NullablePaginateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

