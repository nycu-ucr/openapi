/*
 * Nudm_SSAU
 *
 * Nudm Service Specific Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.9.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ServiceSpecificAuthorization

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

type ServiceSpecificAuthorizationRemoveApiService service

/*
ServiceSpecificAuthorizationRemoveApiService Remove the authorization of specific service's configuration.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeIdentity - Represents the scope of the UE for which the Service Specific configuration authorization to be removed. Contains the GPSI of the user or the external group ID.
 * @param ServiceType - Represents the specific service for which the Service Specific configuration authorization to be removed.
 * @param ServiceSpecificAuthorizationRemoveData -

@return ServiceSpecificAuthorizationRemovalResponse
*/

// ServiceSpecificAuthorizationRemovalRequest
type ServiceSpecificAuthorizationRemovalRequest struct {
	UeIdentity                             *string
	ServiceType                            *models.ServiceType
	ServiceSpecificAuthorizationRemoveData *models.ServiceSpecificAuthorizationRemoveData
}

func (r *ServiceSpecificAuthorizationRemovalRequest) SetUeIdentity(UeIdentity string) {
	r.UeIdentity = &UeIdentity
}
func (r *ServiceSpecificAuthorizationRemovalRequest) SetServiceType(ServiceType models.ServiceType) {
	r.ServiceType = &ServiceType
}
func (r *ServiceSpecificAuthorizationRemovalRequest) SetServiceSpecificAuthorizationRemoveData(ServiceSpecificAuthorizationRemoveData models.ServiceSpecificAuthorizationRemoveData) {
	r.ServiceSpecificAuthorizationRemoveData = &ServiceSpecificAuthorizationRemoveData
}

type ServiceSpecificAuthorizationRemovalResponse struct {
}

type ServiceSpecificAuthorizationRemovalError struct {
	ProblemDetails models.ProblemDetails
}

func (a *ServiceSpecificAuthorizationRemoveApiService) ServiceSpecificAuthorizationRemoval(ctx context.Context, request *ServiceSpecificAuthorizationRemovalRequest) (*ServiceSpecificAuthorizationRemovalResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ServiceSpecificAuthorizationRemovalResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{ueIdentity}/{serviceType}/remove"
	localVarPath = strings.Replace(localVarPath, "{"+"ueIdentity"+"}", openapi.StringOfValue(*request.UeIdentity), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceType"+"}", openapi.StringOfValue(*request.ServiceType), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.ServiceSpecificAuthorizationRemoveData

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
	case 400:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 502:
		var v ServiceSpecificAuthorizationRemovalError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v ServiceSpecificAuthorizationRemovalError
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
