// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a VPN Gateway within a Virtual Hub, which enables Site-to-Site communication.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/vpn_gateway.html.markdown.
 */
export class VpnGateway extends pulumi.CustomResource {
    /**
     * Get an existing VpnGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnGatewayState, opts?: pulumi.CustomResourceOptions): VpnGateway {
        return new VpnGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/vpnGateway:VpnGateway';

    /**
     * Returns true if the given object is an instance of VpnGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnGateway.__pulumiType;
    }

    /**
     * A `bgpSettings` block as defined below.
     */
    public readonly bgpSettings!: pulumi.Output<outputs.network.VpnGatewayBgpSetting[]>;
    /**
     * The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Scale Unit for this VPN Gateway. Defaults to `1`.
     */
    public readonly scaleUnit!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the VPN Gateway.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    public readonly virtualHubId!: pulumi.Output<string>;

    /**
     * Create a VpnGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnGatewayArgs | VpnGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpnGatewayState | undefined;
            inputs["bgpSettings"] = state ? state.bgpSettings : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["scaleUnit"] = state ? state.scaleUnit : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualHubId"] = state ? state.virtualHubId : undefined;
        } else {
            const args = argsOrState as VpnGatewayArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualHubId === undefined) {
                throw new Error("Missing required property 'virtualHubId'");
            }
            inputs["bgpSettings"] = args ? args.bgpSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scaleUnit"] = args ? args.scaleUnit : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualHubId"] = args ? args.virtualHubId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpnGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnGateway resources.
 */
export interface VpnGatewayState {
    /**
     * A `bgpSettings` block as defined below.
     */
    readonly bgpSettings?: pulumi.Input<pulumi.Input<inputs.network.VpnGatewayBgpSetting>[]>;
    /**
     * The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Scale Unit for this VPN Gateway. Defaults to `1`.
     */
    readonly scaleUnit?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the VPN Gateway.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    readonly virtualHubId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnGateway resource.
 */
export interface VpnGatewayArgs {
    /**
     * A `bgpSettings` block as defined below.
     */
    readonly bgpSettings?: pulumi.Input<pulumi.Input<inputs.network.VpnGatewayBgpSetting>[]>;
    /**
     * The Azure location where this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Name which should be used for this VPN Gateway. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group in which this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Scale Unit for this VPN Gateway. Defaults to `1`.
     */
    readonly scaleUnit?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the VPN Gateway.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Hub within which this VPN Gateway should be created. Changing this forces a new resource to be created.
     */
    readonly virtualHubId: pulumi.Input<string>;
}
