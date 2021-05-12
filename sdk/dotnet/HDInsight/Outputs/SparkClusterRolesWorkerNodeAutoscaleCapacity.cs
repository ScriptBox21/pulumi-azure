// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class SparkClusterRolesWorkerNodeAutoscaleCapacity
    {
        /// <summary>
        /// The maximum number of worker nodes to autoscale to based on the cluster's activity.
        /// </summary>
        public readonly int MaxInstanceCount;
        /// <summary>
        /// The minimum number of worker nodes to autoscale to based on the cluster's activity.
        /// </summary>
        public readonly int MinInstanceCount;

        [OutputConstructor]
        private SparkClusterRolesWorkerNodeAutoscaleCapacity(
            int maxInstanceCount,

            int minInstanceCount)
        {
            MaxInstanceCount = maxInstanceCount;
            MinInstanceCount = minInstanceCount;
        }
    }
}
