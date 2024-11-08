/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V17.7.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AccessToken

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

type AccessTokenRequestApiService service

/*
AccessTokenRequestApiService Access Token Request
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param GrantType -
 * @param NfInstanceId - String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
 * @param Scope -
 * @param ContentEncoding - Content-Encoding, described in IETF RFC 7231
 * @param AcceptEncoding - Accept-Encoding, described in IETF RFC 7231
 * @param NfType -
 * @param TargetNfType -
 * @param TargetNfInstanceId - String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
 * @param RequesterPlmn -
 * @param RequesterPlmnList -
 * @param RequesterSnssaiList -
 * @param RequesterFqdn - Fully Qualified Domain Name
 * @param RequesterSnpnList -
 * @param TargetPlmn -
 * @param TargetSnpn -
 * @param TargetSnssaiList -
 * @param TargetNsiList -
 * @param TargetNfSetId - NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \\\"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\\\", or  \\\"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\\\" with  <MCC> encoded as defined in clause 5.4.2 (\\\"Mcc\\\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \\\"0\\\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
 * @param TargetNfServiceSetId - NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \\\"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\\\", or  \\\"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\\\" with  <MCC> encoded as defined in clause 5.4.2 (\\\"Mcc\\\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \\\"0\\\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\\\"Nid\\\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.
 * @param HnrfAccessTokenUri - String providing an URI formatted according to RFC 3986.
 * @param SourceNfInstanceId - String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.

@return AccessTokenRequestResponse
*/

// AccessTokenRequestRequest
type AccessTokenRequestRequest struct {
	GrantType            *string
	NfInstanceId         *string
	Scope                *string
	ContentEncoding      *string
	AcceptEncoding       *string
	NfType               *models.NrfNfManagementNfType
	TargetNfType         *models.NrfNfManagementNfType
	TargetNfInstanceId   *string
	RequesterPlmn        *models.PlmnId
	RequesterPlmnList    []models.PlmnId
	RequesterSnssaiList  []models.Snssai
	RequesterFqdn        *string
	RequesterSnpnList    []models.PlmnIdNid
	TargetPlmn           *models.PlmnId
	TargetSnpn           *models.PlmnIdNid
	TargetSnssaiList     []models.Snssai
	TargetNsiList        []string
	TargetNfSetId        *string
	TargetNfServiceSetId *string
	HnrfAccessTokenUri   *string
	SourceNfInstanceId   *string
}

func (r *AccessTokenRequestRequest) SetGrantType(GrantType string) {
	r.GrantType = &GrantType
}
func (r *AccessTokenRequestRequest) SetNfInstanceId(NfInstanceId string) {
	r.NfInstanceId = &NfInstanceId
}
func (r *AccessTokenRequestRequest) SetScope(Scope string) {
	r.Scope = &Scope
}
func (r *AccessTokenRequestRequest) SetContentEncoding(ContentEncoding string) {
	r.ContentEncoding = &ContentEncoding
}
func (r *AccessTokenRequestRequest) SetAcceptEncoding(AcceptEncoding string) {
	r.AcceptEncoding = &AcceptEncoding
}
func (r *AccessTokenRequestRequest) SetNfType(NfType models.NrfNfManagementNfType) {
	r.NfType = &NfType
}
func (r *AccessTokenRequestRequest) SetTargetNfType(TargetNfType models.NrfNfManagementNfType) {
	r.TargetNfType = &TargetNfType
}
func (r *AccessTokenRequestRequest) SetTargetNfInstanceId(TargetNfInstanceId string) {
	r.TargetNfInstanceId = &TargetNfInstanceId
}
func (r *AccessTokenRequestRequest) SetRequesterPlmn(RequesterPlmn models.PlmnId) {
	r.RequesterPlmn = &RequesterPlmn
}
func (r *AccessTokenRequestRequest) SetRequesterPlmnList(RequesterPlmnList []models.PlmnId) {
	r.RequesterPlmnList = RequesterPlmnList
}
func (r *AccessTokenRequestRequest) SetRequesterSnssaiList(RequesterSnssaiList []models.Snssai) {
	r.RequesterSnssaiList = RequesterSnssaiList
}
func (r *AccessTokenRequestRequest) SetRequesterFqdn(RequesterFqdn string) {
	r.RequesterFqdn = &RequesterFqdn
}
func (r *AccessTokenRequestRequest) SetRequesterSnpnList(RequesterSnpnList []models.PlmnIdNid) {
	r.RequesterSnpnList = RequesterSnpnList
}
func (r *AccessTokenRequestRequest) SetTargetPlmn(TargetPlmn models.PlmnId) {
	r.TargetPlmn = &TargetPlmn
}
func (r *AccessTokenRequestRequest) SetTargetSnpn(TargetSnpn models.PlmnIdNid) {
	r.TargetSnpn = &TargetSnpn
}
func (r *AccessTokenRequestRequest) SetTargetSnssaiList(TargetSnssaiList []models.Snssai) {
	r.TargetSnssaiList = TargetSnssaiList
}
func (r *AccessTokenRequestRequest) SetTargetNsiList(TargetNsiList []string) {
	r.TargetNsiList = TargetNsiList
}
func (r *AccessTokenRequestRequest) SetTargetNfSetId(TargetNfSetId string) {
	r.TargetNfSetId = &TargetNfSetId
}
func (r *AccessTokenRequestRequest) SetTargetNfServiceSetId(TargetNfServiceSetId string) {
	r.TargetNfServiceSetId = &TargetNfServiceSetId
}
func (r *AccessTokenRequestRequest) SetHnrfAccessTokenUri(HnrfAccessTokenUri string) {
	r.HnrfAccessTokenUri = &HnrfAccessTokenUri
}
func (r *AccessTokenRequestRequest) SetSourceNfInstanceId(SourceNfInstanceId string) {
	r.SourceNfInstanceId = &SourceNfInstanceId
}

