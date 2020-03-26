// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Batch pool
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/batch_pool.html.markdown.
        /// </summary>
        [Obsolete("Use GetPool.InvokeAsync() instead")]
        public static Task<GetPoolResult> GetPool(GetPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPoolResult>("azure:batch/getPool:getPool", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetPool
    {
        /// <summary>
        /// Use this data source to access information about an existing Batch pool
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/batch_pool.html.markdown.
        /// </summary>
        public static Task<GetPoolResult> InvokeAsync(GetPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPoolResult>("azure:batch/getPool:getPool", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        [Input("certificates")]
        private List<Inputs.GetPoolCertificatesArgs>? _certificates;

        /// <summary>
        /// One or more `certificate` blocks that describe the certificates installed on each compute node in the pool.
        /// </summary>
        public List<Inputs.GetPoolCertificatesArgs> Certificates
        {
            get => _certificates ?? (_certificates = new List<Inputs.GetPoolCertificatesArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("networkConfiguration")]
        public Inputs.GetPoolNetworkConfigurationArgs? NetworkConfiguration { get; set; }

        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `start_task` block that describes the start task settings for the Batch pool.
        /// </summary>
        [Input("startTask")]
        public Inputs.GetPoolStartTaskArgs? StartTask { get; set; }

        public GetPoolArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPoolResult
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// A `auto_scale` block that describes the scale settings when using auto scale.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolAutoScalesResult> AutoScales;
        /// <summary>
        /// One or more `certificate` blocks that describe the certificates installed on each compute node in the pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolCertificatesResult> Certificates;
        /// <summary>
        /// The container configuration used in the pool's VMs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolContainerConfigurationsResult> ContainerConfigurations;
        public readonly string DisplayName;
        /// <summary>
        /// A `fixed_scale` block that describes the scale settings when using fixed scale.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolFixedScalesResult> FixedScales;
        /// <summary>
        /// The maximum number of tasks that can run concurrently on a single compute node in the pool.
        /// </summary>
        public readonly int MaxTasksPerNode;
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.GetPoolNetworkConfigurationResult NetworkConfiguration;
        /// <summary>
        /// The Sku of the node agents in the Batch pool.
        /// </summary>
        public readonly string NodeAgentSkuId;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A `start_task` block that describes the start task settings for the Batch pool.
        /// </summary>
        public readonly Outputs.GetPoolStartTaskResult? StartTask;
        /// <summary>
        /// The reference of the storage image used by the nodes in the Batch pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolStorageImageReferencesResult> StorageImageReferences;
        /// <summary>
        /// The size of the VM created in the Batch pool.
        /// </summary>
        public readonly string VmSize;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPoolResult(
            string accountName,
            ImmutableArray<Outputs.GetPoolAutoScalesResult> autoScales,
            ImmutableArray<Outputs.GetPoolCertificatesResult> certificates,
            ImmutableArray<Outputs.GetPoolContainerConfigurationsResult> containerConfigurations,
            string displayName,
            ImmutableArray<Outputs.GetPoolFixedScalesResult> fixedScales,
            int maxTasksPerNode,
            ImmutableDictionary<string, string> metadata,
            string name,
            Outputs.GetPoolNetworkConfigurationResult networkConfiguration,
            string nodeAgentSkuId,
            string resourceGroupName,
            Outputs.GetPoolStartTaskResult? startTask,
            ImmutableArray<Outputs.GetPoolStorageImageReferencesResult> storageImageReferences,
            string vmSize,
            string id)
        {
            AccountName = accountName;
            AutoScales = autoScales;
            Certificates = certificates;
            ContainerConfigurations = containerConfigurations;
            DisplayName = displayName;
            FixedScales = fixedScales;
            MaxTasksPerNode = maxTasksPerNode;
            Metadata = metadata;
            Name = name;
            NetworkConfiguration = networkConfiguration;
            NodeAgentSkuId = nodeAgentSkuId;
            ResourceGroupName = resourceGroupName;
            StartTask = startTask;
            StorageImageReferences = storageImageReferences;
            VmSize = vmSize;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetPoolCertificatesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The fully qualified ID of the certificate installed on the pool.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The location of the certificate store on the compute node into which the certificate is installed, either `CurrentUser` or `LocalMachine`.
        /// </summary>
        [Input("storeLocation", required: true)]
        public string StoreLocation { get; set; } = null!;

        /// <summary>
        /// The name of the certificate store on the compute node into which the certificate is installed.
        /// </summary>
        [Input("storeName")]
        public string? StoreName { get; set; }

        [Input("visibilities")]
        private List<string>? _visibilities;

        /// <summary>
        /// Which user accounts on the compute node have access to the private data of the certificate.
        /// </summary>
        public List<string> Visibilities
        {
            get => _visibilities ?? (_visibilities = new List<string>());
            set => _visibilities = value;
        }

        public GetPoolCertificatesArgs()
        {
        }
    }

    public sealed class GetPoolNetworkConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The inbound NAT pools that are used to address specific ports on the individual compute node externally.
        /// </summary>
        [Input("endpointConfiguration")]
        public GetPoolNetworkConfigurationEndpointConfigurationArgs? EndpointConfiguration { get; set; }

        /// <summary>
        /// The ARM resource identifier of the virtual network subnet which the compute nodes of the pool are joined too.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        public GetPoolNetworkConfigurationArgs()
        {
        }
    }

    public sealed class GetPoolNetworkConfigurationEndpointConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The port number on the compute node.
        /// </summary>
        [Input("backendPort")]
        public int? BackendPort { get; set; }

        /// <summary>
        /// The range of external ports that are used to provide inbound access to the backendPort on the individual compute nodes in the format of `1000-1100`.
        /// </summary>
        [Input("frontendPortRange")]
        public string? FrontendPortRange { get; set; }

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("networkSecurityGroupRules")]
        private List<GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesArgs>? _networkSecurityGroupRules;

        /// <summary>
        /// The list of network security group rules that are applied to the endpoint.
        /// </summary>
        public List<GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesArgs> NetworkSecurityGroupRules
        {
            get => _networkSecurityGroupRules ?? (_networkSecurityGroupRules = new List<GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesArgs>());
            set => _networkSecurityGroupRules = value;
        }

        /// <summary>
        /// The protocol of the endpoint.
        /// </summary>
        [Input("protocol")]
        public string? Protocol { get; set; }

        public GetPoolNetworkConfigurationEndpointConfigurationArgs()
        {
        }
    }

    public sealed class GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The action that should be taken for a specified IP address, subnet range or tag.
        /// </summary>
        [Input("access")]
        public string? Access { get; set; }

        /// <summary>
        /// The priority for this rule.
        /// </summary>
        [Input("priority")]
        public int? Priority { get; set; }

        /// <summary>
        /// The source address prefix or tag to match for the rule.
        /// </summary>
        [Input("sourceAddressPrefix")]
        public string? SourceAddressPrefix { get; set; }

        public GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesArgs()
        {
        }
    }

    public sealed class GetPoolStartTaskArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The command line executed by the start task.
        /// </summary>
        [Input("commandLine", required: true)]
        public string CommandLine { get; set; } = null!;

        [Input("environment")]
        private Dictionary<string, string>? _environment;

        /// <summary>
        /// A map of strings (key,value) that represents the environment variables to set in the start task.
        /// </summary>
        public Dictionary<string, string> Environment
        {
            get => _environment ?? (_environment = new Dictionary<string, string>());
            set => _environment = value;
        }

        /// <summary>
        /// The number of retry count.
        /// </summary>
        [Input("maxTaskRetryCount")]
        public int? MaxTaskRetryCount { get; set; }

        [Input("resourceFiles")]
        private List<GetPoolStartTaskResourceFilesArgs>? _resourceFiles;

        /// <summary>
        /// One or more `resource_file` blocks that describe the files to be downloaded to a compute node.
        /// </summary>
        public List<GetPoolStartTaskResourceFilesArgs> ResourceFiles
        {
            get => _resourceFiles ?? (_resourceFiles = new List<GetPoolStartTaskResourceFilesArgs>());
            set => _resourceFiles = value;
        }

        [Input("userIdentities")]
        private List<GetPoolStartTaskUserIdentitiesArgs>? _userIdentities;

        /// <summary>
        /// A `user_identity` block that describes the user identity under which the start task runs.
        /// </summary>
        public List<GetPoolStartTaskUserIdentitiesArgs> UserIdentities
        {
            get => _userIdentities ?? (_userIdentities = new List<GetPoolStartTaskUserIdentitiesArgs>());
            set => _userIdentities = value;
        }

        /// <summary>
        /// A flag that indicates if the Batch pool should wait for the start task to be completed.
        /// </summary>
        [Input("waitForSuccess")]
        public bool? WaitForSuccess { get; set; }

        public GetPoolStartTaskArgs()
        {
        }
    }

    public sealed class GetPoolStartTaskResourceFilesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The storage container name in the auto storage account.
        /// </summary>
        [Input("autoStorageContainerName")]
        public string? AutoStorageContainerName { get; set; }

        /// <summary>
        /// The blob prefix used when downloading blobs from an Azure Storage container.
        /// </summary>
        [Input("blobPrefix")]
        public string? BlobPrefix { get; set; }

        /// <summary>
        /// The file permission mode attribute represented as a string in octal format (e.g. `"0644"`).
        /// </summary>
        [Input("fileMode")]
        public string? FileMode { get; set; }

        /// <summary>
        /// The location on the compute node to which to download the file, relative to the task's working directory. If the `http_url` property is specified, the `file_path` is required and describes the path which the file will be downloaded to, including the filename. Otherwise, if the `auto_storage_container_name` or `storage_container_url` property is specified.
        /// </summary>
        [Input("filePath")]
        public string? FilePath { get; set; }

        /// <summary>
        /// The URL of the file to download. If the URL is Azure Blob Storage, it must be readable using anonymous access.
        /// </summary>
        [Input("httpUrl")]
        public string? HttpUrl { get; set; }

        /// <summary>
        /// The URL of the blob container within Azure Blob Storage.
        /// </summary>
        [Input("storageContainerUrl")]
        public string? StorageContainerUrl { get; set; }

        public GetPoolStartTaskResourceFilesArgs()
        {
        }
    }

    public sealed class GetPoolStartTaskUserIdentitiesArgs : Pulumi.InvokeArgs
    {
        [Input("autoUsers")]
        private List<GetPoolStartTaskUserIdentitiesAutoUsersArgs>? _autoUsers;

        /// <summary>
        /// A `auto_user` block that describes the user identity under which the start task runs.
        /// </summary>
        public List<GetPoolStartTaskUserIdentitiesAutoUsersArgs> AutoUsers
        {
            get => _autoUsers ?? (_autoUsers = new List<GetPoolStartTaskUserIdentitiesAutoUsersArgs>());
            set => _autoUsers = value;
        }

        /// <summary>
        /// The user name to log into the registry server.
        /// </summary>
        [Input("userName")]
        public string? UserName { get; set; }

        public GetPoolStartTaskUserIdentitiesArgs()
        {
        }
    }

    public sealed class GetPoolStartTaskUserIdentitiesAutoUsersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The elevation level of the user identity under which the start task runs.
        /// </summary>
        [Input("elevationLevel")]
        public string? ElevationLevel { get; set; }

        /// <summary>
        /// The scope of the user identity under which the start task runs.
        /// </summary>
        [Input("scope")]
        public string? Scope { get; set; }

        public GetPoolStartTaskUserIdentitiesAutoUsersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetPoolAutoScalesResult
    {
        /// <summary>
        /// The interval to wait before evaluating if the pool needs to be scaled.
        /// </summary>
        public readonly string EvaluationInterval;
        /// <summary>
        /// The autoscale formula that needs to be used for scaling the Batch pool.
        /// </summary>
        public readonly string Formula;

        [OutputConstructor]
        private GetPoolAutoScalesResult(
            string evaluationInterval,
            string formula)
        {
            EvaluationInterval = evaluationInterval;
            Formula = formula;
        }
    }

    [OutputType]
    public sealed class GetPoolCertificatesResult
    {
        /// <summary>
        /// The fully qualified ID of the certificate installed on the pool.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the certificate store on the compute node into which the certificate is installed, either `CurrentUser` or `LocalMachine`.
        /// </summary>
        public readonly string StoreLocation;
        /// <summary>
        /// The name of the certificate store on the compute node into which the certificate is installed.
        /// </summary>
        public readonly string? StoreName;
        /// <summary>
        /// Which user accounts on the compute node have access to the private data of the certificate.
        /// </summary>
        public readonly ImmutableArray<string> Visibilities;

        [OutputConstructor]
        private GetPoolCertificatesResult(
            string id,
            string storeLocation,
            string? storeName,
            ImmutableArray<string> visibilities)
        {
            Id = id;
            StoreLocation = storeLocation;
            StoreName = storeName;
            Visibilities = visibilities;
        }
    }

    [OutputType]
    public sealed class GetPoolContainerConfigurationsContainerRegistriesResult
    {
        /// <summary>
        /// The password to log into the registry server.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The container registry URL. The default is "docker.io".
        /// </summary>
        public readonly string RegistryServer;
        /// <summary>
        /// The user name to log into the registry server.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetPoolContainerConfigurationsContainerRegistriesResult(
            string password,
            string registryServer,
            string userName)
        {
            Password = password;
            RegistryServer = registryServer;
            UserName = userName;
        }
    }

    [OutputType]
    public sealed class GetPoolContainerConfigurationsResult
    {
        /// <summary>
        /// Additional container registries from which container images can be pulled by the pool's VMs.
        /// </summary>
        public readonly ImmutableArray<GetPoolContainerConfigurationsContainerRegistriesResult> ContainerRegistries;
        /// <summary>
        /// The type of container configuration.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPoolContainerConfigurationsResult(
            ImmutableArray<GetPoolContainerConfigurationsContainerRegistriesResult> containerRegistries,
            string type)
        {
            ContainerRegistries = containerRegistries;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetPoolFixedScalesResult
    {
        /// <summary>
        /// The timeout for resize operations.
        /// </summary>
        public readonly string ResizeTimeout;
        /// <summary>
        /// The number of nodes in the Batch pool.
        /// </summary>
        public readonly int TargetDedicatedNodes;
        /// <summary>
        /// The number of low priority nodes in the Batch pool.
        /// </summary>
        public readonly int TargetLowPriorityNodes;

        [OutputConstructor]
        private GetPoolFixedScalesResult(
            string resizeTimeout,
            int targetDedicatedNodes,
            int targetLowPriorityNodes)
        {
            ResizeTimeout = resizeTimeout;
            TargetDedicatedNodes = targetDedicatedNodes;
            TargetLowPriorityNodes = targetLowPriorityNodes;
        }
    }

    [OutputType]
    public sealed class GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesResult
    {
        /// <summary>
        /// The action that should be taken for a specified IP address, subnet range or tag.
        /// </summary>
        public readonly string Access;
        /// <summary>
        /// The priority for this rule.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The source address prefix or tag to match for the rule.
        /// </summary>
        public readonly string SourceAddressPrefix;

        [OutputConstructor]
        private GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesResult(
            string access,
            int priority,
            string sourceAddressPrefix)
        {
            Access = access;
            Priority = priority;
            SourceAddressPrefix = sourceAddressPrefix;
        }
    }

    [OutputType]
    public sealed class GetPoolNetworkConfigurationEndpointConfigurationResult
    {
        /// <summary>
        /// The port number on the compute node.
        /// </summary>
        public readonly int BackendPort;
        /// <summary>
        /// The range of external ports that are used to provide inbound access to the backendPort on the individual compute nodes in the format of `1000-1100`.
        /// </summary>
        public readonly string FrontendPortRange;
        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of network security group rules that are applied to the endpoint.
        /// </summary>
        public readonly ImmutableArray<GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesResult> NetworkSecurityGroupRules;
        /// <summary>
        /// The protocol of the endpoint.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private GetPoolNetworkConfigurationEndpointConfigurationResult(
            int backendPort,
            string frontendPortRange,
            string name,
            ImmutableArray<GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesResult> networkSecurityGroupRules,
            string protocol)
        {
            BackendPort = backendPort;
            FrontendPortRange = frontendPortRange;
            Name = name;
            NetworkSecurityGroupRules = networkSecurityGroupRules;
            Protocol = protocol;
        }
    }

    [OutputType]
    public sealed class GetPoolNetworkConfigurationResult
    {
        /// <summary>
        /// The inbound NAT pools that are used to address specific ports on the individual compute node externally.
        /// </summary>
        public readonly GetPoolNetworkConfigurationEndpointConfigurationResult EndpointConfiguration;
        /// <summary>
        /// The ARM resource identifier of the virtual network subnet which the compute nodes of the pool are joined too.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private GetPoolNetworkConfigurationResult(
            GetPoolNetworkConfigurationEndpointConfigurationResult endpointConfiguration,
            string subnetId)
        {
            EndpointConfiguration = endpointConfiguration;
            SubnetId = subnetId;
        }
    }

    [OutputType]
    public sealed class GetPoolStartTaskResourceFilesResult
    {
        /// <summary>
        /// The storage container name in the auto storage account.
        /// </summary>
        public readonly string AutoStorageContainerName;
        /// <summary>
        /// The blob prefix used when downloading blobs from an Azure Storage container.
        /// </summary>
        public readonly string BlobPrefix;
        /// <summary>
        /// The file permission mode attribute represented as a string in octal format (e.g. `"0644"`).
        /// </summary>
        public readonly string FileMode;
        /// <summary>
        /// The location on the compute node to which to download the file, relative to the task's working directory. If the `http_url` property is specified, the `file_path` is required and describes the path which the file will be downloaded to, including the filename. Otherwise, if the `auto_storage_container_name` or `storage_container_url` property is specified.
        /// </summary>
        public readonly string FilePath;
        /// <summary>
        /// The URL of the file to download. If the URL is Azure Blob Storage, it must be readable using anonymous access.
        /// </summary>
        public readonly string HttpUrl;
        /// <summary>
        /// The URL of the blob container within Azure Blob Storage.
        /// </summary>
        public readonly string StorageContainerUrl;

        [OutputConstructor]
        private GetPoolStartTaskResourceFilesResult(
            string autoStorageContainerName,
            string blobPrefix,
            string fileMode,
            string filePath,
            string httpUrl,
            string storageContainerUrl)
        {
            AutoStorageContainerName = autoStorageContainerName;
            BlobPrefix = blobPrefix;
            FileMode = fileMode;
            FilePath = filePath;
            HttpUrl = httpUrl;
            StorageContainerUrl = storageContainerUrl;
        }
    }

    [OutputType]
    public sealed class GetPoolStartTaskResult
    {
        /// <summary>
        /// The command line executed by the start task.
        /// </summary>
        public readonly string CommandLine;
        /// <summary>
        /// A map of strings (key,value) that represents the environment variables to set in the start task.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Environment;
        /// <summary>
        /// The number of retry count.
        /// </summary>
        public readonly int? MaxTaskRetryCount;
        /// <summary>
        /// One or more `resource_file` blocks that describe the files to be downloaded to a compute node.
        /// </summary>
        public readonly ImmutableArray<GetPoolStartTaskResourceFilesResult> ResourceFiles;
        /// <summary>
        /// A `user_identity` block that describes the user identity under which the start task runs.
        /// </summary>
        public readonly ImmutableArray<GetPoolStartTaskUserIdentitiesResult> UserIdentities;
        /// <summary>
        /// A flag that indicates if the Batch pool should wait for the start task to be completed.
        /// </summary>
        public readonly bool? WaitForSuccess;

        [OutputConstructor]
        private GetPoolStartTaskResult(
            string commandLine,
            ImmutableDictionary<string, string>? environment,
            int? maxTaskRetryCount,
            ImmutableArray<GetPoolStartTaskResourceFilesResult> resourceFiles,
            ImmutableArray<GetPoolStartTaskUserIdentitiesResult> userIdentities,
            bool? waitForSuccess)
        {
            CommandLine = commandLine;
            Environment = environment;
            MaxTaskRetryCount = maxTaskRetryCount;
            ResourceFiles = resourceFiles;
            UserIdentities = userIdentities;
            WaitForSuccess = waitForSuccess;
        }
    }

    [OutputType]
    public sealed class GetPoolStartTaskUserIdentitiesAutoUsersResult
    {
        /// <summary>
        /// The elevation level of the user identity under which the start task runs.
        /// </summary>
        public readonly string ElevationLevel;
        /// <summary>
        /// The scope of the user identity under which the start task runs.
        /// </summary>
        public readonly string Scope;

        [OutputConstructor]
        private GetPoolStartTaskUserIdentitiesAutoUsersResult(
            string elevationLevel,
            string scope)
        {
            ElevationLevel = elevationLevel;
            Scope = scope;
        }
    }

    [OutputType]
    public sealed class GetPoolStartTaskUserIdentitiesResult
    {
        /// <summary>
        /// A `auto_user` block that describes the user identity under which the start task runs.
        /// </summary>
        public readonly ImmutableArray<GetPoolStartTaskUserIdentitiesAutoUsersResult> AutoUsers;
        /// <summary>
        /// The user name to log into the registry server.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetPoolStartTaskUserIdentitiesResult(
            ImmutableArray<GetPoolStartTaskUserIdentitiesAutoUsersResult> autoUsers,
            string userName)
        {
            AutoUsers = autoUsers;
            UserName = userName;
        }
    }

    [OutputType]
    public sealed class GetPoolStorageImageReferencesResult
    {
        /// <summary>
        /// The fully qualified ID of the certificate installed on the pool.
        /// </summary>
        public readonly string Id;
        public readonly string Offer;
        public readonly string Publisher;
        public readonly string Sku;
        public readonly string Version;

        [OutputConstructor]
        private GetPoolStorageImageReferencesResult(
            string id,
            string offer,
            string publisher,
            string sku,
            string version)
        {
            Id = id;
            Offer = offer;
            Publisher = publisher;
            Sku = sku;
            Version = version;
        }
    }
    }
}
