// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package backup

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure File Share Backup Policy within a Recovery Services vault.
//
// > **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/backup_policy_file_share.markdown.
type PolicyFileShare struct {
	pulumi.CustomResourceState

	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackupOutput `pulumi:"backup"`
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDailyOutput `pulumi:"retentionDaily"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewPolicyFileShare registers a new resource with the given unique name, arguments, and options.
func NewPolicyFileShare(ctx *pulumi.Context,
	name string, args *PolicyFileShareArgs, opts ...pulumi.ResourceOption) (*PolicyFileShare, error) {
	if args == nil || args.Backup == nil {
		return nil, errors.New("missing required argument 'Backup'")
	}
	if args == nil || args.RecoveryVaultName == nil {
		return nil, errors.New("missing required argument 'RecoveryVaultName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RetentionDaily == nil {
		return nil, errors.New("missing required argument 'RetentionDaily'")
	}
	if args == nil {
		args = &PolicyFileShareArgs{}
	}
	var resource PolicyFileShare
	err := ctx.RegisterResource("azure:backup/policyFileShare:PolicyFileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyFileShare gets an existing PolicyFileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyFileShareState, opts ...pulumi.ResourceOption) (*PolicyFileShare, error) {
	var resource PolicyFileShare
	err := ctx.ReadResource("azure:backup/policyFileShare:PolicyFileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyFileShare resources.
type policyFileShareState struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup *PolicyFileShareBackup `pulumi:"backup"`
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily *PolicyFileShareRetentionDaily `pulumi:"retentionDaily"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone *string `pulumi:"timezone"`
}

type PolicyFileShareState struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackupPtrInput
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringPtrInput
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDailyPtrInput
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrInput
}

func (PolicyFileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFileShareState)(nil)).Elem()
}

type policyFileShareArgs struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackup `pulumi:"backup"`
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDaily `pulumi:"retentionDaily"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a PolicyFileShare resource.
type PolicyFileShareArgs struct {
	// Configures the Policy backup frequency and times as documented in the `backup` block below.
	Backup PolicyFileShareBackupInput
	// Specifies the name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringInput
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Configures the policy daily retention as documented in the `retentionDaily` block below.
	RetentionDaily PolicyFileShareRetentionDailyInput
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrInput
}

func (PolicyFileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyFileShareArgs)(nil)).Elem()
}
