// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access the properties of a Log Profile.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/monitor_log_profile.html.markdown.
        /// </summary>
        [Obsolete("Use GetLogProfile.InvokeAsync() instead")]
        public static Task<GetLogProfileResult> GetLogProfile(GetLogProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLogProfileResult>("azure:monitoring/getLogProfile:getLogProfile", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetLogProfile
    {
        /// <summary>
        /// Use this data source to access the properties of a Log Profile.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/monitor_log_profile.html.markdown.
        /// </summary>
        public static Task<GetLogProfileResult> InvokeAsync(GetLogProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLogProfileResult>("azure:monitoring/getLogProfile:getLogProfile", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetLogProfileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Log Profile.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetLogProfileArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetLogProfileResult
    {
        /// <summary>
        /// List of categories of the logs.
        /// </summary>
        public readonly ImmutableArray<string> Categories;
        /// <summary>
        /// List of regions for which Activity Log events are stored or streamed.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetLogProfileRetentionPoliciesResult> RetentionPolicies;
        /// <summary>
        /// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.
        /// </summary>
        public readonly string ServicebusRuleId;
        /// <summary>
        /// The resource id of the storage account in which the Activity Log is stored.
        /// </summary>
        public readonly string StorageAccountId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetLogProfileResult(
            ImmutableArray<string> categories,
            ImmutableArray<string> locations,
            string name,
            ImmutableArray<Outputs.GetLogProfileRetentionPoliciesResult> retentionPolicies,
            string servicebusRuleId,
            string storageAccountId,
            string id)
        {
            Categories = categories;
            Locations = locations;
            Name = name;
            RetentionPolicies = retentionPolicies;
            ServicebusRuleId = servicebusRuleId;
            StorageAccountId = storageAccountId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetLogProfileRetentionPoliciesResult
    {
        /// <summary>
        /// The number of days for the retention policy.
        /// </summary>
        public readonly int Days;
        /// <summary>
        /// A boolean value indicating whether the retention policy is enabled.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private GetLogProfileRetentionPoliciesResult(
            int days,
            bool enabled)
        {
            Days = days;
            Enabled = enabled;
        }
    }
    }
}
