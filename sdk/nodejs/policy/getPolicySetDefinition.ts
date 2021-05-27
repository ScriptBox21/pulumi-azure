// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Policy Set Definition.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.policy.getPolicySetDefinition({
 *     displayName: "Policy Set Definition Example",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getPolicySetDefinition(args?: GetPolicySetDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicySetDefinitionResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:policy/getPolicySetDefinition:getPolicySetDefinition", {
        "displayName": args.displayName,
        "managementGroupName": args.managementGroupName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicySetDefinition.
 */
export interface GetPolicySetDefinitionArgs {
    /**
     * Specifies the display name of the Policy Set Definition. Conflicts with `name`.
     */
    displayName?: string;
    /**
     * Only retrieve Policy Set Definitions from this Management Group.
     */
    managementGroupName?: string;
    /**
     * Specifies the name of the Policy Set Definition. Conflicts with `displayName`.
     */
    name?: string;
}

/**
 * A collection of values returned by getPolicySetDefinition.
 */
export interface GetPolicySetDefinitionResult {
    /**
     * The description of this policy definition group.
     */
    readonly description: string;
    /**
     * The display name of this policy definition group.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly managementGroupName?: string;
    /**
     * Any Metadata defined in the Policy Set Definition.
     */
    readonly metadata: string;
    /**
     * The name of this policy definition group.
     */
    readonly name: string;
    /**
     * The mapping of the parameter values for the referenced policy rule. The keys are the parameter names.
     */
    readonly parameters: string;
    /**
     * One or more `policyDefinitionGroup` blocks as defined below.
     */
    readonly policyDefinitionGroups: outputs.policy.GetPolicySetDefinitionPolicyDefinitionGroup[];
    /**
     * One or more `policyDefinitionReference` blocks as defined below.
     */
    readonly policyDefinitionReferences: outputs.policy.GetPolicySetDefinitionPolicyDefinitionReference[];
    /**
     * The policy definitions contained within the policy set definition.
     */
    readonly policyDefinitions: string;
    /**
     * The Type of the Policy Set Definition.
     */
    readonly policyType: string;
}
