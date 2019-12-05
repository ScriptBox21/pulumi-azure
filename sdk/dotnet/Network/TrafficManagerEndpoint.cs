// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a Traffic Manager Endpoint.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/traffic_manager_endpoint.html.markdown.
    /// </summary>
    public partial class TrafficManagerEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `custom_header` blocks as defined below
        /// </summary>
        [Output("customHeaders")]
        public Output<ImmutableArray<Outputs.TrafficManagerEndpointCustomHeaders>> CustomHeaders { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure location of the Endpoint,
        /// this must be specified for Profiles using the `Performance` routing method
        /// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
        /// For Endpoints of type `azureEndpoints` the value will be taken from the
        /// location of the Azure target resource.
        /// </summary>
        [Output("endpointLocation")]
        public Output<string> EndpointLocation { get; private set; } = null!;

        [Output("endpointMonitorStatus")]
        public Output<string> EndpointMonitorStatus { get; private set; } = null!;

        /// <summary>
        /// The status of the Endpoint, can be set to
        /// either `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Output("endpointStatus")]
        public Output<string> EndpointStatus { get; private set; } = null!;

        /// <summary>
        /// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
        /// </summary>
        [Output("geoMappings")]
        public Output<ImmutableArray<string>> GeoMappings { get; private set; } = null!;

        /// <summary>
        /// This argument specifies the minimum number
        /// of endpoints that must be ‘online’ in the child profile in order for the
        /// parent profile to direct traffic to any of the endpoints in that child
        /// profile. This argument only applies to Endpoints of type `nestedEndpoints`
        /// and defaults to `1`.
        /// </summary>
        [Output("minChildEndpoints")]
        public Output<int?> MinChildEndpoints { get; private set; } = null!;

        /// <summary>
        /// The name of the Traffic Manager endpoint. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the priority of this Endpoint, this must be
        /// specified for Profiles using the `Priority` traffic routing method. Supports
        /// values between 1 and 1000, with no Endpoints sharing the same value. If
        /// omitted the value will be computed in order of creation.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The name of the Traffic Manager Profile to attach
        /// create the Traffic Manager endpoint.
        /// </summary>
        [Output("profileName")]
        public Output<string> ProfileName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the Traffic Manager endpoint.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// One or more `subnet` blocks as defined below
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<Outputs.TrafficManagerEndpointSubnets>> Subnets { get; private set; } = null!;

        /// <summary>
        /// The FQDN DNS name of the target. This argument must be
        /// provided for an endpoint of type `externalEndpoints`, for other types it
        /// will be computed.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        /// <summary>
        /// The resource id of an Azure resource to
        /// target. This argument must be provided for an endpoint of type
        /// `azureEndpoints` or `nestedEndpoints`.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string?> TargetResourceId { get; private set; } = null!;

        /// <summary>
        /// The Endpoint type, must be one of:
        /// - `azureEndpoints`
        /// - `externalEndpoints`
        /// - `nestedEndpoints`
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies how much traffic should be distributed to this
        /// endpoint, this must be specified for Profiles using the  `Weighted` traffic
        /// routing method. Supports values between 1 and 1000.
        /// </summary>
        [Output("weight")]
        public Output<int> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficManagerEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficManagerEndpoint(string name, TrafficManagerEndpointArgs args, CustomResourceOptions? options = null)
            : base("azure:network/trafficManagerEndpoint:TrafficManagerEndpoint", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private TrafficManagerEndpoint(string name, Input<string> id, TrafficManagerEndpointState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/trafficManagerEndpoint:TrafficManagerEndpoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,                Aliases = { new Alias { Type = "azure:trafficmanager/endpoint:Endpoint" } },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TrafficManagerEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficManagerEndpoint Get(string name, Input<string> id, TrafficManagerEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new TrafficManagerEndpoint(name, id, state, options);
        }
    }

    public sealed class TrafficManagerEndpointArgs : Pulumi.ResourceArgs
    {
        [Input("customHeaders")]
        private InputList<Inputs.TrafficManagerEndpointCustomHeadersArgs>? _customHeaders;

        /// <summary>
        /// One or more `custom_header` blocks as defined below
        /// </summary>
        public InputList<Inputs.TrafficManagerEndpointCustomHeadersArgs> CustomHeaders
        {
            get => _customHeaders ?? (_customHeaders = new InputList<Inputs.TrafficManagerEndpointCustomHeadersArgs>());
            set => _customHeaders = value;
        }

        /// <summary>
        /// Specifies the Azure location of the Endpoint,
        /// this must be specified for Profiles using the `Performance` routing method
        /// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
        /// For Endpoints of type `azureEndpoints` the value will be taken from the
        /// location of the Azure target resource.
        /// </summary>
        [Input("endpointLocation")]
        public Input<string>? EndpointLocation { get; set; }

        /// <summary>
        /// The status of the Endpoint, can be set to
        /// either `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Input("endpointStatus")]
        public Input<string>? EndpointStatus { get; set; }

        [Input("geoMappings")]
        private InputList<string>? _geoMappings;

        /// <summary>
        /// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
        /// </summary>
        public InputList<string> GeoMappings
        {
            get => _geoMappings ?? (_geoMappings = new InputList<string>());
            set => _geoMappings = value;
        }

        /// <summary>
        /// This argument specifies the minimum number
        /// of endpoints that must be ‘online’ in the child profile in order for the
        /// parent profile to direct traffic to any of the endpoints in that child
        /// profile. This argument only applies to Endpoints of type `nestedEndpoints`
        /// and defaults to `1`.
        /// </summary>
        [Input("minChildEndpoints")]
        public Input<int>? MinChildEndpoints { get; set; }

        /// <summary>
        /// The name of the Traffic Manager endpoint. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the priority of this Endpoint, this must be
        /// specified for Profiles using the `Priority` traffic routing method. Supports
        /// values between 1 and 1000, with no Endpoints sharing the same value. If
        /// omitted the value will be computed in order of creation.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The name of the Traffic Manager Profile to attach
        /// create the Traffic Manager endpoint.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the Traffic Manager endpoint.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("subnets")]
        private InputList<Inputs.TrafficManagerEndpointSubnetsArgs>? _subnets;

        /// <summary>
        /// One or more `subnet` blocks as defined below
        /// </summary>
        public InputList<Inputs.TrafficManagerEndpointSubnetsArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.TrafficManagerEndpointSubnetsArgs>());
            set => _subnets = value;
        }

        /// <summary>
        /// The FQDN DNS name of the target. This argument must be
        /// provided for an endpoint of type `externalEndpoints`, for other types it
        /// will be computed.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// The resource id of an Azure resource to
        /// target. This argument must be provided for an endpoint of type
        /// `azureEndpoints` or `nestedEndpoints`.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The Endpoint type, must be one of:
        /// - `azureEndpoints`
        /// - `externalEndpoints`
        /// - `nestedEndpoints`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies how much traffic should be distributed to this
        /// endpoint, this must be specified for Profiles using the  `Weighted` traffic
        /// routing method. Supports values between 1 and 1000.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public TrafficManagerEndpointArgs()
        {
        }
    }

    public sealed class TrafficManagerEndpointState : Pulumi.ResourceArgs
    {
        [Input("customHeaders")]
        private InputList<Inputs.TrafficManagerEndpointCustomHeadersGetArgs>? _customHeaders;

        /// <summary>
        /// One or more `custom_header` blocks as defined below
        /// </summary>
        public InputList<Inputs.TrafficManagerEndpointCustomHeadersGetArgs> CustomHeaders
        {
            get => _customHeaders ?? (_customHeaders = new InputList<Inputs.TrafficManagerEndpointCustomHeadersGetArgs>());
            set => _customHeaders = value;
        }

        /// <summary>
        /// Specifies the Azure location of the Endpoint,
        /// this must be specified for Profiles using the `Performance` routing method
        /// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
        /// For Endpoints of type `azureEndpoints` the value will be taken from the
        /// location of the Azure target resource.
        /// </summary>
        [Input("endpointLocation")]
        public Input<string>? EndpointLocation { get; set; }

        [Input("endpointMonitorStatus")]
        public Input<string>? EndpointMonitorStatus { get; set; }

        /// <summary>
        /// The status of the Endpoint, can be set to
        /// either `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Input("endpointStatus")]
        public Input<string>? EndpointStatus { get; set; }

        [Input("geoMappings")]
        private InputList<string>? _geoMappings;

        /// <summary>
        /// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
        /// </summary>
        public InputList<string> GeoMappings
        {
            get => _geoMappings ?? (_geoMappings = new InputList<string>());
            set => _geoMappings = value;
        }

        /// <summary>
        /// This argument specifies the minimum number
        /// of endpoints that must be ‘online’ in the child profile in order for the
        /// parent profile to direct traffic to any of the endpoints in that child
        /// profile. This argument only applies to Endpoints of type `nestedEndpoints`
        /// and defaults to `1`.
        /// </summary>
        [Input("minChildEndpoints")]
        public Input<int>? MinChildEndpoints { get; set; }

        /// <summary>
        /// The name of the Traffic Manager endpoint. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the priority of this Endpoint, this must be
        /// specified for Profiles using the `Priority` traffic routing method. Supports
        /// values between 1 and 1000, with no Endpoints sharing the same value. If
        /// omitted the value will be computed in order of creation.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The name of the Traffic Manager Profile to attach
        /// create the Traffic Manager endpoint.
        /// </summary>
        [Input("profileName")]
        public Input<string>? ProfileName { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Traffic Manager endpoint.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("subnets")]
        private InputList<Inputs.TrafficManagerEndpointSubnetsGetArgs>? _subnets;

        /// <summary>
        /// One or more `subnet` blocks as defined below
        /// </summary>
        public InputList<Inputs.TrafficManagerEndpointSubnetsGetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.TrafficManagerEndpointSubnetsGetArgs>());
            set => _subnets = value;
        }

        /// <summary>
        /// The FQDN DNS name of the target. This argument must be
        /// provided for an endpoint of type `externalEndpoints`, for other types it
        /// will be computed.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// The resource id of an Azure resource to
        /// target. This argument must be provided for an endpoint of type
        /// `azureEndpoints` or `nestedEndpoints`.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The Endpoint type, must be one of:
        /// - `azureEndpoints`
        /// - `externalEndpoints`
        /// - `nestedEndpoints`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies how much traffic should be distributed to this
        /// endpoint, this must be specified for Profiles using the  `Weighted` traffic
        /// routing method. Supports values between 1 and 1000.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public TrafficManagerEndpointState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class TrafficManagerEndpointCustomHeadersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Traffic Manager endpoint. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public TrafficManagerEndpointCustomHeadersArgs()
        {
        }
    }

    public sealed class TrafficManagerEndpointCustomHeadersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Traffic Manager endpoint. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public TrafficManagerEndpointCustomHeadersGetArgs()
        {
        }
    }

    public sealed class TrafficManagerEndpointSubnetsArgs : Pulumi.ResourceArgs
    {
        [Input("first", required: true)]
        public Input<string> First { get; set; } = null!;

        [Input("last")]
        public Input<string>? Last { get; set; }

        [Input("scope")]
        public Input<int>? Scope { get; set; }

        public TrafficManagerEndpointSubnetsArgs()
        {
        }
    }

    public sealed class TrafficManagerEndpointSubnetsGetArgs : Pulumi.ResourceArgs
    {
        [Input("first", required: true)]
        public Input<string> First { get; set; } = null!;

        [Input("last")]
        public Input<string>? Last { get; set; }

        [Input("scope")]
        public Input<int>? Scope { get; set; }

        public TrafficManagerEndpointSubnetsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class TrafficManagerEndpointCustomHeaders
    {
        /// <summary>
        /// The name of the Traffic Manager endpoint. Changing this forces a
        /// new resource to be created.
        /// </summary>
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private TrafficManagerEndpointCustomHeaders(
            string name,
            string value)
        {
            Name = name;
            Value = value;
        }
    }

    [OutputType]
    public sealed class TrafficManagerEndpointSubnets
    {
        public readonly string First;
        public readonly string? Last;
        public readonly int? Scope;

        [OutputConstructor]
        private TrafficManagerEndpointSubnets(
            string first,
            string? last,
            int? scope)
        {
            First = first;
            Last = last;
            Scope = scope;
        }
    }
    }
}