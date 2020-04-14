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
    public sealed class GetScheduledQueryRulesAlertTriggerResult
    {
        public readonly ImmutableArray<Outputs.GetScheduledQueryRulesAlertTriggerMetricTriggerResult> MetricTriggers;
        /// <summary>
        /// Evaluation operation for rule.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Result or count threshold based on which rule should be triggered.
        /// </summary>
        public readonly double Threshold;

        [OutputConstructor]
        private GetScheduledQueryRulesAlertTriggerResult(
            ImmutableArray<Outputs.GetScheduledQueryRulesAlertTriggerMetricTriggerResult> metricTriggers,

            string @operator,

            double threshold)
        {
            MetricTriggers = metricTriggers;
            Operator = @operator;
            Threshold = threshold;
        }
    }
}
