// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Role Definition.
 *
 */
export function getRoleDefinition(args?: GetRoleDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleDefinitionResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:authorization/getRoleDefinition:getRoleDefinition", {
        "name": args.name,
        "roleDefinitionId": args.roleDefinitionId,
        "scope": args.scope,
    }, opts);
}

/**
 * A collection of arguments for invoking getRoleDefinition.
 */
export interface GetRoleDefinitionArgs {
    /**
     * Specifies the Name of either a built-in or custom Role Definition.
     */
    readonly name?: string;
    /**
     * Specifies the ID of the Role Definition as a UUID/GUID.
     */
    readonly roleDefinitionId?: string;
    /**
     * Specifies the Scope at which the Custom Role Definition exists.
     */
    readonly scope?: string;
}

/**
 * A collection of values returned by getRoleDefinition.
 */
export interface GetRoleDefinitionResult {
    /**
     * One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
     */
    readonly assignableScopes: string[];
    /**
     * the Description of the built-in Role.
     */
    readonly description: string;
    readonly name: string;
    /**
     * a `permissions` block as documented below.
     */
    readonly permissions: outputs.authorization.GetRoleDefinitionPermission[];
    readonly roleDefinitionId: string;
    readonly scope?: string;
    /**
     * the Type of the Role.
     */
    readonly type: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
