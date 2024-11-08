/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/nycu-ucr/openapi"
	"github.com/nycu-ucr/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type Query5GVnGroupConfigurationDocumentApiService service

/*
Query5GVnGroupConfigurationDocumentApiService Retrieve a 5GVnGroup configuration
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ExternalGroupId -

@return Get5GVnGroupConfigurationResponse
*/

// Get5GVnGroupConfigurationRequest
type Get5GVnGroupConfigurationRequest struct {
	ExternalGroupId *string
}

func (r *Get5GVnGroupConfigurationRequest) SetExternalGroupId(ExternalGroupId string) {
	r.ExternalGroupId = &ExternalGroupId
}

type Get5GVnGroupConfigurationResponse struct {
	Model5GVnGroupConfiguration models.Model5GVnGroupConfiguration
}

type Get5GVnGroupConfigurationError struct {
}

func (a *Query5GVnGroupConfigurationDocumentApiService) Get5GVnGroupConfiguration(ctx context.Context, request *Get5GVnGroupConfigurationRequest) (*Get5GVnGroupConfigurationResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Get5GVnGroupConfigurationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/5g-vn-groups/{externalGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", openapi.StringOfValue(*request.ExternalGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.Model5GVnGroupConfiguration, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}
