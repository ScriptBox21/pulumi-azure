// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DevSpace Controller.
 *
 * > **NOTE:** Microsoft will be retiring Azure Dev Spaces on 31 October 2023, please see the product [documentation](https://azure.microsoft.com/en-us/updates/azure-dev-spaces-is-retiring-on-31-october-2023/) for more information.
 *
 * ## Import
 *
 * DevSpace Controller's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:devspace/controller:Controller controller1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevSpaces/controllers/controller1Name
 * ```
 */
export class Controller extends pulumi.CustomResource {
    /**
     * Get an existing Controller resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ControllerState, opts?: pulumi.CustomResourceOptions): Controller {
        return new Controller(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:devspace/controller:Controller';

    /**
     * Returns true if the given object is an instance of Controller.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Controller {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Controller.__pulumiType;
    }

    /**
     * DNS name for accessing DataPlane services.
     */
    public /*out*/ readonly dataPlaneFqdn!: pulumi.Output<string>;
    /**
     * The host suffix for the DevSpace Controller.
     */
    public /*out*/ readonly hostSuffix!: pulumi.Output<string>;
    /**
     * Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    public readonly targetContainerHostCredentialsBase64!: pulumi.Output<string>;
    /**
     * The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    public readonly targetContainerHostResourceId!: pulumi.Output<string>;

    /**
     * Create a Controller resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ControllerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ControllerArgs | ControllerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ControllerState | undefined;
            inputs["dataPlaneFqdn"] = state ? state.dataPlaneFqdn : undefined;
            inputs["hostSuffix"] = state ? state.hostSuffix : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetContainerHostCredentialsBase64"] = state ? state.targetContainerHostCredentialsBase64 : undefined;
            inputs["targetContainerHostResourceId"] = state ? state.targetContainerHostResourceId : undefined;
        } else {
            const args = argsOrState as ControllerArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.skuName === undefined) {
                throw new Error("Missing required property 'skuName'");
            }
            if (!args || args.targetContainerHostCredentialsBase64 === undefined) {
                throw new Error("Missing required property 'targetContainerHostCredentialsBase64'");
            }
            if (!args || args.targetContainerHostResourceId === undefined) {
                throw new Error("Missing required property 'targetContainerHostResourceId'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetContainerHostCredentialsBase64"] = args ? args.targetContainerHostCredentialsBase64 : undefined;
            inputs["targetContainerHostResourceId"] = args ? args.targetContainerHostResourceId : undefined;
            inputs["dataPlaneFqdn"] = undefined /*out*/;
            inputs["hostSuffix"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Controller.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Controller resources.
 */
export interface ControllerState {
    /**
     * DNS name for accessing DataPlane services.
     */
    readonly dataPlaneFqdn?: pulumi.Input<string>;
    /**
     * The host suffix for the DevSpace Controller.
     */
    readonly hostSuffix?: pulumi.Input<string>;
    /**
     * Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostCredentialsBase64?: pulumi.Input<string>;
    /**
     * The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostResourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Controller resource.
 */
export interface ControllerArgs {
    /**
     * Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostCredentialsBase64: pulumi.Input<string>;
    /**
     * The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostResourceId: pulumi.Input<string>;
}
