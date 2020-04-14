// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Disk Encryption Set.
 * 
 * > **NOTE**: Disk Encryption Sets are in Public Preview and at this time is only available in `Canada Central`, `North Europe` and `West Central US` regions - [more information can be found in the preview documentation](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/disk-encryption).
 * 
 * > **NOTE:** At this time the Key Vault used to store the Active Key for this Disk Encryption Set must have both Soft Delete & Purge Protection enabled - which are not yet supported by this provider.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/disk_encryption_set.html.markdown.
 */
export class DiskEncryptionSet extends pulumi.CustomResource {
    /**
     * Get an existing DiskEncryptionSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskEncryptionSetState, opts?: pulumi.CustomResourceOptions): DiskEncryptionSet {
        return new DiskEncryptionSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/diskEncryptionSet:DiskEncryptionSet';

    /**
     * Returns true if the given object is an instance of DiskEncryptionSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiskEncryptionSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiskEncryptionSet.__pulumiType;
    }

    /**
     * A `identity` block defined below.
     */
    public readonly identity!: pulumi.Output<outputs.compute.DiskEncryptionSetIdentity>;
    /**
     * Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
     */
    public readonly keyVaultKeyId!: pulumi.Output<string>;
    /**
     * Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Disk Encryption Set. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the Disk Encryption Set.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a DiskEncryptionSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskEncryptionSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskEncryptionSetArgs | DiskEncryptionSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DiskEncryptionSetState | undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["keyVaultKeyId"] = state ? state.keyVaultKeyId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DiskEncryptionSetArgs | undefined;
            if (!args || args.identity === undefined) {
                throw new Error("Missing required property 'identity'");
            }
            if (!args || args.keyVaultKeyId === undefined) {
                throw new Error("Missing required property 'keyVaultKeyId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["keyVaultKeyId"] = args ? args.keyVaultKeyId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DiskEncryptionSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiskEncryptionSet resources.
 */
export interface DiskEncryptionSetState {
    /**
     * A `identity` block defined below.
     */
    readonly identity?: pulumi.Input<inputs.compute.DiskEncryptionSetIdentity>;
    /**
     * Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
     */
    readonly keyVaultKeyId?: pulumi.Input<string>;
    /**
     * Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Disk Encryption Set. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Disk Encryption Set.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DiskEncryptionSet resource.
 */
export interface DiskEncryptionSetArgs {
    /**
     * A `identity` block defined below.
     */
    readonly identity: pulumi.Input<inputs.compute.DiskEncryptionSetIdentity>;
    /**
     * Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
     */
    readonly keyVaultKeyId: pulumi.Input<string>;
    /**
     * Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Disk Encryption Set. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Disk Encryption Set.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
