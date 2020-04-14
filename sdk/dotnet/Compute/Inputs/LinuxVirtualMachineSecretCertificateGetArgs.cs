// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class LinuxVirtualMachineSecretCertificateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Secret URL of a Key Vault Certificate.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public LinuxVirtualMachineSecretCertificateGetArgs()
        {
        }
    }
}
