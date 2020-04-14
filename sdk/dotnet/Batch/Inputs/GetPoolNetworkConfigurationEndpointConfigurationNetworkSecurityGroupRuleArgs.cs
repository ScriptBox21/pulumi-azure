// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch.Inputs
{

    public sealed class GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The action that should be taken for a specified IP address, subnet range or tag.
        /// </summary>
        [Input("access", required: true)]
        public string Access { get; set; } = null!;

        /// <summary>
        /// The priority for this rule.
        /// </summary>
        [Input("priority", required: true)]
        public int Priority { get; set; }

        /// <summary>
        /// The source address prefix or tag to match for the rule.
        /// </summary>
        [Input("sourceAddressPrefix", required: true)]
        public string SourceAddressPrefix { get; set; } = null!;

        public GetPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRuleArgs()
        {
        }
    }
}
