// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Network Interface.
type NetworkInterface struct {
	pulumi.CustomResourceState

	// If the Virtual Machine using this Network Interface is part of an Availability Set, then this list will have the union of all DNS servers from all Network Interfaces that are part of the Availability Set.
	AppliedDnsServers pulumi.StringArrayOutput `pulumi:"appliedDnsServers"`
	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	DnsServers pulumi.StringArrayOutput `pulumi:"dnsServers"`
	// Should Accelerated Networking be enabled? Defaults to `false`.
	EnableAcceleratedNetworking pulumi.BoolPtrOutput `pulumi:"enableAcceleratedNetworking"`
	// Should IP Forwarding be enabled? Defaults to `false`.
	EnableIpForwarding pulumi.BoolPtrOutput `pulumi:"enableIpForwarding"`
	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	InternalDnsNameLabel pulumi.StringOutput `pulumi:"internalDnsNameLabel"`
	// One or more `ipConfiguration` blocks as defined below.
	IpConfigurations NetworkInterfaceIpConfigurationArrayOutput `pulumi:"ipConfigurations"`
	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Media Access Control (MAC) Address of the Network Interface.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The name of the Network Interface. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Static IP Address which should be used.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// The private IP addresses of the network interface.
	PrivateIpAddresses pulumi.StringArrayOutput `pulumi:"privateIpAddresses"`
	// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Virtual Machine which this Network Interface is connected to.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil || args.IpConfigurations == nil {
		return nil, errors.New("missing required argument 'IpConfigurations'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkInterfaceArgs{}
	}
	var resource NetworkInterface
	err := ctx.RegisterResource("azure:network/networkInterface:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("azure:network/networkInterface:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
	// If the Virtual Machine using this Network Interface is part of an Availability Set, then this list will have the union of all DNS servers from all Network Interfaces that are part of the Availability Set.
	AppliedDnsServers []string `pulumi:"appliedDnsServers"`
	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	DnsServers []string `pulumi:"dnsServers"`
	// Should Accelerated Networking be enabled? Defaults to `false`.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Should IP Forwarding be enabled? Defaults to `false`.
	EnableIpForwarding *bool `pulumi:"enableIpForwarding"`
	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	InternalDnsNameLabel *string `pulumi:"internalDnsNameLabel"`
	// One or more `ipConfiguration` blocks as defined below.
	IpConfigurations []NetworkInterfaceIpConfiguration `pulumi:"ipConfigurations"`
	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Media Access Control (MAC) Address of the Network Interface.
	MacAddress *string `pulumi:"macAddress"`
	// The name of the Network Interface. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Static IP Address which should be used.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The private IP addresses of the network interface.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Machine which this Network Interface is connected to.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type NetworkInterfaceState struct {
	// If the Virtual Machine using this Network Interface is part of an Availability Set, then this list will have the union of all DNS servers from all Network Interfaces that are part of the Availability Set.
	AppliedDnsServers pulumi.StringArrayInput
	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	DnsServers pulumi.StringArrayInput
	// Should Accelerated Networking be enabled? Defaults to `false`.
	EnableAcceleratedNetworking pulumi.BoolPtrInput
	// Should IP Forwarding be enabled? Defaults to `false`.
	EnableIpForwarding pulumi.BoolPtrInput
	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	InternalDnsNameLabel pulumi.StringPtrInput
	// One or more `ipConfiguration` blocks as defined below.
	IpConfigurations NetworkInterfaceIpConfigurationArrayInput
	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Media Access Control (MAC) Address of the Network Interface.
	MacAddress pulumi.StringPtrInput
	// The name of the Network Interface. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Static IP Address which should be used.
	PrivateIpAddress pulumi.StringPtrInput
	// The private IP addresses of the network interface.
	PrivateIpAddresses pulumi.StringArrayInput
	// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Machine which this Network Interface is connected to.
	VirtualMachineId pulumi.StringPtrInput
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	DnsServers []string `pulumi:"dnsServers"`
	// Should Accelerated Networking be enabled? Defaults to `false`.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Should IP Forwarding be enabled? Defaults to `false`.
	EnableIpForwarding *bool `pulumi:"enableIpForwarding"`
	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	InternalDnsNameLabel *string `pulumi:"internalDnsNameLabel"`
	// One or more `ipConfiguration` blocks as defined below.
	IpConfigurations []NetworkInterfaceIpConfiguration `pulumi:"ipConfigurations"`
	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Network Interface. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	DnsServers pulumi.StringArrayInput
	// Should Accelerated Networking be enabled? Defaults to `false`.
	EnableAcceleratedNetworking pulumi.BoolPtrInput
	// Should IP Forwarding be enabled? Defaults to `false`.
	EnableIpForwarding pulumi.BoolPtrInput
	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	InternalDnsNameLabel pulumi.StringPtrInput
	// One or more `ipConfiguration` blocks as defined below.
	IpConfigurations NetworkInterfaceIpConfigurationArrayInput
	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Network Interface. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}
