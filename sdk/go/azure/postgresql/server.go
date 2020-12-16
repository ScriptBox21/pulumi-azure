// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a PostgreSQL Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/postgresql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = postgresql.NewServer(ctx, "exampleServer", &postgresql.ServerArgs{
// 			Location:                     exampleResourceGroup.Location,
// 			ResourceGroupName:            exampleResourceGroup.Name,
// 			AdministratorLogin:           pulumi.String("psqladminun"),
// 			AdministratorLoginPassword:   pulumi.String("H@Sh1CoR3!"),
// 			SkuName:                      pulumi.String("GP_Gen5_4"),
// 			Version:                      pulumi.String("9.6"),
// 			StorageMb:                    pulumi.Int(640000),
// 			BackupRetentionDays:          pulumi.Int(7),
// 			GeoRedundantBackupEnabled:    pulumi.Bool(true),
// 			AutoGrowEnabled:              pulumi.Bool(true),
// 			PublicNetworkAccessEnabled:   pulumi.Bool(false),
// 			SslEnforcementEnabled:        pulumi.Bool(true),
// 			SslMinimalTlsVersionEnforced: pulumi.String("TLS1_2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// PostgreSQL Server's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:postgresql/server:Server server1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1
// ```
type Server struct {
	pulumi.CustomResourceState

	// The Administrator Login for the PostgreSQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the PostgreSQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword pulumi.StringPtrOutput `pulumi:"administratorLoginPassword"`
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled pulumi.BoolOutput `pulumi:"autoGrowEnabled"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntOutput `pulumi:"backupRetentionDays"`
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default.`
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// For creation modes other then default the source server ID to use.
	CreationSourceServerId pulumi.StringPtrOutput `pulumi:"creationSourceServerId"`
	// The FQDN of the PostgreSQL Server.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not support for the Basic tier. Changing this forces a new resource to be created.
	GeoRedundantBackupEnabled pulumi.BoolOutput `pulumi:"geoRedundantBackupEnabled"`
	// An `identity` block as defined below.
	Identity ServerIdentityPtrOutput `pulumi:"identity"`
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled pulumi.BoolPtrOutput `pulumi:"infrastructureEncryptionEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which to create the PostgreSQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// When `createMode` is `PointInTimeRestore` the point in time to restore from `creationSourceServerId`.
	RestorePointInTime pulumi.StringPtrOutput `pulumi:"restorePointInTime"`
	// Specifies the SKU Name for this PostgreSQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#sku).
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// Deprecated: this has been renamed to the boolean `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement pulumi.StringOutput `pulumi:"sslEnforcement"`
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled pulumi.BoolPtrOutput `pulumi:"sslEnforcementEnabled"`
	// The mimimun TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced pulumi.StringPtrOutput `pulumi:"sslMinimalTlsVersionEnforced"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#StorageProfile).
	StorageMb pulumi.IntOutput `pulumi:"storageMb"`
	// Deprecated: all storage_profile properties have been move to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile ServerStorageProfileOutput `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy ServerThreatDetectionPolicyPtrOutput `pulumi:"threatDetectionPolicy"`
	// Specifies the version of PostgreSQL to use. Valid values are `9.5`, `9.6`, `10`, `10.0`, and `11`. Changing this forces a new resource to be created.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource Server
	err := ctx.RegisterResource("azure:postgresql/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure:postgresql/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// The Administrator Login for the PostgreSQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the PostgreSQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled *bool `pulumi:"autoGrowEnabled"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default.`
	CreateMode *string `pulumi:"createMode"`
	// For creation modes other then default the source server ID to use.
	CreationSourceServerId *string `pulumi:"creationSourceServerId"`
	// The FQDN of the PostgreSQL Server.
	Fqdn *string `pulumi:"fqdn"`
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not support for the Basic tier. Changing this forces a new resource to be created.
	GeoRedundantBackupEnabled *bool `pulumi:"geoRedundantBackupEnabled"`
	// An `identity` block as defined below.
	Identity *ServerIdentity `pulumi:"identity"`
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled *bool `pulumi:"infrastructureEncryptionEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which to create the PostgreSQL Server. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// When `createMode` is `PointInTimeRestore` the point in time to restore from `creationSourceServerId`.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// Specifies the SKU Name for this PostgreSQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#sku).
	SkuName *string `pulumi:"skuName"`
	// Deprecated: this has been renamed to the boolean `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled *bool `pulumi:"sslEnforcementEnabled"`
	// The mimimun TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced *string `pulumi:"sslMinimalTlsVersionEnforced"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#StorageProfile).
	StorageMb *int `pulumi:"storageMb"`
	// Deprecated: all storage_profile properties have been move to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile *ServerStorageProfile `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy *ServerThreatDetectionPolicy `pulumi:"threatDetectionPolicy"`
	// Specifies the version of PostgreSQL to use. Valid values are `9.5`, `9.6`, `10`, `10.0`, and `11`. Changing this forces a new resource to be created.
	Version *string `pulumi:"version"`
}

type ServerState struct {
	// The Administrator Login for the PostgreSQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `administratorLogin` for the PostgreSQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword pulumi.StringPtrInput
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled pulumi.BoolPtrInput
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntPtrInput
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default.`
	CreateMode pulumi.StringPtrInput
	// For creation modes other then default the source server ID to use.
	CreationSourceServerId pulumi.StringPtrInput
	// The FQDN of the PostgreSQL Server.
	Fqdn pulumi.StringPtrInput
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not support for the Basic tier. Changing this forces a new resource to be created.
	GeoRedundantBackupEnabled pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity ServerIdentityPtrInput
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the PostgreSQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// When `createMode` is `PointInTimeRestore` the point in time to restore from `creationSourceServerId`.
	RestorePointInTime pulumi.StringPtrInput
	// Specifies the SKU Name for this PostgreSQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#sku).
	SkuName pulumi.StringPtrInput
	// Deprecated: this has been renamed to the boolean `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement pulumi.StringPtrInput
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled pulumi.BoolPtrInput
	// The mimimun TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced pulumi.StringPtrInput
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#StorageProfile).
	StorageMb pulumi.IntPtrInput
	// Deprecated: all storage_profile properties have been move to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile ServerStorageProfilePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy ServerThreatDetectionPolicyPtrInput
	// Specifies the version of PostgreSQL to use. Valid values are `9.5`, `9.6`, `10`, `10.0`, and `11`. Changing this forces a new resource to be created.
	Version pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The Administrator Login for the PostgreSQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Password associated with the `administratorLogin` for the PostgreSQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled *bool `pulumi:"autoGrowEnabled"`
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default.`
	CreateMode *string `pulumi:"createMode"`
	// For creation modes other then default the source server ID to use.
	CreationSourceServerId *string `pulumi:"creationSourceServerId"`
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not support for the Basic tier. Changing this forces a new resource to be created.
	GeoRedundantBackupEnabled *bool `pulumi:"geoRedundantBackupEnabled"`
	// An `identity` block as defined below.
	Identity *ServerIdentity `pulumi:"identity"`
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled *bool `pulumi:"infrastructureEncryptionEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which to create the PostgreSQL Server. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When `createMode` is `PointInTimeRestore` the point in time to restore from `creationSourceServerId`.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// Specifies the SKU Name for this PostgreSQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#sku).
	SkuName string `pulumi:"skuName"`
	// Deprecated: this has been renamed to the boolean `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement *string `pulumi:"sslEnforcement"`
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled *bool `pulumi:"sslEnforcementEnabled"`
	// The mimimun TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced *string `pulumi:"sslMinimalTlsVersionEnforced"`
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#StorageProfile).
	StorageMb *int `pulumi:"storageMb"`
	// Deprecated: all storage_profile properties have been move to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile *ServerStorageProfile `pulumi:"storageProfile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy *ServerThreatDetectionPolicy `pulumi:"threatDetectionPolicy"`
	// Specifies the version of PostgreSQL to use. Valid values are `9.5`, `9.6`, `10`, `10.0`, and `11`. Changing this forces a new resource to be created.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The Administrator Login for the PostgreSQL Server. Required when `createMode` is `Default`. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `administratorLogin` for the PostgreSQL Server. Required when `createMode` is `Default`.
	AdministratorLoginPassword pulumi.StringPtrInput
	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is `true`.
	AutoGrowEnabled pulumi.BoolPtrInput
	// Backup retention days for the server, supported values are between `7` and `35` days.
	BackupRetentionDays pulumi.IntPtrInput
	// The creation mode. Can be used to restore or replicate existing servers. Possible values are `Default`, `Replica`, `GeoRestore`, and `PointInTimeRestore`. Defaults to `Default.`
	CreateMode pulumi.StringPtrInput
	// For creation modes other then default the source server ID to use.
	CreationSourceServerId pulumi.StringPtrInput
	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not support for the Basic tier. Changing this forces a new resource to be created.
	GeoRedundantBackupEnabled pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity ServerIdentityPtrInput
	// Whether or not infrastructure is encrypted for this server. Defaults to `false`. Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the PostgreSQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// When `createMode` is `PointInTimeRestore` the point in time to restore from `creationSourceServerId`.
	RestorePointInTime pulumi.StringPtrInput
	// Specifies the SKU Name for this PostgreSQL Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#sku).
	SkuName pulumi.StringInput
	// Deprecated: this has been renamed to the boolean `ssl_enforcement_enabled` and will be removed in version 3.0 of the provider.
	SslEnforcement pulumi.StringPtrInput
	// Specifies if SSL should be enforced on connections. Possible values are `true` and `false`.
	SslEnforcementEnabled pulumi.BoolPtrInput
	// The mimimun TLS version to support on the sever. Possible values are `TLSEnforcementDisabled`, `TLS1_0`, `TLS1_1`, and `TLS1_2`. Defaults to `TLSEnforcementDisabled`.
	SslMinimalTlsVersionEnforced pulumi.StringPtrInput
	// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#StorageProfile).
	StorageMb pulumi.IntPtrInput
	// Deprecated: all storage_profile properties have been move to the top level. This block will be removed in version 3.0 of the provider.
	StorageProfile ServerStorageProfilePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy ServerThreatDetectionPolicyPtrInput
	// Specifies the version of PostgreSQL to use. Valid values are `9.5`, `9.6`, `10`, `10.0`, and `11`. Changing this forces a new resource to be created.
	Version pulumi.StringInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (Server) ElementType() reflect.Type {
	return reflect.TypeOf((*Server)(nil)).Elem()
}

func (i Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

type ServerOutput struct {
	*pulumi.OutputState
}

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerOutput)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
