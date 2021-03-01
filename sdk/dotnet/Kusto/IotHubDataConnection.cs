// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Kusto
{
    /// <summary>
    /// Manages a Kusto (also known as Azure Data Explorer) IotHub Data Connection
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
    ///             Location = "West Europe",
    ///         });
    ///         var exampleCluster = new Azure.Kusto.Cluster("exampleCluster", new Azure.Kusto.ClusterArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = new Azure.Kusto.Inputs.ClusterSkuArgs
    ///             {
    ///                 Name = "Standard_D13_v2",
    ///                 Capacity = 2,
    ///             },
    ///         });
    ///         var exampleDatabase = new Azure.Kusto.Database("exampleDatabase", new Azure.Kusto.DatabaseArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ClusterName = exampleCluster.Name,
    ///             HotCachePeriod = "P7D",
    ///             SoftDeletePeriod = "P31D",
    ///         });
    ///         var exampleIoTHub = new Azure.Iot.IoTHub("exampleIoTHub", new Azure.Iot.IoTHubArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = new Azure.Iot.Inputs.IoTHubSkuArgs
    ///             {
    ///                 Name = "B1",
    ///                 Capacity = 1,
    ///             },
    ///         });
    ///         var exampleSharedAccessPolicy = new Azure.Iot.SharedAccessPolicy("exampleSharedAccessPolicy", new Azure.Iot.SharedAccessPolicyArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubName = exampleIoTHub.Name,
    ///             RegistryRead = true,
    ///         });
    ///         var exampleConsumerGroup = new Azure.Iot.ConsumerGroup("exampleConsumerGroup", new Azure.Iot.ConsumerGroupArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IothubName = exampleIoTHub.Name,
    ///             EventhubEndpointName = "events",
    ///         });
    ///         var exampleIotHubDataConnection = new Azure.Kusto.IotHubDataConnection("exampleIotHubDataConnection", new Azure.Kusto.IotHubDataConnectionArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             ClusterName = exampleCluster.Name,
    ///             DatabaseName = exampleDatabase.Name,
    ///             IothubId = exampleIoTHub.Id,
    ///             ConsumerGroup = exampleConsumerGroup.Name,
    ///             SharedAccessPolicyName = exampleSharedAccessPolicy.Name,
    ///             EventSystemProperties = 
    ///             {
    ///                 "message-id",
    ///                 "sequence-number",
    ///                 "to",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Kusto IotHub Data Connections can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:kusto/iotHubDataConnection:IotHubDataConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/dataConnection1
    /// ```
    /// </summary>
    [AzureResourceType("azure:kusto/iotHubDataConnection:IotHubDataConnection")]
    public partial class IotHubDataConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Output("consumerGroup")]
        public Output<string> ConsumerGroup { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
        /// </summary>
        [Output("eventSystemProperties")]
        public Output<ImmutableArray<string>> EventSystemProperties { get; private set; } = null!;

        /// <summary>
        /// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Output("iothubId")]
        public Output<string> IothubId { get; private set; } = null!;

        /// <summary>
        /// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sharedAccessPolicyName")]
        public Output<string> SharedAccessPolicyName { get; private set; } = null!;


        /// <summary>
        /// Create a IotHubDataConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotHubDataConnection(string name, IotHubDataConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure:kusto/iotHubDataConnection:IotHubDataConnection", name, args ?? new IotHubDataConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotHubDataConnection(string name, Input<string> id, IotHubDataConnectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:kusto/iotHubDataConnection:IotHubDataConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IotHubDataConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotHubDataConnection Get(string name, Input<string> id, IotHubDataConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new IotHubDataConnection(name, id, state, options);
        }
    }

    public sealed class IotHubDataConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("consumerGroup", required: true)]
        public Input<string> ConsumerGroup { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("eventSystemProperties")]
        private InputList<string>? _eventSystemProperties;

        /// <summary>
        /// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> EventSystemProperties
        {
            get => _eventSystemProperties ?? (_eventSystemProperties = new InputList<string>());
            set => _eventSystemProperties = value;
        }

        /// <summary>
        /// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubId", required: true)]
        public Input<string> IothubId { get; set; } = null!;

        /// <summary>
        /// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sharedAccessPolicyName", required: true)]
        public Input<string> SharedAccessPolicyName { get; set; } = null!;

        public IotHubDataConnectionArgs()
        {
        }
    }

    public sealed class IotHubDataConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("consumerGroup")]
        public Input<string>? ConsumerGroup { get; set; }

        /// <summary>
        /// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("eventSystemProperties")]
        private InputList<string>? _eventSystemProperties;

        /// <summary>
        /// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> EventSystemProperties
        {
            get => _eventSystemProperties ?? (_eventSystemProperties = new InputList<string>());
            set => _eventSystemProperties = value;
        }

        /// <summary>
        /// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubId")]
        public Input<string>? IothubId { get; set; }

        /// <summary>
        /// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sharedAccessPolicyName")]
        public Input<string>? SharedAccessPolicyName { get; set; }

        public IotHubDataConnectionState()
        {
        }
    }
}
