// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Bot
{
    /// <summary>
    /// Manages a Directline integration for a Bot Channel
    /// </summary>
    public partial class ChannelDirectLine : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Output("botName")]
        public Output<string> BotName { get; private set; } = null!;

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A site represents a client application that you want to connect to your bot. Multiple `site` blocks may be defined as below
        /// </summary>
        [Output("sites")]
        public Output<ImmutableArray<Outputs.ChannelDirectLineSite>> Sites { get; private set; } = null!;


        /// <summary>
        /// Create a ChannelDirectLine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ChannelDirectLine(string name, ChannelDirectLineArgs args, CustomResourceOptions? options = null)
            : base("azure:bot/channelDirectLine:ChannelDirectLine", name, args ?? new ChannelDirectLineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ChannelDirectLine(string name, Input<string> id, ChannelDirectLineState? state = null, CustomResourceOptions? options = null)
            : base("azure:bot/channelDirectLine:ChannelDirectLine", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ChannelDirectLine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ChannelDirectLine Get(string name, Input<string> id, ChannelDirectLineState? state = null, CustomResourceOptions? options = null)
        {
            return new ChannelDirectLine(name, id, state, options);
        }
    }

    public sealed class ChannelDirectLineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Input("botName", required: true)]
        public Input<string> BotName { get; set; } = null!;

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("sites", required: true)]
        private InputList<Inputs.ChannelDirectLineSiteArgs>? _sites;

        /// <summary>
        /// A site represents a client application that you want to connect to your bot. Multiple `site` blocks may be defined as below
        /// </summary>
        public InputList<Inputs.ChannelDirectLineSiteArgs> Sites
        {
            get => _sites ?? (_sites = new InputList<Inputs.ChannelDirectLineSiteArgs>());
            set => _sites = value;
        }

        public ChannelDirectLineArgs()
        {
        }
    }

    public sealed class ChannelDirectLineState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Input("botName")]
        public Input<string>? BotName { get; set; }

        /// <summary>
        /// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("sites")]
        private InputList<Inputs.ChannelDirectLineSiteGetArgs>? _sites;

        /// <summary>
        /// A site represents a client application that you want to connect to your bot. Multiple `site` blocks may be defined as below
        /// </summary>
        public InputList<Inputs.ChannelDirectLineSiteGetArgs> Sites
        {
            get => _sites ?? (_sites = new InputList<Inputs.ChannelDirectLineSiteGetArgs>());
            set => _sites = value;
        }

        public ChannelDirectLineState()
        {
        }
    }
}
