// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual Hub Route Table.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.5.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNetworkSecurityGroup, err := network.NewNetworkSecurityGroup(ctx, "exampleNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.5.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSubnetNetworkSecurityGroupAssociation(ctx, "exampleSubnetNetworkSecurityGroupAssociation", &network.SubnetNetworkSecurityGroupAssociationArgs{
// 			SubnetId:               exampleSubnet.ID(),
// 			NetworkSecurityGroupId: exampleNetworkSecurityGroup.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualWan, err := network.NewVirtualWan(ctx, "exampleVirtualWan", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			AddressPrefix:     pulumi.String("10.0.2.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualHubConnection, err := network.NewVirtualHubConnection(ctx, "exampleVirtualHubConnection", &network.VirtualHubConnectionArgs{
// 			VirtualHubId:           exampleVirtualHub.ID(),
// 			RemoteVirtualNetworkId: exampleVirtualNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualHubRouteTable(ctx, "exampleVirtualHubRouteTable", &network.VirtualHubRouteTableArgs{
// 			VirtualHubId: exampleVirtualHub.ID(),
// 			Labels: pulumi.StringArray{
// 				pulumi.String("label1"),
// 			},
// 			Routes: network.VirtualHubRouteTableRouteArray{
// 				&network.VirtualHubRouteTableRouteArgs{
// 					Name:             pulumi.String("example-route"),
// 					DestinationsType: pulumi.String("CIDR"),
// 					Destinations: pulumi.StringArray{
// 						pulumi.String("10.0.0.0/16"),
// 					},
// 					NextHopType: pulumi.String("ResourceId"),
// 					NextHop:     exampleVirtualHubConnection.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Virtual Hub Route Tables can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/virtualHubRouteTable:VirtualHubRouteTable example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/routeTable1
// ```
type VirtualHubRouteTable struct {
	pulumi.CustomResourceState

	// List of labels associated with this route table.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The name which should be used for Virtual Hub Route Table. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `route` block as defined below.
	Routes VirtualHubRouteTableRouteArrayOutput `pulumi:"routes"`
	// The ID of the Virtual Hub within which this route table should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
}

// NewVirtualHubRouteTable registers a new resource with the given unique name, arguments, and options.
func NewVirtualHubRouteTable(ctx *pulumi.Context,
	name string, args *VirtualHubRouteTableArgs, opts ...pulumi.ResourceOption) (*VirtualHubRouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource VirtualHubRouteTable
	err := ctx.RegisterResource("azure:network/virtualHubRouteTable:VirtualHubRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHubRouteTable gets an existing VirtualHubRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHubRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubRouteTableState, opts ...pulumi.ResourceOption) (*VirtualHubRouteTable, error) {
	var resource VirtualHubRouteTable
	err := ctx.ReadResource("azure:network/virtualHubRouteTable:VirtualHubRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHubRouteTable resources.
type virtualHubRouteTableState struct {
	// List of labels associated with this route table.
	Labels []string `pulumi:"labels"`
	// The name which should be used for Virtual Hub Route Table. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `route` block as defined below.
	Routes []VirtualHubRouteTableRoute `pulumi:"routes"`
	// The ID of the Virtual Hub within which this route table should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

type VirtualHubRouteTableState struct {
	// List of labels associated with this route table.
	Labels pulumi.StringArrayInput
	// The name which should be used for Virtual Hub Route Table. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `route` block as defined below.
	Routes VirtualHubRouteTableRouteArrayInput
	// The ID of the Virtual Hub within which this route table should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
}

func (VirtualHubRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubRouteTableState)(nil)).Elem()
}

type virtualHubRouteTableArgs struct {
	// List of labels associated with this route table.
	Labels []string `pulumi:"labels"`
	// The name which should be used for Virtual Hub Route Table. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `route` block as defined below.
	Routes []VirtualHubRouteTableRoute `pulumi:"routes"`
	// The ID of the Virtual Hub within which this route table should be created. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
}

// The set of arguments for constructing a VirtualHubRouteTable resource.
type VirtualHubRouteTableArgs struct {
	// List of labels associated with this route table.
	Labels pulumi.StringArrayInput
	// The name which should be used for Virtual Hub Route Table. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `route` block as defined below.
	Routes VirtualHubRouteTableRouteArrayInput
	// The ID of the Virtual Hub within which this route table should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
}

func (VirtualHubRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubRouteTableArgs)(nil)).Elem()
}

type VirtualHubRouteTableInput interface {
	pulumi.Input

	ToVirtualHubRouteTableOutput() VirtualHubRouteTableOutput
	ToVirtualHubRouteTableOutputWithContext(ctx context.Context) VirtualHubRouteTableOutput
}

func (*VirtualHubRouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTable)(nil))
}

func (i *VirtualHubRouteTable) ToVirtualHubRouteTableOutput() VirtualHubRouteTableOutput {
	return i.ToVirtualHubRouteTableOutputWithContext(context.Background())
}

func (i *VirtualHubRouteTable) ToVirtualHubRouteTableOutputWithContext(ctx context.Context) VirtualHubRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableOutput)
}

func (i *VirtualHubRouteTable) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return i.ToVirtualHubRouteTablePtrOutputWithContext(context.Background())
}

