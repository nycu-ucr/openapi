/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.10.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents DN performance data.
type PerfData struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	AvgTrafficRate string `json:"avgTrafficRate,omitempty" yaml:"avgTrafficRate" bson:"avgTrafficRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxTrafficRate string `json:"maxTrafficRate,omitempty" yaml:"maxTrafficRate" bson:"maxTrafficRate,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AvePacketDelay int32 `json:"avePacketDelay,omitempty" yaml:"avePacketDelay" bson:"avePacketDelay,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	MaxPacketDelay int32 `json:"maxPacketDelay,omitempty" yaml:"maxPacketDelay" bson:"maxPacketDelay,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.
	AvgPacketLossRate int32 `json:"avgPacketLossRate,omitempty" yaml:"avgPacketLossRate" bson:"avgPacketLossRate,omitempty"`
}
