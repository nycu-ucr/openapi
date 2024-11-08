/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.507 V17.9.0; 5G System; Access and Mobility Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.507/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents updated policies that the PCF provides in a notification or in a reply to an Update Request.
type PcfAmPolicyControlPolicyUpdate struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri" yaml:"resourceUri" bson:"resourceUri,omitempty"`
	// Request Triggers that the PCF subscribes.
	Triggers      []PcfAmPolicyControlRequestTrigger `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	ServAreaRes   *ServiceAreaRestriction            `json:"servAreaRes,omitempty" yaml:"servAreaRes" bson:"servAreaRes,omitempty"`
	WlServAreaRes *WirelineServiceAreaRestriction    `json:"wlServAreaRes,omitempty" yaml:"wlServAreaRes" bson:"wlServAreaRes,omitempty"`
	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413.
	Rfsp int32 `json:"rfsp,omitempty" yaml:"rfsp" bson:"rfsp,omitempty"`
	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413.
	TargetRfsp int32             `json:"targetRfsp,omitempty" yaml:"targetRfsp" bson:"targetRfsp,omitempty"`
	SmfSelInfo *SmfSelectionData `json:"smfSelInfo,omitempty" yaml:"smfSelInfo" bson:"smfSelInfo,omitempty"`
	UeAmbr     *Ambr             `json:"ueAmbr,omitempty" yaml:"ueAmbr" bson:"ueAmbr,omitempty"`
	// One or more UE-Slice-MBR(s) for S-NSSAI(s) of serving PLMN the allowed NSSAI as part of the AMF Access and Mobility Policy as determined by the PCF.
	UeSliceMbrs []*UeSliceMbr `json:"ueSliceMbrs,omitempty" yaml:"ueSliceMbrs" bson:"ueSliceMbrs,omitempty"`
	// Contains the presence reporting area(s) for which reporting was requested. The praId attribute within the PresenceInfo data type is the key of the map.
	Pras           map[string]*PresenceInfoRm                 `json:"pras,omitempty" yaml:"pras" bson:"pras,omitempty"`
	PcfUeInfo      *PcfUeCallbackInfo                         `json:"pcfUeInfo,omitempty" yaml:"pcfUeInfo" bson:"pcfUeInfo,omitempty"`
	MatchPdus      []PduSessionInfo                           `json:"matchPdus,omitempty" yaml:"matchPdus" bson:"matchPdus,omitempty"`
	AsTimeDisParam *PcfAmPolicyControlAsTimeDistributionParam `json:"asTimeDisParam,omitempty" yaml:"asTimeDisParam" bson:"asTimeDisParam,omitempty"`
}
