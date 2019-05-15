// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Dev Test Lab.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = pulumi.output(azure.devtest.getLab({
 *     name: "example-lab",
 *     resourceGroupName: "example-resources",
 * }));
 * 
 * export const uniqueIdentifier = test.uniqueIdentifier;
 * ```
 */
export function getLab(args: GetLabArgs, opts?: pulumi.InvokeOptions): Promise<GetLabResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:devtest/getLab:getLab", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getLab.
 */
export interface GetLabArgs {
    /**
     * The name of the Dev Test Lab.
     */
    readonly name: string;
    /**
     * The Name of the Resource Group where the Dev Test Lab exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getLab.
 */
export interface GetLabResult {
    /**
     * The ID of the Storage Account used for Artifact Storage.
     */
    readonly artifactsStorageAccountId: string;
    /**
     * The ID of the Default Premium Storage Account for this Dev Test Lab.
     */
    readonly defaultPremiumStorageAccountId: string;
    /**
     * The ID of the Default Storage Account for this Dev Test Lab.
     */
    readonly defaultStorageAccountId: string;
    /**
     * The ID of the Key used for this Dev Test Lab.
     */
    readonly keyVaultId: string;
    /**
     * The Azure location where the Dev Test Lab exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * The ID of the Storage Account used for Storage of Premium Data Disk.
     */
    readonly premiumDataDiskStorageAccountId: string;
    readonly resourceGroupName: string;
    /**
     * The type of storage used by the Dev Test Lab.
     */
    readonly storageType: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * The unique immutable identifier of the Dev Test Lab.
     */
    readonly uniqueIdentifier: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
