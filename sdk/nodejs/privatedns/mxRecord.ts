// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Enables you to manage DNS MX Records within Azure Private DNS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleZone = new azure.privatedns.Zone("exampleZone", {resourceGroupName: exampleResourceGroup.name});
 * const exampleMxRecord = new azure.privatedns.MxRecord("exampleMxRecord", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     zoneName: exampleZone.name,
 *     ttl: 300,
 *     records: [
 *         {
 *             preference: 10,
 *             exchange: "mx1.contoso.com",
 *         },
 *         {
 *             preference: 20,
 *             exchange: "backupmx.contoso.com",
 *         },
 *     ],
 *     tags: {
 *         Environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Private DNS MX Records can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:privatedns/mxRecord:MxRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/contoso.com/MX/@
 * ```
 */
export class MxRecord extends pulumi.CustomResource {
    /**
     * Get an existing MxRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MxRecordState, opts?: pulumi.CustomResourceOptions): MxRecord {
        return new MxRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:privatedns/mxRecord:MxRecord';

    /**
     * Returns true if the given object is an instance of MxRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MxRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MxRecord.__pulumiType;
    }

    /**
     * The FQDN of the DNS MX Record.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    public readonly records!: pulumi.Output<outputs.privatedns.MxRecordRecord[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly zoneName!: pulumi.Output<string>;

    /**
     * Create a MxRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MxRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MxRecordArgs | MxRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MxRecordState | undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["records"] = state ? state.records : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["zoneName"] = state ? state.zoneName : undefined;
        } else {
            const args = argsOrState as MxRecordArgs | undefined;
            if ((!args || args.records === undefined) && !opts.urn) {
                throw new Error("Missing required property 'records'");
            }
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
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["zoneName"] = args ? args.zoneName : undefined;
            inputs["fqdn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MxRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MxRecord resources.
 */
export interface MxRecordState {
    /**
     * The FQDN of the DNS MX Record.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    readonly records?: pulumi.Input<pulumi.Input<inputs.privatedns.MxRecordRecord>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly ttl?: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MxRecord resource.
 */
export interface MxRecordArgs {
    /**
     * The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    readonly records: pulumi.Input<pulumi.Input<inputs.privatedns.MxRecordRecord>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly ttl: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName: pulumi.Input<string>;
}
