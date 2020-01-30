// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package core

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the configuration of the AzureRM provider.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/client_config.html.markdown.
func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	var rv GetClientConfigResult
	err := ctx.Invoke("azure:core/getClientConfig:getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getClientConfig.
type GetClientConfigResult struct {
	ClientId string `pulumi:"clientId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	ObjectId string `pulumi:"objectId"`
	ServicePrincipalApplicationId string `pulumi:"servicePrincipalApplicationId"`
	ServicePrincipalObjectId string `pulumi:"servicePrincipalObjectId"`
	SubscriptionId string `pulumi:"subscriptionId"`
	TenantId string `pulumi:"tenantId"`
}

