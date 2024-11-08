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

type QueryAMFSubscriptionInfoDocumentApiService service

/*
QueryAMFSubscriptionInfoDocumentApiService Retrieve AMF subscription Info for a group of UEs or any UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeGroupId -
 * @param SubsId -

@return GetAmfGroupSubscriptionsResponse
*/

// GetAmfGroupSubscriptionsRequest
type GetAmfGroupSubscriptionsRequest struct {
	UeGroupId *string
	SubsId    *string
}

func (r *GetAmfGroupSubscriptionsRequest) SetUeGroupId(UeGroupId string) {
	r.UeGroupId = &UeGroupId
}
func (r *GetAmfGroupSubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type GetAmfGroupSubscriptionsResponse struct {
	AmfSubscriptionInfo []models.AmfSubscriptionInfo
}

type GetAmfGroupSubscriptionsError struct {
}

func (a *QueryAMFSubscriptionInfoDocumentApiService) GetAmfGroupSubscriptions(ctx context.Context, request *GetAmfGroupSubscriptionsRequest) (*GetAmfGroupSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetAmfGroupSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/amf-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", openapi.StringOfValue(*request.UeGroupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

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
		err = openapi.Deserialize(&localVarReturnValue.AmfSubscriptionInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
QueryAMFSubscriptionInfoDocumentApiService Retrieve AMF subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -

@return GetAmfSubscriptionInfoResponse
*/

// GetAmfSubscriptionInfoRequest
type GetAmfSubscriptionInfoRequest struct {
	UeId   *string
	SubsId *string
}

func (r *GetAmfSubscriptionInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *GetAmfSubscriptionInfoRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type GetAmfSubscriptionInfoResponse struct {
	AmfSubscriptionInfo []models.AmfSubscriptionInfo
}

type GetAmfSubscriptionInfoError struct {
}

func (a *QueryAMFSubscriptionInfoDocumentApiService) GetAmfSubscriptionInfo(ctx context.Context, request *GetAmfSubscriptionInfoRequest) (*GetAmfSubscriptionInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetAmfSubscriptionInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

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
		err = openapi.Deserialize(&localVarReturnValue.AmfSubscriptionInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}
