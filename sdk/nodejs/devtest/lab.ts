// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Dev Test Lab.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/dev_test_lab.html.markdown.
 */
export class Lab extends pulumi.CustomResource {
    /**
     * Get an existing Lab resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LabState, opts?: pulumi.CustomResourceOptions): Lab {
        return new Lab(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:devtest/lab:Lab';

    /**
     * Returns true if the given object is an instance of Lab.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lab {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lab.__pulumiType;
    }

    /**
     * The ID of the Storage Account used for Artifact Storage.
     */
    public /*out*/ readonly artifactsStorageAccountId!: pulumi.Output<string>;
    /**
     * The ID of the Default Premium Storage Account for this Dev Test Lab.
     */
    public /*out*/ readonly defaultPremiumStorageAccountId!: pulumi.Output<string>;
    /**
     * The ID of the Default Storage Account for this Dev Test Lab.
     */
    public /*out*/ readonly defaultStorageAccountId!: pulumi.Output<string>;
    /**
     * The ID of the Key used for this Dev Test Lab.
     */
    public /*out*/ readonly keyVaultId!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Storage Account used for Storage of Premium Data Disk.
     */
    public /*out*/ readonly premiumDataDiskStorageAccountId!: pulumi.Output<string>;
    /**
     * The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The unique immutable identifier of the Dev Test Lab.
     */
    public /*out*/ readonly uniqueIdentifier!: pulumi.Output<string>;

    /**
     * Create a Lab resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LabArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LabArgs | LabState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LabState | undefined;
            inputs["artifactsStorageAccountId"] = state ? state.artifactsStorageAccountId : undefined;
            inputs["defaultPremiumStorageAccountId"] = state ? state.defaultPremiumStorageAccountId : undefined;
            inputs["defaultStorageAccountId"] = state ? state.defaultStorageAccountId : undefined;
            inputs["keyVaultId"] = state ? state.keyVaultId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["premiumDataDiskStorageAccountId"] = state ? state.premiumDataDiskStorageAccountId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["storageType"] = state ? state.storageType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uniqueIdentifier"] = state ? state.uniqueIdentifier : undefined;
        } else {
            const args = argsOrState as LabArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["artifactsStorageAccountId"] = undefined /*out*/;
            inputs["defaultPremiumStorageAccountId"] = undefined /*out*/;
            inputs["defaultStorageAccountId"] = undefined /*out*/;
            inputs["keyVaultId"] = undefined /*out*/;
            inputs["premiumDataDiskStorageAccountId"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Lab.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lab resources.
 */
export interface LabState {
    /**
     * The ID of the Storage Account used for Artifact Storage.
     */
    readonly artifactsStorageAccountId?: pulumi.Input<string>;
    /**
     * The ID of the Default Premium Storage Account for this Dev Test Lab.
     */
    readonly defaultPremiumStorageAccountId?: pulumi.Input<string>;
    /**
     * The ID of the Default Storage Account for this Dev Test Lab.
     */
    readonly defaultStorageAccountId?: pulumi.Input<string>;
    /**
     * The ID of the Key used for this Dev Test Lab.
     */
    readonly keyVaultId?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account used for Storage of Premium Data Disk.
     */
    readonly premiumDataDiskStorageAccountId?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique immutable identifier of the Dev Test Lab.
     */
    readonly uniqueIdentifier?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Lab resource.
 */
export interface LabArgs {
    /**
     * Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Dev Test Lab. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The type of storage used by the Dev Test Lab. Possible values are `Standard` and `Premium`. Defaults to `Premium`. Changing this forces a new resource to be created.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
