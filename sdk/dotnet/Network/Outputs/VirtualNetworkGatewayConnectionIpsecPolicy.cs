// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class VirtualNetworkGatewayConnectionIpsecPolicy
    {
        /// <summary>
        /// The DH group used in IKE phase 1 for initial SA. Valid
        /// options are `DHGroup1`, `DHGroup14`, `DHGroup2`, `DHGroup2048`, `DHGroup24`,
        /// `ECP256`, `ECP384`, or `None`.
        /// </summary>
        public readonly string DhGroup;
        /// <summary>
        /// The IKE encryption algorithm. Valid
        /// options are `AES128`, `AES192`, `AES256`, `DES`, or `DES3`.
        /// </summary>
        public readonly string IkeEncryption;
        /// <summary>
        /// The IKE integrity algorithm. Valid
        /// options are `MD5`, `SHA1`, `SHA256`, or `SHA384`.
        /// </summary>
        public readonly string IkeIntegrity;
        /// <summary>
        /// The IPSec encryption algorithm. Valid
        /// options are `AES128`, `AES192`, `AES256`, `DES`, `DES3`, `GCMAES128`, `GCMAES192`, `GCMAES256`, or `None`.
        /// </summary>
        public readonly string IpsecEncryption;
        /// <summary>
        /// The IPSec integrity algorithm. Valid
        /// options are `GCMAES128`, `GCMAES192`, `GCMAES256`, `MD5`, `SHA1`, or `SHA256`.
        /// </summary>
        public readonly string IpsecIntegrity;
        /// <summary>
        /// The DH group used in IKE phase 2 for new child SA.
        /// Valid options are `ECP256`, `ECP384`, `PFS1`, `PFS2`, `PFS2048`, `PFS24`,
        /// or `None`.
        /// </summary>
        public readonly string PfsGroup;
        /// <summary>
        /// The IPSec SA payload size in KB. Must be at least
        /// `1024` KB. Defaults to `102400000` KB.
        /// </summary>
        public readonly int? SaDatasize;
        /// <summary>
        /// The IPSec SA lifetime in seconds. Must be at least
        /// `300` seconds. Defaults to `27000` seconds.
        /// </summary>
        public readonly int? SaLifetime;

        [OutputConstructor]
        private VirtualNetworkGatewayConnectionIpsecPolicy(
            string dhGroup,

            string ikeEncryption,

            string ikeIntegrity,

            string ipsecEncryption,

            string ipsecIntegrity,

            string pfsGroup,

            int? saDatasize,

            int? saLifetime)
        {
            DhGroup = dhGroup;
            IkeEncryption = ikeEncryption;
            IkeIntegrity = ikeIntegrity;
            IpsecEncryption = ipsecEncryption;
            IpsecIntegrity = ipsecIntegrity;
            PfsGroup = pfsGroup;
            SaDatasize = saDatasize;
            SaLifetime = saLifetime;
        }
    }
}
