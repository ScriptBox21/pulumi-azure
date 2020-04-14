// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DevTest.Inputs
{

    public sealed class VirtualNetworkSubnetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Can this subnet be used for creating Virtual Machines? Possible values are `Allow`, `Default` and `Deny`.
        /// </summary>
        [Input("useInVirtualMachineCreation")]
        public Input<string>? UseInVirtualMachineCreation { get; set; }

        /// <summary>
        /// Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are `Allow`, `Default` and `Deny`.
        /// </summary>
        [Input("usePublicIpAddress")]
        public Input<string>? UsePublicIpAddress { get; set; }

        public VirtualNetworkSubnetGetArgs()
        {
        }
    }
}
