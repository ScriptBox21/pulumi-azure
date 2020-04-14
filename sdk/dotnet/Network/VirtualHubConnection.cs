// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a Connection for a Virtual Hub.
    /// </summary>
    public partial class VirtualHubConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Is the Virtual Hub traffic allowed to transit via the Remote Virtual Network? Changing this forces a new resource to be created.
        /// </summary>
        [Output("hubToVitualNetworkTrafficAllowed")]
        public Output<bool?> HubToVitualNetworkTrafficAllowed { get; private set; } = null!;

        /// <summary>
        /// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created.
        /// </summary>
        [Output("internetSecurityEnabled")]
        public Output<bool?> InternetSecurityEnabled { get; private set; } = null!;

        /// <summary>
        /// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("remoteVirtualNetworkId")]
        public Output<string> RemoteVirtualNetworkId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualHubId")]
        public Output<string> VirtualHubId { get; private set; } = null!;

        /// <summary>
        /// Is Remote Virtual Network traffic allowed to transit the Hub's Virtual Network Gateway's? Changing this forces a new resource to be created.
        /// </summary>
        [Output("vitualNetworkToHubGatewaysTrafficAllowed")]
        public Output<bool?> VitualNetworkToHubGatewaysTrafficAllowed { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHubConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHubConnection(string name, VirtualHubConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure:network/virtualHubConnection:VirtualHubConnection", name, args ?? new VirtualHubConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualHubConnection(string name, Input<string> id, VirtualHubConnectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/virtualHubConnection:VirtualHubConnection", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualHubConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHubConnection Get(string name, Input<string> id, VirtualHubConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualHubConnection(name, id, state, options);
        }
    }

    public sealed class VirtualHubConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the Virtual Hub traffic allowed to transit via the Remote Virtual Network? Changing this forces a new resource to be created.
        /// </summary>
        [Input("hubToVitualNetworkTrafficAllowed")]
        public Input<bool>? HubToVitualNetworkTrafficAllowed { get; set; }

        /// <summary>
        /// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created.
        /// </summary>
        [Input("internetSecurityEnabled")]
        public Input<bool>? InternetSecurityEnabled { get; set; }

        /// <summary>
        /// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("remoteVirtualNetworkId", required: true)]
        public Input<string> RemoteVirtualNetworkId { get; set; } = null!;

        /// <summary>
        /// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualHubId", required: true)]
        public Input<string> VirtualHubId { get; set; } = null!;

        /// <summary>
        /// Is Remote Virtual Network traffic allowed to transit the Hub's Virtual Network Gateway's? Changing this forces a new resource to be created.
        /// </summary>
        [Input("vitualNetworkToHubGatewaysTrafficAllowed")]
        public Input<bool>? VitualNetworkToHubGatewaysTrafficAllowed { get; set; }

        public VirtualHubConnectionArgs()
        {
        }
    }

    public sealed class VirtualHubConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the Virtual Hub traffic allowed to transit via the Remote Virtual Network? Changing this forces a new resource to be created.
        /// </summary>
        [Input("hubToVitualNetworkTrafficAllowed")]
        public Input<bool>? HubToVitualNetworkTrafficAllowed { get; set; }

        /// <summary>
        /// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created.
        /// </summary>
        [Input("internetSecurityEnabled")]
        public Input<bool>? InternetSecurityEnabled { get; set; }

        /// <summary>
        /// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("remoteVirtualNetworkId")]
        public Input<string>? RemoteVirtualNetworkId { get; set; }

        /// <summary>
        /// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualHubId")]
        public Input<string>? VirtualHubId { get; set; }

        /// <summary>
        /// Is Remote Virtual Network traffic allowed to transit the Hub's Virtual Network Gateway's? Changing this forces a new resource to be created.
        /// </summary>
        [Input("vitualNetworkToHubGatewaysTrafficAllowed")]
        public Input<bool>? VitualNetworkToHubGatewaysTrafficAllowed { get; set; }

        public VirtualHubConnectionState()
        {
        }
    }
}
