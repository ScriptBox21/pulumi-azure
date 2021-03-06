// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//          AccountActiveDirectoryArgs{...}
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
//          AccountActiveDirectoryArgs{...}
//
//  or:
//
//          nil
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
	return o.ApplyT(func(v *AccountActiveDirectory) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

// The name of the Active Directory domain.
func (o AccountActiveDirectoryPtrOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return &v.Domain
	}).(pulumi.StringPtrOutput)
}

// The Organizational Unit (OU) within the Active Directory Domain.
func (o AccountActiveDirectoryPtrOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationalUnit
	}).(pulumi.StringPtrOutput)
}

// The password associated with the `username`.
func (o AccountActiveDirectoryPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
func (o AccountActiveDirectoryPtrOutput) SmbServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return &v.SmbServerName
	}).(pulumi.StringPtrOutput)
}

// The Username of Active Directory Domain Administrator.
func (o AccountActiveDirectoryPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type VolumeDataProtectionReplication struct {
	// The endpoint type, default value is `dst` for destination.
	EndpointType *string `pulumi:"endpointType"`
	// Location of the primary volume.
	RemoteVolumeLocation string `pulumi:"remoteVolumeLocation"`
	// Resource ID of the primary volume.
	RemoteVolumeResourceId string `pulumi:"remoteVolumeResourceId"`
	// Replication frequency, supported values are '10minutes', 'hourly', 'daily', values are case sensitive.
	ReplicationFrequency string `pulumi:"replicationFrequency"`
}

// VolumeDataProtectionReplicationInput is an input type that accepts VolumeDataProtectionReplicationArgs and VolumeDataProtectionReplicationOutput values.
// You can construct a concrete instance of `VolumeDataProtectionReplicationInput` via:
//
//          VolumeDataProtectionReplicationArgs{...}
type VolumeDataProtectionReplicationInput interface {
	pulumi.Input

	ToVolumeDataProtectionReplicationOutput() VolumeDataProtectionReplicationOutput
	ToVolumeDataProtectionReplicationOutputWithContext(context.Context) VolumeDataProtectionReplicationOutput
}

type VolumeDataProtectionReplicationArgs struct {
	// The endpoint type, default value is `dst` for destination.
	EndpointType pulumi.StringPtrInput `pulumi:"endpointType"`
	// Location of the primary volume.
	RemoteVolumeLocation pulumi.StringInput `pulumi:"remoteVolumeLocation"`
	// Resource ID of the primary volume.
	RemoteVolumeResourceId pulumi.StringInput `pulumi:"remoteVolumeResourceId"`
	// Replication frequency, supported values are '10minutes', 'hourly', 'daily', values are case sensitive.
	ReplicationFrequency pulumi.StringInput `pulumi:"replicationFrequency"`
}

func (VolumeDataProtectionReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeDataProtectionReplication)(nil)).Elem()
}

func (i VolumeDataProtectionReplicationArgs) ToVolumeDataProtectionReplicationOutput() VolumeDataProtectionReplicationOutput {
	return i.ToVolumeDataProtectionReplicationOutputWithContext(context.Background())
}

func (i VolumeDataProtectionReplicationArgs) ToVolumeDataProtectionReplicationOutputWithContext(ctx context.Context) VolumeDataProtectionReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeDataProtectionReplicationOutput)
}

func (i VolumeDataProtectionReplicationArgs) ToVolumeDataProtectionReplicationPtrOutput() VolumeDataProtectionReplicationPtrOutput {
	return i.ToVolumeDataProtectionReplicationPtrOutputWithContext(context.Background())
}

func (i VolumeDataProtectionReplicationArgs) ToVolumeDataProtectionReplicationPtrOutputWithContext(ctx context.Context) VolumeDataProtectionReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeDataProtectionReplicationOutput).ToVolumeDataProtectionReplicationPtrOutputWithContext(ctx)
}

