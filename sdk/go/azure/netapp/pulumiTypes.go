// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AccountActiveDirectory struct {
	// A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
	DnsServers []string `pulumi:"dnsServers"`
	// The name of the Active Directory domain.
	Domain string `pulumi:"domain"`
	// The Organizational Unit (OU) within the Active Directory Domain.
	OrganizationalUnit *string `pulumi:"organizationalUnit"`
	// The password associated with the `username`.
	Password string `pulumi:"password"`
	// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
	SmbServerName string `pulumi:"smbServerName"`
	// The Username of Active Directory Domain Administrator.
	Username string `pulumi:"username"`
}

// AccountActiveDirectoryInput is an input type that accepts AccountActiveDirectoryArgs and AccountActiveDirectoryOutput values.
// You can construct a concrete instance of `AccountActiveDirectoryInput` via:
//
// 		 AccountActiveDirectoryArgs{...}
//
type AccountActiveDirectoryInput interface {
	pulumi.Input

	ToAccountActiveDirectoryOutput() AccountActiveDirectoryOutput
	ToAccountActiveDirectoryOutputWithContext(context.Context) AccountActiveDirectoryOutput
}

type AccountActiveDirectoryArgs struct {
	// A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
	// The name of the Active Directory domain.
	Domain pulumi.StringInput `pulumi:"domain"`
	// The Organizational Unit (OU) within the Active Directory Domain.
	OrganizationalUnit pulumi.StringPtrInput `pulumi:"organizationalUnit"`
	// The password associated with the `username`.
	Password pulumi.StringInput `pulumi:"password"`
	// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
	SmbServerName pulumi.StringInput `pulumi:"smbServerName"`
	// The Username of Active Directory Domain Administrator.
	Username pulumi.StringInput `pulumi:"username"`
}

func (AccountActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountActiveDirectory)(nil)).Elem()
}

func (i AccountActiveDirectoryArgs) ToAccountActiveDirectoryOutput() AccountActiveDirectoryOutput {
	return i.ToAccountActiveDirectoryOutputWithContext(context.Background())
}

func (i AccountActiveDirectoryArgs) ToAccountActiveDirectoryOutputWithContext(ctx context.Context) AccountActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountActiveDirectoryOutput)
}

func (i AccountActiveDirectoryArgs) ToAccountActiveDirectoryPtrOutput() AccountActiveDirectoryPtrOutput {
	return i.ToAccountActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i AccountActiveDirectoryArgs) ToAccountActiveDirectoryPtrOutputWithContext(ctx context.Context) AccountActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountActiveDirectoryOutput).ToAccountActiveDirectoryPtrOutputWithContext(ctx)
}

// AccountActiveDirectoryPtrInput is an input type that accepts AccountActiveDirectoryArgs, AccountActiveDirectoryPtr and AccountActiveDirectoryPtrOutput values.
// You can construct a concrete instance of `AccountActiveDirectoryPtrInput` via:
//
// 		 AccountActiveDirectoryArgs{...}
//
//  or:
//
// 		 nil
//
type AccountActiveDirectoryPtrInput interface {
	pulumi.Input

	ToAccountActiveDirectoryPtrOutput() AccountActiveDirectoryPtrOutput
	ToAccountActiveDirectoryPtrOutputWithContext(context.Context) AccountActiveDirectoryPtrOutput
}

type accountActiveDirectoryPtrType AccountActiveDirectoryArgs

func AccountActiveDirectoryPtr(v *AccountActiveDirectoryArgs) AccountActiveDirectoryPtrInput {
	return (*accountActiveDirectoryPtrType)(v)
}

func (*accountActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountActiveDirectory)(nil)).Elem()
}

func (i *accountActiveDirectoryPtrType) ToAccountActiveDirectoryPtrOutput() AccountActiveDirectoryPtrOutput {
	return i.ToAccountActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *accountActiveDirectoryPtrType) ToAccountActiveDirectoryPtrOutputWithContext(ctx context.Context) AccountActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountActiveDirectoryPtrOutput)
}

type AccountActiveDirectoryOutput struct{ *pulumi.OutputState }

func (AccountActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountActiveDirectory)(nil)).Elem()
}

func (o AccountActiveDirectoryOutput) ToAccountActiveDirectoryOutput() AccountActiveDirectoryOutput {
	return o
}

func (o AccountActiveDirectoryOutput) ToAccountActiveDirectoryOutputWithContext(ctx context.Context) AccountActiveDirectoryOutput {
	return o
}

func (o AccountActiveDirectoryOutput) ToAccountActiveDirectoryPtrOutput() AccountActiveDirectoryPtrOutput {
	return o.ToAccountActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o AccountActiveDirectoryOutput) ToAccountActiveDirectoryPtrOutputWithContext(ctx context.Context) AccountActiveDirectoryPtrOutput {
	return o.ApplyT(func(v AccountActiveDirectory) *AccountActiveDirectory {
		return &v
	}).(AccountActiveDirectoryPtrOutput)
}

