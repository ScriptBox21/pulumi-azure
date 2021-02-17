// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Network Connection Monitor.
//
// > **NOTE:** Any Network Connection Monitor resource created with API versions 2019-06-01 or earlier (v1) are now incompatible with this provider, which now only supports v2.
//
// ## Import
//
// Network Connection Monitors can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkConnectionMonitor:NetworkConnectionMonitor example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkWatchers/watcher1/connectionMonitors/connectionMonitor1
// ```
type NetworkConnectionMonitor struct {
	pulumi.CustomResourceState

	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart pulumi.BoolOutput `pulumi:"autoStart"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination NetworkConnectionMonitorDestinationOutput `pulumi:"destination"`
	// A `endpoint` block as defined below.
	Endpoints NetworkConnectionMonitorEndpointArrayOutput `pulumi:"endpoints"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds pulumi.IntOutput `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId pulumi.StringOutput `pulumi:"networkWatcherId"`
	// The description of the Network Connection Monitor.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds pulumi.StringArrayOutput `pulumi:"outputWorkspaceResourceIds"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source NetworkConnectionMonitorSourceOutput `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `testConfiguration` block as defined below.
	TestConfigurations NetworkConnectionMonitorTestConfigurationArrayOutput `pulumi:"testConfigurations"`
	// A `testGroup` block as defined below.
	TestGroups NetworkConnectionMonitorTestGroupArrayOutput `pulumi:"testGroups"`
}

// NewNetworkConnectionMonitor registers a new resource with the given unique name, arguments, and options.
func NewNetworkConnectionMonitor(ctx *pulumi.Context,
	name string, args *NetworkConnectionMonitorArgs, opts ...pulumi.ResourceOption) (*NetworkConnectionMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoints == nil {
		return nil, errors.New("invalid value for required argument 'Endpoints'")
	}
	if args.NetworkWatcherId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherId'")
	}
	if args.TestConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'TestConfigurations'")
	}
	if args.TestGroups == nil {
		return nil, errors.New("invalid value for required argument 'TestGroups'")
	}
	var resource NetworkConnectionMonitor
	err := ctx.RegisterResource("azure:network/networkConnectionMonitor:NetworkConnectionMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkConnectionMonitor gets an existing NetworkConnectionMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkConnectionMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkConnectionMonitorState, opts ...pulumi.ResourceOption) (*NetworkConnectionMonitor, error) {
	var resource NetworkConnectionMonitor
	err := ctx.ReadResource("azure:network/networkConnectionMonitor:NetworkConnectionMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkConnectionMonitor resources.
type networkConnectionMonitorState struct {
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart *bool `pulumi:"autoStart"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination *NetworkConnectionMonitorDestination `pulumi:"destination"`
	// A `endpoint` block as defined below.
	Endpoints []NetworkConnectionMonitorEndpoint `pulumi:"endpoints"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId *string `pulumi:"networkWatcherId"`
	// The description of the Network Connection Monitor.
	Notes *string `pulumi:"notes"`
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds []string `pulumi:"outputWorkspaceResourceIds"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source *NetworkConnectionMonitorSource `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags map[string]string `pulumi:"tags"`
	// A `testConfiguration` block as defined below.
	TestConfigurations []NetworkConnectionMonitorTestConfiguration `pulumi:"testConfigurations"`
	// A `testGroup` block as defined below.
	TestGroups []NetworkConnectionMonitorTestGroup `pulumi:"testGroups"`
}

type NetworkConnectionMonitorState struct {
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart pulumi.BoolPtrInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination NetworkConnectionMonitorDestinationPtrInput
	// A `endpoint` block as defined below.
	Endpoints NetworkConnectionMonitorEndpointArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds pulumi.IntPtrInput
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId pulumi.StringPtrInput
	// The description of the Network Connection Monitor.
	Notes pulumi.StringPtrInput
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds pulumi.StringArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source NetworkConnectionMonitorSourcePtrInput
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapInput
	// A `testConfiguration` block as defined below.
	TestConfigurations NetworkConnectionMonitorTestConfigurationArrayInput
	// A `testGroup` block as defined below.
	TestGroups NetworkConnectionMonitorTestGroupArrayInput
}

func (NetworkConnectionMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionMonitorState)(nil)).Elem()
}

type networkConnectionMonitorArgs struct {
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart *bool `pulumi:"autoStart"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination *NetworkConnectionMonitorDestination `pulumi:"destination"`
	// A `endpoint` block as defined below.
	Endpoints []NetworkConnectionMonitorEndpoint `pulumi:"endpoints"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId string `pulumi:"networkWatcherId"`
	// The description of the Network Connection Monitor.
	Notes *string `pulumi:"notes"`
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds []string `pulumi:"outputWorkspaceResourceIds"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source *NetworkConnectionMonitorSource `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags map[string]string `pulumi:"tags"`
	// A `testConfiguration` block as defined below.
	TestConfigurations []NetworkConnectionMonitorTestConfiguration `pulumi:"testConfigurations"`
	// A `testGroup` block as defined below.
	TestGroups []NetworkConnectionMonitorTestGroup `pulumi:"testGroups"`
}

// The set of arguments for constructing a NetworkConnectionMonitor resource.
type NetworkConnectionMonitorArgs struct {
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart pulumi.BoolPtrInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination NetworkConnectionMonitorDestinationPtrInput
	// A `endpoint` block as defined below.
	Endpoints NetworkConnectionMonitorEndpointArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds pulumi.IntPtrInput
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId pulumi.StringInput
	// The description of the Network Connection Monitor.
	Notes pulumi.StringPtrInput
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds pulumi.StringArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source NetworkConnectionMonitorSourcePtrInput
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapInput
	// A `testConfiguration` block as defined below.
	TestConfigurations NetworkConnectionMonitorTestConfigurationArrayInput
	// A `testGroup` block as defined below.
	TestGroups NetworkConnectionMonitorTestGroupArrayInput
}

func (NetworkConnectionMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionMonitorArgs)(nil)).Elem()
}

type NetworkConnectionMonitorInput interface {
	pulumi.Input

	ToNetworkConnectionMonitorOutput() NetworkConnectionMonitorOutput
	ToNetworkConnectionMonitorOutputWithContext(ctx context.Context) NetworkConnectionMonitorOutput
}

func (*NetworkConnectionMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConnectionMonitor)(nil))
}

func (i *NetworkConnectionMonitor) ToNetworkConnectionMonitorOutput() NetworkConnectionMonitorOutput {
	return i.ToNetworkConnectionMonitorOutputWithContext(context.Background())
}

func (i *NetworkConnectionMonitor) ToNetworkConnectionMonitorOutputWithContext(ctx context.Context) NetworkConnectionMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionMonitorOutput)
}

func (i *NetworkConnectionMonitor) ToNetworkConnectionMonitorPtrOutput() NetworkConnectionMonitorPtrOutput {
	return i.ToNetworkConnectionMonitorPtrOutputWithContext(context.Background())
}

func (i *NetworkConnectionMonitor) ToNetworkConnectionMonitorPtrOutputWithContext(ctx context.Context) NetworkConnectionMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionMonitorPtrOutput)
}

type NetworkConnectionMonitorPtrInput interface {
	pulumi.Input

	ToNetworkConnectionMonitorPtrOutput() NetworkConnectionMonitorPtrOutput
	ToNetworkConnectionMonitorPtrOutputWithContext(ctx context.Context) NetworkConnectionMonitorPtrOutput
}

type networkConnectionMonitorPtrType NetworkConnectionMonitorArgs

func (*networkConnectionMonitorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnectionMonitor)(nil))
}

func (i *networkConnectionMonitorPtrType) ToNetworkConnectionMonitorPtrOutput() NetworkConnectionMonitorPtrOutput {
	return i.ToNetworkConnectionMonitorPtrOutputWithContext(context.Background())
}

func (i *networkConnectionMonitorPtrType) ToNetworkConnectionMonitorPtrOutputWithContext(ctx context.Context) NetworkConnectionMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionMonitorPtrOutput)
}

// NetworkConnectionMonitorArrayInput is an input type that accepts NetworkConnectionMonitorArray and NetworkConnectionMonitorArrayOutput values.
// You can construct a concrete instance of `NetworkConnectionMonitorArrayInput` via:
//
//          NetworkConnectionMonitorArray{ NetworkConnectionMonitorArgs{...} }
type NetworkConnectionMonitorArrayInput interface {
	pulumi.Input

	ToNetworkConnectionMonitorArrayOutput() NetworkConnectionMonitorArrayOutput
	ToNetworkConnectionMonitorArrayOutputWithContext(context.Context) NetworkConnectionMonitorArrayOutput
}

type NetworkConnectionMonitorArray []NetworkConnectionMonitorInput

func (NetworkConnectionMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkConnectionMonitor)(nil))
}

func (i NetworkConnectionMonitorArray) ToNetworkConnectionMonitorArrayOutput() NetworkConnectionMonitorArrayOutput {
	return i.ToNetworkConnectionMonitorArrayOutputWithContext(context.Background())
}

func (i NetworkConnectionMonitorArray) ToNetworkConnectionMonitorArrayOutputWithContext(ctx context.Context) NetworkConnectionMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionMonitorArrayOutput)
}

// NetworkConnectionMonitorMapInput is an input type that accepts NetworkConnectionMonitorMap and NetworkConnectionMonitorMapOutput values.
// You can construct a concrete instance of `NetworkConnectionMonitorMapInput` via:
//
//          NetworkConnectionMonitorMap{ "key": NetworkConnectionMonitorArgs{...} }
type NetworkConnectionMonitorMapInput interface {
	pulumi.Input

	ToNetworkConnectionMonitorMapOutput() NetworkConnectionMonitorMapOutput
	ToNetworkConnectionMonitorMapOutputWithContext(context.Context) NetworkConnectionMonitorMapOutput
}

type NetworkConnectionMonitorMap map[string]NetworkConnectionMonitorInput

func (NetworkConnectionMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkConnectionMonitor)(nil))
}

func (i NetworkConnectionMonitorMap) ToNetworkConnectionMonitorMapOutput() NetworkConnectionMonitorMapOutput {
	return i.ToNetworkConnectionMonitorMapOutputWithContext(context.Background())
}

func (i NetworkConnectionMonitorMap) ToNetworkConnectionMonitorMapOutputWithContext(ctx context.Context) NetworkConnectionMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionMonitorMapOutput)
}

type NetworkConnectionMonitorOutput struct {
	*pulumi.OutputState
}

func (NetworkConnectionMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConnectionMonitor)(nil))
}

func (o NetworkConnectionMonitorOutput) ToNetworkConnectionMonitorOutput() NetworkConnectionMonitorOutput {
	return o
}

func (o NetworkConnectionMonitorOutput) ToNetworkConnectionMonitorOutputWithContext(ctx context.Context) NetworkConnectionMonitorOutput {
	return o
}

func (o NetworkConnectionMonitorOutput) ToNetworkConnectionMonitorPtrOutput() NetworkConnectionMonitorPtrOutput {
	return o.ToNetworkConnectionMonitorPtrOutputWithContext(context.Background())
}

func (o NetworkConnectionMonitorOutput) ToNetworkConnectionMonitorPtrOutputWithContext(ctx context.Context) NetworkConnectionMonitorPtrOutput {
	return o.ApplyT(func(v NetworkConnectionMonitor) *NetworkConnectionMonitor {
		return &v
	}).(NetworkConnectionMonitorPtrOutput)
}

type NetworkConnectionMonitorPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkConnectionMonitorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnectionMonitor)(nil))
}

func (o NetworkConnectionMonitorPtrOutput) ToNetworkConnectionMonitorPtrOutput() NetworkConnectionMonitorPtrOutput {
	return o
}

func (o NetworkConnectionMonitorPtrOutput) ToNetworkConnectionMonitorPtrOutputWithContext(ctx context.Context) NetworkConnectionMonitorPtrOutput {
	return o
}

type NetworkConnectionMonitorArrayOutput struct{ *pulumi.OutputState }

func (NetworkConnectionMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkConnectionMonitor)(nil))
}

func (o NetworkConnectionMonitorArrayOutput) ToNetworkConnectionMonitorArrayOutput() NetworkConnectionMonitorArrayOutput {
	return o
}

func (o NetworkConnectionMonitorArrayOutput) ToNetworkConnectionMonitorArrayOutputWithContext(ctx context.Context) NetworkConnectionMonitorArrayOutput {
	return o
}

func (o NetworkConnectionMonitorArrayOutput) Index(i pulumi.IntInput) NetworkConnectionMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkConnectionMonitor {
		return vs[0].([]NetworkConnectionMonitor)[vs[1].(int)]
	}).(NetworkConnectionMonitorOutput)
}

type NetworkConnectionMonitorMapOutput struct{ *pulumi.OutputState }

func (NetworkConnectionMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkConnectionMonitor)(nil))
}

func (o NetworkConnectionMonitorMapOutput) ToNetworkConnectionMonitorMapOutput() NetworkConnectionMonitorMapOutput {
	return o
}

func (o NetworkConnectionMonitorMapOutput) ToNetworkConnectionMonitorMapOutputWithContext(ctx context.Context) NetworkConnectionMonitorMapOutput {
	return o
}

func (o NetworkConnectionMonitorMapOutput) MapIndex(k pulumi.StringInput) NetworkConnectionMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkConnectionMonitor {
		return vs[0].(map[string]NetworkConnectionMonitor)[vs[1].(string)]
	}).(NetworkConnectionMonitorOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkConnectionMonitorOutput{})
	pulumi.RegisterOutputType(NetworkConnectionMonitorPtrOutput{})
	pulumi.RegisterOutputType(NetworkConnectionMonitorArrayOutput{})
	pulumi.RegisterOutputType(NetworkConnectionMonitorMapOutput{})
}