// VolumeDataProtectionReplicationPtrInput is an input type that accepts VolumeDataProtectionReplicationArgs, VolumeDataProtectionReplicationPtr and VolumeDataProtectionReplicationPtrOutput values.
// You can construct a concrete instance of `VolumeDataProtectionReplicationPtrInput` via:
//
//          VolumeDataProtectionReplicationArgs{...}
//
//  or:
//
//          nil
type VolumeDataProtectionReplicationPtrInput interface {
	pulumi.Input

	ToVolumeDataProtectionReplicationPtrOutput() VolumeDataProtectionReplicationPtrOutput
	ToVolumeDataProtectionReplicationPtrOutputWithContext(context.Context) VolumeDataProtectionReplicationPtrOutput
}

type volumeDataProtectionReplicationPtrType VolumeDataProtectionReplicationArgs

func VolumeDataProtectionReplicationPtr(v *VolumeDataProtectionReplicationArgs) VolumeDataProtectionReplicationPtrInput {
	return (*volumeDataProtectionReplicationPtrType)(v)
}

func (*volumeDataProtectionReplicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeDataProtectionReplication)(nil)).Elem()
}

func (i *volumeDataProtectionReplicationPtrType) ToVolumeDataProtectionReplicationPtrOutput() VolumeDataProtectionReplicationPtrOutput {
	return i.ToVolumeDataProtectionReplicationPtrOutputWithContext(context.Background())
}

func (i *volumeDataProtectionReplicationPtrType) ToVolumeDataProtectionReplicationPtrOutputWithContext(ctx context.Context) VolumeDataProtectionReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeDataProtectionReplicationPtrOutput)
}

type VolumeDataProtectionReplicationOutput struct{ *pulumi.OutputState }

func (VolumeDataProtectionReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeDataProtectionReplication)(nil)).Elem()
}

func (o VolumeDataProtectionReplicationOutput) ToVolumeDataProtectionReplicationOutput() VolumeDataProtectionReplicationOutput {
	return o
}

func (o VolumeDataProtectionReplicationOutput) ToVolumeDataProtectionReplicationOutputWithContext(ctx context.Context) VolumeDataProtectionReplicationOutput {
	return o
}

func (o VolumeDataProtectionReplicationOutput) ToVolumeDataProtectionReplicationPtrOutput() VolumeDataProtectionReplicationPtrOutput {
	return o.ToVolumeDataProtectionReplicationPtrOutputWithContext(context.Background())
}

func (o VolumeDataProtectionReplicationOutput) ToVolumeDataProtectionReplicationPtrOutputWithContext(ctx context.Context) VolumeDataProtectionReplicationPtrOutput {
	return o.ApplyT(func(v VolumeDataProtectionReplication) *VolumeDataProtectionReplication {
		return &v
	}).(VolumeDataProtectionReplicationPtrOutput)
}

// The endpoint type, default value is `dst` for destination.
func (o VolumeDataProtectionReplicationOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeDataProtectionReplication) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

// Location of the primary volume.
func (o VolumeDataProtectionReplicationOutput) RemoteVolumeLocation() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeDataProtectionReplication) string { return v.RemoteVolumeLocation }).(pulumi.StringOutput)
}

// Resource ID of the primary volume.
func (o VolumeDataProtectionReplicationOutput) RemoteVolumeResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeDataProtectionReplication) string { return v.RemoteVolumeResourceId }).(pulumi.StringOutput)
}

// Replication frequency, supported values are '10minutes', 'hourly', 'daily', values are case sensitive.
func (o VolumeDataProtectionReplicationOutput) ReplicationFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeDataProtectionReplication) string { return v.ReplicationFrequency }).(pulumi.StringOutput)
}

type VolumeDataProtectionReplicationPtrOutput struct{ *pulumi.OutputState }

func (VolumeDataProtectionReplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeDataProtectionReplication)(nil)).Elem()
}

func (o VolumeDataProtectionReplicationPtrOutput) ToVolumeDataProtectionReplicationPtrOutput() VolumeDataProtectionReplicationPtrOutput {
	return o
}

func (o VolumeDataProtectionReplicationPtrOutput) ToVolumeDataProtectionReplicationPtrOutputWithContext(ctx context.Context) VolumeDataProtectionReplicationPtrOutput {
	return o
}

func (o VolumeDataProtectionReplicationPtrOutput) Elem() VolumeDataProtectionReplicationOutput {
	return o.ApplyT(func(v *VolumeDataProtectionReplication) VolumeDataProtectionReplication { return *v }).(VolumeDataProtectionReplicationOutput)
}

