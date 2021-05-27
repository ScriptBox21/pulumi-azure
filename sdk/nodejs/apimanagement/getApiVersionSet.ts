// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Uses this data source to access information about an API Version Set within an API Management Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.apimanagement.getApiVersionSet({
 *     resourceGroupName: "example-resources",
 *     apiManagementName: "example-api",
 *     name: "example-api-version-set",
 * });
 * export const apiManagementApiVersionSetId = example.then(example => example.id);
 * ```
 */
export function getApiVersionSet(args: GetApiVersionSetArgs, opts?: pulumi.InvokeOptions): Promise<GetApiVersionSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:apimanagement/getApiVersionSet:getApiVersionSet", {
        "apiManagementName": args.apiManagementName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiVersionSet.
 */
export interface GetApiVersionSetArgs {
    /**
     * The name of the API Management Service where the API Version Set exists.
     */
    apiManagementName: string;
    /**
     * The name of the API Version Set.
     */
    name: string;
    /**
     * The name of the Resource Group in which the parent API Management Service exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getApiVersionSet.
 */
export interface GetApiVersionSetResult {
    readonly apiManagementName: string;
    /**
     * The description of API Version Set.
     */
    readonly description: string;
    /**
     * The display name of this API Version Set.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The name of the Header which should be read from Inbound Requests which defines the API Version.
     */
    readonly versionHeaderName: string;
    /**
     * The name of the Query String which should be read from Inbound Requests which defines the API Version.
     */
    readonly versionQueryName: string;
    readonly versioningScheme: string;
}
