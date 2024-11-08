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

// Represents notifications on events that occurred.
type NsmfEventExposureNotification struct {
	// Notification correlation ID
	NotifId string `json:"notifId" yaml:"notifId" bson:"notifId,omitempty"`
	// Notifications about Individual Events
	EventNotifs []SmfEventExposureEventNotification `json:"eventNotifs" yaml:"eventNotifs" bson:"eventNotifs,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AckUri string `json:"ackUri,omitempty" yaml:"ackUri" bson:"ackUri,omitempty"`
}
