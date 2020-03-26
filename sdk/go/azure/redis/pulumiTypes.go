// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type CachePatchSchedule struct {
	DayOfWeek    string `pulumi:"dayOfWeek"`
	StartHourUtc *int   `pulumi:"startHourUtc"`
}

type CachePatchScheduleInput interface {
	pulumi.Input

	ToCachePatchScheduleOutput() CachePatchScheduleOutput
	ToCachePatchScheduleOutputWithContext(context.Context) CachePatchScheduleOutput
}

type CachePatchScheduleArgs struct {
	DayOfWeek    pulumi.StringInput `pulumi:"dayOfWeek"`
	StartHourUtc pulumi.IntPtrInput `pulumi:"startHourUtc"`
}

func (CachePatchScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CachePatchSchedule)(nil)).Elem()
}

func (i CachePatchScheduleArgs) ToCachePatchScheduleOutput() CachePatchScheduleOutput {
	return i.ToCachePatchScheduleOutputWithContext(context.Background())
}

func (i CachePatchScheduleArgs) ToCachePatchScheduleOutputWithContext(ctx context.Context) CachePatchScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachePatchScheduleOutput)
}

type CachePatchScheduleArrayInput interface {
	pulumi.Input

	ToCachePatchScheduleArrayOutput() CachePatchScheduleArrayOutput
	ToCachePatchScheduleArrayOutputWithContext(context.Context) CachePatchScheduleArrayOutput
}

type CachePatchScheduleArray []CachePatchScheduleInput

func (CachePatchScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CachePatchSchedule)(nil)).Elem()
}

func (i CachePatchScheduleArray) ToCachePatchScheduleArrayOutput() CachePatchScheduleArrayOutput {
	return i.ToCachePatchScheduleArrayOutputWithContext(context.Background())
}

func (i CachePatchScheduleArray) ToCachePatchScheduleArrayOutputWithContext(ctx context.Context) CachePatchScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachePatchScheduleArrayOutput)
}

type CachePatchScheduleOutput struct{ *pulumi.OutputState }

func (CachePatchScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CachePatchSchedule)(nil)).Elem()
}

func (o CachePatchScheduleOutput) ToCachePatchScheduleOutput() CachePatchScheduleOutput {
	return o
}

func (o CachePatchScheduleOutput) ToCachePatchScheduleOutputWithContext(ctx context.Context) CachePatchScheduleOutput {
	return o
}

func (o CachePatchScheduleOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v CachePatchSchedule) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

func (o CachePatchScheduleOutput) StartHourUtc() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CachePatchSchedule) *int { return v.StartHourUtc }).(pulumi.IntPtrOutput)
}

type CachePatchScheduleArrayOutput struct{ *pulumi.OutputState }

func (CachePatchScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CachePatchSchedule)(nil)).Elem()
}

func (o CachePatchScheduleArrayOutput) ToCachePatchScheduleArrayOutput() CachePatchScheduleArrayOutput {
	return o
}

func (o CachePatchScheduleArrayOutput) ToCachePatchScheduleArrayOutputWithContext(ctx context.Context) CachePatchScheduleArrayOutput {
	return o
}

func (o CachePatchScheduleArrayOutput) Index(i pulumi.IntInput) CachePatchScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CachePatchSchedule {
		return vs[0].([]CachePatchSchedule)[vs[1].(int)]
	}).(CachePatchScheduleOutput)
}

