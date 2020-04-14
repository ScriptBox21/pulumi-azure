// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class ManagedDiskEncryptionSettingsDiskEncryptionKeyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL to the Key Vault Secret used as the Disk Encryption Key. This can be found as `id` on the `azure.keyvault.Secret` resource.
        /// </summary>
        [Input("secretUrl", required: true)]
        public Input<string> SecretUrl { get; set; } = null!;

        /// <summary>
        /// The ID of the source Key Vault.
        /// </summary>
        [Input("sourceVaultId", required: true)]
        public Input<string> SourceVaultId { get; set; } = null!;

        public ManagedDiskEncryptionSettingsDiskEncryptionKeyGetArgs()
        {
        }
    }
}
