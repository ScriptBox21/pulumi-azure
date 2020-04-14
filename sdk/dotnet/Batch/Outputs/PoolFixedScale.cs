// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch.Outputs
{

    [OutputType]
    public sealed class PoolFixedScale
    {
        /// <summary>
        /// The timeout for resize operations. Defaults to `PT15M`.
        /// </summary>
        public readonly string? ResizeTimeout;
        /// <summary>
        /// The number of nodes in the Batch pool. Defaults to `1`.
        /// </summary>
        public readonly int? TargetDedicatedNodes;
        /// <summary>
        /// The number of low priority nodes in the Batch pool. Defaults to `0`.
        /// </summary>
        public readonly int? TargetLowPriorityNodes;

        [OutputConstructor]
        private PoolFixedScale(
            string? resizeTimeout,

            int? targetDedicatedNodes,

            int? targetLowPriorityNodes)
        {
            ResizeTimeout = resizeTimeout;
            TargetDedicatedNodes = targetDedicatedNodes;
            TargetLowPriorityNodes = targetLowPriorityNodes;
        }
    }
}
