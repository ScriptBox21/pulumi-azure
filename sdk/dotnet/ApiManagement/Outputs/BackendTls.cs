// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Outputs
{

    [OutputType]
    public sealed class BackendTls
    {
        /// <summary>
        /// Flag indicating whether SSL certificate chain validation should be done when using self-signed certificates for the backend host.
        /// </summary>
        public readonly bool? ValidateCertificateChain;
        /// <summary>
        /// Flag indicating whether SSL certificate name validation should be done when using self-signed certificates for the backend host.
        /// </summary>
        public readonly bool? ValidateCertificateName;

        [OutputConstructor]
        private BackendTls(
            bool? validateCertificateChain,

            bool? validateCertificateName)
        {
            ValidateCertificateChain = validateCertificateChain;
            ValidateCertificateName = validateCertificateName;
        }
    }
}
