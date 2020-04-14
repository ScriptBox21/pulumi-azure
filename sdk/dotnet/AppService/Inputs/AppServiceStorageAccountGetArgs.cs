// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class AppServiceStorageAccountGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key for the storage account.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// The name of the storage account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The path to mount the storage within the site's runtime environment.
        /// </summary>
        [Input("mountPath")]
        public Input<string>? MountPath { get; set; }

        /// <summary>
        /// The name of the storage account identifier.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the file share (container name, for Blob storage).
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        /// <summary>
        /// The type of storage. Possible values are `AzureBlob` and `AzureFiles`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AppServiceStorageAccountGetArgs()
        {
        }
    }
}
