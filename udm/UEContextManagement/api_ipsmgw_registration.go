/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package UEContextManagement

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

type IPSMGWRegistrationApiService service

/*
IPSMGWRegistrationApiService Register an IP-SM-GW
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - Identifier of the UE
 * @param IpSmGwRegistration -

@return IpSmGwRegistrationResponse
*/

// IpSmGwRegistrationRequest
type IpSmGwRegistrationRequest struct {
	UeId               *string
	IpSmGwRegistration *models.IpSmGwRegistration
}

func (r *IpSmGwRegistrationRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *IpSmGwRegistrationRequest) SetIpSmGwRegistration(IpSmGwRegistration models.IpSmGwRegistration) {
	r.IpSmGwRegistration = &IpSmGwRegistration
}

type IpSmGwRegistrationResponse struct {
	Location           string
	IpSmGwRegistration models.IpSmGwRegistration
}

type IpSmGwRegistrationError struct {
	ProblemDetails models.ProblemDetails
}

func (a *IPSMGWRegistrationApiService) IpSmGwRegistration(ctx context.Context, request *IpSmGwRegistrationRequest) (*IpSmGwRegistrationResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IpSmGwRegistrationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{ueId}/registrations/ip-sm-gw"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.IpSmGwRegistration

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
		err = openapi.Deserialize(&localVarReturnValue.IpSmGwRegistration, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		return &localVarReturnValue, nil
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.IpSmGwRegistration, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 400:
		var v IpSmGwRegistrationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v IpSmGwRegistrationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v IpSmGwRegistrationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v IpSmGwRegistrationError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v IpSmGwRegistrationError
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
