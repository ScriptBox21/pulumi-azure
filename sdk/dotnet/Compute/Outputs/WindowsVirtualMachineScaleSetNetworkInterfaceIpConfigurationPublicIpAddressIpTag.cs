// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag
    {
        /// <summary>
        /// The IP Tag associated with the Public IP, such as `SQL` or `Storage`.
        /// </summary>
        public readonly string Tag;
        /// <summary>
        /// The Type of IP Tag, such as `FirstPartyUsage`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private WindowsVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag(
            string tag,

            string type)
        {
            Tag = tag;
            Type = type;
        }
    }
}
