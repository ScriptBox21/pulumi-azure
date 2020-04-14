// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ServerStorageProfile struct {
	// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
	AutoGrow *string `pulumi:"autoGrow"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`, not supported for the `basic` tier.
	GeoRedundantBackup *string `pulumi:"geoRedundantBackup"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb int `pulumi:"storageMb"`
}

// ServerStorageProfileInput is an input type that accepts ServerStorageProfileArgs and ServerStorageProfileOutput values.
// You can construct a concrete instance of `ServerStorageProfileInput` via:
//
// 		 ServerStorageProfileArgs{...}
//
type ServerStorageProfileInput interface {
	pulumi.Input

	ToServerStorageProfileOutput() ServerStorageProfileOutput
	ToServerStorageProfileOutputWithContext(context.Context) ServerStorageProfileOutput
}

type ServerStorageProfileArgs struct {
	// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
	AutoGrow pulumi.StringPtrInput `pulumi:"autoGrow"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntPtrInput `pulumi:"backupRetentionDays"`
	// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`, not supported for the `basic` tier.
	GeoRedundantBackup pulumi.StringPtrInput `pulumi:"geoRedundantBackup"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
	StorageMb pulumi.IntInput `pulumi:"storageMb"`
}

func (ServerStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerStorageProfile)(nil)).Elem()
}

func (i ServerStorageProfileArgs) ToServerStorageProfileOutput() ServerStorageProfileOutput {
	return i.ToServerStorageProfileOutputWithContext(context.Background())
}

func (i ServerStorageProfileArgs) ToServerStorageProfileOutputWithContext(ctx context.Context) ServerStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerStorageProfileOutput)
}

func (i ServerStorageProfileArgs) ToServerStorageProfilePtrOutput() ServerStorageProfilePtrOutput {
	return i.ToServerStorageProfilePtrOutputWithContext(context.Background())
}

func (i ServerStorageProfileArgs) ToServerStorageProfilePtrOutputWithContext(ctx context.Context) ServerStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerStorageProfileOutput).ToServerStorageProfilePtrOutputWithContext(ctx)
}

// ServerStorageProfilePtrInput is an input type that accepts ServerStorageProfileArgs, ServerStorageProfilePtr and ServerStorageProfilePtrOutput values.
// You can construct a concrete instance of `ServerStorageProfilePtrInput` via:
//
// 		 ServerStorageProfileArgs{...}
//
//  or:
//
// 		 nil
//
type ServerStorageProfilePtrInput interface {
	pulumi.Input

	ToServerStorageProfilePtrOutput() ServerStorageProfilePtrOutput
	ToServerStorageProfilePtrOutputWithContext(context.Context) ServerStorageProfilePtrOutput
}

type serverStorageProfilePtrType ServerStorageProfileArgs

func ServerStorageProfilePtr(v *ServerStorageProfileArgs) ServerStorageProfilePtrInput {
	return (*serverStorageProfilePtrType)(v)
}

func (*serverStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerStorageProfile)(nil)).Elem()
}

func (i *serverStorageProfilePtrType) ToServerStorageProfilePtrOutput() ServerStorageProfilePtrOutput {
	return i.ToServerStorageProfilePtrOutputWithContext(context.Background())
}

func (i *serverStorageProfilePtrType) ToServerStorageProfilePtrOutputWithContext(ctx context.Context) ServerStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerStorageProfilePtrOutput)
}

type ServerStorageProfileOutput struct{ *pulumi.OutputState }

func (ServerStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerStorageProfile)(nil)).Elem()
}

func (o ServerStorageProfileOutput) ToServerStorageProfileOutput() ServerStorageProfileOutput {
	return o
}

func (o ServerStorageProfileOutput) ToServerStorageProfileOutputWithContext(ctx context.Context) ServerStorageProfileOutput {
	return o
}

func (o ServerStorageProfileOutput) ToServerStorageProfilePtrOutput() ServerStorageProfilePtrOutput {
	return o.ToServerStorageProfilePtrOutputWithContext(context.Background())
}

func (o ServerStorageProfileOutput) ToServerStorageProfilePtrOutputWithContext(ctx context.Context) ServerStorageProfilePtrOutput {
	return o.ApplyT(func(v ServerStorageProfile) *ServerStorageProfile {
		return &v
	}).(ServerStorageProfilePtrOutput)
}

// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
func (o ServerStorageProfileOutput) AutoGrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerStorageProfile) *string { return v.AutoGrow }).(pulumi.StringPtrOutput)
}

// Backup retention days for the server, supported values are between `7` and `35` days.
func (o ServerStorageProfileOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServerStorageProfile) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`, not supported for the `basic` tier.
func (o ServerStorageProfileOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerStorageProfile) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
func (o ServerStorageProfileOutput) StorageMb() pulumi.IntOutput {
	return o.ApplyT(func(v ServerStorageProfile) int { return v.StorageMb }).(pulumi.IntOutput)
}

type ServerStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (ServerStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerStorageProfile)(nil)).Elem()
}

func (o ServerStorageProfilePtrOutput) ToServerStorageProfilePtrOutput() ServerStorageProfilePtrOutput {
	return o
}

func (o ServerStorageProfilePtrOutput) ToServerStorageProfilePtrOutputWithContext(ctx context.Context) ServerStorageProfilePtrOutput {
	return o
}

func (o ServerStorageProfilePtrOutput) Elem() ServerStorageProfileOutput {
	return o.ApplyT(func(v *ServerStorageProfile) ServerStorageProfile { return *v }).(ServerStorageProfileOutput)
}

// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
func (o ServerStorageProfilePtrOutput) AutoGrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerStorageProfile) *string { return v.AutoGrow }).(pulumi.StringPtrOutput)
}

// Backup retention days for the server, supported values are between `7` and `35` days.
func (o ServerStorageProfilePtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServerStorageProfile) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`, not supported for the `basic` tier.
func (o ServerStorageProfilePtrOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerStorageProfile) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `4194304` MB(4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mysql/servers/create#StorageProfile).
func (o ServerStorageProfilePtrOutput) StorageMb() pulumi.IntOutput {
	return o.ApplyT(func(v ServerStorageProfile) int { return v.StorageMb }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerStorageProfileOutput{})
	pulumi.RegisterOutputType(ServerStorageProfilePtrOutput{})
}
