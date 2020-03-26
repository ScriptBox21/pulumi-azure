// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CostManagement
{
    /// <summary>
    /// Manages an Azure Cost Management Export for a Resource Group.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cost_management_export_resource_group.html.markdown.
    /// </summary>
    public partial class ResourceGroupExport : Pulumi.CustomResource
    {
        /// <summary>
        /// Is the cost management export active? Default is `true`.
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// A `delivery_info` block as defined below.
        /// </summary>
        [Output("deliveryInfo")]
        public Output<Outputs.ResourceGroupExportDeliveryInfo> DeliveryInfo { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `query` block as defined below.
        /// </summary>
        [Output("query")]
        public Output<Outputs.ResourceGroupExportQuery> Query { get; private set; } = null!;

        /// <summary>
        /// The date the export will stop capturing information. 
        /// </summary>
        [Output("recurrencePeriodEnd")]
        public Output<string> RecurrencePeriodEnd { get; private set; } = null!;

        /// <summary>
        /// The date the export will start capturing information.
        /// </summary>
        [Output("recurrencePeriodStart")]
        public Output<string> RecurrencePeriodStart { get; private set; } = null!;

        /// <summary>
        /// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
        /// </summary>
        [Output("recurrenceType")]
        public Output<string> RecurrenceType { get; private set; } = null!;

        /// <summary>
        /// The id of the resource group in which to export information.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceGroupExport resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceGroupExport(string name, ResourceGroupExportArgs args, CustomResourceOptions? options = null)
            : base("azure:costmanagement/resourceGroupExport:ResourceGroupExport", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ResourceGroupExport(string name, Input<string> id, ResourceGroupExportState? state = null, CustomResourceOptions? options = null)
            : base("azure:costmanagement/resourceGroupExport:ResourceGroupExport", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceGroupExport resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceGroupExport Get(string name, Input<string> id, ResourceGroupExportState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceGroupExport(name, id, state, options);
        }
    }

    public sealed class ResourceGroupExportArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the cost management export active? Default is `true`.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// A `delivery_info` block as defined below.
        /// </summary>
        [Input("deliveryInfo", required: true)]
        public Input<Inputs.ResourceGroupExportDeliveryInfoArgs> DeliveryInfo { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `query` block as defined below.
        /// </summary>
        [Input("query", required: true)]
        public Input<Inputs.ResourceGroupExportQueryArgs> Query { get; set; } = null!;

        /// <summary>
        /// The date the export will stop capturing information. 
        /// </summary>
        [Input("recurrencePeriodEnd", required: true)]
        public Input<string> RecurrencePeriodEnd { get; set; } = null!;

        /// <summary>
        /// The date the export will start capturing information.
        /// </summary>
        [Input("recurrencePeriodStart", required: true)]
        public Input<string> RecurrencePeriodStart { get; set; } = null!;

        /// <summary>
        /// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
        /// </summary>
        [Input("recurrenceType", required: true)]
        public Input<string> RecurrenceType { get; set; } = null!;

        /// <summary>
        /// The id of the resource group in which to export information.
        /// </summary>
        [Input("resourceGroupId", required: true)]
        public Input<string> ResourceGroupId { get; set; } = null!;

        public ResourceGroupExportArgs()
        {
        }
    }

    public sealed class ResourceGroupExportState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the cost management export active? Default is `true`.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// A `delivery_info` block as defined below.
        /// </summary>
        [Input("deliveryInfo")]
        public Input<Inputs.ResourceGroupExportDeliveryInfoGetArgs>? DeliveryInfo { get; set; }

        /// <summary>
        /// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `query` block as defined below.
        /// </summary>
        [Input("query")]
        public Input<Inputs.ResourceGroupExportQueryGetArgs>? Query { get; set; }

        /// <summary>
        /// The date the export will stop capturing information. 
        /// </summary>
        [Input("recurrencePeriodEnd")]
        public Input<string>? RecurrencePeriodEnd { get; set; }

        /// <summary>
        /// The date the export will start capturing information.
        /// </summary>
        [Input("recurrencePeriodStart")]
        public Input<string>? RecurrencePeriodStart { get; set; }

        /// <summary>
        /// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
        /// </summary>
        [Input("recurrenceType")]
        public Input<string>? RecurrenceType { get; set; }

        /// <summary>
        /// The id of the resource group in which to export information.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        public ResourceGroupExportState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ResourceGroupExportDeliveryInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the container where exports will be uploaded.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// The path of the directory where exports will be uploaded.
        /// </summary>
        [Input("rootFolderPath", required: true)]
        public Input<string> RootFolderPath { get; set; } = null!;

        /// <summary>
        /// The storage account id where exports will be delivered.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public ResourceGroupExportDeliveryInfoArgs()
        {
        }
    }

    public sealed class ResourceGroupExportDeliveryInfoGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the container where exports will be uploaded.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// The path of the directory where exports will be uploaded.
        /// </summary>
        [Input("rootFolderPath", required: true)]
        public Input<string> RootFolderPath { get; set; } = null!;

        /// <summary>
        /// The storage account id where exports will be delivered.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public ResourceGroupExportDeliveryInfoGetArgs()
        {
        }
    }

    public sealed class ResourceGroupExportQueryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `YearToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastYear`, `Custom`.
        /// </summary>
        [Input("timeFrame", required: true)]
        public Input<string> TimeFrame { get; set; } = null!;

        /// <summary>
        /// The type of the query.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ResourceGroupExportQueryArgs()
        {
        }
    }

    public sealed class ResourceGroupExportQueryGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `YearToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastYear`, `Custom`.
        /// </summary>
        [Input("timeFrame", required: true)]
        public Input<string> TimeFrame { get; set; } = null!;

        /// <summary>
        /// The type of the query.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ResourceGroupExportQueryGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ResourceGroupExportDeliveryInfo
    {
        /// <summary>
        /// The name of the container where exports will be uploaded.
        /// </summary>
        public readonly string ContainerName;
        /// <summary>
        /// The path of the directory where exports will be uploaded.
        /// </summary>
        public readonly string RootFolderPath;
        /// <summary>
        /// The storage account id where exports will be delivered.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private ResourceGroupExportDeliveryInfo(
            string containerName,
            string rootFolderPath,
            string storageAccountId)
        {
            ContainerName = containerName;
            RootFolderPath = rootFolderPath;
            StorageAccountId = storageAccountId;
        }
    }

    [OutputType]
    public sealed class ResourceGroupExportQuery
    {
        /// <summary>
        /// The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `YearToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastYear`, `Custom`.
        /// </summary>
        public readonly string TimeFrame;
        /// <summary>
        /// The type of the query.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ResourceGroupExportQuery(
            string timeFrame,
            string type)
        {
            TimeFrame = timeFrame;
            Type = type;
        }
    }
    }
}