type CacheRedisConfiguration struct {
	AofBackupEnabled            *bool   `pulumi:"aofBackupEnabled"`
	AofStorageConnectionString0 *string `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1 *string `pulumi:"aofStorageConnectionString1"`
	// If set to `false`, the Redis instance will be accessible without authentication. Defaults to `true`.
	EnableAuthentication *bool `pulumi:"enableAuthentication"`
	// Returns the max number of connected clients at the same time.
	Maxclients *int `pulumi:"maxclients"`
	// Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
	MaxfragmentationmemoryReserved *int `pulumi:"maxfragmentationmemoryReserved"`
	// The max-memory delta for this Redis instance. Defaults are shown below.
	MaxmemoryDelta *int `pulumi:"maxmemoryDelta"`
	// How Redis will select what to remove when `maxmemory` is reached. Defaults are shown below.
	MaxmemoryPolicy *string `pulumi:"maxmemoryPolicy"`
	// Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
	MaxmemoryReserved *int `pulumi:"maxmemoryReserved"`
	// Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. [Reference](https://redis.io/topics/notifications#configuration)
	NotifyKeyspaceEvents *string `pulumi:"notifyKeyspaceEvents"`
	// Is Backup Enabled? Only supported on Premium SKU's.
	RdbBackupEnabled *bool `pulumi:"rdbBackupEnabled"`
	// The Backup Frequency in Minutes. Only supported on Premium SKU's. Possible values are: `15`, `30`, `60`, `360`, `720` and `1440`.
	RdbBackupFrequency *int `pulumi:"rdbBackupFrequency"`
	// The maximum number of snapshots to create as a backup. Only supported for Premium SKU's.
	RdbBackupMaxSnapshotCount *int `pulumi:"rdbBackupMaxSnapshotCount"`
	// The Connection String to the Storage Account. Only supported for Premium SKU's. In the format: `DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}`.
	RdbStorageConnectionString *string `pulumi:"rdbStorageConnectionString"`
}

type CacheRedisConfigurationInput interface {
	pulumi.Input

	ToCacheRedisConfigurationOutput() CacheRedisConfigurationOutput
	ToCacheRedisConfigurationOutputWithContext(context.Context) CacheRedisConfigurationOutput
}

type CacheRedisConfigurationArgs struct {
	AofBackupEnabled            pulumi.BoolPtrInput   `pulumi:"aofBackupEnabled"`
	AofStorageConnectionString0 pulumi.StringPtrInput `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1 pulumi.StringPtrInput `pulumi:"aofStorageConnectionString1"`
	// If set to `false`, the Redis instance will be accessible without authentication. Defaults to `true`.
	EnableAuthentication pulumi.BoolPtrInput `pulumi:"enableAuthentication"`
	// Returns the max number of connected clients at the same time.
	Maxclients pulumi.IntPtrInput `pulumi:"maxclients"`
	// Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
	MaxfragmentationmemoryReserved pulumi.IntPtrInput `pulumi:"maxfragmentationmemoryReserved"`
	// The max-memory delta for this Redis instance. Defaults are shown below.
	MaxmemoryDelta pulumi.IntPtrInput `pulumi:"maxmemoryDelta"`
	// How Redis will select what to remove when `maxmemory` is reached. Defaults are shown below.
	MaxmemoryPolicy pulumi.StringPtrInput `pulumi:"maxmemoryPolicy"`
	// Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
	MaxmemoryReserved pulumi.IntPtrInput `pulumi:"maxmemoryReserved"`
	// Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. [Reference](https://redis.io/topics/notifications#configuration)
	NotifyKeyspaceEvents pulumi.StringPtrInput `pulumi:"notifyKeyspaceEvents"`
	// Is Backup Enabled? Only supported on Premium SKU's.
	RdbBackupEnabled pulumi.BoolPtrInput `pulumi:"rdbBackupEnabled"`
	// The Backup Frequency in Minutes. Only supported on Premium SKU's. Possible values are: `15`, `30`, `60`, `360`, `720` and `1440`.
	RdbBackupFrequency pulumi.IntPtrInput `pulumi:"rdbBackupFrequency"`
	// The maximum number of snapshots to create as a backup. Only supported for Premium SKU's.
	RdbBackupMaxSnapshotCount pulumi.IntPtrInput `pulumi:"rdbBackupMaxSnapshotCount"`
	// The Connection String to the Storage Account. Only supported for Premium SKU's. In the format: `DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}`.
	RdbStorageConnectionString pulumi.StringPtrInput `pulumi:"rdbStorageConnectionString"`
}

