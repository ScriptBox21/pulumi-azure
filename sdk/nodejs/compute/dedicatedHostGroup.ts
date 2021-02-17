// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage a Dedicated Host Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleDedicatedHostGroup = new azure.compute.DedicatedHostGroup("exampleDedicatedHostGroup", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     platformFaultDomainCount: 1,
 * });
 * ```
 *
 * ## Import
 *
 * Dedicated Host Group can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/dedicatedHostGroup:DedicatedHostGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/hostGroups/group1
 * ```
 */
export class DedicatedHostGroup extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedHostGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedHostGroupState, opts?: pulumi.CustomResourceOptions): DedicatedHostGroup {
        return new DedicatedHostGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/dedicatedHostGroup:DedicatedHostGroup';

    /**
     * Returns true if the given object is an instance of DedicatedHostGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedHostGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedHostGroup.__pulumiType;
    }

    /**
     * The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
     */
    public readonly platformFaultDomainCount!: pulumi.Output<number>;
    /**
     * Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
     */
    public readonly zones!: pulumi.Output<string | undefined>;

    /**
     * Create a DedicatedHostGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedHostGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedHostGroupArgs | DedicatedHostGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DedicatedHostGroupState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platformFaultDomainCount"] = state ? state.platformFaultDomainCount : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as DedicatedHostGroupArgs | undefined;
            if ((!args || args.platformFaultDomainCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'platformFaultDomainCount'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platformFaultDomainCount"] = args ? args.platformFaultDomainCount : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DedicatedHostGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedHostGroup resources.
 */
export interface DedicatedHostGroupState {
    /**
     * The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
     */
    readonly platformFaultDomainCount?: pulumi.Input<number>;
    /**
     * Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
     */
    readonly zones?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DedicatedHostGroup resource.
 */
export interface DedicatedHostGroupArgs {
    /**
     * The Azure location where the Dedicated Host Group exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Dedicated Host Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of fault domains that the Dedicated Host Group spans. Changing this forces a new resource to be created.
     */
    readonly platformFaultDomainCount: pulumi.Input<number>;
    /**
     * Specifies the name of the resource group the Dedicated Host Group is located in. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Availability Zones in which the Dedicated Host Group should be located. Changing this forces a new resource to be created.
     */
    readonly zones?: pulumi.Input<string>;
}
