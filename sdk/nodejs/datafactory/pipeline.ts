// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Pipeline inside a Azure Data Factory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const examplePipeline = new azure.datafactory.Pipeline("examplePipeline", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 * });
 * ```
 * ### With Activities
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const test = new azure.datafactory.Pipeline("test", {
 *     resourceGroupName: azurerm_resource_group.test.name,
 *     dataFactoryName: azurerm_data_factory.test.name,
 *     variables: {
 *         bob: "item1",
 *     },
 *     activitiesJson: `[
 * 	{
 * 		"name": "Append variable1",
 * 		"type": "AppendVariable",
 * 		"dependsOn": [],
 * 		"userProperties": [],
 * 		"typeProperties": {
 * 			"variableName": "bob",
 * 			"value": "something"
 * 		}
 * 	}
 * ]
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Pipeline's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/pipeline:Pipeline example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/pipelines/example
 * ```
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineState, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/pipeline:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
     */
    public readonly activitiesJson!: pulumi.Output<string | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Pipeline.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Pipeline.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The folder that this Pipeline is in. If not specified, the Pipeline will appear at the root level.
     */
    public readonly folder!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of parameters to associate with the Data Factory Pipeline.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A map of variables to associate with the Data Factory Pipeline.
     */
    public readonly variables!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs | PipelineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineState | undefined;
            inputs["activitiesJson"] = state ? state.activitiesJson : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as PipelineArgs | undefined;
            if ((!args || args.dataFactoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["activitiesJson"] = args ? args.activitiesJson : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["variables"] = args ? args.variables : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Pipeline.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pipeline resources.
 */
export interface PipelineState {
    /**
     * A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
     */
    readonly activitiesJson?: pulumi.Input<string>;
    /**
     * List of tags that can be used for describing the Data Factory Pipeline.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
     */
    readonly dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Pipeline.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The folder that this Pipeline is in. If not specified, the Pipeline will appear at the root level.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Pipeline.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A map of variables to associate with the Data Factory Pipeline.
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * A JSON object that contains the activities that will be associated with the Data Factory Pipeline.
     */
    readonly activitiesJson?: pulumi.Input<string>;
    /**
     * List of tags that can be used for describing the Data Factory Pipeline.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Pipeline with. Changing this forces a new resource.
     */
    readonly dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Pipeline.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The folder that this Pipeline is in. If not specified, the Pipeline will appear at the root level.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Pipeline. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Pipeline.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Pipeline. Changing this forces a new resource
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A map of variables to associate with the Data Factory Pipeline.
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