func (CacheRedisConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheRedisConfiguration)(nil)).Elem()
}

func (i CacheRedisConfigurationArgs) ToCacheRedisConfigurationOutput() CacheRedisConfigurationOutput {
	return i.ToCacheRedisConfigurationOutputWithContext(context.Background())
}

func (i CacheRedisConfigurationArgs) ToCacheRedisConfigurationOutputWithContext(ctx context.Context) CacheRedisConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheRedisConfigurationOutput)
}

func (i CacheRedisConfigurationArgs) ToCacheRedisConfigurationPtrOutput() CacheRedisConfigurationPtrOutput {
	return i.ToCacheRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i CacheRedisConfigurationArgs) ToCacheRedisConfigurationPtrOutputWithContext(ctx context.Context) CacheRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheRedisConfigurationOutput).ToCacheRedisConfigurationPtrOutputWithContext(ctx)
}

type CacheRedisConfigurationPtrInput interface {
	pulumi.Input

	ToCacheRedisConfigurationPtrOutput() CacheRedisConfigurationPtrOutput
	ToCacheRedisConfigurationPtrOutputWithContext(context.Context) CacheRedisConfigurationPtrOutput
}

type cacheRedisConfigurationPtrType CacheRedisConfigurationArgs

func CacheRedisConfigurationPtr(v *CacheRedisConfigurationArgs) CacheRedisConfigurationPtrInput {
	return (*cacheRedisConfigurationPtrType)(v)
}

func (*cacheRedisConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheRedisConfiguration)(nil)).Elem()
}

func (i *cacheRedisConfigurationPtrType) ToCacheRedisConfigurationPtrOutput() CacheRedisConfigurationPtrOutput {
	return i.ToCacheRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i *cacheRedisConfigurationPtrType) ToCacheRedisConfigurationPtrOutputWithContext(ctx context.Context) CacheRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheRedisConfigurationPtrOutput)
}

type CacheRedisConfigurationOutput struct{ *pulumi.OutputState }

func (CacheRedisConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheRedisConfiguration)(nil)).Elem()
}

func (o CacheRedisConfigurationOutput) ToCacheRedisConfigurationOutput() CacheRedisConfigurationOutput {
	return o
}

func (o CacheRedisConfigurationOutput) ToCacheRedisConfigurationOutputWithContext(ctx context.Context) CacheRedisConfigurationOutput {
	return o
}

func (o CacheRedisConfigurationOutput) ToCacheRedisConfigurationPtrOutput() CacheRedisConfigurationPtrOutput {
	return o.ToCacheRedisConfigurationPtrOutputWithContext(context.Background())
}

func (o CacheRedisConfigurationOutput) ToCacheRedisConfigurationPtrOutputWithContext(ctx context.Context) CacheRedisConfigurationPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *CacheRedisConfiguration {
		return &v
	}).(CacheRedisConfigurationPtrOutput)
}
func (o CacheRedisConfigurationOutput) AofBackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *bool { return v.AofBackupEnabled }).(pulumi.BoolPtrOutput)
}

func (o CacheRedisConfigurationOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.AofStorageConnectionString0 }).(pulumi.StringPtrOutput)
}

func (o CacheRedisConfigurationOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.AofStorageConnectionString1 }).(pulumi.StringPtrOutput)
}

// If set to `false`, the Redis instance will be accessible without authentication. Defaults to `true`.
func (o CacheRedisConfigurationOutput) EnableAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *bool { return v.EnableAuthentication }).(pulumi.BoolPtrOutput)
}

// Returns the max number of connected clients at the same time.
func (o CacheRedisConfigurationOutput) Maxclients() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.Maxclients }).(pulumi.IntPtrOutput)
}

// Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
func (o CacheRedisConfigurationOutput) MaxfragmentationmemoryReserved() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.MaxfragmentationmemoryReserved }).(pulumi.IntPtrOutput)
}

