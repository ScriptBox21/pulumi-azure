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
    public sealed class MetricAlertDynamicCriteria
    {
        /// <summary>
        /// The statistic that runs over the metric values. Possible values are `Average`, `Count`, `Minimum`, `Maximum` and `Total`.
        /// </summary>
        public readonly string Aggregation;
        /// <summary>
        /// The extent of deviation required to trigger an alert. Possible values are `Low`, `Medium` and `High`.
        /// </summary>
        public readonly string AlertSensitivity;
        /// <summary>
        /// One or more `dimension` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricAlertDynamicCriteriaDimension> Dimensions;
        /// <summary>
        /// The number of violations to trigger an alert. Should be smaller or equal to `evaluation_total_count`.
        /// </summary>
        public readonly int? EvaluationFailureCount;
        /// <summary>
        /// The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (`window_size`) and the selected number of aggregated points.
        /// </summary>
        public readonly int? EvaluationTotalCount;
        /// <summary>
        /// The [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) date from which to start learning the metric historical data and calculate the dynamic thresholds.
        /// </summary>
        public readonly string? IgnoreDataBefore;
        /// <summary>
        /// One of the metric names to be monitored.
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// One of the metric namespaces to be monitored.
        /// </summary>
        public readonly string MetricNamespace;
        /// <summary>
        /// The criteria operator. Possible values are `LessThan`, `GreaterThan` and `GreaterOrLessThan`.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Skip the metric validation to allow creating an alert rule on a custom metric that isn't yet emitted? Defaults to `false`.
        /// </summary>
        public readonly bool? SkipMetricValidation;

        [OutputConstructor]
        private MetricAlertDynamicCriteria(
            string aggregation,

            string alertSensitivity,

            ImmutableArray<Outputs.MetricAlertDynamicCriteriaDimension> dimensions,

            int? evaluationFailureCount,

            int? evaluationTotalCount,

            string? ignoreDataBefore,

            string metricName,

            string metricNamespace,

            string @operator,

            bool? skipMetricValidation)
        {
            Aggregation = aggregation;
            AlertSensitivity = alertSensitivity;
            Dimensions = dimensions;
            EvaluationFailureCount = evaluationFailureCount;
            EvaluationTotalCount = evaluationTotalCount;
            IgnoreDataBefore = ignoreDataBefore;
            MetricName = metricName;
            MetricNamespace = metricNamespace;
            Operator = @operator;
            SkipMetricValidation = skipMetricValidation;
        }
    }
}
