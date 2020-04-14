// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class ScaleSetOsProfileSecretVaultCertificateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the certificate store on the Virtual Machine where the certificate should be added to.
        /// </summary>
        [Input("certificateStore")]
        public Input<string>? CertificateStore { get; set; }

        /// <summary>
        /// It is the Base64 encoding of a JSON Object that which is encoded in UTF-8 of which the contents need to be `data`, `dataType` and `password`.
        /// </summary>
        [Input("certificateUrl", required: true)]
        public Input<string> CertificateUrl { get; set; } = null!;

        public ScaleSetOsProfileSecretVaultCertificateGetArgs()
        {
        }
    }
}
