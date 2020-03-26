// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Virtual Machine.
//
// ## Disclaimers
//
// > **Note:** The `compute.VirtualMachine` resource has been superseded by the `compute.LinuxVirtualMachine` and `compute.WindowsVirtualMachine` resources. The existing `compute.VirtualMachine` resource will continue to be available throughout the 2.x releases however is in a feature-frozen state to maintain compatibility - new functionality will instead be added to the `compute.LinuxVirtualMachine` and `compute.WindowsVirtualMachine` resources.
//
// > **Note:** Data Disks can be attached either directly on the `compute.VirtualMachine` resource, or using the `compute.DataDiskAttachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_machine.html.markdown.
type VirtualMachine struct {
	pulumi.CustomResourceState

	// A `additionalCapabilities` block.
	AdditionalCapabilities VirtualMachineAdditionalCapabilitiesPtrOutput `pulumi:"additionalCapabilities"`
	// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringOutput `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block.
	BootDiagnostics VirtualMachineBootDiagnosticsPtrOutput `pulumi:"bootDiagnostics"`
	// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteDataDisksOnTermination pulumi.BoolPtrOutput `pulumi:"deleteDataDisksOnTermination"`
	// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteOsDiskOnTermination pulumi.BoolPtrOutput `pulumi:"deleteOsDiskOnTermination"`
	// A `identity` block.
	Identity VirtualMachineIdentityOutput `pulumi:"identity"`
	// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of Network Interface ID's which should be associated with the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// An `osProfile` block. Required when `createOption` in the `storageOsDisk` block is set to `FromImage`.
	OsProfile VirtualMachineOsProfilePtrOutput `pulumi:"osProfile"`
	// A `osProfileLinuxConfig` block.
	OsProfileLinuxConfig VirtualMachineOsProfileLinuxConfigPtrOutput `pulumi:"osProfileLinuxConfig"`
	// One or more `osProfileSecrets` blocks.
	OsProfileSecrets VirtualMachineOsProfileSecretArrayOutput `pulumi:"osProfileSecrets"`
	// A `osProfileWindowsConfig` block.
	OsProfileWindowsConfig VirtualMachineOsProfileWindowsConfigPtrOutput `pulumi:"osProfileWindowsConfig"`
	// A `plan` block.
	Plan VirtualMachinePlanPtrOutput `pulumi:"plan"`
	// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
	PrimaryNetworkInterfaceId pulumi.StringPtrOutput `pulumi:"primaryNetworkInterfaceId"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `storageDataDisk` blocks.
	StorageDataDisks VirtualMachineStorageDataDiskArrayOutput `pulumi:"storageDataDisks"`
	// A `storageImageReference` block.
	StorageImageReference VirtualMachineStorageImageReferenceOutput `pulumi:"storageImageReference"`
	// A `storageOsDisk` block.
	StorageOsDisk VirtualMachineStorageOsDiskOutput `pulumi:"storageOsDisk"`
	// A mapping of tags to assign to the Virtual Machine.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
	VmSize pulumi.StringOutput `pulumi:"vmSize"`
	// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil || args.NetworkInterfaceIds == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceIds'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageOsDisk == nil {
		return nil, errors.New("missing required argument 'StorageOsDisk'")
	}
	if args == nil || args.VmSize == nil {
		return nil, errors.New("missing required argument 'VmSize'")
	}
	if args == nil {
		args = &VirtualMachineArgs{}
	}
	var resource VirtualMachine
	err := ctx.RegisterResource("azure:compute/virtualMachine:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachine gets an existing VirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure:compute/virtualMachine:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachine resources.
type virtualMachineState struct {
	// A `additionalCapabilities` block.
	AdditionalCapabilities *VirtualMachineAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId *string `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block.
	BootDiagnostics *VirtualMachineBootDiagnostics `pulumi:"bootDiagnostics"`
	// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteDataDisksOnTermination *bool `pulumi:"deleteDataDisksOnTermination"`
	// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteOsDiskOnTermination *bool `pulumi:"deleteOsDiskOnTermination"`
	// A `identity` block.
	Identity *VirtualMachineIdentity `pulumi:"identity"`
	// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of Network Interface ID's which should be associated with the Virtual Machine.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// An `osProfile` block. Required when `createOption` in the `storageOsDisk` block is set to `FromImage`.
	OsProfile *VirtualMachineOsProfile `pulumi:"osProfile"`
	// A `osProfileLinuxConfig` block.
	OsProfileLinuxConfig *VirtualMachineOsProfileLinuxConfig `pulumi:"osProfileLinuxConfig"`
	// One or more `osProfileSecrets` blocks.
	OsProfileSecrets []VirtualMachineOsProfileSecret `pulumi:"osProfileSecrets"`
	// A `osProfileWindowsConfig` block.
	OsProfileWindowsConfig *VirtualMachineOsProfileWindowsConfig `pulumi:"osProfileWindowsConfig"`
	// A `plan` block.
	Plan *VirtualMachinePlan `pulumi:"plan"`
	// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
	PrimaryNetworkInterfaceId *string `pulumi:"primaryNetworkInterfaceId"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `storageDataDisk` blocks.
	StorageDataDisks []VirtualMachineStorageDataDisk `pulumi:"storageDataDisks"`
	// A `storageImageReference` block.
	StorageImageReference *VirtualMachineStorageImageReference `pulumi:"storageImageReference"`
	// A `storageOsDisk` block.
	StorageOsDisk *VirtualMachineStorageOsDisk `pulumi:"storageOsDisk"`
	// A mapping of tags to assign to the Virtual Machine.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
	VmSize *string `pulumi:"vmSize"`
	// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
	Zones *string `pulumi:"zones"`
}

type VirtualMachineState struct {
	// A `additionalCapabilities` block.
	AdditionalCapabilities VirtualMachineAdditionalCapabilitiesPtrInput
	// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrInput
	// A `bootDiagnostics` block.
	BootDiagnostics VirtualMachineBootDiagnosticsPtrInput
	// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteDataDisksOnTermination pulumi.BoolPtrInput
	// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteOsDiskOnTermination pulumi.BoolPtrInput
	// A `identity` block.
	Identity VirtualMachineIdentityPtrInput
	// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
	LicenseType pulumi.StringPtrInput
	// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of Network Interface ID's which should be associated with the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayInput
	// An `osProfile` block. Required when `createOption` in the `storageOsDisk` block is set to `FromImage`.
	OsProfile VirtualMachineOsProfilePtrInput
	// A `osProfileLinuxConfig` block.
	OsProfileLinuxConfig VirtualMachineOsProfileLinuxConfigPtrInput
	// One or more `osProfileSecrets` blocks.
	OsProfileSecrets VirtualMachineOsProfileSecretArrayInput
	// A `osProfileWindowsConfig` block.
	OsProfileWindowsConfig VirtualMachineOsProfileWindowsConfigPtrInput
	// A `plan` block.
	Plan VirtualMachinePlanPtrInput
	// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
	PrimaryNetworkInterfaceId pulumi.StringPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `storageDataDisk` blocks.
	StorageDataDisks VirtualMachineStorageDataDiskArrayInput
	// A `storageImageReference` block.
	StorageImageReference VirtualMachineStorageImageReferencePtrInput
	// A `storageOsDisk` block.
	StorageOsDisk VirtualMachineStorageOsDiskPtrInput
	// A mapping of tags to assign to the Virtual Machine.
	Tags pulumi.StringMapInput
	// Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
	VmSize pulumi.StringPtrInput
	// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
	Zones pulumi.StringPtrInput
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	// A `additionalCapabilities` block.
	AdditionalCapabilities *VirtualMachineAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId *string `pulumi:"availabilitySetId"`
	// A `bootDiagnostics` block.
	BootDiagnostics *VirtualMachineBootDiagnostics `pulumi:"bootDiagnostics"`
	// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteDataDisksOnTermination *bool `pulumi:"deleteDataDisksOnTermination"`
	// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteOsDiskOnTermination *bool `pulumi:"deleteOsDiskOnTermination"`
	// A `identity` block.
	Identity *VirtualMachineIdentity `pulumi:"identity"`
	// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of Network Interface ID's which should be associated with the Virtual Machine.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// An `osProfile` block. Required when `createOption` in the `storageOsDisk` block is set to `FromImage`.
	OsProfile *VirtualMachineOsProfile `pulumi:"osProfile"`
	// A `osProfileLinuxConfig` block.
	OsProfileLinuxConfig *VirtualMachineOsProfileLinuxConfig `pulumi:"osProfileLinuxConfig"`
	// One or more `osProfileSecrets` blocks.
	OsProfileSecrets []VirtualMachineOsProfileSecret `pulumi:"osProfileSecrets"`
	// A `osProfileWindowsConfig` block.
	OsProfileWindowsConfig *VirtualMachineOsProfileWindowsConfig `pulumi:"osProfileWindowsConfig"`
	// A `plan` block.
	Plan *VirtualMachinePlan `pulumi:"plan"`
	// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
	PrimaryNetworkInterfaceId *string `pulumi:"primaryNetworkInterfaceId"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `storageDataDisk` blocks.
	StorageDataDisks []VirtualMachineStorageDataDisk `pulumi:"storageDataDisks"`
	// A `storageImageReference` block.
	StorageImageReference *VirtualMachineStorageImageReference `pulumi:"storageImageReference"`
	// A `storageOsDisk` block.
	StorageOsDisk VirtualMachineStorageOsDisk `pulumi:"storageOsDisk"`
	// A mapping of tags to assign to the Virtual Machine.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
	VmSize string `pulumi:"vmSize"`
	// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// A `additionalCapabilities` block.
	AdditionalCapabilities VirtualMachineAdditionalCapabilitiesPtrInput
	// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	AvailabilitySetId pulumi.StringPtrInput
	// A `bootDiagnostics` block.
	BootDiagnostics VirtualMachineBootDiagnosticsPtrInput
	// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteDataDisksOnTermination pulumi.BoolPtrInput
	// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
	DeleteOsDiskOnTermination pulumi.BoolPtrInput
	// A `identity` block.
	Identity VirtualMachineIdentityPtrInput
	// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
	LicenseType pulumi.StringPtrInput
	// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of Network Interface ID's which should be associated with the Virtual Machine.
	NetworkInterfaceIds pulumi.StringArrayInput
	// An `osProfile` block. Required when `createOption` in the `storageOsDisk` block is set to `FromImage`.
	OsProfile VirtualMachineOsProfilePtrInput
	// A `osProfileLinuxConfig` block.
	OsProfileLinuxConfig VirtualMachineOsProfileLinuxConfigPtrInput
	// One or more `osProfileSecrets` blocks.
	OsProfileSecrets VirtualMachineOsProfileSecretArrayInput
	// A `osProfileWindowsConfig` block.
	OsProfileWindowsConfig VirtualMachineOsProfileWindowsConfigPtrInput
	// A `plan` block.
	Plan VirtualMachinePlanPtrInput
	// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
	PrimaryNetworkInterfaceId pulumi.StringPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `storageDataDisk` blocks.
	StorageDataDisks VirtualMachineStorageDataDiskArrayInput
	// A `storageImageReference` block.
	StorageImageReference VirtualMachineStorageImageReferencePtrInput
	// A `storageOsDisk` block.
	StorageOsDisk VirtualMachineStorageOsDiskInput
	// A mapping of tags to assign to the Virtual Machine.
	Tags pulumi.StringMapInput
	// Specifies the [size of the Virtual Machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).
	VmSize pulumi.StringInput
	// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
	Zones pulumi.StringPtrInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}