func (i *VirtualHubRouteTable) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTablePtrOutput)
}

type VirtualHubRouteTablePtrInput interface {
	pulumi.Input

	ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput
	ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput
}

type virtualHubRouteTablePtrType VirtualHubRouteTableArgs

func (*virtualHubRouteTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubRouteTable)(nil))
}

func (i *virtualHubRouteTablePtrType) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return i.ToVirtualHubRouteTablePtrOutputWithContext(context.Background())
}

func (i *virtualHubRouteTablePtrType) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTablePtrOutput)
}

// VirtualHubRouteTableArrayInput is an input type that accepts VirtualHubRouteTableArray and VirtualHubRouteTableArrayOutput values.
// You can construct a concrete instance of `VirtualHubRouteTableArrayInput` via:
//
//          VirtualHubRouteTableArray{ VirtualHubRouteTableArgs{...} }
type VirtualHubRouteTableArrayInput interface {
	pulumi.Input

	ToVirtualHubRouteTableArrayOutput() VirtualHubRouteTableArrayOutput
	ToVirtualHubRouteTableArrayOutputWithContext(context.Context) VirtualHubRouteTableArrayOutput
}

type VirtualHubRouteTableArray []VirtualHubRouteTableInput

func (VirtualHubRouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VirtualHubRouteTable)(nil))
}

func (i VirtualHubRouteTableArray) ToVirtualHubRouteTableArrayOutput() VirtualHubRouteTableArrayOutput {
	return i.ToVirtualHubRouteTableArrayOutputWithContext(context.Background())
}

func (i VirtualHubRouteTableArray) ToVirtualHubRouteTableArrayOutputWithContext(ctx context.Context) VirtualHubRouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableArrayOutput)
}

// VirtualHubRouteTableMapInput is an input type that accepts VirtualHubRouteTableMap and VirtualHubRouteTableMapOutput values.
// You can construct a concrete instance of `VirtualHubRouteTableMapInput` via:
//
//          VirtualHubRouteTableMap{ "key": VirtualHubRouteTableArgs{...} }
type VirtualHubRouteTableMapInput interface {
	pulumi.Input

	ToVirtualHubRouteTableMapOutput() VirtualHubRouteTableMapOutput
	ToVirtualHubRouteTableMapOutputWithContext(context.Context) VirtualHubRouteTableMapOutput
}

type VirtualHubRouteTableMap map[string]VirtualHubRouteTableInput

func (VirtualHubRouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VirtualHubRouteTable)(nil))
}

func (i VirtualHubRouteTableMap) ToVirtualHubRouteTableMapOutput() VirtualHubRouteTableMapOutput {
	return i.ToVirtualHubRouteTableMapOutputWithContext(context.Background())
}

func (i VirtualHubRouteTableMap) ToVirtualHubRouteTableMapOutputWithContext(ctx context.Context) VirtualHubRouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableMapOutput)
}

type VirtualHubRouteTableOutput struct {
	*pulumi.OutputState
}

func (VirtualHubRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTable)(nil))
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTableOutput() VirtualHubRouteTableOutput {
	return o
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTableOutputWithContext(ctx context.Context) VirtualHubRouteTableOutput {
	return o
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return o.ToVirtualHubRouteTablePtrOutputWithContext(context.Background())
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return o.ApplyT(func(v VirtualHubRouteTable) *VirtualHubRouteTable {
		return &v
	}).(VirtualHubRouteTablePtrOutput)
}

type VirtualHubRouteTablePtrOutput struct {
	*pulumi.OutputState
}

func (VirtualHubRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubRouteTable)(nil))
}

func (o VirtualHubRouteTablePtrOutput) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return o
}

func (o VirtualHubRouteTablePtrOutput) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return o
}

type VirtualHubRouteTableArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteTable)(nil))
}

func (o VirtualHubRouteTableArrayOutput) ToVirtualHubRouteTableArrayOutput() VirtualHubRouteTableArrayOutput {
	return o
}

func (o VirtualHubRouteTableArrayOutput) ToVirtualHubRouteTableArrayOutputWithContext(ctx context.Context) VirtualHubRouteTableArrayOutput {
	return o
}

func (o VirtualHubRouteTableArrayOutput) Index(i pulumi.IntInput) VirtualHubRouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubRouteTable {
		return vs[0].([]VirtualHubRouteTable)[vs[1].(int)]
	}).(VirtualHubRouteTableOutput)
}

type VirtualHubRouteTableMapOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualHubRouteTable)(nil))
}

func (o VirtualHubRouteTableMapOutput) ToVirtualHubRouteTableMapOutput() VirtualHubRouteTableMapOutput {
	return o
}

func (o VirtualHubRouteTableMapOutput) ToVirtualHubRouteTableMapOutputWithContext(ctx context.Context) VirtualHubRouteTableMapOutput {
	return o
}

func (o VirtualHubRouteTableMapOutput) MapIndex(k pulumi.StringInput) VirtualHubRouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualHubRouteTable {
		return vs[0].(map[string]VirtualHubRouteTable)[vs[1].(string)]
	}).(VirtualHubRouteTableOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubRouteTableOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTablePtrOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableMapOutput{})
}
