// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class ScheduledQueryRulesLogCriteriaArgs : Pulumi.ResourceArgs
    {
        [Input("dimensions", required: true)]
        private InputList<Inputs.ScheduledQueryRulesLogCriteriaDimensionArgs>? _dimensions;

        /// <summary>
        /// A `dimension` block as defined below.
        /// </summary>
        public InputList<Inputs.ScheduledQueryRulesLogCriteriaDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.ScheduledQueryRulesLogCriteriaDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Name of the metric.  Supported metrics are listed in the Azure Monitor [Microsoft.OperationalInsights/workspaces](https://docs.microsoft.com/en-us/azure/azure-monitor/platform/metrics-supported#microsoftoperationalinsightsworkspaces) metrics namespace.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        public ScheduledQueryRulesLogCriteriaArgs()
        {
        }
    }
}
