// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognitive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountNetworkAcls struct {
	// The Default Action to use when no rules match from `ipRules` / `virtualNetworkSubnetIds`. Possible values are `Allow` and `Deny`.
	DefaultAction string `pulumi:"defaultAction"`
	// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
	IpRules []string `pulumi:"ipRules"`
	// One or more Subnet ID's which should be able to access this Cognitive Account.
	VirtualNetworkSubnetIds []string `pulumi:"virtualNetworkSubnetIds"`
}

// AccountNetworkAclsInput is an input type that accepts AccountNetworkAclsArgs and AccountNetworkAclsOutput values.
// You can construct a concrete instance of `AccountNetworkAclsInput` via:
//
//          AccountNetworkAclsArgs{...}
type AccountNetworkAclsInput interface {
	pulumi.Input

	ToAccountNetworkAclsOutput() AccountNetworkAclsOutput
	ToAccountNetworkAclsOutputWithContext(context.Context) AccountNetworkAclsOutput
}

type AccountNetworkAclsArgs struct {
	// The Default Action to use when no rules match from `ipRules` / `virtualNetworkSubnetIds`. Possible values are `Allow` and `Deny`.
	DefaultAction pulumi.StringInput `pulumi:"defaultAction"`
	// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
	IpRules pulumi.StringArrayInput `pulumi:"ipRules"`
	// One or more Subnet ID's which should be able to access this Cognitive Account.
	VirtualNetworkSubnetIds pulumi.StringArrayInput `pulumi:"virtualNetworkSubnetIds"`
}

func (AccountNetworkAclsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountNetworkAcls)(nil)).Elem()
}

func (i AccountNetworkAclsArgs) ToAccountNetworkAclsOutput() AccountNetworkAclsOutput {
	return i.ToAccountNetworkAclsOutputWithContext(context.Background())
}

func (i AccountNetworkAclsArgs) ToAccountNetworkAclsOutputWithContext(ctx context.Context) AccountNetworkAclsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountNetworkAclsOutput)
}

func (i AccountNetworkAclsArgs) ToAccountNetworkAclsPtrOutput() AccountNetworkAclsPtrOutput {
	return i.ToAccountNetworkAclsPtrOutputWithContext(context.Background())
}

func (i AccountNetworkAclsArgs) ToAccountNetworkAclsPtrOutputWithContext(ctx context.Context) AccountNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountNetworkAclsOutput).ToAccountNetworkAclsPtrOutputWithContext(ctx)
}

// AccountNetworkAclsPtrInput is an input type that accepts AccountNetworkAclsArgs, AccountNetworkAclsPtr and AccountNetworkAclsPtrOutput values.
// You can construct a concrete instance of `AccountNetworkAclsPtrInput` via:
//
//          AccountNetworkAclsArgs{...}
//
//  or:
//
//          nil
type AccountNetworkAclsPtrInput interface {
	pulumi.Input

	ToAccountNetworkAclsPtrOutput() AccountNetworkAclsPtrOutput
	ToAccountNetworkAclsPtrOutputWithContext(context.Context) AccountNetworkAclsPtrOutput
}

type accountNetworkAclsPtrType AccountNetworkAclsArgs

func AccountNetworkAclsPtr(v *AccountNetworkAclsArgs) AccountNetworkAclsPtrInput {
	return (*accountNetworkAclsPtrType)(v)
}

func (*accountNetworkAclsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountNetworkAcls)(nil)).Elem()
}

func (i *accountNetworkAclsPtrType) ToAccountNetworkAclsPtrOutput() AccountNetworkAclsPtrOutput {
	return i.ToAccountNetworkAclsPtrOutputWithContext(context.Background())
}

func (i *accountNetworkAclsPtrType) ToAccountNetworkAclsPtrOutputWithContext(ctx context.Context) AccountNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountNetworkAclsPtrOutput)
}

type AccountNetworkAclsOutput struct{ *pulumi.OutputState }

func (AccountNetworkAclsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountNetworkAcls)(nil)).Elem()
}

func (o AccountNetworkAclsOutput) ToAccountNetworkAclsOutput() AccountNetworkAclsOutput {
	return o
}

func (o AccountNetworkAclsOutput) ToAccountNetworkAclsOutputWithContext(ctx context.Context) AccountNetworkAclsOutput {
	return o
}

func (o AccountNetworkAclsOutput) ToAccountNetworkAclsPtrOutput() AccountNetworkAclsPtrOutput {
	return o.ToAccountNetworkAclsPtrOutputWithContext(context.Background())
}

func (o AccountNetworkAclsOutput) ToAccountNetworkAclsPtrOutputWithContext(ctx context.Context) AccountNetworkAclsPtrOutput {
	return o.ApplyT(func(v AccountNetworkAcls) *AccountNetworkAcls {
		return &v
	}).(AccountNetworkAclsPtrOutput)
}

// The Default Action to use when no rules match from `ipRules` / `virtualNetworkSubnetIds`. Possible values are `Allow` and `Deny`.
func (o AccountNetworkAclsOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v AccountNetworkAcls) string { return v.DefaultAction }).(pulumi.StringOutput)
}

// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
func (o AccountNetworkAclsOutput) IpRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccountNetworkAcls) []string { return v.IpRules }).(pulumi.StringArrayOutput)
}

// One or more Subnet ID's which should be able to access this Cognitive Account.
func (o AccountNetworkAclsOutput) VirtualNetworkSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccountNetworkAcls) []string { return v.VirtualNetworkSubnetIds }).(pulumi.StringArrayOutput)
}

type AccountNetworkAclsPtrOutput struct{ *pulumi.OutputState }

func (AccountNetworkAclsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountNetworkAcls)(nil)).Elem()
}

func (o AccountNetworkAclsPtrOutput) ToAccountNetworkAclsPtrOutput() AccountNetworkAclsPtrOutput {
	return o
}

func (o AccountNetworkAclsPtrOutput) ToAccountNetworkAclsPtrOutputWithContext(ctx context.Context) AccountNetworkAclsPtrOutput {
	return o
}

func (o AccountNetworkAclsPtrOutput) Elem() AccountNetworkAclsOutput {
	return o.ApplyT(func(v *AccountNetworkAcls) AccountNetworkAcls { return *v }).(AccountNetworkAclsOutput)
}

// The Default Action to use when no rules match from `ipRules` / `virtualNetworkSubnetIds`. Possible values are `Allow` and `Deny`.
func (o AccountNetworkAclsPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountNetworkAcls) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
func (o AccountNetworkAclsPtrOutput) IpRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccountNetworkAcls) []string {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(pulumi.StringArrayOutput)
}

// One or more Subnet ID's which should be able to access this Cognitive Account.
func (o AccountNetworkAclsPtrOutput) VirtualNetworkSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccountNetworkAcls) []string {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkSubnetIds
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountNetworkAclsOutput{})
	pulumi.RegisterOutputType(AccountNetworkAclsPtrOutput{})
}