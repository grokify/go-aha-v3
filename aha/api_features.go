/*
Aha.io API

Articles that matter on social publishing platform

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aha

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FeaturesAPIService FeaturesAPI service
type FeaturesAPIService service

type ApiGetFeatureRequest struct {
	ctx        context.Context
	ApiService *FeaturesAPIService
	featureId  string
}

func (r ApiGetFeatureRequest) Execute() (*FeatureWrap, *http.Response, error) {
	return r.ApiService.GetFeatureExecute(r)
}

/*
GetFeature

Get a specific feature

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param featureId Numeric ID, or key of the feature to be retrieved
	@return ApiGetFeatureRequest
*/
func (a *FeaturesAPIService) GetFeature(ctx context.Context, featureId string) ApiGetFeatureRequest {
	return ApiGetFeatureRequest{
		ApiService: a,
		ctx:        ctx,
		featureId:  featureId,
	}
}

// Execute executes the request
//
//	@return FeatureWrap
func (a *FeaturesAPIService) GetFeatureExecute(r ApiGetFeatureRequest) (*FeatureWrap, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FeatureWrap
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeaturesAPIService.GetFeature")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/features/{feature_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"feature_id"+"}", url.PathEscape(parameterValueToString(r.featureId, "featureId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetFeaturesRequest struct {
	ctx            context.Context
	ApiService     *FeaturesAPIService
	q              *string
	updatedSince   *time.Time
	tag            *string
	assignedToUser *string
	page           *int32
	perPage        *int32
}

// Sub-string to match against feature name or ID
func (r ApiGetFeaturesRequest) Q(q string) ApiGetFeaturesRequest {
	r.q = &q
	return r
}

// UTC timestamp (in ISO8601 format) that the updated_at field must be larger than.
func (r ApiGetFeaturesRequest) UpdatedSince(updatedSince time.Time) ApiGetFeaturesRequest {
	r.updatedSince = &updatedSince
	return r
}

// A string tag value.
func (r ApiGetFeaturesRequest) Tag(tag string) ApiGetFeaturesRequest {
	r.tag = &tag
	return r
}

// The ID or email address of user to return assigned features for.
func (r ApiGetFeaturesRequest) AssignedToUser(assignedToUser string) ApiGetFeaturesRequest {
	r.assignedToUser = &assignedToUser
	return r
}

// A specific page of results.
func (r ApiGetFeaturesRequest) Page(page int32) ApiGetFeaturesRequest {
	r.page = &page
	return r
}

// Number of results per page.
func (r ApiGetFeaturesRequest) PerPage(perPage int32) ApiGetFeaturesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetFeaturesRequest) Execute() (*FeaturesResponse, *http.Response, error) {
	return r.ApiService.GetFeaturesExecute(r)
}

/*
GetFeatures Get all features

Get all features

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetFeaturesRequest
*/
func (a *FeaturesAPIService) GetFeatures(ctx context.Context) ApiGetFeaturesRequest {
	return ApiGetFeaturesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FeaturesResponse
func (a *FeaturesAPIService) GetFeaturesExecute(r ApiGetFeaturesRequest) (*FeaturesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FeaturesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeaturesAPIService.GetFeatures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/features"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.updatedSince != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updated_since", r.updatedSince, "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.assignedToUser != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assigned_to_user", r.assignedToUser, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetReleaseFeaturesRequest struct {
	ctx        context.Context
	ApiService *FeaturesAPIService
	releaseId  string
	page       *int32
	perPage    *int32
}

// A specific page of results.
func (r ApiGetReleaseFeaturesRequest) Page(page int32) ApiGetReleaseFeaturesRequest {
	r.page = &page
	return r
}

// Number of results per page.
func (r ApiGetReleaseFeaturesRequest) PerPage(perPage int32) ApiGetReleaseFeaturesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetReleaseFeaturesRequest) Execute() (*FeaturesResponse, *http.Response, error) {
	return r.ApiService.GetReleaseFeaturesExecute(r)
}

/*
GetReleaseFeatures Get all features for a release

Get all features for a release

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param releaseId Numeric ID, or key of the release to retrieve features for
	@return ApiGetReleaseFeaturesRequest
*/
func (a *FeaturesAPIService) GetReleaseFeatures(ctx context.Context, releaseId string) ApiGetReleaseFeaturesRequest {
	return ApiGetReleaseFeaturesRequest{
		ApiService: a,
		ctx:        ctx,
		releaseId:  releaseId,
	}
}

// Execute executes the request
//
//	@return FeaturesResponse
func (a *FeaturesAPIService) GetReleaseFeaturesExecute(r ApiGetReleaseFeaturesRequest) (*FeaturesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FeaturesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeaturesAPIService.GetReleaseFeatures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/releases/{release_id}/features"
	localVarPath = strings.Replace(localVarPath, "{"+"release_id"+"}", url.PathEscape(parameterValueToString(r.releaseId, "releaseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateFeatureRequest struct {
	ctx           context.Context
	ApiService    *FeaturesAPIService
	featureId     string
	featureUpdate *FeatureUpdate
}

// Feature properties to update
func (r ApiUpdateFeatureRequest) FeatureUpdate(featureUpdate FeatureUpdate) ApiUpdateFeatureRequest {
	r.featureUpdate = &featureUpdate
	return r
}

func (r ApiUpdateFeatureRequest) Execute() (*FeatureWrap, *http.Response, error) {
	return r.ApiService.UpdateFeatureExecute(r)
}

/*
UpdateFeature Update a feature's custom fields with tag-like value

Update a feature's custom fields with tag-like value

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param featureId Numeric ID, or key of the feature to be retrieved
	@return ApiUpdateFeatureRequest
*/
func (a *FeaturesAPIService) UpdateFeature(ctx context.Context, featureId string) ApiUpdateFeatureRequest {
	return ApiUpdateFeatureRequest{
		ApiService: a,
		ctx:        ctx,
		featureId:  featureId,
	}
}

// Execute executes the request
//
//	@return FeatureWrap
func (a *FeaturesAPIService) UpdateFeatureExecute(r ApiUpdateFeatureRequest) (*FeatureWrap, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FeatureWrap
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeaturesAPIService.UpdateFeature")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/features/{feature_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"feature_id"+"}", url.PathEscape(parameterValueToString(r.featureId, "featureId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.featureUpdate == nil {
		return localVarReturnValue, nil, reportError("featureUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.featureUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
