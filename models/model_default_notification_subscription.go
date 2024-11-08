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

// Data structure for specifying the notifications the NF service subscribes by default, along with callback URI
type DefaultNotificationSubscription struct {
	NotificationType NrfNfManagementNotificationType `json:"notificationType" yaml:"notificationType" bson:"notificationType,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackUri string `json:"callbackUri" yaml:"callbackUri" bson:"callbackUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	InterPlmnCallbackUri string             `json:"interPlmnCallbackUri,omitempty" yaml:"interPlmnCallbackUri" bson:"interPlmnCallbackUri,omitempty"`
	N1MessageClass       N1MessageClass     `json:"n1MessageClass,omitempty" yaml:"n1MessageClass" bson:"n1MessageClass,omitempty"`
	N2InformationClass   N2InformationClass `json:"n2InformationClass,omitempty" yaml:"n2InformationClass" bson:"n2InformationClass,omitempty"`
	Versions             []string           `json:"versions,omitempty" yaml:"versions" bson:"versions,omitempty"`
	Binding              string             `json:"binding,omitempty" yaml:"binding" bson:"binding,omitempty"`
	AcceptedEncoding     string             `json:"acceptedEncoding,omitempty" yaml:"acceptedEncoding" bson:"acceptedEncoding,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	// A map of service specific information. The name of the corresponding service (as specified in ServiceName data type) is the key.
	ServiceInfoList map[string]DefSubServiceInfo `json:"serviceInfoList,omitempty" yaml:"serviceInfoList" bson:"serviceInfoList,omitempty"`
}
