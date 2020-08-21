# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ScheduledQueryRulesLog']


class ScheduledQueryRulesLog(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_resource_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['ScheduledQueryRulesLogCriteriaArgs']]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a LogToMetricAction Scheduled Query Rules resource within Azure Monitor.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ScheduledQueryRulesLogCriteriaArgs']] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] data_source_id: The resource uri over which log search query is to be run.
        :param pulumi.Input[str] description: The description of the scheduled query rule.
        :param pulumi.Input[bool] enabled: Whether this scheduled query rule is enabled.  Default is `true`.
        :param pulumi.Input[str] name: The name of the scheduled query rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the scheduled query rule instance.
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

            __props__['authorized_resource_ids'] = authorized_resource_ids
            if criteria is None:
                raise TypeError("Missing required property 'criteria'")
            __props__['criteria'] = criteria
            if data_source_id is None:
                raise TypeError("Missing required property 'data_source_id'")
            __props__['data_source_id'] = data_source_id
            __props__['description'] = description
            __props__['enabled'] = enabled
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(ScheduledQueryRulesLog, __self__).__init__(
            'azure:monitoring/scheduledQueryRulesLog:ScheduledQueryRulesLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized_resource_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            criteria: Optional[pulumi.Input[pulumi.InputType['ScheduledQueryRulesLogCriteriaArgs']]] = None,
            data_source_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ScheduledQueryRulesLog':
        """
        Get an existing ScheduledQueryRulesLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ScheduledQueryRulesLogCriteriaArgs']] criteria: A `criteria` block as defined below.
        :param pulumi.Input[str] data_source_id: The resource uri over which log search query is to be run.
        :param pulumi.Input[str] description: The description of the scheduled query rule.
        :param pulumi.Input[bool] enabled: Whether this scheduled query rule is enabled.  Default is `true`.
        :param pulumi.Input[str] name: The name of the scheduled query rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the scheduled query rule instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorized_resource_ids"] = authorized_resource_ids
        __props__["criteria"] = criteria
        __props__["data_source_id"] = data_source_id
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return ScheduledQueryRulesLog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizedResourceIds")
    def authorized_resource_ids(self) -> Optional[List[str]]:
        return pulumi.get(self, "authorized_resource_ids")

    @property
    @pulumi.getter
    def criteria(self) -> 'outputs.ScheduledQueryRulesLogCriteria':
        """
        A `criteria` block as defined below.
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> str:
        """
        The resource uri over which log search query is to be run.
        """
        return pulumi.get(self, "data_source_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the scheduled query rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether this scheduled query rule is enabled.  Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the scheduled query rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to create the scheduled query rule instance.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

