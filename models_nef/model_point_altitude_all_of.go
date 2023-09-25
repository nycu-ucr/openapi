/*
 * 3gpp-pfd-management
 *
 * API for PFD management. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

import (
	"github.com/nycu-ucr/openapi/models"
)

type PointAltitudeAllOf struct {
	Point models.GeographicalCoordinates `json:"point"`

	Altitude float64 `json:"altitude"`
}