// The endpoint type, default value is `dst` for destination.
func (o VolumeDataProtectionReplicationPtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeDataProtectionReplication) *string {
		if v == nil {
			return nil
		}
		return v.EndpointType
	}).(pulumi.StringPtrOutput)
}

// Location of the primary volume.
func (o VolumeDataProtectionReplicationPtrOutput) RemoteVolumeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeDataProtectionReplication) *string {
		if v == nil {
			return nil
		}
		return &v.RemoteVolumeLocation
	}).(pulumi.StringPtrOutput)
}

// Resource ID of the primary volume.
func (o VolumeDataProtectionReplicationPtrOutput) RemoteVolumeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeDataProtectionReplication) *string {
		if v == nil {
			return nil
		}
		return &v.RemoteVolumeResourceId
	}).(pulumi.StringPtrOutput)
}

// Replication frequency, supported values are '10minutes', 'hourly', 'daily', values are case sensitive.
func (o VolumeDataProtectionReplicationPtrOutput) ReplicationFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeDataProtectionReplication) *string {
		if v == nil {
			return nil
		}
		return &v.ReplicationFrequency
	}).(pulumi.StringPtrOutput)
}

type VolumeExportPolicyRule struct {
	// A list of allowed clients IPv4 addresses.
	AllowedClients []string `pulumi:"allowedClients"`
	// Is the CIFS protocol allowed?
	//
	// Deprecated: Deprecated in favour of `protocols_enabled`
	CifsEnabled *bool `pulumi:"cifsEnabled"`
	// Is the NFSv3 protocol allowed?
	//
	// Deprecated: Deprecated in favour of `protocols_enabled`
	Nfsv3Enabled *bool `pulumi:"nfsv3Enabled"`
	// Is the NFSv4 protocol allowed?
	//
	// Deprecated: Deprecated in favour of `protocols_enabled`
	Nfsv4Enabled *bool `pulumi:"nfsv4Enabled"`
	// A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifsEnabled`, `nfsv3Enabled` and `nfsv4Enabled`.
	ProtocolsEnabled *string `pulumi:"protocolsEnabled"`
	// Is root access permitted to this volume?
	RootAccessEnabled *bool `pulumi:"rootAccessEnabled"`
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
//          VolumeExportPolicyRuleArgs{...}
type VolumeExportPolicyRuleInput interface {
	pulumi.Input

	ToVolumeExportPolicyRuleOutput() VolumeExportPolicyRuleOutput
	ToVolumeExportPolicyRuleOutputWithContext(context.Context) VolumeExportPolicyRuleOutput
}

type VolumeExportPolicyRuleArgs struct {
	// A list of allowed clients IPv4 addresses.
	AllowedClients pulumi.StringArrayInput `pulumi:"allowedClients"`
	// Is the CIFS protocol allowed?
	//
	// Deprecated: Deprecated in favour of `protocols_enabled`
	CifsEnabled pulumi.BoolPtrInput `pulumi:"cifsEnabled"`
	// Is the NFSv3 protocol allowed?
	//
	// Deprecated: Deprecated in favour of `protocols_enabled`
	Nfsv3Enabled pulumi.BoolPtrInput `pulumi:"nfsv3Enabled"`
	// Is the NFSv4 protocol allowed?
	//
	// Deprecated: Deprecated in favour of `protocols_enabled`
	Nfsv4Enabled pulumi.BoolPtrInput `pulumi:"nfsv4Enabled"`
	// A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifsEnabled`, `nfsv3Enabled` and `nfsv4Enabled`.
	ProtocolsEnabled pulumi.StringPtrInput `pulumi:"protocolsEnabled"`
	// Is root access permitted to this volume?
	RootAccessEnabled pulumi.BoolPtrInput `pulumi:"rootAccessEnabled"`
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
//          VolumeExportPolicyRuleArray{ VolumeExportPolicyRuleArgs{...} }
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
//
// Deprecated: Deprecated in favour of `protocols_enabled`
func (o VolumeExportPolicyRuleOutput) CifsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.CifsEnabled }).(pulumi.BoolPtrOutput)
}

