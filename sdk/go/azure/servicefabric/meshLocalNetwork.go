// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Service Fabric Mesh Local Network can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicefabric/meshLocalNetwork:MeshLocalNetwork network1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabricMesh/networks/network1
// ```
type MeshLocalNetwork struct {
	pulumi.CustomResourceState

	// A description of this Service Fabric Mesh Local Network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The address space for the local container network.
	NetworkAddressPrefix pulumi.StringOutput `pulumi:"networkAddressPrefix"`
	// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewMeshLocalNetwork registers a new resource with the given unique name, arguments, and options.
func NewMeshLocalNetwork(ctx *pulumi.Context,
	name string, args *MeshLocalNetworkArgs, opts ...pulumi.ResourceOption) (*MeshLocalNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkAddressPrefix == nil {
		return nil, errors.New("invalid value for required argument 'NetworkAddressPrefix'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource MeshLocalNetwork
	err := ctx.RegisterResource("azure:servicefabric/meshLocalNetwork:MeshLocalNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMeshLocalNetwork gets an existing MeshLocalNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMeshLocalNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeshLocalNetworkState, opts ...pulumi.ResourceOption) (*MeshLocalNetwork, error) {
	var resource MeshLocalNetwork
	err := ctx.ReadResource("azure:servicefabric/meshLocalNetwork:MeshLocalNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MeshLocalNetwork resources.
type meshLocalNetworkState struct {
	// A description of this Service Fabric Mesh Local Network.
	Description *string `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The address space for the local container network.
	NetworkAddressPrefix *string `pulumi:"networkAddressPrefix"`
	// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type MeshLocalNetworkState struct {
	// A description of this Service Fabric Mesh Local Network.
	Description pulumi.StringPtrInput
	// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The address space for the local container network.
	NetworkAddressPrefix pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshLocalNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*meshLocalNetworkState)(nil)).Elem()
}

type meshLocalNetworkArgs struct {
	// A description of this Service Fabric Mesh Local Network.
	Description *string `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The address space for the local container network.
	NetworkAddressPrefix string `pulumi:"networkAddressPrefix"`
	// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MeshLocalNetwork resource.
type MeshLocalNetworkArgs struct {
	// A description of this Service Fabric Mesh Local Network.
	Description pulumi.StringPtrInput
	// Specifies the Azure Region where the Service Fabric Mesh Local Network should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Local Network. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The address space for the local container network.
	NetworkAddressPrefix pulumi.StringInput
	// The name of the Resource Group in which the Service Fabric Mesh Local Network exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshLocalNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meshLocalNetworkArgs)(nil)).Elem()
}

type MeshLocalNetworkInput interface {
	pulumi.Input

	ToMeshLocalNetworkOutput() MeshLocalNetworkOutput
	ToMeshLocalNetworkOutputWithContext(ctx context.Context) MeshLocalNetworkOutput
}

func (*MeshLocalNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshLocalNetwork)(nil))
}

func (i *MeshLocalNetwork) ToMeshLocalNetworkOutput() MeshLocalNetworkOutput {
	return i.ToMeshLocalNetworkOutputWithContext(context.Background())
}

func (i *MeshLocalNetwork) ToMeshLocalNetworkOutputWithContext(ctx context.Context) MeshLocalNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshLocalNetworkOutput)
}

func (i *MeshLocalNetwork) ToMeshLocalNetworkPtrOutput() MeshLocalNetworkPtrOutput {
	return i.ToMeshLocalNetworkPtrOutputWithContext(context.Background())
}

func (i *MeshLocalNetwork) ToMeshLocalNetworkPtrOutputWithContext(ctx context.Context) MeshLocalNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshLocalNetworkPtrOutput)
}

type MeshLocalNetworkPtrInput interface {
	pulumi.Input

	ToMeshLocalNetworkPtrOutput() MeshLocalNetworkPtrOutput
	ToMeshLocalNetworkPtrOutputWithContext(ctx context.Context) MeshLocalNetworkPtrOutput
}

type meshLocalNetworkPtrType MeshLocalNetworkArgs

func (*meshLocalNetworkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MeshLocalNetwork)(nil))
}

func (i *meshLocalNetworkPtrType) ToMeshLocalNetworkPtrOutput() MeshLocalNetworkPtrOutput {
	return i.ToMeshLocalNetworkPtrOutputWithContext(context.Background())
}

func (i *meshLocalNetworkPtrType) ToMeshLocalNetworkPtrOutputWithContext(ctx context.Context) MeshLocalNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshLocalNetworkPtrOutput)
}

// MeshLocalNetworkArrayInput is an input type that accepts MeshLocalNetworkArray and MeshLocalNetworkArrayOutput values.
// You can construct a concrete instance of `MeshLocalNetworkArrayInput` via:
//
//          MeshLocalNetworkArray{ MeshLocalNetworkArgs{...} }
type MeshLocalNetworkArrayInput interface {
	pulumi.Input

	ToMeshLocalNetworkArrayOutput() MeshLocalNetworkArrayOutput
	ToMeshLocalNetworkArrayOutputWithContext(context.Context) MeshLocalNetworkArrayOutput
}

