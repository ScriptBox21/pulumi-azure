// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Storage Table Entity.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.storage.getTableEntity({
 *     partitionKey: "example-parition-key",
 *     rowKey: "example-row-key",
 *     storageAccountName: "example-storage-account-name",
 *     tableName: "example-table-name",
 * }));
 * ```
 */
export function getTableEntity(args: GetTableEntityArgs, opts?: pulumi.InvokeOptions): Promise<GetTableEntityResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:storage/getTableEntity:getTableEntity", {
        "partitionKey": args.partitionKey,
        "rowKey": args.rowKey,
        "storageAccountName": args.storageAccountName,
        "tableName": args.tableName,
    }, opts);
}

/**
 * A collection of arguments for invoking getTableEntity.
 */
export interface GetTableEntityArgs {
    /**
     * The key for the partition where the entity will be retrieved.
     */
    partitionKey: string;
    /**
     * The key for the row where the entity will be retrieved.
     */
    rowKey: string;
    /**
     * The name of the Storage Account where the Table exists.
     */
    storageAccountName: string;
    /**
     * The name of the Table.
     */
    tableName: string;
}

/**
 * A collection of values returned by getTableEntity.
 */
export interface GetTableEntityResult {
    /**
     * A map of key/value pairs that describe the entity to be stored in the storage table.
     */
    readonly entity: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly partitionKey: string;
    readonly rowKey: string;
    readonly storageAccountName: string;
    readonly tableName: string;
}
