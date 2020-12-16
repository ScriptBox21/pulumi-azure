// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type VaultIdentity struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	// The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
	Type string `pulumi:"type"`
}

// VaultIdentityInput is an input type that accepts VaultIdentityArgs and VaultIdentityOutput values.
// You can construct a concrete instance of `VaultIdentityInput` via:
//
//          VaultIdentityArgs{...}
type VaultIdentityInput interface {
	pulumi.Input

	ToVaultIdentityOutput() VaultIdentityOutput
	ToVaultIdentityOutputWithContext(context.Context) VaultIdentityOutput
}

type VaultIdentityArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	// The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (VaultIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultIdentity)(nil)).Elem()
}

func (i VaultIdentityArgs) ToVaultIdentityOutput() VaultIdentityOutput {
	return i.ToVaultIdentityOutputWithContext(context.Background())
}

func (i VaultIdentityArgs) ToVaultIdentityOutputWithContext(ctx context.Context) VaultIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultIdentityOutput)
}

func (i VaultIdentityArgs) ToVaultIdentityPtrOutput() VaultIdentityPtrOutput {
	return i.ToVaultIdentityPtrOutputWithContext(context.Background())
}

func (i VaultIdentityArgs) ToVaultIdentityPtrOutputWithContext(ctx context.Context) VaultIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultIdentityOutput).ToVaultIdentityPtrOutputWithContext(ctx)
}

// VaultIdentityPtrInput is an input type that accepts VaultIdentityArgs, VaultIdentityPtr and VaultIdentityPtrOutput values.
// You can construct a concrete instance of `VaultIdentityPtrInput` via:
//
//          VaultIdentityArgs{...}
//
//  or:
//
//          nil
type VaultIdentityPtrInput interface {
	pulumi.Input

	ToVaultIdentityPtrOutput() VaultIdentityPtrOutput
	ToVaultIdentityPtrOutputWithContext(context.Context) VaultIdentityPtrOutput
}

type vaultIdentityPtrType VaultIdentityArgs

func VaultIdentityPtr(v *VaultIdentityArgs) VaultIdentityPtrInput {
	return (*vaultIdentityPtrType)(v)
}

func (*vaultIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultIdentity)(nil)).Elem()
}

func (i *vaultIdentityPtrType) ToVaultIdentityPtrOutput() VaultIdentityPtrOutput {
	return i.ToVaultIdentityPtrOutputWithContext(context.Background())
}

func (i *vaultIdentityPtrType) ToVaultIdentityPtrOutputWithContext(ctx context.Context) VaultIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultIdentityPtrOutput)
}

type VaultIdentityOutput struct{ *pulumi.OutputState }

func (VaultIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultIdentity)(nil)).Elem()
}

func (o VaultIdentityOutput) ToVaultIdentityOutput() VaultIdentityOutput {
	return o
}

func (o VaultIdentityOutput) ToVaultIdentityOutputWithContext(ctx context.Context) VaultIdentityOutput {
	return o
}

func (o VaultIdentityOutput) ToVaultIdentityPtrOutput() VaultIdentityPtrOutput {
	return o.ToVaultIdentityPtrOutputWithContext(context.Background())
}

func (o VaultIdentityOutput) ToVaultIdentityPtrOutputWithContext(ctx context.Context) VaultIdentityPtrOutput {
	return o.ApplyT(func(v VaultIdentity) *VaultIdentity {
		return &v
	}).(VaultIdentityPtrOutput)
}
func (o VaultIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o VaultIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
func (o VaultIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VaultIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type VaultIdentityPtrOutput struct{ *pulumi.OutputState }

func (VaultIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultIdentity)(nil)).Elem()
}

func (o VaultIdentityPtrOutput) ToVaultIdentityPtrOutput() VaultIdentityPtrOutput {
	return o
}

func (o VaultIdentityPtrOutput) ToVaultIdentityPtrOutputWithContext(ctx context.Context) VaultIdentityPtrOutput {
	return o
}

func (o VaultIdentityPtrOutput) Elem() VaultIdentityOutput {
	return o.ApplyT(func(v *VaultIdentity) VaultIdentity { return *v }).(VaultIdentityOutput)
}

func (o VaultIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o VaultIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
func (o VaultIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VaultIdentityOutput{})
	pulumi.RegisterOutputType(VaultIdentityPtrOutput{})
}
