// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Bgp Connection for a Virtual Hub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualHub = new azure.network.VirtualHub("exampleVirtualHub", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: "Standard",
 * });
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.5.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefix: "10.5.1.0/24",
 * });
 * const exampleVirtualHubIp = new azure.network.VirtualHubIp("exampleVirtualHubIp", {
 *     virtualHubId: exampleVirtualHub.id,
 *     privateIpAddress: "10.5.1.18",
 *     privateIpAllocationMethod: "Static",
 *     publicIpAddressId: examplePublicIp.id,
 *     subnetId: exampleSubnet.id,
 * });
 * const exampleBgpConnection = new azure.network.BgpConnection("exampleBgpConnection", {
 *     virtualHubId: exampleVirtualHub.id,
 *     peerAsn: 65514,
 *     peerIp: "169.254.21.5",
 * }, {
 *     dependsOn: [exampleVirtualHubIp],
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Hub Bgp Connections can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/bgpConnection:BgpConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/connection1
 * ```
 */
export class BgpConnection extends pulumi.CustomResource {
    /**
     * Get an existing BgpConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpConnectionState, opts?: pulumi.CustomResourceOptions): BgpConnection {
        return new BgpConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/bgpConnection:BgpConnection';

    /**
     * Returns true if the given object is an instance of BgpConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BgpConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BgpConnection.__pulumiType;
    }

    /**
     * The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    public readonly peerAsn!: pulumi.Output<number>;
    /**
     * The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    public readonly peerIp!: pulumi.Output<string>;
    /**
     * The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
     */
    public readonly virtualHubId!: pulumi.Output<string>;

    /**
     * Create a BgpConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BgpConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BgpConnectionArgs | BgpConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BgpConnectionState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["peerAsn"] = state ? state.peerAsn : undefined;
            inputs["peerIp"] = state ? state.peerIp : undefined;
            inputs["virtualHubId"] = state ? state.virtualHubId : undefined;
        } else {
            const args = argsOrState as BgpConnectionArgs | undefined;
            if ((!args || args.peerAsn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerAsn'");
            }
            if ((!args || args.peerIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerIp'");
            }
            if ((!args || args.virtualHubId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualHubId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["peerAsn"] = args ? args.peerAsn : undefined;
            inputs["peerIp"] = args ? args.peerIp : undefined;
            inputs["virtualHubId"] = args ? args.virtualHubId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BgpConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpConnection resources.
 */
export interface BgpConnectionState {
    /**
     * The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    readonly peerAsn?: pulumi.Input<number>;
    /**
     * The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    readonly peerIp?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
     */
    readonly virtualHubId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BgpConnection resource.
 */
export interface BgpConnectionArgs {
    /**
     * The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    readonly peerAsn: pulumi.Input<number>;
    /**
     * The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
     */
    readonly peerIp: pulumi.Input<string>;
    /**
     * The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
     */
    readonly virtualHubId: pulumi.Input<string>;
}
