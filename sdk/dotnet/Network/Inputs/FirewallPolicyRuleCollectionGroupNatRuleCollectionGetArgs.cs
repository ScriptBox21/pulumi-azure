// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallPolicyRuleCollectionGroupNatRuleCollectionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take for the nat rules in this collection. Currently, the only possible value is `Dnat`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this nat rule collection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The priority of the nat rule collection. The range is `100` - `65000`.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleGetArgs>? _rules;

        /// <summary>
        /// A `rule` (nat rule) block as defined above.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleGetArgs>());
            set => _rules = value;
        }

        public FirewallPolicyRuleCollectionGroupNatRuleCollectionGetArgs()
        {
        }
    }
}
