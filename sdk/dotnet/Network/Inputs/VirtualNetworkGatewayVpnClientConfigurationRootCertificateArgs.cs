// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VirtualNetworkGatewayVpnClientConfigurationRootCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-defined name of the revoked certificate.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The SHA1 thumbprint of the certificate to be
        /// revoked.
        /// </summary>
        [Input("publicCertData", required: true)]
        public Input<string> PublicCertData { get; set; } = null!;

        public VirtualNetworkGatewayVpnClientConfigurationRootCertificateArgs()
        {
        }
    }
}
