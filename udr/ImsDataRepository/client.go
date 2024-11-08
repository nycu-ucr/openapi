/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V18.3.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ImsDataRepository

import (
	"crypto/tls"

	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/net/http2"
)

// APIClient manages communication with the Nudr_DataRepository API OpenAPI file API v2.3.0-alpha.4
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	ChargingInfoRetrievalApi                     *ChargingInfoRetrievalApiService
	IMSProfileDataRetrievalApi                   *IMSProfileDataRetrievalApiService
	RetrievalOfAssociatedIMSPrivateIdentitiesApi *RetrievalOfAssociatedIMSPrivateIdentitiesApiService
	RetrievalOfAssociatedIMSPublicIdentitiesApi  *RetrievalOfAssociatedIMSPublicIdentitiesApiService
	SCSCFRegistrationAndDeregistrationApi        *SCSCFRegistrationAndDeregistrationApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.httpClient == nil {
		cfg.httpClient = http.DefaultClient
		cfg.httpClient.Transport = &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // nolint:gosec
			},
		}
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.ChargingInfoRetrievalApi = (*ChargingInfoRetrievalApiService)(&c.common)
	c.IMSProfileDataRetrievalApi = (*IMSProfileDataRetrievalApiService)(&c.common)
	c.RetrievalOfAssociatedIMSPrivateIdentitiesApi = (*RetrievalOfAssociatedIMSPrivateIdentitiesApiService)(&c.common)
	c.RetrievalOfAssociatedIMSPublicIdentitiesApi = (*RetrievalOfAssociatedIMSPublicIdentitiesApiService)(&c.common)
	c.SCSCFRegistrationAndDeregistrationApi = (*SCSCFRegistrationAndDeregistrationApiService)(&c.common)

	return c
}
