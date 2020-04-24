// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sql.Outputs
{

    [OutputType]
    public sealed class SqlServerIdentity
    {
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// The Tenant ID for the Service Principal associated with the Identity of this SQL Server.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// Specifies the identity type of the SQL Server. At this time the only allowed value is `SystemAssigned`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SqlServerIdentity(
            string? principalId,

            string? tenantId,

            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}