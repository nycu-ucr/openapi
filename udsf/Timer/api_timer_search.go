/*
 * Nudsf_Timer
 *
 * Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.598 UDSF Services, V17.6.0.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.598/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Timer

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

type TimerSearchApiService service

/*
TimerSearchApiService Timers search with get
Retrieve one or multiple TimerIDs based on filter
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param RealmId - Identifier of the Realm
 * @param StorageId - Identifier of the Storage
 * @param Filter - Query filter using conditions on tags
 * @param ExpiredFilter - Used to query for expired timers.
 * @param SupportedFeatures - Features required to be supported by the target NF

@return SearchTimerResponse
*/

// SearchTimerRequest
type SearchTimerRequest struct {
	RealmId           *string
	StorageId         *string
	Filter            *models.SearchExpression
	ExpiredFilter     *string
	SupportedFeatures *string
}

func (r *SearchTimerRequest) SetRealmId(RealmId string) {
	r.RealmId = &RealmId
}
func (r *SearchTimerRequest) SetStorageId(StorageId string) {
	r.StorageId = &StorageId
}
func (r *SearchTimerRequest) SetFilter(Filter models.SearchExpression) {
	r.Filter = &Filter
}
func (r *SearchTimerRequest) SetExpiredFilter(ExpiredFilter string) {
	r.ExpiredFilter = &ExpiredFilter
}
func (r *SearchTimerRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type SearchTimerResponse struct {
	TimerIdList models.TimerIdList
}

type SearchTimerError struct {
	ProblemDetails models.ProblemDetails
}

func (a *TimerSearchApiService) SearchTimer(ctx context.Context, request *SearchTimerRequest) (*SearchTimerResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchTimerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{realmId}/{storageId}/timers"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", openapi.StringOfValue(*request.RealmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", openapi.StringOfValue(*request.StorageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Filter != nil {
		localVarQueryParams.Add("filter", openapi.ParameterToString(request.Filter, "application/json"))
	}
	if request.ExpiredFilter != nil {
		localVarQueryParams.Add("expired-filter", openapi.ParameterToString(request.ExpiredFilter, "multi"))
	}
	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
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
		err = openapi.Deserialize(&localVarReturnValue.TimerIdList, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 400:
		var v SearchTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v SearchTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v SearchTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v SearchTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v SearchTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v SearchTimerError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v SearchTimerError
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
