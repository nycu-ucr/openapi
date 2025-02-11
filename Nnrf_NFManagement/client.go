//+build !debug

/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// APIClient manages communication with the NRF NFManagement Service API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	NFInstanceIDDocumentApi    *NFInstanceIDDocumentApiService
	NFInstancesStoreApi        *NFInstancesStoreApiService
	SubscriptionIDDocumentApi  *SubscriptionIDDocumentApiService
	SubscriptionsCollectionApi *SubscriptionsCollectionApiService
	NotificationApi            *NotificationApiService
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
	c.NFInstanceIDDocumentApi = (*NFInstanceIDDocumentApiService)(&c.common)
	c.NFInstancesStoreApi = (*NFInstancesStoreApiService)(&c.common)
	c.SubscriptionIDDocumentApi = (*SubscriptionIDDocumentApiService)(&c.common)
	c.SubscriptionsCollectionApi = (*SubscriptionsCollectionApiService)(&c.common)
	c.NotificationApi = (*NotificationApiService)(&c.common)
	return c
}
