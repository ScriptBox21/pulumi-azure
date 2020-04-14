// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class RServerClusterRolesEdgeNodeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Password associated with the local administrator for the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;

        /// <summary>
        /// A list of SSH Keys which should be used for the local administrator on the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        /// <summary>
        /// The ID of the Subnet within the Virtual Network where the Edge Node should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The Username of the local administrator for the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// The ID of the Virtual Network where the Edge Node should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        /// <summary>
        /// The Size of the Virtual Machine which should be used as the Edge Node. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesEdgeNodeGetArgs()
        {
        }
    }
}
