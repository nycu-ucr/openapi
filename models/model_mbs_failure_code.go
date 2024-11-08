/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.537 V17.3.0; 5G System; Multicast/Broadcast Policy Control Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.537/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type MbsFailureCode string

// List of MbsFailureCode
const (
	MbsFailureCode_NF_MALFUNCTION              MbsFailureCode = "NF_MALFUNCTION"
	MbsFailureCode_NF_RESOURCES_UNAVAILABLE    MbsFailureCode = "NF_RESOURCES_UNAVAILABLE"
	MbsFailureCode_RESOURCE_ALLOCATION_FAILURE MbsFailureCode = "RESOURCE_ALLOCATION_FAILURE"
	MbsFailureCode_MBS_QOS_VALIDATION_FAILURE  MbsFailureCode = "MBS_QOS_VALIDATION_FAILURE"
	MbsFailureCode_NO_MBS_QOS_FLOW             MbsFailureCode = "NO_MBS_QOS_FLOW"
	MbsFailureCode_MBS_QOS_DECISION_ERROR      MbsFailureCode = "MBS_QOS_DECISION_ERROR"
	MbsFailureCode_MBS_POLICY_PARAM_ERROR      MbsFailureCode = "MBS_POLICY_PARAM_ERROR"
)
