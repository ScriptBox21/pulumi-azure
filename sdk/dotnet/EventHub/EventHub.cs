// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    /// <summary>
    /// Manages a Event Hubs as a nested resource within a Event Hubs namespace.
    /// </summary>
    public partial class EventHub : Pulumi.CustomResource
    {
        /// <summary>
        /// A `capture_description` block as defined below.
        /// </summary>
        [Output("captureDescription")]
        public Output<Outputs.EventHubCaptureDescription?> CaptureDescription { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
        /// </summary>
        [Output("messageRetention")]
        public Output<int> MessageRetention { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Output("partitionCount")]
        public Output<int> PartitionCount { get; private set; } = null!;

        /// <summary>
        /// The identifiers for partitions created for Event Hubs.
        /// </summary>
        [Output("partitionIds")]
        public Output<ImmutableArray<string>> PartitionIds { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a EventHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHub(string name, EventHubArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHub:EventHub", name, args ?? new EventHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHub(string name, Input<string> id, EventHubState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/eventHub:EventHub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHub Get(string name, Input<string> id, EventHubState? state = null, CustomResourceOptions? options = null)
        {
            return new EventHub(name, id, state, options);
        }
    }

    public sealed class EventHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `capture_description` block as defined below.
        /// </summary>
        [Input("captureDescription")]
        public Input<Inputs.EventHubCaptureDescriptionArgs>? CaptureDescription { get; set; }

        /// <summary>
        /// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
        /// </summary>
        [Input("messageRetention", required: true)]
        public Input<int> MessageRetention { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("partitionCount", required: true)]
        public Input<int> PartitionCount { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public EventHubArgs()
        {
        }
    }

    public sealed class EventHubState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `capture_description` block as defined below.
        /// </summary>
        [Input("captureDescription")]
        public Input<Inputs.EventHubCaptureDescriptionGetArgs>? CaptureDescription { get; set; }

        /// <summary>
        /// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
        /// </summary>
        [Input("messageRetention")]
        public Input<int>? MessageRetention { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        [Input("partitionIds")]
        private InputList<string>? _partitionIds;

        /// <summary>
        /// The identifiers for partitions created for Event Hubs.
        /// </summary>
        public InputList<string> PartitionIds
        {
            get => _partitionIds ?? (_partitionIds = new InputList<string>());
            set => _partitionIds = value;
        }

        /// <summary>
        /// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public EventHubState()
        {
        }
    }
}
