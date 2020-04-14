// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class LinuxVirtualMachineScaleSetSecretArgs : Pulumi.ResourceArgs
    {
        [Input("certificates", required: true)]
        private InputList<Inputs.LinuxVirtualMachineScaleSetSecretCertificateArgs>? _certificates;

        /// <summary>
        /// One or more `certificate` blocks as defined above.
        /// </summary>
        public InputList<Inputs.LinuxVirtualMachineScaleSetSecretCertificateArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.LinuxVirtualMachineScaleSetSecretCertificateArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// The ID of the Key Vault from which all Secrets should be sourced.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        public LinuxVirtualMachineScaleSetSecretArgs()
        {
        }
    }
}
