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
    public sealed class ScaleSetOsProfileSecret
    {
        /// <summary>
        /// Specifies the key vault to use.
        /// </summary>
        public readonly string SourceVaultId;
        /// <summary>
        /// A collection of Vault Certificates as documented below
        /// </summary>
        public readonly ImmutableArray<Outputs.ScaleSetOsProfileSecretVaultCertificate> VaultCertificates;

        [OutputConstructor]
        private ScaleSetOsProfileSecret(
            string sourceVaultId,

            ImmutableArray<Outputs.ScaleSetOsProfileSecretVaultCertificate> vaultCertificates)
        {
            SourceVaultId = sourceVaultId;
            VaultCertificates = vaultCertificates;
        }
    }
}
