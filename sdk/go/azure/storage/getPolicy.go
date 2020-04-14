// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Storage Management Policy.
func GetPolicy(ctx *pulumi.Context, args *GetPolicyArgs, opts ...pulumi.InvokeOption) (*GetPolicyResult, error) {
	var rv GetPolicyResult
	err := ctx.Invoke("azure:storage/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type GetPolicyArgs struct {
	// Specifies the id of the storage account to retrieve the management policy for.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// A collection of values returned by getPolicy.
type GetPolicyResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A `rule` block as documented below.
	Rules            []GetPolicyRule `pulumi:"rules"`
	StorageAccountId string          `pulumi:"storageAccountId"`
}
