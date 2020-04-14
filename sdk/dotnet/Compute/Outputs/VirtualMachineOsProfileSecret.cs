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
    public sealed class VirtualMachineOsProfileSecret
    {
        /// <summary>
        /// Specifies the ID of the Key Vault to use.
        /// </summary>
        public readonly string SourceVaultId;
        /// <summary>
        /// One or more `vault_certificates` blocks.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineOsProfileSecretVaultCertificate> VaultCertificates;

        [OutputConstructor]
        private VirtualMachineOsProfileSecret(
            string sourceVaultId,

            ImmutableArray<Outputs.VirtualMachineOsProfileSecretVaultCertificate> vaultCertificates)
        {
            SourceVaultId = sourceVaultId;
            VaultCertificates = vaultCertificates;
        }
    }
}