// Is the NFSv3 protocol allowed?
//
// Deprecated: Deprecated in favour of `protocols_enabled`
func (o VolumeExportPolicyRuleOutput) Nfsv3Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.Nfsv3Enabled }).(pulumi.BoolPtrOutput)
}

// Is the NFSv4 protocol allowed?
//
// Deprecated: Deprecated in favour of `protocols_enabled`
func (o VolumeExportPolicyRuleOutput) Nfsv4Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.Nfsv4Enabled }).(pulumi.BoolPtrOutput)
}

// A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifsEnabled`, `nfsv3Enabled` and `nfsv4Enabled`.
func (o VolumeExportPolicyRuleOutput) ProtocolsEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *string { return v.ProtocolsEnabled }).(pulumi.StringPtrOutput)
}

// Is root access permitted to this volume?
func (o VolumeExportPolicyRuleOutput) RootAccessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeExportPolicyRule) *bool { return v.RootAccessEnabled }).(pulumi.BoolPtrOutput)
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

type GetVolumeDataProtectionReplication struct {
	// The endpoint type.
	EndpointType string `pulumi:"endpointType"`
	// Location of the primary volume.
	RemoteVolumeLocation string `pulumi:"remoteVolumeLocation"`
	// Resource ID of the primary volume.
	RemoteVolumeResourceId string `pulumi:"remoteVolumeResourceId"`
	// Frequency of replication.
	ReplicationFrequency string `pulumi:"replicationFrequency"`
	// Deprecated: This property is not in use and will be removed in version 3.0 of the provider. Please use `replication_frequency` instead
	ReplicationSchedule string `pulumi:"replicationSchedule"`
}

// GetVolumeDataProtectionReplicationInput is an input type that accepts GetVolumeDataProtectionReplicationArgs and GetVolumeDataProtectionReplicationOutput values.
// You can construct a concrete instance of `GetVolumeDataProtectionReplicationInput` via:
//
//          GetVolumeDataProtectionReplicationArgs{...}
type GetVolumeDataProtectionReplicationInput interface {
	pulumi.Input

	ToGetVolumeDataProtectionReplicationOutput() GetVolumeDataProtectionReplicationOutput
	ToGetVolumeDataProtectionReplicationOutputWithContext(context.Context) GetVolumeDataProtectionReplicationOutput
}

type GetVolumeDataProtectionReplicationArgs struct {
	// The endpoint type.
	EndpointType pulumi.StringInput `pulumi:"endpointType"`
	// Location of the primary volume.
	RemoteVolumeLocation pulumi.StringInput `pulumi:"remoteVolumeLocation"`
	// Resource ID of the primary volume.
	RemoteVolumeResourceId pulumi.StringInput `pulumi:"remoteVolumeResourceId"`
	// Frequency of replication.
	ReplicationFrequency pulumi.StringInput `pulumi:"replicationFrequency"`
	// Deprecated: This property is not in use and will be removed in version 3.0 of the provider. Please use `replication_frequency` instead
	ReplicationSchedule pulumi.StringInput `pulumi:"replicationSchedule"`
}

func (GetVolumeDataProtectionReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVolumeDataProtectionReplication)(nil)).Elem()
}

func (i GetVolumeDataProtectionReplicationArgs) ToGetVolumeDataProtectionReplicationOutput() GetVolumeDataProtectionReplicationOutput {
	return i.ToGetVolumeDataProtectionReplicationOutputWithContext(context.Background())
}

func (i GetVolumeDataProtectionReplicationArgs) ToGetVolumeDataProtectionReplicationOutputWithContext(ctx context.Context) GetVolumeDataProtectionReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVolumeDataProtectionReplicationOutput)
}

// GetVolumeDataProtectionReplicationArrayInput is an input type that accepts GetVolumeDataProtectionReplicationArray and GetVolumeDataProtectionReplicationArrayOutput values.
// You can construct a concrete instance of `GetVolumeDataProtectionReplicationArrayInput` via:
//
//          GetVolumeDataProtectionReplicationArray{ GetVolumeDataProtectionReplicationArgs{...} }
type GetVolumeDataProtectionReplicationArrayInput interface {
	pulumi.Input

	ToGetVolumeDataProtectionReplicationArrayOutput() GetVolumeDataProtectionReplicationArrayOutput
	ToGetVolumeDataProtectionReplicationArrayOutputWithContext(context.Context) GetVolumeDataProtectionReplicationArrayOutput
}

