// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an ExpressRoute circuit.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West US",
 *     name: "exprtTest",
 * });
 * const testExpressRouteCircuit = new azure.network.ExpressRouteCircuit("test", {
 *     bandwidthInMbps: 50,
 *     location: testResourceGroup.location,
 *     name: "expressRoute1",
 *     peeringLocation: "Silicon Valley",
 *     resourceGroupName: testResourceGroup.name,
 *     serviceProviderName: "Equinix",
 *     sku: {
 *         family: "MeteredData",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 */
export class ExpressRouteCircuit extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRouteCircuit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExpressRouteCircuitState, opts?: pulumi.CustomResourceOptions): ExpressRouteCircuit {
        return new ExpressRouteCircuit(name, <any>state, { ...opts, id: id });
    }

    /**
     * Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
     */
    public readonly allowClassicOperations: pulumi.Output<boolean | undefined>;
    /**
     * The bandwidth in Mbps of the circuit being created.
     */
    public readonly bandwidthInMbps: pulumi.Output<number>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of the peering location and **not** the Azure resource location.
     */
    public readonly peeringLocation: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * The string needed by the service provider to provision the ExpressRoute circuit.
     */
    public /*out*/ readonly serviceKey: pulumi.Output<string>;
    /**
     * The name of the ExpressRoute Service Provider.
     */
    public readonly serviceProviderName: pulumi.Output<string>;
    /**
     * The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
     */
    public /*out*/ readonly serviceProviderProvisioningState: pulumi.Output<string>;
    /**
     * A `sku` block for the ExpressRoute circuit as documented below.
     */
    public readonly sku: pulumi.Output<{ family: string, tier: string }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a ExpressRouteCircuit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRouteCircuitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExpressRouteCircuitArgs | ExpressRouteCircuitState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ExpressRouteCircuitState = argsOrState as ExpressRouteCircuitState | undefined;
            inputs["allowClassicOperations"] = state ? state.allowClassicOperations : undefined;
            inputs["bandwidthInMbps"] = state ? state.bandwidthInMbps : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["peeringLocation"] = state ? state.peeringLocation : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serviceKey"] = state ? state.serviceKey : undefined;
            inputs["serviceProviderName"] = state ? state.serviceProviderName : undefined;
            inputs["serviceProviderProvisioningState"] = state ? state.serviceProviderProvisioningState : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ExpressRouteCircuitArgs | undefined;
            if (!args || args.bandwidthInMbps === undefined) {
                throw new Error("Missing required property 'bandwidthInMbps'");
            }
            if (!args || args.peeringLocation === undefined) {
                throw new Error("Missing required property 'peeringLocation'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceProviderName === undefined) {
                throw new Error("Missing required property 'serviceProviderName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["allowClassicOperations"] = args ? args.allowClassicOperations : undefined;
            inputs["bandwidthInMbps"] = args ? args.bandwidthInMbps : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peeringLocation"] = args ? args.peeringLocation : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceProviderName"] = args ? args.serviceProviderName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["serviceKey"] = undefined /*out*/;
            inputs["serviceProviderProvisioningState"] = undefined /*out*/;
        }
        super("azure:network/expressRouteCircuit:ExpressRouteCircuit", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExpressRouteCircuit resources.
 */
export interface ExpressRouteCircuitState {
    /**
     * Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
     */
    readonly allowClassicOperations?: pulumi.Input<boolean>;
    /**
     * The bandwidth in Mbps of the circuit being created.
     */
    readonly bandwidthInMbps?: pulumi.Input<number>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the peering location and **not** the Azure resource location.
     */
    readonly peeringLocation?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The string needed by the service provider to provision the ExpressRoute circuit.
     */
    readonly serviceKey?: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute Service Provider.
     */
    readonly serviceProviderName?: pulumi.Input<string>;
    /**
     * The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
     */
    readonly serviceProviderProvisioningState?: pulumi.Input<string>;
    /**
     * A `sku` block for the ExpressRoute circuit as documented below.
     */
    readonly sku?: pulumi.Input<{ family: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ExpressRouteCircuit resource.
 */
export interface ExpressRouteCircuitArgs {
    /**
     * Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
     */
    readonly allowClassicOperations?: pulumi.Input<boolean>;
    /**
     * The bandwidth in Mbps of the circuit being created.
     */
    readonly bandwidthInMbps: pulumi.Input<number>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the peering location and **not** the Azure resource location.
     */
    readonly peeringLocation: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute Service Provider.
     */
    readonly serviceProviderName: pulumi.Input<string>;
    /**
     * A `sku` block for the ExpressRoute circuit as documented below.
     */
    readonly sku: pulumi.Input<{ family: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
