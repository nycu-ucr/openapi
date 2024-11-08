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

// Identifies the (event exposure or analytics) event that the processing instructions  shall apply to. Contains all event IDs related to DCCF.
type DccfEvent struct {
	NwdafEvent NwdafEvent             `json:"nwdafEvent,omitempty" yaml:"nwdafEvent" bson:"nwdafEvent,omitempty"`
	SmfEvent   SmfEvent               `json:"smfEvent,omitempty" yaml:"smfEvent" bson:"smfEvent,omitempty"`
	AmfEvent   AmfEventType           `json:"amfEvent,omitempty" yaml:"amfEvent" bson:"amfEvent,omitempty"`
	NefEvent   NefEvent               `json:"nefEvent,omitempty" yaml:"nefEvent" bson:"nefEvent,omitempty"`
	UdmEvent   UdmEeEventType         `json:"udmEvent,omitempty" yaml:"udmEvent" bson:"udmEvent,omitempty"`
	AfEvent    AfEventExposureAfEvent `json:"afEvent,omitempty" yaml:"afEvent" bson:"afEvent,omitempty"`
	SacEvent   *SacEvent              `json:"sacEvent,omitempty" yaml:"sacEvent" bson:"sacEvent,omitempty"`
	NrfEvent   NotificationEventType  `json:"nrfEvent,omitempty" yaml:"nrfEvent" bson:"nrfEvent,omitempty"`
}
