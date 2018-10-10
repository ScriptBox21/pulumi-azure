// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the azurerm package
 */
export class Provider extends pulumi.ProviderResource {

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            inputs["clientId"] = (args ? args.clientId : undefined) || (utilities.getEnv("ARM_CLIENT_ID") || "");
            inputs["clientSecret"] = (args ? args.clientSecret : undefined) || (utilities.getEnv("ARM_CLIENT_SECRET") || "");
            inputs["environment"] = (args ? args.environment : undefined) || (utilities.getEnv("ARM_ENVIRONMENT") || "public");
            inputs["msiEndpoint"] = (args ? args.msiEndpoint : undefined) || (utilities.getEnv("ARM_MSI_ENDPOINT") || "");
            inputs["skipCredentialsValidation"] = pulumi.output((args ? args.skipCredentialsValidation : undefined) || (utilities.getEnvBoolean("ARM_SKIP_CREDENTIALS_VALIDATION") || false)).apply(JSON.stringify);
            inputs["skipProviderRegistration"] = pulumi.output((args ? args.skipProviderRegistration : undefined) || (utilities.getEnvBoolean("ARM_SKIP_PROVIDER_REGISTRATION") || false)).apply(JSON.stringify);
            inputs["subscriptionId"] = (args ? args.subscriptionId : undefined) || (utilities.getEnv("ARM_SUBSCRIPTION_ID") || "");
            inputs["tenantId"] = (args ? args.tenantId : undefined) || (utilities.getEnv("ARM_TENANT_ID") || "");
            inputs["useMsi"] = pulumi.output((args ? args.useMsi : undefined) || (utilities.getEnvBoolean("ARM_USE_MSI") || false)).apply(JSON.stringify);
        }
        super("azure", name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    readonly clientId?: pulumi.Input<string>;
    readonly clientSecret?: pulumi.Input<string>;
    readonly environment?: pulumi.Input<string>;
    readonly msiEndpoint?: pulumi.Input<string>;
    readonly skipCredentialsValidation?: pulumi.Input<boolean>;
    readonly skipProviderRegistration?: pulumi.Input<boolean>;
    readonly subscriptionId?: pulumi.Input<string>;
    readonly tenantId?: pulumi.Input<string>;
    readonly useMsi?: pulumi.Input<boolean>;
}