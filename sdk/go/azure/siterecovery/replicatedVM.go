// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a VM replicated using Azure Site Recovery (Azure to Azure only). A replicated VM keeps a copiously updated image of the VM in another region in order to be able to start the VM in that region in case of a disaster.
type ReplicatedVM struct {
	pulumi.CustomResourceState

	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVMManagedDiskArrayOutput `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name                        pulumi.StringOutput `pulumi:"name"`
	RecoveryReplicationPolicyId pulumi.StringOutput `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringOutput `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringOutput `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId pulumi.StringOutput `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringPtrOutput `pulumi:"targetAvailabilitySetId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringOutput `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringOutput `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringOutput `pulumi:"targetResourceGroupId"`
}

// NewReplicatedVM registers a new resource with the given unique name, arguments, and options.
func NewReplicatedVM(ctx *pulumi.Context,
	name string, args *ReplicatedVMArgs, opts ...pulumi.ResourceOption) (*ReplicatedVM, error) {
	if args == nil || args.RecoveryReplicationPolicyId == nil {
		return nil, errors.New("missing required argument 'RecoveryReplicationPolicyId'")
	}
	if args == nil || args.RecoveryVaultName == nil {
		return nil, errors.New("missing required argument 'RecoveryVaultName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SourceRecoveryFabricName == nil {
		return nil, errors.New("missing required argument 'SourceRecoveryFabricName'")
	}
	if args == nil || args.SourceRecoveryProtectionContainerName == nil {
		return nil, errors.New("missing required argument 'SourceRecoveryProtectionContainerName'")
	}
	if args == nil || args.SourceVmId == nil {
		return nil, errors.New("missing required argument 'SourceVmId'")
	}
	if args == nil || args.TargetRecoveryFabricId == nil {
		return nil, errors.New("missing required argument 'TargetRecoveryFabricId'")
	}
	if args == nil || args.TargetRecoveryProtectionContainerId == nil {
		return nil, errors.New("missing required argument 'TargetRecoveryProtectionContainerId'")
	}
	if args == nil || args.TargetResourceGroupId == nil {
		return nil, errors.New("missing required argument 'TargetResourceGroupId'")
	}
	if args == nil {
		args = &ReplicatedVMArgs{}
	}
	var resource ReplicatedVM
	err := ctx.RegisterResource("azure:siterecovery/replicatedVM:ReplicatedVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicatedVM gets an existing ReplicatedVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicatedVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicatedVMState, opts ...pulumi.ResourceOption) (*ReplicatedVM, error) {
	var resource ReplicatedVM
	err := ctx.ReadResource("azure:siterecovery/replicatedVM:ReplicatedVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicatedVM resources.
type replicatedVMState struct {
	// One or more `managedDisk` block.
	ManagedDisks []ReplicatedVMManagedDisk `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name                        *string `pulumi:"name"`
	RecoveryReplicationPolicyId *string `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName *string `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName *string `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId *string `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId *string `pulumi:"targetAvailabilitySetId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId *string `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId *string `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId *string `pulumi:"targetResourceGroupId"`
}

type ReplicatedVMState struct {
	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVMManagedDiskArrayInput
	// The name of the network mapping.
	Name                        pulumi.StringPtrInput
	RecoveryReplicationPolicyId pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringPtrInput
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringPtrInput
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringPtrInput
	// Id of the VM to replicate
	SourceVmId pulumi.StringPtrInput
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringPtrInput
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringPtrInput
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringPtrInput
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringPtrInput
}

func (ReplicatedVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicatedVMState)(nil)).Elem()
}

type replicatedVMArgs struct {
	// One or more `managedDisk` block.
	ManagedDisks []ReplicatedVMManagedDisk `pulumi:"managedDisks"`
	// The name of the network mapping.
	Name                        *string `pulumi:"name"`
	RecoveryReplicationPolicyId string  `pulumi:"recoveryReplicationPolicyId"`
	// The name of the vault that should be updated.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName string `pulumi:"sourceRecoveryFabricName"`
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName string `pulumi:"sourceRecoveryProtectionContainerName"`
	// Id of the VM to replicate
	SourceVmId string `pulumi:"sourceVmId"`
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId *string `pulumi:"targetAvailabilitySetId"`
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId string `pulumi:"targetRecoveryFabricId"`
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId string `pulumi:"targetRecoveryProtectionContainerId"`
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId string `pulumi:"targetResourceGroupId"`
}

// The set of arguments for constructing a ReplicatedVM resource.
type ReplicatedVMArgs struct {
	// One or more `managedDisk` block.
	ManagedDisks ReplicatedVMManagedDiskArrayInput
	// The name of the network mapping.
	Name                        pulumi.StringPtrInput
	RecoveryReplicationPolicyId pulumi.StringInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput
	// Name of fabric that should contains this replication.
	SourceRecoveryFabricName pulumi.StringInput
	// Name of the protection container to use.
	SourceRecoveryProtectionContainerName pulumi.StringInput
	// Id of the VM to replicate
	SourceVmId pulumi.StringInput
	// Id of availability set that the new VM should belong to when a failover is done.
	TargetAvailabilitySetId pulumi.StringPtrInput
	// Id of fabric where the VM replication should be handled when a failover is done.
	TargetRecoveryFabricId pulumi.StringInput
	// Id of protection container where the VM replication should be created when a failover is done.
	TargetRecoveryProtectionContainerId pulumi.StringInput
	// Id of resource group where the VM should be created when a failover is done.
	TargetResourceGroupId pulumi.StringInput
}

func (ReplicatedVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicatedVMArgs)(nil)).Elem()
}
