// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages Security Center Automation and Continuous Export. This resource supports three types of destination in the `action`, Logic Apps, Log Analytics and Event Hubs
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "westeurope"});
 * const exampleEventHubNamespace = new azure.eventhub.EventHubNamespace("exampleEventHubNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 *     capacity: 2,
 * });
 * const exampleAutomation = new azure.securitycenter.Automation("exampleAutomation", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     actions: [{
 *         type: "EventHub",
 *         resourceId: exampleEventHubNamespace.id,
 *         connectionString: exampleEventHubNamespace.defaultPrimaryConnectionString,
 *     }],
 *     sources: [{
 *         eventSource: "Alerts",
 *         ruleSets: [{
 *             rules: [{
 *                 propertyPath: "properties.metadata.severity",
 *                 operator: "Equals",
 *                 expectedValue: "High",
 *                 propertyType: "String",
 *             }],
 *         }],
 *     }],
 *     scopes: [current.then(current => `/subscriptions/${current.subscriptionId}`)],
 * });
 * ```
 *
 * ## Import
 *
 * Security Center Automations can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:securitycenter/automation:Automation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Security/automations/automation1
 * ```
 */
export class Automation extends pulumi.CustomResource {
    /**
     * Get an existing Automation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutomationState, opts?: pulumi.CustomResourceOptions): Automation {
        return new Automation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:securitycenter/automation:Automation';

    /**
     * Returns true if the given object is an instance of Automation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Automation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Automation.__pulumiType;
    }

    /**
     * One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
     */
    public readonly actions!: pulumi.Output<outputs.securitycenter.AutomationAction[]>;
    /**
     * Specifies the description for the Security Center Automation.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Boolean to enable or disable this Security Center Automation.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
     */
    public readonly sources!: pulumi.Output<outputs.securitycenter.AutomationSource[]>;
    /**
     * A mapping of tags assigned to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Automation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutomationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutomationArgs | AutomationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AutomationState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["sources"] = state ? state.sources : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AutomationArgs | undefined;
            if ((!args || args.actions === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.scopes === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'scopes'");
            }
            if ((!args || args.sources === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sources'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
            inputs["sources"] = args ? args.sources : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Automation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Automation resources.
 */
export interface AutomationState {
    /**
     * One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
     */
    readonly actions?: pulumi.Input<pulumi.Input<inputs.securitycenter.AutomationAction>[]>;
    /**
     * Specifies the description for the Security Center Automation.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Boolean to enable or disable this Security Center Automation.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
     */
    readonly sources?: pulumi.Input<pulumi.Input<inputs.securitycenter.AutomationSource>[]>;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Automation resource.
 */
export interface AutomationArgs {
    /**
     * One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
     */
    readonly actions: pulumi.Input<pulumi.Input<inputs.securitycenter.AutomationAction>[]>;
    /**
     * Specifies the description for the Security Center Automation.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Boolean to enable or disable this Security Center Automation.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
     */
    readonly scopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
     */
    readonly sources: pulumi.Input<pulumi.Input<inputs.securitycenter.AutomationSource>[]>;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
