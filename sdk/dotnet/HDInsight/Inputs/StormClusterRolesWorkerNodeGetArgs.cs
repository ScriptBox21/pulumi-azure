// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class StormClusterRolesWorkerNodeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The minimum number of instances which should be run for the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        /// <summary>
        /// The Password associated with the local administrator for the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;

        /// <summary>
        /// A list of SSH Keys which should be used for the local administrator on the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        /// <summary>
        /// The ID of the Subnet within the Virtual Network where the Worker Nodes should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The number of instances which should be run for the Worker Nodes.
        /// </summary>
        [Input("targetInstanceCount", required: true)]
        public Input<int> TargetInstanceCount { get; set; } = null!;

        /// <summary>
        /// The Username of the local administrator for the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// The ID of the Virtual Network where the Worker Nodes should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        /// <summary>
        /// The Size of the Virtual Machine which should be used as the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public StormClusterRolesWorkerNodeGetArgs()
        {
        }
    }
}
