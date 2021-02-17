// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Stream Analytics Stream Input IoTHub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = azure.core.getResourceGroup({
 *     name: "example-resources",
 * });
 * const exampleJob = azure.streamanalytics.getJob({
 *     name: "example-job",
 *     resourceGroupName: azurerm_resource_group.example.name,
 * });
 * const exampleIoTHub = new azure.iot.IoTHub("exampleIoTHub", {
 *     resourceGroupName: azurerm_resource_group.example.name,
 *     location: azurerm_resource_group.example.location,
 *     sku: {
 *         name: "S1",
 *         capacity: "1",
 *     },
 * });
 * const exampleStreamInputIotHub = new azure.streamanalytics.StreamInputIotHub("exampleStreamInputIotHub", {
 *     streamAnalyticsJobName: exampleJob.then(exampleJob => exampleJob.name),
 *     resourceGroupName: exampleJob.then(exampleJob => exampleJob.resourceGroupName),
 *     endpoint: "messages/events",
 *     eventhubConsumerGroupName: `$Default`,
 *     iothubNamespace: exampleIoTHub.name,
 *     sharedAccessPolicyKey: exampleIoTHub.sharedAccessPolicies.apply(sharedAccessPolicies => sharedAccessPolicies[0].primaryKey),
 *     sharedAccessPolicyName: "iothubowner",
 *     serialization: {
 *         type: "Json",
 *         encoding: "UTF8",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Stream Analytics Stream Input IoTHub's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:streamanalytics/streamInputIotHub:StreamInputIotHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
 * ```
 */
export class StreamInputIotHub extends pulumi.CustomResource {
    /**
     * Get an existing StreamInputIotHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamInputIotHubState, opts?: pulumi.CustomResourceOptions): StreamInputIotHub {
        return new StreamInputIotHub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:streamanalytics/streamInputIotHub:StreamInputIotHub';

    /**
     * Returns true if the given object is an instance of StreamInputIotHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamInputIotHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamInputIotHub.__pulumiType;
    }

    /**
     * The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
     */
    public readonly eventhubConsumerGroupName!: pulumi.Output<string>;
    /**
     * The name or the URI of the IoT Hub.
     */
    public readonly iothubNamespace!: pulumi.Output<string>;
    /**
     * The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `serialization` block as defined below.
     */
    public readonly serialization!: pulumi.Output<outputs.streamanalytics.StreamInputIotHubSerialization>;
    /**
     * The shared access policy key for the specified shared access policy.
     */
    public readonly sharedAccessPolicyKey!: pulumi.Output<string>;
    /**
     * The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    public readonly sharedAccessPolicyName!: pulumi.Output<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    public readonly streamAnalyticsJobName!: pulumi.Output<string>;

    /**
     * Create a StreamInputIotHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamInputIotHubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamInputIotHubArgs | StreamInputIotHubState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamInputIotHubState | undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["eventhubConsumerGroupName"] = state ? state.eventhubConsumerGroupName : undefined;
            inputs["iothubNamespace"] = state ? state.iothubNamespace : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serialization"] = state ? state.serialization : undefined;
            inputs["sharedAccessPolicyKey"] = state ? state.sharedAccessPolicyKey : undefined;
            inputs["sharedAccessPolicyName"] = state ? state.sharedAccessPolicyName : undefined;
            inputs["streamAnalyticsJobName"] = state ? state.streamAnalyticsJobName : undefined;
        } else {
            const args = argsOrState as StreamInputIotHubArgs | undefined;
            if ((!args || args.endpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoint'");
            }
            if ((!args || args.eventhubConsumerGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventhubConsumerGroupName'");
            }
            if ((!args || args.iothubNamespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iothubNamespace'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serialization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serialization'");
            }
            if ((!args || args.sharedAccessPolicyKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sharedAccessPolicyKey'");
            }
            if ((!args || args.sharedAccessPolicyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sharedAccessPolicyName'");
            }
            if ((!args || args.streamAnalyticsJobName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamAnalyticsJobName'");
            }
            inputs["endpoint"] = args ? args.endpoint : undefined;
            inputs["eventhubConsumerGroupName"] = args ? args.eventhubConsumerGroupName : undefined;
            inputs["iothubNamespace"] = args ? args.iothubNamespace : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serialization"] = args ? args.serialization : undefined;
            inputs["sharedAccessPolicyKey"] = args ? args.sharedAccessPolicyKey : undefined;
            inputs["sharedAccessPolicyName"] = args ? args.sharedAccessPolicyName : undefined;
            inputs["streamAnalyticsJobName"] = args ? args.streamAnalyticsJobName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StreamInputIotHub.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StreamInputIotHub resources.
 */
export interface StreamInputIotHubState {
    /**
     * The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
     */
    readonly eventhubConsumerGroupName?: pulumi.Input<string>;
    /**
     * The name or the URI of the IoT Hub.
     */
    readonly iothubNamespace?: pulumi.Input<string>;
    /**
     * The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `serialization` block as defined below.
     */
    readonly serialization?: pulumi.Input<inputs.streamanalytics.StreamInputIotHubSerialization>;
    /**
     * The shared access policy key for the specified shared access policy.
     */
    readonly sharedAccessPolicyKey?: pulumi.Input<string>;
    /**
     * The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    readonly sharedAccessPolicyName?: pulumi.Input<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    readonly streamAnalyticsJobName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StreamInputIotHub resource.
 */
export interface StreamInputIotHubArgs {
    /**
     * The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
     */
    readonly endpoint: pulumi.Input<string>;
    /**
     * The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
     */
    readonly eventhubConsumerGroupName: pulumi.Input<string>;
    /**
     * The name or the URI of the IoT Hub.
     */
    readonly iothubNamespace: pulumi.Input<string>;
    /**
     * The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `serialization` block as defined below.
     */
    readonly serialization: pulumi.Input<inputs.streamanalytics.StreamInputIotHubSerialization>;
    /**
     * The shared access policy key for the specified shared access policy.
     */
    readonly sharedAccessPolicyKey: pulumi.Input<string>;
    /**
     * The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    readonly sharedAccessPolicyName: pulumi.Input<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    readonly streamAnalyticsJobName: pulumi.Input<string>;
}
