// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayFrontendIpConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the Frontend IP Configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Private IP Address to use for the Application Gateway.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The Allocation Method for the Private IP Address. Possible values are `Dynamic` and `Static`.
        /// </summary>
        [Input("privateIpAddressAllocation")]
        public Input<string>? PrivateIpAddressAllocation { get; set; }

        /// <summary>
        /// The ID of a Public IP Address which the Application Gateway should use.
        /// </summary>
        [Input("publicIpAddressId")]
        public Input<string>? PublicIpAddressId { get; set; }

        /// <summary>
        /// The ID of the Subnet which the Application Gateway should be connected to.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public ApplicationGatewayFrontendIpConfigurationArgs()
        {
        }
    }
}
