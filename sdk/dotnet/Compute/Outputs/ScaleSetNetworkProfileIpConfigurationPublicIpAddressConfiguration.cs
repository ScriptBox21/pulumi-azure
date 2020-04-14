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
    public sealed class ScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration
    {
        /// <summary>
        /// The domain name label for the dns settings.
        /// </summary>
        public readonly string DomainNameLabel;
        /// <summary>
        /// The idle timeout in minutes. This value must be between 4 and 30.
        /// </summary>
        public readonly int IdleTimeout;
        /// <summary>
        /// The name of the public ip address configuration
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration(
            string domainNameLabel,

            int idleTimeout,

            string name)
        {
            DomainNameLabel = domainNameLabel;
            IdleTimeout = idleTimeout;
            Name = name;
        }
    }
}
