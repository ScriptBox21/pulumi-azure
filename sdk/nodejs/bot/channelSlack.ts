// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Slack integration for a Bot Channel
 * 
 * > **Note** A bot can only have a single Slack Channel associated with it.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const current = azure.core.getClientConfig();
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "northeurope",
 * });
 * const exampleChannelsRegistration = new azure.bot.ChannelsRegistration("example", {
 *     location: "global",
 *     microsoftAppId: current.clientId,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "F0",
 * });
 * const exampleChannelSlack = new azure.bot.ChannelSlack("example", {
 *     botName: exampleChannelsRegistration.name,
 *     clientId: "exampleId",
 *     clientSecret: "exampleSecret",
 *     location: exampleChannelsRegistration.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     verificationToken: "exampleVerificationToken",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channel_slack.markdown.
 */
export class ChannelSlack extends pulumi.CustomResource {
    /**
     * Get an existing ChannelSlack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ChannelSlackState, opts?: pulumi.CustomResourceOptions): ChannelSlack {
        return new ChannelSlack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:bot/channelSlack:ChannelSlack';

    /**
     * Returns true if the given object is an instance of ChannelSlack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChannelSlack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChannelSlack.__pulumiType;
    }

    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    public readonly botName!: pulumi.Output<string>;
    /**
     * The Client ID that will be used to authenticate with Slack.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The Client Secret that will be used to authenticate with Slack.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The Slack Landing Page URL.
     */
    public readonly landingPageUrl!: pulumi.Output<string | undefined>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Verification Token that will be used to authenticate with Slack.
     */
    public readonly verificationToken!: pulumi.Output<string>;

    /**
     * Create a ChannelSlack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChannelSlackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ChannelSlackArgs | ChannelSlackState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ChannelSlackState | undefined;
            inputs["botName"] = state ? state.botName : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["landingPageUrl"] = state ? state.landingPageUrl : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["verificationToken"] = state ? state.verificationToken : undefined;
        } else {
            const args = argsOrState as ChannelSlackArgs | undefined;
            if (!args || args.botName === undefined) {
                throw new Error("Missing required property 'botName'");
            }
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.clientSecret === undefined) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.verificationToken === undefined) {
                throw new Error("Missing required property 'verificationToken'");
            }
            inputs["botName"] = args ? args.botName : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["landingPageUrl"] = args ? args.landingPageUrl : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["verificationToken"] = args ? args.verificationToken : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ChannelSlack.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ChannelSlack resources.
 */
export interface ChannelSlackState {
    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    readonly botName?: pulumi.Input<string>;
    /**
     * The Client ID that will be used to authenticate with Slack.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The Client Secret that will be used to authenticate with Slack.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The Slack Landing Page URL.
     */
    readonly landingPageUrl?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Verification Token that will be used to authenticate with Slack.
     */
    readonly verificationToken?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ChannelSlack resource.
 */
export interface ChannelSlackArgs {
    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    readonly botName: pulumi.Input<string>;
    /**
     * The Client ID that will be used to authenticate with Slack.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * The Client Secret that will be used to authenticate with Slack.
     */
    readonly clientSecret: pulumi.Input<string>;
    /**
     * The Slack Landing Page URL.
     */
    readonly landingPageUrl?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Verification Token that will be used to authenticate with Slack.
     */
    readonly verificationToken: pulumi.Input<string>;
}
