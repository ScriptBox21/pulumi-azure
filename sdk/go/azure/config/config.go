// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func GetAuxiliaryTenantIds(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:auxiliaryTenantIds")
}

// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
// Certificate
func GetClientCertificatePassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:clientCertificatePassword")
}

// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
// Principal using a Client Certificate.
func GetClientCertificatePath(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:clientCertificatePath")
}

// The Client ID which should be used.
func GetClientId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:clientId")
}

// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
func GetClientSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:clientSecret")
}

// This will disable the x-ms-correlation-request-id header.
func GetDisableCorrelationRequestId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azure:disableCorrelationRequestId")
}

// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
func GetDisableTerraformPartnerId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azure:disableTerraformPartnerId")
}

// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
// public.
func GetEnvironment(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azure:environment")
	if err == nil {
		return v
	}
	return getEnvOrDefault("public", nil, "AZURE_ENVIRONMENT", "ARM_ENVIRONMENT").(string)
}
func GetFeatures(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:features")
}
func GetLocation(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azure:location")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "ARM_LOCATION").(string)
}

// The Hostname which should be used for the Azure Metadata Service.
func GetMetadataHost(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azure:metadataHost")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "ARM_METADATA_HOSTNAME").(string)
}

// Deprecated - replaced by `metadata_host`.
//
// Deprecated: use `metadata_host` instead
func GetMetadataUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azure:metadataUrl")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "ARM_METADATA_URL").(string)
}

// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
// automatically.
func GetMsiEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:msiEndpoint")
}

// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
func GetPartnerId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:partnerId")
}

// [DEPRECATED] This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
//
// Deprecated: This field is deprecated and will be removed in version 3.0 of the Azure Provider
func GetSkipCredentialsValidation(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azure:skipCredentialsValidation")
}

// Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
// registered?
func GetSkipProviderRegistration(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "azure:skipProviderRegistration")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "ARM_SKIP_PROVIDER_REGISTRATION").(bool)
}

// Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
func GetStorageUseAzuread(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "azure:storageUseAzuread")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "ARM_STORAGE_USE_AZUREAD").(bool)
}

// The Subscription ID which should be used.
func GetSubscriptionId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azure:subscriptionId")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "ARM_SUBSCRIPTION_ID").(string)
}

// The Tenant ID which should be used.
func GetTenantId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure:tenantId")
}

// Allowed Managed Service Identity be used for Authentication.
func GetUseMsi(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azure:useMsi")
}
