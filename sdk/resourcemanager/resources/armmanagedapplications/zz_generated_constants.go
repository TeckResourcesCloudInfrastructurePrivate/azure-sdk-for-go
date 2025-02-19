//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

const (
	moduleName    = "armmanagedapplications"
	moduleVersion = "v1.0.0"
)

// ApplicationArtifactType - The managed application artifact type.
type ApplicationArtifactType string

const (
	ApplicationArtifactTypeTemplate ApplicationArtifactType = "Template"
	ApplicationArtifactTypeCustom   ApplicationArtifactType = "Custom"
)

// PossibleApplicationArtifactTypeValues returns the possible values for the ApplicationArtifactType const type.
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return []ApplicationArtifactType{
		ApplicationArtifactTypeTemplate,
		ApplicationArtifactTypeCustom,
	}
}

// ApplicationLockLevel - The managed application lock level.
type ApplicationLockLevel string

const (
	ApplicationLockLevelCanNotDelete ApplicationLockLevel = "CanNotDelete"
	ApplicationLockLevelReadOnly     ApplicationLockLevel = "ReadOnly"
	ApplicationLockLevelNone         ApplicationLockLevel = "None"
)

// PossibleApplicationLockLevelValues returns the possible values for the ApplicationLockLevel const type.
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return []ApplicationLockLevel{
		ApplicationLockLevelCanNotDelete,
		ApplicationLockLevelReadOnly,
		ApplicationLockLevelNone,
	}
}

// ProvisioningState - Provisioning status of the managed application.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreated   ProvisioningState = "Created"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateReady     ProvisioningState = "Ready"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateReady,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}
