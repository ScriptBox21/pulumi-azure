// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing API Management API.
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	var rv LookupApiResult
	err := ctx.Invoke("azure:apimanagement/getApi:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApi.
type LookupApiArgs struct {
	// The name of the API Management Service in which the API Management API exists.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The name of the API Management API.
	Name string `pulumi:"name"`
	// The Name of the Resource Group in which the API Management Service exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Revision of the API Management API.
	Revision string `pulumi:"revision"`
}

// A collection of values returned by getApi.
type LookupApiResult struct {
	ApiManagementName string `pulumi:"apiManagementName"`
	// A description of the API Management API, which may include HTML formatting tags.
	Description string `pulumi:"description"`
	// The display name of the API.
	DisplayName string `pulumi:"displayName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Is this the current API Revision?
	IsCurrent bool `pulumi:"isCurrent"`
	// Is this API Revision online/accessible via the Gateway?
	IsOnline bool   `pulumi:"isOnline"`
	Name     string `pulumi:"name"`
	// The Path for this API Management API.
	Path string `pulumi:"path"`
	// A list of protocols the operations in this API can be invoked.
	Protocols         []string `pulumi:"protocols"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Revision          string   `pulumi:"revision"`
	// Absolute URL of the backend service implementing this API.
	ServiceUrl string `pulumi:"serviceUrl"`
	// Should this API expose a SOAP frontend, rather than a HTTP frontend?
	SoapPassThrough bool `pulumi:"soapPassThrough"`
	// A `subscriptionKeyParameterNames` block as documented below.
	SubscriptionKeyParameterNames []GetApiSubscriptionKeyParameterName `pulumi:"subscriptionKeyParameterNames"`
	// The Version number of this API, if this API is versioned.
	Version string `pulumi:"version"`
	// The ID of the Version Set which this API is associated with.
	VersionSetId string `pulumi:"versionSetId"`
}
