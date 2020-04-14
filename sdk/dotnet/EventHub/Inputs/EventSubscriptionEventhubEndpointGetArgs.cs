// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub.Inputs
{

    public sealed class EventSubscriptionEventhubEndpointGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the eventhub where the Event Subscription will receive events.
        /// </summary>
        [Input("eventhubId", required: true)]
        public Input<string> EventhubId { get; set; } = null!;

        public EventSubscriptionEventhubEndpointGetArgs()
        {
        }
    }
}
