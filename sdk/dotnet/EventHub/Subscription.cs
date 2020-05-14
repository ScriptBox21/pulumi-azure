// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    /// <summary>
    /// Manages a ServiceBus Subscription.
    /// </summary>
    [Obsolete(@"azure.eventhub.Subscription has been deprecated in favor of azure.servicebus.Subscription")]
    public partial class Subscription : Pulumi.CustomResource
    {
        /// <summary>
        /// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
        /// </summary>
        [Output("autoDeleteOnIdle")]
        public Output<string> AutoDeleteOnIdle { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
        /// </summary>
        [Output("deadLetteringOnMessageExpiration")]
        public Output<bool?> DeadLetteringOnMessageExpiration { get; private set; } = null!;

        /// <summary>
        /// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Output("defaultMessageTtl")]
        public Output<string> DefaultMessageTtl { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
        /// </summary>
        [Output("enableBatchedOperations")]
        public Output<bool?> EnableBatchedOperations { get; private set; } = null!;

        /// <summary>
        /// The name of a Queue or Topic to automatically forward Dead Letter messages to.
        /// </summary>
        [Output("forwardDeadLetteredMessagesTo")]
        public Output<string?> ForwardDeadLetteredMessagesTo { get; private set; } = null!;

        /// <summary>
        /// The name of a Queue or Topic to automatically forward messages to.
        /// </summary>
        [Output("forwardTo")]
        public Output<string?> ForwardTo { get; private set; } = null!;

        /// <summary>
        /// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
        /// </summary>
        [Output("lockDuration")]
        public Output<string> LockDuration { get; private set; } = null!;

        /// <summary>
        /// The maximum number of deliveries.
        /// </summary>
        [Output("maxDeliveryCount")]
        public Output<int> MaxDeliveryCount { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("requiresSession")]
        public Output<bool?> RequiresSession { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/subscription:Subscription", name, args ?? new SubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/subscription:Subscription", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, state, options);
        }
    }

    public sealed class SubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
        /// </summary>
        [Input("deadLetteringOnMessageExpiration")]
        public Input<bool>? DeadLetteringOnMessageExpiration { get; set; }

        /// <summary>
        /// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Input("defaultMessageTtl")]
        public Input<string>? DefaultMessageTtl { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
        /// </summary>
        [Input("enableBatchedOperations")]
        public Input<bool>? EnableBatchedOperations { get; set; }

        /// <summary>
        /// The name of a Queue or Topic to automatically forward Dead Letter messages to.
        /// </summary>
        [Input("forwardDeadLetteredMessagesTo")]
        public Input<string>? ForwardDeadLetteredMessagesTo { get; set; }

        /// <summary>
        /// The name of a Queue or Topic to automatically forward messages to.
        /// </summary>
        [Input("forwardTo")]
        public Input<string>? ForwardTo { get; set; }

        /// <summary>
        /// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
        /// </summary>
        [Input("lockDuration")]
        public Input<string>? LockDuration { get; set; }

        /// <summary>
        /// The maximum number of deliveries.
        /// </summary>
        [Input("maxDeliveryCount", required: true)]
        public Input<int> MaxDeliveryCount { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("requiresSession")]
        public Input<bool>? RequiresSession { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public SubscriptionArgs()
        {
        }
    }

    public sealed class SubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The idle interval after which the topic is automatically deleted as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The minimum duration is `5` minutes or `P5M`.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Subscription has dead letter support when a message expires. Defaults to `false`.
        /// </summary>
        [Input("deadLetteringOnMessageExpiration")]
        public Input<bool>? DeadLetteringOnMessageExpiration { get; set; }

        /// <summary>
        /// The Default message timespan to live as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Input("defaultMessageTtl")]
        public Input<string>? DefaultMessageTtl { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Subscription supports batched operations. Defaults to `false`.
        /// </summary>
        [Input("enableBatchedOperations")]
        public Input<bool>? EnableBatchedOperations { get; set; }

        /// <summary>
        /// The name of a Queue or Topic to automatically forward Dead Letter messages to.
        /// </summary>
        [Input("forwardDeadLetteredMessagesTo")]
        public Input<string>? ForwardDeadLetteredMessagesTo { get; set; }

        /// <summary>
        /// The name of a Queue or Topic to automatically forward messages to.
        /// </summary>
        [Input("forwardTo")]
        public Input<string>? ForwardTo { get; set; }

        /// <summary>
        /// The lock duration for the subscription as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). The default value is `1` minute or `P1M`.
        /// </summary>
        [Input("lockDuration")]
        public Input<string>? LockDuration { get; set; }

        /// <summary>
        /// The maximum number of deliveries.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace to create this Subscription in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// Boolean flag which controls whether this Subscription supports the concept of a session. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("requiresSession")]
        public Input<bool>? RequiresSession { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The name of the ServiceBus Topic to create this Subscription in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public SubscriptionState()
        {
        }
    }
}
