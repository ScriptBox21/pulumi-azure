// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class KafkaClusterRolesWorkerNode
    {
        /// <summary>
        /// The minimum number of instances which should be run for the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? MinInstanceCount;
        /// <summary>
        /// The number of Data Disks which should be assigned to each Worker Node, which can be between 1 and 8. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int NumberOfDisksPerNode;
        /// <summary>
        /// The Password associated with the local administrator for the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// A list of SSH Keys which should be used for the local administrator on the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> SshKeys;
        /// <summary>
        /// The ID of the Subnet within the Virtual Network where the Worker Nodes should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The number of instances which should be run for the Worker Nodes.
        /// </summary>
        public readonly int TargetInstanceCount;
        /// <summary>
        /// The Username of the local administrator for the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// The ID of the Virtual Network where the Worker Nodes should be provisioned within. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? VirtualNetworkId;
        /// <summary>
        /// The Size of the Virtual Machine which should be used as the Worker Nodes. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string VmSize;

        [OutputConstructor]
        private KafkaClusterRolesWorkerNode(
            int? minInstanceCount,

            int numberOfDisksPerNode,

            string? password,

            ImmutableArray<string> sshKeys,

            string? subnetId,

            int targetInstanceCount,

            string username,

            string? virtualNetworkId,

            string vmSize)
        {
            MinInstanceCount = minInstanceCount;
            NumberOfDisksPerNode = numberOfDisksPerNode;
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            TargetInstanceCount = targetInstanceCount;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }
}
