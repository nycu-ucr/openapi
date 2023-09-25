/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RegistrationMessageType string

// List of RegistrationMessageType
const (
	RegistrationMessageType_INITIAL        RegistrationMessageType = "INITIAL"
	RegistrationMessageType_MOBILITY       RegistrationMessageType = "MOBILITY"
	RegistrationMessageType_PERIODIC       RegistrationMessageType = "PERIODIC"
	RegistrationMessageType_EMERGENCY      RegistrationMessageType = "EMERGENCY"
	RegistrationMessageType_DEREGISTRATION RegistrationMessageType = "DEREGISTRATION"
)
