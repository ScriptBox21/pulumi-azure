// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VpnGatewayBgpSettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ASN of the BGP Speaker. Changing this forces a new resource to be created.
        /// </summary>
        [Input("asn", required: true)]
        public Input<int> Asn { get; set; } = null!;

        /// <summary>
        /// The Address which should be used for the BGP Peering.
        /// </summary>
        [Input("bgpPeeringAddress")]
        public Input<string>? BgpPeeringAddress { get; set; }

        /// <summary>
        /// The weight added to Routes learned from this BGP Speaker. Changing this forces a new resource to be created.
        /// </summary>
        [Input("peerWeight", required: true)]
        public Input<int> PeerWeight { get; set; } = null!;

        public VpnGatewayBgpSettingArgs()
        {
        }
    }
}
