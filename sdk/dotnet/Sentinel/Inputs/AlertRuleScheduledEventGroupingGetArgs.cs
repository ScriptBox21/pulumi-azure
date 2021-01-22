// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel.Inputs
{

    public sealed class AlertRuleScheduledEventGroupingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The aggregation type of grouping the events.
        /// </summary>
        [Input("aggregationMethod", required: true)]
        public Input<string> AggregationMethod { get; set; } = null!;

        public AlertRuleScheduledEventGroupingGetArgs()
        {
        }
    }
}
