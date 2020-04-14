// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AnalysisServices.Outputs
{

    [OutputType]
    public sealed class ServerIpv4FirewallRule
    {
        /// <summary>
        /// Specifies the name of the firewall rule.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// End of the firewall rule range as IPv4 address.
        /// </summary>
        public readonly string RangeEnd;
        /// <summary>
        /// Start of the firewall rule range as IPv4 address.
        /// </summary>
        public readonly string RangeStart;

        [OutputConstructor]
        private ServerIpv4FirewallRule(
            string name,

            string rangeEnd,

            string rangeStart)
        {
            Name = name;
            RangeEnd = rangeEnd;
            RangeStart = rangeStart;
        }
    }
}
