// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appservice

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing App Service Environment
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/app_service_environment.html.markdown.
func GetAppServiceEnvironment(ctx *pulumi.Context, args *GetAppServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*GetAppServiceEnvironmentResult, error) {
	var rv GetAppServiceEnvironmentResult
	err := ctx.Invoke("azure:appservice/getAppServiceEnvironment:getAppServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppServiceEnvironment.
type GetAppServiceEnvironmentArgs struct {
	// The name of the App Service Environment.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the App Service Environment exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAppServiceEnvironment.
type GetAppServiceEnvironmentResult struct {
	// The number of app instances per App Service Environment Front End
	FrontEndScaleFactor int `pulumi:"frontEndScaleFactor"`
	// id is the provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The Pricing Tier (Isolated SKU) of the App Service Environment.
	PricingTier       string `pulumi:"pricingTier"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}
