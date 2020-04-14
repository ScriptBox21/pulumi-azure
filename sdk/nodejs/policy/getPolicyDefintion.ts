// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about a Policy Definition, both custom and built in. Retrieves Policy Definitions from your current subscription by default.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/policy_definition.html.markdown.
 */
export function getPolicyDefintion(args: GetPolicyDefintionArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyDefintionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:policy/getPolicyDefintion:getPolicyDefintion", {
        "displayName": args.displayName,
        "managementGroupId": args.managementGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicyDefintion.
 */
export interface GetPolicyDefintionArgs {
    /**
     * Specifies the name of the Policy Definition.
     */
    readonly displayName: string;
    /**
     * Only retrieve Policy Definitions from this Management Group.
     */
    readonly managementGroupId?: string;
}

/**
 * A collection of values returned by getPolicyDefintion.
 */
export interface GetPolicyDefintionResult {
    /**
     * The Description of the Policy.
     */
    readonly description: string;
    readonly displayName: string;
    readonly managementGroupId?: string;
    /**
     * Any Metadata defined in the Policy.
     */
    readonly metadata: string;
    /**
     * The Name of the Policy Definition.
     */
    readonly name: string;
    /**
     * Any Parameters defined in the Policy.
     */
    readonly parameters: string;
    /**
     * The Rule as defined (in JSON) in the Policy.
     */
    readonly policyRule: string;
    /**
     * The Type of the Policy, such as `Microsoft.Authorization/policyDefinitions`.
     */
    readonly policyType: string;
    /**
     * The Type of Policy.
     */
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
