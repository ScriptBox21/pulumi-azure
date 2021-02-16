// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Outputs
{

    [OutputType]
    public sealed class SecurityDeviceGroupRangeRule
    {
        /// <summary>
        /// Specifies the time range. represented in ISO 8601 duration format.
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// The maximum threshold in the given time window.
        /// </summary>
        public readonly int Max;
        /// <summary>
        /// The minimum threshold in the given time window.
        /// </summary>
        public readonly int Min;
        /// <summary>
        /// The type of supported rule type. Possible Values are `ActiveConnectionsNotInAllowedRange`, `AmqpC2DMessagesNotInAllowedRange`, `MqttC2DMessagesNotInAllowedRange`, `HttpC2DMessagesNotInAllowedRange`, `AmqpC2DRejectedMessagesNotInAllowedRange`, `MqttC2DRejectedMessagesNotInAllowedRange`, `HttpC2DRejectedMessagesNotInAllowedRange`, `AmqpD2CMessagesNotInAllowedRange`, `MqttD2CMessagesNotInAllowedRange`, `HttpD2CMessagesNotInAllowedRange`, `DirectMethodInvokesNotInAllowedRange`, `FailedLocalLoginsNotInAllowedRange`, `FileUploadsNotInAllowedRange`, `QueuePurgesNotInAllowedRange`, `TwinUpdatesNotInAllowedRange` and `UnauthorizedOperationsNotInAllowedRange`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SecurityDeviceGroupRangeRule(
            string duration,

            int max,

            int min,

            string type)
        {
            Duration = duration;
            Max = max;
            Min = min;
            Type = type;
        }
    }
}
