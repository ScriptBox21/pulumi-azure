// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Enables you to manage DNS AAAA Records within Azure DNS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleZone = new azure.dns.Zone("exampleZone", {resourceGroupName: exampleResourceGroup.name});
 * const exampleAaaaRecord = new azure.dns.AaaaRecord("exampleAaaaRecord", {
 *     zoneName: exampleZone.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ttl: 300,
 * });
 * ```
 * ### Alias Record)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleZone = new azure.dns.Zone("exampleZone", {resourceGroupName: exampleResourceGroup.name});
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Dynamic",
 *     ipVersion: "IPv6",
 * });
 * const exampleAaaaRecord = new azure.dns.AaaaRecord("exampleAaaaRecord", {
 *     zoneName: exampleZone.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ttl: 300,
 *     targetResourceId: examplePublicIp.id,
 * });
 * ```
 *
 * ## Import
 *
 * AAAA records can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:dns/aaaaRecord:AaaaRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/AAAA/myrecord1
 * ```
 */
export class AaaaRecord extends pulumi.CustomResource {
    /**
     * Get an existing AaaaRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AaaaRecordState, opts?: pulumi.CustomResourceOptions): AaaaRecord {
        return new AaaaRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:dns/aaaaRecord:AaaaRecord';

    /**
     * Returns true if the given object is an instance of AaaaRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AaaaRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AaaaRecord.__pulumiType;
    }

    /**
     * The FQDN of the DNS AAAA Record.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The name of the DNS AAAA Record.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of IPv4 Addresses. Conflicts with `targetResourceId`.
     */
    public readonly records!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Azure resource id of the target object. Conflicts with `records`
     */
    public readonly targetResourceId!: pulumi.Output<string | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly zoneName!: pulumi.Output<string>;

    /**
     * Create a AaaaRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AaaaRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AaaaRecordArgs | AaaaRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AaaaRecordState | undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["records"] = state ? state.records : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetResourceId"] = state ? state.targetResourceId : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["zoneName"] = state ? state.zoneName : undefined;
        } else {
            const args = argsOrState as AaaaRecordArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            if ((!args || args.zoneName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["records"] = args ? args.records : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetResourceId"] = args ? args.targetResourceId : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["zoneName"] = args ? args.zoneName : undefined;
            inputs["fqdn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AaaaRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AaaaRecord resources.
 */
export interface AaaaRecordState {
    /**
     * The FQDN of the DNS AAAA Record.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * The name of the DNS AAAA Record.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of IPv4 Addresses. Conflicts with `targetResourceId`.
     */
    readonly records?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Azure resource id of the target object. Conflicts with `records`
     */
    readonly targetResourceId?: pulumi.Input<string>;
    readonly ttl?: pulumi.Input<number>;
    /**
     * Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AaaaRecord resource.
 */
export interface AaaaRecordArgs {
    /**
     * The name of the DNS AAAA Record.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of IPv4 Addresses. Conflicts with `targetResourceId`.
     */
    readonly records?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Azure resource id of the target object. Conflicts with `records`
     */
    readonly targetResourceId?: pulumi.Input<string>;
    readonly ttl: pulumi.Input<number>;
    /**
     * Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName: pulumi.Input<string>;
}
