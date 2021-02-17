// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a user assigned identity.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "eastus"});
 * const exampleUserAssignedIdentity = new azure.authorization.UserAssignedIdentity("exampleUserAssignedIdentity", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * ```
 *
 * ## Import
 *
 * User Assigned Identities can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:msi/userAssignedIdentity:UserAssignedIdentity exampleIdentity /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acceptanceTestResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity
 * ```
 *
 * @deprecated azure.msi.UserAssignedIdentity has been deprecated in favor of azure.authorization.UserAssignedIdentity
 */
export class UserAssignedIdentity extends pulumi.CustomResource {
    /**
     * Get an existing UserAssignedIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserAssignedIdentityState, opts?: pulumi.CustomResourceOptions): UserAssignedIdentity {
        pulumi.log.warn("UserAssignedIdentity is deprecated: azure.msi.UserAssignedIdentity has been deprecated in favor of azure.authorization.UserAssignedIdentity")
        return new UserAssignedIdentity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:msi/userAssignedIdentity:UserAssignedIdentity';

    /**
     * Returns true if the given object is an instance of UserAssignedIdentity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserAssignedIdentity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserAssignedIdentity.__pulumiType;
    }

    /**
     * Client ID associated with the user assigned identity.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * The location/region where the user assigned identity is
     * created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the user assigned identity. Changing this forces a
     * new identity to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Service Principal ID associated with the user assigned identity.
     */
    public /*out*/ readonly principalId!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the user assigned identity.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a UserAssignedIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.msi.UserAssignedIdentity has been deprecated in favor of azure.authorization.UserAssignedIdentity */
    constructor(name: string, args: UserAssignedIdentityArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.msi.UserAssignedIdentity has been deprecated in favor of azure.authorization.UserAssignedIdentity */
    constructor(name: string, argsOrState?: UserAssignedIdentityArgs | UserAssignedIdentityState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("UserAssignedIdentity is deprecated: azure.msi.UserAssignedIdentity has been deprecated in favor of azure.authorization.UserAssignedIdentity")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserAssignedIdentityState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["principalId"] = state ? state.principalId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as UserAssignedIdentityArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["clientId"] = undefined /*out*/;
            inputs["principalId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserAssignedIdentity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserAssignedIdentity resources.
 */
export interface UserAssignedIdentityState {
    /**
     * Client ID associated with the user assigned identity.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The location/region where the user assigned identity is
     * created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the user assigned identity. Changing this forces a
     * new identity to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Service Principal ID associated with the user assigned identity.
     */
    readonly principalId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the user assigned identity.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a UserAssignedIdentity resource.
 */
export interface UserAssignedIdentityArgs {
    /**
     * The location/region where the user assigned identity is
     * created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the user assigned identity. Changing this forces a
     * new identity to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the user assigned identity.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