// The max-memory delta for this Redis instance. Defaults are shown below.
func (o CacheRedisConfigurationOutput) MaxmemoryDelta() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.MaxmemoryDelta }).(pulumi.IntPtrOutput)
}

// How Redis will select what to remove when `maxmemory` is reached. Defaults are shown below.
func (o CacheRedisConfigurationOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.MaxmemoryPolicy }).(pulumi.StringPtrOutput)
}

// Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
func (o CacheRedisConfigurationOutput) MaxmemoryReserved() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.MaxmemoryReserved }).(pulumi.IntPtrOutput)
}

// Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. [Reference](https://redis.io/topics/notifications#configuration)
func (o CacheRedisConfigurationOutput) NotifyKeyspaceEvents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.NotifyKeyspaceEvents }).(pulumi.StringPtrOutput)
}

// Is Backup Enabled? Only supported on Premium SKU's.
func (o CacheRedisConfigurationOutput) RdbBackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *bool { return v.RdbBackupEnabled }).(pulumi.BoolPtrOutput)
}

// The Backup Frequency in Minutes. Only supported on Premium SKU's. Possible values are: `15`, `30`, `60`, `360`, `720` and `1440`.
func (o CacheRedisConfigurationOutput) RdbBackupFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.RdbBackupFrequency }).(pulumi.IntPtrOutput)
}

// The maximum number of snapshots to create as a backup. Only supported for Premium SKU's.
func (o CacheRedisConfigurationOutput) RdbBackupMaxSnapshotCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.RdbBackupMaxSnapshotCount }).(pulumi.IntPtrOutput)
}

// The Connection String to the Storage Account. Only supported for Premium SKU's. In the format: `DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}`.
func (o CacheRedisConfigurationOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.RdbStorageConnectionString }).(pulumi.StringPtrOutput)
}

type CacheRedisConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CacheRedisConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheRedisConfiguration)(nil)).Elem()
}

func (o CacheRedisConfigurationPtrOutput) ToCacheRedisConfigurationPtrOutput() CacheRedisConfigurationPtrOutput {
	return o
}

func (o CacheRedisConfigurationPtrOutput) ToCacheRedisConfigurationPtrOutputWithContext(ctx context.Context) CacheRedisConfigurationPtrOutput {
	return o
}

func (o CacheRedisConfigurationPtrOutput) Elem() CacheRedisConfigurationOutput {
	return o.ApplyT(func(v *CacheRedisConfiguration) CacheRedisConfiguration { return *v }).(CacheRedisConfigurationOutput)
}

func (o CacheRedisConfigurationPtrOutput) AofBackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *bool { return v.AofBackupEnabled }).(pulumi.BoolPtrOutput)
}

func (o CacheRedisConfigurationPtrOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.AofStorageConnectionString0 }).(pulumi.StringPtrOutput)
}

func (o CacheRedisConfigurationPtrOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.AofStorageConnectionString1 }).(pulumi.StringPtrOutput)
}

// If set to `false`, the Redis instance will be accessible without authentication. Defaults to `true`.
func (o CacheRedisConfigurationPtrOutput) EnableAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *bool { return v.EnableAuthentication }).(pulumi.BoolPtrOutput)
}

// Returns the max number of connected clients at the same time.
func (o CacheRedisConfigurationPtrOutput) Maxclients() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.Maxclients }).(pulumi.IntPtrOutput)
}

// Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
func (o CacheRedisConfigurationPtrOutput) MaxfragmentationmemoryReserved() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.MaxfragmentationmemoryReserved }).(pulumi.IntPtrOutput)
}

// The max-memory delta for this Redis instance. Defaults are shown below.
func (o CacheRedisConfigurationPtrOutput) MaxmemoryDelta() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.MaxmemoryDelta }).(pulumi.IntPtrOutput)
}

