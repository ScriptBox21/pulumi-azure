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
    public sealed class VirtualHubRouteTableRoute
    {
        /// <summary>
        /// A list of destination addresses for this route.
        /// </summary>
        public readonly ImmutableArray<string> Destinations;
        /// <summary>
        /// The type of destinations. Possible values are `CIDR`, `ResourceId` and `Service`.
        /// </summary>
        public readonly string DestinationsType;
        /// <summary>
        /// The name which should be used for this route.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The next hop's resource ID.
        /// </summary>
        public readonly string NextHop;
        /// <summary>
        /// The type of next hop. Currently the only possible value is `ResourceId`. Defaults to `ResourceId`.
        /// </summary>
        public readonly string? NextHopType;

        [OutputConstructor]
        private VirtualHubRouteTableRoute(
            ImmutableArray<string> destinations,

            string destinationsType,

            string name,

            string nextHop,

            string? nextHopType)
        {
            Destinations = destinations;
            DestinationsType = destinationsType;
            Name = name;
            NextHop = nextHop;
            NextHopType = nextHopType;
        }
    }
}
