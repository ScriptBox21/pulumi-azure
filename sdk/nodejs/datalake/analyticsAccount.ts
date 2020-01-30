// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Data Lake Analytics Account.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "northeurope",
 * });
 * const exampleStore = new azure.datalake.Store("example", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleAnalyticsAccount = new azure.datalake.AnalyticsAccount("example", {
 *     defaultStoreAccountName: exampleStore.name,
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/data_lake_analytics_account.html.markdown.
 */
export class AnalyticsAccount extends pulumi.CustomResource {
    /**
     * Get an existing AnalyticsAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyticsAccountState, opts?: pulumi.CustomResourceOptions): AnalyticsAccount {
        return new AnalyticsAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datalake/analyticsAccount:AnalyticsAccount';

    /**
     * Returns true if the given object is an instance of AnalyticsAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnalyticsAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnalyticsAccount.__pulumiType;
    }

    /**
     * Specifies the data lake store to use by default. Changing this forces a new resource to be created.
     */
    public readonly defaultStoreAccountName!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Data Lake Analytics Account. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Data Lake Analytics Account.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The monthly commitment tier for Data Lake Analytics Account. Accepted values are `Consumption`, `Commitment_100000AUHours`, `Commitment_10000AUHours`, `Commitment_1000AUHours`, `Commitment_100AUHours`, `Commitment_500000AUHours`, `Commitment_50000AUHours`, `Commitment_5000AUHours`, or `Commitment_500AUHours`.
     */
    public readonly tier!: pulumi.Output<string | undefined>;

    /**
     * Create a AnalyticsAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyticsAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnalyticsAccountArgs | AnalyticsAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AnalyticsAccountState | undefined;
            inputs["defaultStoreAccountName"] = state ? state.defaultStoreAccountName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tier"] = state ? state.tier : undefined;
        } else {
            const args = argsOrState as AnalyticsAccountArgs | undefined;
            if (!args || args.defaultStoreAccountName === undefined) {
                throw new Error("Missing required property 'defaultStoreAccountName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["defaultStoreAccountName"] = args ? args.defaultStoreAccountName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tier"] = args ? args.tier : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AnalyticsAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnalyticsAccount resources.
 */
export interface AnalyticsAccountState {
    /**
     * Specifies the data lake store to use by default. Changing this forces a new resource to be created.
     */
    readonly defaultStoreAccountName?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Lake Analytics Account. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Data Lake Analytics Account.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The monthly commitment tier for Data Lake Analytics Account. Accepted values are `Consumption`, `Commitment_100000AUHours`, `Commitment_10000AUHours`, `Commitment_1000AUHours`, `Commitment_100AUHours`, `Commitment_500000AUHours`, `Commitment_50000AUHours`, `Commitment_5000AUHours`, or `Commitment_500AUHours`.
     */
    readonly tier?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AnalyticsAccount resource.
 */
export interface AnalyticsAccountArgs {
    /**
     * Specifies the data lake store to use by default. Changing this forces a new resource to be created.
     */
    readonly defaultStoreAccountName: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Lake Analytics Account. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Data Lake Analytics Account.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The monthly commitment tier for Data Lake Analytics Account. Accepted values are `Consumption`, `Commitment_100000AUHours`, `Commitment_10000AUHours`, `Commitment_1000AUHours`, `Commitment_100AUHours`, `Commitment_500000AUHours`, `Commitment_50000AUHours`, `Commitment_5000AUHours`, or `Commitment_500AUHours`.
     */
    readonly tier?: pulumi.Input<string>;
}
