// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a HDInsight Spark Cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
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
 * const exampleSparkCluster = new azure.hdinsight.SparkCluster("exampleSparkCluster", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     clusterVersion: "3.6",
 *     tier: "Standard",
 *     componentVersion: {
 *         spark: "2.3",
 *     },
 *     gateway: {
 *         enabled: true,
 *         username: "acctestusrgw",
 *         password: "Password123!",
 *     },
 *     storageAccounts: [{
 *         storageContainerId: exampleContainer.id,
 *         storageAccountKey: exampleAccount.primaryAccessKey,
 *         isDefault: true,
 *     }],
 *     roles: {
 *         headNode: {
 *             vmSize: "Standard_A3",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *         },
 *         workerNode: {
 *             vmSize: "Standard_A3",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *             targetInstanceCount: 3,
 *         },
 *         zookeeperNode: {
 *             vmSize: "Medium",
 *             username: "acctestusrvm",
 *             password: "AccTestvdSC4daf986!",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * HDInsight Spark Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:hdinsight/sparkCluster:SparkCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1}
 * ```
 */
export class SparkCluster extends pulumi.CustomResource {
    /**
     * Get an existing SparkCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SparkClusterState, opts?: pulumi.CustomResourceOptions): SparkCluster {
        return new SparkCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:hdinsight/sparkCluster:SparkCluster';

    /**
     * Returns true if the given object is an instance of SparkCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SparkCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SparkCluster.__pulumiType;
    }

    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    public readonly clusterVersion!: pulumi.Output<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    public readonly componentVersion!: pulumi.Output<outputs.hdinsight.SparkClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    public readonly gateway!: pulumi.Output<outputs.hdinsight.SparkClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Spark Cluster.
     */
    public /*out*/ readonly httpsEndpoint!: pulumi.Output<string>;
    /**
     * Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A `metastores` block as defined below.
     */
    public readonly metastores!: pulumi.Output<outputs.hdinsight.SparkClusterMetastores | undefined>;
    /**
     * A `monitor` block as defined below.
     */
    public readonly monitor!: pulumi.Output<outputs.hdinsight.SparkClusterMonitor | undefined>;
    /**
     * Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `roles` block as defined below.
     */
    public readonly roles!: pulumi.Output<outputs.hdinsight.SparkClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Spark Cluster.
     */
    public /*out*/ readonly sshEndpoint!: pulumi.Output<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    public readonly storageAccountGen2!: pulumi.Output<outputs.hdinsight.SparkClusterStorageAccountGen2 | undefined>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    public readonly storageAccounts!: pulumi.Output<outputs.hdinsight.SparkClusterStorageAccount[] | undefined>;
    /**
     * A map of Tags which should be assigned to this HDInsight Spark Cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    public readonly tier!: pulumi.Output<string>;
    public readonly tlsMinVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a SparkCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SparkClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SparkClusterArgs | SparkClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SparkClusterState | undefined;
            inputs["clusterVersion"] = state ? state.clusterVersion : undefined;
            inputs["componentVersion"] = state ? state.componentVersion : undefined;
            inputs["gateway"] = state ? state.gateway : undefined;
            inputs["httpsEndpoint"] = state ? state.httpsEndpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["metastores"] = state ? state.metastores : undefined;
            inputs["monitor"] = state ? state.monitor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["sshEndpoint"] = state ? state.sshEndpoint : undefined;
            inputs["storageAccountGen2"] = state ? state.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = state ? state.storageAccounts : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
        } else {
            const args = argsOrState as SparkClusterArgs | undefined;
            if (!args || args.clusterVersion === undefined) {
                throw new Error("Missing required property 'clusterVersion'");
            }
            if (!args || args.componentVersion === undefined) {
                throw new Error("Missing required property 'componentVersion'");
            }
            if (!args || args.gateway === undefined) {
                throw new Error("Missing required property 'gateway'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.roles === undefined) {
                throw new Error("Missing required property 'roles'");
            }
            if (!args || args.tier === undefined) {
                throw new Error("Missing required property 'tier'");
            }
            inputs["clusterVersion"] = args ? args.clusterVersion : undefined;
            inputs["componentVersion"] = args ? args.componentVersion : undefined;
            inputs["gateway"] = args ? args.gateway : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metastores"] = args ? args.metastores : undefined;
            inputs["monitor"] = args ? args.monitor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["storageAccountGen2"] = args ? args.storageAccountGen2 : undefined;
            inputs["storageAccounts"] = args ? args.storageAccounts : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            inputs["httpsEndpoint"] = undefined /*out*/;
            inputs["sshEndpoint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SparkCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SparkCluster resources.
 */
export interface SparkClusterState {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    readonly clusterVersion?: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    readonly componentVersion?: pulumi.Input<inputs.hdinsight.SparkClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    readonly gateway?: pulumi.Input<inputs.hdinsight.SparkClusterGateway>;
    /**
     * The HTTPS Connectivity Endpoint for this HDInsight Spark Cluster.
     */
    readonly httpsEndpoint?: pulumi.Input<string>;
    /**
     * Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    readonly metastores?: pulumi.Input<inputs.hdinsight.SparkClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    readonly monitor?: pulumi.Input<inputs.hdinsight.SparkClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `roles` block as defined below.
     */
    readonly roles?: pulumi.Input<inputs.hdinsight.SparkClusterRoles>;
    /**
     * The SSH Connectivity Endpoint for this HDInsight Spark Cluster.
     */
    readonly sshEndpoint?: pulumi.Input<string>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    readonly storageAccountGen2?: pulumi.Input<inputs.hdinsight.SparkClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.SparkClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Spark Cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    readonly tier?: pulumi.Input<string>;
    readonly tlsMinVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SparkCluster resource.
 */
export interface SparkClusterArgs {
    /**
     * Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
     */
    readonly clusterVersion: pulumi.Input<string>;
    /**
     * A `componentVersion` block as defined below.
     */
    readonly componentVersion: pulumi.Input<inputs.hdinsight.SparkClusterComponentVersion>;
    /**
     * A `gateway` block as defined below.
     */
    readonly gateway: pulumi.Input<inputs.hdinsight.SparkClusterGateway>;
    /**
     * Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `metastores` block as defined below.
     */
    readonly metastores?: pulumi.Input<inputs.hdinsight.SparkClusterMetastores>;
    /**
     * A `monitor` block as defined below.
     */
    readonly monitor?: pulumi.Input<inputs.hdinsight.SparkClusterMonitor>;
    /**
     * Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `roles` block as defined below.
     */
    readonly roles: pulumi.Input<inputs.hdinsight.SparkClusterRoles>;
    /**
     * A `storageAccountGen2` block as defined below.
     */
    readonly storageAccountGen2?: pulumi.Input<inputs.hdinsight.SparkClusterStorageAccountGen2>;
    /**
     * One or more `storageAccount` block as defined below.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.hdinsight.SparkClusterStorageAccount>[]>;
    /**
     * A map of Tags which should be assigned to this HDInsight Spark Cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
     */
    readonly tier: pulumi.Input<string>;
    readonly tlsMinVersion?: pulumi.Input<string>;
}
