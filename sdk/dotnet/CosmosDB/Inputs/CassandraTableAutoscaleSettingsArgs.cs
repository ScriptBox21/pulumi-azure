// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Inputs
{

    public sealed class CassandraTableAutoscaleSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum throughput of the Cassandra Table (RU/s). Must be between `4,000` and `1,000,000`. Must be set in increments of `1,000`. Conflicts with `throughput`.
        /// </summary>
        [Input("maxThroughput")]
        public Input<int>? MaxThroughput { get; set; }

        public CassandraTableAutoscaleSettingsArgs()
        {
        }
    }
}
