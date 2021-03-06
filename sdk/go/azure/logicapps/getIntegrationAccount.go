// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Logic App Integration Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/logicapps"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := logicapps.LookupIntegrationAccount(ctx, &logicapps.LookupIntegrationAccountArgs{
// 			Name:              "example-account",
// 			ResourceGroupName: "example-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupIntegrationAccount(ctx *pulumi.Context, args *LookupIntegrationAccountArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountResult, error) {
	var rv LookupIntegrationAccountResult
	err := ctx.Invoke("azure:logicapps/getIntegrationAccount:getIntegrationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIntegrationAccount.
type LookupIntegrationAccountArgs struct {
	// The name of this Logic App Integration Account.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getIntegrationAccount.
type LookupIntegrationAccountResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region where the Logic App Integration Account exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku name of the Logic App Integration Account.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags assigned to the Logic App Integration Account.
	Tags map[string]string `pulumi:"tags"`
}
