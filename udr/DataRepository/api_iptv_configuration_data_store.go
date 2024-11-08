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

type IPTVConfigurationDataStoreApiService service

/*
IPTVConfigurationDataStoreApiService Retrieve IPTV configuration Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ConfigIds - Each element identifies a configuration.
 * @param Dnns - Each element identifies a DNN.
 * @param Snssais - Each element identifies a slice.
 * @param Supis - Each element identifies the user.
 * @param InterGroupIds - Each element identifies a group of users.

@return ReadIPTVCongifurationDataResponse
*/

// ReadIPTVCongifurationDataRequest
type ReadIPTVCongifurationDataRequest struct {
	ConfigIds     []string
	Dnns          []string
	Snssais       []models.Snssai
	Supis         []string
	InterGroupIds []string
}

func (r *ReadIPTVCongifurationDataRequest) SetConfigIds(ConfigIds []string) {
	r.ConfigIds = ConfigIds
}
func (r *ReadIPTVCongifurationDataRequest) SetDnns(Dnns []string) {
	r.Dnns = Dnns
}
func (r *ReadIPTVCongifurationDataRequest) SetSnssais(Snssais []models.Snssai) {
	r.Snssais = Snssais
}
func (r *ReadIPTVCongifurationDataRequest) SetSupis(Supis []string) {
	r.Supis = Supis
}
func (r *ReadIPTVCongifurationDataRequest) SetInterGroupIds(InterGroupIds []string) {
	r.InterGroupIds = InterGroupIds
}

type ReadIPTVCongifurationDataResponse struct {
	IptvConfigData []models.IptvConfigData
}

type ReadIPTVCongifurationDataError struct {
	ProblemDetails models.ProblemDetails
}

func (a *IPTVConfigurationDataStoreApiService) ReadIPTVCongifurationData(ctx context.Context, request *ReadIPTVCongifurationDataRequest) (*ReadIPTVCongifurationDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReadIPTVCongifurationDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/application-data/iptvConfigData"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.ConfigIds != nil {
		if len(request.ConfigIds) < 1 {
			return &localVarReturnValue, openapi.ReportError("ConfigIds must have at least 1 elements")
		}
		localVarQueryParams.Add("config-ids", openapi.ParameterToString(request.ConfigIds, "multi"))
	}
	if request.Dnns != nil {
		if len(request.Dnns) < 1 {
			return &localVarReturnValue, openapi.ReportError("Dnns must have at least 1 elements")
		}
		localVarQueryParams.Add("dnns", openapi.ParameterToString(request.Dnns, "multi"))
	}
	if request.Snssais != nil {
		if len(request.Snssais) < 1 {
			return &localVarReturnValue, openapi.ReportError("Snssais must have at least 1 elements")
		}
		localVarQueryParams.Add("snssais", openapi.ParameterToString(request.Snssais, "application/json"))
	}
	if request.Supis != nil {
		if len(request.Supis) < 1 {
			return &localVarReturnValue, openapi.ReportError("Supis must have at least 1 elements")
		}
		localVarQueryParams.Add("supis", openapi.ParameterToString(request.Supis, "multi"))
	}
	if request.InterGroupIds != nil {
		if len(request.InterGroupIds) < 1 {
			return &localVarReturnValue, openapi.ReportError("InterGroupIds must have at least 1 elements")
		}
		localVarQueryParams.Add("inter-group-ids", openapi.ParameterToString(request.InterGroupIds, "multi"))
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
		err = openapi.Deserialize(&localVarReturnValue.IptvConfigData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 400:
		var v ReadIPTVCongifurationDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v ReadIPTVCongifurationDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v ReadIPTVCongifurationDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ReadIPTVCongifurationDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 414:
		var v ReadIPTVCongifurationDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v ReadIPTVCongifurationDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v ReadIPTVCongifurationDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v ReadIPTVCongifurationDataError
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
