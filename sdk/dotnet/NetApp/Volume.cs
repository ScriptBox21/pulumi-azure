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
    /// Manages a NetApp Volume.
    /// 
    /// ## NetApp Volume Usage
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
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
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
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AccountName = exampleAccount.Name,
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
    ///             SubnetId = exampleSubnet.Id,
    ///             Protocols = 
    ///             {
    ///                 "NFSv4.1",
    ///             },
    ///             StorageQuotaInGb = 100,
    ///             DataProtectionReplication = new Azure.NetApp.Inputs.VolumeDataProtectionReplicationArgs
    ///             {
    ///                 EndpointType = "dst",
    ///                 RemoteVolumeLocation = azurerm_resource_group.Example_primary.Location,
    ///                 RemoteVolumeResourceId = azurerm_netapp_volume.Example_primary.Id,
    ///                 ReplicationFrequency = "_10minutely",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// NetApp Volumes can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:netapp/volume:Volume example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1
    /// ```
    /// </summary>
    [AzureResourceType("azure:netapp/volume:Volume")]
    public partial class Volume : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        [Output("dataProtectionReplication")]
        public Output<Outputs.VolumeDataProtectionReplication?> DataProtectionReplication { get; private set; } = null!;

        /// <summary>
        /// One or more `export_policy_rule` block defined below.
        /// </summary>
        [Output("exportPolicyRules")]
        public Output<ImmutableArray<Outputs.VolumeExportPolicyRule>> ExportPolicyRules { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A list of IPv4 Addresses which should be used to mount the volume.
        /// </summary>
        [Output("mountIpAddresses")]
        public Output<ImmutableArray<string>> MountIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the NetApp Volume. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("poolName")]
        public Output<string> PoolName { get; private set; } = null!;

        /// <summary>
        /// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        /// </summary>
        [Output("serviceLevel")]
        public Output<string> ServiceLevel { get; private set; } = null!;

        /// <summary>
        /// The maximum Storage Quota allowed for a file system in Gigabytes.
        /// </summary>
        [Output("storageQuotaInGb")]
        public Output<int> StorageQuotaInGb { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
        /// </summary>
        [Output("volumePath")]
        public Output<string> VolumePath { get; private set; } = null!;


        /// <summary>
        /// Create a Volume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Volume(string name, VolumeArgs args, CustomResourceOptions? options = null)
            : base("azure:netapp/volume:Volume", name, args ?? new VolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Volume(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
            : base("azure:netapp/volume:Volume", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Volume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Volume Get(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new Volume(name, id, state, options);
        }
    }

    public sealed class VolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("dataProtectionReplication")]
        public Input<Inputs.VolumeDataProtectionReplicationArgs>? DataProtectionReplication { get; set; }

        [Input("exportPolicyRules")]
        private InputList<Inputs.VolumeExportPolicyRuleArgs>? _exportPolicyRules;

        /// <summary>
        /// One or more `export_policy_rule` block defined below.
        /// </summary>
        public InputList<Inputs.VolumeExportPolicyRuleArgs> ExportPolicyRules
        {
            get => _exportPolicyRules ?? (_exportPolicyRules = new InputList<Inputs.VolumeExportPolicyRuleArgs>());
            set => _exportPolicyRules = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the NetApp Volume. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("poolName", required: true)]
        public Input<string> PoolName { get; set; } = null!;

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        /// </summary>
        [Input("serviceLevel", required: true)]
        public Input<string> ServiceLevel { get; set; } = null!;

        /// <summary>
        /// The maximum Storage Quota allowed for a file system in Gigabytes.
        /// </summary>
        [Input("storageQuotaInGb", required: true)]
        public Input<int> StorageQuotaInGb { get; set; } = null!;

        /// <summary>
        /// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

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
        /// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
        /// </summary>
        [Input("volumePath", required: true)]
        public Input<string> VolumePath { get; set; } = null!;

        public VolumeArgs()
        {
        }
    }

    public sealed class VolumeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("dataProtectionReplication")]
        public Input<Inputs.VolumeDataProtectionReplicationGetArgs>? DataProtectionReplication { get; set; }

        [Input("exportPolicyRules")]
        private InputList<Inputs.VolumeExportPolicyRuleGetArgs>? _exportPolicyRules;

        /// <summary>
        /// One or more `export_policy_rule` block defined below.
        /// </summary>
        public InputList<Inputs.VolumeExportPolicyRuleGetArgs> ExportPolicyRules
        {
            get => _exportPolicyRules ?? (_exportPolicyRules = new InputList<Inputs.VolumeExportPolicyRuleGetArgs>());
            set => _exportPolicyRules = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("mountIpAddresses")]
        private InputList<string>? _mountIpAddresses;

        /// <summary>
        /// A list of IPv4 Addresses which should be used to mount the volume.
        /// </summary>
        public InputList<string> MountIpAddresses
        {
            get => _mountIpAddresses ?? (_mountIpAddresses = new InputList<string>());
            set => _mountIpAddresses = value;
        }

        /// <summary>
        /// The name of the NetApp Volume. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("poolName")]
        public Input<string>? PoolName { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        /// </summary>
        [Input("serviceLevel")]
        public Input<string>? ServiceLevel { get; set; }

        /// <summary>
        /// The maximum Storage Quota allowed for a file system in Gigabytes.
        /// </summary>
        [Input("storageQuotaInGb")]
        public Input<int>? StorageQuotaInGb { get; set; }

        /// <summary>
        /// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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
        /// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
        /// </summary>
        [Input("volumePath")]
        public Input<string>? VolumePath { get; set; }

        public VolumeState()
        {
        }
    }
}
