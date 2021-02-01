// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault
{
    public static class GetCertificate
    {
        /// <summary>
        /// Use this data source to access information about an existing Key Vault Certificate.
        /// 
        /// &gt; **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
        /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleKeyVault = Output.Create(Azure.KeyVault.GetKeyVault.InvokeAsync(new Azure.KeyVault.GetKeyVaultArgs
        ///         {
        ///             Name = "examplekv",
        ///             ResourceGroupName = "some-resource-group",
        ///         }));
        ///         var exampleCertificate = exampleKeyVault.Apply(exampleKeyVault =&gt; Output.Create(Azure.KeyVault.GetCertificate.InvokeAsync(new Azure.KeyVault.GetCertificateArgs
        ///         {
        ///             Name = "secret-sauce",
        ///             KeyVaultId = exampleKeyVault.Id,
        ///         })));
        ///         this.CertificateThumbprint = exampleCertificate.Apply(exampleCertificate =&gt; exampleCertificate.Thumbprint);
        ///     }
        /// 
        ///     [Output("certificateThumbprint")]
        ///     public Output&lt;string&gt; CertificateThumbprint { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("azure:keyvault/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithVersion());
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the ID of the Key Vault instance where the Secret resides, available on the `azure.keyvault.KeyVault` Data Source / Resource.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public string KeyVaultId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Key Vault Certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the version of the certificate to look up.  (Defaults to latest)
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        public GetCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// The raw Key Vault Certificate data represented as a hexadecimal string.
        /// </summary>
        public readonly string CertificateData;
        /// <summary>
        /// The raw Key Vault Certificate data represented as a base64 string.
        /// </summary>
        public readonly string CertificateDataBase64;
        /// <summary>
        /// A `certificate_policy` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCertificateCertificatePolicyResult> CertificatePolicies;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KeyVaultId;
        /// <summary>
        /// The name of the Certificate Issuer.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the associated Key Vault Secret.
        /// </summary>
        public readonly string SecretId;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// The current version of the Key Vault Certificate.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetCertificateResult(
            string certificateData,

            string certificateDataBase64,

            ImmutableArray<Outputs.GetCertificateCertificatePolicyResult> certificatePolicies,

            string id,

            string keyVaultId,

            string name,

            string secretId,

            ImmutableDictionary<string, string> tags,

            string thumbprint,

            string version)
        {
            CertificateData = certificateData;
            CertificateDataBase64 = certificateDataBase64;
            CertificatePolicies = certificatePolicies;
            Id = id;
            KeyVaultId = keyVaultId;
            Name = name;
            SecretId = secretId;
            Tags = tags;
            Thumbprint = thumbprint;
            Version = version;
        }
    }
}
