// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Key Vault.
 *
 * ## Disclaimers
 *
 * > **Note:** It's possible to define Key Vault Access Policies both within the `azure.keyvault.KeyVault` resource via the `accessPolicy` block and by using the `azure.keyvault.AccessPolicy` resource. However it's not possible to use both methods to manage Access Policies within a KeyVault, since there'll be conflicts.
 *
 * > **Note:** This provider will automatically recover a soft-deleted Key Vault during Creation if one is found - you can opt out of this using the `features` configuration within the Provider configuration block.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     enabledForDiskEncryption: true,
 *     tenantId: current.then(current => current.tenantId),
 *     softDeleteRetentionDays: 7,
 *     purgeProtectionEnabled: false,
 *     skuName: "standard",
 *     accessPolicies: [{
 *         tenantId: current.then(current => current.tenantId),
 *         objectId: current.then(current => current.objectId),
 *         keyPermissions: ["Get"],
 *         secretPermissions: ["Get"],
 *         storagePermissions: ["Get"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Key Vault's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:keyvault/keyVault:KeyVault example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/vault1
 * ```
 */
export class KeyVault extends pulumi.CustomResource {
    /**
     * Get an existing KeyVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyVaultState, opts?: pulumi.CustomResourceOptions): KeyVault {
        return new KeyVault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:keyvault/keyVault:KeyVault';

    /**
     * Returns true if the given object is an instance of KeyVault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyVault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyVault.__pulumiType;
    }

    /**
     * A list of up to 16 objects describing access policies, as described below.
     */
    public readonly accessPolicies!: pulumi.Output<outputs.keyvault.KeyVaultAccessPolicy[]>;
    /**
     * One or more `contact` block as defined below.
     */
    public readonly contacts!: pulumi.Output<outputs.keyvault.KeyVaultContact[] | undefined>;
    /**
     * Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
     */
    public readonly enableRbacAuthorization!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
     */
    public readonly enabledForDeployment!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
     */
    public readonly enabledForDiskEncryption!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
     */
    public readonly enabledForTemplateDeployment!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Key Vault. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `networkAcls` block as defined below.
     */
    public readonly networkAcls!: pulumi.Output<outputs.keyvault.KeyVaultNetworkAcls>;
    /**
     * Is Purge Protection enabled for this Key Vault? Defaults to `false`.
     */
    public readonly purgeProtectionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * @deprecated Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
     */
    public readonly softDeleteEnabled!: pulumi.Output<boolean>;
    /**
     * The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
     */
    public readonly softDeleteRetentionDays!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The URI of the Key Vault, used for performing operations on keys and secrets.
     */
    public /*out*/ readonly vaultUri!: pulumi.Output<string>;

    /**
     * Create a KeyVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyVaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyVaultArgs | KeyVaultState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyVaultState | undefined;
            inputs["accessPolicies"] = state ? state.accessPolicies : undefined;
            inputs["contacts"] = state ? state.contacts : undefined;
            inputs["enableRbacAuthorization"] = state ? state.enableRbacAuthorization : undefined;
            inputs["enabledForDeployment"] = state ? state.enabledForDeployment : undefined;
            inputs["enabledForDiskEncryption"] = state ? state.enabledForDiskEncryption : undefined;
            inputs["enabledForTemplateDeployment"] = state ? state.enabledForTemplateDeployment : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkAcls"] = state ? state.networkAcls : undefined;
            inputs["purgeProtectionEnabled"] = state ? state.purgeProtectionEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["softDeleteEnabled"] = state ? state.softDeleteEnabled : undefined;
            inputs["softDeleteRetentionDays"] = state ? state.softDeleteRetentionDays : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["vaultUri"] = state ? state.vaultUri : undefined;
        } else {
            const args = argsOrState as KeyVaultArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.skuName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'skuName'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["accessPolicies"] = args ? args.accessPolicies : undefined;
            inputs["contacts"] = args ? args.contacts : undefined;
            inputs["enableRbacAuthorization"] = args ? args.enableRbacAuthorization : undefined;
            inputs["enabledForDeployment"] = args ? args.enabledForDeployment : undefined;
            inputs["enabledForDiskEncryption"] = args ? args.enabledForDiskEncryption : undefined;
            inputs["enabledForTemplateDeployment"] = args ? args.enabledForTemplateDeployment : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkAcls"] = args ? args.networkAcls : undefined;
            inputs["purgeProtectionEnabled"] = args ? args.purgeProtectionEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["softDeleteEnabled"] = args ? args.softDeleteEnabled : undefined;
            inputs["softDeleteRetentionDays"] = args ? args.softDeleteRetentionDays : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["vaultUri"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KeyVault.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyVault resources.
 */
export interface KeyVaultState {
    /**
     * A list of up to 16 objects describing access policies, as described below.
     */
    accessPolicies?: pulumi.Input<pulumi.Input<inputs.keyvault.KeyVaultAccessPolicy>[]>;
    /**
     * One or more `contact` block as defined below.
     */
    contacts?: pulumi.Input<pulumi.Input<inputs.keyvault.KeyVaultContact>[]>;
    /**
     * Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
     */
    enableRbacAuthorization?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
     */
    enabledForDeployment?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
     */
    enabledForDiskEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
     */
    enabledForTemplateDeployment?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Key Vault. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `networkAcls` block as defined below.
     */
    networkAcls?: pulumi.Input<inputs.keyvault.KeyVaultNetworkAcls>;
    /**
     * Is Purge Protection enabled for this Key Vault? Defaults to `false`.
     */
    purgeProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
     */
    skuName?: pulumi.Input<string>;
    /**
     * @deprecated Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
     */
    softDeleteEnabled?: pulumi.Input<boolean>;
    /**
     * The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
     */
    softDeleteRetentionDays?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The URI of the Key Vault, used for performing operations on keys and secrets.
     */
    vaultUri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyVault resource.
 */
export interface KeyVaultArgs {
    /**
     * A list of up to 16 objects describing access policies, as described below.
     */
    accessPolicies?: pulumi.Input<pulumi.Input<inputs.keyvault.KeyVaultAccessPolicy>[]>;
    /**
     * One or more `contact` block as defined below.
     */
    contacts?: pulumi.Input<pulumi.Input<inputs.keyvault.KeyVaultContact>[]>;
    /**
     * Boolean flag to specify whether Azure Key Vault uses Role Based Access Control (RBAC) for authorization of data actions. Defaults to `false`.
     */
    enableRbacAuthorization?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault. Defaults to `false`.
     */
    enabledForDeployment?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys. Defaults to `false`.
     */
    enabledForDiskEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean flag to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault. Defaults to `false`.
     */
    enabledForTemplateDeployment?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Key Vault. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `networkAcls` block as defined below.
     */
    networkAcls?: pulumi.Input<inputs.keyvault.KeyVaultNetworkAcls>;
    /**
     * Is Purge Protection enabled for this Key Vault? Defaults to `false`.
     */
    purgeProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Key Vault. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The Name of the SKU used for this Key Vault. Possible values are `standard` and `premium`.
     */
    skuName: pulumi.Input<string>;
    /**
     * @deprecated Azure has removed support for disabling Soft Delete as of 2020-12-15, as such this field is no longer configurable and can be safely removed. This field will be removed in version 3.0 of the Azure Provider.
     */
    softDeleteEnabled?: pulumi.Input<boolean>;
    /**
     * The number of days that items should be retained for once soft-deleted. This value can be between `7` and `90` (the default) days.
     */
    softDeleteRetentionDays?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
     */
    tenantId: pulumi.Input<string>;
}
