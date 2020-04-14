// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Outputs
{

    [OutputType]
    public sealed class ElasticPoolSku
    {
        /// <summary>
        /// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// The `family` of hardware `Gen4` or `Gen5`.
        /// </summary>
        public readonly string? Family;
        /// <summary>
        /// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either `vCore` based `tier` + `family` pattern (e.g. GP_Gen4, BC_Gen5) or the `DTU` based `BasicPool`, `StandardPool`, or `PremiumPool` pattern.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The tier of the particular SKU. Possible values are `GeneralPurpose`, `BusinessCritical`, `Basic`, `Standard`, or `Premium`. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private ElasticPoolSku(
            int capacity,

            string? family,

            string name,

            string tier)
        {
            Capacity = capacity;
            Family = family;
            Name = name;
            Tier = tier;
        }
    }
}
