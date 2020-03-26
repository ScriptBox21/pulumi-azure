// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package devtest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Virtual Network within a DevTest Lab.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/dev_test_virtual_network.html.markdown.
type VirtualNetwork struct {
	pulumi.CustomResourceState

	// A description for the Virtual Network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringOutput `pulumi:"labName"`
	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `subnet` block as defined below.
	Subnet VirtualNetworkSubnetOutput `pulumi:"subnet"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The unique immutable identifier of the Dev Test Virtual Network.
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
}

// NewVirtualNetwork registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualNetworkArgs{}
	}
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure:devtest/virtualNetwork:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetwork gets an existing VirtualNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azure:devtest/virtualNetwork:VirtualNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetwork resources.
type virtualNetworkState struct {
	// A description for the Virtual Network.
	Description *string `pulumi:"description"`
	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName *string `pulumi:"labName"`
	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `subnet` block as defined below.
	Subnet *VirtualNetworkSubnet `pulumi:"subnet"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of the Dev Test Virtual Network.
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

type VirtualNetworkState struct {
	// A description for the Virtual Network.
	Description pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringPtrInput
	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `subnet` block as defined below.
	Subnet VirtualNetworkSubnetPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of the Dev Test Virtual Network.
	UniqueIdentifier pulumi.StringPtrInput
}

func (VirtualNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkState)(nil)).Elem()
}

type virtualNetworkArgs struct {
	// A description for the Virtual Network.
	Description *string `pulumi:"description"`
	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName string `pulumi:"labName"`
	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `subnet` block as defined below.
	Subnet *VirtualNetworkSubnet `pulumi:"subnet"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualNetwork resource.
type VirtualNetworkArgs struct {
	// A description for the Virtual Network.
	Description pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringInput
	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `subnet` block as defined below.
	Subnet VirtualNetworkSubnetPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkArgs)(nil)).Elem()
}
