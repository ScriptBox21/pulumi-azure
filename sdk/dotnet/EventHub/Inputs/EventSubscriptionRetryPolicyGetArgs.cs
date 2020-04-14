// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub.Inputs
{

    public sealed class EventSubscriptionRetryPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the time to live (in minutes) for events.
        /// </summary>
        [Input("eventTimeToLive", required: true)]
        public Input<int> EventTimeToLive { get; set; } = null!;

        /// <summary>
        /// Specifies the maximum number of delivery retry attempts for events.
        /// </summary>
        [Input("maxDeliveryAttempts", required: true)]
        public Input<int> MaxDeliveryAttempts { get; set; } = null!;

        public EventSubscriptionRetryPolicyGetArgs()
        {
        }
    }
}
