// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package devtest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Windows Virtual Machine within a Dev Test Lab.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/dev_test_windows_virtual_machine.html.markdown.
type WindowsVirtualMachine struct {
	pulumi.CustomResourceState

	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrOutput `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrOutput `pulumi:"disallowPublicIpAddress"`
	// The FQDN of the Virtual Machine.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReferenceOutput `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules WindowsVirtualMachineInboundNatRuleArrayOutput `pulumi:"inboundNatRules"`
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringOutput `pulumi:"labName"`
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName pulumi.StringOutput `pulumi:"labSubnetName"`
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId pulumi.StringOutput `pulumi:"labVirtualNetworkId"`
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Any notes about the Virtual Machine.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password pulumi.StringOutput `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringOutput `pulumi:"size"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewWindowsVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewWindowsVirtualMachine(ctx *pulumi.Context,
	name string, args *WindowsVirtualMachineArgs, opts ...pulumi.ResourceOption) (*WindowsVirtualMachine, error) {
	if args == nil || args.GalleryImageReference == nil {
		return nil, errors.New("missing required argument 'GalleryImageReference'")
	}
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.LabSubnetName == nil {
		return nil, errors.New("missing required argument 'LabSubnetName'")
	}
	if args == nil || args.LabVirtualNetworkId == nil {
		return nil, errors.New("missing required argument 'LabVirtualNetworkId'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil || args.StorageType == nil {
		return nil, errors.New("missing required argument 'StorageType'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &WindowsVirtualMachineArgs{}
	}
	var resource WindowsVirtualMachine
	err := ctx.RegisterResource("azure:devtest/windowsVirtualMachine:WindowsVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWindowsVirtualMachine gets an existing WindowsVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWindowsVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WindowsVirtualMachineState, opts ...pulumi.ResourceOption) (*WindowsVirtualMachine, error) {
	var resource WindowsVirtualMachine
	err := ctx.ReadResource("azure:devtest/windowsVirtualMachine:WindowsVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WindowsVirtualMachine resources.
type windowsVirtualMachineState struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim *bool `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// The FQDN of the Virtual Machine.
	Fqdn *string `pulumi:"fqdn"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference *WindowsVirtualMachineGalleryImageReference `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules []WindowsVirtualMachineInboundNatRule `pulumi:"inboundNatRules"`
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName *string `pulumi:"labName"`
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName *string `pulumi:"labSubnetName"`
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId *string `pulumi:"labVirtualNetworkId"`
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Any notes about the Virtual Machine.
	Notes *string `pulumi:"notes"`
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size *string `pulumi:"size"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType *string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username *string `pulumi:"username"`
}

type WindowsVirtualMachineState struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrInput
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrInput
	// The FQDN of the Virtual Machine.
	Fqdn pulumi.StringPtrInput
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReferencePtrInput
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules WindowsVirtualMachineInboundNatRuleArrayInput
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringPtrInput
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName pulumi.StringPtrInput
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId pulumi.StringPtrInput
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Any notes about the Virtual Machine.
	Notes pulumi.StringPtrInput
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringPtrInput
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier pulumi.StringPtrInput
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringPtrInput
}

func (WindowsVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsVirtualMachineState)(nil)).Elem()
}

type windowsVirtualMachineArgs struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim *bool `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReference `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules []WindowsVirtualMachineInboundNatRule `pulumi:"inboundNatRules"`
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName string `pulumi:"labName"`
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName string `pulumi:"labSubnetName"`
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId string `pulumi:"labVirtualNetworkId"`
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Any notes about the Virtual Machine.
	Notes *string `pulumi:"notes"`
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password string `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size string `pulumi:"size"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a WindowsVirtualMachine resource.
type WindowsVirtualMachineArgs struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrInput
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrInput
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReferenceInput
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules WindowsVirtualMachineInboundNatRuleArrayInput
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringInput
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName pulumi.StringInput
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId pulumi.StringInput
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Any notes about the Virtual Machine.
	Notes pulumi.StringPtrInput
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password pulumi.StringInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringInput
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringInput
}

func (WindowsVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsVirtualMachineArgs)(nil)).Elem()
}
