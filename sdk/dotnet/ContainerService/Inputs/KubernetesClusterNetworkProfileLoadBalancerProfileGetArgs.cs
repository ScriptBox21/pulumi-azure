// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterNetworkProfileLoadBalancerProfileGetArgs : Pulumi.ResourceArgs
    {
        [Input("effectiveOutboundIps")]
        private InputList<string>? _effectiveOutboundIps;

        /// <summary>
        /// The outcome (resource IDs) of the specified arguments.
        /// </summary>
        public InputList<string> EffectiveOutboundIps
        {
            get => _effectiveOutboundIps ?? (_effectiveOutboundIps = new InputList<string>());
            set => _effectiveOutboundIps = value;
        }

        /// <summary>
        /// Count of desired managed outbound IPs for the cluster load balancer. Must be in the range of [1, 100].
        /// </summary>
        [Input("managedOutboundIpCount")]
        public Input<int>? ManagedOutboundIpCount { get; set; }

        [Input("outboundIpAddressIds")]
        private InputList<string>? _outboundIpAddressIds;

        /// <summary>
        /// The ID of the Public IP Addresses which should be used for outbound communication for the cluster load balancer.
        /// </summary>
        public InputList<string> OutboundIpAddressIds
        {
            get => _outboundIpAddressIds ?? (_outboundIpAddressIds = new InputList<string>());
            set => _outboundIpAddressIds = value;
        }

        [Input("outboundIpPrefixIds")]
        private InputList<string>? _outboundIpPrefixIds;

        /// <summary>
        /// The ID of the outbound Public IP Address Prefixes which should be used for the cluster load balancer.
        /// </summary>
        public InputList<string> OutboundIpPrefixIds
        {
            get => _outboundIpPrefixIds ?? (_outboundIpPrefixIds = new InputList<string>());
            set => _outboundIpPrefixIds = value;
        }

        public KubernetesClusterNetworkProfileLoadBalancerProfileGetArgs()
        {
        }
    }
}
