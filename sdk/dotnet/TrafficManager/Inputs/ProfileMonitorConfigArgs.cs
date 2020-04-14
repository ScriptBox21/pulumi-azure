// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.TrafficManager.Inputs
{

    public sealed class ProfileMonitorConfigArgs : Pulumi.ResourceArgs
    {
        [Input("customHeaders")]
        private InputList<Inputs.ProfileMonitorConfigCustomHeaderArgs>? _customHeaders;

        /// <summary>
        /// One or more `custom_header` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ProfileMonitorConfigCustomHeaderArgs> CustomHeaders
        {
            get => _customHeaders ?? (_customHeaders = new InputList<Inputs.ProfileMonitorConfigCustomHeaderArgs>());
            set => _customHeaders = value;
        }

        [Input("expectedStatusCodeRanges")]
        private InputList<string>? _expectedStatusCodeRanges;

        /// <summary>
        /// A list of status code ranges in the format of `100-101`.
        /// </summary>
        public InputList<string> ExpectedStatusCodeRanges
        {
            get => _expectedStatusCodeRanges ?? (_expectedStatusCodeRanges = new InputList<string>());
            set => _expectedStatusCodeRanges = value;
        }

        /// <summary>
        /// The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: `30` (normal probing) and `10` (fast probing). The default value is `30`.
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        /// <summary>
        /// The path used by the monitoring checks. Required when `protocol` is set to `HTTP` or `HTTPS` - cannot be set when `protocol` is set to `TCP`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The port number used by the monitoring checks.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol used by the monitoring checks, supported values are `HTTP`, `HTTPS` and `TCP`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If `interval_in_seconds` is set to `30`, then `timeout_in_seconds` can be between `5` and `10`. The default value is `10`. If `interval_in_seconds` is set to `10`, then valid values are between `5` and `9` and `timeout_in_seconds` is required.
        /// </summary>
        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        /// <summary>
        /// The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between `0` and `9`. The default value is `3`
        /// </summary>
        [Input("toleratedNumberOfFailures")]
        public Input<int>? ToleratedNumberOfFailures { get; set; }

        public ProfileMonitorConfigArgs()
        {
        }
    }
}
