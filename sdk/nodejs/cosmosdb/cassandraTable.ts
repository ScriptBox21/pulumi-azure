// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Cassandra Table within a Cosmos DB Cassandra Keyspace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = azure.core.getResourceGroup({
 *     name: "tflex-cosmosdb-account-rg",
 * });
 * const exampleAccount = new azure.cosmosdb.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.then(exampleResourceGroup => exampleResourceGroup.name),
 *     location: exampleResourceGroup.then(exampleResourceGroup => exampleResourceGroup.location),
 *     offerType: "Standard",
 *     capabilities: [{
 *         name: "EnableCassandra",
 *     }],
 *     consistencyPolicy: {
 *         consistencyLevel: "Strong",
 *     },
 *     geoLocations: [{
 *         location: "West US",
 *         failoverPriority: 0,
 *     }],
 * });
 * const exampleCassandraKeyspace = new azure.cosmosdb.CassandraKeyspace("exampleCassandraKeyspace", {
 *     resourceGroupName: exampleAccount.resourceGroupName,
 *     accountName: exampleAccount.name,
 *     throughput: 400,
 * });
 * const exampleCassandraTable = new azure.cosmosdb.CassandraTable("exampleCassandraTable", {
 *     cassandraKeyspaceId: exampleCassandraKeyspace.id,
 *     schema: {
 *         columns: [
 *             {
 *                 name: "test1",
 *                 type: "ascii",
 *             },
 *             {
 *                 name: "test2",
 *                 type: "int",
 *             },
 *         ],
 *         partitionKeys: [{
 *             name: "test1",
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Cosmos Cassandra Table can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:cosmosdb/cassandraTable:CassandraTable ks1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1/tables/table1
 * ```
 */
export class CassandraTable extends pulumi.CustomResource {
    /**
     * Get an existing CassandraTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CassandraTableState, opts?: pulumi.CustomResourceOptions): CassandraTable {
        return new CassandraTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:cosmosdb/cassandraTable:CassandraTable';

    /**
     * Returns true if the given object is an instance of CassandraTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CassandraTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CassandraTable.__pulumiType;
    }

    public readonly autoscaleSettings!: pulumi.Output<outputs.cosmosdb.CassandraTableAutoscaleSettings | undefined>;
    /**
     * The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
     */
    public readonly cassandraKeyspaceId!: pulumi.Output<string>;
    public readonly defaultTtl!: pulumi.Output<number>;
    /**
     * Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `schema` block as defined below. Changing this forces a new resource to be created.
     */
    public readonly schema!: pulumi.Output<outputs.cosmosdb.CassandraTableSchema>;
    public readonly throughput!: pulumi.Output<number>;

    /**
     * Create a CassandraTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CassandraTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CassandraTableArgs | CassandraTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CassandraTableState | undefined;
            inputs["autoscaleSettings"] = state ? state.autoscaleSettings : undefined;
            inputs["cassandraKeyspaceId"] = state ? state.cassandraKeyspaceId : undefined;
            inputs["defaultTtl"] = state ? state.defaultTtl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["throughput"] = state ? state.throughput : undefined;
        } else {
            const args = argsOrState as CassandraTableArgs | undefined;
            if ((!args || args.cassandraKeyspaceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'cassandraKeyspaceId'");
            }
            if ((!args || args.schema === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'schema'");
            }
            inputs["autoscaleSettings"] = args ? args.autoscaleSettings : undefined;
            inputs["cassandraKeyspaceId"] = args ? args.cassandraKeyspaceId : undefined;
            inputs["defaultTtl"] = args ? args.defaultTtl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["throughput"] = args ? args.throughput : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CassandraTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CassandraTable resources.
 */
export interface CassandraTableState {
    readonly autoscaleSettings?: pulumi.Input<inputs.cosmosdb.CassandraTableAutoscaleSettings>;
    /**
     * The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
     */
    readonly cassandraKeyspaceId?: pulumi.Input<string>;
    readonly defaultTtl?: pulumi.Input<number>;
    /**
     * Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `schema` block as defined below. Changing this forces a new resource to be created.
     */
    readonly schema?: pulumi.Input<inputs.cosmosdb.CassandraTableSchema>;
    readonly throughput?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CassandraTable resource.
 */
export interface CassandraTableArgs {
    readonly autoscaleSettings?: pulumi.Input<inputs.cosmosdb.CassandraTableAutoscaleSettings>;
    /**
     * The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
     */
    readonly cassandraKeyspaceId: pulumi.Input<string>;
    readonly defaultTtl?: pulumi.Input<number>;
    /**
     * Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `schema` block as defined below. Changing this forces a new resource to be created.
     */
    readonly schema: pulumi.Input<inputs.cosmosdb.CassandraTableSchema>;
    readonly throughput?: pulumi.Input<number>;
}
