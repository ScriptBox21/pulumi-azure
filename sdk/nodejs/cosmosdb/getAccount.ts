// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.cosmosdb.getAccount({
 *     name: "tfex-cosmosdb-account",
 *     resourceGroupName: "tfex-cosmosdb-account-rg",
 * });
 * export const cosmosdbAccountEndpoint = data.azurerm_cosmosdb_account.jobs.endpoint;
 * ```
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:cosmosdb/getAccount:getAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountArgs {
    /**
     * Specifies the name of the CosmosDB Account.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group in which the CosmosDB Account resides.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * Capabilities enabled on this Cosmos DB account.
     */
    readonly capabilities: outputs.cosmosdb.GetAccountCapability[];
    readonly consistencyPolicies: outputs.cosmosdb.GetAccountConsistencyPolicy[];
    /**
     * If automatic failover is enabled for this CosmosDB Account.
     */
    readonly enableAutomaticFailover: boolean;
    /**
     * If Free Tier pricing option is enabled for this CosmosDB Account.
     */
    readonly enableFreeTier: boolean;
    /**
     * If multi-master is enabled for this Cosmos DB account.
     */
    readonly enableMultipleWriteLocations: boolean;
    /**
     * The endpoint used to connect to the CosmosDB account.
     */
    readonly endpoint: string;
    readonly geoLocations: outputs.cosmosdb.GetAccountGeoLocation[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The current IP Filter for this CosmosDB account
     */
    readonly ipRangeFilter: string;
    /**
     * If virtual network filtering is enabled for this Cosmos DB account.
     */
    readonly isVirtualNetworkFilterEnabled: boolean;
    /**
     * The Key Vault key URI for CMK encryption.
     */
    readonly keyVaultKeyId: string;
    /**
     * The Kind of the CosmosDB account.
     */
    readonly kind: string;
    /**
     * The name of the Azure region hosting replicated data.
     */
    readonly location: string;
    readonly name: string;
    /**
     * The Offer Type to used by this CosmosDB Account.
     */
    readonly offerType: string;
    /**
     * The Primary master key for the CosmosDB Account.
     */
    readonly primaryKey: string;
    /**
     * @deprecated This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    readonly primaryMasterKey: string;
    /**
     * The Primary read-only master Key for the CosmosDB Account.
     */
    readonly primaryReadonlyKey: string;
    /**
     * @deprecated This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    readonly primaryReadonlyMasterKey: string;
    /**
     * A list of read endpoints available for this CosmosDB account.
     */
    readonly readEndpoints: string[];
    readonly resourceGroupName: string;
    /**
     * The Secondary master key for the CosmosDB Account.
     */
    readonly secondaryKey: string;
    /**
     * @deprecated This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    readonly secondaryMasterKey: string;
    /**
     * The Secondary read-only master key for the CosmosDB Account.
     */
    readonly secondaryReadonlyKey: string;
    /**
     * @deprecated This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    readonly secondaryReadonlyMasterKey: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * Subnets that are allowed to access this CosmosDB account.
     */
    readonly virtualNetworkRules: outputs.cosmosdb.GetAccountVirtualNetworkRule[];
    /**
     * A list of write endpoints available for this CosmosDB account.
     */
    readonly writeEndpoints: string[];
}
