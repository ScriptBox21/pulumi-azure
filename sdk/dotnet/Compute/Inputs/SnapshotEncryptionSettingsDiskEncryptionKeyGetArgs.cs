// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class SnapshotEncryptionSettingsDiskEncryptionKeyGetArgs : Pulumi.ResourceArgs
    {
        [Input("secretUrl", required: true)]
        public Input<string> SecretUrl { get; set; } = null!;

        [Input("sourceVaultId", required: true)]
        public Input<string> SourceVaultId { get; set; } = null!;

        public SnapshotEncryptionSettingsDiskEncryptionKeyGetArgs()
        {
        }
    }
}
