// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Azure App Configuration.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const rg = new azure.core.ResourceGroup("rg", {
 *     location: "West Europe",
 * });
 * const appconf = new azure.appconfiguration.ConfigurationStore("appconf", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_configuration.html.markdown.
 */
export class ConfigurationStore extends pulumi.CustomResource {
    /**
     * Get an existing ConfigurationStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationStoreState, opts?: pulumi.CustomResourceOptions): ConfigurationStore {
        return new ConfigurationStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appconfiguration/configurationStore:ConfigurationStore';

    /**
     * Returns true if the given object is an instance of ConfigurationStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigurationStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigurationStore.__pulumiType;
    }

    /**
     * The URL of the App Configuration.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the App Configuration. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An `accessKey` block as defined below containing the primary read access key.
     */
    public /*out*/ readonly primaryReadKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStorePrimaryReadKey[]>;
    /**
     * An `accessKey` block as defined below containing the primary write access key.
     */
    public /*out*/ readonly primaryWriteKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStorePrimaryWriteKey[]>;
    /**
     * The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * An `accessKey` block as defined below containing the secondary read access key.
     */
    public /*out*/ readonly secondaryReadKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStoreSecondaryReadKey[]>;
    /**
     * An `accessKey` block as defined below containing the secondary write access key.
     */
    public /*out*/ readonly secondaryWriteKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStoreSecondaryWriteKey[]>;
    /**
     * The SKU name of the the App Configuration. Possible values are `free` and `standard`.
     */
    public readonly sku!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a ConfigurationStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigurationStoreArgs | ConfigurationStoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConfigurationStoreState | undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["primaryReadKeys"] = state ? state.primaryReadKeys : undefined;
            inputs["primaryWriteKeys"] = state ? state.primaryWriteKeys : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryReadKeys"] = state ? state.secondaryReadKeys : undefined;
            inputs["secondaryWriteKeys"] = state ? state.secondaryWriteKeys : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ConfigurationStoreArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["endpoint"] = undefined /*out*/;
            inputs["primaryReadKeys"] = undefined /*out*/;
            inputs["primaryWriteKeys"] = undefined /*out*/;
            inputs["secondaryReadKeys"] = undefined /*out*/;
            inputs["secondaryWriteKeys"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConfigurationStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigurationStore resources.
 */
export interface ConfigurationStoreState {
    /**
     * The URL of the App Configuration.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the App Configuration. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An `accessKey` block as defined below containing the primary read access key.
     */
    readonly primaryReadKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStorePrimaryReadKey>[]>;
    /**
     * An `accessKey` block as defined below containing the primary write access key.
     */
    readonly primaryWriteKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStorePrimaryWriteKey>[]>;
    /**
     * The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * An `accessKey` block as defined below containing the secondary read access key.
     */
    readonly secondaryReadKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStoreSecondaryReadKey>[]>;
    /**
     * An `accessKey` block as defined below containing the secondary write access key.
     */
    readonly secondaryWriteKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStoreSecondaryWriteKey>[]>;
    /**
     * The SKU name of the the App Configuration. Possible values are `free` and `standard`.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ConfigurationStore resource.
 */
export interface ConfigurationStoreArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the App Configuration. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU name of the the App Configuration. Possible values are `free` and `standard`.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
