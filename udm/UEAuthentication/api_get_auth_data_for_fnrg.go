/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.10.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package UEAuthentication

import (
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type GetAuthDataForFNRGApiService service

/*
GetAuthDataForFNRGApiService Get authentication data for the FN-RG
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param SupiOrSuci - SUPI or SUCI of the user
 * @param AuthenticatedInd - Authenticated indication
 * @param SupportedFeatures - Supported Features
 * @param PlmnId - serving PLMN ID
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfModifiedSince - Validator for conditional requests, as described in RFC 7232, 3.3

@return GetRgAuthDataResponse
*/

// GetRgAuthDataRequest
type GetRgAuthDataRequest struct {
	SupiOrSuci        *string
	AuthenticatedInd  *bool
	SupportedFeatures *string
	PlmnId            *models.PlmnId
	IfNoneMatch       *string
	IfModifiedSince   *string
}

func (r *GetRgAuthDataRequest) SetSupiOrSuci(SupiOrSuci string) {
	r.SupiOrSuci = &SupiOrSuci
}
func (r *GetRgAuthDataRequest) SetAuthenticatedInd(AuthenticatedInd bool) {
	r.AuthenticatedInd = &AuthenticatedInd
}
func (r *GetRgAuthDataRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}
func (r *GetRgAuthDataRequest) SetPlmnId(PlmnId models.PlmnId) {
	r.PlmnId = &PlmnId
}
func (r *GetRgAuthDataRequest) SetIfNoneMatch(IfNoneMatch string) {
	r.IfNoneMatch = &IfNoneMatch
}
func (r *GetRgAuthDataRequest) SetIfModifiedSince(IfModifiedSince string) {
	r.IfModifiedSince = &IfModifiedSince
}

type GetRgAuthDataResponse struct {
	UdmUeauRgAuthCtx models.UdmUeauRgAuthCtx
}

type GetRgAuthDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *GetAuthDataForFNRGApiService) GetRgAuthData(ctx context.Context, request *GetRgAuthDataRequest) (*GetRgAuthDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetRgAuthDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{supiOrSuci}/security-information-rg"
	localVarPath = strings.Replace(localVarPath, "{"+"supiOrSuci"+"}", openapi.StringOfValue(*request.SupiOrSuci), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.AuthenticatedInd == nil {
		return nil, openapi.ReportError("AuthenticatedInd must be non nil")
	} else {
		localVarQueryParams.Add("authenticated-ind", openapi.ParameterToString(request.AuthenticatedInd, "multi"))
	}
	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	if request.PlmnId != nil {
		localVarQueryParams.Add("plmn-id", openapi.ParameterToString(request.PlmnId, "application/json"))
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
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(request.IfNoneMatch, "csv")
	}

	if request.IfModifiedSince != nil {
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(request.IfModifiedSince, "csv")
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
		err = openapi.Deserialize(&localVarReturnValue.UdmUeauRgAuthCtx, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 400:
		var v GetRgAuthDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v GetRgAuthDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetRgAuthDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetRgAuthDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetRgAuthDataError
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
