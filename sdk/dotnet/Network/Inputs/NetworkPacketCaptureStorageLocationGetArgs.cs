// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class NetworkPacketCaptureStorageLocationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A valid local path on the targeting VM. Must include the name of the capture file (*.cap). For linux virtual machine it must start with `/var/captures`.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// The ID of the storage account to save the packet capture session
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        /// <summary>
        /// The URI of the storage path to save the packet capture.
        /// </summary>
        [Input("storagePath")]
        public Input<string>? StoragePath { get; set; }

        public NetworkPacketCaptureStorageLocationGetArgs()
        {
        }
    }
}
