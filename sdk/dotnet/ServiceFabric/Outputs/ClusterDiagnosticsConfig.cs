// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric.Outputs
{

    [OutputType]
    public sealed class ClusterDiagnosticsConfig
    {
        /// <summary>
        /// The Blob Endpoint of the Storage Account.
        /// </summary>
        public readonly string BlobEndpoint;
        /// <summary>
        /// The protected diagnostics storage key name, such as `StorageAccountKey1`.
        /// </summary>
        public readonly string ProtectedAccountKeyName;
        /// <summary>
        /// The Queue Endpoint of the Storage Account.
        /// </summary>
        public readonly string QueueEndpoint;
        /// <summary>
        /// The name of the Storage Account where the Diagnostics should be sent to.
        /// </summary>
        public readonly string StorageAccountName;
        /// <summary>
        /// The Table Endpoint of the Storage Account.
        /// </summary>
        public readonly string TableEndpoint;

        [OutputConstructor]
        private ClusterDiagnosticsConfig(
            string blobEndpoint,

            string protectedAccountKeyName,

            string queueEndpoint,

            string storageAccountName,

            string tableEndpoint)
        {
            BlobEndpoint = blobEndpoint;
            ProtectedAccountKeyName = protectedAccountKeyName;
            QueueEndpoint = queueEndpoint;
            StorageAccountName = storageAccountName;
            TableEndpoint = tableEndpoint;
        }
    }
}
