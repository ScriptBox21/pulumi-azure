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
    public sealed class InteractiveQueryClusterRoles
    {
        /// <summary>
        /// A `head_node` block as defined above.
        /// </summary>
        public readonly Outputs.InteractiveQueryClusterRolesHeadNode HeadNode;
        /// <summary>
        /// A `worker_node` block as defined below.
        /// </summary>
        public readonly Outputs.InteractiveQueryClusterRolesWorkerNode WorkerNode;
        /// <summary>
        /// A `zookeeper_node` block as defined below.
        /// </summary>
        public readonly Outputs.InteractiveQueryClusterRolesZookeeperNode ZookeeperNode;

        [OutputConstructor]
        private InteractiveQueryClusterRoles(
            Outputs.InteractiveQueryClusterRolesHeadNode headNode,

            Outputs.InteractiveQueryClusterRolesWorkerNode workerNode,

            Outputs.InteractiveQueryClusterRolesZookeeperNode zookeeperNode)
        {
            HeadNode = headNode;
            WorkerNode = workerNode;
            ZookeeperNode = zookeeperNode;
        }
    }
}