type MeshLocalNetworkArray []MeshLocalNetworkInput

func (MeshLocalNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MeshLocalNetwork)(nil))
}

func (i MeshLocalNetworkArray) ToMeshLocalNetworkArrayOutput() MeshLocalNetworkArrayOutput {
	return i.ToMeshLocalNetworkArrayOutputWithContext(context.Background())
}

func (i MeshLocalNetworkArray) ToMeshLocalNetworkArrayOutputWithContext(ctx context.Context) MeshLocalNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshLocalNetworkArrayOutput)
}

// MeshLocalNetworkMapInput is an input type that accepts MeshLocalNetworkMap and MeshLocalNetworkMapOutput values.
// You can construct a concrete instance of `MeshLocalNetworkMapInput` via:
//
//          MeshLocalNetworkMap{ "key": MeshLocalNetworkArgs{...} }
type MeshLocalNetworkMapInput interface {
	pulumi.Input

	ToMeshLocalNetworkMapOutput() MeshLocalNetworkMapOutput
	ToMeshLocalNetworkMapOutputWithContext(context.Context) MeshLocalNetworkMapOutput
}

type MeshLocalNetworkMap map[string]MeshLocalNetworkInput

func (MeshLocalNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MeshLocalNetwork)(nil))
}

func (i MeshLocalNetworkMap) ToMeshLocalNetworkMapOutput() MeshLocalNetworkMapOutput {
	return i.ToMeshLocalNetworkMapOutputWithContext(context.Background())
}

func (i MeshLocalNetworkMap) ToMeshLocalNetworkMapOutputWithContext(ctx context.Context) MeshLocalNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshLocalNetworkMapOutput)
}

type MeshLocalNetworkOutput struct {
	*pulumi.OutputState
}

func (MeshLocalNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshLocalNetwork)(nil))
}

func (o MeshLocalNetworkOutput) ToMeshLocalNetworkOutput() MeshLocalNetworkOutput {
	return o
}

func (o MeshLocalNetworkOutput) ToMeshLocalNetworkOutputWithContext(ctx context.Context) MeshLocalNetworkOutput {
	return o
}

func (o MeshLocalNetworkOutput) ToMeshLocalNetworkPtrOutput() MeshLocalNetworkPtrOutput {
	return o.ToMeshLocalNetworkPtrOutputWithContext(context.Background())
}

func (o MeshLocalNetworkOutput) ToMeshLocalNetworkPtrOutputWithContext(ctx context.Context) MeshLocalNetworkPtrOutput {
	return o.ApplyT(func(v MeshLocalNetwork) *MeshLocalNetwork {
		return &v
	}).(MeshLocalNetworkPtrOutput)
}

type MeshLocalNetworkPtrOutput struct {
	*pulumi.OutputState
}

func (MeshLocalNetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MeshLocalNetwork)(nil))
}

func (o MeshLocalNetworkPtrOutput) ToMeshLocalNetworkPtrOutput() MeshLocalNetworkPtrOutput {
	return o
}

func (o MeshLocalNetworkPtrOutput) ToMeshLocalNetworkPtrOutputWithContext(ctx context.Context) MeshLocalNetworkPtrOutput {
	return o
}

type MeshLocalNetworkArrayOutput struct{ *pulumi.OutputState }

func (MeshLocalNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MeshLocalNetwork)(nil))
}

func (o MeshLocalNetworkArrayOutput) ToMeshLocalNetworkArrayOutput() MeshLocalNetworkArrayOutput {
	return o
}

func (o MeshLocalNetworkArrayOutput) ToMeshLocalNetworkArrayOutputWithContext(ctx context.Context) MeshLocalNetworkArrayOutput {
	return o
}

func (o MeshLocalNetworkArrayOutput) Index(i pulumi.IntInput) MeshLocalNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MeshLocalNetwork {
		return vs[0].([]MeshLocalNetwork)[vs[1].(int)]
	}).(MeshLocalNetworkOutput)
}

type MeshLocalNetworkMapOutput struct{ *pulumi.OutputState }

func (MeshLocalNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MeshLocalNetwork)(nil))
}

func (o MeshLocalNetworkMapOutput) ToMeshLocalNetworkMapOutput() MeshLocalNetworkMapOutput {
	return o
}

func (o MeshLocalNetworkMapOutput) ToMeshLocalNetworkMapOutputWithContext(ctx context.Context) MeshLocalNetworkMapOutput {
	return o
}

func (o MeshLocalNetworkMapOutput) MapIndex(k pulumi.StringInput) MeshLocalNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MeshLocalNetwork {
		return vs[0].(map[string]MeshLocalNetwork)[vs[1].(string)]
	}).(MeshLocalNetworkOutput)
}

func init() {
	pulumi.RegisterOutputType(MeshLocalNetworkOutput{})
	pulumi.RegisterOutputType(MeshLocalNetworkPtrOutput{})
	pulumi.RegisterOutputType(MeshLocalNetworkArrayOutput{})
	pulumi.RegisterOutputType(MeshLocalNetworkMapOutput{})
}
