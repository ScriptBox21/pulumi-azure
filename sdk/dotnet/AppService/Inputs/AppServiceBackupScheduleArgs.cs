// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class AppServiceBackupScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets how often the backup should be executed.
        /// </summary>
        [Input("frequencyInterval", required: true)]
        public Input<int> FrequencyInterval { get; set; } = null!;

        /// <summary>
        /// Sets the unit of time for how often the backup should be executed. Possible values are `Day` or `Hour`.
        /// </summary>
        [Input("frequencyUnit", required: true)]
        public Input<string> FrequencyUnit { get; set; } = null!;

        /// <summary>
        /// Should at least one backup always be kept in the Storage Account by the Retention Policy, regardless of how old it is?
        /// </summary>
        [Input("keepAtLeastOneBackup")]
        public Input<bool>? KeepAtLeastOneBackup { get; set; }

        /// <summary>
        /// Specifies the number of days after which Backups should be deleted.
        /// </summary>
        [Input("retentionPeriodInDays")]
        public Input<int>? RetentionPeriodInDays { get; set; }

        /// <summary>
        /// Sets when the schedule should start working.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public AppServiceBackupScheduleArgs()
        {
        }
    }
}
