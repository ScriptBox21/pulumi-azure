// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Media Services Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "GRS",
 * });
 * const exampleServiceAccount = new azure.media.ServiceAccount("exampleServiceAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     storageAccounts: [{
 *         id: exampleAccount.id,
 *         isPrimary: true,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Media Services Accounts can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:media/serviceAccount:ServiceAccount account /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Media/mediaservices/account1
 * ```
 */
export class ServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceAccountState, opts?: pulumi.CustomResourceOptions): ServiceAccount {
        return new ServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:media/serviceAccount:ServiceAccount';

    /**
     * Returns true if the given object is an instance of ServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAccount.__pulumiType;
    }

    /**
     * An `identity` block is documented below.
     */
    public readonly identity!: pulumi.Output<outputs.media.ServiceAccountIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * One or more `storageAccount` blocks as defined below.
     */
    public readonly storageAccounts!: pulumi.Output<outputs.media.ServiceAccountStorageAccount[]>;
    /**
     * Specifies the storage authentication type. 
     * Possible value is  `ManagedIdentity` or `System`.
     */
    public readonly storageAuthenticationType!: pulumi.Output<string>;
    /**
     * A mapping of tags assigned to the resource.
     * ---
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceAccountArgs | ServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceAccountState | undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["storageAccounts"] = state ? state.storageAccounts : undefined;
            inputs["storageAuthenticationType"] = state ? state.storageAuthenticationType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ServiceAccountArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.storageAccounts === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'storageAccounts'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccounts"] = args ? args.storageAccounts : undefined;
            inputs["storageAuthenticationType"] = args ? args.storageAuthenticationType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure:mediaservices/account:Account" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ServiceAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceAccount resources.
 */
export interface ServiceAccountState {
    /**
     * An `identity` block is documented below.
     */
    readonly identity?: pulumi.Input<inputs.media.ServiceAccountIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * One or more `storageAccount` blocks as defined below.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.media.ServiceAccountStorageAccount>[]>;
    /**
     * Specifies the storage authentication type. 
     * Possible value is  `ManagedIdentity` or `System`.
     */
    readonly storageAuthenticationType?: pulumi.Input<string>;
    /**
     * A mapping of tags assigned to the resource.
     * ---
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ServiceAccount resource.
 */
export interface ServiceAccountArgs {
    /**
     * An `identity` block is documented below.
     */
    readonly identity?: pulumi.Input<inputs.media.ServiceAccountIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * One or more `storageAccount` blocks as defined below.
     */
    readonly storageAccounts: pulumi.Input<pulumi.Input<inputs.media.ServiceAccountStorageAccount>[]>;
    /**
     * Specifies the storage authentication type. 
     * Possible value is  `ManagedIdentity` or `System`.
     */
    readonly storageAuthenticationType?: pulumi.Input<string>;
    /**
     * A mapping of tags assigned to the resource.
     * ---
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
