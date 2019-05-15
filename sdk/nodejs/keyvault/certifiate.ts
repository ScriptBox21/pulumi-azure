// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Key Vault Certificate.
 * 
 * ## Example Usage (Importing a PFX)
 * 
 * > **Note:** this example assumed the PFX file is located in the same directory at `certificate-to-import.pfx`.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "key-vault-certificate-example",
 * });
 * const current = pulumi.output(azure.core.getClientConfig({}));
 * const testKeyVault = new azure.keyvault.KeyVault("test", {
 *     accessPolicies: [{
 *         certificatePermissions: [
 *             "create",
 *             "delete",
 *             "deleteissuers",
 *             "get",
 *             "getissuers",
 *             "import",
 *             "list",
 *             "listissuers",
 *             "managecontacts",
 *             "manageissuers",
 *             "setissuers",
 *             "update",
 *         ],
 *         keyPermissions: [
 *             "backup",
 *             "create",
 *             "decrypt",
 *             "delete",
 *             "encrypt",
 *             "get",
 *             "import",
 *             "list",
 *             "purge",
 *             "recover",
 *             "restore",
 *             "sign",
 *             "unwrapKey",
 *             "update",
 *             "verify",
 *             "wrapKey",
 *         ],
 *         objectId: current.servicePrincipalObjectId,
 *         secretPermissions: [
 *             "backup",
 *             "delete",
 *             "get",
 *             "list",
 *             "purge",
 *             "recover",
 *             "restore",
 *             "set",
 *         ],
 *         tenantId: current.tenantId,
 *     }],
 *     location: testResourceGroup.location,
 *     name: "keyvaultcertexample",
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         name: "standard",
 *     },
 *     tags: {
 *         environment: "Production",
 *     },
 *     tenantId: current.tenantId,
 * });
 * const testCertifiate = new azure.keyvault.Certifiate("test", {
 *     certificate: {
 *         contents: (() => {
 *             throw "tf2pulumi error: NYI: call to filebase64";
 *             return (() => { throw "NYI: call to filebase64"; })();
 *         })(),
 *         password: "",
 *     },
 *     certificatePolicy: {
 *         issuerParameters: {
 *             name: "Self",
 *         },
 *         keyProperties: {
 *             exportable: true,
 *             keySize: 2048,
 *             keyType: "RSA",
 *             reuseKey: false,
 *         },
 *         secretProperties: {
 *             contentType: "application/x-pkcs12",
 *         },
 *     },
 *     keyVaultId: testKeyVault.id,
 *     name: "imported-cert",
 * });
 * ```
 * 
 * ## Example Usage (Generating a new certificate)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "key-vault-certificate-example",
 * });
 * const current = pulumi.output(azure.core.getClientConfig({}));
 * const testKeyVault = new azure.keyvault.KeyVault("test", {
 *     accessPolicies: [{
 *         certificatePermissions: [
 *             "create",
 *             "delete",
 *             "deleteissuers",
 *             "get",
 *             "getissuers",
 *             "import",
 *             "list",
 *             "listissuers",
 *             "managecontacts",
 *             "manageissuers",
 *             "setissuers",
 *             "update",
 *         ],
 *         keyPermissions: [
 *             "backup",
 *             "create",
 *             "decrypt",
 *             "delete",
 *             "encrypt",
 *             "get",
 *             "import",
 *             "list",
 *             "purge",
 *             "recover",
 *             "restore",
 *             "sign",
 *             "unwrapKey",
 *             "update",
 *             "verify",
 *             "wrapKey",
 *         ],
 *         objectId: current.servicePrincipalObjectId,
 *         secretPermissions: [
 *             "backup",
 *             "delete",
 *             "get",
 *             "list",
 *             "purge",
 *             "recover",
 *             "restore",
 *             "set",
 *         ],
 *         tenantId: current.tenantId,
 *     }],
 *     location: testResourceGroup.location,
 *     name: "keyvaultcertexample",
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         name: "standard",
 *     },
 *     tags: {
 *         environment: "Production",
 *     },
 *     tenantId: current.tenantId,
 * });
 * const testCertifiate = new azure.keyvault.Certifiate("test", {
 *     certificatePolicy: {
 *         issuerParameters: {
 *             name: "Self",
 *         },
 *         keyProperties: {
 *             exportable: true,
 *             keySize: 2048,
 *             keyType: "RSA",
 *             reuseKey: true,
 *         },
 *         lifetimeActions: [{
 *             action: {
 *                 actionType: "AutoRenew",
 *             },
 *             trigger: {
 *                 daysBeforeExpiry: 30,
 *             },
 *         }],
 *         secretProperties: {
 *             contentType: "application/x-pkcs12",
 *         },
 *         x509CertificateProperties: {
 *             // Server Authentication = 1.3.6.1.5.5.7.3.1
 *             // Client Authentication = 1.3.6.1.5.5.7.3.2
 *             extendedKeyUsages: ["1.3.6.1.5.5.7.3.1"],
 *             keyUsages: [
 *                 "cRLSign",
 *                 "dataEncipherment",
 *                 "digitalSignature",
 *                 "keyAgreement",
 *                 "keyCertSign",
 *                 "keyEncipherment",
 *             ],
 *             subject: "CN=hello-world",
 *             subjectAlternativeNames: {
 *                 dnsNames: [
 *                     "internal.contoso.com",
 *                     "domain.hello.world",
 *                 ],
 *             },
 *             validityInMonths: 12,
 *         },
 *     },
 *     keyVaultId: testKeyVault.id,
 *     name: "generated-cert",
 * });
 * ```
 */
