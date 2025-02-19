//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork

const (
	moduleName    = "armmobilenetwork"
	moduleVersion = "v0.6.0"
)

// BillingSKU - The SKU of the packet core control plane resource. The SKU list may change over time when a new SKU gets added
// or an exiting SKU gets removed.
type BillingSKU string

const (
	// BillingSKUEdgeSite2GBPS - Edge site 2Gbps plan
	BillingSKUEdgeSite2GBPS BillingSKU = "EdgeSite2GBPS"
	// BillingSKUEdgeSite3GBPS - Edge site 3Gbps plan
	BillingSKUEdgeSite3GBPS BillingSKU = "EdgeSite3GBPS"
	// BillingSKUEdgeSite4GBPS - Edge site 4Gbps plan
	BillingSKUEdgeSite4GBPS BillingSKU = "EdgeSite4GBPS"
	// BillingSKUEvaluationPackage - Evaluation package plan
	BillingSKUEvaluationPackage BillingSKU = "EvaluationPackage"
	// BillingSKUFlagshipStarterPackage - Flagship starter package plan
	BillingSKUFlagshipStarterPackage BillingSKU = "FlagshipStarterPackage"
	// BillingSKULargePackage - Large package plan
	BillingSKULargePackage BillingSKU = "LargePackage"
	// BillingSKUMediumPackage - Medium package plan
	BillingSKUMediumPackage BillingSKU = "MediumPackage"
)

// PossibleBillingSKUValues returns the possible values for the BillingSKU const type.
func PossibleBillingSKUValues() []BillingSKU {
	return []BillingSKU{
		BillingSKUEdgeSite2GBPS,
		BillingSKUEdgeSite3GBPS,
		BillingSKUEdgeSite4GBPS,
		BillingSKUEvaluationPackage,
		BillingSKUFlagshipStarterPackage,
		BillingSKULargePackage,
		BillingSKUMediumPackage,
	}
}

// CoreNetworkType - The core network technology generation (5G core or EPC / 4G core).
type CoreNetworkType string

const (
	// CoreNetworkTypeEPC - EPC / 4G core
	CoreNetworkTypeEPC CoreNetworkType = "EPC"
	// CoreNetworkTypeFiveGC - 5G core
	CoreNetworkTypeFiveGC CoreNetworkType = "5GC"
)

// PossibleCoreNetworkTypeValues returns the possible values for the CoreNetworkType const type.
func PossibleCoreNetworkTypeValues() []CoreNetworkType {
	return []CoreNetworkType{
		CoreNetworkTypeEPC,
		CoreNetworkTypeFiveGC,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// NaptEnabled - Whether network address and port translation is enabled.
type NaptEnabled string

const (
	// NaptEnabledDisabled - NAPT is disabled
	NaptEnabledDisabled NaptEnabled = "Disabled"
	// NaptEnabledEnabled - NAPT is enabled
	NaptEnabledEnabled NaptEnabled = "Enabled"
)

// PossibleNaptEnabledValues returns the possible values for the NaptEnabled const type.
func PossibleNaptEnabledValues() []NaptEnabled {
	return []NaptEnabled{
		NaptEnabledDisabled,
		NaptEnabledEnabled,
	}
}

// PduSessionType - PDU session type (IPv4/IPv6).
type PduSessionType string

const (
	PduSessionTypeIPv4 PduSessionType = "IPv4"
	PduSessionTypeIPv6 PduSessionType = "IPv6"
)

// PossiblePduSessionTypeValues returns the possible values for the PduSessionType const type.
func PossiblePduSessionTypeValues() []PduSessionType {
	return []PduSessionType{
		PduSessionTypeIPv4,
		PduSessionTypeIPv6,
	}
}

// PlatformType - The platform type where packet core is deployed. The contents of this enum can change.
type PlatformType string

const (
	// PlatformTypeAKSHCI - If this option is chosen, you must set one of "azureStackEdgeDevice", "connectedCluster" or "customLocation".
	// If multiple are set then "customLocation" will take precedence over "connectedCluster" which takes precedence over "azureStackEdgeDevice".
	PlatformTypeAKSHCI PlatformType = "AKS-HCI"
	// PlatformTypeBaseVM - If this option is chosen, you must set one of "connectedCluster" or "customLocation". If multiple
	// are set then "customLocation" will take precedence over "connectedCluster".
	PlatformTypeBaseVM PlatformType = "BaseVM"
)

// PossiblePlatformTypeValues returns the possible values for the PlatformType const type.
func PossiblePlatformTypeValues() []PlatformType {
	return []PlatformType{
		PlatformTypeAKSHCI,
		PlatformTypeBaseVM,
	}
}

// PreemptionCapability - Preemption capability.
type PreemptionCapability string

const (
	// PreemptionCapabilityMayPreempt - May preempt
	PreemptionCapabilityMayPreempt PreemptionCapability = "MayPreempt"
	// PreemptionCapabilityNotPreempt - Cannot preempt
	PreemptionCapabilityNotPreempt PreemptionCapability = "NotPreempt"
)

// PossiblePreemptionCapabilityValues returns the possible values for the PreemptionCapability const type.
func PossiblePreemptionCapabilityValues() []PreemptionCapability {
	return []PreemptionCapability{
		PreemptionCapabilityMayPreempt,
		PreemptionCapabilityNotPreempt,
	}
}

// PreemptionVulnerability - Preemption vulnerability.
type PreemptionVulnerability string

const (
	// PreemptionVulnerabilityNotPreemptable - Cannot be preempted
	PreemptionVulnerabilityNotPreemptable PreemptionVulnerability = "NotPreemptable"
	// PreemptionVulnerabilityPreemptable - May be preempted
	PreemptionVulnerabilityPreemptable PreemptionVulnerability = "Preemptable"
)

// PossiblePreemptionVulnerabilityValues returns the possible values for the PreemptionVulnerability const type.
func PossiblePreemptionVulnerabilityValues() []PreemptionVulnerability {
	return []PreemptionVulnerability{
		PreemptionVulnerabilityNotPreemptable,
		PreemptionVulnerabilityPreemptable,
	}
}

// ProvisioningState - The current provisioning state.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
	}
}

