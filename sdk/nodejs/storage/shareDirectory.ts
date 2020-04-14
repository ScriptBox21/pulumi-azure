// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Directory within an Azure Storage File Share.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_share_directory.html.markdown.
 */
export class ShareDirectory extends pulumi.CustomResource {
    /**
     * Get an existing ShareDirectory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShareDirectoryState, opts?: pulumi.CustomResourceOptions): ShareDirectory {
        return new ShareDirectory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/shareDirectory:ShareDirectory';

    /**
     * Returns true if the given object is an instance of ShareDirectory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShareDirectory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShareDirectory.__pulumiType;
    }

    /**
     * A mapping of metadata to assign to this Directory.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
     */
    public readonly shareName!: pulumi.Output<string>;
    /**
     * The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
     */
    public readonly storageAccountName!: pulumi.Output<string>;

    /**
     * Create a ShareDirectory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShareDirectoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShareDirectoryArgs | ShareDirectoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ShareDirectoryState | undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["shareName"] = state ? state.shareName : undefined;
            inputs["storageAccountName"] = state ? state.storageAccountName : undefined;
        } else {
            const args = argsOrState as ShareDirectoryArgs | undefined;
            if (!args || args.shareName === undefined) {
                throw new Error("Missing required property 'shareName'");
            }
            if (!args || args.storageAccountName === undefined) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["shareName"] = args ? args.shareName : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ShareDirectory.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShareDirectory resources.
 */
export interface ShareDirectoryState {
    /**
     * A mapping of metadata to assign to this Directory.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
     */
    readonly shareName?: pulumi.Input<string>;
    /**
     * The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
     */
    readonly storageAccountName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ShareDirectory resource.
 */
export interface ShareDirectoryArgs {
    /**
     * A mapping of metadata to assign to this Directory.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
     */
    readonly shareName: pulumi.Input<string>;
    /**
     * The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
     */
    readonly storageAccountName: pulumi.Input<string>;
}
