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

type QueryServiceSpecificAuthorizationDataDocumentApiService service

/*
QueryServiceSpecificAuthorizationDataDocumentApiService Retrieve ServiceSpecific Authorization Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE ID
 * @param ServiceType - Service Type
 * @param SingleNssai - single NSSAI
 * @param Dnn - DNN
 * @param MtcProviderInformation - MTC Provider Information
 * @param AfId - Application Function Identifier
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfModifiedSince - Validator for conditional requests, as described in RFC 7232, 3.3

@return GetSSAuDataResponse
*/

// GetSSAuDataRequest
type GetSSAuDataRequest struct {
	UeId                   *string
	ServiceType            *models.ServiceType
	SingleNssai            *models.Snssai
	Dnn                    *string
	MtcProviderInformation *string
	AfId                   *string
	IfNoneMatch            *string
	IfModifiedSince        *string
}

func (r *GetSSAuDataRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *GetSSAuDataRequest) SetServiceType(ServiceType models.ServiceType) {
	r.ServiceType = &ServiceType
}
func (r *GetSSAuDataRequest) SetSingleNssai(SingleNssai models.Snssai) {
	r.SingleNssai = &SingleNssai
}
func (r *GetSSAuDataRequest) SetDnn(Dnn string) {
	r.Dnn = &Dnn
}
func (r *GetSSAuDataRequest) SetMtcProviderInformation(MtcProviderInformation string) {
	r.MtcProviderInformation = &MtcProviderInformation
}
func (r *GetSSAuDataRequest) SetAfId(AfId string) {
	r.AfId = &AfId
}
func (r *GetSSAuDataRequest) SetIfNoneMatch(IfNoneMatch string) {
	r.IfNoneMatch = &IfNoneMatch
}
func (r *GetSSAuDataRequest) SetIfModifiedSince(IfModifiedSince string) {
	r.IfModifiedSince = &IfModifiedSince
}

type GetSSAuDataResponse struct {
	CacheControl      string
	ETag              string
	LastModified      string
	AuthorizationData models.AuthorizationData
}

type GetSSAuDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *QueryServiceSpecificAuthorizationDataDocumentApiService) GetSSAuData(ctx context.Context, request *GetSSAuDataRequest) (*GetSSAuDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetSSAuDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/service-specific-authorization-data/{serviceType}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceType"+"}", openapi.StringOfValue(*request.ServiceType), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SingleNssai == nil {
		return nil, openapi.ReportError("SingleNssai must be non nil")
	} else {
		localVarQueryParams.Add("single-nssai", openapi.ParameterToString(request.SingleNssai, "application/json"))
	}
	if request.Dnn == nil {
		return nil, openapi.ReportError("Dnn must be non nil")
	} else {
		localVarQueryParams.Add("dnn", openapi.ParameterToString(request.Dnn, "multi"))
	}
	if request.MtcProviderInformation != nil {
		localVarQueryParams.Add("mtc-provider-information", openapi.ParameterToString(request.MtcProviderInformation, "multi"))
	}
	if request.AfId != nil {
		localVarQueryParams.Add("af-id", openapi.ParameterToString(request.AfId, "multi"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if request.IfNoneMatch != nil {
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(request.IfNoneMatch, "multi")
	}

	if request.IfModifiedSince != nil {
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(request.IfModifiedSince, "multi")
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
		err = openapi.Deserialize(&localVarReturnValue.AuthorizationData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
		return &localVarReturnValue, nil
	case 403:
		var v GetSSAuDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetSSAuDataError
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
