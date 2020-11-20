// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Data Share Data Lake Gen2 Dataset.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as azuread from "@pulumi/azuread";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.datashare.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleShare = new azure.datashare.Share("exampleShare", {
 *     accountId: exampleAccount.id,
 *     kind: "CopyBased",
 * });
 * const exampleStorage_accountAccount = new azure.storage.Account("exampleStorage/accountAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountKind: "BlobStorage",
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleDataLakeGen2Filesystem = new azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", {storageAccountId: exampleStorage / accountAccount.id});
 * const exampleServicePrincipal = exampleAccount.name.apply(name => azuread.getServicePrincipal({
 *     displayName: name,
 * }));
 * const exampleAssignment = new azure.authorization.Assignment("exampleAssignment", {
 *     scope: exampleStorage / accountAccount.id,
 *     roleDefinitionName: "Storage Blob Data Reader",
 *     principalId: exampleServicePrincipal.objectId,
 * });
 * const exampleDatasetDataLakeGen2 = new azure.datashare.DatasetDataLakeGen2("exampleDatasetDataLakeGen2", {
 *     shareId: exampleShare.id,
 *     storageAccountId: exampleStorage / accountAccount.id,
 *     fileSystemName: exampleDataLakeGen2Filesystem.name,
 *     filePath: "myfile.txt",
 * }, {
 *     dependsOn: [exampleAssignment],
 * });
 * ```
 *
 * ## Import
 *
 * Data Share Data Lake Gen2 Datasets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
 * ```
 */
export class DatasetDataLakeGen2 extends pulumi.CustomResource {
    /**
     * Get an existing DatasetDataLakeGen2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetDataLakeGen2State, opts?: pulumi.CustomResourceOptions): DatasetDataLakeGen2 {
        return new DatasetDataLakeGen2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2';

    /**
     * Returns true if the given object is an instance of DatasetDataLakeGen2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetDataLakeGen2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetDataLakeGen2.__pulumiType;
    }

    /**
     * The name of the Data Share Dataset.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    public readonly filePath!: pulumi.Output<string | undefined>;
    /**
     * The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    public readonly fileSystemName!: pulumi.Output<string>;
    /**
     * The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    public readonly folderPath!: pulumi.Output<string | undefined>;
    /**
     * The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    public readonly shareId!: pulumi.Output<string>;
    /**
     * The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string>;

    /**
     * Create a DatasetDataLakeGen2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetDataLakeGen2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetDataLakeGen2Args | DatasetDataLakeGen2State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatasetDataLakeGen2State | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["filePath"] = state ? state.filePath : undefined;
            inputs["fileSystemName"] = state ? state.fileSystemName : undefined;
            inputs["folderPath"] = state ? state.folderPath : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["shareId"] = state ? state.shareId : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
        } else {
            const args = argsOrState as DatasetDataLakeGen2Args | undefined;
            if (!args || args.fileSystemName === undefined) {
                throw new Error("Missing required property 'fileSystemName'");
            }
            if (!args || args.shareId === undefined) {
                throw new Error("Missing required property 'shareId'");
            }
            if (!args || args.storageAccountId === undefined) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            inputs["filePath"] = args ? args.filePath : undefined;
            inputs["fileSystemName"] = args ? args.fileSystemName : undefined;
            inputs["folderPath"] = args ? args.folderPath : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["shareId"] = args ? args.shareId : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["displayName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatasetDataLakeGen2.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatasetDataLakeGen2 resources.
 */
export interface DatasetDataLakeGen2State {
    /**
     * The name of the Data Share Dataset.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly filePath?: pulumi.Input<string>;
    /**
     * The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly fileSystemName?: pulumi.Input<string>;
    /**
     * The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly folderPath?: pulumi.Input<string>;
    /**
     * The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly shareId?: pulumi.Input<string>;
    /**
     * The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly storageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatasetDataLakeGen2 resource.
 */
export interface DatasetDataLakeGen2Args {
    /**
     * The path of the file in the data lake file system to be shared with the receiver. Conflicts with `folderPath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly filePath?: pulumi.Input<string>;
    /**
     * The name of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly fileSystemName: pulumi.Input<string>;
    /**
     * The folder path in the data lake file system to be shared with the receiver. Conflicts with `filePath` Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly folderPath?: pulumi.Input<string>;
    /**
     * The name which should be used for this Data Share Data Lake Gen2 Dataset. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen2 Dataset should be created. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly shareId: pulumi.Input<string>;
    /**
     * The resource id of the storage account of the data lake file system to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen2 Dataset to be created.
     */
    readonly storageAccountId: pulumi.Input<string>;
}
