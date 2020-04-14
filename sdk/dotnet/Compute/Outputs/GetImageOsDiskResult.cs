// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class GetImageOsDiskResult
    {
        /// <summary>
        /// the URI in Azure storage of the blob used to create the image.
        /// </summary>
        public readonly string BlobUri;
        /// <summary>
        /// the caching mode for the Data Disk, such as `ReadWrite`, `ReadOnly`, or `None`.
        /// </summary>
        public readonly string Caching;
        /// <summary>
        /// the ID of the Managed Disk used as the Data Disk Image.
        /// </summary>
        public readonly string ManagedDiskId;
        /// <summary>
        /// the State of the OS used in the Image, such as `Generalized`.
        /// </summary>
        public readonly string OsState;
        /// <summary>
        /// the type of Operating System used on the OS Disk. such as `Linux` or `Windows`.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// the size of this Data Disk in GB.
        /// </summary>
        public readonly int SizeGb;

        [OutputConstructor]
        private GetImageOsDiskResult(
            string blobUri,

            string caching,

            string managedDiskId,

            string osState,

            string osType,

            int sizeGb)
        {
            BlobUri = blobUri;
            Caching = caching;
            ManagedDiskId = managedDiskId;
            OsState = osState;
            OsType = osType;
            SizeGb = sizeGb;
        }
    }
}
