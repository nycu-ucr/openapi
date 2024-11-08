/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ServiceTypeUnrelatedClass struct {
	// LCS service type.
	ServiceType               int32                     `json:"serviceType" yaml:"serviceType" bson:"serviceType,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty" yaml:"allowedGeographicArea" bson:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty" yaml:"privacyCheckRelatedAction" bson:"privacyCheckRelatedAction,omitempty"`
	CodeWordInd               CodeWordInd               `json:"codeWordInd,omitempty" yaml:"codeWordInd" bson:"codeWordInd,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty" yaml:"validTimePeriod" bson:"validTimePeriod,omitempty"`
	CodeWordList              []string                  `json:"codeWordList,omitempty" yaml:"codeWordList" bson:"codeWordList,omitempty"`
}
