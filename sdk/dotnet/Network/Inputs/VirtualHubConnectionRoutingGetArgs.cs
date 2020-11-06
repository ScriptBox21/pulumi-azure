// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VirtualHubConnectionRoutingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the route table associated with this Virtual Hub connection.
        /// </summary>
        [Input("associatedRouteTableId")]
        public Input<string>? AssociatedRouteTableId { get; set; }

        /// <summary>
        /// A `propagated_route_table` block as defined below.
        /// </summary>
        [Input("propagatedRouteTable")]
        public Input<Inputs.VirtualHubConnectionRoutingPropagatedRouteTableGetArgs>? PropagatedRouteTable { get; set; }

        [Input("staticVnetRoutes")]
        private InputList<Inputs.VirtualHubConnectionRoutingStaticVnetRouteGetArgs>? _staticVnetRoutes;

        /// <summary>
        /// A `static_vnet_route` block as defined below.
        /// </summary>
        public InputList<Inputs.VirtualHubConnectionRoutingStaticVnetRouteGetArgs> StaticVnetRoutes
        {
            get => _staticVnetRoutes ?? (_staticVnetRoutes = new InputList<Inputs.VirtualHubConnectionRoutingStaticVnetRouteGetArgs>());
            set => _staticVnetRoutes = value;
        }

        public VirtualHubConnectionRoutingGetArgs()
        {
        }
    }
}
