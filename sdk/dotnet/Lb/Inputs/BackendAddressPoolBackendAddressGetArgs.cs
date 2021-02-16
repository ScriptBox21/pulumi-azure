// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb.Inputs
{

    public sealed class BackendAddressPoolBackendAddressGetArgs : Pulumi.ResourceArgs
    {
        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Backend Address Pool.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("virtualNetworkId", required: true)]
        public Input<string> VirtualNetworkId { get; set; } = null!;

        public BackendAddressPoolBackendAddressGetArgs()
        {
        }
    }
}
