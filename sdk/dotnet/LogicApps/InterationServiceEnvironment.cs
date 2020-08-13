// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    /// <summary>
    /// Manages private and isolated Logic App instances within an Azure virtual network.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "westeurope",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/22",
    ///             },
    ///         });
    ///         var isesubnet1 = new Azure.Network.Subnet("isesubnet1", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.0/26",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "integrationServiceEnvironments",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Name = "Microsoft.Logic/integrationServiceEnvironments",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var isesubnet2 = new Azure.Network.Subnet("isesubnet2", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.64/26",
    ///             },
    ///         });
    ///         var isesubnet3 = new Azure.Network.Subnet("isesubnet3", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.128/26",
    ///             },
    ///         });
    ///         var isesubnet4 = new Azure.Network.Subnet("isesubnet4", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.192/26",
    ///             },
    ///         });
    ///         var exampleInterationServiceEnvironment = new Azure.LogicApps.InterationServiceEnvironment("exampleInterationServiceEnvironment", new Azure.LogicApps.InterationServiceEnvironmentArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SkuName = "Developer_0",
    ///             AccessEndpointType = "Internal",
    ///             VirtualNetworkSubnetIds = 
    ///             {
    ///                 isesubnet1.Id,
    ///                 isesubnet2.Id,
    ///                 isesubnet3.Id,
    ///                 isesubnet4.Id,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "environment", "development" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class InterationServiceEnvironment : Pulumi.CustomResource
    {
        /// <summary>
        /// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Output("accessEndpointType")]
        public Output<string> AccessEndpointType { get; private set; } = null!;

        /// <summary>
        /// The list of access endpoint ip addresses of connector.
        /// </summary>
        [Output("connectorEndpointIpAddresses")]
        public Output<ImmutableArray<string>> ConnectorEndpointIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The list of outgoing ip addresses of connector.
        /// </summary>
        [Output("connectorOutboundIpAddresses")]
        public Output<ImmutableArray<string>> ConnectorOutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
        /// </summary>
        [Output("skuName")]
        public Output<string?> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Integration Service Environment.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Output("virtualNetworkSubnetIds")]
        public Output<ImmutableArray<string>> VirtualNetworkSubnetIds { get; private set; } = null!;

        /// <summary>
        /// The list of access endpoint ip addresses of workflow.
        /// </summary>
        [Output("workflowEndpointIpAddresses")]
        public Output<ImmutableArray<string>> WorkflowEndpointIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The list of outgoing ip addresses of workflow.
        /// </summary>
        [Output("workflowOutboundIpAddresses")]
        public Output<ImmutableArray<string>> WorkflowOutboundIpAddresses { get; private set; } = null!;


        /// <summary>
        /// Create a InterationServiceEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterationServiceEnvironment(string name, InterationServiceEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("azure:logicapps/interationServiceEnvironment:InterationServiceEnvironment", name, args ?? new InterationServiceEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InterationServiceEnvironment(string name, Input<string> id, InterationServiceEnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("azure:logicapps/interationServiceEnvironment:InterationServiceEnvironment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InterationServiceEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterationServiceEnvironment Get(string name, Input<string> id, InterationServiceEnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InterationServiceEnvironment(name, id, state, options);
        }
    }

    public sealed class InterationServiceEnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("accessEndpointType", required: true)]
        public Input<string> AccessEndpointType { get; set; } = null!;

        /// <summary>
        /// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Integration Service Environment.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("virtualNetworkSubnetIds", required: true)]
        private InputList<string>? _virtualNetworkSubnetIds;

        /// <summary>
        /// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        public InputList<string> VirtualNetworkSubnetIds
        {
            get => _virtualNetworkSubnetIds ?? (_virtualNetworkSubnetIds = new InputList<string>());
            set => _virtualNetworkSubnetIds = value;
        }

        public InterationServiceEnvironmentArgs()
        {
        }
    }

    public sealed class InterationServiceEnvironmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("accessEndpointType")]
        public Input<string>? AccessEndpointType { get; set; }

        [Input("connectorEndpointIpAddresses")]
        private InputList<string>? _connectorEndpointIpAddresses;

        /// <summary>
        /// The list of access endpoint ip addresses of connector.
        /// </summary>
        public InputList<string> ConnectorEndpointIpAddresses
        {
            get => _connectorEndpointIpAddresses ?? (_connectorEndpointIpAddresses = new InputList<string>());
            set => _connectorEndpointIpAddresses = value;
        }

        [Input("connectorOutboundIpAddresses")]
        private InputList<string>? _connectorOutboundIpAddresses;

        /// <summary>
        /// The list of outgoing ip addresses of connector.
        /// </summary>
        public InputList<string> ConnectorOutboundIpAddresses
        {
            get => _connectorOutboundIpAddresses ?? (_connectorOutboundIpAddresses = new InputList<string>());
            set => _connectorOutboundIpAddresses = value;
        }

        /// <summary>
        /// The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Integration Service Environment.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("virtualNetworkSubnetIds")]
        private InputList<string>? _virtualNetworkSubnetIds;

        /// <summary>
        /// A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
        /// </summary>
        public InputList<string> VirtualNetworkSubnetIds
        {
            get => _virtualNetworkSubnetIds ?? (_virtualNetworkSubnetIds = new InputList<string>());
            set => _virtualNetworkSubnetIds = value;
        }

        [Input("workflowEndpointIpAddresses")]
        private InputList<string>? _workflowEndpointIpAddresses;

        /// <summary>
        /// The list of access endpoint ip addresses of workflow.
        /// </summary>
        public InputList<string> WorkflowEndpointIpAddresses
        {
            get => _workflowEndpointIpAddresses ?? (_workflowEndpointIpAddresses = new InputList<string>());
            set => _workflowEndpointIpAddresses = value;
        }

        [Input("workflowOutboundIpAddresses")]
        private InputList<string>? _workflowOutboundIpAddresses;

        /// <summary>
        /// The list of outgoing ip addresses of workflow.
        /// </summary>
        public InputList<string> WorkflowOutboundIpAddresses
        {
            get => _workflowOutboundIpAddresses ?? (_workflowOutboundIpAddresses = new InputList<string>());
            set => _workflowOutboundIpAddresses = value;
        }

        public InterationServiceEnvironmentState()
        {
        }
    }
}
