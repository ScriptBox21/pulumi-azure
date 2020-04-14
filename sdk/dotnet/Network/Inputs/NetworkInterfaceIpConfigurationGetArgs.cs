// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class NetworkInterfaceIpConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name used for this IP Configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Is this the Primary IP Configuration? Must be `true` for the first `ip_configuration` when multiple are specified. Defaults to `false`.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// The Static IP Address which should be used.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The allocation method used for the Private IP Address. Possible values are `Dynamic` and `Static`.
        /// </summary>
        [Input("privateIpAddressAllocation", required: true)]
        public Input<string> PrivateIpAddressAllocation { get; set; } = null!;

        /// <summary>
        /// The IP Version to use. Possible values are `IPv4` or `IPv6`. Defaults to `IPv4`.
        /// </summary>
        [Input("privateIpAddressVersion")]
        public Input<string>? PrivateIpAddressVersion { get; set; }

        /// <summary>
        /// Reference to a Public IP Address to associate with this NIC
        /// </summary>
        [Input("publicIpAddressId")]
        public Input<string>? PublicIpAddressId { get; set; }

        /// <summary>
        /// The ID of the Subnet where this Network Interface should be located in.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public NetworkInterfaceIpConfigurationGetArgs()
        {
        }
    }
}
