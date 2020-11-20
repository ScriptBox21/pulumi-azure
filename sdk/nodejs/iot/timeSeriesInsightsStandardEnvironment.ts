// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure IoT Time Series Insights Standard Environment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "northeurope"});
 * const exampleTimeSeriesInsightsStandardEnvironment = new azure.iot.TimeSeriesInsightsStandardEnvironment("exampleTimeSeriesInsightsStandardEnvironment", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "S1_1",
 *     dataRetentionTime: "P30D",
 * });
 * ```
 *
 * ## Import
 *
 * Azure IoT Time Series Insights Standard Environment can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
 * ```
 */
export class TimeSeriesInsightsStandardEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing TimeSeriesInsightsStandardEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TimeSeriesInsightsStandardEnvironmentState, opts?: pulumi.CustomResourceOptions): TimeSeriesInsightsStandardEnvironment {
        return new TimeSeriesInsightsStandardEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment';

    /**
     * Returns true if the given object is an instance of TimeSeriesInsightsStandardEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TimeSeriesInsightsStandardEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TimeSeriesInsightsStandardEnvironment.__pulumiType;
    }

    /**
     * Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
     */
    public readonly dataRetentionTime!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
     */
    public readonly partitionKey!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
     */
    public readonly storageLimitExceededBehavior!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a TimeSeriesInsightsStandardEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TimeSeriesInsightsStandardEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TimeSeriesInsightsStandardEnvironmentArgs | TimeSeriesInsightsStandardEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TimeSeriesInsightsStandardEnvironmentState | undefined;
            inputs["dataRetentionTime"] = state ? state.dataRetentionTime : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partitionKey"] = state ? state.partitionKey : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["storageLimitExceededBehavior"] = state ? state.storageLimitExceededBehavior : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as TimeSeriesInsightsStandardEnvironmentArgs | undefined;
            if (!args || args.dataRetentionTime === undefined) {
                throw new Error("Missing required property 'dataRetentionTime'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.skuName === undefined) {
                throw new Error("Missing required property 'skuName'");
            }
            inputs["dataRetentionTime"] = args ? args.dataRetentionTime : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partitionKey"] = args ? args.partitionKey : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["storageLimitExceededBehavior"] = args ? args.storageLimitExceededBehavior : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TimeSeriesInsightsStandardEnvironment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TimeSeriesInsightsStandardEnvironment resources.
 */
export interface TimeSeriesInsightsStandardEnvironmentState {
    /**
     * Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
     */
    readonly dataRetentionTime?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
     */
    readonly partitionKey?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
     */
    readonly storageLimitExceededBehavior?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a TimeSeriesInsightsStandardEnvironment resource.
 */
export interface TimeSeriesInsightsStandardEnvironmentArgs {
    /**
     * Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
     */
    readonly dataRetentionTime: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created. Must be globally unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
     */
    readonly partitionKey?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(\_).The fist part is the `name`, valid values include: `S1` and `S2`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `S1_1`). Changing this forces a new resource to be created.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include `PauseIngress` and `PurgeOldData`. Defaults to `PurgeOldData`.
     */
    readonly storageLimitExceededBehavior?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
