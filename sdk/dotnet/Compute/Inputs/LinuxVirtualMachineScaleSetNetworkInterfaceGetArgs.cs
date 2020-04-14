// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class LinuxVirtualMachineScaleSetNetworkInterfaceGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// A list of IP Addresses of DNS Servers which should be assigned to the Network Interface.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// Does this Network Interface support Accelerated Networking? Defaults to `false`.
        /// </summary>
        [Input("enableAcceleratedNetworking")]
        public Input<bool>? EnableAcceleratedNetworking { get; set; }

        /// <summary>
        /// Does this Network Interface support IP Forwarding? Defaults to `false`.
        /// </summary>
        [Input("enableIpForwarding")]
        public Input<bool>? EnableIpForwarding { get; set; }

        [Input("ipConfigurations", required: true)]
        private InputList<Inputs.LinuxVirtualMachineScaleSetNetworkInterfaceIpConfigurationGetArgs>? _ipConfigurations;

        /// <summary>
        /// One or more `ip_configuration` blocks as defined above.
        /// </summary>
        public InputList<Inputs.LinuxVirtualMachineScaleSetNetworkInterfaceIpConfigurationGetArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.LinuxVirtualMachineScaleSetNetworkInterfaceIpConfigurationGetArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The Name which should be used for this Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of a Network Security Group which should be assigned to this Network Interface.
        /// </summary>
        [Input("networkSecurityGroupId")]
        public Input<string>? NetworkSecurityGroupId { get; set; }

        /// <summary>
        /// Is this the Primary IP Configuration?
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        public LinuxVirtualMachineScaleSetNetworkInterfaceGetArgs()
        {
        }
    }
}
