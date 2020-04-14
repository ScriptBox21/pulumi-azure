// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayRewriteRuleSetRewriteRuleConditionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Perform a case in-sensitive comparison. Defaults to `false`
        /// </summary>
        [Input("ignoreCase")]
        public Input<bool>? IgnoreCase { get; set; }

        /// <summary>
        /// Negate the result of the condition evaluation. Defaults to `false`
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// The pattern, either fixed string or regular expression, that evaluates the truthfulness of the condition.
        /// </summary>
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        /// <summary>
        /// The [variable](https://docs.microsoft.com/en-us/azure/application-gateway/rewrite-http-headers#server-variables) of the condition.
        /// </summary>
        [Input("variable", required: true)]
        public Input<string> Variable { get; set; } = null!;

        public ApplicationGatewayRewriteRuleSetRewriteRuleConditionGetArgs()
        {
        }
    }
}
