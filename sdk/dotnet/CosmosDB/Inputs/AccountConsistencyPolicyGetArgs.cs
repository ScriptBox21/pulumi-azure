// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Inputs
{

    public sealed class AccountConsistencyPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Consistency Level to use for this CosmosDB Account - can be either `BoundedStaleness`, `Eventual`, `Session`, `Strong` or `ConsistentPrefix`.
        /// </summary>
        [Input("consistencyLevel", required: true)]
        public Input<string> ConsistencyLevel { get; set; } = null!;

        /// <summary>
        /// When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is `5` - `86400` (1 day). Defaults to `5`. Required when `consistency_level` is set to `BoundedStaleness`.
        /// </summary>
        [Input("maxIntervalInSeconds")]
        public Input<int>? MaxIntervalInSeconds { get; set; }

        /// <summary>
        /// When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is `10` – `2147483647`. Defaults to `100`. Required when `consistency_level` is set to `BoundedStaleness`.
        /// </summary>
        [Input("maxStalenessPrefix")]
        public Input<int>? MaxStalenessPrefix { get; set; }

        public AccountConsistencyPolicyGetArgs()
        {
        }
    }
}
