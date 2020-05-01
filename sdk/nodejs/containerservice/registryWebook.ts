// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Azure Container Registry Webhook.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_registry_webhook.html.markdown.
 */
/** @deprecated azure.containerservice.RegistryWebook has been deprecated in favour of azure.containerservice.RegistryWebhook */
export class RegistryWebook extends pulumi.CustomResource {
    /**
     * Get an existing RegistryWebook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryWebookState, opts?: pulumi.CustomResourceOptions): RegistryWebook {
        pulumi.log.warn("RegistryWebook is deprecated: azure.containerservice.RegistryWebook has been deprecated in favour of azure.containerservice.RegistryWebhook")
        return new RegistryWebook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:containerservice/registryWebook:RegistryWebook';

    /**
     * Returns true if the given object is an instance of RegistryWebook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryWebook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryWebook.__pulumiType;
    }

    /**
     * A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
     */
    public readonly actions!: pulumi.Output<string[]>;
    /**
     * Custom headers that will be added to the webhook notifications request.
     */
    public readonly customHeaders!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
     */
    public readonly registryName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Specifies the service URI for the Webhook to post notifications.
     */
    public readonly serviceUri!: pulumi.Output<string>;
    /**
     * Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a RegistryWebook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.containerservice.RegistryWebook has been deprecated in favour of azure.containerservice.RegistryWebhook */
    constructor(name: string, args: RegistryWebookArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.containerservice.RegistryWebook has been deprecated in favour of azure.containerservice.RegistryWebhook */
    constructor(name: string, argsOrState?: RegistryWebookArgs | RegistryWebookState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("RegistryWebook is deprecated: azure.containerservice.RegistryWebook has been deprecated in favour of azure.containerservice.RegistryWebhook")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegistryWebookState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["customHeaders"] = state ? state.customHeaders : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["registryName"] = state ? state.registryName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["serviceUri"] = state ? state.serviceUri : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RegistryWebookArgs | undefined;
            if (!args || args.actions === undefined) {
                throw new Error("Missing required property 'actions'");
            }
            if (!args || args.registryName === undefined) {
                throw new Error("Missing required property 'registryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceUri === undefined) {
                throw new Error("Missing required property 'serviceUri'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["customHeaders"] = args ? args.customHeaders : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["serviceUri"] = args ? args.serviceUri : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegistryWebook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegistryWebook resources.
 */
export interface RegistryWebookState {
    /**
     * A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
     */
    readonly actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom headers that will be added to the webhook notifications request.
     */
    readonly customHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
     */
    readonly registryName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * Specifies the service URI for the Webhook to post notifications.
     */
    readonly serviceUri?: pulumi.Input<string>;
    /**
     * Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
     */
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a RegistryWebook resource.
 */
export interface RegistryWebookArgs {
    /**
     * A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
     */
    readonly actions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom headers that will be added to the webhook notifications request.
     */
    readonly customHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * Specifies the service URI for the Webhook to post notifications.
     */
    readonly serviceUri: pulumi.Input<string>;
    /**
     * Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
     */
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
