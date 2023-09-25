/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type Trigger struct {
	TriggerType      TriggerType     `json:"triggerType" yaml:"triggerType" bson:"triggerType" mapstructure:"TriggerType"`
	TriggerCategory  TriggerCategory `json:"triggerCategory" yaml:"triggerCategory" bson:"triggerCategory" mapstructure:"TriggerCategory"`
	TimeLimit        int32           `json:"timeLimit,omitempty" yaml:"timeLimit" bson:"timeLimit" mapstructure:"TimeLimit"`
	VolumeLimit      int32           `json:"volumeLimit,omitempty" yaml:"volumeLimit" bson:"volumeLimit" mapstructure:"VolumeLimit"`
	VolumeLimit64    int32           `json:"volumeLimit64,omitempty" yaml:"volumeLimit64" bson:"volumeLimit64" mapstructure:"VolumeLimit64"`
	EventLimit       int32           `json:"eventLimit,omitempty" yaml:"eventLimit" bson:"eventLimit" mapstructure:"EventLimit"`
	MaxNumberOfccc   int32           `json:"maxNumberOfccc,omitempty" yaml:"maxNumberOfccc" bson:"maxNumberOfccc" mapstructure:"MaxNumberOfccc"`
	TariffTimeChange *time.Time      `json:"tariffTimeChange,omitempty" yaml:"tariffTimeChange" bson:"tariffTimeChange" mapstructure:"TariffTimeChange"`
}
