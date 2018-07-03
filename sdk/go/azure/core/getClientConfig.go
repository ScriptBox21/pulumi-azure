// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the configuration of the Azure Resource Manager
// provider.
func LookupClientConfig(ctx *pulumi.Context) (*GetClientConfigResult, error) {
	outputs, err := ctx.Invoke("azure:core/getClientConfig:getClientConfig", nil)
	if err != nil {
		return nil, err
	}
	return &GetClientConfigResult{
		ClientId: outputs["clientId"],
		ServicePrincipalApplicationId: outputs["servicePrincipalApplicationId"],
		ServicePrincipalObjectId: outputs["servicePrincipalObjectId"],
		SubscriptionId: outputs["subscriptionId"],
		TenantId: outputs["tenantId"],
		Id: outputs["id"],
	}, nil
}

// A collection of values returned by getClientConfig.
type GetClientConfigResult struct {
	ClientId interface{}
	ServicePrincipalApplicationId interface{}
	ServicePrincipalObjectId interface{}
	SubscriptionId interface{}
	TenantId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
