// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a subnet. Subnets represent network segments within the IP space defined by the virtual network.
    /// 
    /// &gt; **NOTE on Virtual Networks and Subnet's:** This provider currently
    /// provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
    /// At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/subnet.html.markdown.
    /// </summary>
    public partial class Subnet : Pulumi.CustomResource
    {
        /// <summary>
        /// The address prefix to use for the subnet.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// One or more `delegation` blocks as defined below.
        /// </summary>
        [Output("delegations")]
        public Output<ImmutableArray<Outputs.SubnetDelegations>> Delegations { get; private set; } = null!;

        /// <summary>
        /// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
        /// </summary>
        [Output("enforcePrivateLinkEndpointNetworkPolicies")]
        public Output<bool?> EnforcePrivateLinkEndpointNetworkPolicies { get; private set; } = null!;

        /// <summary>
        /// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforce_private_link_endpoint_network_policies`.
        /// </summary>
        [Output("enforcePrivateLinkServiceNetworkPolicies")]
        public Output<bool?> EnforcePrivateLinkServiceNetworkPolicies { get; private set; } = null!;

        /// <summary>
        /// The name of the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
        /// </summary>
        [Output("serviceEndpoints")]
        public Output<ImmutableArray<string>> ServiceEndpoints { get; private set; } = null!;

        /// <summary>
        /// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualNetworkName")]
        public Output<string> VirtualNetworkName { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("azure:network/subnet:Subnet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/subnet:Subnet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subnet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnet Get(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
        {
            return new Subnet(name, id, state, options);
        }
    }

    public sealed class SubnetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address prefix to use for the subnet.
        /// </summary>
        [Input("addressPrefix", required: true)]
        public Input<string> AddressPrefix { get; set; } = null!;

        [Input("delegations")]
        private InputList<Inputs.SubnetDelegationsArgs>? _delegations;

        /// <summary>
        /// One or more `delegation` blocks as defined below.
        /// </summary>
        public InputList<Inputs.SubnetDelegationsArgs> Delegations
        {
            get => _delegations ?? (_delegations = new InputList<Inputs.SubnetDelegationsArgs>());
            set => _delegations = value;
        }

        /// <summary>
        /// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
        /// </summary>
        [Input("enforcePrivateLinkEndpointNetworkPolicies")]
        public Input<bool>? EnforcePrivateLinkEndpointNetworkPolicies { get; set; }

        /// <summary>
        /// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforce_private_link_endpoint_network_policies`.
        /// </summary>
        [Input("enforcePrivateLinkServiceNetworkPolicies")]
        public Input<bool>? EnforcePrivateLinkServiceNetworkPolicies { get; set; }

        /// <summary>
        /// The name of the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("serviceEndpoints")]
        private InputList<string>? _serviceEndpoints;

        /// <summary>
        /// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
        /// </summary>
        public InputList<string> ServiceEndpoints
        {
            get => _serviceEndpoints ?? (_serviceEndpoints = new InputList<string>());
            set => _serviceEndpoints = value;
        }

        /// <summary>
        /// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualNetworkName", required: true)]
        public Input<string> VirtualNetworkName { get; set; } = null!;

        public SubnetArgs()
        {
        }
    }

    public sealed class SubnetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address prefix to use for the subnet.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        [Input("delegations")]
        private InputList<Inputs.SubnetDelegationsGetArgs>? _delegations;

        /// <summary>
        /// One or more `delegation` blocks as defined below.
        /// </summary>
        public InputList<Inputs.SubnetDelegationsGetArgs> Delegations
        {
            get => _delegations ?? (_delegations = new InputList<Inputs.SubnetDelegationsGetArgs>());
            set => _delegations = value;
        }

        /// <summary>
        /// Enable or Disable network policies for the private link endpoint on the subnet. Default value is `false`. Conflicts with enforce_private_link_service_network_policies.
        /// </summary>
        [Input("enforcePrivateLinkEndpointNetworkPolicies")]
        public Input<bool>? EnforcePrivateLinkEndpointNetworkPolicies { get; set; }

        /// <summary>
        /// Enable or Disable network policies for the private link service on the subnet. Default valule is `false`. Conflicts with `enforce_private_link_endpoint_network_policies`.
        /// </summary>
        [Input("enforcePrivateLinkServiceNetworkPolicies")]
        public Input<bool>? EnforcePrivateLinkServiceNetworkPolicies { get; set; }

        /// <summary>
        /// The name of the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("serviceEndpoints")]
        private InputList<string>? _serviceEndpoints;

        /// <summary>
        /// The list of Service endpoints to associate with the subnet. Possible values include: `Microsoft.AzureActiveDirectory`, `Microsoft.AzureCosmosDB`, `Microsoft.ContainerRegistry`, `Microsoft.EventHub`, `Microsoft.KeyVault`, `Microsoft.ServiceBus`, `Microsoft.Sql`, `Microsoft.Storage` and `Microsoft.Web`.
        /// </summary>
        public InputList<string> ServiceEndpoints
        {
            get => _serviceEndpoints ?? (_serviceEndpoints = new InputList<string>());
            set => _serviceEndpoints = value;
        }

        /// <summary>
        /// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualNetworkName")]
        public Input<string>? VirtualNetworkName { get; set; }

        public SubnetState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SubnetDelegationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for this delegation.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// A `service_delegation` block as defined below.
        /// </summary>
        [Input("serviceDelegation", required: true)]
        public Input<SubnetDelegationsServiceDelegationArgs> ServiceDelegation { get; set; } = null!;

        public SubnetDelegationsArgs()
        {
        }
    }

    public sealed class SubnetDelegationsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for this delegation.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// A `service_delegation` block as defined below.
        /// </summary>
        [Input("serviceDelegation", required: true)]
        public Input<SubnetDelegationsServiceDelegationGetArgs> ServiceDelegation { get; set; } = null!;

        public SubnetDelegationsGetArgs()
        {
        }
    }

    public sealed class SubnetDelegationsServiceDelegationArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;

        /// <summary>
        /// A list of Actions which should be delegated. This list is specific to the service to delegate to. Possible values include `Microsoft.Network/networkinterfaces/*`, `Microsoft.Network/virtualNetworks/subnets/action`, `Microsoft.Network/virtualNetworks/subnets/join/action`, `Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action` and `Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action`.
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        /// <summary>
        /// The name of service to delegate to. Possible values include `Microsoft.BareMetal/AzureVMware`, `Microsoft.BareMetal/CrayServers`, `Microsoft.Batch/batchAccounts`, `Microsoft.ContainerInstance/containerGroups`, `Microsoft.Databricks/workspaces`, `Microsoft.DBforPostgreSQL/serversv2`, `Microsoft.HardwareSecurityModules/dedicatedHSMs`, `Microsoft.Logic/integrationServiceEnvironments`, `Microsoft.Netapp/volumes`, `Microsoft.ServiceFabricMesh/networks`, `Microsoft.Sql/managedInstances`, `Microsoft.Sql/servers`, `Microsoft.StreamAnalytics/streamingJobs`, `Microsoft.Web/hostingEnvironments` and `Microsoft.Web/serverFarms`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public SubnetDelegationsServiceDelegationArgs()
        {
        }
    }

    public sealed class SubnetDelegationsServiceDelegationGetArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;

        /// <summary>
        /// A list of Actions which should be delegated. This list is specific to the service to delegate to. Possible values include `Microsoft.Network/networkinterfaces/*`, `Microsoft.Network/virtualNetworks/subnets/action`, `Microsoft.Network/virtualNetworks/subnets/join/action`, `Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action` and `Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action`.
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        /// <summary>
        /// The name of service to delegate to. Possible values include `Microsoft.BareMetal/AzureVMware`, `Microsoft.BareMetal/CrayServers`, `Microsoft.Batch/batchAccounts`, `Microsoft.ContainerInstance/containerGroups`, `Microsoft.Databricks/workspaces`, `Microsoft.DBforPostgreSQL/serversv2`, `Microsoft.HardwareSecurityModules/dedicatedHSMs`, `Microsoft.Logic/integrationServiceEnvironments`, `Microsoft.Netapp/volumes`, `Microsoft.ServiceFabricMesh/networks`, `Microsoft.Sql/managedInstances`, `Microsoft.Sql/servers`, `Microsoft.StreamAnalytics/streamingJobs`, `Microsoft.Web/hostingEnvironments` and `Microsoft.Web/serverFarms`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public SubnetDelegationsServiceDelegationGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SubnetDelegations
    {
        /// <summary>
        /// A name for this delegation.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A `service_delegation` block as defined below.
        /// </summary>
        public readonly SubnetDelegationsServiceDelegation ServiceDelegation;

        [OutputConstructor]
        private SubnetDelegations(
            string name,
            SubnetDelegationsServiceDelegation serviceDelegation)
        {
            Name = name;
            ServiceDelegation = serviceDelegation;
        }
    }

    [OutputType]
    public sealed class SubnetDelegationsServiceDelegation
    {
        /// <summary>
        /// A list of Actions which should be delegated. This list is specific to the service to delegate to. Possible values include `Microsoft.Network/networkinterfaces/*`, `Microsoft.Network/virtualNetworks/subnets/action`, `Microsoft.Network/virtualNetworks/subnets/join/action`, `Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action` and `Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action`.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// The name of service to delegate to. Possible values include `Microsoft.BareMetal/AzureVMware`, `Microsoft.BareMetal/CrayServers`, `Microsoft.Batch/batchAccounts`, `Microsoft.ContainerInstance/containerGroups`, `Microsoft.Databricks/workspaces`, `Microsoft.DBforPostgreSQL/serversv2`, `Microsoft.HardwareSecurityModules/dedicatedHSMs`, `Microsoft.Logic/integrationServiceEnvironments`, `Microsoft.Netapp/volumes`, `Microsoft.ServiceFabricMesh/networks`, `Microsoft.Sql/managedInstances`, `Microsoft.Sql/servers`, `Microsoft.StreamAnalytics/streamingJobs`, `Microsoft.Web/hostingEnvironments` and `Microsoft.Web/serverFarms`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private SubnetDelegationsServiceDelegation(
            ImmutableArray<string> actions,
            string name)
        {
            Actions = actions;
            Name = name;
        }
    }
    }
}
