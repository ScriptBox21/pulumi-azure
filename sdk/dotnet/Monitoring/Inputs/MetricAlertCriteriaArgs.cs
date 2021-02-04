// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class MetricAlertCriteriaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The statistic that runs over the metric values. Possible values are `Average`, `Count`, `Minimum`, `Maximum` and `Total`.
        /// </summary>
        [Input("aggregation", required: true)]
        public Input<string> Aggregation { get; set; } = null!;

        [Input("dimensions")]
        private InputList<Inputs.MetricAlertCriteriaDimensionArgs>? _dimensions;

        /// <summary>
        /// One or more `dimension` blocks as defined below.
        /// </summary>
        public InputList<Inputs.MetricAlertCriteriaDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.MetricAlertCriteriaDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// One of the metric names to be monitored.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        /// <summary>
        /// One of the metric namespaces to be monitored.
        /// </summary>
        [Input("metricNamespace", required: true)]
        public Input<string> MetricNamespace { get; set; } = null!;

        /// <summary>
        /// The criteria operator. Possible values are `Equals`, `NotEquals`, `GreaterThan`, `GreaterThanOrEqual`, `LessThan` and `LessThanOrEqual`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// Skip the metric validation to allow creating an alert rule on a custom metric that isn't yet emitted? Defaults to `false`.
        /// </summary>
        [Input("skipMetricValidation")]
        public Input<bool>? SkipMetricValidation { get; set; }

        /// <summary>
        /// The criteria threshold value that activates the alert.
        /// </summary>
        [Input("threshold", required: true)]
        public Input<double> Threshold { get; set; } = null!;

        public MetricAlertCriteriaArgs()
        {
        }
    }
}
