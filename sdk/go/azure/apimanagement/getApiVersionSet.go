// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Uses this data source to access information about an API Version Set within an API Management Service.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/api_management_api_version_set.html.markdown.
func LookupApiVersionSet(ctx *pulumi.Context, args *LookupApiVersionSetArgs, opts ...pulumi.InvokeOption) (*LookupApiVersionSetResult, error) {
	var rv LookupApiVersionSetResult
	err := ctx.Invoke("azure:apimanagement/getApiVersionSet:getApiVersionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiVersionSet.
type LookupApiVersionSetArgs struct {
	// The name of the API Management Service where the API Version Set exists.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The name of the API Version Set.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the parent API Management Service exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getApiVersionSet.
type LookupApiVersionSetResult struct {
	ApiManagementName string `pulumi:"apiManagementName"`
	// The description of API Version Set.
	Description string `pulumi:"description"`
	// The display name of this API Version Set.
	DisplayName string `pulumi:"displayName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName string `pulumi:"versionHeaderName"`
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName string `pulumi:"versionQueryName"`
	VersioningScheme string `pulumi:"versioningScheme"`
}

