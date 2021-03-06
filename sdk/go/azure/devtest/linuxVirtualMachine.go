// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linux Virtual Machine within a Dev Test Lab.
//
// ## Import
//
// Dev Test Linux Virtual Machines can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:devtest/linuxVirtualMachine:LinuxVirtualMachine machine1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
// ```
type LinuxVirtualMachine struct {
	pulumi.CustomResourceState

	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrOutput `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrOutput `pulumi:"disallowPublicIpAddress"`
	// The FQDN of the Virtual Machine.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference LinuxVirtualMachineGalleryImageReferenceOutput `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules LinuxVirtualMachineInboundNatRuleArrayOutput `pulumi:"inboundNatRules"`
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
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringOutput `pulumi:"size"`
	// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	SshKey pulumi.StringPtrOutput `pulumi:"sshKey"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewLinuxVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewLinuxVirtualMachine(ctx *pulumi.Context,
	name string, args *LinuxVirtualMachineArgs, opts ...pulumi.ResourceOption) (*LinuxVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryImageReference == nil {
		return nil, errors.New("invalid value for required argument 'GalleryImageReference'")
	}
	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.LabSubnetName == nil {
		return nil, errors.New("invalid value for required argument 'LabSubnetName'")
	}
	if args.LabVirtualNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'LabVirtualNetworkId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource LinuxVirtualMachine
	err := ctx.RegisterResource("azure:devtest/linuxVirtualMachine:LinuxVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinuxVirtualMachine gets an existing LinuxVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinuxVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinuxVirtualMachineState, opts ...pulumi.ResourceOption) (*LinuxVirtualMachine, error) {
	var resource LinuxVirtualMachine
	err := ctx.ReadResource("azure:devtest/linuxVirtualMachine:LinuxVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinuxVirtualMachine resources.
type linuxVirtualMachineState struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim *bool `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// The FQDN of the Virtual Machine.
	Fqdn *string `pulumi:"fqdn"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference *LinuxVirtualMachineGalleryImageReference `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules []LinuxVirtualMachineInboundNatRule `pulumi:"inboundNatRules"`
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
	// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	SshKey *string `pulumi:"sshKey"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType *string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username *string `pulumi:"username"`
}

type LinuxVirtualMachineState struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrInput
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrInput
	// The FQDN of the Virtual Machine.
	Fqdn pulumi.StringPtrInput
	// A `galleryImageReference` block as defined below.
	GalleryImageReference LinuxVirtualMachineGalleryImageReferencePtrInput
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules LinuxVirtualMachineInboundNatRuleArrayInput
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
	// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	SshKey pulumi.StringPtrInput
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier pulumi.StringPtrInput
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringPtrInput
}

func (LinuxVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*linuxVirtualMachineState)(nil)).Elem()
}

type linuxVirtualMachineArgs struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim *bool `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference LinuxVirtualMachineGalleryImageReference `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules []LinuxVirtualMachineInboundNatRule `pulumi:"inboundNatRules"`
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
	Password *string `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size string `pulumi:"size"`
	// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	SshKey *string `pulumi:"sshKey"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a LinuxVirtualMachine resource.
type LinuxVirtualMachineArgs struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrInput
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrInput
	// A `galleryImageReference` block as defined below.
	GalleryImageReference LinuxVirtualMachineGalleryImageReferenceInput
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules LinuxVirtualMachineInboundNatRuleArrayInput
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
	Password pulumi.StringPtrInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringInput
	// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	SshKey pulumi.StringPtrInput
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringInput
}

func (LinuxVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linuxVirtualMachineArgs)(nil)).Elem()
}

type LinuxVirtualMachineInput interface {
	pulumi.Input

	ToLinuxVirtualMachineOutput() LinuxVirtualMachineOutput
	ToLinuxVirtualMachineOutputWithContext(ctx context.Context) LinuxVirtualMachineOutput
}

func (*LinuxVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxVirtualMachine)(nil))
}

func (i *LinuxVirtualMachine) ToLinuxVirtualMachineOutput() LinuxVirtualMachineOutput {
	return i.ToLinuxVirtualMachineOutputWithContext(context.Background())
}

func (i *LinuxVirtualMachine) ToLinuxVirtualMachineOutputWithContext(ctx context.Context) LinuxVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVirtualMachineOutput)
}

func (i *LinuxVirtualMachine) ToLinuxVirtualMachinePtrOutput() LinuxVirtualMachinePtrOutput {
	return i.ToLinuxVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *LinuxVirtualMachine) ToLinuxVirtualMachinePtrOutputWithContext(ctx context.Context) LinuxVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVirtualMachinePtrOutput)
}

type LinuxVirtualMachinePtrInput interface {
	pulumi.Input

	ToLinuxVirtualMachinePtrOutput() LinuxVirtualMachinePtrOutput
	ToLinuxVirtualMachinePtrOutputWithContext(ctx context.Context) LinuxVirtualMachinePtrOutput
}

type linuxVirtualMachinePtrType LinuxVirtualMachineArgs

func (*linuxVirtualMachinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxVirtualMachine)(nil))
}

func (i *linuxVirtualMachinePtrType) ToLinuxVirtualMachinePtrOutput() LinuxVirtualMachinePtrOutput {
	return i.ToLinuxVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *linuxVirtualMachinePtrType) ToLinuxVirtualMachinePtrOutputWithContext(ctx context.Context) LinuxVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVirtualMachinePtrOutput)
}

