// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    /// <summary>
    /// Manages a AutoScale Setting which can be applied to Virtual Machine Scale Sets, App Services and other scalable resources.
    /// </summary>
    public partial class AutoscaleSetting : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the AutoScale Setting. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies a `notification` block as defined below.
        /// </summary>
        [Output("notification")]
        public Output<Outputs.AutoscaleSettingNotification?> Notification { get; private set; } = null!;

        /// <summary>
        /// Specifies one or more (up to 20) `profile` blocks as defined below.
        /// </summary>
        [Output("profiles")]
        public Output<ImmutableArray<Outputs.AutoscaleSettingProfile>> Profiles { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the resource ID of the resource that the autoscale setting should be added to.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a AutoscaleSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoscaleSetting(string name, AutoscaleSettingArgs args, CustomResourceOptions? options = null)
            : base("azure:monitoring/autoscaleSetting:AutoscaleSetting", name, args ?? new AutoscaleSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoscaleSetting(string name, Input<string> id, AutoscaleSettingState? state = null, CustomResourceOptions? options = null)
            : base("azure:monitoring/autoscaleSetting:AutoscaleSetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AutoscaleSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoscaleSetting Get(string name, Input<string> id, AutoscaleSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new AutoscaleSetting(name, id, state, options);
        }
    }

    public sealed class AutoscaleSettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the AutoScale Setting. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies a `notification` block as defined below.
        /// </summary>
        [Input("notification")]
        public Input<Inputs.AutoscaleSettingNotificationArgs>? Notification { get; set; }

        [Input("profiles", required: true)]
        private InputList<Inputs.AutoscaleSettingProfileArgs>? _profiles;

        /// <summary>
        /// Specifies one or more (up to 20) `profile` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AutoscaleSettingProfileArgs> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<Inputs.AutoscaleSettingProfileArgs>());
            set => _profiles = value;
        }

        /// <summary>
        /// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the resource ID of the resource that the autoscale setting should be added to.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        public AutoscaleSettingArgs()
        {
        }
    }

    public sealed class AutoscaleSettingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the AutoScale Setting. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies a `notification` block as defined below.
        /// </summary>
        [Input("notification")]
        public Input<Inputs.AutoscaleSettingNotificationGetArgs>? Notification { get; set; }

        [Input("profiles")]
        private InputList<Inputs.AutoscaleSettingProfileGetArgs>? _profiles;

        /// <summary>
        /// Specifies one or more (up to 20) `profile` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AutoscaleSettingProfileGetArgs> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<Inputs.AutoscaleSettingProfileGetArgs>());
            set => _profiles = value;
        }

        /// <summary>
        /// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the resource ID of the resource that the autoscale setting should be added to.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        public AutoscaleSettingState()
        {
        }
    }
}
