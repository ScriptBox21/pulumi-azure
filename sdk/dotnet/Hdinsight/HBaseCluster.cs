// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight
{
    /// <summary>
    /// Manages a HDInsight HBase Cluster.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/hdinsight_hbase_cluster.html.markdown.
    /// </summary>
    public partial class HBaseCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterVersion")]
        public Output<string> ClusterVersion { get; private set; } = null!;

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Output("componentVersion")]
        public Output<Outputs.HBaseClusterComponentVersion> ComponentVersion { get; private set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.HBaseClusterGateway> Gateway { get; private set; } = null!;

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
        /// </summary>
        [Output("httpsEndpoint")]
        public Output<string> HttpsEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Output("roles")]
        public Output<Outputs.HBaseClusterRoles> Roles { get; private set; } = null!;

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
        /// </summary>
        [Output("sshEndpoint")]
        public Output<string> SshEndpoint { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.HBaseClusterStorageAccounts>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Output("storageAccountGen2")]
        public Output<Outputs.HBaseClusterStorageAccountGen2?> StorageAccountGen2 { get; private set; } = null!;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight HBase Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a HBaseCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HBaseCluster(string name, HBaseClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:hdinsight/hBaseCluster:HBaseCluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private HBaseCluster(string name, Input<string> id, HBaseClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:hdinsight/hBaseCluster:HBaseCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HBaseCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HBaseCluster Get(string name, Input<string> id, HBaseClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new HBaseCluster(name, id, state, options);
        }
    }

    public sealed class HBaseClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion", required: true)]
        public Input<string> ClusterVersion { get; set; } = null!;

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Input("componentVersion", required: true)]
        public Input<Inputs.HBaseClusterComponentVersionArgs> ComponentVersion { get; set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<Inputs.HBaseClusterGatewayArgs> Gateway { get; set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles", required: true)]
        public Input<Inputs.HBaseClusterRolesArgs> Roles { get; set; } = null!;

        [Input("storageAccounts")]
        private InputList<Inputs.HBaseClusterStorageAccountsArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.HBaseClusterStorageAccountsArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.HBaseClusterStorageAccountsArgs>());
            set => _storageAccounts = value;
        }

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.HBaseClusterStorageAccountGen2Args>? StorageAccountGen2 { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight HBase Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public HBaseClusterArgs()
        {
        }
    }

    public sealed class HBaseClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// A `component_version` block as defined below.
        /// </summary>
        [Input("componentVersion")]
        public Input<Inputs.HBaseClusterComponentVersionGetArgs>? ComponentVersion { get; set; }

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.HBaseClusterGatewayGetArgs>? Gateway { get; set; }

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
        /// </summary>
        [Input("httpsEndpoint")]
        public Input<string>? HttpsEndpoint { get; set; }

        /// <summary>
        /// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles")]
        public Input<Inputs.HBaseClusterRolesGetArgs>? Roles { get; set; }

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
        /// </summary>
        [Input("sshEndpoint")]
        public Input<string>? SshEndpoint { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.HBaseClusterStorageAccountsGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.HBaseClusterStorageAccountsGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.HBaseClusterStorageAccountsGetArgs>());
            set => _storageAccounts = value;
        }

        /// <summary>
        /// A `storage_account_gen2` block as defined below.
        /// </summary>
        [Input("storageAccountGen2")]
        public Input<Inputs.HBaseClusterStorageAccountGen2GetArgs>? StorageAccountGen2 { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight HBase Cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public HBaseClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class HBaseClusterComponentVersionArgs : Pulumi.ResourceArgs
    {
        [Input("hbase", required: true)]
        public Input<string> Hbase { get; set; } = null!;

        public HBaseClusterComponentVersionArgs()
        {
        }
    }

    public sealed class HBaseClusterComponentVersionGetArgs : Pulumi.ResourceArgs
    {
        [Input("hbase", required: true)]
        public Input<string> Hbase { get; set; } = null!;

        public HBaseClusterComponentVersionGetArgs()
        {
        }
    }

    public sealed class HBaseClusterGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public HBaseClusterGatewayArgs()
        {
        }
    }

    public sealed class HBaseClusterGatewayGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public HBaseClusterGatewayGetArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesArgs : Pulumi.ResourceArgs
    {
        [Input("headNode", required: true)]
        public Input<HBaseClusterRolesHeadNodeArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<HBaseClusterRolesWorkerNodeArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<HBaseClusterRolesZookeeperNodeArgs> ZookeeperNode { get; set; } = null!;

        public HBaseClusterRolesArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesGetArgs : Pulumi.ResourceArgs
    {
        [Input("headNode", required: true)]
        public Input<HBaseClusterRolesHeadNodeGetArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<HBaseClusterRolesWorkerNodeGetArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<HBaseClusterRolesZookeeperNodeGetArgs> ZookeeperNode { get; set; } = null!;

        public HBaseClusterRolesGetArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesHeadNodeArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public HBaseClusterRolesHeadNodeArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesHeadNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public HBaseClusterRolesHeadNodeGetArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesWorkerNodeArgs : Pulumi.ResourceArgs
    {
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("targetInstanceCount", required: true)]
        public Input<int> TargetInstanceCount { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public HBaseClusterRolesWorkerNodeArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesWorkerNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("targetInstanceCount", required: true)]
        public Input<int> TargetInstanceCount { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public HBaseClusterRolesWorkerNodeGetArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesZookeeperNodeArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public HBaseClusterRolesZookeeperNodeArgs()
        {
        }
    }

    public sealed class HBaseClusterRolesZookeeperNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public HBaseClusterRolesZookeeperNodeGetArgs()
        {
        }
    }

    public sealed class HBaseClusterStorageAccountGen2Args : Pulumi.ResourceArgs
    {
        [Input("filesystemId", required: true)]
        public Input<string> FilesystemId { get; set; } = null!;

        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("managedIdentityResourceId", required: true)]
        public Input<string> ManagedIdentityResourceId { get; set; } = null!;

        [Input("storageResourceId", required: true)]
        public Input<string> StorageResourceId { get; set; } = null!;

        public HBaseClusterStorageAccountGen2Args()
        {
        }
    }

    public sealed class HBaseClusterStorageAccountGen2GetArgs : Pulumi.ResourceArgs
    {
        [Input("filesystemId", required: true)]
        public Input<string> FilesystemId { get; set; } = null!;

        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("managedIdentityResourceId", required: true)]
        public Input<string> ManagedIdentityResourceId { get; set; } = null!;

        [Input("storageResourceId", required: true)]
        public Input<string> StorageResourceId { get; set; } = null!;

        public HBaseClusterStorageAccountGen2GetArgs()
        {
        }
    }

    public sealed class HBaseClusterStorageAccountsArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public HBaseClusterStorageAccountsArgs()
        {
        }
    }

    public sealed class HBaseClusterStorageAccountsGetArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public HBaseClusterStorageAccountsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class HBaseClusterComponentVersion
    {
        public readonly string Hbase;

        [OutputConstructor]
        private HBaseClusterComponentVersion(string hbase)
        {
            Hbase = hbase;
        }
    }

    [OutputType]
    public sealed class HBaseClusterGateway
    {
        public readonly bool Enabled;
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private HBaseClusterGateway(
            bool enabled,
            string password,
            string username)
        {
            Enabled = enabled;
            Password = password;
            Username = username;
        }
    }

    [OutputType]
    public sealed class HBaseClusterRoles
    {
        public readonly HBaseClusterRolesHeadNode HeadNode;
        public readonly HBaseClusterRolesWorkerNode WorkerNode;
        public readonly HBaseClusterRolesZookeeperNode ZookeeperNode;

        [OutputConstructor]
        private HBaseClusterRoles(
            HBaseClusterRolesHeadNode headNode,
            HBaseClusterRolesWorkerNode workerNode,
            HBaseClusterRolesZookeeperNode zookeeperNode)
        {
            HeadNode = headNode;
            WorkerNode = workerNode;
            ZookeeperNode = zookeeperNode;
        }
    }

    [OutputType]
    public sealed class HBaseClusterRolesHeadNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private HBaseClusterRolesHeadNode(
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class HBaseClusterRolesWorkerNode
    {
        public readonly int? MinInstanceCount;
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly int TargetInstanceCount;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private HBaseClusterRolesWorkerNode(
            int? minInstanceCount,
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            int targetInstanceCount,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            MinInstanceCount = minInstanceCount;
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            TargetInstanceCount = targetInstanceCount;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class HBaseClusterRolesZookeeperNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private HBaseClusterRolesZookeeperNode(
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class HBaseClusterStorageAccountGen2
    {
        public readonly string FilesystemId;
        public readonly bool IsDefault;
        public readonly string ManagedIdentityResourceId;
        public readonly string StorageResourceId;

        [OutputConstructor]
        private HBaseClusterStorageAccountGen2(
            string filesystemId,
            bool isDefault,
            string managedIdentityResourceId,
            string storageResourceId)
        {
            FilesystemId = filesystemId;
            IsDefault = isDefault;
            ManagedIdentityResourceId = managedIdentityResourceId;
            StorageResourceId = storageResourceId;
        }
    }

    [OutputType]
    public sealed class HBaseClusterStorageAccounts
    {
        public readonly bool IsDefault;
        public readonly string StorageAccountKey;
        public readonly string StorageContainerId;

        [OutputConstructor]
        private HBaseClusterStorageAccounts(
            bool isDefault,
            string storageAccountKey,
            string storageContainerId)
        {
            IsDefault = isDefault;
            StorageAccountKey = storageAccountKey;
            StorageContainerId = storageContainerId;
        }
    }
    }
}
