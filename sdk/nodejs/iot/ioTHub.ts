// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IotHub
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West US",
 *     name: "resourceGroup1",
 * });
 * const testAccount = new azure.storage.Account("test", {
 *     accountReplicationType: "LRS",
 *     accountTier: "Standard",
 *     location: testResourceGroup.location,
 *     name: "teststa",
 *     resourceGroupName: testResourceGroup.name,
 * });
 * const testIoTHub = new azure.iot.IoTHub("test", {
 *     endpoints: [{
 *         batchFrequencyInSeconds: 60,
 *         connectionString: testAccount.primaryBlobConnectionString,
 *         containerName: "test",
 *         encoding: "Avro",
 *         fileNameFormat: "{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}",
 *         maxChunkSizeInBytes: 10485760,
 *         name: "export",
 *         type: "AzureIotHub.StorageContainer",
 *     }],
 *     fallbackRoute: {
 *         enabled: true,
 *     },
 *     location: testResourceGroup.location,
 *     name: "test",
 *     resourceGroupName: testResourceGroup.name,
 *     routes: [{
 *         condition: "true",
 *         enabled: true,
 *         endpointNames: ["export"],
 *         name: "export",
 *         source: "DeviceMessages",
 *     }],
 *     sku: {
 *         capacity: 1,
 *         name: "S1",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         purpose: "testing",
 *     },
 * });
 * const testContainer = new azure.storage.Container("test", {
 *     containerAccessType: "private",
 *     name: "test",
 *     resourceGroupName: testResourceGroup.name,
 *     storageAccountName: testAccount.name,
 * });
 * ```
 */
export class IoTHub extends pulumi.CustomResource {
    /**
     * Get an existing IoTHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IoTHubState, opts?: pulumi.CustomResourceOptions): IoTHub {
        return new IoTHub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:iot/ioTHub:IoTHub';

    /**
     * Returns true if the given object is an instance of IoTHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IoTHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IoTHub.__pulumiType;
    }

    /**
     * An `endpoint` block as defined below.
     */
    public readonly endpoints!: pulumi.Output<{ batchFrequencyInSeconds?: number, connectionString: string, containerName?: string, encoding?: string, fileNameFormat?: string, maxChunkSizeInBytes?: number, name: string, type: string }[] | undefined>;
    /**
     * The EventHub compatible endpoint for events data
     */
    public /*out*/ readonly eventHubEventsEndpoint!: pulumi.Output<string>;
    /**
     * The EventHub compatible path for events data
     */
    public /*out*/ readonly eventHubEventsPath!: pulumi.Output<string>;
    /**
     * The EventHub compatible endpoint for operational data
     */
    public /*out*/ readonly eventHubOperationsEndpoint!: pulumi.Output<string>;
    /**
     * The EventHub compatible path for operational data
     */
    public /*out*/ readonly eventHubOperationsPath!: pulumi.Output<string>;
    /**
     * A `fallback_route` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
     */
    public readonly fallbackRoute!: pulumi.Output<{ condition?: string, enabled: boolean, endpointNames: string[], source?: string }>;
    /**
     * The hostname of the IotHub Resource.
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * One or more `ip_filter_rule` blocks as defined below.
     */
    public readonly ipFilterRules!: pulumi.Output<{ action: string, ipMask: string, name: string }[] | undefined>;
    /**
     * Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `route` block as defined below.
     */
    public readonly routes!: pulumi.Output<{ condition?: string, enabled: boolean, endpointNames: string[], name: string, source: string }[] | undefined>;
    /**
     * One or more `shared_access_policy` blocks as defined below.
     */
    public /*out*/ readonly sharedAccessPolicies!: pulumi.Output<{ keyName: string, permissions: string, primaryKey: string, secondaryKey: string }[]>;
    /**
     * A `sku` block as defined below.
     */
    public readonly sku!: pulumi.Output<{ capacity: number, name: string, tier: string }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IoTHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IoTHubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IoTHubArgs | IoTHubState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IoTHubState | undefined;
            inputs["endpoints"] = state ? state.endpoints : undefined;
            inputs["eventHubEventsEndpoint"] = state ? state.eventHubEventsEndpoint : undefined;
            inputs["eventHubEventsPath"] = state ? state.eventHubEventsPath : undefined;
            inputs["eventHubOperationsEndpoint"] = state ? state.eventHubOperationsEndpoint : undefined;
            inputs["eventHubOperationsPath"] = state ? state.eventHubOperationsPath : undefined;
            inputs["fallbackRoute"] = state ? state.fallbackRoute : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["ipFilterRules"] = state ? state.ipFilterRules : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["routes"] = state ? state.routes : undefined;
            inputs["sharedAccessPolicies"] = state ? state.sharedAccessPolicies : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as IoTHubArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["endpoints"] = args ? args.endpoints : undefined;
            inputs["fallbackRoute"] = args ? args.fallbackRoute : undefined;
            inputs["ipFilterRules"] = args ? args.ipFilterRules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routes"] = args ? args.routes : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["eventHubEventsEndpoint"] = undefined /*out*/;
            inputs["eventHubEventsPath"] = undefined /*out*/;
            inputs["eventHubOperationsEndpoint"] = undefined /*out*/;
            inputs["eventHubOperationsPath"] = undefined /*out*/;
            inputs["hostname"] = undefined /*out*/;
            inputs["sharedAccessPolicies"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        super(IoTHub.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IoTHub resources.
 */
export interface IoTHubState {
    /**
     * An `endpoint` block as defined below.
     */
    readonly endpoints?: pulumi.Input<pulumi.Input<{ batchFrequencyInSeconds?: pulumi.Input<number>, connectionString: pulumi.Input<string>, containerName?: pulumi.Input<string>, encoding?: pulumi.Input<string>, fileNameFormat?: pulumi.Input<string>, maxChunkSizeInBytes?: pulumi.Input<number>, name: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * The EventHub compatible endpoint for events data
     */
    readonly eventHubEventsEndpoint?: pulumi.Input<string>;
    /**
     * The EventHub compatible path for events data
     */
    readonly eventHubEventsPath?: pulumi.Input<string>;
    /**
     * The EventHub compatible endpoint for operational data
     */
    readonly eventHubOperationsEndpoint?: pulumi.Input<string>;
    /**
     * The EventHub compatible path for operational data
     */
    readonly eventHubOperationsPath?: pulumi.Input<string>;
    /**
     * A `fallback_route` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
     */
    readonly fallbackRoute?: pulumi.Input<{ condition?: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, endpointNames?: pulumi.Input<pulumi.Input<string>[]>, source?: pulumi.Input<string> }>;
    /**
     * The hostname of the IotHub Resource.
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * One or more `ip_filter_rule` blocks as defined below.
     */
    readonly ipFilterRules?: pulumi.Input<pulumi.Input<{ action: pulumi.Input<string>, ipMask: pulumi.Input<string>, name: pulumi.Input<string> }>[]>;
    /**
     * Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `route` block as defined below.
     */
    readonly routes?: pulumi.Input<pulumi.Input<{ condition?: pulumi.Input<string>, enabled: pulumi.Input<boolean>, endpointNames: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, source: pulumi.Input<string> }>[]>;
    /**
     * One or more `shared_access_policy` blocks as defined below.
     */
    readonly sharedAccessPolicies?: pulumi.Input<pulumi.Input<{ keyName?: pulumi.Input<string>, permissions?: pulumi.Input<string>, primaryKey?: pulumi.Input<string>, secondaryKey?: pulumi.Input<string> }>[]>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku?: pulumi.Input<{ capacity: pulumi.Input<number>, name: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IoTHub resource.
 */
export interface IoTHubArgs {
    /**
     * An `endpoint` block as defined below.
     */
    readonly endpoints?: pulumi.Input<pulumi.Input<{ batchFrequencyInSeconds?: pulumi.Input<number>, connectionString: pulumi.Input<string>, containerName?: pulumi.Input<string>, encoding?: pulumi.Input<string>, fileNameFormat?: pulumi.Input<string>, maxChunkSizeInBytes?: pulumi.Input<number>, name: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * A `fallback_route` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
     */
    readonly fallbackRoute?: pulumi.Input<{ condition?: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, endpointNames?: pulumi.Input<pulumi.Input<string>[]>, source?: pulumi.Input<string> }>;
    /**
     * One or more `ip_filter_rule` blocks as defined below.
     */
    readonly ipFilterRules?: pulumi.Input<pulumi.Input<{ action: pulumi.Input<string>, ipMask: pulumi.Input<string>, name: pulumi.Input<string> }>[]>;
    /**
     * Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `route` block as defined below.
     */
    readonly routes?: pulumi.Input<pulumi.Input<{ condition?: pulumi.Input<string>, enabled: pulumi.Input<boolean>, endpointNames: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, source: pulumi.Input<string> }>[]>;
    /**
     * A `sku` block as defined below.
     */
    readonly sku: pulumi.Input<{ capacity: pulumi.Input<number>, name: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
