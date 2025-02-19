/*
Data Catalog Service - Asset Details

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateAssetRequest struct {
	ctx                          context.Context
	ApiService                   *DefaultApiService
	xRequestDatacatalogWriteCred *string
	createAssetRequest           *CreateAssetRequest
}

// This header carries credential information related to accessing the relevant destination catalog.
func (r ApiCreateAssetRequest) XRequestDatacatalogWriteCred(xRequestDatacatalogWriteCred string) ApiCreateAssetRequest {
	r.xRequestDatacatalogWriteCred = &xRequestDatacatalogWriteCred
	return r
}

// Write Asset Request
func (r ApiCreateAssetRequest) CreateAssetRequest(createAssetRequest CreateAssetRequest) ApiCreateAssetRequest {
	r.createAssetRequest = &createAssetRequest
	return r
}

func (r ApiCreateAssetRequest) Execute() (*CreateAssetResponse, *http.Response, error) {
	return r.ApiService.CreateAssetExecute(r)
}

/*
CreateAsset This REST API writes data asset information to the data catalog configured in fybrik

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAssetRequest
*/
func (a *DefaultApiService) CreateAsset(ctx context.Context) ApiCreateAssetRequest {
	return ApiCreateAssetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return CreateAssetResponse
func (a *DefaultApiService) CreateAssetExecute(r ApiCreateAssetRequest) (*CreateAssetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateAssetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateAsset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/createAsset"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xRequestDatacatalogWriteCred == nil {
		return localVarReturnValue, nil, reportError("xRequestDatacatalogWriteCred is required and must be specified")
	}
	if r.createAssetRequest == nil {
		return localVarReturnValue, nil, reportError("createAssetRequest is required and must be specified")
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
	localVarHeaderParams["X-Request-Datacatalog-Write-Cred"] = parameterToString(*r.xRequestDatacatalogWriteCred, "")
	// body params
	localVarPostBody = r.createAssetRequest
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

type ApiGetAssetInfoRequest struct {
	ctx                     context.Context
	ApiService              *DefaultApiService
	xRequestDatacatalogCred *string
	getAssetRequest         *GetAssetRequest
}

// This header carries credential information related to relevant catalog from which the asset information needs to be retrieved.
func (r ApiGetAssetInfoRequest) XRequestDatacatalogCred(xRequestDatacatalogCred string) ApiGetAssetInfoRequest {
	r.xRequestDatacatalogCred = &xRequestDatacatalogCred
	return r
}

// Data Catalog Request Object.
func (r ApiGetAssetInfoRequest) GetAssetRequest(getAssetRequest GetAssetRequest) ApiGetAssetInfoRequest {
	r.getAssetRequest = &getAssetRequest
	return r
}

func (r ApiGetAssetInfoRequest) Execute() (*GetAssetResponse, *http.Response, error) {
	return r.ApiService.GetAssetInfoExecute(r)
}

/*
GetAssetInfo This REST API gets data asset information from the data catalog configured in fybrik for the data sets indicated in FybrikApplication yaml

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAssetInfoRequest
*/
func (a *DefaultApiService) GetAssetInfo(ctx context.Context) ApiGetAssetInfoRequest {
	return ApiGetAssetInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GetAssetResponse
func (a *DefaultApiService) GetAssetInfoExecute(r ApiGetAssetInfoRequest) (*GetAssetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAssetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetAssetInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/getAssetInfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xRequestDatacatalogCred == nil {
		return localVarReturnValue, nil, reportError("xRequestDatacatalogCred is required and must be specified")
	}
	if r.getAssetRequest == nil {
		return localVarReturnValue, nil, reportError("getAssetRequest is required and must be specified")
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
	localVarHeaderParams["X-Request-Datacatalog-Cred"] = parameterToString(*r.xRequestDatacatalogCred, "")
	// body params
	localVarPostBody = r.getAssetRequest
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
