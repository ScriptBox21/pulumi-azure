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
    public sealed class ServiceHostnameConfigurationProxy
    {
        /// <summary>
        /// The Base64 Encoded Certificate.
        /// </summary>
        public readonly string? Certificate;
        /// <summary>
        /// The password associated with the certificate provided above.
        /// </summary>
        public readonly string? CertificatePassword;
        /// <summary>
        /// Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to `false`.
        /// </summary>
        public readonly bool? DefaultSslBinding;
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
        private ServiceHostnameConfigurationProxy(
            string? certificate,

            string? certificatePassword,

            bool? defaultSslBinding,

            string hostName,

            string? keyVaultId,

            bool? negotiateClientCertificate)
        {
            Certificate = certificate;
            CertificatePassword = certificatePassword;
            DefaultSslBinding = defaultSslBinding;
            HostName = hostName;
            KeyVaultId = keyVaultId;
            NegotiateClientCertificate = negotiateClientCertificate;
        }
    }
}
