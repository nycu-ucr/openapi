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

// Information of a PCF NF Instance
type PcfInfo struct {
	// Identifier of a group of NFs.
	GroupId    string          `json:"groupId,omitempty" yaml:"groupId" bson:"groupId,omitempty"`
	DnnList    []string        `json:"dnnList,omitempty" yaml:"dnnList" bson:"dnnList,omitempty"`
	SupiRanges []SupiRange     `json:"supiRanges,omitempty" yaml:"supiRanges" bson:"supiRanges,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty" yaml:"gpsiRanges" bson:"gpsiRanges,omitempty"`
	// Fully Qualified Domain Name
	RxDiamHost string `json:"rxDiamHost,omitempty" yaml:"rxDiamHost" bson:"rxDiamHost,omitempty"`
	// Fully Qualified Domain Name
	RxDiamRealm     string                          `json:"rxDiamRealm,omitempty" yaml:"rxDiamRealm" bson:"rxDiamRealm,omitempty"`
	V2xSupportInd   bool                            `json:"v2xSupportInd,omitempty" yaml:"v2xSupportInd" bson:"v2xSupportInd,omitempty"`
	ProseSupportInd bool                            `json:"proseSupportInd,omitempty" yaml:"proseSupportInd" bson:"proseSupportInd,omitempty"`
	ProseCapability *NrfNfManagementProSeCapability `json:"proseCapability,omitempty" yaml:"proseCapability" bson:"proseCapability,omitempty"`
	V2xCapability   *V2xCapability                  `json:"v2xCapability,omitempty" yaml:"v2xCapability" bson:"v2xCapability,omitempty"`
}
