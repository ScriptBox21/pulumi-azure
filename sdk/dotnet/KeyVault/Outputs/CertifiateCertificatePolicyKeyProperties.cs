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
    public sealed class CertifiateCertificatePolicyKeyProperties
    {
        /// <summary>
        /// Is this Certificate Exportable? Changing this forces a new resource to be created.
        /// </summary>
        public readonly bool Exportable;
        /// <summary>
        /// The size of the Key used in the Certificate. Possible values include `2048` and `4096`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int KeySize;
        /// <summary>
        /// Specifies the Type of Key, such as `RSA`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string KeyType;
        /// <summary>
        /// Is the key reusable? Changing this forces a new resource to be created.
        /// </summary>
        public readonly bool ReuseKey;

        [OutputConstructor]
        private CertifiateCertificatePolicyKeyProperties(
            bool exportable,

            int keySize,

            string keyType,

            bool reuseKey)
        {
            Exportable = exportable;
            KeySize = keySize;
            KeyType = keyType;
            ReuseKey = reuseKey;
        }
    }
}
