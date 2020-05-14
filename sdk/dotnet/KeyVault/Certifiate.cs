// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault
{
    /// <summary>
    /// Manages a Key Vault Certificate.
    /// </summary>
    [Obsolete(@"azure.keyvault.Certifiate has been deprecated in favor of azure.keyvault.Certificate")]
    public partial class Certifiate : Pulumi.CustomResource
    {
        /// <summary>
        /// A `certificate` block as defined below, used to Import an existing certificate.
        /// </summary>
        [Output("certificate")]
        public Output<Outputs.CertifiateCertificate?> KeyVaultCertificate { get; private set; } = null!;

        /// <summary>
        /// The raw Key Vault Certificate data represented as a hexadecimal string.
        /// </summary>
        [Output("certificateData")]
        public Output<string> CertificateData { get; private set; } = null!;

        /// <summary>
        /// A `certificate_policy` block as defined below.
        /// </summary>
        [Output("certificatePolicy")]
        public Output<Outputs.CertifiateCertificatePolicy> CertificatePolicy { get; private set; } = null!;

        /// <summary>
        /// The ID of the Key Vault where the Certificate should be created.
        /// </summary>
        [Output("keyVaultId")]
        public Output<string> KeyVaultId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated Key Vault Secret.
        /// </summary>
        [Output("secretId")]
        public Output<string> SecretId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        /// </summary>
        [Output("thumbprint")]
        public Output<string> Thumbprint { get; private set; } = null!;

        /// <summary>
        /// The current version of the Key Vault Certificate.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Certifiate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certifiate(string name, CertifiateArgs args, CustomResourceOptions? options = null)
            : base("azure:keyvault/certifiate:Certifiate", name, args ?? new CertifiateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certifiate(string name, Input<string> id, CertifiateState? state = null, CustomResourceOptions? options = null)
            : base("azure:keyvault/certifiate:Certifiate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certifiate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certifiate Get(string name, Input<string> id, CertifiateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certifiate(name, id, state, options);
        }
    }

    public sealed class CertifiateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `certificate` block as defined below, used to Import an existing certificate.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.CertifiateCertificateArgs>? KeyVaultCertificate { get; set; }

        /// <summary>
        /// A `certificate_policy` block as defined below.
        /// </summary>
        [Input("certificatePolicy", required: true)]
        public Input<Inputs.CertifiateCertificatePolicyArgs> CertificatePolicy { get; set; } = null!;

        /// <summary>
        /// The ID of the Key Vault where the Certificate should be created.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CertifiateArgs()
        {
        }
    }

    public sealed class CertifiateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `certificate` block as defined below, used to Import an existing certificate.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.CertifiateCertificateGetArgs>? KeyVaultCertificate { get; set; }

        /// <summary>
        /// The raw Key Vault Certificate data represented as a hexadecimal string.
        /// </summary>
        [Input("certificateData")]
        public Input<string>? CertificateData { get; set; }

        /// <summary>
        /// A `certificate_policy` block as defined below.
        /// </summary>
        [Input("certificatePolicy")]
        public Input<Inputs.CertifiateCertificatePolicyGetArgs>? CertificatePolicy { get; set; }

        /// <summary>
        /// The ID of the Key Vault where the Certificate should be created.
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the associated Key Vault Secret.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        /// <summary>
        /// The current version of the Key Vault Certificate.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public CertifiateState()
        {
        }
    }
}