// How Redis will select what to remove when `maxmemory` is reached. Defaults are shown below.
func (o CacheRedisConfigurationPtrOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.MaxmemoryPolicy }).(pulumi.StringPtrOutput)
}

// Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
func (o CacheRedisConfigurationPtrOutput) MaxmemoryReserved() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.MaxmemoryReserved }).(pulumi.IntPtrOutput)
}

// Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. [Reference](https://redis.io/topics/notifications#configuration)
func (o CacheRedisConfigurationPtrOutput) NotifyKeyspaceEvents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.NotifyKeyspaceEvents }).(pulumi.StringPtrOutput)
}

// Is Backup Enabled? Only supported on Premium SKU's.
func (o CacheRedisConfigurationPtrOutput) RdbBackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *bool { return v.RdbBackupEnabled }).(pulumi.BoolPtrOutput)
}

// The Backup Frequency in Minutes. Only supported on Premium SKU's. Possible values are: `15`, `30`, `60`, `360`, `720` and `1440`.
func (o CacheRedisConfigurationPtrOutput) RdbBackupFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.RdbBackupFrequency }).(pulumi.IntPtrOutput)
}

// The maximum number of snapshots to create as a backup. Only supported for Premium SKU's.
func (o CacheRedisConfigurationPtrOutput) RdbBackupMaxSnapshotCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *int { return v.RdbBackupMaxSnapshotCount }).(pulumi.IntPtrOutput)
}

// The Connection String to the Storage Account. Only supported for Premium SKU's. In the format: `DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}`.
func (o CacheRedisConfigurationPtrOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheRedisConfiguration) *string { return v.RdbStorageConnectionString }).(pulumi.StringPtrOutput)
}

type GetCachePatchSchedule struct {
	// the Weekday name for the patch item
	DayOfWeek string `pulumi:"dayOfWeek"`
	// The Start Hour for maintenance in UTC
	StartHourUtc int `pulumi:"startHourUtc"`
}

type GetCachePatchScheduleInput interface {
	pulumi.Input

	ToGetCachePatchScheduleOutput() GetCachePatchScheduleOutput
	ToGetCachePatchScheduleOutputWithContext(context.Context) GetCachePatchScheduleOutput
}

type GetCachePatchScheduleArgs struct {
	// the Weekday name for the patch item
	DayOfWeek pulumi.StringInput `pulumi:"dayOfWeek"`
	// The Start Hour for maintenance in UTC
	StartHourUtc pulumi.IntInput `pulumi:"startHourUtc"`
}

func (GetCachePatchScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCachePatchSchedule)(nil)).Elem()
}

func (i GetCachePatchScheduleArgs) ToGetCachePatchScheduleOutput() GetCachePatchScheduleOutput {
	return i.ToGetCachePatchScheduleOutputWithContext(context.Background())
}

func (i GetCachePatchScheduleArgs) ToGetCachePatchScheduleOutputWithContext(ctx context.Context) GetCachePatchScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCachePatchScheduleOutput)
}

type GetCachePatchScheduleArrayInput interface {
	pulumi.Input

	ToGetCachePatchScheduleArrayOutput() GetCachePatchScheduleArrayOutput
	ToGetCachePatchScheduleArrayOutputWithContext(context.Context) GetCachePatchScheduleArrayOutput
}

type GetCachePatchScheduleArray []GetCachePatchScheduleInput

func (GetCachePatchScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCachePatchSchedule)(nil)).Elem()
}

func (i GetCachePatchScheduleArray) ToGetCachePatchScheduleArrayOutput() GetCachePatchScheduleArrayOutput {
	return i.ToGetCachePatchScheduleArrayOutputWithContext(context.Background())
}

func (i GetCachePatchScheduleArray) ToGetCachePatchScheduleArrayOutputWithContext(ctx context.Context) GetCachePatchScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCachePatchScheduleArrayOutput)
}

type GetCachePatchScheduleOutput struct{ *pulumi.OutputState }

func (GetCachePatchScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCachePatchSchedule)(nil)).Elem()
}

