/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RedirectInformation struct {
	// Indicates the redirect is enable.
	RedirectEnabled bool `json:"redirectEnabled,omitempty" bson:"redirectEnabled"`

	RedirectAddressType RedirectAddressType `json:"redirectAddressType,omitempty" bson:"redirectAddressType"`

	// Indicates the address of the redirect server.
	RedirectServerAddress string `json:"redirectServerAddress,omitempty" bson:"redirectServerAddress"`
}