// LinuxVirtualMachineArrayInput is an input type that accepts LinuxVirtualMachineArray and LinuxVirtualMachineArrayOutput values.
// You can construct a concrete instance of `LinuxVirtualMachineArrayInput` via:
//
//          LinuxVirtualMachineArray{ LinuxVirtualMachineArgs{...} }
type LinuxVirtualMachineArrayInput interface {
	pulumi.Input

	ToLinuxVirtualMachineArrayOutput() LinuxVirtualMachineArrayOutput
	ToLinuxVirtualMachineArrayOutputWithContext(context.Context) LinuxVirtualMachineArrayOutput
}

type LinuxVirtualMachineArray []LinuxVirtualMachineInput

func (LinuxVirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinuxVirtualMachine)(nil))
}

func (i LinuxVirtualMachineArray) ToLinuxVirtualMachineArrayOutput() LinuxVirtualMachineArrayOutput {
	return i.ToLinuxVirtualMachineArrayOutputWithContext(context.Background())
}

func (i LinuxVirtualMachineArray) ToLinuxVirtualMachineArrayOutputWithContext(ctx context.Context) LinuxVirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVirtualMachineArrayOutput)
}

// LinuxVirtualMachineMapInput is an input type that accepts LinuxVirtualMachineMap and LinuxVirtualMachineMapOutput values.
// You can construct a concrete instance of `LinuxVirtualMachineMapInput` via:
//
//          LinuxVirtualMachineMap{ "key": LinuxVirtualMachineArgs{...} }
type LinuxVirtualMachineMapInput interface {
	pulumi.Input

	ToLinuxVirtualMachineMapOutput() LinuxVirtualMachineMapOutput
	ToLinuxVirtualMachineMapOutputWithContext(context.Context) LinuxVirtualMachineMapOutput
}

type LinuxVirtualMachineMap map[string]LinuxVirtualMachineInput

func (LinuxVirtualMachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinuxVirtualMachine)(nil))
}

func (i LinuxVirtualMachineMap) ToLinuxVirtualMachineMapOutput() LinuxVirtualMachineMapOutput {
	return i.ToLinuxVirtualMachineMapOutputWithContext(context.Background())
}

func (i LinuxVirtualMachineMap) ToLinuxVirtualMachineMapOutputWithContext(ctx context.Context) LinuxVirtualMachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVirtualMachineMapOutput)
}

type LinuxVirtualMachineOutput struct {
	*pulumi.OutputState
}

func (LinuxVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxVirtualMachine)(nil))
}

func (o LinuxVirtualMachineOutput) ToLinuxVirtualMachineOutput() LinuxVirtualMachineOutput {
	return o
}

func (o LinuxVirtualMachineOutput) ToLinuxVirtualMachineOutputWithContext(ctx context.Context) LinuxVirtualMachineOutput {
	return o
}

func (o LinuxVirtualMachineOutput) ToLinuxVirtualMachinePtrOutput() LinuxVirtualMachinePtrOutput {
	return o.ToLinuxVirtualMachinePtrOutputWithContext(context.Background())
}

func (o LinuxVirtualMachineOutput) ToLinuxVirtualMachinePtrOutputWithContext(ctx context.Context) LinuxVirtualMachinePtrOutput {
	return o.ApplyT(func(v LinuxVirtualMachine) *LinuxVirtualMachine {
		return &v
	}).(LinuxVirtualMachinePtrOutput)
}

type LinuxVirtualMachinePtrOutput struct {
	*pulumi.OutputState
}

func (LinuxVirtualMachinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxVirtualMachine)(nil))
}

func (o LinuxVirtualMachinePtrOutput) ToLinuxVirtualMachinePtrOutput() LinuxVirtualMachinePtrOutput {
	return o
}

func (o LinuxVirtualMachinePtrOutput) ToLinuxVirtualMachinePtrOutputWithContext(ctx context.Context) LinuxVirtualMachinePtrOutput {
	return o
}

type LinuxVirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (LinuxVirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinuxVirtualMachine)(nil))
}

func (o LinuxVirtualMachineArrayOutput) ToLinuxVirtualMachineArrayOutput() LinuxVirtualMachineArrayOutput {
	return o
}

func (o LinuxVirtualMachineArrayOutput) ToLinuxVirtualMachineArrayOutputWithContext(ctx context.Context) LinuxVirtualMachineArrayOutput {
	return o
}

func (o LinuxVirtualMachineArrayOutput) Index(i pulumi.IntInput) LinuxVirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinuxVirtualMachine {
		return vs[0].([]LinuxVirtualMachine)[vs[1].(int)]
	}).(LinuxVirtualMachineOutput)
}

type LinuxVirtualMachineMapOutput struct{ *pulumi.OutputState }

func (LinuxVirtualMachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinuxVirtualMachine)(nil))
}

func (o LinuxVirtualMachineMapOutput) ToLinuxVirtualMachineMapOutput() LinuxVirtualMachineMapOutput {
	return o
}

func (o LinuxVirtualMachineMapOutput) ToLinuxVirtualMachineMapOutputWithContext(ctx context.Context) LinuxVirtualMachineMapOutput {
	return o
}

func (o LinuxVirtualMachineMapOutput) MapIndex(k pulumi.StringInput) LinuxVirtualMachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinuxVirtualMachine {
		return vs[0].(map[string]LinuxVirtualMachine)[vs[1].(string)]
	}).(LinuxVirtualMachineOutput)
}

func init() {
	pulumi.RegisterOutputType(LinuxVirtualMachineOutput{})
	pulumi.RegisterOutputType(LinuxVirtualMachinePtrOutput{})
	pulumi.RegisterOutputType(LinuxVirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(LinuxVirtualMachineMapOutput{})
}
