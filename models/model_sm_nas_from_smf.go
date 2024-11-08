/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Represents information on the SM congestion control applied SM NAS messages that SMF sends  to UE for PDU Session.
type SmNasFromSmf struct {
	SmNasType string `json:"smNasType" yaml:"smNasType" bson:"smNasType,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp" yaml:"timeStamp" bson:"timeStamp,omitempty"`
	// indicating a time in seconds.
	BackoffTimer    int32           `json:"backoffTimer" yaml:"backoffTimer" bson:"backoffTimer,omitempty"`
	AppliedSmccType AppliedSmccType `json:"appliedSmccType" yaml:"appliedSmccType" bson:"appliedSmccType,omitempty"`
}
