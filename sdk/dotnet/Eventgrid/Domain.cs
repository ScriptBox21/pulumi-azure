// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid
{
    /// <summary>
    /// Manages an EventGrid Domain
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventgrid_domain.html.markdown.
    /// </summary>
    public partial class Domain : Pulumi.CustomResource
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Domain.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Output("inputMappingDefaultValues")]
        public Output<Outputs.DomainInputMappingDefaultValues?> InputMappingDefaultValues { get; private set; } = null!;

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Output("inputMappingFields")]
        public Output<Outputs.DomainInputMappingFields?> InputMappingFields { get; private set; } = null!;

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("inputSchema")]
        public Output<string?> InputSchema { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Primary Shared Access Key associated with the EventGrid Domain.
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Shared Access Key associated with the EventGrid Domain.
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("azure:eventgrid/domain:Domain", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventgrid/domain:Domain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,                Aliases = { new Alias { Type = "azure:eventhub/domain:Domain" } },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Input("inputMappingDefaultValues")]
        public Input<Inputs.DomainInputMappingDefaultValuesArgs>? InputMappingDefaultValues { get; set; }

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Input("inputMappingFields")]
        public Input<Inputs.DomainInputMappingFieldsArgs>? InputMappingFields { get; set; }

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
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

        public DomainArgs()
        {
        }
    }

    public sealed class DomainState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Domain.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        [Input("inputMappingDefaultValues")]
        public Input<Inputs.DomainInputMappingDefaultValuesGetArgs>? InputMappingDefaultValues { get; set; }

        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        [Input("inputMappingFields")]
        public Input<Inputs.DomainInputMappingFieldsGetArgs>? InputMappingFields { get; set; }

        /// <summary>
        /// Specifies the schema in which incoming events will be published to this domain. Allowed values are `cloudeventv01schema`, `customeventschema`, or `eventgridschema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("inputSchema")]
        public Input<string>? InputSchema { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Primary Shared Access Key associated with the EventGrid Domain.
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Shared Access Key associated with the EventGrid Domain.
        /// </summary>
        [Input("secondaryAccessKey")]
        public Input<string>? SecondaryAccessKey { get; set; }

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

        public DomainState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DomainInputMappingDefaultValuesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataVersion")]
        public Input<string>? DataVersion { get; set; }

        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        public DomainInputMappingDefaultValuesArgs()
        {
        }
    }

    public sealed class DomainInputMappingDefaultValuesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataVersion")]
        public Input<string>? DataVersion { get; set; }

        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        public DomainInputMappingDefaultValuesGetArgs()
        {
        }
    }

    public sealed class DomainInputMappingFieldsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataVersion")]
        public Input<string>? DataVersion { get; set; }

        /// <summary>
        /// Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventTime")]
        public Input<string>? EventTime { get; set; }

        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public DomainInputMappingFieldsArgs()
        {
        }
    }

    public sealed class DomainInputMappingFieldsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataVersion")]
        public Input<string>? DataVersion { get; set; }

        /// <summary>
        /// Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventTime")]
        public Input<string>? EventTime { get; set; }

        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public DomainInputMappingFieldsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DomainInputMappingDefaultValues
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? DataVersion;
        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? EventType;
        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Subject;

        [OutputConstructor]
        private DomainInputMappingDefaultValues(
            string? dataVersion,
            string? eventType,
            string? subject)
        {
            DataVersion = dataVersion;
            EventType = eventType;
            Subject = subject;
        }
    }

    [OutputType]
    public sealed class DomainInputMappingFields
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? DataVersion;
        /// <summary>
        /// Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? EventTime;
        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? EventType;
        /// <summary>
        /// Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Subject;
        /// <summary>
        /// Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Topic;

        [OutputConstructor]
        private DomainInputMappingFields(
            string? dataVersion,
            string? eventTime,
            string? eventType,
            string? id,
            string? subject,
            string? topic)
        {
            DataVersion = dataVersion;
            EventTime = eventTime;
            EventType = eventType;
            Id = id;
            Subject = subject;
            Topic = topic;
        }
    }
    }
}