export class Certifiate extends pulumi.CustomResource {
    /**
     * Get an existing Certifiate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertifiateState, opts?: pulumi.CustomResourceOptions): Certifiate {
        return new Certifiate(name, <any>state, { ...opts, id: id });
    }

    /**
     * A `certificate` block as defined below, used to Import an existing certificate.
     */
    public readonly certificate!: pulumi.Output<{ contents: string, password?: string } | undefined>;
    /**
     * The raw Key Vault Certificate.
     */
    public /*out*/ readonly certificateData!: pulumi.Output<string>;
    /**
     * A `certificate_policy` block as defined below.
     */
    public readonly certificatePolicy!: pulumi.Output<{ issuerParameters: { name: string }, keyProperties: { exportable: boolean, keySize: number, keyType: string, reuseKey: boolean }, lifetimeActions?: { action: { actionType: string }, trigger: { daysBeforeExpiry?: number, lifetimePercentage?: number } }[], secretProperties: { contentType: string }, x509CertificateProperties: { extendedKeyUsages: string[], keyUsages: string[], subject: string, subjectAlternativeNames: { dnsNames?: string[], emails?: string[], upns?: string[] }, validityInMonths: number } }>;
    /**
     * The ID of the Key Vault where the Certificate should be created.
     */
    public readonly keyVaultId!: pulumi.Output<string>;
    /**
     * The name of the Certificate Issuer. Possible values include `Self`, or the name of a certificate issuing authority supported by Azure. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the associated Key Vault Secret.
     */
    public /*out*/ readonly secretId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;
    /**
     * The X509 Thumbprint of the Key Vault Certificate returned as hex string.
     */
    public /*out*/ readonly thumbprint!: pulumi.Output<string>;
    public readonly vaultUri!: pulumi.Output<string>;
    /**
     * The current version of the Key Vault Certificate.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Certifiate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertifiateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertifiateArgs | CertifiateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CertifiateState | undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateData"] = state ? state.certificateData : undefined;
            inputs["certificatePolicy"] = state ? state.certificatePolicy : undefined;
            inputs["keyVaultId"] = state ? state.keyVaultId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["secretId"] = state ? state.secretId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["thumbprint"] = state ? state.thumbprint : undefined;
            inputs["vaultUri"] = state ? state.vaultUri : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as CertifiateArgs | undefined;
            if (!args || args.certificatePolicy === undefined) {
                throw new Error("Missing required property 'certificatePolicy'");
            }
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["certificatePolicy"] = args ? args.certificatePolicy : undefined;
            inputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vaultUri"] = args ? args.vaultUri : undefined;
            inputs["certificateData"] = undefined /*out*/;
            inputs["secretId"] = undefined /*out*/;
            inputs["thumbprint"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("azure:keyvault/certifiate:Certifiate", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certifiate resources.
 */
export interface CertifiateState {
    /**
     * A `certificate` block as defined below, used to Import an existing certificate.
     */
    readonly certificate?: pulumi.Input<{ contents: pulumi.Input<string>, password?: pulumi.Input<string> }>;
    /**
     * The raw Key Vault Certificate.
     */
    readonly certificateData?: pulumi.Input<string>;
    /**
     * A `certificate_policy` block as defined below.
     */
    readonly certificatePolicy?: pulumi.Input<{ issuerParameters: pulumi.Input<{ name: pulumi.Input<string> }>, keyProperties: pulumi.Input<{ exportable: pulumi.Input<boolean>, keySize: pulumi.Input<number>, keyType: pulumi.Input<string>, reuseKey: pulumi.Input<boolean> }>, lifetimeActions?: pulumi.Input<pulumi.Input<{ action: pulumi.Input<{ actionType: pulumi.Input<string> }>, trigger: pulumi.Input<{ daysBeforeExpiry?: pulumi.Input<number>, lifetimePercentage?: pulumi.Input<number> }> }>[]>, secretProperties: pulumi.Input<{ contentType: pulumi.Input<string> }>, x509CertificateProperties?: pulumi.Input<{ extendedKeyUsages?: pulumi.Input<pulumi.Input<string>[]>, keyUsages: pulumi.Input<pulumi.Input<string>[]>, subject: pulumi.Input<string>, subjectAlternativeNames?: pulumi.Input<{ dnsNames?: pulumi.Input<pulumi.Input<string>[]>, emails?: pulumi.Input<pulumi.Input<string>[]>, upns?: pulumi.Input<pulumi.Input<string>[]> }>, validityInMonths: pulumi.Input<number> }> }>;
    /**
     * The ID of the Key Vault where the Certificate should be created.
     */
    readonly keyVaultId?: pulumi.Input<string>;
    /**
     * The name of the Certificate Issuer. Possible values include `Self`, or the name of a certificate issuing authority supported by Azure. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the associated Key Vault Secret.
     */
    readonly secretId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The X509 Thumbprint of the Key Vault Certificate returned as hex string.
     */
    readonly thumbprint?: pulumi.Input<string>;
    readonly vaultUri?: pulumi.Input<string>;
    /**
     * The current version of the Key Vault Certificate.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certifiate resource.
 */
export interface CertifiateArgs {
    /**
     * A `certificate` block as defined below, used to Import an existing certificate.
     */
    readonly certificate?: pulumi.Input<{ contents: pulumi.Input<string>, password?: pulumi.Input<string> }>;
    /**
     * A `certificate_policy` block as defined below.
     */
    readonly certificatePolicy: pulumi.Input<{ issuerParameters: pulumi.Input<{ name: pulumi.Input<string> }>, keyProperties: pulumi.Input<{ exportable: pulumi.Input<boolean>, keySize: pulumi.Input<number>, keyType: pulumi.Input<string>, reuseKey: pulumi.Input<boolean> }>, lifetimeActions?: pulumi.Input<pulumi.Input<{ action: pulumi.Input<{ actionType: pulumi.Input<string> }>, trigger: pulumi.Input<{ daysBeforeExpiry?: pulumi.Input<number>, lifetimePercentage?: pulumi.Input<number> }> }>[]>, secretProperties: pulumi.Input<{ contentType: pulumi.Input<string> }>, x509CertificateProperties?: pulumi.Input<{ extendedKeyUsages?: pulumi.Input<pulumi.Input<string>[]>, keyUsages: pulumi.Input<pulumi.Input<string>[]>, subject: pulumi.Input<string>, subjectAlternativeNames?: pulumi.Input<{ dnsNames?: pulumi.Input<pulumi.Input<string>[]>, emails?: pulumi.Input<pulumi.Input<string>[]>, upns?: pulumi.Input<pulumi.Input<string>[]> }>, validityInMonths: pulumi.Input<number> }> }>;
    /**
     * The ID of the Key Vault where the Certificate should be created.
     */
    readonly keyVaultId?: pulumi.Input<string>;
    /**
     * The name of the Certificate Issuer. Possible values include `Self`, or the name of a certificate issuing authority supported by Azure. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly vaultUri?: pulumi.Input<string>;
}
