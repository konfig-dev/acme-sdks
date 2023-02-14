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

// PaginateRequest struct for PaginateRequest
type PaginateRequest struct {
	First *int32 `json:"first,omitempty"`
	After *string `json:"after,omitempty"`
}

// NewPaginateRequest instantiates a new PaginateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginateRequest() *PaginateRequest {
	this := PaginateRequest{}
	return &this
}

// NewPaginateRequestWithDefaults instantiates a new PaginateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginateRequestWithDefaults() *PaginateRequest {
	this := PaginateRequest{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *PaginateRequest) GetFirst() int32 {
	if o == nil || isNil(o.First) {
		var ret int32
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateRequest) GetFirstOk() (*int32, bool) {
	if o == nil || isNil(o.First) {
    return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *PaginateRequest) HasFirst() bool {
	if o != nil && !isNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given int32 and assigns it to the First field.
func (o *PaginateRequest) SetFirst(v int32) {
	o.First = &v
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *PaginateRequest) GetAfter() string {
	if o == nil || isNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateRequest) GetAfterOk() (*string, bool) {
	if o == nil || isNil(o.After) {
    return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *PaginateRequest) HasAfter() bool {
	if o != nil && !isNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *PaginateRequest) SetAfter(v string) {
	o.After = &v
}

func (o PaginateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !isNil(o.After) {
		toSerialize["after"] = o.After
	}
	return json.Marshal(toSerialize)
}

type NullablePaginateRequest struct {
	value *PaginateRequest
	isSet bool
}

func (v NullablePaginateRequest) Get() *PaginateRequest {
	return v.value
}

func (v *NullablePaginateRequest) Set(val *PaginateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginateRequest(val *PaginateRequest) *NullablePaginateRequest {
	return &NullablePaginateRequest{value: val, isSet: true}
}

func (v NullablePaginateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


