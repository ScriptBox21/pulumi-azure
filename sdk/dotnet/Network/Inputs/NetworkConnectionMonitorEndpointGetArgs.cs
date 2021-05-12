// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class NetworkConnectionMonitorEndpointGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address or domain name of the Network Connection Monitor endpoint.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The test coverage for the Network Connection Monitor endpoint. Possible values are `AboveAverage`, `Average`, `BelowAverage`, `Default`, `Full` and `Low`.
        /// </summary>
        [Input("coverageLevel")]
        public Input<string>? CoverageLevel { get; set; }

        [Input("excludedIpAddresses")]
        private InputList<string>? _excludedIpAddresses;

        /// <summary>
        /// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be excluded to the Network Connection Monitor endpoint.
        /// </summary>
        public InputList<string> ExcludedIpAddresses
        {
            get => _excludedIpAddresses ?? (_excludedIpAddresses = new InputList<string>());
            set => _excludedIpAddresses = value;
        }

        /// <summary>
        /// A `filter` block as defined below.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.NetworkConnectionMonitorEndpointFilterGetArgs>? Filter { get; set; }

        [Input("includedIpAddresses")]
        private InputList<string>? _includedIpAddresses;

        /// <summary>
        /// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be included to the Network Connection Monitor endpoint.
        /// </summary>
        public InputList<string> IncludedIpAddresses
        {
            get => _includedIpAddresses ?? (_includedIpAddresses = new InputList<string>());
            set => _includedIpAddresses = value;
        }

        /// <summary>
        /// The name of the endpoint for the Network Connection Monitor .
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The resource ID which is used as the endpoint by the Network Connection Monitor.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The endpoint type of the Network Connection Monitor. Possible values are `AzureSubnet`, `AzureVM`, `AzureVNet`, `ExternalAddress`, `MMAWorkspaceMachine` and `MMAWorkspaceNetwork`.
        /// </summary>
        [Input("targetResourceType")]
        public Input<string>? TargetResourceType { get; set; }

        /// <summary>
        /// The ID of the Virtual Machine which is used as the endpoint by the Network Connection Monitor. This property is deprecated in favour of `target_resource_id`.
        /// </summary>
        [Input("virtualMachineId")]
        public Input<string>? VirtualMachineId { get; set; }

        public NetworkConnectionMonitorEndpointGetArgs()
        {
        }
    }
}
