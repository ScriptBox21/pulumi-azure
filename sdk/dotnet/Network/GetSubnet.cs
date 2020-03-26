// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Subnet within a Virtual Network.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/subnet.html.markdown.
        /// </summary>
        [Obsolete("Use GetSubnet.InvokeAsync() instead")]
        public static Task<GetSubnetResult> GetSubnet(GetSubnetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetResult>("azure:network/getSubnet:getSubnet", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSubnet
    {
        /// <summary>
        /// Use this data source to access information about an existing Subnet within a Virtual Network.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/subnet.html.markdown.
        /// </summary>
        public static Task<GetSubnetResult> InvokeAsync(GetSubnetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetResult>("azure:network/getSubnet:getSubnet", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSubnetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Subnet.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Virtual Network is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Virtual Network this Subnet is located within.
        /// </summary>
        [Input("virtualNetworkName", required: true)]
        public string VirtualNetworkName { get; set; } = null!;

        public GetSubnetArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSubnetResult
    {
        /// <summary>
        /// The address prefix used for the subnet.
        /// </summary>
        public readonly string AddressPrefix;
        /// <summary>
        /// Enable or Disable network policies for the private link endpoint on the subnet.
        /// </summary>
        public readonly bool EnforcePrivateLinkEndpointNetworkPolicies;
        /// <summary>
        /// Enable or Disable network policies for the private link service on the subnet.
        /// </summary>
        public readonly bool EnforcePrivateLinkServiceNetworkPolicies;
        public readonly string Name;
        /// <summary>
        /// The ID of the Network Security Group associated with the subnet.
        /// </summary>
        public readonly string NetworkSecurityGroupId;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The ID of the Route Table associated with this subnet.
        /// </summary>
        public readonly string RouteTableId;
        /// <summary>
        /// A list of Service Endpoints within this subnet.
        /// </summary>
        public readonly ImmutableArray<string> ServiceEndpoints;
        public readonly string VirtualNetworkName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSubnetResult(
            string addressPrefix,
            bool enforcePrivateLinkEndpointNetworkPolicies,
            bool enforcePrivateLinkServiceNetworkPolicies,
            string name,
            string networkSecurityGroupId,
            string resourceGroupName,
            string routeTableId,
            ImmutableArray<string> serviceEndpoints,
            string virtualNetworkName,
            string id)
        {
            AddressPrefix = addressPrefix;
            EnforcePrivateLinkEndpointNetworkPolicies = enforcePrivateLinkEndpointNetworkPolicies;
            EnforcePrivateLinkServiceNetworkPolicies = enforcePrivateLinkServiceNetworkPolicies;
            Name = name;
            NetworkSecurityGroupId = networkSecurityGroupId;
            ResourceGroupName = resourceGroupName;
            RouteTableId = routeTableId;
            ServiceEndpoints = serviceEndpoints;
            VirtualNetworkName = virtualNetworkName;
            Id = id;
        }
    }
}
