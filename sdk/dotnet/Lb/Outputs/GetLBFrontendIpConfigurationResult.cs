// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb.Outputs
{

    [OutputType]
    public sealed class GetLBFrontendIpConfigurationResult
    {
        /// <summary>
        /// The id of the Frontend IP Configuration.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the name of the Load Balancer.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private IP Address to assign to the Load Balancer.
        /// </summary>
        public readonly string PrivateIpAddress;
        /// <summary>
        /// The allocation method for the Private IP Address used by this Load Balancer.
        /// </summary>
        public readonly string PrivateIpAddressAllocation;
        /// <summary>
        /// The Private IP Address Version, either `IPv4` or `IPv6`.
        /// </summary>
        public readonly string PrivateIpAddressVersion;
        /// <summary>
        /// The ID of a  Public IP Address which is associated with this Load Balancer.
        /// </summary>
        public readonly string PublicIpAddressId;
        /// <summary>
        /// The ID of the Subnet which is associated with the IP Configuration.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetLBFrontendIpConfigurationResult(
            string id,

            string name,

            string privateIpAddress,

            string privateIpAddressAllocation,

            string privateIpAddressVersion,

            string publicIpAddressId,

            string subnetId,

            ImmutableArray<string> zones)
        {
            Id = id;
            Name = name;
            PrivateIpAddress = privateIpAddress;
            PrivateIpAddressAllocation = privateIpAddressAllocation;
            PrivateIpAddressVersion = privateIpAddressVersion;
            PublicIpAddressId = publicIpAddressId;
            SubnetId = subnetId;
            Zones = zones;
        }
    }
}
