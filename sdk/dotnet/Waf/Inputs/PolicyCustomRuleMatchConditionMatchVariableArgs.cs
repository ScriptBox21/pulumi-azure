// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Waf.Inputs
{

    public sealed class PolicyCustomRuleMatchConditionMatchVariableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes field of the matchVariable collection
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        /// <summary>
        /// The name of the Match Variable
        /// </summary>
        [Input("variableName", required: true)]
        public Input<string> VariableName { get; set; } = null!;

        public PolicyCustomRuleMatchConditionMatchVariableArgs()
        {
        }
    }
}
