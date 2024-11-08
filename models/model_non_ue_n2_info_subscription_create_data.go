/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Data within a create subscription request for non-UE specific N2 information notification
type NonUeN2InfoSubscriptionCreateData struct {
	GlobalRanNodeList  []GlobalRanNodeId  `json:"globalRanNodeList,omitempty" yaml:"globalRanNodeList" bson:"globalRanNodeList,omitempty"`
	AnTypeList         []AccessType       `json:"anTypeList,omitempty" yaml:"anTypeList" bson:"anTypeList,omitempty"`
	N2InformationClass N2InformationClass `json:"n2InformationClass" yaml:"n2InformationClass" bson:"n2InformationClass,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	N2NotifyCallbackUri string `json:"n2NotifyCallbackUri" yaml:"n2NotifyCallbackUri" bson:"n2NotifyCallbackUri,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfId string `json:"nfId,omitempty" yaml:"nfId" bson:"nfId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
}
