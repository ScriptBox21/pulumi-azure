// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package lb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Load Balancer Backend Address Pool.
//
// > **NOTE:** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb_backend_address_pool.html.markdown.
type BackendAddressPool struct {
	pulumi.CustomResourceState

	// The Backend IP Configurations associated with this Backend Address Pool.
	BackendIpConfigurations pulumi.StringArrayOutput `pulumi:"backendIpConfigurations"`
	// The Load Balancing Rules associated with this Backend Address Pool.
	LoadBalancingRules pulumi.StringArrayOutput `pulumi:"loadBalancingRules"`
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the name of the Backend Address Pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewBackendAddressPool registers a new resource with the given unique name, arguments, and options.
func NewBackendAddressPool(ctx *pulumi.Context,
	name string, args *BackendAddressPoolArgs, opts ...pulumi.ResourceOption) (*BackendAddressPool, error) {
	if args == nil || args.LoadbalancerId == nil {
		return nil, errors.New("missing required argument 'LoadbalancerId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &BackendAddressPoolArgs{}
	}
	var resource BackendAddressPool
	err := ctx.RegisterResource("azure:lb/backendAddressPool:BackendAddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendAddressPool gets an existing BackendAddressPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendAddressPoolState, opts ...pulumi.ResourceOption) (*BackendAddressPool, error) {
	var resource BackendAddressPool
	err := ctx.ReadResource("azure:lb/backendAddressPool:BackendAddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendAddressPool resources.
type backendAddressPoolState struct {
	// The Backend IP Configurations associated with this Backend Address Pool.
	BackendIpConfigurations []string `pulumi:"backendIpConfigurations"`
	// The Load Balancing Rules associated with this Backend Address Pool.
	LoadBalancingRules []string `pulumi:"loadBalancingRules"`
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the Backend Address Pool.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type BackendAddressPoolState struct {
	// The Backend IP Configurations associated with this Backend Address Pool.
	BackendIpConfigurations pulumi.StringArrayInput
	// The Load Balancing Rules associated with this Backend Address Pool.
	LoadBalancingRules pulumi.StringArrayInput
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the name of the Backend Address Pool.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringPtrInput
}

func (BackendAddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendAddressPoolState)(nil)).Elem()
}

type backendAddressPoolArgs struct {
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the Backend Address Pool.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a BackendAddressPool resource.
type BackendAddressPoolArgs struct {
	// The ID of the Load Balancer in which to create the Backend Address Pool.
	LoadbalancerId pulumi.StringInput
	// Specifies the name of the Backend Address Pool.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringInput
}

func (BackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendAddressPoolArgs)(nil)).Elem()
}
