/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.531 V17.8.0; 5G System; Network Slice Selection Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.531/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NSSAIAvailability

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

type NFInstanceIDDocumentApiService service

/*
NFInstanceIDDocumentApiService Deletes an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfId - Identifier of the NF service consumer instance

@return NSSAIAvailabilityDeleteResponse
*/

// NSSAIAvailabilityDeleteRequest
type NSSAIAvailabilityDeleteRequest struct {
	NfId *string
}

func (r *NSSAIAvailabilityDeleteRequest) SetNfId(NfId string) {
	r.NfId = &NfId
}

type NSSAIAvailabilityDeleteResponse struct {
}

type NSSAIAvailabilityDeleteError struct {
	Location             string
	Var3gppSbiTargetNfId string
	ProblemDetails       models.ProblemDetails
	RedirectResponse     models.RedirectResponse
}

func (a *NFInstanceIDDocumentApiService) NSSAIAvailabilityDelete(ctx context.Context, request *NSSAIAvailabilityDeleteRequest) (*NSSAIAvailabilityDeleteResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NSSAIAvailabilityDeleteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nssai-availability/{nfId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nfId"+"}", openapi.StringOfValue(*request.NfId), -1)

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
	case 307:
		var v NSSAIAvailabilityDeleteError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v NSSAIAvailabilityDeleteError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v NSSAIAvailabilityDeleteError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v NSSAIAvailabilityDeleteError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v NSSAIAvailabilityDeleteError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v NSSAIAvailabilityDeleteError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v NSSAIAvailabilityDeleteError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v NSSAIAvailabilityDeleteError
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
NFInstanceIDDocumentApiService Updates an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfId - Identifier of the NF service consumer instance
 * @param PatchItem - JSON Patch instructions to update at the NSSF, the S-NSSAIs supported per TA
 * @param ContentEncoding - Content-Encoding, described in IETF RFC 7231
 * @param AcceptEncoding - Accept-Encoding, described in IETF RFC 7231

@return NSSAIAvailabilityPatchResponse
*/

// NSSAIAvailabilityPatchRequest
type NSSAIAvailabilityPatchRequest struct {
	NfId            *string
	PatchItem       []models.PatchItem
	ContentEncoding *string
	AcceptEncoding  *string
}

func (r *NSSAIAvailabilityPatchRequest) SetNfId(NfId string) {
	r.NfId = &NfId
}
func (r *NSSAIAvailabilityPatchRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *NSSAIAvailabilityPatchRequest) SetContentEncoding(ContentEncoding string) {
	r.ContentEncoding = &ContentEncoding
}
func (r *NSSAIAvailabilityPatchRequest) SetAcceptEncoding(AcceptEncoding string) {
	r.AcceptEncoding = &AcceptEncoding
}

type NSSAIAvailabilityPatchResponse struct {
	AcceptEncoding                  string
	ContentEncoding                 string
	AuthorizedNssaiAvailabilityInfo models.AuthorizedNssaiAvailabilityInfo
}

type NSSAIAvailabilityPatchError struct {
	Location             string
	Var3gppSbiTargetNfId string
	ProblemDetails       models.ProblemDetails
	RedirectResponse     models.RedirectResponse
}

func (a *NFInstanceIDDocumentApiService) NSSAIAvailabilityPatch(ctx context.Context, request *NSSAIAvailabilityPatchRequest) (*NSSAIAvailabilityPatchResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NSSAIAvailabilityPatchResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nssai-availability/{nfId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nfId"+"}", openapi.StringOfValue(*request.NfId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json-patch+json:"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if request.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = openapi.ParameterToString(request.ContentEncoding, "csv")
	}

	if request.AcceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = openapi.ParameterToString(request.AcceptEncoding, "csv")
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
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.AuthorizedNssaiAvailabilityInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.AcceptEncoding = localVarHTTPResponse.Header.Get("Accept-Encoding")
		localVarReturnValue.ContentEncoding = localVarHTTPResponse.Header.Get("Content-Encoding")
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 307:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v NSSAIAvailabilityPatchError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v NSSAIAvailabilityPatchError
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
NFInstanceIDDocumentApiService Updates/replaces the NSSF with the S-NSSAIs the NF service consumer (e.g AMF)supports per TA
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfId - Identifier of the NF service consumer instance
 * @param NssaiAvailabilityInfo - Parameters to update/replace at the NSSF, the S-NSSAIs supported per TA
 * @param ContentEncoding - Content-Encoding, described in IETF RFC 7231
 * @param AcceptEncoding - Accept-Encoding, described in IETF RFC 7231

@return NSSAIAvailabilityPutResponse
*/

// NSSAIAvailabilityPutRequest
type NSSAIAvailabilityPutRequest struct {
	NfId                  *string
	NssaiAvailabilityInfo *models.NssaiAvailabilityInfo
	ContentEncoding       *string
	AcceptEncoding        *string
}

func (r *NSSAIAvailabilityPutRequest) SetNfId(NfId string) {
	r.NfId = &NfId
}
func (r *NSSAIAvailabilityPutRequest) SetNssaiAvailabilityInfo(NssaiAvailabilityInfo models.NssaiAvailabilityInfo) {
	r.NssaiAvailabilityInfo = &NssaiAvailabilityInfo
}
func (r *NSSAIAvailabilityPutRequest) SetContentEncoding(ContentEncoding string) {
	r.ContentEncoding = &ContentEncoding
}
func (r *NSSAIAvailabilityPutRequest) SetAcceptEncoding(AcceptEncoding string) {
	r.AcceptEncoding = &AcceptEncoding
}

type NSSAIAvailabilityPutResponse struct {
	AcceptEncoding                  string
	ContentEncoding                 string
	AuthorizedNssaiAvailabilityInfo models.AuthorizedNssaiAvailabilityInfo
}

type NSSAIAvailabilityPutError struct {
	Location             string
	Var3gppSbiTargetNfId string
	ProblemDetails       models.ProblemDetails
	RedirectResponse     models.RedirectResponse
}

func (a *NFInstanceIDDocumentApiService) NSSAIAvailabilityPut(ctx context.Context, request *NSSAIAvailabilityPutRequest) (*NSSAIAvailabilityPutResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NSSAIAvailabilityPutResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nssai-availability/{nfId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nfId"+"}", openapi.StringOfValue(*request.NfId), -1)

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

	if request.ContentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = openapi.ParameterToString(request.ContentEncoding, "csv")
	}

	if request.AcceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = openapi.ParameterToString(request.AcceptEncoding, "csv")
	}

	// body params
	localVarPostBody = request.NssaiAvailabilityInfo

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
		err = openapi.Deserialize(&localVarReturnValue.AuthorizedNssaiAvailabilityInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.AcceptEncoding = localVarHTTPResponse.Header.Get("Accept-Encoding")
		localVarReturnValue.ContentEncoding = localVarHTTPResponse.Header.Get("Content-Encoding")
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 307:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v NSSAIAvailabilityPutError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v NSSAIAvailabilityPutError
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
