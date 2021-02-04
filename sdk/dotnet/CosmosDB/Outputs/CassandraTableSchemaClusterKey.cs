// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Outputs
{

    [OutputType]
    public sealed class CassandraTableSchemaClusterKey
    {
        /// <summary>
        /// Name of the cluster key to be created.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Order of the key. Currently supported values are `Asc` and `Desc`.
        /// </summary>
        public readonly string OrderBy;

        [OutputConstructor]
        private CassandraTableSchemaClusterKey(
            string name,

            string orderBy)
        {
            Name = name;
            OrderBy = orderBy;
        }
    }
}
