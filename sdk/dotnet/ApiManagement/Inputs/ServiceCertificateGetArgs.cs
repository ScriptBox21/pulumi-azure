// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Inputs
{

    public sealed class ServiceCertificateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password for the certificate.
        /// </summary>
        [Input("certificatePassword", required: true)]
        public Input<string> CertificatePassword { get; set; } = null!;

        /// <summary>
        /// The Base64 Encoded PFX Certificate.
        /// </summary>
        [Input("encodedCertificate", required: true)]
        public Input<string> EncodedCertificate { get; set; } = null!;

        /// <summary>
        /// The name of the Certificate Store where this certificate should be stored. Possible values are `CertificateAuthority` and `Root`.
        /// </summary>
        [Input("storeName", required: true)]
        public Input<string> StoreName { get; set; } = null!;

        public ServiceCertificateGetArgs()
        {
        }
    }
}
