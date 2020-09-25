// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Data Share Kusto Cluster Dataset.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.datashare.getDatasetKustoCluster({
 *     name: "example-dskc",
 *     shareId: "example-share-id",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getDatasetKustoCluster(args: GetDatasetKustoClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetDatasetKustoClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:datashare/getDatasetKustoCluster:getDatasetKustoCluster", {
        "name": args.name,
        "shareId": args.shareId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatasetKustoCluster.
 */
export interface GetDatasetKustoClusterArgs {
    /**
     * The name of this Data Share Kusto Cluster Dataset.
     */
    readonly name: string;
    /**
     * The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created.
     */
    readonly shareId: string;
}

/**
 * A collection of values returned by getDatasetKustoCluster.
 */
export interface GetDatasetKustoClusterResult {
    /**
     * The name of the Data Share Dataset.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The resource ID of the Kusto Cluster to be shared with the receiver.
     */
    readonly kustoClusterId: string;
    /**
     * The location of the Kusto Cluster.
     */
    readonly kustoClusterLocation: string;
    readonly name: string;
    readonly shareId: string;
}
