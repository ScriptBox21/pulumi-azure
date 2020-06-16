// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare.Outputs
{

    [OutputType]
    public sealed class GetShareSnapshotScheduleResult
    {
        /// <summary>
        /// The name of this Data Share.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The interval of the synchronization with the source data.
        /// </summary>
        public readonly string Recurrence;
        /// <summary>
        /// The synchronization with the source data's start time.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private GetShareSnapshotScheduleResult(
            string name,

            string recurrence,

            string startTime)
        {
            Name = name;
            Recurrence = recurrence;
            StartTime = startTime;
        }
    }
}