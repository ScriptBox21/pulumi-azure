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
    public sealed class VirtualNetworkSubnet
    {
        /// <summary>
        /// The address prefix to use for the subnet.
        /// </summary>
        public readonly string AddressPrefix;
        /// <summary>
        /// The Resource ID of DDoS Protection Plan.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the virtual network. Changing this forces a
        /// new resource to be created.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Network Security Group to associate with
        /// the subnet. (Referenced by `id`, ie. `azurerm_network_security_group.example.id`)
        /// </summary>
        public readonly string? SecurityGroup;

        [OutputConstructor]
        private VirtualNetworkSubnet(
            string addressPrefix,

            string? id,

            string name,

            string? securityGroup)
        {
            AddressPrefix = addressPrefix;
            Id = id;
            Name = name;
            SecurityGroup = securityGroup;
        }
    }
}
