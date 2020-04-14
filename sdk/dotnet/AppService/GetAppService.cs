// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    public static class GetAppService
    {
        /// <summary>
        /// Use this data source to access information about an existing App Service.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppServiceResult> InvokeAsync(GetAppServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServiceResult>("azure:appservice/getAppService:getAppService", args ?? new GetAppServiceArgs(), options.WithVersion());
    }


    public sealed class GetAppServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the App Service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the App Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAppServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServiceResult
    {
        /// <summary>
        /// The ID of the App Service Plan within which the App Service exists.
        /// </summary>
        public readonly string AppServicePlanId;
        /// <summary>
        /// A key-value pair of App Settings for the App Service.
        /// </summary>
        public readonly ImmutableDictionary<string, string> AppSettings;
        /// <summary>
        /// Does the App Service send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        public readonly bool ClientAffinityEnabled;
        /// <summary>
        /// Does the App Service require client certificates for incoming requests?
        /// </summary>
        public readonly bool ClientCertEnabled;
        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAppServiceConnectionStringResult> ConnectionStrings;
        /// <summary>
        /// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
        /// </summary>
        public readonly string DefaultSiteHostname;
        /// <summary>
        /// Is the App Service Enabled?
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Can the App Service only be accessed via HTTPS?
        /// </summary>
        public readonly bool HttpsOnly;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure location where the App Service exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the Connection String.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
        /// </summary>
        public readonly string OutboundIpAddresses;
        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        /// </summary>
        public readonly string PossibleOutboundIpAddresses;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A `site_config` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAppServiceSiteConfigResult> SiteConfigs;
        public readonly ImmutableArray<Outputs.GetAppServiceSiteCredentialResult> SiteCredentials;
        public readonly ImmutableArray<Outputs.GetAppServiceSourceControlResult> SourceControls;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetAppServiceResult(
            string appServicePlanId,

            ImmutableDictionary<string, string> appSettings,

            bool clientAffinityEnabled,

            bool clientCertEnabled,

            ImmutableArray<Outputs.GetAppServiceConnectionStringResult> connectionStrings,

            string defaultSiteHostname,

            bool enabled,

            bool httpsOnly,

            string id,

            string location,

            string name,

            string outboundIpAddresses,

            string possibleOutboundIpAddresses,

            string resourceGroupName,

            ImmutableArray<Outputs.GetAppServiceSiteConfigResult> siteConfigs,

            ImmutableArray<Outputs.GetAppServiceSiteCredentialResult> siteCredentials,

            ImmutableArray<Outputs.GetAppServiceSourceControlResult> sourceControls,

            ImmutableDictionary<string, string> tags)
        {
            AppServicePlanId = appServicePlanId;
            AppSettings = appSettings;
            ClientAffinityEnabled = clientAffinityEnabled;
            ClientCertEnabled = clientCertEnabled;
            ConnectionStrings = connectionStrings;
            DefaultSiteHostname = defaultSiteHostname;
            Enabled = enabled;
            HttpsOnly = httpsOnly;
            Id = id;
            Location = location;
            Name = name;
            OutboundIpAddresses = outboundIpAddresses;
            PossibleOutboundIpAddresses = possibleOutboundIpAddresses;
            ResourceGroupName = resourceGroupName;
            SiteConfigs = siteConfigs;
            SiteCredentials = siteCredentials;
            SourceControls = sourceControls;
            Tags = tags;
        }
    }
}
