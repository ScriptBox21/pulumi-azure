// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Prefix which should be used for the Domain Name Label for each Virtual Machine Instance. Azure concatenates the Domain Name Label and Virtual Machine Index to create a unique Domain Name Label for each Virtual Machine.
        /// </summary>
        [Input("domainNameLabel")]
        public Input<string>? DomainNameLabel { get; set; }

        /// <summary>
        /// The Idle Timeout in Minutes for the Public IP Address. Possible values are in the range `4` to `32`.
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        [Input("ipTags")]
        private InputList<Inputs.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagGetArgs>? _ipTags;

        /// <summary>
        /// One or more `ip_tag` blocks as defined above.
        /// </summary>
        public InputList<Inputs.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagGetArgs> IpTags
        {
            get => _ipTags ?? (_ipTags = new InputList<Inputs.WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagGetArgs>());
            set => _ipTags = value;
        }

        /// <summary>
        /// The Name of the Public IP Address Configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the Public IP Address Prefix from where Public IP Addresses should be allocated. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publicIpPrefixId")]
        public Input<string>? PublicIpPrefixId { get; set; }

        public WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressGetArgs()
        {
        }
    }
}
