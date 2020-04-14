// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Outputs
{

    [OutputType]
    public sealed class DiagnosticSettingMetric
    {
        /// <summary>
        /// The name of a Diagnostic Metric Category for this Resource.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Is this Diagnostic Metric enabled? Defaults to `true`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// A `retention_policy` block as defined below.
        /// </summary>
        public readonly Outputs.DiagnosticSettingMetricRetentionPolicy RetentionPolicy;

        [OutputConstructor]
        private DiagnosticSettingMetric(
            string category,

            bool? enabled,

            Outputs.DiagnosticSettingMetricRetentionPolicy retentionPolicy)
        {
            Category = category;
            Enabled = enabled;
            RetentionPolicy = retentionPolicy;
        }
    }
}
