/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.534 V17.3.0; 5G System; Access and Mobility Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.534/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AMPolicyAuthorization

// APIClient manages communication with the Npcf_AMPolicyAuthorization Service API API v1.0.2
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AMPolicyEventsSubscriptionDocumentApi     *AMPolicyEventsSubscriptionDocumentApiService
	ApplicationAMContextsCollectionApi        *ApplicationAMContextsCollectionApiService
	IndividualApplicationAMContextDocumentApi *IndividualApplicationAMContextDocumentApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AMPolicyEventsSubscriptionDocumentApi = (*AMPolicyEventsSubscriptionDocumentApiService)(&c.common)
	c.ApplicationAMContextsCollectionApi = (*ApplicationAMContextsCollectionApiService)(&c.common)
	c.IndividualApplicationAMContextDocumentApi = (*IndividualApplicationAMContextDocumentApiService)(&c.common)

	return c
}
