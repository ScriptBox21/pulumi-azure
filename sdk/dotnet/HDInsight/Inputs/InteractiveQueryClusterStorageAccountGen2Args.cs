// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class InteractiveQueryClusterStorageAccountGen2Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Gen2 Filesystem. Changing this forces a new resource to be created.
        /// </summary>
        [Input("filesystemId", required: true)]
        public Input<string> FilesystemId { get; set; } = null!;

        /// <summary>
        /// Is this the Default Storage Account for the HDInsight Hadoop Cluster? Changing this forces a new resource to be created.
        /// </summary>
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        /// <summary>
        /// The ID of Managed Identity to use for accessing the Gen2 filesystem. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedIdentityResourceId", required: true)]
        public Input<string> ManagedIdentityResourceId { get; set; } = null!;

        /// <summary>
        /// The ID of the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageResourceId", required: true)]
        public Input<string> StorageResourceId { get; set; } = null!;

        public InteractiveQueryClusterStorageAccountGen2Args()
        {
        }
    }
}
