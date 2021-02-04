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
    public sealed class GetBackendAddressPoolBackendAddressResult
    {
        /// <summary>
        /// The IP address pre-allocated for this Backend Address with in the Virtual Network of `virtual_network_id`.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Specifies the name of the Backend Address Pool.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the Virtual Network that is pre-allocated for this Backend Address.
        /// </summary>
        public readonly string VirtualNetworkId;

        [OutputConstructor]
        private GetBackendAddressPoolBackendAddressResult(
            string ipAddress,

            string name,

            string virtualNetworkId)
        {
            IpAddress = ipAddress;
            Name = name;
            VirtualNetworkId = virtualNetworkId;
        }
    }
}
