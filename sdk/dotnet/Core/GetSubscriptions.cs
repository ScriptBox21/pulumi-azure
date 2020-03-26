// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about all the Subscriptions currently available.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/subscriptions.html.markdown.
        /// </summary>
        [Obsolete("Use GetSubscriptions.InvokeAsync() instead")]
        public static Task<GetSubscriptionsResult> GetSubscriptions(GetSubscriptionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionsResult>("azure:core/getSubscriptions:getSubscriptions", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSubscriptions
    {
        /// <summary>
        /// Use this data source to access information about all the Subscriptions currently available.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/subscriptions.html.markdown.
        /// </summary>
        public static Task<GetSubscriptionsResult> InvokeAsync(GetSubscriptionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionsResult>("azure:core/getSubscriptions:getSubscriptions", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSubscriptionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A case-insensitive value which must be contained within the `display_name` field, used to filter the results
        /// </summary>
        [Input("displayNameContains")]
        public string? DisplayNameContains { get; set; }

        /// <summary>
        /// A case-insensitive prefix which can be used to filter on the `display_name` field
        /// </summary>
        [Input("displayNamePrefix")]
        public string? DisplayNamePrefix { get; set; }

        public GetSubscriptionsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSubscriptionsResult
    {
        public readonly string? DisplayNameContains;
        public readonly string? DisplayNamePrefix;
        /// <summary>
        /// One or more `subscription` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubscriptionsSubscriptionsResult> Subscriptions;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSubscriptionsResult(
            string? displayNameContains,
            string? displayNamePrefix,
            ImmutableArray<Outputs.GetSubscriptionsSubscriptionsResult> subscriptions,
            string id)
        {
            DisplayNameContains = displayNameContains;
            DisplayNamePrefix = displayNamePrefix;
            Subscriptions = subscriptions;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetSubscriptionsSubscriptionsResult
    {
        /// <summary>
        /// The subscription display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The subscription location placement ID.
        /// </summary>
        public readonly string LocationPlacementId;
        /// <summary>
        /// The subscription quota ID.
        /// </summary>
        public readonly string QuotaId;
        /// <summary>
        /// The subscription spending limit.
        /// </summary>
        public readonly string SpendingLimit;
        /// <summary>
        /// The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The subscription GUID.
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// The subscription tenant ID.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetSubscriptionsSubscriptionsResult(
            string displayName,
            string locationPlacementId,
            string quotaId,
            string spendingLimit,
            string state,
            string subscriptionId,
            string tenantId)
        {
            DisplayName = displayName;
            LocationPlacementId = locationPlacementId;
            QuotaId = quotaId;
            SpendingLimit = spendingLimit;
            State = state;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
        }
    }
    }
}
