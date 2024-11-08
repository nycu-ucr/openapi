/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmfNon3GppAccessRegistrationModification struct {
	Guami     *Guami `json:"guami" yaml:"guami" bson:"guami,omitempty"`
	PurgeFlag bool   `json:"purgeFlag,omitempty" yaml:"purgeFlag" bson:"purgeFlag,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei           string          `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty" yaml:"imsVoPs" bson:"imsVoPs,omitempty"`
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty" yaml:"backupAmfInfo" bson:"backupAmfInfo,omitempty"`
}
