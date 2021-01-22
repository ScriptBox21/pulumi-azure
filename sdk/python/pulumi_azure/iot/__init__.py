# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .consumer_group import *
from .dps_shared_access_policy import *
from .endpoint_eventhub import *
from .endpoint_servicebus_queue import *
from .endpoint_servicebus_topic import *
from .endpoint_storage_container import *
from .fallback_route import *
from .get_dps import *
from .get_dps_shared_access_policy import *
from .get_iot_hub import *
from .get_shared_access_policy import *
from .io_t_hub import *
from .iot_hub_certificate import *
from .iot_hub_dps import *
from .route import *
from .shared_access_policy import *
from .time_series_insights_access_policy import *
from .time_series_insights_gen2_environment import *
from .time_series_insights_reference_data_set import *
from .time_series_insights_standard_environment import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:iot/consumerGroup:ConsumerGroup":
                return ConsumerGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/dpsSharedAccessPolicy:DpsSharedAccessPolicy":
                return DpsSharedAccessPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/endpointEventhub:EndpointEventhub":
                return EndpointEventhub(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/endpointServicebusQueue:EndpointServicebusQueue":
                return EndpointServicebusQueue(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/endpointServicebusTopic:EndpointServicebusTopic":
                return EndpointServicebusTopic(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/endpointStorageContainer:EndpointStorageContainer":
                return EndpointStorageContainer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/fallbackRoute:FallbackRoute":
                return FallbackRoute(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/ioTHub:IoTHub":
                return IoTHub(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/iotHubCertificate:IotHubCertificate":
                return IotHubCertificate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/iotHubDps:IotHubDps":
                return IotHubDps(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/route:Route":
                return Route(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/sharedAccessPolicy:SharedAccessPolicy":
                return SharedAccessPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/timeSeriesInsightsAccessPolicy:TimeSeriesInsightsAccessPolicy":
                return TimeSeriesInsightsAccessPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment":
                return TimeSeriesInsightsGen2Environment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet":
                return TimeSeriesInsightsReferenceDataSet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment":
                return TimeSeriesInsightsStandardEnvironment(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "iot/consumerGroup", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/dpsSharedAccessPolicy", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/endpointEventhub", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/endpointServicebusQueue", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/endpointServicebusTopic", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/endpointStorageContainer", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/fallbackRoute", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/ioTHub", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/iotHubCertificate", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/iotHubDps", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/route", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/sharedAccessPolicy", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/timeSeriesInsightsAccessPolicy", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/timeSeriesInsightsGen2Environment", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/timeSeriesInsightsReferenceDataSet", _module_instance)
    pulumi.runtime.register_resource_module("azure", "iot/timeSeriesInsightsStandardEnvironment", _module_instance)

_register_module()
