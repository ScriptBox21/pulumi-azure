// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IotHub Enrichment
 *
 * > **NOTE:** Enrichment can be defined either directly on the `azure.iot.IoTHub` resource, or using the `azure.iot.Enrichment` resources - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleContainer = new azure.storage.Container("exampleContainer", {
 *     storageAccountName: exampleAccount.name,
 *     containerAccessType: "private",
 * });
 * const exampleIoTHub = new azure.iot.IoTHub("exampleIoTHub", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: {
 *         name: "S1",
 *         capacity: "1",
 *     },
 *     tags: {
 *         purpose: "testing",
 *     },
 * });
 * const exampleEndpointStorageContainer = new azure.iot.EndpointStorageContainer("exampleEndpointStorageContainer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     iothubName: exampleIoTHub.name,
 *     connectionString: exampleAccount.primaryBlobConnectionString,
 *     batchFrequencyInSeconds: 60,
 *     maxChunkSizeInBytes: 10485760,
 *     containerName: exampleContainer.name,
 *     encoding: "Avro",
 *     fileNameFormat: "{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}",
 * });
 * const exampleRoute = new azure.iot.Route("exampleRoute", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     iothubName: exampleIoTHub.name,
 *     source: "DeviceMessages",
 *     condition: "true",
 *     endpointNames: [exampleEndpointStorageContainer.name],
 *     enabled: true,
 * });
 * const exampleEnrichment = new azure.iot.Enrichment("exampleEnrichment", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     iothubName: exampleIoTHub.name,
 *     key: "example",
 *     value: "my value",
 *     endpointNames: [exampleEndpointStorageContainer.name],
 * });
 * ```
 *
 * ## Import
 *
 * IoTHub Enrichment can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:iot/enrichment:Enrichment enrichment1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/Enrichments/enrichment1
 * ```
 */
export class Enrichment extends pulumi.CustomResource {
    /**
     * Get an existing Enrichment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnrichmentState, opts?: pulumi.CustomResourceOptions): Enrichment {
        return new Enrichment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:iot/enrichment:Enrichment';

    /**
     * Returns true if the given object is an instance of Enrichment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Enrichment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Enrichment.__pulumiType;
    }

    /**
     * The list of endpoints which will be enriched.
     */
    public readonly endpointNames!: pulumi.Output<string[]>;
    public readonly iothubName!: pulumi.Output<string>;
    /**
     * The key of the enrichment.
     */
    public readonly key!: pulumi.Output<string>;
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Enrichment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnrichmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnrichmentArgs | EnrichmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EnrichmentState | undefined;
            inputs["endpointNames"] = state ? state.endpointNames : undefined;
            inputs["iothubName"] = state ? state.iothubName : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as EnrichmentArgs | undefined;
            if ((!args || args.endpointNames === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'endpointNames'");
            }
            if ((!args || args.iothubName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'iothubName'");
            }
            if ((!args || args.key === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.value === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'value'");
            }
            inputs["endpointNames"] = args ? args.endpointNames : undefined;
            inputs["iothubName"] = args ? args.iothubName : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Enrichment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Enrichment resources.
 */
export interface EnrichmentState {
    /**
     * The list of endpoints which will be enriched.
     */
    readonly endpointNames?: pulumi.Input<pulumi.Input<string>[]>;
    readonly iothubName?: pulumi.Input<string>;
    /**
     * The key of the enrichment.
     */
    readonly key?: pulumi.Input<string>;
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Enrichment resource.
 */
export interface EnrichmentArgs {
    /**
     * The list of endpoints which will be enriched.
     */
    readonly endpointNames: pulumi.Input<pulumi.Input<string>[]>;
    readonly iothubName: pulumi.Input<string>;
    /**
     * The key of the enrichment.
     */
    readonly key: pulumi.Input<string>;
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
     */
    readonly value: pulumi.Input<string>;
}
