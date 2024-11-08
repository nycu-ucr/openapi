/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.522 V17.7.0; 5G System; Network Exposure Function Northbound APIs.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.522/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type.
type RouteToLocation struct {
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai      string            `json:"dnai" yaml:"dnai" bson:"dnai,omitempty"`
	RouteInfo *RouteInformation `json:"routeInfo,omitempty" yaml:"routeInfo" bson:"routeInfo,omitempty"`
	// Identifies the routing profile Id.
	RouteProfId string `json:"routeProfId,omitempty" yaml:"routeProfId" bson:"routeProfId,omitempty"`
}
