// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor.Inputs
{

    public sealed class FirewallPolicyManagedRuleOverrideExclusionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The variable type to be excluded. Possible values are `QueryStringArgNames`, `RequestBodyPostArgNames`, `RequestCookieNames`, `RequestHeaderNames`.
        /// </summary>
        [Input("matchVariable", required: true)]
        public Input<string> MatchVariable { get; set; } = null!;

        /// <summary>
        /// Comparison operator to apply to the selector when specifying which elements in the collection this exclusion applies to. Possible values are: `Equals`, `Contains`, `StartsWith`, `EndsWith`, `EqualsAny`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// Selector for the value in the `match_variable` attribute this exclusion applies to.
        /// </summary>
        [Input("selector", required: true)]
        public Input<string> Selector { get; set; } = null!;

        public FirewallPolicyManagedRuleOverrideExclusionGetArgs()
        {
        }
    }
}
