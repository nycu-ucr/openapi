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

type EventId string

// List of EventId
const (
	EventId_LOAD_LEVEL_INFORMATION EventId = "LOAD_LEVEL_INFORMATION"
	EventId_NETWORK_PERFORMANCE    EventId = "NETWORK_PERFORMANCE"
	EventId_NF_LOAD                EventId = "NF_LOAD"
	EventId_SERVICE_EXPERIENCE     EventId = "SERVICE_EXPERIENCE"
	EventId_UE_MOBILITY            EventId = "UE_MOBILITY"
	EventId_UE_COMMUNICATION       EventId = "UE_COMMUNICATION"
	EventId_QOS_SUSTAINABILITY     EventId = "QOS_SUSTAINABILITY"
	EventId_ABNORMAL_BEHAVIOUR     EventId = "ABNORMAL_BEHAVIOUR"
	EventId_USER_DATA_CONGESTION   EventId = "USER_DATA_CONGESTION"
	EventId_NSI_LOAD_LEVEL         EventId = "NSI_LOAD_LEVEL"
	EventId_SM_CONGESTION          EventId = "SM_CONGESTION"
	EventId_DISPERSION             EventId = "DISPERSION"
	EventId_RED_TRANS_EXP          EventId = "RED_TRANS_EXP"
	EventId_WLAN_PERFORMANCE       EventId = "WLAN_PERFORMANCE"
	EventId_DN_PERFORMANCE         EventId = "DN_PERFORMANCE"
)
