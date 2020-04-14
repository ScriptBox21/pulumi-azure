// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.TrafficManager.Outputs
{

    [OutputType]
    public sealed class ProfileMonitorConfig
    {
        /// <summary>
        /// One or more `custom_header` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileMonitorConfigCustomHeader> CustomHeaders;
        /// <summary>
        /// A list of status code ranges in the format of `100-101`.
        /// </summary>
        public readonly ImmutableArray<string> ExpectedStatusCodeRanges;
        /// <summary>
        /// The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: `30` (normal probing) and `10` (fast probing). The default value is `30`.
        /// </summary>
        public readonly int? IntervalInSeconds;
        /// <summary>
        /// The path used by the monitoring checks. Required when `protocol` is set to `HTTP` or `HTTPS` - cannot be set when `protocol` is set to `TCP`.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The port number used by the monitoring checks.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The protocol used by the monitoring checks, supported values are `HTTP`, `HTTPS` and `TCP`.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If `interval_in_seconds` is set to `30`, then `timeout_in_seconds` can be between `5` and `10`. The default value is `10`. If `interval_in_seconds` is set to `10`, then valid values are between `5` and `9` and `timeout_in_seconds` is required.
        /// </summary>
        public readonly int? TimeoutInSeconds;
        /// <summary>
        /// The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between `0` and `9`. The default value is `3`
        /// </summary>
        public readonly int? ToleratedNumberOfFailures;

        [OutputConstructor]
        private ProfileMonitorConfig(
            ImmutableArray<Outputs.ProfileMonitorConfigCustomHeader> customHeaders,

            ImmutableArray<string> expectedStatusCodeRanges,

            int? intervalInSeconds,

            string? path,

            int port,

            string protocol,

            int? timeoutInSeconds,

            int? toleratedNumberOfFailures)
        {
            CustomHeaders = customHeaders;
            ExpectedStatusCodeRanges = expectedStatusCodeRanges;
            IntervalInSeconds = intervalInSeconds;
            Path = path;
            Port = port;
            Protocol = protocol;
            TimeoutInSeconds = timeoutInSeconds;
            ToleratedNumberOfFailures = toleratedNumberOfFailures;
        }
    }
}
