// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Outputs
{

    [OutputType]
    public sealed class IntegrationRuntimeSsisCatalogInfo
    {
        /// <summary>
        /// Administrator login name for the SQL Server.
        /// </summary>
        public readonly string AdministratorLogin;
        /// <summary>
        /// Administrator login password for the SQL Server.
        /// </summary>
        public readonly string AdministratorPassword;
        /// <summary>
        /// Pricing tier for the database that will be created for the SSIS catalog. Valid values are: `Basic`, `Standard`, `Premium` and `PremiumRS`.
        /// </summary>
        public readonly string? PricingTier;
        /// <summary>
        /// The endpoint of an Azure SQL Server that will be used to host the SSIS catalog.
        /// </summary>
        public readonly string ServerEndpoint;

        [OutputConstructor]
        private IntegrationRuntimeSsisCatalogInfo(
            string administratorLogin,

            string administratorPassword,

            string? pricingTier,

            string serverEndpoint)
        {
            AdministratorLogin = administratorLogin;
            AdministratorPassword = administratorPassword;
            PricingTier = pricingTier;
            ServerEndpoint = serverEndpoint;
        }
    }
}