type GetVolumeDataProtectionReplicationArray []GetVolumeDataProtectionReplicationInput

func (GetVolumeDataProtectionReplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVolumeDataProtectionReplication)(nil)).Elem()
}

func (i GetVolumeDataProtectionReplicationArray) ToGetVolumeDataProtectionReplicationArrayOutput() GetVolumeDataProtectionReplicationArrayOutput {
	return i.ToGetVolumeDataProtectionReplicationArrayOutputWithContext(context.Background())
}

func (i GetVolumeDataProtectionReplicationArray) ToGetVolumeDataProtectionReplicationArrayOutputWithContext(ctx context.Context) GetVolumeDataProtectionReplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVolumeDataProtectionReplicationArrayOutput)
}

type GetVolumeDataProtectionReplicationOutput struct{ *pulumi.OutputState }

func (GetVolumeDataProtectionReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVolumeDataProtectionReplication)(nil)).Elem()
}

func (o GetVolumeDataProtectionReplicationOutput) ToGetVolumeDataProtectionReplicationOutput() GetVolumeDataProtectionReplicationOutput {
	return o
}

func (o GetVolumeDataProtectionReplicationOutput) ToGetVolumeDataProtectionReplicationOutputWithContext(ctx context.Context) GetVolumeDataProtectionReplicationOutput {
	return o
}

// The endpoint type.
func (o GetVolumeDataProtectionReplicationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v GetVolumeDataProtectionReplication) string { return v.EndpointType }).(pulumi.StringOutput)
}

// Location of the primary volume.
func (o GetVolumeDataProtectionReplicationOutput) RemoteVolumeLocation() pulumi.StringOutput {
	return o.ApplyT(func(v GetVolumeDataProtectionReplication) string { return v.RemoteVolumeLocation }).(pulumi.StringOutput)
}

// Resource ID of the primary volume.
func (o GetVolumeDataProtectionReplicationOutput) RemoteVolumeResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVolumeDataProtectionReplication) string { return v.RemoteVolumeResourceId }).(pulumi.StringOutput)
}

// Frequency of replication.
func (o GetVolumeDataProtectionReplicationOutput) ReplicationFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v GetVolumeDataProtectionReplication) string { return v.ReplicationFrequency }).(pulumi.StringOutput)
}

// Deprecated: This property is not in use and will be removed in version 3.0 of the provider. Please use `replication_frequency` instead
func (o GetVolumeDataProtectionReplicationOutput) ReplicationSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v GetVolumeDataProtectionReplication) string { return v.ReplicationSchedule }).(pulumi.StringOutput)
}

type GetVolumeDataProtectionReplicationArrayOutput struct{ *pulumi.OutputState }

func (GetVolumeDataProtectionReplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVolumeDataProtectionReplication)(nil)).Elem()
}

func (o GetVolumeDataProtectionReplicationArrayOutput) ToGetVolumeDataProtectionReplicationArrayOutput() GetVolumeDataProtectionReplicationArrayOutput {
	return o
}

func (o GetVolumeDataProtectionReplicationArrayOutput) ToGetVolumeDataProtectionReplicationArrayOutputWithContext(ctx context.Context) GetVolumeDataProtectionReplicationArrayOutput {
	return o
}

func (o GetVolumeDataProtectionReplicationArrayOutput) Index(i pulumi.IntInput) GetVolumeDataProtectionReplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVolumeDataProtectionReplication {
		return vs[0].([]GetVolumeDataProtectionReplication)[vs[1].(int)]
	}).(GetVolumeDataProtectionReplicationOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AccountActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(VolumeDataProtectionReplicationOutput{})
	pulumi.RegisterOutputType(VolumeDataProtectionReplicationPtrOutput{})
	pulumi.RegisterOutputType(VolumeExportPolicyRuleOutput{})
	pulumi.RegisterOutputType(VolumeExportPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(GetVolumeDataProtectionReplicationOutput{})
	pulumi.RegisterOutputType(GetVolumeDataProtectionReplicationArrayOutput{})
}
