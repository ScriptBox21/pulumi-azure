// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class GetVirtualNetworkGatewayIpConfigurationResult
    {
        /// <summary>
        /// Specifies the name of the Virtual Network Gateway.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines how the private IP address
        /// of the gateways virtual interface is assigned.
        /// </summary>
        public readonly string PrivateIpAddressAllocation;
        /// <summary>
        /// The ID of the Public IP Address associated
        /// with the Virtual Network Gateway.
        /// </summary>
        public readonly string PublicIpAddressId;
        /// <summary>
        /// The ID of the gateway subnet of a virtual network in
        /// which the virtual network gateway will be created. It is mandatory that
        /// the associated subnet is named `GatewaySubnet`. Therefore, each virtual
        /// network can contain at most a single Virtual Network Gateway.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private GetVirtualNetworkGatewayIpConfigurationResult(
            string name,

            string privateIpAddressAllocation,

            string publicIpAddressId,

            string subnetId)
        {
            Name = name;
            PrivateIpAddressAllocation = privateIpAddressAllocation;
            PublicIpAddressId = publicIpAddressId;
            SubnetId = subnetId;
        }
    }
}
