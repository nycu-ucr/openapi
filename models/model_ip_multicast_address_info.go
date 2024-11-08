/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the IP multicast addressing information.
type IpMulticastAddressInfo struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	SrcIpv4Addr string `json:"srcIpv4Addr,omitempty" yaml:"srcIpv4Addr" bson:"srcIpv4Addr,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4MulAddr string `json:"ipv4MulAddr,omitempty" yaml:"ipv4MulAddr" bson:"ipv4MulAddr,omitempty"`
	SrcIpv6Addr string `json:"srcIpv6Addr,omitempty" yaml:"srcIpv6Addr" bson:"srcIpv6Addr,omitempty"`
	Ipv6MulAddr string `json:"ipv6MulAddr,omitempty" yaml:"ipv6MulAddr" bson:"ipv6MulAddr,omitempty"`
}
