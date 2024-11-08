/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 17.6.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// IMS Private Identity and its identity type (IMPI or IMSI)
type PrivateIdentity struct {
	// String containing a Private User Identity or a Private Service Identity
	PrivateIdentity     string              `json:"privateIdentity" yaml:"privateIdentity" bson:"privateIdentity,omitempty"`
	PrivateIdentityType PrivateIdentityType `json:"privateIdentityType" yaml:"privateIdentityType" bson:"privateIdentityType,omitempty"`
}
