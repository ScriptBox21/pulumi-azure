// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp
{
    /// <summary>
    /// Manages a NetApp Snapshot.
    /// 
    /// ## NetApp Snapshot Usage
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "netapp",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Name = "Microsoft.Netapp/volumes",
    ///                         Actions = 
    ///                         {
    ///                             "Microsoft.Network/networkinterfaces/*",
    ///                             "Microsoft.Network/virtualNetworks/subnets/join/action",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleAccount = new Azure.NetApp.Account("exampleAccount", new Azure.NetApp.AccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var examplePool = new Azure.NetApp.Pool("examplePool", new Azure.NetApp.PoolArgs
    ///         {
    ///             AccountName = exampleAccount.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ServiceLevel = "Premium",
    ///             SizeInTb = 4,
    ///         });
    ///         var exampleVolume = new Azure.NetApp.Volume("exampleVolume", new Azure.NetApp.VolumeArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AccountName = exampleAccount.Name,
    ///             PoolName = examplePool.Name,
    ///             VolumePath = "my-unique-file-path",
    ///             ServiceLevel = "Premium",
    ///             SubnetId = azurerm_subnet.Test.Id,
    ///             StorageQuotaInGb = 100,
    ///         });
    ///         var exampleSnapshot = new Azure.NetApp.Snapshot("exampleSnapshot", new Azure.NetApp.SnapshotArgs
    ///         {
    ///             AccountName = exampleAccount.Name,
    ///             PoolName = examplePool.Name,
    ///             VolumeName = exampleVolume.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// NetApp Snapshot can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:netapp/snapshot:Snapshot example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/snapshots/snapshot1
    /// ```
    /// </summary>
    [AzureResourceType("azure:netapp/snapshot:Snapshot")]
    public partial class Snapshot : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("poolName")]
        public Output<string> PoolName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("volumeName")]
        public Output<string> VolumeName { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("azure:netapp/snapshot:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
            : base("azure:netapp/snapshot:Snapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, state, options);
        }
    }

    public sealed class SnapshotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("poolName", required: true)]
        public Input<string> PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        [Obsolete(@"This property as been deprecated as the API no longer supports tags and will be removed in version 3.0 of the provider.")]
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("volumeName", required: true)]
        public Input<string> VolumeName { get; set; } = null!;

        public SnapshotArgs()
        {
        }
    }

    public sealed class SnapshotState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the NetApp Snapshot. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("poolName")]
        public Input<string>? PoolName { get; set; }

        /// <summary>
        /// The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        [Obsolete(@"This property as been deprecated as the API no longer supports tags and will be removed in version 3.0 of the provider.")]
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("volumeName")]
        public Input<string>? VolumeName { get; set; }

        public SnapshotState()
        {
        }
    }
}
