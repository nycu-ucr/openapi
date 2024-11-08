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

type HSSSDMSubscriptionInfoDocumentApiService service

/*
HSSSDMSubscriptionInfoDocumentApiService Create HSS SDM Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -
 * @param HssSubscriptionInfo -

@return CreateHSSSDMSubscriptionsResponse
*/

// CreateHSSSDMSubscriptionsRequest
type CreateHSSSDMSubscriptionsRequest struct {
	UeId                *string
	SubsId              *string
	HssSubscriptionInfo *models.HssSubscriptionInfo
}

func (r *CreateHSSSDMSubscriptionsRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *CreateHSSSDMSubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}
func (r *CreateHSSSDMSubscriptionsRequest) SetHssSubscriptionInfo(HssSubscriptionInfo models.HssSubscriptionInfo) {
	r.HssSubscriptionInfo = &HssSubscriptionInfo
}

type CreateHSSSDMSubscriptionsResponse struct {
}

type CreateHSSSDMSubscriptionsError struct {
}

func (a *HSSSDMSubscriptionInfoDocumentApiService) CreateHSSSDMSubscriptions(ctx context.Context, request *CreateHSSSDMSubscriptionsRequest) (*CreateHSSSDMSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateHSSSDMSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.HssSubscriptionInfo

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
	case 204:
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
HSSSDMSubscriptionInfoDocumentApiService Retrieve HSS SDM Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -

@return GetHssSDMSubscriptionInfoResponse
*/

// GetHssSDMSubscriptionInfoRequest
type GetHssSDMSubscriptionInfoRequest struct {
	UeId   *string
	SubsId *string
}

func (r *GetHssSDMSubscriptionInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *GetHssSDMSubscriptionInfoRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type GetHssSDMSubscriptionInfoResponse struct {
	SmfSubscriptionInfo models.SmfSubscriptionInfo
}

type GetHssSDMSubscriptionInfoError struct {
}

func (a *HSSSDMSubscriptionInfoDocumentApiService) GetHssSDMSubscriptionInfo(ctx context.Context, request *GetHssSDMSubscriptionInfoRequest) (*GetHssSDMSubscriptionInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetHssSDMSubscriptionInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions"
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
		err = openapi.Deserialize(&localVarReturnValue.SmfSubscriptionInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
HSSSDMSubscriptionInfoDocumentApiService Modify HSS SDM Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -
 * @param PatchItem -
 * @param SupportedFeatures - Features required to be supported by the target NF

@return ModifyHssSDMSubscriptionInfoResponse
*/

// ModifyHssSDMSubscriptionInfoRequest
type ModifyHssSDMSubscriptionInfoRequest struct {
	UeId              *string
	SubsId            *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *ModifyHssSDMSubscriptionInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *ModifyHssSDMSubscriptionInfoRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}
func (r *ModifyHssSDMSubscriptionInfoRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *ModifyHssSDMSubscriptionInfoRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type ModifyHssSDMSubscriptionInfoResponse struct {
	PatchResult models.PatchResult
}

type ModifyHssSDMSubscriptionInfoError struct {
	ProblemDetails models.ProblemDetails
}

func (a *HSSSDMSubscriptionInfoDocumentApiService) ModifyHssSDMSubscriptionInfo(ctx context.Context, request *ModifyHssSDMSubscriptionInfoRequest) (*ModifyHssSDMSubscriptionInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModifyHssSDMSubscriptionInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.PatchItem

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
	case 204:
		return &localVarReturnValue, nil
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.PatchResult, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 403:
		var v ModifyHssSDMSubscriptionInfoError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}

/*
HSSSDMSubscriptionInfoDocumentApiService Delete HSS SDM Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -

@return RemoveHssSDMSubscriptionsInfoResponse
*/

// RemoveHssSDMSubscriptionsInfoRequest
type RemoveHssSDMSubscriptionsInfoRequest struct {
	UeId   *string
	SubsId *string
}

func (r *RemoveHssSDMSubscriptionsInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *RemoveHssSDMSubscriptionsInfoRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type RemoveHssSDMSubscriptionsInfoResponse struct {
}

type RemoveHssSDMSubscriptionsInfoError struct {
}

func (a *HSSSDMSubscriptionInfoDocumentApiService) RemoveHssSDMSubscriptionsInfo(ctx context.Context, request *RemoveHssSDMSubscriptionsInfoRequest) (*RemoveHssSDMSubscriptionsInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RemoveHssSDMSubscriptionsInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
	case 204:
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}