func (o GetCachePatchScheduleOutput) ToGetCachePatchScheduleOutput() GetCachePatchScheduleOutput {
	return o
}

func (o GetCachePatchScheduleOutput) ToGetCachePatchScheduleOutputWithContext(ctx context.Context) GetCachePatchScheduleOutput {
	return o
}

// the Weekday name for the patch item
func (o GetCachePatchScheduleOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v GetCachePatchSchedule) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

// The Start Hour for maintenance in UTC
func (o GetCachePatchScheduleOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v GetCachePatchSchedule) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type GetCachePatchScheduleArrayOutput struct{ *pulumi.OutputState }

func (GetCachePatchScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCachePatchSchedule)(nil)).Elem()
}

func (o GetCachePatchScheduleArrayOutput) ToGetCachePatchScheduleArrayOutput() GetCachePatchScheduleArrayOutput {
	return o
}

func (o GetCachePatchScheduleArrayOutput) ToGetCachePatchScheduleArrayOutputWithContext(ctx context.Context) GetCachePatchScheduleArrayOutput {
	return o
}

func (o GetCachePatchScheduleArrayOutput) Index(i pulumi.IntInput) GetCachePatchScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCachePatchSchedule {
		return vs[0].([]GetCachePatchSchedule)[vs[1].(int)]
	}).(GetCachePatchScheduleOutput)
}

type GetCacheRedisConfiguration struct {
	AofBackupEnabled            bool   `pulumi:"aofBackupEnabled"`
	AofStorageConnectionString0 string `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1 string `pulumi:"aofStorageConnectionString1"`
	// Specifies if authentication is enabled
	EnableAuthentication bool `pulumi:"enableAuthentication"`
	Maxclients           int  `pulumi:"maxclients"`
	// Value in megabytes reserved to accommodate for memory fragmentation.
	MaxfragmentationmemoryReserved int `pulumi:"maxfragmentationmemoryReserved"`
	// The max-memory delta for this Redis instance.
	MaxmemoryDelta int `pulumi:"maxmemoryDelta"`
	// How Redis will select what to remove when `maxmemory` is reached.
	MaxmemoryPolicy string `pulumi:"maxmemoryPolicy"`
	// The value in megabytes reserved for non-cache usage e.g. failover
	MaxmemoryReserved    int    `pulumi:"maxmemoryReserved"`
	NotifyKeyspaceEvents string `pulumi:"notifyKeyspaceEvents"`
	// Is Backup Enabled? Only supported on Premium SKU's.
	RdbBackupEnabled bool `pulumi:"rdbBackupEnabled"`
	// The Backup Frequency in Minutes. Only supported on Premium SKU's.
	RdbBackupFrequency int `pulumi:"rdbBackupFrequency"`
	// The maximum number of snapshots that can be created as a backup.
	RdbBackupMaxSnapshotCount int `pulumi:"rdbBackupMaxSnapshotCount"`
	// The Connection String to the Storage Account. Only supported for Premium SKU's.
	RdbStorageConnectionString string `pulumi:"rdbStorageConnectionString"`
}

type GetCacheRedisConfigurationInput interface {
	pulumi.Input

	ToGetCacheRedisConfigurationOutput() GetCacheRedisConfigurationOutput
	ToGetCacheRedisConfigurationOutputWithContext(context.Context) GetCacheRedisConfigurationOutput
}

type GetCacheRedisConfigurationArgs struct {
	AofBackupEnabled            pulumi.BoolInput   `pulumi:"aofBackupEnabled"`
	AofStorageConnectionString0 pulumi.StringInput `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1 pulumi.StringInput `pulumi:"aofStorageConnectionString1"`
	// Specifies if authentication is enabled
	EnableAuthentication pulumi.BoolInput `pulumi:"enableAuthentication"`
	Maxclients           pulumi.IntInput  `pulumi:"maxclients"`
	// Value in megabytes reserved to accommodate for memory fragmentation.
	MaxfragmentationmemoryReserved pulumi.IntInput `pulumi:"maxfragmentationmemoryReserved"`
	// The max-memory delta for this Redis instance.
	MaxmemoryDelta pulumi.IntInput `pulumi:"maxmemoryDelta"`
	// How Redis will select what to remove when `maxmemory` is reached.
	MaxmemoryPolicy pulumi.StringInput `pulumi:"maxmemoryPolicy"`
	// The value in megabytes reserved for non-cache usage e.g. failover
	MaxmemoryReserved    pulumi.IntInput    `pulumi:"maxmemoryReserved"`
	NotifyKeyspaceEvents pulumi.StringInput `pulumi:"notifyKeyspaceEvents"`
	// Is Backup Enabled? Only supported on Premium SKU's.
	RdbBackupEnabled pulumi.BoolInput `pulumi:"rdbBackupEnabled"`
	// The Backup Frequency in Minutes. Only supported on Premium SKU's.
	RdbBackupFrequency pulumi.IntInput `pulumi:"rdbBackupFrequency"`
	// The maximum number of snapshots that can be created as a backup.
	RdbBackupMaxSnapshotCount pulumi.IntInput `pulumi:"rdbBackupMaxSnapshotCount"`
	// The Connection String to the Storage Account. Only supported for Premium SKU's.
	RdbStorageConnectionString pulumi.StringInput `pulumi:"rdbStorageConnectionString"`
}

