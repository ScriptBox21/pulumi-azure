// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Event Hubs Authorization Rule within an Event Hub.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const test = pulumi.all([azurerm_eventhub_test.name, azurerm_eventhub_namespace_test.name, azurerm_resource_group_test.name]).apply(([azurerm_eventhub_testName, azurerm_eventhub_namespace_testName, azurerm_resource_group_testName]) => azure.eventhub.getAuthorizationRule({
 *     eventhubName: azurerm_eventhub_testName,
 *     name: "test",
 *     namespaceName: azurerm_eventhub_namespace_testName,
 *     resourceGroupName: azurerm_resource_group_testName,
 * }, { async: true }));
 * ```
 */
export function getAuthorizationRule(args: GetAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:eventhub/getAuthorizationRule:getAuthorizationRule", {
        "eventhubName": args.eventhubName,
        "listen": args.listen,
        "manage": args.manage,
        "name": args.name,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
        "send": args.send,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthorizationRule.
 */
export interface GetAuthorizationRuleArgs {
    /**
     * Specifies the name of the EventHub.
     */
    readonly eventhubName: string;
    readonly listen?: boolean;
    readonly manage?: boolean;
    /**
     * Specifies the name of the EventHub Authorization Rule resource. be created.
     */
    readonly name: string;
    /**
     * Specifies the name of the grandparent EventHub Namespace.
     */
    readonly namespaceName: string;
    /**
     * The name of the resource group in which the EventHub Authorization Rule's grandparent Namespace exists.
     */
    readonly resourceGroupName: string;
    readonly send?: boolean;
}

/**
 * A collection of values returned by getAuthorizationRule.
 */
export interface GetAuthorizationRuleResult {
    readonly eventhubName: string;
    readonly listen?: boolean;
    readonly location: string;
    readonly manage?: boolean;
    readonly name: string;
    readonly namespaceName: string;
    /**
     * The Primary Connection String for the Event Hubs Authorization Rule.
     */
    readonly primaryConnectionString: string;
    /**
     * The alias of the Primary Connection String for the Event Hubs Authorization Rule.
     */
    readonly primaryConnectionStringAlias: string;
    /**
     * The Primary Key for the Event Hubs Authorization Rule.
     */
    readonly primaryKey: string;
    readonly resourceGroupName: string;
    /**
     * The Secondary Connection String for the Event Hubs Authorization Rule.
     */
    readonly secondaryConnectionString: string;
    /**
     * The alias of the Secondary Connection String for the Event Hubs Authorization Rule.
     */
    readonly secondaryConnectionStringAlias: string;
    /**
     * The Secondary Key for the Event Hubs Authorization Rule.
     */
    readonly secondaryKey: string;
    readonly send?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
