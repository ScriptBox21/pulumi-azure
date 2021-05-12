// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class SparkClusterRolesWorkerNodeAutoscaleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `capacity` block as defined below.
        /// </summary>
        [Input("capacity")]
        public Input<Inputs.SparkClusterRolesWorkerNodeAutoscaleCapacityGetArgs>? Capacity { get; set; }

        /// <summary>
        /// A `recurrence` block as defined below.
        /// </summary>
        [Input("recurrence")]
        public Input<Inputs.SparkClusterRolesWorkerNodeAutoscaleRecurrenceGetArgs>? Recurrence { get; set; }

        public SparkClusterRolesWorkerNodeAutoscaleGetArgs()
        {
        }
    }
}