func (GetCacheRedisConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCacheRedisConfiguration)(nil)).Elem()
}

func (i GetCacheRedisConfigurationArgs) ToGetCacheRedisConfigurationOutput() GetCacheRedisConfigurationOutput {
	return i.ToGetCacheRedisConfigurationOutputWithContext(context.Background())
}

func (i GetCacheRedisConfigurationArgs) ToGetCacheRedisConfigurationOutputWithContext(ctx context.Context) GetCacheRedisConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCacheRedisConfigurationOutput)
}

type GetCacheRedisConfigurationArrayInput interface {
	pulumi.Input

	ToGetCacheRedisConfigurationArrayOutput() GetCacheRedisConfigurationArrayOutput
	ToGetCacheRedisConfigurationArrayOutputWithContext(context.Context) GetCacheRedisConfigurationArrayOutput
}

type GetCacheRedisConfigurationArray []GetCacheRedisConfigurationInput

func (GetCacheRedisConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCacheRedisConfiguration)(nil)).Elem()
}

func (i GetCacheRedisConfigurationArray) ToGetCacheRedisConfigurationArrayOutput() GetCacheRedisConfigurationArrayOutput {
	return i.ToGetCacheRedisConfigurationArrayOutputWithContext(context.Background())
}

func (i GetCacheRedisConfigurationArray) ToGetCacheRedisConfigurationArrayOutputWithContext(ctx context.Context) GetCacheRedisConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCacheRedisConfigurationArrayOutput)
}

type GetCacheRedisConfigurationOutput struct{ *pulumi.OutputState }

func (GetCacheRedisConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCacheRedisConfiguration)(nil)).Elem()
}

func (o GetCacheRedisConfigurationOutput) ToGetCacheRedisConfigurationOutput() GetCacheRedisConfigurationOutput {
	return o
}

func (o GetCacheRedisConfigurationOutput) ToGetCacheRedisConfigurationOutputWithContext(ctx context.Context) GetCacheRedisConfigurationOutput {
	return o
}

func (o GetCacheRedisConfigurationOutput) AofBackupEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) bool { return v.AofBackupEnabled }).(pulumi.BoolOutput)
}

func (o GetCacheRedisConfigurationOutput) AofStorageConnectionString0() pulumi.StringOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) string { return v.AofStorageConnectionString0 }).(pulumi.StringOutput)
}

func (o GetCacheRedisConfigurationOutput) AofStorageConnectionString1() pulumi.StringOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) string { return v.AofStorageConnectionString1 }).(pulumi.StringOutput)
}

