/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ChfConvergedChargingPduAddress struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	PduIPv4Address           string   `json:"pduIPv4Address,omitempty" yaml:"pduIPv4Address" bson:"pduIPv4Address,omitempty"`
	PduIPv6AddresswithPrefix string   `json:"pduIPv6AddresswithPrefix,omitempty" yaml:"pduIPv6AddresswithPrefix" bson:"pduIPv6AddresswithPrefix,omitempty"`
	PduAddressprefixlength   int32    `json:"pduAddressprefixlength,omitempty" yaml:"pduAddressprefixlength" bson:"pduAddressprefixlength,omitempty"`
	IPv4dynamicAddressFlag   bool     `json:"iPv4dynamicAddressFlag,omitempty" yaml:"iPv4dynamicAddressFlag" bson:"iPv4dynamicAddressFlag,omitempty"`
	IPv6dynamicPrefixFlag    bool     `json:"iPv6dynamicPrefixFlag,omitempty" yaml:"iPv6dynamicPrefixFlag" bson:"iPv6dynamicPrefixFlag,omitempty"`
	AddIpv6AddrPrefixes      string   `json:"addIpv6AddrPrefixes,omitempty" yaml:"addIpv6AddrPrefixes" bson:"addIpv6AddrPrefixes,omitempty"`
	AddIpv6AddrPrefixList    []string `json:"addIpv6AddrPrefixList,omitempty" yaml:"addIpv6AddrPrefixList" bson:"addIpv6AddrPrefixList,omitempty"`
}
