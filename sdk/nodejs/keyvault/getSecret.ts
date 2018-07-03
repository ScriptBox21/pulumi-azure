// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Returns information about the specified Key Vault Secret.
 * 
 * ~> **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](/docs/state/sensitive-data.html).
 */
export function getSecret(args: GetSecretArgs): Promise<GetSecretResult> {
    return pulumi.runtime.invoke("azure:keyvault/getSecret:getSecret", {
        "name": args.name,
        "vaultUri": args.vaultUri,
    });
}

/**
 * A collection of arguments for invoking getSecret.
 */
export interface GetSecretArgs {
    /**
     * Specifies the name of the Key Vault Secret.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` Data Source / Resource.
     */
    readonly vaultUri: pulumi.Input<string>;
}

/**
 * A collection of values returned by getSecret.
 */
export interface GetSecretResult {
    /**
     * The content type for the Key Vault Secret.
     */
    readonly contentType: string;
    /**
     * Any tags assigned to this resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * The value of the Key Vault Secret.
     */
    readonly value: string;
    /**
     * The current version of the Key Vault Secret.
     */
    readonly version: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
