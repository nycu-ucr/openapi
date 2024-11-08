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

type EventExposureGroupSubscriptionsCollectionApiService service

/*
EventExposureGroupSubscriptionsCollectionApiService Create individual EE subscription for a group of UEs or any UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeGroupId - Group of UEs or any UE
 * @param Body -

@return CreateEeGroupSubscriptionsResponse
*/

// CreateEeGroupSubscriptionsRequest
type CreateEeGroupSubscriptionsRequest struct {
	UeGroupId *string
	Body      *models.UdmEeEeSubscription
}

func (r *CreateEeGroupSubscriptionsRequest) SetUeGroupId(UeGroupId string) {
	r.UeGroupId = &UeGroupId
}
func (r *CreateEeGroupSubscriptionsRequest) SetBody(Body models.UdmEeEeSubscription) {
	r.Body = &Body
}

type CreateEeGroupSubscriptionsResponse struct {
	Location            string
	UdmEeEeSubscription models.UdmEeEeSubscription
}

type CreateEeGroupSubscriptionsError struct {
}

func (a *EventExposureGroupSubscriptionsCollectionApiService) CreateEeGroupSubscriptions(ctx context.Context, request *CreateEeGroupSubscriptionsRequest) (*CreateEeGroupSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateEeGroupSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", openapi.StringOfValue(*request.UeGroupId), -1)

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

	// body params
	localVarPostBody = request.Body

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
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.UdmEeEeSubscription, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
EventExposureGroupSubscriptionsCollectionApiService Retrieves the ee subscriptions of a group of UEs or any UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeGroupId - Group of UEs or any UE
 * @param SupportedFeatures - Supported Features

@return QueryEeGroupSubscriptionsResponse
*/

// QueryEeGroupSubscriptionsRequest
type QueryEeGroupSubscriptionsRequest struct {
	UeGroupId         *string
	SupportedFeatures *string
}

func (r *QueryEeGroupSubscriptionsRequest) SetUeGroupId(UeGroupId string) {
	r.UeGroupId = &UeGroupId
}
func (r *QueryEeGroupSubscriptionsRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type QueryEeGroupSubscriptionsResponse struct {
	UdmEeEeSubscription []models.UdmEeEeSubscription
}

type QueryEeGroupSubscriptionsError struct {
}

func (a *EventExposureGroupSubscriptionsCollectionApiService) QueryEeGroupSubscriptions(ctx context.Context, request *QueryEeGroupSubscriptionsRequest) (*QueryEeGroupSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QueryEeGroupSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", openapi.StringOfValue(*request.UeGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}

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
		err = openapi.Deserialize(&localVarReturnValue.UdmEeEeSubscription, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}
