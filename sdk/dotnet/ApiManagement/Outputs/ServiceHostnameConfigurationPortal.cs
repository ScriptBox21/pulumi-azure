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
    public sealed class ServiceHostnameConfigurationPortal
    {
        /// <summary>
        /// One or more (up to 10) `certificate` blocks as defined below.
        /// </summary>
        public readonly string? Certificate;
        /// <summary>
        /// The password for the certificate.
        /// </summary>
        public readonly string? CertificatePassword;
        /// <summary>
        /// The Hostname to use for the Management API.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
        /// </summary>
        public readonly string? KeyVaultId;
        /// <summary>
        /// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.
        /// </summary>
        public readonly bool? NegotiateClientCertificate;

        [OutputConstructor]
        private ServiceHostnameConfigurationPortal(
            string? certificate,

            string? certificatePassword,

            string hostName,

            string? keyVaultId,

            bool? negotiateClientCertificate)
        {
            Certificate = certificate;
            CertificatePassword = certificatePassword;
            HostName = hostName;
            KeyVaultId = keyVaultId;
            NegotiateClientCertificate = negotiateClientCertificate;
        }
    }
}
