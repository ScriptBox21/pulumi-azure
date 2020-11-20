# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['NetworkConnectionMonitor']


class NetworkConnectionMonitor(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_start: Optional[pulumi.Input[bool]] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorEndpointArgs']]]]] = None,
                 interval_in_seconds: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_watcher_id: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 output_workspace_resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 test_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestConfigurationArgs']]]]] = None,
                 test_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestGroupArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Network Connection Monitors can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/networkConnectionMonitor:NetworkConnectionMonitor example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkWatchers/watcher1/connectionMonitors/connectionMonitor1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorEndpointArgs']]]] endpoints: A `endpoint` block as defined below.
        :param pulumi.Input[str] location: The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_id: The ID of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] notes: The description of the Network Connection Monitor.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] output_workspace_resource_ids: A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Network Connection Monitor.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestConfigurationArgs']]]] test_configurations: A `test_configuration` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestGroupArgs']]]] test_groups: A `test_group` block as defined below.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if auto_start is not None:
                warnings.warn("""The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.""", DeprecationWarning)
                pulumi.log.warn("auto_start is deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.")
            __props__['auto_start'] = auto_start
            if destination is not None:
                warnings.warn("""The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.""", DeprecationWarning)
                pulumi.log.warn("destination is deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.")
            __props__['destination'] = destination
            if endpoints is None:
                raise TypeError("Missing required property 'endpoints'")
            __props__['endpoints'] = endpoints
            if interval_in_seconds is not None:
                warnings.warn("""The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.""", DeprecationWarning)
                pulumi.log.warn("interval_in_seconds is deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.")
            __props__['interval_in_seconds'] = interval_in_seconds
            __props__['location'] = location
            __props__['name'] = name
            if network_watcher_id is None:
                raise TypeError("Missing required property 'network_watcher_id'")
            __props__['network_watcher_id'] = network_watcher_id
            __props__['notes'] = notes
            __props__['output_workspace_resource_ids'] = output_workspace_resource_ids
            if source is not None:
                warnings.warn("""The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.""", DeprecationWarning)
                pulumi.log.warn("source is deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.")
            __props__['source'] = source
            __props__['tags'] = tags
            if test_configurations is None:
                raise TypeError("Missing required property 'test_configurations'")
            __props__['test_configurations'] = test_configurations
            if test_groups is None:
                raise TypeError("Missing required property 'test_groups'")
            __props__['test_groups'] = test_groups
        super(NetworkConnectionMonitor, __self__).__init__(
            'azure:network/networkConnectionMonitor:NetworkConnectionMonitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_start: Optional[pulumi.Input[bool]] = None,
            destination: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorDestinationArgs']]] = None,
            endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorEndpointArgs']]]]] = None,
            interval_in_seconds: Optional[pulumi.Input[int]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_watcher_id: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            output_workspace_resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            source: Optional[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorSourceArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            test_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestConfigurationArgs']]]]] = None,
            test_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestGroupArgs']]]]] = None) -> 'NetworkConnectionMonitor':
        """
        Get an existing NetworkConnectionMonitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorEndpointArgs']]]] endpoints: A `endpoint` block as defined below.
        :param pulumi.Input[str] location: The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
        :param pulumi.Input[str] network_watcher_id: The ID of the Network Watcher. Changing this forces a new resource to be created.
        :param pulumi.Input[str] notes: The description of the Network Connection Monitor.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] output_workspace_resource_ids: A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Network Connection Monitor.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestConfigurationArgs']]]] test_configurations: A `test_configuration` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConnectionMonitorTestGroupArgs']]]] test_groups: A `test_group` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_start"] = auto_start
        __props__["destination"] = destination
        __props__["endpoints"] = endpoints
        __props__["interval_in_seconds"] = interval_in_seconds
        __props__["location"] = location
        __props__["name"] = name
        __props__["network_watcher_id"] = network_watcher_id
        __props__["notes"] = notes
        __props__["output_workspace_resource_ids"] = output_workspace_resource_ids
        __props__["source"] = source
        __props__["tags"] = tags
        __props__["test_configurations"] = test_configurations
        __props__["test_groups"] = test_groups
        return NetworkConnectionMonitor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoStart")
    def auto_start(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "auto_start")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output['outputs.NetworkConnectionMonitorDestination']:
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Sequence['outputs.NetworkConnectionMonitorEndpoint']]:
        """
        A `endpoint` block as defined below.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter(name="intervalInSeconds")
    def interval_in_seconds(self) -> pulumi.Output[int]:
        return pulumi.get(self, "interval_in_seconds")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkWatcherId")
    def network_watcher_id(self) -> pulumi.Output[str]:
        """
        The ID of the Network Watcher. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_watcher_id")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Network Connection Monitor.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="outputWorkspaceResourceIds")
    def output_workspace_resource_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
        """
        return pulumi.get(self, "output_workspace_resource_ids")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output['outputs.NetworkConnectionMonitorSource']:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Network Connection Monitor.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="testConfigurations")
    def test_configurations(self) -> pulumi.Output[Sequence['outputs.NetworkConnectionMonitorTestConfiguration']]:
        """
        A `test_configuration` block as defined below.
        """
        return pulumi.get(self, "test_configurations")

    @property
    @pulumi.getter(name="testGroups")
    def test_groups(self) -> pulumi.Output[Sequence['outputs.NetworkConnectionMonitorTestGroup']]:
        """
        A `test_group` block as defined below.
        """
        return pulumi.get(self, "test_groups")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

