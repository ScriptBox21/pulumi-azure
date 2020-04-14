// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class GroupDiagnosticsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `log_analytics` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logAnalytics", required: true)]
        public Input<Inputs.GroupDiagnosticsLogAnalyticsGetArgs> LogAnalytics { get; set; } = null!;

        public GroupDiagnosticsGetArgs()
        {
        }
    }
}
