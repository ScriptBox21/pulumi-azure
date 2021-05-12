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
    public sealed class InteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceSchedule
    {
        /// <summary>
        /// The days of the week to perform autoscale.
        /// </summary>
        public readonly ImmutableArray<string> Days;
        /// <summary>
        /// The number of worker nodes to autoscale at the specified time.
        /// </summary>
        public readonly int TargetInstanceCount;
        /// <summary>
        /// The time of day to perform the autoscale in 24hour format.
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private InteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceSchedule(
            ImmutableArray<string> days,

            int targetInstanceCount,

            string time)
        {
            Days = days;
            TargetInstanceCount = targetInstanceCount;
            Time = time;
        }
    }
}
