// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VpnServerConfigurationIpsecPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DH Group, used in IKE Phase 1. Possible values include `DHGroup1`, `DHGroup2`, `DHGroup14`, `DHGroup24`, `DHGroup2048`, `ECP256`, `ECP384` and `None`.
        /// </summary>
        [Input("dhGroup", required: true)]
        public Input<string> DhGroup { get; set; } = null!;

        /// <summary>
        /// The IKE encryption algorithm, used for IKE Phase 2. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128` and `GCMAES256`.
        /// </summary>
        [Input("ikeEncryption", required: true)]
        public Input<string> IkeEncryption { get; set; } = null!;

        /// <summary>
        /// The IKE encryption integrity algorithm, used for IKE Phase 2. Possible values include `GCMAES128`, `GCMAES256`, `MD5`, `SHA1`, `SHA256` and `SHA384`.
        /// </summary>
        [Input("ikeIntegrity", required: true)]
        public Input<string> IkeIntegrity { get; set; } = null!;

        /// <summary>
        /// The IPSec encryption algorithm, used for IKE phase 1. Possible values include `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128`, `GCMAES192`, `GCMAES256` and `None`.
        /// </summary>
        [Input("ipsecEncryption", required: true)]
        public Input<string> IpsecEncryption { get; set; } = null!;

        /// <summary>
        /// The IPSec integrity algorithm, used for IKE phase 1. Possible values include `GCMAES128`, `GCMAES192`, `GCMAES256`, `MD5`, `SHA1` and `SHA256`.
        /// </summary>
        [Input("ipsecIntegrity", required: true)]
        public Input<string> IpsecIntegrity { get; set; } = null!;

        /// <summary>
        /// The Pfs Group, used in IKE Phase 2. Possible values include `ECP256`, `ECP384`, `PFS1`, `PFS2`, `PFS14`, `PFS24`, `PFS2048`, `PFSMM` and `None`.
        /// </summary>
        [Input("pfsGroup", required: true)]
        public Input<string> PfsGroup { get; set; } = null!;

        /// <summary>
        /// The IPSec Security Association payload size in KB for a Site-to-Site VPN tunnel.
        /// </summary>
        [Input("saDataSizeKilobytes", required: true)]
        public Input<int> SaDataSizeKilobytes { get; set; } = null!;

        /// <summary>
        /// The IPSec Security Association lifetime in seconds for a Site-to-Site VPN tunnel.
        /// </summary>
        [Input("saLifetimeSeconds", required: true)]
        public Input<int> SaLifetimeSeconds { get; set; } = null!;

        public VpnServerConfigurationIpsecPolicyArgs()
        {
        }
    }
}
