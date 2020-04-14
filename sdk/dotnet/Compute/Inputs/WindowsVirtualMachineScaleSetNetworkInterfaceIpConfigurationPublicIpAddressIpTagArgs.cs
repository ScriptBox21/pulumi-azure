// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP Tag associated with the Public IP, such as `SQL` or `Storage`.
        /// </summary>
        [Input("tag", required: true)]
        public Input<string> Tag { get; set; } = null!;

        /// <summary>
        /// The Type of IP Tag, such as `FirstPartyUsage`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagArgs()
        {
        }
    }
}
