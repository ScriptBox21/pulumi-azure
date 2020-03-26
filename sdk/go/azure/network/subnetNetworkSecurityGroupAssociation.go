// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associates a Network Security Group with a Subnet within a Virtual Network.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/subnet_network_security_group_association.html.markdown.
type SubnetNetworkSecurityGroupAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Network Security Group which should be associated with the Subnet. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringOutput `pulumi:"networkSecurityGroupId"`
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewSubnetNetworkSecurityGroupAssociation registers a new resource with the given unique name, arguments, and options.
func NewSubnetNetworkSecurityGroupAssociation(ctx *pulumi.Context,
	name string, args *SubnetNetworkSecurityGroupAssociationArgs, opts ...pulumi.ResourceOption) (*SubnetNetworkSecurityGroupAssociation, error) {
	if args == nil || args.NetworkSecurityGroupId == nil {
		return nil, errors.New("missing required argument 'NetworkSecurityGroupId'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil {
		args = &SubnetNetworkSecurityGroupAssociationArgs{}
	}
	var resource SubnetNetworkSecurityGroupAssociation
	err := ctx.RegisterResource("azure:network/subnetNetworkSecurityGroupAssociation:SubnetNetworkSecurityGroupAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetNetworkSecurityGroupAssociation gets an existing SubnetNetworkSecurityGroupAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetNetworkSecurityGroupAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetNetworkSecurityGroupAssociationState, opts ...pulumi.ResourceOption) (*SubnetNetworkSecurityGroupAssociation, error) {
	var resource SubnetNetworkSecurityGroupAssociation
	err := ctx.ReadResource("azure:network/subnetNetworkSecurityGroupAssociation:SubnetNetworkSecurityGroupAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetNetworkSecurityGroupAssociation resources.
type subnetNetworkSecurityGroupAssociationState struct {
	// The ID of the Network Security Group which should be associated with the Subnet. Changing this forces a new resource to be created.
	NetworkSecurityGroupId *string `pulumi:"networkSecurityGroupId"`
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
}

type SubnetNetworkSecurityGroupAssociationState struct {
	// The ID of the Network Security Group which should be associated with the Subnet. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringPtrInput
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
}

func (SubnetNetworkSecurityGroupAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetNetworkSecurityGroupAssociationState)(nil)).Elem()
}

type subnetNetworkSecurityGroupAssociationArgs struct {
	// The ID of the Network Security Group which should be associated with the Subnet. Changing this forces a new resource to be created.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SubnetNetworkSecurityGroupAssociation resource.
type SubnetNetworkSecurityGroupAssociationArgs struct {
	// The ID of the Network Security Group which should be associated with the Subnet. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringInput
	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
}

func (SubnetNetworkSecurityGroupAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetNetworkSecurityGroupAssociationArgs)(nil)).Elem()
}
