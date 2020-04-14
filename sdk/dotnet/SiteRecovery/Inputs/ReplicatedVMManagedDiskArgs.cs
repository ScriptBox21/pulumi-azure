// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SiteRecovery.Inputs
{

    public sealed class ReplicatedVMManagedDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of disk that should be replicated.
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        /// <summary>
        /// Storage account that should be used for caching.
        /// </summary>
        [Input("stagingStorageAccountId", required: true)]
        public Input<string> StagingStorageAccountId { get; set; } = null!;

        /// <summary>
        /// What type should the disk be when a failover is done.
        /// </summary>
        [Input("targetDiskType", required: true)]
        public Input<string> TargetDiskType { get; set; } = null!;

        /// <summary>
        /// What type should the disk be that holds the replication data.
        /// </summary>
        [Input("targetReplicaDiskType", required: true)]
        public Input<string> TargetReplicaDiskType { get; set; } = null!;

        /// <summary>
        /// Resource group disk should belong to when a failover is done.
        /// </summary>
        [Input("targetResourceGroupId", required: true)]
        public Input<string> TargetResourceGroupId { get; set; } = null!;

        public ReplicatedVMManagedDiskArgs()
        {
        }
    }
}
