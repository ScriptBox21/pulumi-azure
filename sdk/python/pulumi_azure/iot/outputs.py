# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'IoTHubEndpoint',
    'IoTHubEnrichment',
    'IoTHubFallbackRoute',
    'IoTHubFileUpload',
    'IoTHubIpFilterRule',
    'IoTHubRoute',
    'IoTHubSharedAccessPolicy',
    'IoTHubSku',
    'IotHubDpsLinkedHub',
    'IotHubDpsSku',
    'SecuritySolutionRecommendationsEnabled',
    'TimeSeriesInsightsGen2EnvironmentStorage',
    'TimeSeriesInsightsReferenceDataSetKeyProperty',
]

@pulumi.output_type
class IoTHubEndpoint(dict):
    def __init__(__self__, *,
                 connection_string: str,
                 name: str,
                 type: str,
                 batch_frequency_in_seconds: Optional[int] = None,
                 container_name: Optional[str] = None,
                 encoding: Optional[str] = None,
                 file_name_format: Optional[str] = None,
                 max_chunk_size_in_bytes: Optional[int] = None,
                 resource_group_name: Optional[str] = None):
        """
        :param str connection_string: The connection string for the endpoint.
        :param str name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        :param str type: The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
        :param int batch_frequency_in_seconds: Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param str container_name: The name of storage container in the storage account. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param str encoding: Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param str file_name_format: File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param int max_chunk_size_in_bytes: Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB). This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        :param str resource_group_name: The resource group in which the endpoint will be created.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if batch_frequency_in_seconds is not None:
            pulumi.set(__self__, "batch_frequency_in_seconds", batch_frequency_in_seconds)
        if container_name is not None:
            pulumi.set(__self__, "container_name", container_name)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if file_name_format is not None:
            pulumi.set(__self__, "file_name_format", file_name_format)
        if max_chunk_size_in_bytes is not None:
            pulumi.set(__self__, "max_chunk_size_in_bytes", max_chunk_size_in_bytes)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        The connection string for the endpoint.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="batchFrequencyInSeconds")
    def batch_frequency_in_seconds(self) -> Optional[int]:
        """
        Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        """
        return pulumi.get(self, "batch_frequency_in_seconds")

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> Optional[str]:
        """
        The name of storage container in the storage account. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        """
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter
    def encoding(self) -> Optional[str]:
        """
        Encoding that is used to serialize messages to blobs. Supported values are 'avro' and 'avrodeflate'. Default value is 'avro'. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="fileNameFormat")
    def file_name_format(self) -> Optional[str]:
        """
        File name format for the blob. Default format is ``{iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}``. All parameters are mandatory but can be reordered. This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        """
        return pulumi.get(self, "file_name_format")

    @property
    @pulumi.getter(name="maxChunkSizeInBytes")
    def max_chunk_size_in_bytes(self) -> Optional[int]:
        """
        Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB). This attribute is mandatory for endpoint type `AzureIotHub.StorageContainer`.
        """
        return pulumi.get(self, "max_chunk_size_in_bytes")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[str]:
        """
        The resource group in which the endpoint will be created.
        """
        return pulumi.get(self, "resource_group_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubEnrichment(dict):
    def __init__(__self__, *,
                 endpoint_names: Sequence[str],
                 key: str,
                 value: str):
        """
        :param Sequence[str] endpoint_names: The list of endpoints which will be enriched.
        :param str key: The key of the enrichment.
        :param str value: The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
        """
        pulumi.set(__self__, "endpoint_names", endpoint_names)
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="endpointNames")
    def endpoint_names(self) -> Sequence[str]:
        """
        The list of endpoints which will be enriched.
        """
        return pulumi.get(self, "endpoint_names")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key of the enrichment.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use `$iothubname`) or information from the device twin (ex: `$twin.tags.latitude`)
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubFallbackRoute(dict):
    def __init__(__self__, *,
                 condition: Optional[str] = None,
                 enabled: Optional[bool] = None,
                 endpoint_names: Optional[Sequence[str]] = None,
                 source: Optional[str] = None):
        """
        :param str condition: The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
        :param bool enabled: Used to specify whether the fallback route is enabled.
        :param Sequence[str] endpoint_names: The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
        :param str source: The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if endpoint_names is not None:
            pulumi.set(__self__, "endpoint_names", endpoint_names)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def condition(self) -> Optional[str]:
        """
        The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Used to specify whether the fallback route is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="endpointNames")
    def endpoint_names(self) -> Optional[Sequence[str]]:
        """
        The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
        """
        return pulumi.get(self, "endpoint_names")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        """
        The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
        """
        return pulumi.get(self, "source")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubFileUpload(dict):
    def __init__(__self__, *,
                 connection_string: str,
                 container_name: str,
                 default_ttl: Optional[str] = None,
                 lock_duration: Optional[str] = None,
                 max_delivery_count: Optional[int] = None,
                 notifications: Optional[bool] = None,
                 sas_ttl: Optional[str] = None):
        """
        :param str connection_string: The connection string for the Azure Storage account to which files are uploaded.
        :param str container_name: The name of the root container where you upload files. The container need not exist but should be creatable using the connection_string specified.
        :param str default_ttl: The period of time for which a file upload notification message is available to consume before it is expired by the IoT hub, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 48 hours, and evaluates to 'PT1H' by default.
        :param str lock_duration: The lock duration for the file upload notifications queue, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 5 and 300 seconds, and evaluates to 'PT1M' by default.
        :param int max_delivery_count: The number of times the IoT hub attempts to deliver a file upload notification message. It evaluates to 10 by default.
        :param bool notifications: Used to specify whether file notifications are sent to IoT Hub on upload. It evaluates to false by default.
        :param str sas_ttl: The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 24 hours, and evaluates to 'PT1H' by default.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "container_name", container_name)
        if default_ttl is not None:
            pulumi.set(__self__, "default_ttl", default_ttl)
        if lock_duration is not None:
            pulumi.set(__self__, "lock_duration", lock_duration)
        if max_delivery_count is not None:
            pulumi.set(__self__, "max_delivery_count", max_delivery_count)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)
        if sas_ttl is not None:
            pulumi.set(__self__, "sas_ttl", sas_ttl)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        The connection string for the Azure Storage account to which files are uploaded.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> str:
        """
        The name of the root container where you upload files. The container need not exist but should be creatable using the connection_string specified.
        """
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter(name="defaultTtl")
    def default_ttl(self) -> Optional[str]:
        """
        The period of time for which a file upload notification message is available to consume before it is expired by the IoT hub, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 48 hours, and evaluates to 'PT1H' by default.
        """
        return pulumi.get(self, "default_ttl")

    @property
    @pulumi.getter(name="lockDuration")
    def lock_duration(self) -> Optional[str]:
        """
        The lock duration for the file upload notifications queue, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 5 and 300 seconds, and evaluates to 'PT1M' by default.
        """
        return pulumi.get(self, "lock_duration")

    @property
    @pulumi.getter(name="maxDeliveryCount")
    def max_delivery_count(self) -> Optional[int]:
        """
        The number of times the IoT hub attempts to deliver a file upload notification message. It evaluates to 10 by default.
        """
        return pulumi.get(self, "max_delivery_count")

    @property
    @pulumi.getter
    def notifications(self) -> Optional[bool]:
        """
        Used to specify whether file notifications are sent to IoT Hub on upload. It evaluates to false by default.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter(name="sasTtl")
    def sas_ttl(self) -> Optional[str]:
        """
        The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an [ISO 8601 timespan duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). This value must be between 1 minute and 24 hours, and evaluates to 'PT1H' by default.
        """
        return pulumi.get(self, "sas_ttl")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubIpFilterRule(dict):
    def __init__(__self__, *,
                 action: str,
                 ip_mask: str,
                 name: str):
        """
        :param str action: The desired action for requests captured by this rule. Possible values are  `Accept`, `Reject`
        :param str ip_mask: The IP address range in CIDR notation for the rule.
        :param str name: The name of the filter.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "ip_mask", ip_mask)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        The desired action for requests captured by this rule. Possible values are  `Accept`, `Reject`
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="ipMask")
    def ip_mask(self) -> str:
        """
        The IP address range in CIDR notation for the rule.
        """
        return pulumi.get(self, "ip_mask")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubRoute(dict):
    def __init__(__self__, *,
                 enabled: bool,
                 endpoint_names: Sequence[str],
                 name: str,
                 source: str,
                 condition: Optional[str] = None):
        """
        :param bool enabled: Used to specify whether a route is enabled.
        :param Sequence[str] endpoint_names: The list of endpoints to which messages that satisfy the condition are routed.
        :param str name: The name of the route.
        :param str source: The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
        :param str condition: The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "endpoint_names", endpoint_names)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "source", source)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Used to specify whether a route is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="endpointNames")
    def endpoint_names(self) -> Sequence[str]:
        """
        The list of endpoints to which messages that satisfy the condition are routed.
        """
        return pulumi.get(self, "endpoint_names")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the route.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> str:
        """
        The source that the routing rule is to be applied to, such as `DeviceMessages`. Possible values include: `RoutingSourceInvalid`, `RoutingSourceDeviceMessages`, `RoutingSourceTwinChangeEvents`, `RoutingSourceDeviceLifecycleEvents`, `RoutingSourceDeviceJobLifecycleEvents`.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def condition(self) -> Optional[str]:
        """
        The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
        """
        return pulumi.get(self, "condition")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubSharedAccessPolicy(dict):
    def __init__(__self__, *,
                 key_name: Optional[str] = None,
                 permissions: Optional[str] = None,
                 primary_key: Optional[str] = None,
                 secondary_key: Optional[str] = None):
        """
        :param str key_name: The name of the shared access policy.
        :param str permissions: The permissions assigned to the shared access policy.
        :param str primary_key: The primary key.
        :param str secondary_key: The secondary key.
        """
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if primary_key is not None:
            pulumi.set(__self__, "primary_key", primary_key)
        if secondary_key is not None:
            pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[str]:
        """
        The name of the shared access policy.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def permissions(self) -> Optional[str]:
        """
        The permissions assigned to the shared access policy.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> Optional[str]:
        """
        The primary key.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> Optional[str]:
        """
        The secondary key.
        """
        return pulumi.get(self, "secondary_key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IoTHubSku(dict):
    def __init__(__self__, *,
                 capacity: int,
                 name: str):
        """
        :param int capacity: The number of provisioned IoT Hub units.
        :param str name: The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> int:
        """
        The number of provisioned IoT Hub units.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the sku. Possible values are `B1`, `B2`, `B3`, `F1`, `S1`, `S2`, and `S3`.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotHubDpsLinkedHub(dict):
    def __init__(__self__, *,
                 connection_string: str,
                 location: str,
                 allocation_weight: Optional[int] = None,
                 apply_allocation_policy: Optional[bool] = None,
                 hostname: Optional[str] = None):
        """
        :param str connection_string: The connection string to connect to the IoT Hub. Changing this forces a new resource.
        :param str location: The location of the IoT hub. Changing this forces a new resource.
        :param int allocation_weight: The weight applied to the IoT Hub. Defaults to 0.
        :param bool apply_allocation_policy: Determines whether to apply allocation policies to the IoT Hub. Defaults to false.
        :param str hostname: The IoT Hub hostname.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "location", location)
        if allocation_weight is not None:
            pulumi.set(__self__, "allocation_weight", allocation_weight)
        if apply_allocation_policy is not None:
            pulumi.set(__self__, "apply_allocation_policy", apply_allocation_policy)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        The connection string to connect to the IoT Hub. Changing this forces a new resource.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the IoT hub. Changing this forces a new resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="allocationWeight")
    def allocation_weight(self) -> Optional[int]:
        """
        The weight applied to the IoT Hub. Defaults to 0.
        """
        return pulumi.get(self, "allocation_weight")

    @property
    @pulumi.getter(name="applyAllocationPolicy")
    def apply_allocation_policy(self) -> Optional[bool]:
        """
        Determines whether to apply allocation policies to the IoT Hub. Defaults to false.
        """
        return pulumi.get(self, "apply_allocation_policy")

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        """
        The IoT Hub hostname.
        """
        return pulumi.get(self, "hostname")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotHubDpsSku(dict):
    def __init__(__self__, *,
                 capacity: int,
                 name: str):
        """
        :param int capacity: The number of provisioned IoT Device Provisioning Service units.
        :param str name: The name of the sku. Currently can only be set to `S1`.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> int:
        """
        The number of provisioned IoT Device Provisioning Service units.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the sku. Currently can only be set to `S1`.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SecuritySolutionRecommendationsEnabled(dict):
    def __init__(__self__, *,
                 acr_authentication: Optional[bool] = None,
                 agent_send_unutilized_msg: Optional[bool] = None,
                 baseline: Optional[bool] = None,
                 edge_hub_mem_optimize: Optional[bool] = None,
                 edge_logging_option: Optional[bool] = None,
                 inconsistent_module_settings: Optional[bool] = None,
                 install_agent: Optional[bool] = None,
                 ip_filter_deny_all: Optional[bool] = None,
                 ip_filter_permissive_rule: Optional[bool] = None,
                 open_ports: Optional[bool] = None,
                 permissive_firewall_policy: Optional[bool] = None,
                 permissive_input_firewall_rules: Optional[bool] = None,
                 permissive_output_firewall_rules: Optional[bool] = None,
                 privileged_docker_options: Optional[bool] = None,
                 shared_credentials: Optional[bool] = None,
                 vulnerable_tls_cipher_suite: Optional[bool] = None):
        """
        :param bool acr_authentication: Is Principal Authentication enabled for the ACR repository? Defaults to `true`.
        :param bool agent_send_unutilized_msg: Is Agent send underutilized messages enabled? Defaults to `true`.
        :param bool baseline: Is Security related system configuration issues identified? Defaults to `true`.
        :param bool edge_hub_mem_optimize: Is IoT Edge Hub memory optimized? Defaults to `true`.
        :param bool edge_logging_option: Is logging configured for IoT Edge module? Defaults to `true`.
        :param bool inconsistent_module_settings: Is inconsistent module settings enabled for SecurityGroup? Defaults to `true`.
        :param bool install_agent: is Azure IoT Security agent installed? Defaults to `true`.
        :param bool ip_filter_deny_all: Is Default IP filter policy denied? Defaults to `true`.
        :param bool ip_filter_permissive_rule: Is IP filter rule source allowable IP range too large? Defaults to `true`.
        :param bool open_ports: Is any ports open on the device? Defaults to `true`.
        :param bool permissive_firewall_policy: Does firewall policy exist which allow necessary communication to/from the device? Defaults to `true`.
        :param bool permissive_input_firewall_rules: Is only necessary addresses or ports are permitted in? Defaults to `true`.
        :param bool permissive_output_firewall_rules: Is only necessary addresses or ports are permitted out? Defaults to `true`.
        :param bool privileged_docker_options: Is high level permissions are needed for the module? Defaults to `true`.
        :param bool shared_credentials: Is any credentials shared among devices? Defaults to `true`.
        :param bool vulnerable_tls_cipher_suite: Does TLS cipher suite need to be updated? Defaults to `true`.
        """
        if acr_authentication is not None:
            pulumi.set(__self__, "acr_authentication", acr_authentication)
        if agent_send_unutilized_msg is not None:
            pulumi.set(__self__, "agent_send_unutilized_msg", agent_send_unutilized_msg)
        if baseline is not None:
            pulumi.set(__self__, "baseline", baseline)
        if edge_hub_mem_optimize is not None:
            pulumi.set(__self__, "edge_hub_mem_optimize", edge_hub_mem_optimize)
        if edge_logging_option is not None:
            pulumi.set(__self__, "edge_logging_option", edge_logging_option)
        if inconsistent_module_settings is not None:
            pulumi.set(__self__, "inconsistent_module_settings", inconsistent_module_settings)
        if install_agent is not None:
            pulumi.set(__self__, "install_agent", install_agent)
        if ip_filter_deny_all is not None:
            pulumi.set(__self__, "ip_filter_deny_all", ip_filter_deny_all)
        if ip_filter_permissive_rule is not None:
            pulumi.set(__self__, "ip_filter_permissive_rule", ip_filter_permissive_rule)
        if open_ports is not None:
            pulumi.set(__self__, "open_ports", open_ports)
        if permissive_firewall_policy is not None:
            pulumi.set(__self__, "permissive_firewall_policy", permissive_firewall_policy)
        if permissive_input_firewall_rules is not None:
            pulumi.set(__self__, "permissive_input_firewall_rules", permissive_input_firewall_rules)
        if permissive_output_firewall_rules is not None:
            pulumi.set(__self__, "permissive_output_firewall_rules", permissive_output_firewall_rules)
        if privileged_docker_options is not None:
            pulumi.set(__self__, "privileged_docker_options", privileged_docker_options)
        if shared_credentials is not None:
            pulumi.set(__self__, "shared_credentials", shared_credentials)
        if vulnerable_tls_cipher_suite is not None:
            pulumi.set(__self__, "vulnerable_tls_cipher_suite", vulnerable_tls_cipher_suite)

    @property
    @pulumi.getter(name="acrAuthentication")
    def acr_authentication(self) -> Optional[bool]:
        """
        Is Principal Authentication enabled for the ACR repository? Defaults to `true`.
        """
        return pulumi.get(self, "acr_authentication")

    @property
    @pulumi.getter(name="agentSendUnutilizedMsg")
    def agent_send_unutilized_msg(self) -> Optional[bool]:
        """
        Is Agent send underutilized messages enabled? Defaults to `true`.
        """
        return pulumi.get(self, "agent_send_unutilized_msg")

    @property
    @pulumi.getter
    def baseline(self) -> Optional[bool]:
        """
        Is Security related system configuration issues identified? Defaults to `true`.
        """
        return pulumi.get(self, "baseline")

    @property
    @pulumi.getter(name="edgeHubMemOptimize")
    def edge_hub_mem_optimize(self) -> Optional[bool]:
        """
        Is IoT Edge Hub memory optimized? Defaults to `true`.
        """
        return pulumi.get(self, "edge_hub_mem_optimize")

    @property
    @pulumi.getter(name="edgeLoggingOption")
    def edge_logging_option(self) -> Optional[bool]:
        """
        Is logging configured for IoT Edge module? Defaults to `true`.
        """
        return pulumi.get(self, "edge_logging_option")

    @property
    @pulumi.getter(name="inconsistentModuleSettings")
    def inconsistent_module_settings(self) -> Optional[bool]:
        """
        Is inconsistent module settings enabled for SecurityGroup? Defaults to `true`.
        """
        return pulumi.get(self, "inconsistent_module_settings")

    @property
    @pulumi.getter(name="installAgent")
    def install_agent(self) -> Optional[bool]:
        """
        is Azure IoT Security agent installed? Defaults to `true`.
        """
        return pulumi.get(self, "install_agent")

    @property
    @pulumi.getter(name="ipFilterDenyAll")
    def ip_filter_deny_all(self) -> Optional[bool]:
        """
        Is Default IP filter policy denied? Defaults to `true`.
        """
        return pulumi.get(self, "ip_filter_deny_all")

    @property
    @pulumi.getter(name="ipFilterPermissiveRule")
    def ip_filter_permissive_rule(self) -> Optional[bool]:
        """
        Is IP filter rule source allowable IP range too large? Defaults to `true`.
        """
        return pulumi.get(self, "ip_filter_permissive_rule")

    @property
    @pulumi.getter(name="openPorts")
    def open_ports(self) -> Optional[bool]:
        """
        Is any ports open on the device? Defaults to `true`.
        """
        return pulumi.get(self, "open_ports")

    @property
    @pulumi.getter(name="permissiveFirewallPolicy")
    def permissive_firewall_policy(self) -> Optional[bool]:
        """
        Does firewall policy exist which allow necessary communication to/from the device? Defaults to `true`.
        """
        return pulumi.get(self, "permissive_firewall_policy")

    @property
    @pulumi.getter(name="permissiveInputFirewallRules")
    def permissive_input_firewall_rules(self) -> Optional[bool]:
        """
        Is only necessary addresses or ports are permitted in? Defaults to `true`.
        """
        return pulumi.get(self, "permissive_input_firewall_rules")

    @property
    @pulumi.getter(name="permissiveOutputFirewallRules")
    def permissive_output_firewall_rules(self) -> Optional[bool]:
        """
        Is only necessary addresses or ports are permitted out? Defaults to `true`.
        """
        return pulumi.get(self, "permissive_output_firewall_rules")

    @property
    @pulumi.getter(name="privilegedDockerOptions")
    def privileged_docker_options(self) -> Optional[bool]:
        """
        Is high level permissions are needed for the module? Defaults to `true`.
        """
        return pulumi.get(self, "privileged_docker_options")

    @property
    @pulumi.getter(name="sharedCredentials")
    def shared_credentials(self) -> Optional[bool]:
        """
        Is any credentials shared among devices? Defaults to `true`.
        """
        return pulumi.get(self, "shared_credentials")

    @property
    @pulumi.getter(name="vulnerableTlsCipherSuite")
    def vulnerable_tls_cipher_suite(self) -> Optional[bool]:
        """
        Does TLS cipher suite need to be updated? Defaults to `true`.
        """
        return pulumi.get(self, "vulnerable_tls_cipher_suite")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TimeSeriesInsightsGen2EnvironmentStorage(dict):
    def __init__(__self__, *,
                 key: str,
                 name: str):
        """
        :param str key: Access key of storage account for Azure IoT Time Series Insights Gen2 Environment
        :param str name: Name of storage account for Azure IoT Time Series Insights Gen2 Environment
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Access key of storage account for Azure IoT Time Series Insights Gen2 Environment
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of storage account for Azure IoT Time Series Insights Gen2 Environment
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TimeSeriesInsightsReferenceDataSetKeyProperty(dict):
    def __init__(__self__, *,
                 name: str,
                 type: str):
        """
        :param str name: The name of the key property. Changing this forces a new resource to be created.
        :param str type: The data type of the key property. Valid values include `Bool`, `DateTime`, `Double`, `String`. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the key property. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The data type of the key property. Valid values include `Bool`, `DateTime`, `Double`, `String`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


