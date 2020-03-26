// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceBus
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/servicebus_topic_authorization_rule.html.markdown.
        /// </summary>
        [Obsolete("Use GetTopicAuthorizationRule.InvokeAsync() instead")]
        public static Task<GetTopicAuthorizationRuleResult> GetTopicAuthorizationRule(GetTopicAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicAuthorizationRuleResult>("azure:servicebus/getTopicAuthorizationRule:getTopicAuthorizationRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetTopicAuthorizationRule
    {
        /// <summary>
        /// Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/servicebus_topic_authorization_rule.html.markdown.
        /// </summary>
        public static Task<GetTopicAuthorizationRuleResult> InvokeAsync(GetTopicAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicAuthorizationRuleResult>("azure:servicebus/getTopicAuthorizationRule:getTopicAuthorizationRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetTopicAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ServiceBus Topic Authorization Rule resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the ServiceBus Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Topic.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetTopicAuthorizationRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetTopicAuthorizationRuleResult
    {
        public readonly bool Listen;
        public readonly bool Manage;
        public readonly string Name;
        public readonly string NamespaceName;
        /// <summary>
        /// The Primary Connection String for the ServiceBus Topic authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The Primary Key for the ServiceBus Topic authorization Rule.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Connection String for the ServiceBus Topic authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// The Secondary Key for the ServiceBus Topic authorization Rule.
        /// </summary>
        public readonly string SecondaryKey;
        public readonly bool Send;
        public readonly string TopicName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTopicAuthorizationRuleResult(
            bool listen,
            bool manage,
            string name,
            string namespaceName,
            string primaryConnectionString,
            string primaryKey,
            string resourceGroupName,
            string secondaryConnectionString,
            string secondaryKey,
            bool send,
            string topicName,
            string id)
        {
            Listen = listen;
            Manage = manage;
            Name = name;
            NamespaceName = namespaceName;
            PrimaryConnectionString = primaryConnectionString;
            PrimaryKey = primaryKey;
            ResourceGroupName = resourceGroupName;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryKey = secondaryKey;
            Send = send;
            TopicName = topicName;
            Id = id;
        }
    }
}
