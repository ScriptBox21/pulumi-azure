// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Directline integration for a Bot Channel
 *
 * ## Import
 *
 * The Directline Channel for a Bot can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:bot/channelDirectLine:ChannelDirectLine example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/DirectlineChannel
 * ```
 */
export class ChannelDirectLine extends pulumi.CustomResource {
    /**
     * Get an existing ChannelDirectLine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ChannelDirectLineState, opts?: pulumi.CustomResourceOptions): ChannelDirectLine {
        return new ChannelDirectLine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:bot/channelDirectLine:ChannelDirectLine';

    /**
     * Returns true if the given object is an instance of ChannelDirectLine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChannelDirectLine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChannelDirectLine.__pulumiType;
    }

    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    public readonly botName!: pulumi.Output<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A site represents a client application that you want to connect to your bot. Multiple `site` blocks may be defined as below
     */
    public readonly sites!: pulumi.Output<outputs.bot.ChannelDirectLineSite[]>;

    /**
     * Create a ChannelDirectLine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChannelDirectLineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ChannelDirectLineArgs | ChannelDirectLineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ChannelDirectLineState | undefined;
            inputs["botName"] = state ? state.botName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sites"] = state ? state.sites : undefined;
        } else {
            const args = argsOrState as ChannelDirectLineArgs | undefined;
            if (!args || args.botName === undefined) {
                throw new Error("Missing required property 'botName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sites === undefined) {
                throw new Error("Missing required property 'sites'");
            }
            inputs["botName"] = args ? args.botName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sites"] = args ? args.sites : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ChannelDirectLine.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ChannelDirectLine resources.
 */
export interface ChannelDirectLineState {
    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    readonly botName?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A site represents a client application that you want to connect to your bot. Multiple `site` blocks may be defined as below
     */
    readonly sites?: pulumi.Input<pulumi.Input<inputs.bot.ChannelDirectLineSite>[]>;
}

/**
 * The set of arguments for constructing a ChannelDirectLine resource.
 */
export interface ChannelDirectLineArgs {
    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    readonly botName: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A site represents a client application that you want to connect to your bot. Multiple `site` blocks may be defined as below
     */
    readonly sites: pulumi.Input<pulumi.Input<inputs.bot.ChannelDirectLineSite>[]>;
}
