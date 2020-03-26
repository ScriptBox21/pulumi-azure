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
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub.html.markdown.
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
            : base("azure:eventhub/eventHub:EventHub", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    namespace Inputs
    {

    public sealed class EventHubCaptureDescriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `destination` block as defined below.
        /// </summary>
        [Input("destination", required: true)]
        public Input<EventHubCaptureDescriptionDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Specifies if the Capture Description is Enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Specifies the Encoding used for the Capture Description. Possible values are `Avro` and `AvroDeflate`.
        /// </summary>
        [Input("encoding", required: true)]
        public Input<string> Encoding { get; set; } = null!;

        /// <summary>
        /// Specifies the time interval in seconds at which the capture will happen. Values can be between `60` and `900` seconds. Defaults to `300` seconds.
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        /// <summary>
        /// Specifies the amount of data built up in your EventHub before a Capture Operation occurs. Value should be between `10485760` and `524288000`  bytes. Defaults to `314572800` bytes.
        /// </summary>
        [Input("sizeLimitInBytes")]
        public Input<int>? SizeLimitInBytes { get; set; }

        /// <summary>
        /// Specifies if empty files should not be emitted if no events occur during the Capture time window.  Defaults to `false`.
        /// </summary>
        [Input("skipEmptyArchives")]
        public Input<bool>? SkipEmptyArchives { get; set; }

        public EventHubCaptureDescriptionArgs()
        {
        }
    }

    public sealed class EventHubCaptureDescriptionDestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Blob naming convention for archiving. e.g. `{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}`. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
        /// </summary>
        [Input("archiveNameFormat", required: true)]
        public Input<string> ArchiveNameFormat { get; set; } = null!;

        /// <summary>
        /// The name of the Container within the Blob Storage Account where messages should be archived.
        /// </summary>
        [Input("blobContainerName", required: true)]
        public Input<string> BlobContainerName { get; set; } = null!;

        /// <summary>
        /// The Name of the Destination where the capture should take place. At this time the only supported value is `EventHubArchive.AzureBlockBlob`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the Blob Storage Account where messages should be archived.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public EventHubCaptureDescriptionDestinationArgs()
        {
        }
    }

    public sealed class EventHubCaptureDescriptionDestinationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Blob naming convention for archiving. e.g. `{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}`. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
        /// </summary>
        [Input("archiveNameFormat", required: true)]
        public Input<string> ArchiveNameFormat { get; set; } = null!;

        /// <summary>
        /// The name of the Container within the Blob Storage Account where messages should be archived.
        /// </summary>
        [Input("blobContainerName", required: true)]
        public Input<string> BlobContainerName { get; set; } = null!;

        /// <summary>
        /// The Name of the Destination where the capture should take place. At this time the only supported value is `EventHubArchive.AzureBlockBlob`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the Blob Storage Account where messages should be archived.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public EventHubCaptureDescriptionDestinationGetArgs()
        {
        }
    }

    public sealed class EventHubCaptureDescriptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `destination` block as defined below.
        /// </summary>
        [Input("destination", required: true)]
        public Input<EventHubCaptureDescriptionDestinationGetArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Specifies if the Capture Description is Enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Specifies the Encoding used for the Capture Description. Possible values are `Avro` and `AvroDeflate`.
        /// </summary>
        [Input("encoding", required: true)]
        public Input<string> Encoding { get; set; } = null!;

        /// <summary>
        /// Specifies the time interval in seconds at which the capture will happen. Values can be between `60` and `900` seconds. Defaults to `300` seconds.
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        /// <summary>
        /// Specifies the amount of data built up in your EventHub before a Capture Operation occurs. Value should be between `10485760` and `524288000`  bytes. Defaults to `314572800` bytes.
        /// </summary>
        [Input("sizeLimitInBytes")]
        public Input<int>? SizeLimitInBytes { get; set; }

        /// <summary>
        /// Specifies if empty files should not be emitted if no events occur during the Capture time window.  Defaults to `false`.
        /// </summary>
        [Input("skipEmptyArchives")]
        public Input<bool>? SkipEmptyArchives { get; set; }

        public EventHubCaptureDescriptionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EventHubCaptureDescription
    {
        /// <summary>
        /// A `destination` block as defined below.
        /// </summary>
        public readonly EventHubCaptureDescriptionDestination Destination;
        /// <summary>
        /// Specifies if the Capture Description is Enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Specifies the Encoding used for the Capture Description. Possible values are `Avro` and `AvroDeflate`.
        /// </summary>
        public readonly string Encoding;
        /// <summary>
        /// Specifies the time interval in seconds at which the capture will happen. Values can be between `60` and `900` seconds. Defaults to `300` seconds.
        /// </summary>
        public readonly int? IntervalInSeconds;
        /// <summary>
        /// Specifies the amount of data built up in your EventHub before a Capture Operation occurs. Value should be between `10485760` and `524288000`  bytes. Defaults to `314572800` bytes.
        /// </summary>
        public readonly int? SizeLimitInBytes;
        /// <summary>
        /// Specifies if empty files should not be emitted if no events occur during the Capture time window.  Defaults to `false`.
        /// </summary>
        public readonly bool? SkipEmptyArchives;

        [OutputConstructor]
        private EventHubCaptureDescription(
            EventHubCaptureDescriptionDestination destination,
            bool enabled,
            string encoding,
            int? intervalInSeconds,
            int? sizeLimitInBytes,
            bool? skipEmptyArchives)
        {
            Destination = destination;
            Enabled = enabled;
            Encoding = encoding;
            IntervalInSeconds = intervalInSeconds;
            SizeLimitInBytes = sizeLimitInBytes;
            SkipEmptyArchives = skipEmptyArchives;
        }
    }

    [OutputType]
    public sealed class EventHubCaptureDescriptionDestination
    {
        /// <summary>
        /// The Blob naming convention for archiving. e.g. `{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}`. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
        /// </summary>
        public readonly string ArchiveNameFormat;
        /// <summary>
        /// The name of the Container within the Blob Storage Account where messages should be archived.
        /// </summary>
        public readonly string BlobContainerName;
        /// <summary>
        /// The Name of the Destination where the capture should take place. At this time the only supported value is `EventHubArchive.AzureBlockBlob`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the Blob Storage Account where messages should be archived.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private EventHubCaptureDescriptionDestination(
            string archiveNameFormat,
            string blobContainerName,
            string name,
            string storageAccountId)
        {
            ArchiveNameFormat = archiveNameFormat;
            BlobContainerName = blobContainerName;
            Name = name;
            StorageAccountId = storageAccountId;
        }
    }
    }
}