// Specifies if authentication is enabled
func (o GetCacheRedisConfigurationOutput) EnableAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) bool { return v.EnableAuthentication }).(pulumi.BoolOutput)
}

func (o GetCacheRedisConfigurationOutput) Maxclients() pulumi.IntOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) int { return v.Maxclients }).(pulumi.IntOutput)
}

// Value in megabytes reserved to accommodate for memory fragmentation.
func (o GetCacheRedisConfigurationOutput) MaxfragmentationmemoryReserved() pulumi.IntOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) int { return v.MaxfragmentationmemoryReserved }).(pulumi.IntOutput)
}

// The max-memory delta for this Redis instance.
func (o GetCacheRedisConfigurationOutput) MaxmemoryDelta() pulumi.IntOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) int { return v.MaxmemoryDelta }).(pulumi.IntOutput)
}

// How Redis will select what to remove when `maxmemory` is reached.
func (o GetCacheRedisConfigurationOutput) MaxmemoryPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) string { return v.MaxmemoryPolicy }).(pulumi.StringOutput)
}

// The value in megabytes reserved for non-cache usage e.g. failover
func (o GetCacheRedisConfigurationOutput) MaxmemoryReserved() pulumi.IntOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) int { return v.MaxmemoryReserved }).(pulumi.IntOutput)
}

func (o GetCacheRedisConfigurationOutput) NotifyKeyspaceEvents() pulumi.StringOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) string { return v.NotifyKeyspaceEvents }).(pulumi.StringOutput)
}

// Is Backup Enabled? Only supported on Premium SKU's.
func (o GetCacheRedisConfigurationOutput) RdbBackupEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) bool { return v.RdbBackupEnabled }).(pulumi.BoolOutput)
}

// The Backup Frequency in Minutes. Only supported on Premium SKU's.
func (o GetCacheRedisConfigurationOutput) RdbBackupFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) int { return v.RdbBackupFrequency }).(pulumi.IntOutput)
}

// The maximum number of snapshots that can be created as a backup.
func (o GetCacheRedisConfigurationOutput) RdbBackupMaxSnapshotCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) int { return v.RdbBackupMaxSnapshotCount }).(pulumi.IntOutput)
}

// The Connection String to the Storage Account. Only supported for Premium SKU's.
func (o GetCacheRedisConfigurationOutput) RdbStorageConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v GetCacheRedisConfiguration) string { return v.RdbStorageConnectionString }).(pulumi.StringOutput)
}

type GetCacheRedisConfigurationArrayOutput struct{ *pulumi.OutputState }

func (GetCacheRedisConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCacheRedisConfiguration)(nil)).Elem()
}

func (o GetCacheRedisConfigurationArrayOutput) ToGetCacheRedisConfigurationArrayOutput() GetCacheRedisConfigurationArrayOutput {
	return o
}

func (o GetCacheRedisConfigurationArrayOutput) ToGetCacheRedisConfigurationArrayOutputWithContext(ctx context.Context) GetCacheRedisConfigurationArrayOutput {
	return o
}

func (o GetCacheRedisConfigurationArrayOutput) Index(i pulumi.IntInput) GetCacheRedisConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCacheRedisConfiguration {
		return vs[0].([]GetCacheRedisConfiguration)[vs[1].(int)]
	}).(GetCacheRedisConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(CachePatchScheduleOutput{})
	pulumi.RegisterOutputType(CachePatchScheduleArrayOutput{})
	pulumi.RegisterOutputType(CacheRedisConfigurationOutput{})
	pulumi.RegisterOutputType(CacheRedisConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GetCachePatchScheduleOutput{})
	pulumi.RegisterOutputType(GetCachePatchScheduleArrayOutput{})
	pulumi.RegisterOutputType(GetCacheRedisConfigurationOutput{})
	pulumi.RegisterOutputType(GetCacheRedisConfigurationArrayOutput{})
}
