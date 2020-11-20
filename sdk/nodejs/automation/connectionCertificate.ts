// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Automation Connection with type `Azure`.
 *
 * ## Import
 *
 * Automation Connection can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:automation/connectionCertificate:ConnectionCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
 * ```
 */
export class ConnectionCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionCertificateState, opts?: pulumi.CustomResourceOptions): ConnectionCertificate {
        return new ConnectionCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:automation/connectionCertificate:ConnectionCertificate';

    /**
     * Returns true if the given object is an instance of ConnectionCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionCertificate.__pulumiType;
    }

    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    public readonly automationAccountName!: pulumi.Output<string>;
    /**
     * The name of the automation certificate.
     */
    public readonly automationCertificateName!: pulumi.Output<string>;
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
     * The id of subscription where the automation certificate exists.
     */
    public readonly subscriptionId!: pulumi.Output<string>;

    /**
     * Create a ConnectionCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionCertificateArgs | ConnectionCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConnectionCertificateState | undefined;
            inputs["automationAccountName"] = state ? state.automationAccountName : undefined;
            inputs["automationCertificateName"] = state ? state.automationCertificateName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subscriptionId"] = state ? state.subscriptionId : undefined;
        } else {
            const args = argsOrState as ConnectionCertificateArgs | undefined;
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.automationCertificateName === undefined) {
                throw new Error("Missing required property 'automationCertificateName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.subscriptionId === undefined) {
                throw new Error("Missing required property 'subscriptionId'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["automationCertificateName"] = args ? args.automationCertificateName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subscriptionId"] = args ? args.subscriptionId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConnectionCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectionCertificate resources.
 */
export interface ConnectionCertificateState {
    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName?: pulumi.Input<string>;
    /**
     * The name of the automation certificate.
     */
    readonly automationCertificateName?: pulumi.Input<string>;
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
     * The id of subscription where the automation certificate exists.
     */
    readonly subscriptionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConnectionCertificate resource.
 */
export interface ConnectionCertificateArgs {
    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The name of the automation certificate.
     */
    readonly automationCertificateName: pulumi.Input<string>;
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
     * The id of subscription where the automation certificate exists.
     */
    readonly subscriptionId: pulumi.Input<string>;
}
