// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a managed disk.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/managed_disk.html.markdown.
type ManagedDisk struct {
	pulumi.CustomResourceState

	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
	CreateOption pulumi.StringOutput `pulumi:"createOption"`
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk. Changing this forces a new resource to be created.
	DiskEncryptionSetId pulumi.StringPtrOutput `pulumi:"diskEncryptionSetId"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite pulumi.IntOutput `pulumi:"diskIopsReadWrite"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite pulumi.IntOutput `pulumi:"diskMbpsReadWrite"`
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
	DiskSizeGb pulumi.IntOutput `pulumi:"diskSizeGb"`
	// A `encryptionSettings` block as defined below.
	EncryptionSettings ManagedDiskEncryptionSettingsPtrOutput `pulumi:"encryptionSettings"`
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId pulumi.StringPtrOutput `pulumi:"imageReferenceId"`
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId pulumi.StringPtrOutput `pulumi:"sourceResourceId"`
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri pulumi.StringOutput `pulumi:"sourceUri"`
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType pulumi.StringOutput `pulumi:"storageAccountType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewManagedDisk registers a new resource with the given unique name, arguments, and options.
func NewManagedDisk(ctx *pulumi.Context,
	name string, args *ManagedDiskArgs, opts ...pulumi.ResourceOption) (*ManagedDisk, error) {
	if args == nil || args.CreateOption == nil {
		return nil, errors.New("missing required argument 'CreateOption'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageAccountType == nil {
		return nil, errors.New("missing required argument 'StorageAccountType'")
	}
	if args == nil {
		args = &ManagedDiskArgs{}
	}
	var resource ManagedDisk
	err := ctx.RegisterResource("azure:compute/managedDisk:ManagedDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedDisk gets an existing ManagedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedDiskState, opts ...pulumi.ResourceOption) (*ManagedDisk, error) {
	var resource ManagedDisk
	err := ctx.ReadResource("azure:compute/managedDisk:ManagedDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedDisk resources.
type managedDiskState struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
	CreateOption *string `pulumi:"createOption"`
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk. Changing this forces a new resource to be created.
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite *int `pulumi:"diskIopsReadWrite"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite *int `pulumi:"diskMbpsReadWrite"`
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// A `encryptionSettings` block as defined below.
	EncryptionSettings *ManagedDiskEncryptionSettings `pulumi:"encryptionSettings"`
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId *string `pulumi:"imageReferenceId"`
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType *string `pulumi:"osType"`
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri *string `pulumi:"sourceUri"`
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType *string `pulumi:"storageAccountType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones *string `pulumi:"zones"`
}

type ManagedDiskState struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
	CreateOption pulumi.StringPtrInput
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk. Changing this forces a new resource to be created.
	DiskEncryptionSetId pulumi.StringPtrInput
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite pulumi.IntPtrInput
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite pulumi.IntPtrInput
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
	DiskSizeGb pulumi.IntPtrInput
	// A `encryptionSettings` block as defined below.
	EncryptionSettings ManagedDiskEncryptionSettingsPtrInput
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId pulumi.StringPtrInput
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri pulumi.StringPtrInput
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones pulumi.StringPtrInput
}

func (ManagedDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDiskState)(nil)).Elem()
}

type managedDiskArgs struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
	CreateOption string `pulumi:"createOption"`
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk. Changing this forces a new resource to be created.
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite *int `pulumi:"diskIopsReadWrite"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite *int `pulumi:"diskMbpsReadWrite"`
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
	DiskSizeGb *int `pulumi:"diskSizeGb"`
	// A `encryptionSettings` block as defined below.
	EncryptionSettings *ManagedDiskEncryptionSettings `pulumi:"encryptionSettings"`
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId *string `pulumi:"imageReferenceId"`
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType *string `pulumi:"osType"`
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri *string `pulumi:"sourceUri"`
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType string `pulumi:"storageAccountType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a ManagedDisk resource.
type ManagedDiskArgs struct {
	// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include:
	CreateOption pulumi.StringInput
	// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk. Changing this forces a new resource to be created.
	DiskEncryptionSetId pulumi.StringPtrInput
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIopsReadWrite pulumi.IntPtrInput
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
	DiskMbpsReadWrite pulumi.IntPtrInput
	// Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size.
	DiskSizeGb pulumi.IntPtrInput
	// A `encryptionSettings` block as defined below.
	EncryptionSettings ManagedDiskEncryptionSettingsPtrInput
	// ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
	ImageReferenceId pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
	OsType pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Disk should exist.
	ResourceGroupName pulumi.StringInput
	// The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
	SourceResourceId pulumi.StringPtrInput
	// URI to a valid VHD file to be used when `createOption` is `Import`.
	SourceUri pulumi.StringPtrInput
	// The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
	StorageAccountType pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A collection containing the availability zone to allocate the Managed Disk in.
	Zones pulumi.StringPtrInput
}

func (ManagedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDiskArgs)(nil)).Elem()
}
