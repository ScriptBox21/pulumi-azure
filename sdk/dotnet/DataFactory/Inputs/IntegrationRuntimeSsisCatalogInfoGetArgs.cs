// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Inputs
{

    public sealed class IntegrationRuntimeSsisCatalogInfoGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrator login name for the SQL Server.
        /// </summary>
        [Input("administratorLogin", required: true)]
        public Input<string> AdministratorLogin { get; set; } = null!;

        /// <summary>
        /// Administrator login password for the SQL Server.
        /// </summary>
        [Input("administratorPassword", required: true)]
        public Input<string> AdministratorPassword { get; set; } = null!;

        /// <summary>
        /// Pricing tier for the database that will be created for the SSIS catalog. Valid values are: `Basic`, `Standard`, `Premium` and `PremiumRS`.
        /// </summary>
        [Input("pricingTier")]
        public Input<string>? PricingTier { get; set; }

        /// <summary>
        /// The endpoint of an Azure SQL Server that will be used to host the SSIS catalog.
        /// </summary>
        [Input("serverEndpoint", required: true)]
        public Input<string> ServerEndpoint { get; set; } = null!;

        public IntegrationRuntimeSsisCatalogInfoGetArgs()
        {
        }
    }
}
