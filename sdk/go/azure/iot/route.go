// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IotHub Route
// 
// > **NOTE:** Routes can be defined either directly on the `iot.IoTHub` resource, or using the `iot.Route` resourcs - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/iothub_route.html.markdown.
type Route struct {
	pulumi.CustomResourceState

	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition pulumi.StringPtrOutput `pulumi:"condition"`
	// Specifies whether a route is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames pulumi.StringOutput `pulumi:"endpointNames"`
	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringOutput `pulumi:"iothubName"`
	// The name of the route.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
	Source pulumi.StringOutput `pulumi:"source"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.EndpointNames == nil {
		return nil, errors.New("missing required argument 'EndpointNames'")
	}
	if args == nil || args.IothubName == nil {
		return nil, errors.New("missing required argument 'IothubName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Source == nil {
		return nil, errors.New("missing required argument 'Source'")
	}
	if args == nil {
		args = &RouteArgs{}
	}
	var resource Route
	err := ctx.RegisterResource("azure:iot/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("azure:iot/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `pulumi:"condition"`
	// Specifies whether a route is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames *string `pulumi:"endpointNames"`
	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IothubName *string `pulumi:"iothubName"`
	// The name of the route.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
	Source *string `pulumi:"source"`
}

type RouteState struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition pulumi.StringPtrInput
	// Specifies whether a route is enabled.
	Enabled pulumi.BoolPtrInput
	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames pulumi.StringPtrInput
	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringPtrInput
	// The name of the route.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
	Source pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `pulumi:"condition"`
	// Specifies whether a route is enabled.
	Enabled bool `pulumi:"enabled"`
	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames string `pulumi:"endpointNames"`
	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IothubName string `pulumi:"iothubName"`
	// The name of the route.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
	Source string `pulumi:"source"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition pulumi.StringPtrInput
	// Specifies whether a route is enabled.
	Enabled pulumi.BoolInput
	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames pulumi.StringInput
	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringInput
	// The name of the route.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The source that the routing rule is to be applied to. Possible values include: `DeviceJobLifecycleEvents`, `DeviceLifecycleEvents`, `DeviceMessages`, `Invalid`, `TwinChangeEvents`.
	Source pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

