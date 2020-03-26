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
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/key_vault_certificate.html.markdown.
    /// </summary>
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
            : base("azure:keyvault/certifiate:Certifiate", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    namespace Inputs
    {

    public sealed class CertifiateCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base64-encoded certificate contents. Changing this forces a new resource to be created.
        /// </summary>
        [Input("contents", required: true)]
        public Input<string> Contents { get; set; } = null!;

        /// <summary>
        /// The password associated with the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        public CertifiateCertificateArgs()
        {
        }
    }

    public sealed class CertifiateCertificateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base64-encoded certificate contents. Changing this forces a new resource to be created.
        /// </summary>
        [Input("contents", required: true)]
        public Input<string> Contents { get; set; } = null!;

        /// <summary>
        /// The password associated with the certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        public CertifiateCertificateGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `issuer_parameters` block as defined below.
        /// </summary>
        [Input("issuerParameters", required: true)]
        public Input<CertifiateCertificatePolicyIssuerParametersArgs> IssuerParameters { get; set; } = null!;

        /// <summary>
        /// A `key_properties` block as defined below.
        /// </summary>
        [Input("keyProperties", required: true)]
        public Input<CertifiateCertificatePolicyKeyPropertiesArgs> KeyProperties { get; set; } = null!;

        [Input("lifetimeActions")]
        private InputList<CertifiateCertificatePolicyLifetimeActionsArgs>? _lifetimeActions;

        /// <summary>
        /// A `lifetime_action` block as defined below.
        /// </summary>
        public InputList<CertifiateCertificatePolicyLifetimeActionsArgs> LifetimeActions
        {
            get => _lifetimeActions ?? (_lifetimeActions = new InputList<CertifiateCertificatePolicyLifetimeActionsArgs>());
            set => _lifetimeActions = value;
        }

        /// <summary>
        /// A `secret_properties` block as defined below.
        /// </summary>
        [Input("secretProperties", required: true)]
        public Input<CertifiateCertificatePolicySecretPropertiesArgs> SecretProperties { get; set; } = null!;

        /// <summary>
        /// A `x509_certificate_properties` block as defined below.
        /// </summary>
        [Input("x509CertificateProperties")]
        public Input<CertifiateCertificatePolicyX509CertificatePropertiesArgs>? X509CertificateProperties { get; set; }

        public CertifiateCertificatePolicyArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `issuer_parameters` block as defined below.
        /// </summary>
        [Input("issuerParameters", required: true)]
        public Input<CertifiateCertificatePolicyIssuerParametersGetArgs> IssuerParameters { get; set; } = null!;

        /// <summary>
        /// A `key_properties` block as defined below.
        /// </summary>
        [Input("keyProperties", required: true)]
        public Input<CertifiateCertificatePolicyKeyPropertiesGetArgs> KeyProperties { get; set; } = null!;

        [Input("lifetimeActions")]
        private InputList<CertifiateCertificatePolicyLifetimeActionsGetArgs>? _lifetimeActions;

        /// <summary>
        /// A `lifetime_action` block as defined below.
        /// </summary>
        public InputList<CertifiateCertificatePolicyLifetimeActionsGetArgs> LifetimeActions
        {
            get => _lifetimeActions ?? (_lifetimeActions = new InputList<CertifiateCertificatePolicyLifetimeActionsGetArgs>());
            set => _lifetimeActions = value;
        }

        /// <summary>
        /// A `secret_properties` block as defined below.
        /// </summary>
        [Input("secretProperties", required: true)]
        public Input<CertifiateCertificatePolicySecretPropertiesGetArgs> SecretProperties { get; set; } = null!;

        /// <summary>
        /// A `x509_certificate_properties` block as defined below.
        /// </summary>
        [Input("x509CertificateProperties")]
        public Input<CertifiateCertificatePolicyX509CertificatePropertiesGetArgs>? X509CertificateProperties { get; set; }

        public CertifiateCertificatePolicyGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyIssuerParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public CertifiateCertificatePolicyIssuerParametersArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyIssuerParametersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public CertifiateCertificatePolicyIssuerParametersGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyKeyPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is this Certificate Exportable? Changing this forces a new resource to be created.
        /// </summary>
        [Input("exportable", required: true)]
        public Input<bool> Exportable { get; set; } = null!;

        /// <summary>
        /// The size of the Key used in the Certificate. Possible values include `2048` and `4096`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keySize", required: true)]
        public Input<int> KeySize { get; set; } = null!;

        /// <summary>
        /// Specifies the Type of Key, such as `RSA`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        /// <summary>
        /// Is the key reusable? Changing this forces a new resource to be created.
        /// </summary>
        [Input("reuseKey", required: true)]
        public Input<bool> ReuseKey { get; set; } = null!;

        public CertifiateCertificatePolicyKeyPropertiesArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyKeyPropertiesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is this Certificate Exportable? Changing this forces a new resource to be created.
        /// </summary>
        [Input("exportable", required: true)]
        public Input<bool> Exportable { get; set; } = null!;

        /// <summary>
        /// The size of the Key used in the Certificate. Possible values include `2048` and `4096`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keySize", required: true)]
        public Input<int> KeySize { get; set; } = null!;

        /// <summary>
        /// Specifies the Type of Key, such as `RSA`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        /// <summary>
        /// Is the key reusable? Changing this forces a new resource to be created.
        /// </summary>
        [Input("reuseKey", required: true)]
        public Input<bool> ReuseKey { get; set; } = null!;

        public CertifiateCertificatePolicyKeyPropertiesGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyLifetimeActionsActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Type of action to be performed when the lifetime trigger is triggerec. Possible values include `AutoRenew` and `EmailContacts`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        public CertifiateCertificatePolicyLifetimeActionsActionArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyLifetimeActionsActionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Type of action to be performed when the lifetime trigger is triggerec. Possible values include `AutoRenew` and `EmailContacts`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        public CertifiateCertificatePolicyLifetimeActionsActionGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyLifetimeActionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `action` block as defined below.
        /// </summary>
        [Input("action", required: true)]
        public Input<CertifiateCertificatePolicyLifetimeActionsActionArgs> Action { get; set; } = null!;

        /// <summary>
        /// A `trigger` block as defined below.
        /// </summary>
        [Input("trigger", required: true)]
        public Input<CertifiateCertificatePolicyLifetimeActionsTriggerArgs> Trigger { get; set; } = null!;

        public CertifiateCertificatePolicyLifetimeActionsArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyLifetimeActionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `action` block as defined below.
        /// </summary>
        [Input("action", required: true)]
        public Input<CertifiateCertificatePolicyLifetimeActionsActionGetArgs> Action { get; set; } = null!;

        /// <summary>
        /// A `trigger` block as defined below.
        /// </summary>
        [Input("trigger", required: true)]
        public Input<CertifiateCertificatePolicyLifetimeActionsTriggerGetArgs> Trigger { get; set; } = null!;

        public CertifiateCertificatePolicyLifetimeActionsGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyLifetimeActionsTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days before the Certificate expires that the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with `lifetime_percentage`.
        /// </summary>
        [Input("daysBeforeExpiry")]
        public Input<int>? DaysBeforeExpiry { get; set; }

        /// <summary>
        /// The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with `days_before_expiry`.
        /// </summary>
        [Input("lifetimePercentage")]
        public Input<int>? LifetimePercentage { get; set; }

        public CertifiateCertificatePolicyLifetimeActionsTriggerArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyLifetimeActionsTriggerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days before the Certificate expires that the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with `lifetime_percentage`.
        /// </summary>
        [Input("daysBeforeExpiry")]
        public Input<int>? DaysBeforeExpiry { get; set; }

        /// <summary>
        /// The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with `days_before_expiry`.
        /// </summary>
        [Input("lifetimePercentage")]
        public Input<int>? LifetimePercentage { get; set; }

        public CertifiateCertificatePolicyLifetimeActionsTriggerGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicySecretPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Content-Type of the Certificate, such as `application/x-pkcs12` for a PFX or `application/x-pem-file` for a PEM. Changing this forces a new resource to be created.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        public CertifiateCertificatePolicySecretPropertiesArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicySecretPropertiesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Content-Type of the Certificate, such as `application/x-pkcs12` for a PFX or `application/x-pem-file` for a PEM. Changing this forces a new resource to be created.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        public CertifiateCertificatePolicySecretPropertiesGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyX509CertificatePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("extendedKeyUsages")]
        private InputList<string>? _extendedKeyUsages;

        /// <summary>
        /// A list of Extended/Enhanced Key Usages. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> ExtendedKeyUsages
        {
            get => _extendedKeyUsages ?? (_extendedKeyUsages = new InputList<string>());
            set => _extendedKeyUsages = value;
        }

        [Input("keyUsages", required: true)]
        private InputList<string>? _keyUsages;

        /// <summary>
        /// A list of uses associated with this Key. Possible values include `cRLSign`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `encipherOnly`, `keyAgreement`, `keyCertSign`, `keyEncipherment` and `nonRepudiation` and are case-sensitive. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> KeyUsages
        {
            get => _keyUsages ?? (_keyUsages = new InputList<string>());
            set => _keyUsages = value;
        }

        /// <summary>
        /// The Certificate's Subject. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        /// <summary>
        /// A `subject_alternative_names` block as defined below.
        /// </summary>
        [Input("subjectAlternativeNames")]
        public Input<CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs>? SubjectAlternativeNames { get; set; }

        /// <summary>
        /// The Certificates Validity Period in Months. Changing this forces a new resource to be created.
        /// </summary>
        [Input("validityInMonths", required: true)]
        public Input<int> ValidityInMonths { get; set; } = null!;

        public CertifiateCertificatePolicyX509CertificatePropertiesArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyX509CertificatePropertiesGetArgs : Pulumi.ResourceArgs
    {
        [Input("extendedKeyUsages")]
        private InputList<string>? _extendedKeyUsages;

        /// <summary>
        /// A list of Extended/Enhanced Key Usages. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> ExtendedKeyUsages
        {
            get => _extendedKeyUsages ?? (_extendedKeyUsages = new InputList<string>());
            set => _extendedKeyUsages = value;
        }

        [Input("keyUsages", required: true)]
        private InputList<string>? _keyUsages;

        /// <summary>
        /// A list of uses associated with this Key. Possible values include `cRLSign`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `encipherOnly`, `keyAgreement`, `keyCertSign`, `keyEncipherment` and `nonRepudiation` and are case-sensitive. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> KeyUsages
        {
            get => _keyUsages ?? (_keyUsages = new InputList<string>());
            set => _keyUsages = value;
        }

        /// <summary>
        /// The Certificate's Subject. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        /// <summary>
        /// A `subject_alternative_names` block as defined below.
        /// </summary>
        [Input("subjectAlternativeNames")]
        public Input<CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesGetArgs>? SubjectAlternativeNames { get; set; }

        /// <summary>
        /// The Certificates Validity Period in Months. Changing this forces a new resource to be created.
        /// </summary>
        [Input("validityInMonths", required: true)]
        public Input<int> ValidityInMonths { get; set; } = null!;

        public CertifiateCertificatePolicyX509CertificatePropertiesGetArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs : Pulumi.ResourceArgs
    {
        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        [Input("emails")]
        private InputList<string>? _emails;

        /// <summary>
        /// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Emails
        {
            get => _emails ?? (_emails = new InputList<string>());
            set => _emails = value;
        }

        [Input("upns")]
        private InputList<string>? _upns;

        /// <summary>
        /// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Upns
        {
            get => _upns ?? (_upns = new InputList<string>());
            set => _upns = value;
        }

        public CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs()
        {
        }
    }

    public sealed class CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        [Input("emails")]
        private InputList<string>? _emails;

        /// <summary>
        /// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Emails
        {
            get => _emails ?? (_emails = new InputList<string>());
            set => _emails = value;
        }

        [Input("upns")]
        private InputList<string>? _upns;

        /// <summary>
        /// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Upns
        {
            get => _upns ?? (_upns = new InputList<string>());
            set => _upns = value;
        }

        public CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class CertifiateCertificate
    {
        /// <summary>
        /// The base64-encoded certificate contents. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Contents;
        /// <summary>
        /// The password associated with the certificate. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Password;

        [OutputConstructor]
        private CertifiateCertificate(
            string contents,
            string? password)
        {
            Contents = contents;
            Password = password;
        }
    }

    [OutputType]
    public sealed class CertifiateCertificatePolicy
    {
        /// <summary>
        /// A `issuer_parameters` block as defined below.
        /// </summary>
        public readonly CertifiateCertificatePolicyIssuerParameters IssuerParameters;
        /// <summary>
        /// A `key_properties` block as defined below.
        /// </summary>
        public readonly CertifiateCertificatePolicyKeyProperties KeyProperties;
        /// <summary>
        /// A `lifetime_action` block as defined below.
        /// </summary>
        public readonly ImmutableArray<CertifiateCertificatePolicyLifetimeActions> LifetimeActions;
        /// <summary>
        /// A `secret_properties` block as defined below.
        /// </summary>
        public readonly CertifiateCertificatePolicySecretProperties SecretProperties;
        /// <summary>
        /// A `x509_certificate_properties` block as defined below.
        /// </summary>
        public readonly CertifiateCertificatePolicyX509CertificateProperties X509CertificateProperties;

        [OutputConstructor]
        private CertifiateCertificatePolicy(
            CertifiateCertificatePolicyIssuerParameters issuerParameters,
            CertifiateCertificatePolicyKeyProperties keyProperties,
            ImmutableArray<CertifiateCertificatePolicyLifetimeActions> lifetimeActions,
            CertifiateCertificatePolicySecretProperties secretProperties,
            CertifiateCertificatePolicyX509CertificateProperties x509CertificateProperties)
        {
            IssuerParameters = issuerParameters;
            KeyProperties = keyProperties;
            LifetimeActions = lifetimeActions;
            SecretProperties = secretProperties;
            X509CertificateProperties = x509CertificateProperties;
        }
    }

    [OutputType]
    public sealed class CertifiateCertificatePolicyIssuerParameters
    {
        /// <summary>
        /// The name of the Certificate Issuer. Possible values include `Self` (for self-signed certificate), or `Unknown` (for a certificate issuing authority like `Let's Encrypt` and Azure direct supported ones). Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private CertifiateCertificatePolicyIssuerParameters(string name)
        {
            Name = name;
        }
    }

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

    [OutputType]
    public sealed class CertifiateCertificatePolicyLifetimeActions
    {
        /// <summary>
        /// A `action` block as defined below.
        /// </summary>
        public readonly CertifiateCertificatePolicyLifetimeActionsAction Action;
        /// <summary>
        /// A `trigger` block as defined below.
        /// </summary>
        public readonly CertifiateCertificatePolicyLifetimeActionsTrigger Trigger;

        [OutputConstructor]
        private CertifiateCertificatePolicyLifetimeActions(
            CertifiateCertificatePolicyLifetimeActionsAction action,
            CertifiateCertificatePolicyLifetimeActionsTrigger trigger)
        {
            Action = action;
            Trigger = trigger;
        }
    }

    [OutputType]
    public sealed class CertifiateCertificatePolicyLifetimeActionsAction
    {
        /// <summary>
        /// The Type of action to be performed when the lifetime trigger is triggerec. Possible values include `AutoRenew` and `EmailContacts`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string ActionType;

        [OutputConstructor]
        private CertifiateCertificatePolicyLifetimeActionsAction(string actionType)
        {
            ActionType = actionType;
        }
    }

    [OutputType]
    public sealed class CertifiateCertificatePolicyLifetimeActionsTrigger
    {
        /// <summary>
        /// The number of days before the Certificate expires that the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with `lifetime_percentage`.
        /// </summary>
        public readonly int? DaysBeforeExpiry;
        /// <summary>
        /// The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with `days_before_expiry`.
        /// </summary>
        public readonly int? LifetimePercentage;

        [OutputConstructor]
        private CertifiateCertificatePolicyLifetimeActionsTrigger(
            int? daysBeforeExpiry,
            int? lifetimePercentage)
        {
            DaysBeforeExpiry = daysBeforeExpiry;
            LifetimePercentage = lifetimePercentage;
        }
    }

    [OutputType]
    public sealed class CertifiateCertificatePolicySecretProperties
    {
        /// <summary>
        /// The Content-Type of the Certificate, such as `application/x-pkcs12` for a PFX or `application/x-pem-file` for a PEM. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string ContentType;

        [OutputConstructor]
        private CertifiateCertificatePolicySecretProperties(string contentType)
        {
            ContentType = contentType;
        }
    }

    [OutputType]
    public sealed class CertifiateCertificatePolicyX509CertificateProperties
    {
        /// <summary>
        /// A list of Extended/Enhanced Key Usages. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> ExtendedKeyUsages;
        /// <summary>
        /// A list of uses associated with this Key. Possible values include `cRLSign`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `encipherOnly`, `keyAgreement`, `keyCertSign`, `keyEncipherment` and `nonRepudiation` and are case-sensitive. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> KeyUsages;
        /// <summary>
        /// The Certificate's Subject. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// A `subject_alternative_names` block as defined below.
        /// </summary>
        public readonly CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames SubjectAlternativeNames;
        /// <summary>
        /// The Certificates Validity Period in Months. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int ValidityInMonths;

        [OutputConstructor]
        private CertifiateCertificatePolicyX509CertificateProperties(
            ImmutableArray<string> extendedKeyUsages,
            ImmutableArray<string> keyUsages,
            string subject,
            CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames subjectAlternativeNames,
            int validityInMonths)
        {
            ExtendedKeyUsages = extendedKeyUsages;
            KeyUsages = keyUsages;
            Subject = subject;
            SubjectAlternativeNames = subjectAlternativeNames;
            ValidityInMonths = validityInMonths;
        }
    }

    [OutputType]
    public sealed class CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames
    {
        /// <summary>
        /// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> DnsNames;
        /// <summary>
        /// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> Emails;
        /// <summary>
        /// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> Upns;

        [OutputConstructor]
        private CertifiateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames(
            ImmutableArray<string> dnsNames,
            ImmutableArray<string> emails,
            ImmutableArray<string> upns)
        {
            DnsNames = dnsNames;
            Emails = emails;
            Upns = upns;
        }
    }
    }
}
