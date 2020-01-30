// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing API Management Group.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/api_management_group.html.markdown.
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azure:apimanagement/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The Name of the API Management Service in which this Group exists.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The Name of the API Management Group.
	Name string `pulumi:"name"`
	// The Name of the Resource Group in which the API Management Service exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getGroup.
type LookupGroupResult struct {
	ApiManagementName string `pulumi:"apiManagementName"`
	// The description of this API Management Group.
	Description string `pulumi:"description"`
	// The display name of this API Management Group.
	DisplayName string `pulumi:"displayName"`
	// The identifier of the external Group.
	ExternalId string `pulumi:"externalId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of this API Management Group, such as `custom` or `external`.
	Type string `pulumi:"type"`
}

