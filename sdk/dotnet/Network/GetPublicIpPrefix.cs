// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetPublicIpPrefix
    {
        /// <summary>
        /// Use this data source to access information about an existing Public IP Prefix.
        /// </summary>
        public static Task<GetPublicIpPrefixResult> InvokeAsync(GetPublicIpPrefixArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicIpPrefixResult>("azure:network/getPublicIpPrefix:getPublicIpPrefix", args ?? new GetPublicIpPrefixArgs(), options.WithVersion());
    }


    public sealed class GetPublicIpPrefixArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the public IP prefix.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("zones")]
        private List<string>? _zones;
        public List<string> Zones
        {
            get => _zones ?? (_zones = new List<string>());
            set => _zones = value;
        }

        public GetPublicIpPrefixArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicIpPrefixResult
    {
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpPrefix;
        /// <summary>
        /// The supported Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the Public IP prefix resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of bits of the prefix.
        /// </summary>
        public readonly int PrefixLength;
        /// <summary>
        /// The name of the resource group in which to create the public IP.
        /// </summary>
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SKU of the Public IP Prefix.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags to assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetPublicIpPrefixResult(
            string id,

            string ipPrefix,

            string location,

            string name,

            int prefixLength,

            string resourceGroupName,

            string sku,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<string> zones)
        {
            Id = id;
            IpPrefix = ipPrefix;
            Location = location;
            Name = name;
            PrefixLength = prefixLength;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
            Zones = zones;
        }
    }
}