// A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
func (o AccountActiveDirectoryOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccountActiveDirectory) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

// The name of the Active Directory domain.
func (o AccountActiveDirectoryOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.Domain }).(pulumi.StringOutput)
}

// The Organizational Unit (OU) within the Active Directory Domain.
func (o AccountActiveDirectoryOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountActiveDirectory) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

// The password associated with the `username`.
func (o AccountActiveDirectoryOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.Password }).(pulumi.StringOutput)
}

// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
func (o AccountActiveDirectoryOutput) SmbServerName() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.SmbServerName }).(pulumi.StringOutput)
}

// The Username of Active Directory Domain Administrator.
func (o AccountActiveDirectoryOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.Username }).(pulumi.StringOutput)
}

type AccountActiveDirectoryPtrOutput struct{ *pulumi.OutputState }

func (AccountActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountActiveDirectory)(nil)).Elem()
}

func (o AccountActiveDirectoryPtrOutput) ToAccountActiveDirectoryPtrOutput() AccountActiveDirectoryPtrOutput {
	return o
}

func (o AccountActiveDirectoryPtrOutput) ToAccountActiveDirectoryPtrOutputWithContext(ctx context.Context) AccountActiveDirectoryPtrOutput {
	return o
}

func (o AccountActiveDirectoryPtrOutput) Elem() AccountActiveDirectoryOutput {
	return o.ApplyT(func(v *AccountActiveDirectory) AccountActiveDirectory { return *v }).(AccountActiveDirectoryOutput)
}

// A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
func (o AccountActiveDirectoryPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AccountActiveDirectory) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

// The name of the Active Directory domain.
func (o AccountActiveDirectoryPtrOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.Domain }).(pulumi.StringOutput)
}

// The Organizational Unit (OU) within the Active Directory Domain.
func (o AccountActiveDirectoryPtrOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountActiveDirectory) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

// The password associated with the `username`.
func (o AccountActiveDirectoryPtrOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.Password }).(pulumi.StringOutput)
}

// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
func (o AccountActiveDirectoryPtrOutput) SmbServerName() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.SmbServerName }).(pulumi.StringOutput)
}

// The Username of Active Directory Domain Administrator.
func (o AccountActiveDirectoryPtrOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v AccountActiveDirectory) string { return v.Username }).(pulumi.StringOutput)
}

type VolumeExportPolicyRule struct {
	// A list of allowed clients IPv4 addresses.
	AllowedClients []string `pulumi:"allowedClients"`
	// Is the CIFS protocol allowed?
	CifsEnabled *bool `pulumi:"cifsEnabled"`
	// Is the NFSv3 protocol allowed?
	Nfsv3Enabled *bool `pulumi:"nfsv3Enabled"`
	// Is the NFSv4 protocol allowed?
	Nfsv4Enabled *bool `pulumi:"nfsv4Enabled"`
	// A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifsEnabled`, `nfsv3Enabled` and `nfsv4Enabled`.
	ProtocolsEnabled *string `pulumi:"protocolsEnabled"`
	// The index number of the rule.
	RuleIndex int `pulumi:"ruleIndex"`
	// Is the file system on unix read only?
	UnixReadOnly *bool `pulumi:"unixReadOnly"`
	// Is the file system on unix read and write?
	UnixReadWrite *bool `pulumi:"unixReadWrite"`
}

// VolumeExportPolicyRuleInput is an input type that accepts VolumeExportPolicyRuleArgs and VolumeExportPolicyRuleOutput values.
// You can construct a concrete instance of `VolumeExportPolicyRuleInput` via:
//
// 		 VolumeExportPolicyRuleArgs{...}
//
type VolumeExportPolicyRuleInput interface {
	pulumi.Input

	ToVolumeExportPolicyRuleOutput() VolumeExportPolicyRuleOutput
	ToVolumeExportPolicyRuleOutputWithContext(context.Context) VolumeExportPolicyRuleOutput
}

type VolumeExportPolicyRuleArgs struct {
	// A list of allowed clients IPv4 addresses.
	AllowedClients pulumi.StringArrayInput `pulumi:"allowedClients"`
	// Is the CIFS protocol allowed?
	CifsEnabled pulumi.BoolPtrInput `pulumi:"cifsEnabled"`
	// Is the NFSv3 protocol allowed?
	Nfsv3Enabled pulumi.BoolPtrInput `pulumi:"nfsv3Enabled"`
	// Is the NFSv4 protocol allowed?
	Nfsv4Enabled pulumi.BoolPtrInput `pulumi:"nfsv4Enabled"`
	// A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifsEnabled`, `nfsv3Enabled` and `nfsv4Enabled`.
	ProtocolsEnabled pulumi.StringPtrInput `pulumi:"protocolsEnabled"`
	// The index number of the rule.
	RuleIndex pulumi.IntInput `pulumi:"ruleIndex"`
	// Is the file system on unix read only?
	UnixReadOnly pulumi.BoolPtrInput `pulumi:"unixReadOnly"`
	// Is the file system on unix read and write?
	UnixReadWrite pulumi.BoolPtrInput `pulumi:"unixReadWrite"`
}

