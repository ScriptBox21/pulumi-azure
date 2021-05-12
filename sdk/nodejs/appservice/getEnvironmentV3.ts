// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing 3rd Generation (v3) App Service Environment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.appservice.getEnvironmentV3({
 *     name: "example-ASE",
 *     resourceGroupName: "example-resource-group",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getEnvironmentV3(args: GetEnvironmentV3Args, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentV3Result> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:appservice/getEnvironmentV3:getEnvironmentV3", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnvironmentV3.
 */
export interface GetEnvironmentV3Args {
    /**
     * The name of this v3 App Service Environment.
     */
    readonly name: string;
    /**
     * The name of the Resource Group where the v3 App Service Environment exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getEnvironmentV3.
 */
export interface GetEnvironmentV3Result {
    /**
     * A `clusterSetting` block as defined below.
     */
    readonly clusterSettings: outputs.appservice.GetEnvironmentV3ClusterSetting[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the Cluster Setting.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The ID of the v3 App Service Environment Subnet.
     */
    readonly subnetId: string;
    /**
     * A mapping of tags assigned to the v3 App Service Environment.
     */
    readonly tags: {[key: string]: string};
}
