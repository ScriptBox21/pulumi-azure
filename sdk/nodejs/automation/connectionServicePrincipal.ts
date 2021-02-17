// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Automation Connection with type `AzureServicePrincipal`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * from "fs";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleClientConfig = azure.core.getClientConfig({});
 * const exampleAccount = new azure.automation.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: [{
 *         name: "Basic",
 *     }],
 * });
 * const exampleConnectionServicePrincipal = new azure.automation.ConnectionServicePrincipal("exampleConnectionServicePrincipal", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     automationAccountName: exampleAccount.name,
 *     applicationId: "00000000-0000-0000-0000-000000000000",
 *     tenantId: exampleClientConfig.then(exampleClientConfig => exampleClientConfig.tenantId),
 *     subscriptionId: exampleClientConfig.then(exampleClientConfig => exampleClientConfig.subscriptionId),
 *     certificateThumbprint: fs.readFileSync("automation_certificate_test.thumb"),
 * });
 * ```
 *
 * ## Import
 *
 * Automation Connection can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:automation/connectionServicePrincipal:ConnectionServicePrincipal conn1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
 * ```
 */
export class ConnectionServicePrincipal extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionServicePrincipal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionServicePrincipalState, opts?: pulumi.CustomResourceOptions): ConnectionServicePrincipal {
        return new ConnectionServicePrincipal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:automation/connectionServicePrincipal:ConnectionServicePrincipal';

    /**
     * Returns true if the given object is an instance of ConnectionServicePrincipal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionServicePrincipal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionServicePrincipal.__pulumiType;
    }

    /**
     * The (Client) ID of the Service Principal.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    public readonly automationAccountName!: pulumi.Output<string>;
    /**
     * The thumbprint of the Service Principal Certificate.
     */
    public readonly certificateThumbprint!: pulumi.Output<string>;
    /**
     * A description for this Connection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Connection. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The subscription GUID.
     */
    public readonly subscriptionId!: pulumi.Output<string>;
    /**
     * The ID of the Tenant the Service Principal is assigned in.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a ConnectionServicePrincipal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionServicePrincipalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionServicePrincipalArgs | ConnectionServicePrincipalState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionServicePrincipalState | undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["automationAccountName"] = state ? state.automationAccountName : undefined;
            inputs["certificateThumbprint"] = state ? state.certificateThumbprint : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subscriptionId"] = state ? state.subscriptionId : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as ConnectionServicePrincipalArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.automationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.certificateThumbprint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateThumbprint'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.subscriptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionId'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["certificateThumbprint"] = args ? args.certificateThumbprint : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subscriptionId"] = args ? args.subscriptionId : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConnectionServicePrincipal.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectionServicePrincipal resources.
 */
export interface ConnectionServicePrincipalState {
    /**
     * The (Client) ID of the Service Principal.
     */
    readonly applicationId?: pulumi.Input<string>;
    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName?: pulumi.Input<string>;
    /**
     * The thumbprint of the Service Principal Certificate.
     */
    readonly certificateThumbprint?: pulumi.Input<string>;
    /**
     * A description for this Connection.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the name of the Connection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The subscription GUID.
     */
    readonly subscriptionId?: pulumi.Input<string>;
    /**
     * The ID of the Tenant the Service Principal is assigned in.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConnectionServicePrincipal resource.
 */
export interface ConnectionServicePrincipalArgs {
    /**
     * The (Client) ID of the Service Principal.
     */
    readonly applicationId: pulumi.Input<string>;
    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The thumbprint of the Service Principal Certificate.
     */
    readonly certificateThumbprint: pulumi.Input<string>;
    /**
     * A description for this Connection.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the name of the Connection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The subscription GUID.
     */
    readonly subscriptionId: pulumi.Input<string>;
    /**
     * The ID of the Tenant the Service Principal is assigned in.
     */
    readonly tenantId: pulumi.Input<string>;
}