func (VolumeExportPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeExportPolicyRule)(nil)).Elem()
}

func (i VolumeExportPolicyRuleArgs) ToVolumeExportPolicyRuleOutput() VolumeExportPolicyRuleOutput {
	return i.ToVolumeExportPolicyRuleOutputWithContext(context.Background())
}

func (i VolumeExportPolicyRuleArgs) ToVolumeExportPolicyRuleOutputWithContext(ctx context.Context) VolumeExportPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeExportPolicyRuleOutput)
}

// VolumeExportPolicyRuleArrayInput is an input type that accepts VolumeExportPolicyRuleArray and VolumeExportPolicyRuleArrayOutput values.
// You can construct a concrete instance of `VolumeExportPolicyRuleArrayInput` via:
//
// 		 VolumeExportPolicyRuleArray{ VolumeExportPolicyRuleArgs{...} }
//
type VolumeExportPolicyRuleArrayInput interface {
	pulumi.Input

	ToVolumeExportPolicyRuleArrayOutput() VolumeExportPolicyRuleArrayOutput
	ToVolumeExportPolicyRuleArrayOutputWithContext(context.Context) VolumeExportPolicyRuleArrayOutput
}

type VolumeExportPolicyRuleArray []VolumeExportPolicyRuleInput

func (VolumeExportPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeExportPolicyRule)(nil)).Elem()
}

func (i VolumeExportPolicyRuleArray) ToVolumeExportPolicyRuleArrayOutput() VolumeExportPolicyRuleArrayOutput {
	return i.ToVolumeExportPolicyRuleArrayOutputWithContext(context.Background())
}

func (i VolumeExportPolicyRuleArray) ToVolumeExportPolicyRuleArrayOutputWithContext(ctx context.Context) VolumeExportPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeExportPolicyRuleArrayOutput)
}

type VolumeExportPolicyRuleOutput struct{ *pulumi.OutputState }

func (VolumeExportPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeExportPolicyRule)(nil)).Elem()
}

func (o VolumeExportPolicyRuleOutput) ToVolumeExportPolicyRuleOutput() VolumeExportPolicyRuleOutput {
	return o
}

func (o VolumeExportPolicyRuleOutput) ToVolumeExportPolicyRuleOutputWithContext(ctx context.Context) VolumeExportPolicyRuleOutput {
	return o
}

// A list of allowed clients IPv4 addresses.
func (o VolumeExportPolicyRuleOutput) AllowedClients() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) []string { return v.AllowedClients }).(pulumi.StringArrayOutput)
}

// Is the CIFS protocol allowed?
func (o VolumeExportPolicyRuleOutput) CifsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.CifsEnabled }).(pulumi.BoolPtrOutput)
}

// Is the NFSv3 protocol allowed?
func (o VolumeExportPolicyRuleOutput) Nfsv3Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.Nfsv3Enabled }).(pulumi.BoolPtrOutput)
}

// Is the NFSv4 protocol allowed?
func (o VolumeExportPolicyRuleOutput) Nfsv4Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.Nfsv4Enabled }).(pulumi.BoolPtrOutput)
}

// A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifsEnabled`, `nfsv3Enabled` and `nfsv4Enabled`.
func (o VolumeExportPolicyRuleOutput) ProtocolsEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *string { return v.ProtocolsEnabled }).(pulumi.StringPtrOutput)
}

// The index number of the rule.
func (o VolumeExportPolicyRuleOutput) RuleIndex() pulumi.IntOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) int { return v.RuleIndex }).(pulumi.IntOutput)
}

// Is the file system on unix read only?
func (o VolumeExportPolicyRuleOutput) UnixReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.UnixReadOnly }).(pulumi.BoolPtrOutput)
}

// Is the file system on unix read and write?
func (o VolumeExportPolicyRuleOutput) UnixReadWrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.UnixReadWrite }).(pulumi.BoolPtrOutput)
}

type VolumeExportPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (VolumeExportPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeExportPolicyRule)(nil)).Elem()
}

func (o VolumeExportPolicyRuleArrayOutput) ToVolumeExportPolicyRuleArrayOutput() VolumeExportPolicyRuleArrayOutput {
	return o
}

func (o VolumeExportPolicyRuleArrayOutput) ToVolumeExportPolicyRuleArrayOutputWithContext(ctx context.Context) VolumeExportPolicyRuleArrayOutput {
	return o
}

func (o VolumeExportPolicyRuleArrayOutput) Index(i pulumi.IntInput) VolumeExportPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeExportPolicyRule {
		return vs[0].([]VolumeExportPolicyRule)[vs[1].(int)]
	}).(VolumeExportPolicyRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AccountActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(VolumeExportPolicyRuleOutput{})
	pulumi.RegisterOutputType(VolumeExportPolicyRuleArrayOutput{})
}
