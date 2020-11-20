// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.servicebus.getTopicAuthorizationRule({
 *     name: "example-tfex_name",
 *     namespaceName: "example-namespace",
 *     resourceGroupName: "example-resources",
 *     topicName: "example-servicebus_topic",
 * }, { async: true }));
 *
 * export const servicebusAuthorizationRuleId = azurem_servicebus_topic_authorization_rule_example.id;
 * ```
 */
export function getTopicAuthorizationRule(args: GetTopicAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:servicebus/getTopicAuthorizationRule:getTopicAuthorizationRule", {
        "name": args.name,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
        "topicName": args.topicName,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopicAuthorizationRule.
 */
export interface GetTopicAuthorizationRuleArgs {
    /**
     * The name of the ServiceBus Topic Authorization Rule resource.
     */
    readonly name: string;
    /**
     * The name of the ServiceBus Namespace.
     */
    readonly namespaceName: string;
    /**
     * The name of the resource group in which the ServiceBus Namespace exists.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the ServiceBus Topic.
     */
    readonly topicName: string;
}

/**
 * A collection of values returned by getTopicAuthorizationRule.
 */
export interface GetTopicAuthorizationRuleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listen: boolean;
    readonly manage: boolean;
    readonly name: string;
    readonly namespaceName: string;
    /**
     * The Primary Connection String for the ServiceBus Topic authorization Rule.
     */
    readonly primaryConnectionString: string;
    /**
     * The Primary Key for the ServiceBus Topic authorization Rule.
     */
    readonly primaryKey: string;
    readonly resourceGroupName: string;
    /**
     * The Secondary Connection String for the ServiceBus Topic authorization Rule.
     */
    readonly secondaryConnectionString: string;
    /**
     * The Secondary Key for the ServiceBus Topic authorization Rule.
     */
    readonly secondaryKey: string;
    readonly send: boolean;
    readonly topicName: string;
}
