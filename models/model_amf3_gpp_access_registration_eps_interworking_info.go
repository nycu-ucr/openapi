/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Amf3GppAccessRegistrationEpsInterworkingInfo struct {
	// A map (list of key-value pairs where Dnn serves as key) of EpsIwkPgws
	EpsIwkPgws map[string]EpsIwkPgw `json:"epsIwkPgws,omitempty" yaml:"epsIwkPgws" bson:"epsIwkPgws" mapstructure:"EpsIwkPgws"`
}
