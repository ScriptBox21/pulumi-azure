# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetScheduledQueryRulesAlertResult:
    """
    A collection of values returned by getScheduledQueryRulesAlert.
    """
    def __init__(__self__, actions=None, authorized_resource_ids=None, data_source_id=None, description=None, enabled=None, frequency=None, id=None, location=None, name=None, query=None, query_type=None, resource_group_name=None, severity=None, tags=None, throttling=None, time_window=None, triggers=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        __self__.actions = actions
        """
        An `action` block as defined below.
        """
        if authorized_resource_ids and not isinstance(authorized_resource_ids, list):
            raise TypeError("Expected argument 'authorized_resource_ids' to be a list")
        __self__.authorized_resource_ids = authorized_resource_ids
        """
        List of Resource IDs referred into query.
        """
        if data_source_id and not isinstance(data_source_id, str):
            raise TypeError("Expected argument 'data_source_id' to be a str")
        __self__.data_source_id = data_source_id
        """
        The resource URI over which log search query is to be run.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The description of the scheduled query rule.
        """
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        """
        Whether this scheduled query rule is enabled.
        """
        if frequency and not isinstance(frequency, float):
            raise TypeError("Expected argument 'frequency' to be a float")
        __self__.frequency = frequency
        """
        Frequency at which rule condition should be evaluated.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if query and not isinstance(query, str):
            raise TypeError("Expected argument 'query' to be a str")
        __self__.query = query
        """
        Log search query.
        """
        if query_type and not isinstance(query_type, str):
            raise TypeError("Expected argument 'query_type' to be a str")
        __self__.query_type = query_type
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if severity and not isinstance(severity, float):
            raise TypeError("Expected argument 'severity' to be a float")
        __self__.severity = severity
        """
        Severity of the alert.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if throttling and not isinstance(throttling, float):
            raise TypeError("Expected argument 'throttling' to be a float")
        __self__.throttling = throttling
        """
        Time for which alerts should be throttled or suppressed.
        """
        if time_window and not isinstance(time_window, float):
            raise TypeError("Expected argument 'time_window' to be a float")
        __self__.time_window = time_window
        """
        Time window for which data needs to be fetched for query.
        """
        if triggers and not isinstance(triggers, list):
            raise TypeError("Expected argument 'triggers' to be a list")
        __self__.triggers = triggers
        """
        A `trigger` block as defined below.
        """
class AwaitableGetScheduledQueryRulesAlertResult(GetScheduledQueryRulesAlertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScheduledQueryRulesAlertResult(
            actions=self.actions,
            authorized_resource_ids=self.authorized_resource_ids,
            data_source_id=self.data_source_id,
            description=self.description,
            enabled=self.enabled,
            frequency=self.frequency,
            id=self.id,
            location=self.location,
            name=self.name,
            query=self.query,
            query_type=self.query_type,
            resource_group_name=self.resource_group_name,
            severity=self.severity,
            tags=self.tags,
            throttling=self.throttling,
            time_window=self.time_window,
            triggers=self.triggers)

def get_scheduled_query_rules_alert(name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access the properties of an AlertingAction scheduled query rule.




    :param str name: Specifies the name of the scheduled query rule.
    :param str resource_group_name: Specifies the name of the resource group where the scheduled query rule is located.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:monitoring/getScheduledQueryRulesAlert:getScheduledQueryRulesAlert', __args__, opts=opts).value

    return AwaitableGetScheduledQueryRulesAlertResult(
        actions=__ret__.get('actions'),
        authorized_resource_ids=__ret__.get('authorizedResourceIds'),
        data_source_id=__ret__.get('dataSourceId'),
        description=__ret__.get('description'),
        enabled=__ret__.get('enabled'),
        frequency=__ret__.get('frequency'),
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        query=__ret__.get('query'),
        query_type=__ret__.get('queryType'),
        resource_group_name=__ret__.get('resourceGroupName'),
        severity=__ret__.get('severity'),
        tags=__ret__.get('tags'),
        throttling=__ret__.get('throttling'),
        time_window=__ret__.get('timeWindow'),
        triggers=__ret__.get('triggers'))
