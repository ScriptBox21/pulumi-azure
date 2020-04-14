// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Outputs
{

    [OutputType]
    public sealed class IoTHubIpFilterRule
    {
        /// <summary>
        /// The desired action for requests captured by this rule. Possible values are  `Accept`, `Reject`
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The IP address range in CIDR notation for the rule.
        /// </summary>
        public readonly string IpMask;
        /// <summary>
        /// The name of the filter.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private IoTHubIpFilterRule(
            string action,

            string ipMask,

            string name)
        {
            Action = action;
            IpMask = ipMask;
            Name = name;
        }
    }
}
