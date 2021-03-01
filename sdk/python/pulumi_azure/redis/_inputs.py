# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'CachePatchScheduleArgs',
    'CacheRedisConfigurationArgs',
    'EnterpriseDatabaseModuleArgs',
]

@pulumi.input_type
class CachePatchScheduleArgs:
    def __init__(__self__, *,
                 day_of_week: pulumi.Input[str],
                 start_hour_utc: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "day_of_week", day_of_week)
        if start_hour_utc is not None:
            pulumi.set(__self__, "start_hour_utc", start_hour_utc)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> pulumi.Input[str]:
        return pulumi.get(self, "day_of_week")

    @day_of_week.setter
    def day_of_week(self, value: pulumi.Input[str]):
        pulumi.set(self, "day_of_week", value)

    @property
    @pulumi.getter(name="startHourUtc")
    def start_hour_utc(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_hour_utc")

    @start_hour_utc.setter
    def start_hour_utc(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_hour_utc", value)


@pulumi.input_type
class CacheRedisConfigurationArgs:
    def __init__(__self__, *,
                 aof_backup_enabled: Optional[pulumi.Input[bool]] = None,
                 aof_storage_connection_string0: Optional[pulumi.Input[str]] = None,
                 aof_storage_connection_string1: Optional[pulumi.Input[str]] = None,
                 enable_authentication: Optional[pulumi.Input[bool]] = None,
                 maxclients: Optional[pulumi.Input[int]] = None,
                 maxfragmentationmemory_reserved: Optional[pulumi.Input[int]] = None,
                 maxmemory_delta: Optional[pulumi.Input[int]] = None,
                 maxmemory_policy: Optional[pulumi.Input[str]] = None,
                 maxmemory_reserved: Optional[pulumi.Input[int]] = None,
                 notify_keyspace_events: Optional[pulumi.Input[str]] = None,
                 rdb_backup_enabled: Optional[pulumi.Input[bool]] = None,
                 rdb_backup_frequency: Optional[pulumi.Input[int]] = None,
                 rdb_backup_max_snapshot_count: Optional[pulumi.Input[int]] = None,
                 rdb_storage_connection_string: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enable_authentication: If set to `false`, the Redis instance will be accessible without authentication. Defaults to `true`.
        :param pulumi.Input[int] maxclients: Returns the max number of connected clients at the same time.
        :param pulumi.Input[int] maxfragmentationmemory_reserved: Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
        :param pulumi.Input[int] maxmemory_delta: The max-memory delta for this Redis instance. Defaults are shown below.
        :param pulumi.Input[str] maxmemory_policy: How Redis will select what to remove when `maxmemory` is reached. Defaults are shown below.
        :param pulumi.Input[int] maxmemory_reserved: Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
        :param pulumi.Input[str] notify_keyspace_events: Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. [Reference](https://redis.io/topics/notifications#configuration)
        :param pulumi.Input[bool] rdb_backup_enabled: Is Backup Enabled? Only supported on Premium SKU's.
        :param pulumi.Input[int] rdb_backup_frequency: The Backup Frequency in Minutes. Only supported on Premium SKU's. Possible values are: `15`, `30`, `60`, `360`, `720` and `1440`.
        :param pulumi.Input[int] rdb_backup_max_snapshot_count: The maximum number of snapshots to create as a backup. Only supported for Premium SKU's.
        :param pulumi.Input[str] rdb_storage_connection_string: The Connection String to the Storage Account. Only supported for Premium SKU's. In the format: `DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}`.
        """
        if aof_backup_enabled is not None:
            pulumi.set(__self__, "aof_backup_enabled", aof_backup_enabled)
        if aof_storage_connection_string0 is not None:
            pulumi.set(__self__, "aof_storage_connection_string0", aof_storage_connection_string0)
        if aof_storage_connection_string1 is not None:
            pulumi.set(__self__, "aof_storage_connection_string1", aof_storage_connection_string1)
        if enable_authentication is not None:
            pulumi.set(__self__, "enable_authentication", enable_authentication)
        if maxclients is not None:
            pulumi.set(__self__, "maxclients", maxclients)
        if maxfragmentationmemory_reserved is not None:
            pulumi.set(__self__, "maxfragmentationmemory_reserved", maxfragmentationmemory_reserved)
        if maxmemory_delta is not None:
            pulumi.set(__self__, "maxmemory_delta", maxmemory_delta)
        if maxmemory_policy is not None:
            pulumi.set(__self__, "maxmemory_policy", maxmemory_policy)
        if maxmemory_reserved is not None:
            pulumi.set(__self__, "maxmemory_reserved", maxmemory_reserved)
        if notify_keyspace_events is not None:
            pulumi.set(__self__, "notify_keyspace_events", notify_keyspace_events)
        if rdb_backup_enabled is not None:
            pulumi.set(__self__, "rdb_backup_enabled", rdb_backup_enabled)
        if rdb_backup_frequency is not None:
            pulumi.set(__self__, "rdb_backup_frequency", rdb_backup_frequency)
        if rdb_backup_max_snapshot_count is not None:
            pulumi.set(__self__, "rdb_backup_max_snapshot_count", rdb_backup_max_snapshot_count)
        if rdb_storage_connection_string is not None:
            pulumi.set(__self__, "rdb_storage_connection_string", rdb_storage_connection_string)

    @property
    @pulumi.getter(name="aofBackupEnabled")
    def aof_backup_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "aof_backup_enabled")

    @aof_backup_enabled.setter
    def aof_backup_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "aof_backup_enabled", value)

    @property
    @pulumi.getter(name="aofStorageConnectionString0")
    def aof_storage_connection_string0(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "aof_storage_connection_string0")

    @aof_storage_connection_string0.setter
    def aof_storage_connection_string0(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aof_storage_connection_string0", value)

    @property
    @pulumi.getter(name="aofStorageConnectionString1")
    def aof_storage_connection_string1(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "aof_storage_connection_string1")

    @aof_storage_connection_string1.setter
    def aof_storage_connection_string1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aof_storage_connection_string1", value)

    @property
    @pulumi.getter(name="enableAuthentication")
    def enable_authentication(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `false`, the Redis instance will be accessible without authentication. Defaults to `true`.
        """
        return pulumi.get(self, "enable_authentication")

    @enable_authentication.setter
    def enable_authentication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_authentication", value)

    @property
    @pulumi.getter
    def maxclients(self) -> Optional[pulumi.Input[int]]:
        """
        Returns the max number of connected clients at the same time.
        """
        return pulumi.get(self, "maxclients")

    @maxclients.setter
    def maxclients(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maxclients", value)

    @property
    @pulumi.getter(name="maxfragmentationmemoryReserved")
    def maxfragmentationmemory_reserved(self) -> Optional[pulumi.Input[int]]:
        """
        Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
        """
        return pulumi.get(self, "maxfragmentationmemory_reserved")

    @maxfragmentationmemory_reserved.setter
    def maxfragmentationmemory_reserved(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maxfragmentationmemory_reserved", value)

    @property
    @pulumi.getter(name="maxmemoryDelta")
    def maxmemory_delta(self) -> Optional[pulumi.Input[int]]:
        """
        The max-memory delta for this Redis instance. Defaults are shown below.
        """
        return pulumi.get(self, "maxmemory_delta")

    @maxmemory_delta.setter
    def maxmemory_delta(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maxmemory_delta", value)

    @property
    @pulumi.getter(name="maxmemoryPolicy")
    def maxmemory_policy(self) -> Optional[pulumi.Input[str]]:
        """
        How Redis will select what to remove when `maxmemory` is reached. Defaults are shown below.
        """
        return pulumi.get(self, "maxmemory_policy")

    @maxmemory_policy.setter
    def maxmemory_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maxmemory_policy", value)

    @property
    @pulumi.getter(name="maxmemoryReserved")
    def maxmemory_reserved(self) -> Optional[pulumi.Input[int]]:
        """
        Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
        """
        return pulumi.get(self, "maxmemory_reserved")

    @maxmemory_reserved.setter
    def maxmemory_reserved(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maxmemory_reserved", value)

    @property
    @pulumi.getter(name="notifyKeyspaceEvents")
    def notify_keyspace_events(self) -> Optional[pulumi.Input[str]]:
        """
        Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. [Reference](https://redis.io/topics/notifications#configuration)
        """
        return pulumi.get(self, "notify_keyspace_events")

    @notify_keyspace_events.setter
    def notify_keyspace_events(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_keyspace_events", value)

    @property
    @pulumi.getter(name="rdbBackupEnabled")
    def rdb_backup_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is Backup Enabled? Only supported on Premium SKU's.
        """
        return pulumi.get(self, "rdb_backup_enabled")

    @rdb_backup_enabled.setter
    def rdb_backup_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "rdb_backup_enabled", value)

    @property
    @pulumi.getter(name="rdbBackupFrequency")
    def rdb_backup_frequency(self) -> Optional[pulumi.Input[int]]:
        """
        The Backup Frequency in Minutes. Only supported on Premium SKU's. Possible values are: `15`, `30`, `60`, `360`, `720` and `1440`.
        """
        return pulumi.get(self, "rdb_backup_frequency")

    @rdb_backup_frequency.setter
    def rdb_backup_frequency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rdb_backup_frequency", value)

    @property
    @pulumi.getter(name="rdbBackupMaxSnapshotCount")
    def rdb_backup_max_snapshot_count(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of snapshots to create as a backup. Only supported for Premium SKU's.
        """
        return pulumi.get(self, "rdb_backup_max_snapshot_count")

    @rdb_backup_max_snapshot_count.setter
    def rdb_backup_max_snapshot_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rdb_backup_max_snapshot_count", value)

    @property
    @pulumi.getter(name="rdbStorageConnectionString")
    def rdb_storage_connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The Connection String to the Storage Account. Only supported for Premium SKU's. In the format: `DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}`.
        """
        return pulumi.get(self, "rdb_storage_connection_string")

    @rdb_storage_connection_string.setter
    def rdb_storage_connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rdb_storage_connection_string", value)


@pulumi.input_type
class EnterpriseDatabaseModuleArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 args: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name which should be used for this module. Possible values are `RediSearch`, `RedisBloom` and `RedisTimeSeries`. Changing this forces a new Redis Enterprise Database to be created.
        :param pulumi.Input[str] args: Configuration options for the module (e.g. `ERROR_RATE 0.00 INITIAL_SIZE 400`).
        """
        pulumi.set(__self__, "name", name)
        if args is not None:
            pulumi.set(__self__, "args", args)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name which should be used for this module. Possible values are `RediSearch`, `RedisBloom` and `RedisTimeSeries`. Changing this forces a new Redis Enterprise Database to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def args(self) -> Optional[pulumi.Input[str]]:
        """
        Configuration options for the module (e.g. `ERROR_RATE 0.00 INITIAL_SIZE 400`).
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


