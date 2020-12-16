// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Kusto (also known as Azure Data Explorer) Cluster
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "East US"});
 * const example = new azure.kusto.Cluster("example", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     sku: {
 *         name: "Standard_D13_v2",
 *         capacity: 2,
 *     },
 *     tags: {
 *         Environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Kusto Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:kusto/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:kusto/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The Kusto Cluster URI to be used for data ingestion.
     */
    public /*out*/ readonly dataIngestionUri!: pulumi.Output<string>;
    /**
     * Specifies if the cluster's disks are encrypted.
     */
    public readonly enableDiskEncryption!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the purge operations are enabled.
     */
    public readonly enablePurge!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the streaming ingest is enabled.
     */
    public readonly enableStreamingIngest!: pulumi.Output<boolean | undefined>;
    /**
     * . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * A identity block.
     */
    public readonly identity!: pulumi.Output<outputs.kusto.ClusterIdentity>;
    /**
     * An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
     */
    public readonly languageExtensions!: pulumi.Output<string[] | undefined>;
    /**
     * The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An `optimizedAutoScale` block as defined below.
     */
    public readonly optimizedAutoScale!: pulumi.Output<outputs.kusto.ClusterOptimizedAutoScale | undefined>;
    /**
     * Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `sku` block as defined below.
     */
    public readonly sku!: pulumi.Output<outputs.kusto.ClusterSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies a list of tenant IDs that are trusted by the cluster.
     */
    public readonly trustedExternalTenants!: pulumi.Output<string[]>;
    /**
     * The FQDN of the Azure Kusto Cluster.
     */
    public /*out*/ readonly uri!: pulumi.Output<string>;
    /**
     * A `virtualNetworkConfiguration` block as defined below.
     */
    public readonly virtualNetworkConfiguration!: pulumi.Output<outputs.kusto.ClusterVirtualNetworkConfiguration | undefined>;
    /**
     * A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["dataIngestionUri"] = state ? state.dataIngestionUri : undefined;
            inputs["enableDiskEncryption"] = state ? state.enableDiskEncryption : undefined;
            inputs["enablePurge"] = state ? state.enablePurge : undefined;
            inputs["enableStreamingIngest"] = state ? state.enableStreamingIngest : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["languageExtensions"] = state ? state.languageExtensions : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["optimizedAutoScale"] = state ? state.optimizedAutoScale : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["trustedExternalTenants"] = state ? state.trustedExternalTenants : undefined;
            inputs["uri"] = state ? state.uri : undefined;
            inputs["virtualNetworkConfiguration"] = state ? state.virtualNetworkConfiguration : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["enableDiskEncryption"] = args ? args.enableDiskEncryption : undefined;
            inputs["enablePurge"] = args ? args.enablePurge : undefined;
            inputs["enableStreamingIngest"] = args ? args.enableStreamingIngest : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["languageExtensions"] = args ? args.languageExtensions : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["optimizedAutoScale"] = args ? args.optimizedAutoScale : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["trustedExternalTenants"] = args ? args.trustedExternalTenants : undefined;
            inputs["virtualNetworkConfiguration"] = args ? args.virtualNetworkConfiguration : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["dataIngestionUri"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The Kusto Cluster URI to be used for data ingestion.
     */
    readonly dataIngestionUri?: pulumi.Input<string>;
    /**
     * Specifies if the cluster's disks are encrypted.
     */
    readonly enableDiskEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies if the purge operations are enabled.
     */
    readonly enablePurge?: pulumi.Input<boolean>;
    /**
     * Specifies if the streaming ingest is enabled.
     */
    readonly enableStreamingIngest?: pulumi.Input<boolean>;
    /**
     * . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * A identity block.
     */
    readonly identity?: pulumi.Input<inputs.kusto.ClusterIdentity>;
    /**
     * An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
     */
    readonly languageExtensions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An `optimizedAutoScale` block as defined below.
     */
    readonly optimizedAutoScale?: pulumi.Input<inputs.kusto.ClusterOptimizedAutoScale>;
    /**
     * Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku?: pulumi.Input<inputs.kusto.ClusterSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a list of tenant IDs that are trusted by the cluster.
     */
    readonly trustedExternalTenants?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The FQDN of the Azure Kusto Cluster.
     */
    readonly uri?: pulumi.Input<string>;
    /**
     * A `virtualNetworkConfiguration` block as defined below.
     */
    readonly virtualNetworkConfiguration?: pulumi.Input<inputs.kusto.ClusterVirtualNetworkConfiguration>;
    /**
     * A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Specifies if the cluster's disks are encrypted.
     */
    readonly enableDiskEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies if the purge operations are enabled.
     */
    readonly enablePurge?: pulumi.Input<boolean>;
    /**
     * Specifies if the streaming ingest is enabled.
     */
    readonly enableStreamingIngest?: pulumi.Input<boolean>;
    /**
     * . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * A identity block.
     */
    readonly identity?: pulumi.Input<inputs.kusto.ClusterIdentity>;
    /**
     * An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
     */
    readonly languageExtensions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An `optimizedAutoScale` block as defined below.
     */
    readonly optimizedAutoScale?: pulumi.Input<inputs.kusto.ClusterOptimizedAutoScale>;
    /**
     * Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku: pulumi.Input<inputs.kusto.ClusterSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a list of tenant IDs that are trusted by the cluster.
     */
    readonly trustedExternalTenants?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `virtualNetworkConfiguration` block as defined below.
     */
    readonly virtualNetworkConfiguration?: pulumi.Input<inputs.kusto.ClusterVirtualNetworkConfiguration>;
    /**
     * A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