// RecommendedVersion - Indicates whether this is the recommended version to use for new packet core control plane deployments.
type RecommendedVersion string

const (
	// RecommendedVersionNotRecommended - This is not the recommended version to use for new packet core control plane deployments.
	RecommendedVersionNotRecommended RecommendedVersion = "NotRecommended"
	// RecommendedVersionRecommended - This is the recommended version to use for new packet core control plane deployments.
	RecommendedVersionRecommended RecommendedVersion = "Recommended"
)

// PossibleRecommendedVersionValues returns the possible values for the RecommendedVersion const type.
func PossibleRecommendedVersionValues() []RecommendedVersion {
	return []RecommendedVersion{
		RecommendedVersionNotRecommended,
		RecommendedVersionRecommended,
	}
}

// SdfDirection - Service data flow direction.
type SdfDirection string

const (
	// SdfDirectionBidirectional - Traffic flowing both to and from the UE.
	SdfDirectionBidirectional SdfDirection = "Bidirectional"
	// SdfDirectionDownlink - Traffic flowing from the data network to the UE.
	SdfDirectionDownlink SdfDirection = "Downlink"
	// SdfDirectionUplink - Traffic flowing from the UE to the data network.
	SdfDirectionUplink SdfDirection = "Uplink"
)

// PossibleSdfDirectionValues returns the possible values for the SdfDirection const type.
func PossibleSdfDirectionValues() []SdfDirection {
	return []SdfDirection{
		SdfDirectionBidirectional,
		SdfDirectionDownlink,
		SdfDirectionUplink,
	}
}

// SimState - The state of the SIM resource.
type SimState string

const (
	// SimStateDisabled - The SIM is disabled because not all configuration required for enabling is present.
	SimStateDisabled SimState = "Disabled"
	// SimStateEnabled - The SIM is enabled.
	SimStateEnabled SimState = "Enabled"
	// SimStateInvalid - The SIM cannot be enabled because some of the associated configuration is invalid.
	SimStateInvalid SimState = "Invalid"
)

// PossibleSimStateValues returns the possible values for the SimState const type.
func PossibleSimStateValues() []SimState {
	return []SimState{
		SimStateDisabled,
		SimStateEnabled,
		SimStateInvalid,
	}
}

// TrafficControlPermission - Traffic control permission.
type TrafficControlPermission string

const (
	// TrafficControlPermissionBlocked - Traffic matching this rule is not allowed to flow.
	TrafficControlPermissionBlocked TrafficControlPermission = "Blocked"
	// TrafficControlPermissionEnabled - Traffic matching this rule is allowed to flow.
	TrafficControlPermissionEnabled TrafficControlPermission = "Enabled"
)

// PossibleTrafficControlPermissionValues returns the possible values for the TrafficControlPermission const type.
func PossibleTrafficControlPermissionValues() []TrafficControlPermission {
	return []TrafficControlPermission{
		TrafficControlPermissionBlocked,
		TrafficControlPermissionEnabled,
	}
}

// VersionState - The state of this packet core control plane version.
type VersionState string

const (
	// VersionStateActive - This version is active and suitable for production use.
	VersionStateActive VersionState = "Active"
	// VersionStateDeprecated - This version is deprecated and is no longer supported.
	VersionStateDeprecated VersionState = "Deprecated"
	// VersionStatePreview - This version is a preview and is not suitable for production use.
	VersionStatePreview VersionState = "Preview"
	// VersionStateUnknown - The state of this version is unknown.
	VersionStateUnknown VersionState = "Unknown"
	// VersionStateValidating - This version is currently being validated.
	VersionStateValidating VersionState = "Validating"
	// VersionStateValidationFailed - This version failed validation.
	VersionStateValidationFailed VersionState = "ValidationFailed"
)

// PossibleVersionStateValues returns the possible values for the VersionState const type.
func PossibleVersionStateValues() []VersionState {
	return []VersionState{
		VersionStateActive,
		VersionStateDeprecated,
		VersionStatePreview,
		VersionStateUnknown,
		VersionStateValidating,
		VersionStateValidationFailed,
	}
}
