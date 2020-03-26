// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DevTest
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Dev Test Lab Virtual Network.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/dev_test_virtual_network.html.markdown.
        /// </summary>
        [Obsolete("Use GetVirtualNetwork.InvokeAsync() instead")]
        public static Task<GetVirtualNetworkResult> GetVirtualNetwork(GetVirtualNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkResult>("azure:devtest/getVirtualNetwork:getVirtualNetwork", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetVirtualNetwork
    {
        /// <summary>
        /// Use this data source to access information about an existing Dev Test Lab Virtual Network.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/dev_test_virtual_network.html.markdown.
        /// </summary>
        public static Task<GetVirtualNetworkResult> InvokeAsync(GetVirtualNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkResult>("azure:devtest/getVirtualNetwork:getVirtualNetwork", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVirtualNetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Dev Test Lab.
        /// </summary>
        [Input("labName", required: true)]
        public string LabName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Virtual Network.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group that contains the Virtual Network.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVirtualNetworkResult
    {
        /// <summary>
        /// The list of subnets enabled for the virtual network as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualNetworkAllowedSubnetsResult> AllowedSubnets;
        public readonly string LabName;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The list of permission overrides for the subnets as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualNetworkSubnetOverridesResult> SubnetOverrides;
        /// <summary>
        /// The unique immutable identifier of the virtual network.
        /// </summary>
        public readonly string UniqueIdentifier;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVirtualNetworkResult(
            ImmutableArray<Outputs.GetVirtualNetworkAllowedSubnetsResult> allowedSubnets,
            string labName,
            string name,
            string resourceGroupName,
            ImmutableArray<Outputs.GetVirtualNetworkSubnetOverridesResult> subnetOverrides,
            string uniqueIdentifier,
            string id)
        {
            AllowedSubnets = allowedSubnets;
            LabName = labName;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SubnetOverrides = subnetOverrides;
            UniqueIdentifier = uniqueIdentifier;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetVirtualNetworkAllowedSubnetsResult
    {
        /// <summary>
        /// Indicates if this subnet allows public IP addresses. Possible values are `Allow`, `Default` and `Deny`.
        /// </summary>
        public readonly string AllowPublicIp;
        /// <summary>
        /// The name of the subnet.
        /// </summary>
        public readonly string LabSubnetName;
        /// <summary>
        /// The resource identifier for the subnet.
        /// </summary>
        public readonly string ResourceId;

        [OutputConstructor]
        private GetVirtualNetworkAllowedSubnetsResult(
            string allowPublicIp,
            string labSubnetName,
            string resourceId)
        {
            AllowPublicIp = allowPublicIp;
            LabSubnetName = labSubnetName;
            ResourceId = resourceId;
        }
    }

    [OutputType]
    public sealed class GetVirtualNetworkSubnetOverridesResult
    {
        /// <summary>
        /// The name of the subnet.
        /// </summary>
        public readonly string LabSubnetName;
        /// <summary>
        /// The resource identifier for the subnet.
        /// </summary>
        public readonly string ResourceId;
        /// <summary>
        /// Indicates if the subnet can be used for VM creation.  Possible values are `Allow`, `Default` and `Deny`.
        /// </summary>
        public readonly string UseInVmCreationPermission;
        public readonly string UsePublicIpAddressPermission;
        /// <summary>
        /// The virtual network pool associated with this subnet.
        /// </summary>
        public readonly string VirtualNetworkPoolName;

        [OutputConstructor]
        private GetVirtualNetworkSubnetOverridesResult(
            string labSubnetName,
            string resourceId,
            string useInVmCreationPermission,
            string usePublicIpAddressPermission,
            string virtualNetworkPoolName)
        {
            LabSubnetName = labSubnetName;
            ResourceId = resourceId;
            UseInVmCreationPermission = useInVmCreationPermission;
            UsePublicIpAddressPermission = usePublicIpAddressPermission;
            VirtualNetworkPoolName = virtualNetworkPoolName;
        }
    }
    }
}
