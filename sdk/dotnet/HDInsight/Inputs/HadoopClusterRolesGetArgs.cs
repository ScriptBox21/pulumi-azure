// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class HadoopClusterRolesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `edge_node` block as defined below.
        /// </summary>
        [Input("edgeNode")]
        public Input<Inputs.HadoopClusterRolesEdgeNodeGetArgs>? EdgeNode { get; set; }

        /// <summary>
        /// A `head_node` block as defined above.
        /// </summary>
        [Input("headNode", required: true)]
        public Input<Inputs.HadoopClusterRolesHeadNodeGetArgs> HeadNode { get; set; } = null!;

        /// <summary>
        /// A `worker_node` block as defined below.
        /// </summary>
        [Input("workerNode", required: true)]
        public Input<Inputs.HadoopClusterRolesWorkerNodeGetArgs> WorkerNode { get; set; } = null!;

        /// <summary>
        /// A `zookeeper_node` block as defined below.
        /// </summary>
        [Input("zookeeperNode", required: true)]
        public Input<Inputs.HadoopClusterRolesZookeeperNodeGetArgs> ZookeeperNode { get; set; } = null!;

        public HadoopClusterRolesGetArgs()
        {
        }
    }
}
