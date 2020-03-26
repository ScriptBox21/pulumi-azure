// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package netapp

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Uses this data source to access information about an existing NetApp Pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/netapp_pool.html.markdown.
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure:netapp/getPool:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPool.
type LookupPoolArgs struct {
	// The name of the NetApp account where the NetApp pool exists.
	AccountName string `pulumi:"accountName"`
	// The name of the NetApp Pool.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the NetApp Pool exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getPool.
type LookupPoolResult struct {
	AccountName string `pulumi:"accountName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region where the NetApp Pool exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service level of the file system.
	ServiceLevel string `pulumi:"serviceLevel"`
	// Provisioned size of the pool in TB.
	SizeInTb int `pulumi:"sizeInTb"`
}
