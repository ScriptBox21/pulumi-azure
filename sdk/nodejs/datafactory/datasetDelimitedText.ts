// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure Delimited Text Dataset inside an Azure Data Factory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "northeurope"});
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleLinkedServiceWeb = new azure.datafactory.LinkedServiceWeb("exampleLinkedServiceWeb", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     authenticationType: "Anonymous",
 *     url: "https://www.bing.com",
 * });
 * const exampleDatasetDelimitedText = new azure.datafactory.DatasetDelimitedText("exampleDatasetDelimitedText", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     linkedServiceName: exampleLinkedServiceWeb.name,
 *     httpServerLocation: {
 *         relativeUrl: "http://www.bing.com",
 *         path: "foo/bar/",
 *         filename: "fizz.txt",
 *     },
 *     columnDelimiter: ",",
 *     rowDelimiter: "NEW",
 *     encoding: "UTF-8",
 *     quoteCharacter: "x",
 *     escapeCharacter: "f",
 *     firstRowAsHeader: true,
 *     nullValue: "NULL",
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Datasets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/datasetDelimitedText:DatasetDelimitedText example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
 * ```
 */
export class DatasetDelimitedText extends pulumi.CustomResource {
    /**
     * Get an existing DatasetDelimitedText resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetDelimitedTextState, opts?: pulumi.CustomResourceOptions): DatasetDelimitedText {
        return new DatasetDelimitedText(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/datasetDelimitedText:DatasetDelimitedText';

    /**
     * Returns true if the given object is an instance of DatasetDelimitedText.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetDelimitedText {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetDelimitedText.__pulumiType;
    }

    /**
     * A map of additional properties to associate with the Data Factory Dataset.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Dataset.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * A `azureBlobStorageLocation` block as defined below.
     */
    public readonly azureBlobStorageLocation!: pulumi.Output<outputs.datafactory.DatasetDelimitedTextAzureBlobStorageLocation | undefined>;
    /**
     * The column delimiter.
     */
    public readonly columnDelimiter!: pulumi.Output<string | undefined>;
    /**
     * The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Dataset.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The encoding format for the file.
     */
    public readonly encoding!: pulumi.Output<string | undefined>;
    /**
     * The escape character.
     */
    public readonly escapeCharacter!: pulumi.Output<string | undefined>;
    /**
     * When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
     */
    public readonly firstRowAsHeader!: pulumi.Output<boolean | undefined>;
    /**
     * The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
     */
    public readonly folder!: pulumi.Output<string | undefined>;
    /**
     * A `httpServerLocation` block as defined below.
     */
    public readonly httpServerLocation!: pulumi.Output<outputs.datafactory.DatasetDelimitedTextHttpServerLocation | undefined>;
    /**
     * The Data Factory Linked Service name in which to associate the Dataset with.
     */
    public readonly linkedServiceName!: pulumi.Output<string>;
    /**
     * Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The null value string.
     */
    public readonly nullValue!: pulumi.Output<string | undefined>;
    /**
     * A map of parameters to associate with the Data Factory Dataset.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The quote character.
     */
    public readonly quoteCharacter!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The row delimiter.
     */
    public readonly rowDelimiter!: pulumi.Output<string | undefined>;
    /**
     * A `schemaColumn` block as defined below.
     */
    public readonly schemaColumns!: pulumi.Output<outputs.datafactory.DatasetDelimitedTextSchemaColumn[] | undefined>;

    /**
     * Create a DatasetDelimitedText resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetDelimitedTextArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetDelimitedTextArgs | DatasetDelimitedTextState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatasetDelimitedTextState | undefined;
            inputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["azureBlobStorageLocation"] = state ? state.azureBlobStorageLocation : undefined;
            inputs["columnDelimiter"] = state ? state.columnDelimiter : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encoding"] = state ? state.encoding : undefined;
            inputs["escapeCharacter"] = state ? state.escapeCharacter : undefined;
            inputs["firstRowAsHeader"] = state ? state.firstRowAsHeader : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["httpServerLocation"] = state ? state.httpServerLocation : undefined;
            inputs["linkedServiceName"] = state ? state.linkedServiceName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nullValue"] = state ? state.nullValue : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["quoteCharacter"] = state ? state.quoteCharacter : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["rowDelimiter"] = state ? state.rowDelimiter : undefined;
            inputs["schemaColumns"] = state ? state.schemaColumns : undefined;
        } else {
            const args = argsOrState as DatasetDelimitedTextArgs | undefined;
            if (!args || args.dataFactoryName === undefined) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if (!args || args.linkedServiceName === undefined) {
                throw new Error("Missing required property 'linkedServiceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["azureBlobStorageLocation"] = args ? args.azureBlobStorageLocation : undefined;
            inputs["columnDelimiter"] = args ? args.columnDelimiter : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encoding"] = args ? args.encoding : undefined;
            inputs["escapeCharacter"] = args ? args.escapeCharacter : undefined;
            inputs["firstRowAsHeader"] = args ? args.firstRowAsHeader : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["httpServerLocation"] = args ? args.httpServerLocation : undefined;
            inputs["linkedServiceName"] = args ? args.linkedServiceName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nullValue"] = args ? args.nullValue : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["quoteCharacter"] = args ? args.quoteCharacter : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["rowDelimiter"] = args ? args.rowDelimiter : undefined;
            inputs["schemaColumns"] = args ? args.schemaColumns : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatasetDelimitedText.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatasetDelimitedText resources.
 */
export interface DatasetDelimitedTextState {
    /**
     * A map of additional properties to associate with the Data Factory Dataset.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Dataset.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `azureBlobStorageLocation` block as defined below.
     */
    readonly azureBlobStorageLocation?: pulumi.Input<inputs.datafactory.DatasetDelimitedTextAzureBlobStorageLocation>;
    /**
     * The column delimiter.
     */
    readonly columnDelimiter?: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
     */
    readonly dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Dataset.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The encoding format for the file.
     */
    readonly encoding?: pulumi.Input<string>;
    /**
     * The escape character.
     */
    readonly escapeCharacter?: pulumi.Input<string>;
    /**
     * When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
     */
    readonly firstRowAsHeader?: pulumi.Input<boolean>;
    /**
     * The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * A `httpServerLocation` block as defined below.
     */
    readonly httpServerLocation?: pulumi.Input<inputs.datafactory.DatasetDelimitedTextHttpServerLocation>;
    /**
     * The Data Factory Linked Service name in which to associate the Dataset with.
     */
    readonly linkedServiceName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The null value string.
     */
    readonly nullValue?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Dataset.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The quote character.
     */
    readonly quoteCharacter?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The row delimiter.
     */
    readonly rowDelimiter?: pulumi.Input<string>;
    /**
     * A `schemaColumn` block as defined below.
     */
    readonly schemaColumns?: pulumi.Input<pulumi.Input<inputs.datafactory.DatasetDelimitedTextSchemaColumn>[]>;
}

/**
 * The set of arguments for constructing a DatasetDelimitedText resource.
 */
export interface DatasetDelimitedTextArgs {
    /**
     * A map of additional properties to associate with the Data Factory Dataset.
     */
    readonly additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Dataset.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `azureBlobStorageLocation` block as defined below.
     */
    readonly azureBlobStorageLocation?: pulumi.Input<inputs.datafactory.DatasetDelimitedTextAzureBlobStorageLocation>;
    /**
     * The column delimiter.
     */
    readonly columnDelimiter?: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
     */
    readonly dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Dataset.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The encoding format for the file.
     */
    readonly encoding?: pulumi.Input<string>;
    /**
     * The escape character.
     */
    readonly escapeCharacter?: pulumi.Input<string>;
    /**
     * When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data.
     */
    readonly firstRowAsHeader?: pulumi.Input<boolean>;
    /**
     * The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * A `httpServerLocation` block as defined below.
     */
    readonly httpServerLocation?: pulumi.Input<inputs.datafactory.DatasetDelimitedTextHttpServerLocation>;
    /**
     * The Data Factory Linked Service name in which to associate the Dataset with.
     */
    readonly linkedServiceName: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The null value string.
     */
    readonly nullValue?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Dataset.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The quote character.
     */
    readonly quoteCharacter?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The row delimiter.
     */
    readonly rowDelimiter?: pulumi.Input<string>;
    /**
     * A `schemaColumn` block as defined below.
     */
    readonly schemaColumns?: pulumi.Input<pulumi.Input<inputs.datafactory.DatasetDelimitedTextSchemaColumn>[]>;
}
