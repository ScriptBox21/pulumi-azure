// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch.Outputs
{

    [OutputType]
    public sealed class PoolNetworkConfiguration
    {
        /// <summary>
        /// A list of inbound NAT pools that can be used to address specific ports on an individual compute node externally. Set as documented in the inbound_nat_pools block below. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<Outputs.PoolNetworkConfigurationEndpointConfiguration> EndpointConfigurations;
        /// <summary>
        /// A list of public ip ids that will be allocated to nodes. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> PublicIps;
        /// <summary>
        /// The ARM resource identifier of the virtual network subnet which the compute nodes of the pool will join. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private PoolNetworkConfiguration(
            ImmutableArray<Outputs.PoolNetworkConfigurationEndpointConfiguration> endpointConfigurations,

            ImmutableArray<string> publicIps,

            string subnetId)
        {
            EndpointConfigurations = endpointConfigurations;
            PublicIps = publicIps;
            SubnetId = subnetId;
        }
    }
}
