// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric.Inputs
{

    public sealed class ClusterAzureActiveDirectoryGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Active Directory Client ID which should be used for the Client Application.
        /// </summary>
        [Input("clientApplicationId", required: true)]
        public Input<string> ClientApplicationId { get; set; } = null!;

        /// <summary>
        /// The Azure Active Directory Cluster Application ID.
        /// </summary>
        [Input("clusterApplicationId", required: true)]
        public Input<string> ClusterApplicationId { get; set; } = null!;

        /// <summary>
        /// The Azure Active Directory Tenant ID.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public ClusterAzureActiveDirectoryGetArgs()
        {
        }
    }
}
