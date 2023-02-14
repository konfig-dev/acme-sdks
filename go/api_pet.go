/*
OpenAPI Petstore

This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package acme

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// PetApiService PetApi service
type PetApiService service

type PetApiAddRequest struct {
	ctx context.Context
	ApiService *PetApiService
	pet *Pet
}

// Pet object that needs to be added to the store
func (r PetApiAddRequest) Pet(pet Pet) PetApiAddRequest {
	r.pet = &pet
	return r
}

func (r PetApiAddRequest) Execute() (*Pet, *http.Response, error) {
	return r.ApiService.AddExecute(r)
}

/*
Add Add a new pet to the store



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PetApiAddRequest
*/
func (a *PetApiService) Add() PetApiAddRequest {
	return PetApiAddRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
	}
}

// Execute executes the request
//  @return Pet
func (a *PetApiService) AddExecute(r PetApiAddRequest) (*Pet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.Add")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pet == nil {
		return localVarReturnValue, nil, reportError("pet is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.pet
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PetApiDeleteRequest struct {
	ctx context.Context
	ApiService *PetApiService
	petId int64
}

func (r PetApiDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Deletes a pet



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param petId Pet id to delete
 @return PetApiDeleteRequest
*/
func (a *PetApiService) Delete(petId int64) PetApiDeleteRequest {
	return PetApiDeleteRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		petId: petId,
	}
}

// Execute executes the request
func (a *PetApiService) DeleteExecute(r PetApiDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.Delete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", url.PathEscape(parameterToString(r.petId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api_key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PetApiFindByStatusRequest struct {
	ctx context.Context
	ApiService *PetApiService
	status *[]string
}

// Status values that need to be considered for filter
// Deprecated
func (r PetApiFindByStatusRequest) Status(status []string) PetApiFindByStatusRequest {
	r.status = &status
	return r
}

func (r PetApiFindByStatusRequest) Execute() ([]Pet, *http.Response, error) {
	return r.ApiService.FindByStatusExecute(r)
}

/*
FindByStatus Finds Pets by status

Multiple status values can be provided with comma separated strings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PetApiFindByStatusRequest
*/
func (a *PetApiService) FindByStatus() PetApiFindByStatusRequest {
	return PetApiFindByStatusRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
	}
}

// Execute executes the request
//  @return []Pet
func (a *PetApiService) FindByStatusExecute(r PetApiFindByStatusRequest) ([]Pet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.FindByStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/findByStatus"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.status == nil {
		return localVarReturnValue, nil, reportError("status is required and must be specified")
	}

	localVarQueryParams.Add("status", parameterToString(*r.status, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PetApiFindByTagsRequest struct {
	ctx context.Context
	ApiService *PetApiService
	tags *[]string
}

// Tags to filter by
func (r PetApiFindByTagsRequest) Tags(tags []string) PetApiFindByTagsRequest {
	r.tags = &tags
	return r
}

func (r PetApiFindByTagsRequest) Execute() ([]Pet, *http.Response, error) {
	return r.ApiService.FindByTagsExecute(r)
}

/*
FindByTags Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PetApiFindByTagsRequest

Deprecated
*/
func (a *PetApiService) FindByTags() PetApiFindByTagsRequest {
	return PetApiFindByTagsRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
	}
}

// Execute executes the request
//  @return []Pet
// Deprecated
func (a *PetApiService) FindByTagsExecute(r PetApiFindByTagsRequest) ([]Pet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.FindByTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/findByTags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tags == nil {
		return localVarReturnValue, nil, reportError("tags is required and must be specified")
	}

	localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PetApiGetByIdRequest struct {
	ctx context.Context
	ApiService *PetApiService
	petId int64
}

func (r PetApiGetByIdRequest) Execute() (*Pet, *http.Response, error) {
	return r.ApiService.GetByIdExecute(r)
}

/*
GetById Find pet by ID

Returns a single pet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param petId ID of pet to return
 @return PetApiGetByIdRequest
*/
func (a *PetApiService) GetById(petId int64) PetApiGetByIdRequest {
	return PetApiGetByIdRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		petId: petId,
	}
}

// Execute executes the request
//  @return Pet
func (a *PetApiService) GetByIdExecute(r PetApiGetByIdRequest) (*Pet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.GetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", url.PathEscape(parameterToString(r.petId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api_key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PetApiUpdateRequest struct {
	ctx context.Context
	ApiService *PetApiService
	pet *Pet
}

// Pet object that needs to be added to the store
func (r PetApiUpdateRequest) Pet(pet Pet) PetApiUpdateRequest {
	r.pet = &pet
	return r
}

func (r PetApiUpdateRequest) Execute() (*Pet, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update an existing pet



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PetApiUpdateRequest
*/
func (a *PetApiService) Update() PetApiUpdateRequest {
	return PetApiUpdateRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
	}
}

// Execute executes the request
//  @return Pet
func (a *PetApiService) UpdateExecute(r PetApiUpdateRequest) (*Pet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pet == nil {
		return localVarReturnValue, nil, reportError("pet is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.pet
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PetApiUpdateWithFormRequest struct {
	ctx context.Context
	ApiService *PetApiService
	petId int64
	name *string
	status *string
}

// Updated name of the pet
func (r PetApiUpdateWithFormRequest) Name(name string) PetApiUpdateWithFormRequest {
	r.name = &name
	return r
}

// Updated status of the pet
func (r PetApiUpdateWithFormRequest) Status(status string) PetApiUpdateWithFormRequest {
	r.status = &status
	return r
}

func (r PetApiUpdateWithFormRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateWithFormExecute(r)
}

/*
UpdateWithForm Updates a pet in the store with form data



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param petId ID of pet that needs to be updated
 @return PetApiUpdateWithFormRequest
*/
func (a *PetApiService) UpdateWithForm(petId int64) PetApiUpdateWithFormRequest {
	return PetApiUpdateWithFormRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		petId: petId,
	}
}

// Execute executes the request
func (a *PetApiService) UpdateWithFormExecute(r PetApiUpdateWithFormRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UpdateWithForm")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", url.PathEscape(parameterToString(r.petId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.status != nil {
		localVarFormParams.Add("status", parameterToString(*r.status, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PetApiUploadImageRequest struct {
	ctx context.Context
	ApiService *PetApiService
	petId int64
	additionalMetadata *string
	file **os.File
}

// Additional data to pass to server
func (r PetApiUploadImageRequest) AdditionalMetadata(additionalMetadata string) PetApiUploadImageRequest {
	r.additionalMetadata = &additionalMetadata
	return r
}

// file to upload
func (r PetApiUploadImageRequest) File(file *os.File) PetApiUploadImageRequest {
	r.file = &file
	return r
}

func (r PetApiUploadImageRequest) Execute() (*ApiResponse, *http.Response, error) {
	return r.ApiService.UploadImageExecute(r)
}

/*
UploadImage uploads an image



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param petId ID of pet to update
 @return PetApiUploadImageRequest
*/
func (a *PetApiService) UploadImage(petId int64) PetApiUploadImageRequest {
	return PetApiUploadImageRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		petId: petId,
	}
}

// Execute executes the request
//  @return ApiResponse
func (a *PetApiService) UploadImageExecute(r PetApiUploadImageRequest) (*ApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UploadImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}/uploadImage"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", url.PathEscape(parameterToString(r.petId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.additionalMetadata != nil {
		localVarFormParams.Add("additionalMetadata", parameterToString(*r.additionalMetadata, ""))
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	var fileLocalVarFile *os.File
	if r.file != nil {
		fileLocalVarFile = *r.file
	}
	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
