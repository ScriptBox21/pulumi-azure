// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Network Interface located in a Virtual Network, usually attached to a Virtual Machine.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West US",
 *     name: "acceptanceTestResourceGroup1",
 * });
 * const testVirtualNetwork = new azure.network.VirtualNetwork("test", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: testResourceGroup.location,
 *     name: "acceptanceTestVirtualNetwork1",
 *     resourceGroupName: testResourceGroup.name,
 * });
 * const testSubnet = new azure.network.Subnet("test", {
 *     addressPrefix: "10.0.2.0/24",
 *     name: "testsubnet",
 *     resourceGroupName: testResourceGroup.name,
 *     virtualNetworkName: testVirtualNetwork.name,
 * });
 * const testNetworkInterface = new azure.network.NetworkInterface("test", {
 *     ipConfigurations: [{
 *         name: "testconfiguration1",
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: testSubnet.id,
 *     }],
 *     location: testResourceGroup.location,
 *     name: "acceptanceTestNetworkInterface1",
 *     resourceGroupName: testResourceGroup.name,
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * ```
 */
export class NetworkInterface extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceState, opts?: pulumi.CustomResourceOptions): NetworkInterface {
        return new NetworkInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkInterface:NetworkInterface';

    /**
     * Returns true if the given object is an instance of NetworkInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterface.__pulumiType;
    }

    /**
     * If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set
     */
    public readonly appliedDnsServers!: pulumi.Output<string[]>;
    /**
     * List of DNS servers IP addresses to use for this NIC, overrides the VNet-level server list
     */
    public readonly dnsServers!: pulumi.Output<string[]>;
    /**
     * Enables Azure Accelerated Networking using SR-IOV. Only certain VM instance sizes are supported. Refer to [Create a Virtual Machine with Accelerated Networking](https://docs.microsoft.com/en-us/azure/virtual-network/create-vm-accelerated-networking-cli). Defaults to `false`.
     */
    public readonly enableAcceleratedNetworking!: pulumi.Output<boolean | undefined>;
    /**
     * Enables IP Forwarding on the NIC. Defaults to `false`.
     */
    public readonly enableIpForwarding!: pulumi.Output<boolean | undefined>;
    /**
     * Relative DNS name for this NIC used for internal communications between VMs in the same VNet
     */
    public readonly internalDnsNameLabel!: pulumi.Output<string>;
    public readonly internalFqdn!: pulumi.Output<string>;
    /**
     * One or more `ip_configuration` associated with this NIC as documented below.
     */
    public readonly ipConfigurations!: pulumi.Output<{ applicationGatewayBackendAddressPoolsIds: string[], applicationSecurityGroupIds: string[], loadBalancerBackendAddressPoolsIds: string[], loadBalancerInboundNatRulesIds: string[], name: string, primary: boolean, privateIpAddress?: string, privateIpAddressAllocation: string, privateIpAddressVersion?: string, publicIpAddressId?: string, subnetId?: string }[]>;
    /**
     * The location/region where the network interface is created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The media access control (MAC) address of the network interface.
     */
    public readonly macAddress!: pulumi.Output<string>;
    /**
     * The name of the network interface. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Network Security Group to associate with the network interface.
     */
    public readonly networkSecurityGroupId!: pulumi.Output<string | undefined>;
    /**
     * The first private IP address of the network interface.
     */
    public /*out*/ readonly privateIpAddress!: pulumi.Output<string>;
    /**
     * The private IP addresses of the network interface.
     */
    public /*out*/ readonly privateIpAddresses!: pulumi.Output<string[]>;
    /**
     * The name of the resource group in which to create the network interface. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;
    /**
     * Reference to a VM with which this NIC has been associated.
     */
    public readonly virtualMachineId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceArgs | NetworkInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkInterfaceState | undefined;
            inputs["appliedDnsServers"] = state ? state.appliedDnsServers : undefined;
            inputs["dnsServers"] = state ? state.dnsServers : undefined;
            inputs["enableAcceleratedNetworking"] = state ? state.enableAcceleratedNetworking : undefined;
            inputs["enableIpForwarding"] = state ? state.enableIpForwarding : undefined;
            inputs["internalDnsNameLabel"] = state ? state.internalDnsNameLabel : undefined;
            inputs["internalFqdn"] = state ? state.internalFqdn : undefined;
            inputs["ipConfigurations"] = state ? state.ipConfigurations : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["macAddress"] = state ? state.macAddress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkSecurityGroupId"] = state ? state.networkSecurityGroupId : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            inputs["privateIpAddresses"] = state ? state.privateIpAddresses : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceArgs | undefined;
            if (!args || args.ipConfigurations === undefined) {
                throw new Error("Missing required property 'ipConfigurations'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["appliedDnsServers"] = args ? args.appliedDnsServers : undefined;
            inputs["dnsServers"] = args ? args.dnsServers : undefined;
            inputs["enableAcceleratedNetworking"] = args ? args.enableAcceleratedNetworking : undefined;
            inputs["enableIpForwarding"] = args ? args.enableIpForwarding : undefined;
            inputs["internalDnsNameLabel"] = args ? args.internalDnsNameLabel : undefined;
            inputs["internalFqdn"] = args ? args.internalFqdn : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["macAddress"] = args ? args.macAddress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkSecurityGroupId"] = args ? args.networkSecurityGroupId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
            inputs["privateIpAddress"] = undefined /*out*/;
            inputs["privateIpAddresses"] = undefined /*out*/;
        }
        super(NetworkInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterface resources.
 */
export interface NetworkInterfaceState {
    /**
     * If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set
     */
    readonly appliedDnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of DNS servers IP addresses to use for this NIC, overrides the VNet-level server list
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables Azure Accelerated Networking using SR-IOV. Only certain VM instance sizes are supported. Refer to [Create a Virtual Machine with Accelerated Networking](https://docs.microsoft.com/en-us/azure/virtual-network/create-vm-accelerated-networking-cli). Defaults to `false`.
     */
    readonly enableAcceleratedNetworking?: pulumi.Input<boolean>;
    /**
     * Enables IP Forwarding on the NIC. Defaults to `false`.
     */
    readonly enableIpForwarding?: pulumi.Input<boolean>;
    /**
     * Relative DNS name for this NIC used for internal communications between VMs in the same VNet
     */
    readonly internalDnsNameLabel?: pulumi.Input<string>;
    readonly internalFqdn?: pulumi.Input<string>;
    /**
     * One or more `ip_configuration` associated with this NIC as documented below.
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<{ applicationGatewayBackendAddressPoolsIds?: pulumi.Input<pulumi.Input<string>[]>, applicationSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>, loadBalancerBackendAddressPoolsIds?: pulumi.Input<pulumi.Input<string>[]>, loadBalancerInboundNatRulesIds?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, primary?: pulumi.Input<boolean>, privateIpAddress?: pulumi.Input<string>, privateIpAddressAllocation: pulumi.Input<string>, privateIpAddressVersion?: pulumi.Input<string>, publicIpAddressId?: pulumi.Input<string>, subnetId?: pulumi.Input<string> }>[]>;
    /**
     * The location/region where the network interface is created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The media access control (MAC) address of the network interface.
     */
    readonly macAddress?: pulumi.Input<string>;
    /**
     * The name of the network interface. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the Network Security Group to associate with the network interface.
     */
    readonly networkSecurityGroupId?: pulumi.Input<string>;
    /**
     * The first private IP address of the network interface.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
    /**
     * The private IP addresses of the network interface.
     */
    readonly privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group in which to create the network interface. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Reference to a VM with which this NIC has been associated.
     */
    readonly virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterface resource.
 */
export interface NetworkInterfaceArgs {
    /**
     * If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set
     */
    readonly appliedDnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of DNS servers IP addresses to use for this NIC, overrides the VNet-level server list
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables Azure Accelerated Networking using SR-IOV. Only certain VM instance sizes are supported. Refer to [Create a Virtual Machine with Accelerated Networking](https://docs.microsoft.com/en-us/azure/virtual-network/create-vm-accelerated-networking-cli). Defaults to `false`.
     */
    readonly enableAcceleratedNetworking?: pulumi.Input<boolean>;
    /**
     * Enables IP Forwarding on the NIC. Defaults to `false`.
     */
    readonly enableIpForwarding?: pulumi.Input<boolean>;
    /**
     * Relative DNS name for this NIC used for internal communications between VMs in the same VNet
     */
    readonly internalDnsNameLabel?: pulumi.Input<string>;
    readonly internalFqdn?: pulumi.Input<string>;
    /**
     * One or more `ip_configuration` associated with this NIC as documented below.
     */
    readonly ipConfigurations: pulumi.Input<pulumi.Input<{ applicationGatewayBackendAddressPoolsIds?: pulumi.Input<pulumi.Input<string>[]>, applicationSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>, loadBalancerBackendAddressPoolsIds?: pulumi.Input<pulumi.Input<string>[]>, loadBalancerInboundNatRulesIds?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, primary?: pulumi.Input<boolean>, privateIpAddress?: pulumi.Input<string>, privateIpAddressAllocation: pulumi.Input<string>, privateIpAddressVersion?: pulumi.Input<string>, publicIpAddressId?: pulumi.Input<string>, subnetId?: pulumi.Input<string> }>[]>;
    /**
     * The location/region where the network interface is created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The media access control (MAC) address of the network interface.
     */
    readonly macAddress?: pulumi.Input<string>;
    /**
     * The name of the network interface. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the Network Security Group to associate with the network interface.
     */
    readonly networkSecurityGroupId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the network interface. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Reference to a VM with which this NIC has been associated.
     */
    readonly virtualMachineId?: pulumi.Input<string>;
}
