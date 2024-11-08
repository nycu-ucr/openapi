/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.531 V17.7.0; 5G System; Network Slice Selection Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.531/
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Range of TACs (Tracking Area Codes)
type TacRange struct {
	Start   string `json:"start,omitempty" yaml:"start" bson:"start,omitempty"`
	End     string `json:"end,omitempty" yaml:"end" bson:"end,omitempty"`
	Pattern string `json:"pattern,omitempty" yaml:"pattern" bson:"pattern,omitempty"`
}
