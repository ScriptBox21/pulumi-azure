// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateDns
{
    /// <summary>
    /// Enables you to manage DNS SRV Records within Azure Private DNS.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/private_dns_srv_record.html.markdown.
    /// </summary>
    public partial class SRVRecord : Pulumi.CustomResource
    {
        /// <summary>
        /// The FQDN of the DNS SRV Record.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The name of the DNS SRV Record. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more `record` blocks as defined below.
        /// </summary>
        [Output("records")]
        public Output<ImmutableArray<Outputs.SRVRecordRecords>> Records { get; private set; } = null!;

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zoneName")]
        public Output<string> ZoneName { get; private set; } = null!;


        /// <summary>
        /// Create a SRVRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SRVRecord(string name, SRVRecordArgs args, CustomResourceOptions? options = null)
            : base("azure:privatedns/sRVRecord:SRVRecord", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SRVRecord(string name, Input<string> id, SRVRecordState? state = null, CustomResourceOptions? options = null)
            : base("azure:privatedns/sRVRecord:SRVRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SRVRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SRVRecord Get(string name, Input<string> id, SRVRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new SRVRecord(name, id, state, options);
        }
    }

    public sealed class SRVRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DNS SRV Record. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records", required: true)]
        private InputList<Inputs.SRVRecordRecordsArgs>? _records;

        /// <summary>
        /// One or more `record` blocks as defined below.
        /// </summary>
        public InputList<Inputs.SRVRecordRecordsArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.SRVRecordRecordsArgs>());
            set => _records = value;
        }

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
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

        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        /// <summary>
        /// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zoneName", required: true)]
        public Input<string> ZoneName { get; set; } = null!;

        public SRVRecordArgs()
        {
        }
    }

    public sealed class SRVRecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FQDN of the DNS SRV Record.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The name of the DNS SRV Record. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records")]
        private InputList<Inputs.SRVRecordRecordsGetArgs>? _records;

        /// <summary>
        /// One or more `record` blocks as defined below.
        /// </summary>
        public InputList<Inputs.SRVRecordRecordsGetArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.SRVRecordRecordsGetArgs>());
            set => _records = value;
        }

        /// <summary>
        /// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
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

        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zoneName")]
        public Input<string>? ZoneName { get; set; }

        public SRVRecordState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SRVRecordRecordsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Port the service is listening on.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The priority of the SRV record.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// The FQDN of the service.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// The Weight of the SRV record.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public SRVRecordRecordsArgs()
        {
        }
    }

    public sealed class SRVRecordRecordsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Port the service is listening on.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The priority of the SRV record.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// The FQDN of the service.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// The Weight of the SRV record.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public SRVRecordRecordsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SRVRecordRecords
    {
        /// <summary>
        /// The Port the service is listening on.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The priority of the SRV record.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The FQDN of the service.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// The Weight of the SRV record.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private SRVRecordRecords(
            int port,
            int priority,
            string target,
            int weight)
        {
            Port = port;
            Priority = priority;
            Target = target;
            Weight = weight;
        }
    }
    }
}
