// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault.Outputs
{

    [OutputType]
    public sealed class CertificateCertificatePolicySecretProperties
    {
        /// <summary>
        /// The Content-Type of the Certificate, such as `application/x-pkcs12` for a PFX or `application/x-pem-file` for a PEM. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string ContentType;

        [OutputConstructor]
        private CertificateCertificatePolicySecretProperties(string contentType)
        {
            ContentType = contentType;
        }
    }
}
