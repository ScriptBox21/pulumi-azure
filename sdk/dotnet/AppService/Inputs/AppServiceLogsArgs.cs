// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class AppServiceLogsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `application_logs` block as defined below.
        /// </summary>
        [Input("applicationLogs")]
        public Input<Inputs.AppServiceLogsApplicationLogsArgs>? ApplicationLogs { get; set; }

        /// <summary>
        /// An `http_logs` block as defined below.
        /// </summary>
        [Input("httpLogs")]
        public Input<Inputs.AppServiceLogsHttpLogsArgs>? HttpLogs { get; set; }

        public AppServiceLogsArgs()
        {
        }
    }
}