type AccessTokenRequestResponse struct {
	AcceptEncoding               string
	CacheControl                 string
	ContentEncoding              string
	Pragma                       string
	NrfAccessTokenAccessTokenRsp models.NrfAccessTokenAccessTokenRsp
}

type AccessTokenRequestError struct {
	CacheControl     string
	Location         string
	Pragma           string
	AccessTokenErr   models.AccessTokenErr
	ProblemDetails   models.ProblemDetails
	RedirectResponse models.RedirectResponse
}

func (a *AccessTokenRequestApiService) AccessTokenRequest(ctx context.Context, request *AccessTokenRequestRequest) (*AccessTokenRequestResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccessTokenRequestResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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

	localVarFormParams.Add("grant_type", openapi.ParameterToString(request.GrantType, ""))

	localVarFormParams.Add("nfInstanceId", openapi.ParameterToString(request.NfInstanceId, ""))

	if request.NfType != nil {
		localVarFormParams.Add("nfType", openapi.ParameterToString(request.NfType, ""))
	}

	if request.TargetNfType != nil {
		localVarFormParams.Add("targetNfType", openapi.ParameterToString(request.TargetNfType, ""))
	}

	localVarFormParams.Add("scope", openapi.ParameterToString(request.Scope, ""))

	if request.TargetNfInstanceId != nil {
		localVarFormParams.Add("targetNfInstanceId", openapi.ParameterToString(request.TargetNfInstanceId, ""))
	}

	if request.RequesterPlmn != nil {
		localVarFormParams.Add("requesterPlmn", openapi.ParameterToString(request.RequesterPlmn, ""))
	}

	if request.RequesterPlmnList != nil {
		localVarFormParams.Add("requesterPlmnList", openapi.ParameterToString(request.RequesterPlmnList, "csv"))
	}

	if request.RequesterSnssaiList != nil {
		localVarFormParams.Add("requesterSnssaiList", openapi.ParameterToString(request.RequesterSnssaiList, "csv"))
	}

	if request.RequesterFqdn != nil {
		localVarFormParams.Add("requesterFqdn", openapi.ParameterToString(request.RequesterFqdn, ""))
	}

	if request.RequesterSnpnList != nil {
		localVarFormParams.Add("requesterSnpnList", openapi.ParameterToString(request.RequesterSnpnList, "csv"))
	}

	if request.TargetPlmn != nil {
		localVarFormParams.Add("targetPlmn", openapi.ParameterToString(request.TargetPlmn, ""))
	}

	if request.TargetSnpn != nil {
		localVarFormParams.Add("targetSnpn", openapi.ParameterToString(request.TargetSnpn, ""))
	}

	if request.TargetSnssaiList != nil {
		localVarFormParams.Add("targetSnssaiList", openapi.ParameterToString(request.TargetSnssaiList, "csv"))
	}

	if request.TargetNsiList != nil {
		localVarFormParams.Add("targetNsiList", openapi.ParameterToString(request.TargetNsiList, "csv"))
	}

	if request.TargetNfSetId != nil {
		localVarFormParams.Add("targetNfSetId", openapi.ParameterToString(request.TargetNfSetId, ""))
	}

	if request.TargetNfServiceSetId != nil {
		localVarFormParams.Add("targetNfServiceSetId", openapi.ParameterToString(request.TargetNfServiceSetId, ""))
	}

	if request.HnrfAccessTokenUri != nil {
		localVarFormParams.Add("hnrfAccessTokenUri", openapi.ParameterToString(request.HnrfAccessTokenUri, ""))
	}

	if request.SourceNfInstanceId != nil {
		localVarFormParams.Add("sourceNfInstanceId", openapi.ParameterToString(request.SourceNfInstanceId, ""))
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
		err = openapi.Deserialize(&localVarReturnValue.NrfAccessTokenAccessTokenRsp, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		localVarReturnValue.Pragma = localVarHTTPResponse.Header.Get("Pragma")
		localVarReturnValue.AcceptEncoding = localVarHTTPResponse.Header.Get("Accept-Encoding")
		localVarReturnValue.ContentEncoding = localVarHTTPResponse.Header.Get("Content-Encoding")
		return &localVarReturnValue, nil
	case 307:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.AccessTokenErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		v.Pragma = localVarHTTPResponse.Header.Get("Pragma")
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 501:
		var v AccessTokenRequestError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v AccessTokenRequestError
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
