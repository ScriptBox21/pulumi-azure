// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayRewriteRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Unique name of the rewrite rule set block
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("rewriteRules")]
        private InputList<Inputs.ApplicationGatewayRewriteRuleSetRewriteRuleArgs>? _rewriteRules;

        /// <summary>
        /// One or more `rewrite_rule` blocks as defined above.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayRewriteRuleSetRewriteRuleArgs> RewriteRules
        {
            get => _rewriteRules ?? (_rewriteRules = new InputList<Inputs.ApplicationGatewayRewriteRuleSetRewriteRuleArgs>());
            set => _rewriteRules = value;
        }

        public ApplicationGatewayRewriteRuleSetArgs()
        {
        }
    }
}
