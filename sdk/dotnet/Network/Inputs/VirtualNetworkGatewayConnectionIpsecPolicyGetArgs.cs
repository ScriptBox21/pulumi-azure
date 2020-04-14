// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VirtualNetworkGatewayConnectionIpsecPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DH group used in IKE phase 1 for initial SA. Valid
        /// options are `DHGroup1`, `DHGroup14`, `DHGroup2`, `DHGroup2048`, `DHGroup24`,
        /// `ECP256`, `ECP384`, or `None`.
        /// </summary>
        [Input("dhGroup", required: true)]
        public Input<string> DhGroup { get; set; } = null!;

        /// <summary>
        /// The IKE encryption algorithm. Valid
        /// options are `AES128`, `AES192`, `AES256`, `DES`, or `DES3`.
        /// </summary>
        [Input("ikeEncryption", required: true)]
        public Input<string> IkeEncryption { get; set; } = null!;

        /// <summary>
        /// The IKE integrity algorithm. Valid
        /// options are `MD5`, `SHA1`, `SHA256`, or `SHA384`.
        /// </summary>
        [Input("ikeIntegrity", required: true)]
        public Input<string> IkeIntegrity { get; set; } = null!;

        /// <summary>
        /// The IPSec encryption algorithm. Valid
        /// options are `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128`, `GCMAES192`, `GCMAES256`, or `None`.
        /// </summary>
        [Input("ipsecEncryption", required: true)]
        public Input<string> IpsecEncryption { get; set; } = null!;

        /// <summary>
        /// The IPSec integrity algorithm. Valid
        /// options are `GCMAES128`, `GCMAES192`, `GCMAES256`, `MD5`, `SHA1`, or `SHA256`.
        /// </summary>
        [Input("ipsecIntegrity", required: true)]
        public Input<string> IpsecIntegrity { get; set; } = null!;

        /// <summary>
        /// The DH group used in IKE phase 2 for new child SA.
        /// Valid options are `ECP256`, `ECP384`, `PFS1`, `PFS2`, `PFS2048`, `PFS24`,
        /// or `None`.
        /// </summary>
        [Input("pfsGroup", required: true)]
        public Input<string> PfsGroup { get; set; } = null!;

        /// <summary>
        /// The IPSec SA payload size in KB. Must be at least
        /// `1024` KB. Defaults to `102400000` KB.
        /// </summary>
        [Input("saDatasize")]
        public Input<int>? SaDatasize { get; set; }

        /// <summary>
        /// The IPSec SA lifetime in seconds. Must be at least
        /// `300` seconds. Defaults to `27000` seconds.
        /// </summary>
        [Input("saLifetime")]
        public Input<int>? SaLifetime { get; set; }

        public VirtualNetworkGatewayConnectionIpsecPolicyGetArgs()
        {
        }
    }
}
