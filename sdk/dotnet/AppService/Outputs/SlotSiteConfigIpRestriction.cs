// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class SlotSiteConfigIpRestriction
    {
        /// <summary>
        /// The IP Address used for this IP Restriction.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// (Optional.The Virtual Network Subnet ID used for this IP Restriction.
        /// </summary>
        public readonly string? VirtualNetworkSubnetId;

        [OutputConstructor]
        private SlotSiteConfigIpRestriction(
            string? ipAddress,

            string? virtualNetworkSubnetId)
        {
            IpAddress = ipAddress;
            VirtualNetworkSubnetId = virtualNetworkSubnetId;
        }
    }
}
