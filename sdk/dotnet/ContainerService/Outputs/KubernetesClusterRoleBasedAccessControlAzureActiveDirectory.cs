// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class KubernetesClusterRoleBasedAccessControlAzureActiveDirectory
    {
        /// <summary>
        /// The Client ID of an Azure Active Directory Application.
        /// </summary>
        public readonly string ClientAppId;
        /// <summary>
        /// The Server ID of an Azure Active Directory Application.
        /// </summary>
        public readonly string ServerAppId;
        /// <summary>
        /// The Server Secret of an Azure Active Directory Application.
        /// </summary>
        public readonly string ServerAppSecret;
        /// <summary>
        /// The Tenant ID used for Azure Active Directory Application. If this isn't specified the Tenant ID of the current Subscription is used.
        /// </summary>
        public readonly string? TenantId;

        [OutputConstructor]
        private KubernetesClusterRoleBasedAccessControlAzureActiveDirectory(
            string clientAppId,

            string serverAppId,

            string serverAppSecret,

            string? tenantId)
        {
            ClientAppId = clientAppId;
            ServerAppId = serverAppId;
            ServerAppSecret = serverAppSecret;
            TenantId = tenantId;
        }
    }
}
