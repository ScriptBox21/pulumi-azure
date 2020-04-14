// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Azure Container Registry.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_registry.html.markdown.
 */
export class Registry extends pulumi.CustomResource {
    /**
     * Get an existing Registry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryState, opts?: pulumi.CustomResourceOptions): Registry {
        return new Registry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:containerservice/registry:Registry';

    /**
     * Returns true if the given object is an instance of Registry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registry.__pulumiType;
    }

    /**
     * Specifies whether the admin user is enabled. Defaults to `false`.
     */
    public readonly adminEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Password associated with the Container Registry Admin account - if the admin account is enabled.
     */
    public /*out*/ readonly adminPassword!: pulumi.Output<string>;
    /**
     * The Username associated with the Container Registry Admin account - if the admin account is enabled.
     */
    public /*out*/ readonly adminUsername!: pulumi.Output<string>;
    /**
     * A list of Azure locations where the container registry should be geo-replicated.
     */
    public readonly georeplicationLocations!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The URL that can be used to log into the container registry.
     */
    public /*out*/ readonly loginServer!: pulumi.Output<string>;
    /**
     * Specifies the name of the Container Registry. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `networkRuleSet` block as documented below.
     */
    public readonly networkRuleSet!: pulumi.Output<outputs.containerservice.RegistryNetworkRuleSet>;
    /**
     * The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
     */
    public readonly sku!: pulumi.Output<string | undefined>;
    /**
     * The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.
     */
    public readonly storageAccountId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Registry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryArgs | RegistryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegistryState | undefined;
            inputs["adminEnabled"] = state ? state.adminEnabled : undefined;
            inputs["adminPassword"] = state ? state.adminPassword : undefined;
            inputs["adminUsername"] = state ? state.adminUsername : undefined;
            inputs["georeplicationLocations"] = state ? state.georeplicationLocations : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["loginServer"] = state ? state.loginServer : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkRuleSet"] = state ? state.networkRuleSet : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RegistryArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["adminEnabled"] = args ? args.adminEnabled : undefined;
            inputs["georeplicationLocations"] = args ? args.georeplicationLocations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkRuleSet"] = args ? args.networkRuleSet : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["adminPassword"] = undefined /*out*/;
            inputs["adminUsername"] = undefined /*out*/;
            inputs["loginServer"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Registry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Registry resources.
 */
export interface RegistryState {
    /**
     * Specifies whether the admin user is enabled. Defaults to `false`.
     */
    readonly adminEnabled?: pulumi.Input<boolean>;
    /**
     * The Password associated with the Container Registry Admin account - if the admin account is enabled.
     */
    readonly adminPassword?: pulumi.Input<string>;
    /**
     * The Username associated with the Container Registry Admin account - if the admin account is enabled.
     */
    readonly adminUsername?: pulumi.Input<string>;
    /**
     * A list of Azure locations where the container registry should be geo-replicated.
     */
    readonly georeplicationLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The URL that can be used to log into the container registry.
     */
    readonly loginServer?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `networkRuleSet` block as documented below.
     */
    readonly networkRuleSet?: pulumi.Input<inputs.containerservice.RegistryNetworkRuleSet>;
    /**
     * The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.
     */
    readonly storageAccountId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Registry resource.
 */
export interface RegistryArgs {
    /**
     * Specifies whether the admin user is enabled. Defaults to `false`.
     */
    readonly adminEnabled?: pulumi.Input<boolean>;
    /**
     * A list of Azure locations where the container registry should be geo-replicated.
     */
    readonly georeplicationLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A `networkRuleSet` block as documented below.
     */
    readonly networkRuleSet?: pulumi.Input<inputs.containerservice.RegistryNetworkRuleSet>;
    /**
     * The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.
     */
    readonly storageAccountId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
