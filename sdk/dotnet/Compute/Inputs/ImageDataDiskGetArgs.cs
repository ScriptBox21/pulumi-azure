// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class ImageDataDiskGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the URI in Azure storage of the blob that you want to use to create the image.
        /// </summary>
        [Input("blobUri")]
        public Input<string>? BlobUri { get; set; }

        /// <summary>
        /// Specifies the caching mode as `ReadWrite`, `ReadOnly`, or `None`. The default is `None`.
        /// </summary>
        [Input("caching")]
        public Input<string>? Caching { get; set; }

        /// <summary>
        /// Specifies the logical unit number of the data disk.
        /// </summary>
        [Input("lun")]
        public Input<int>? Lun { get; set; }

        /// <summary>
        /// Specifies the ID of the managed disk resource that you want to use to create the image.
        /// </summary>
        [Input("managedDiskId")]
        public Input<string>? ManagedDiskId { get; set; }

        /// <summary>
        /// Specifies the size of the image to be created. The target size can't be smaller than the source size.
        /// </summary>
        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        public ImageDataDiskGetArgs()
        {
        }
    }
}
