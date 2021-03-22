// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Azure Security Center Data Connector.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "west europe"});
 * const exampleAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "PerGB2018",
 * });
 * const exampleDataConnectorAzureSecurityCenter = new azure.sentinel.DataConnectorAzureSecurityCenter("exampleDataConnectorAzureSecurityCenter", {logAnalyticsWorkspaceId: exampleAnalyticsWorkspace.id});
 * ```
 *
 * ## Import
 *
 * Azure Security Center Data Connectors can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:sentinel/dataConnectorAzureSecurityCenter:DataConnectorAzureSecurityCenter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
 * ```
 */
export class DataConnectorAzureSecurityCenter extends pulumi.CustomResource {
    /**
     * Get an existing DataConnectorAzureSecurityCenter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataConnectorAzureSecurityCenterState, opts?: pulumi.CustomResourceOptions): DataConnectorAzureSecurityCenter {
        return new DataConnectorAzureSecurityCenter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:sentinel/dataConnectorAzureSecurityCenter:DataConnectorAzureSecurityCenter';

    /**
     * Returns true if the given object is an instance of DataConnectorAzureSecurityCenter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataConnectorAzureSecurityCenter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataConnectorAzureSecurityCenter.__pulumiType;
    }

    /**
     * The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    public readonly logAnalyticsWorkspaceId!: pulumi.Output<string>;
    /**
     * The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    public readonly subscriptionId!: pulumi.Output<string>;

    /**
     * Create a DataConnectorAzureSecurityCenter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataConnectorAzureSecurityCenterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataConnectorAzureSecurityCenterArgs | DataConnectorAzureSecurityCenterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataConnectorAzureSecurityCenterState | undefined;
            inputs["logAnalyticsWorkspaceId"] = state ? state.logAnalyticsWorkspaceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["subscriptionId"] = state ? state.subscriptionId : undefined;
        } else {
            const args = argsOrState as DataConnectorAzureSecurityCenterArgs | undefined;
            if ((!args || args.logAnalyticsWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logAnalyticsWorkspaceId'");
            }
            inputs["logAnalyticsWorkspaceId"] = args ? args.logAnalyticsWorkspaceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["subscriptionId"] = args ? args.subscriptionId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DataConnectorAzureSecurityCenter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataConnectorAzureSecurityCenter resources.
 */
export interface DataConnectorAzureSecurityCenterState {
    /**
     * The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    readonly logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    readonly subscriptionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataConnectorAzureSecurityCenter resource.
 */
export interface DataConnectorAzureSecurityCenterArgs {
    /**
     * The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    readonly logAnalyticsWorkspaceId: pulumi.Input<string>;
    /**
     * The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
     */
    readonly subscriptionId?: pulumi.Input<string>;
}