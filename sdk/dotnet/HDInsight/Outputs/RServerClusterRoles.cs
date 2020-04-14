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
    public sealed class RServerClusterRoles
    {
        /// <summary>
        /// A `edge_node` block as defined above.
        /// </summary>
        public readonly Outputs.RServerClusterRolesEdgeNode EdgeNode;
        /// <summary>
        /// A `head_node` block as defined above.
        /// </summary>
        public readonly Outputs.RServerClusterRolesHeadNode HeadNode;
        /// <summary>
        /// A `worker_node` block as defined below.
        /// </summary>
        public readonly Outputs.RServerClusterRolesWorkerNode WorkerNode;
        /// <summary>
        /// A `zookeeper_node` block as defined below.
        /// </summary>
        public readonly Outputs.RServerClusterRolesZookeeperNode ZookeeperNode;

        [OutputConstructor]
        private RServerClusterRoles(
            Outputs.RServerClusterRolesEdgeNode edgeNode,

            Outputs.RServerClusterRolesHeadNode headNode,

            Outputs.RServerClusterRolesWorkerNode workerNode,

            Outputs.RServerClusterRolesZookeeperNode zookeeperNode)
        {
            EdgeNode = edgeNode;
            HeadNode = headNode;
            WorkerNode = workerNode;
            ZookeeperNode = zookeeperNode;
        }
    }
}
