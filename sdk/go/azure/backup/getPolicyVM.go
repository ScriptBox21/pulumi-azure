// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package backup

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing VM Backup Policy.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/backup_policy_vm.markdown.
func LookupPolicyVM(ctx *pulumi.Context, args *LookupPolicyVMArgs, opts ...pulumi.InvokeOption) (*LookupPolicyVMResult, error) {
	var rv LookupPolicyVMResult
	err := ctx.Invoke("azure:backup/getPolicyVM:getPolicyVM", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyVM.
type LookupPolicyVMArgs struct {
	// Specifies the name of the VM Backup Policy.
	Name string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which the VM Backup Policy resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getPolicyVM.
type LookupPolicyVMResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}
