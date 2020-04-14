// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-defined name of the revoked certificate.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("thumbprint", required: true)]
        public Input<string> Thumbprint { get; set; } = null!;

        public VirtualNetworkGatewayVpnClientConfigurationRevokedCertificateGetArgs()
        {
        }
    }
}
