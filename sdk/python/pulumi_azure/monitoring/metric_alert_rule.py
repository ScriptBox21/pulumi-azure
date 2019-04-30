# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class MetricAlertRule(pulumi.CustomResource):
    aggregation: pulumi.Output[str]
    """
    Defines how the metric data is combined over time. Possible values are `Average`, `Minimum`, `Maximum`, `Total`, and `Last`.
    """
    description: pulumi.Output[str]
    """
    A verbose description of the alert rule that will be included in the alert email.
    """
    email_action: pulumi.Output[dict]
    """
    A `email_action` block as defined below.
    """
    enabled: pulumi.Output[bool]
    """
    If `true`, the alert rule is enabled. Defaults to `true`.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    metric_name: pulumi.Output[str]
    """
    The metric that defines what the rule monitors.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the alert rule. Changing this forces a new resource to be created.
    """
    operator: pulumi.Output[str]
    """
    The operator used to compare the metric data and the threshold. Possible values are `GreaterThan`, `GreaterThanOrEqual`, `LessThan`, and `LessThanOrEqual`.
    """
    period: pulumi.Output[str]
    """
    The period of time formatted in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations) that is used to monitor the alert activity based on the threshold. The period must be between 5 minutes and 1 day.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the alert rule. Changing this forces a new resource to be created.
    """
    resource_id: pulumi.Output[str]
    """
    The ID of the resource monitored by the alert rule.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
    """
    threshold: pulumi.Output[float]
    """
    The threshold value that activates the alert.
    """
    webhook_action: pulumi.Output[dict]
    """
    A `webhook_action` block as defined below.
    """
    def __init__(__self__, resource_name, opts=None, aggregation=None, description=None, email_action=None, enabled=None, location=None, metric_name=None, name=None, operator=None, period=None, resource_group_name=None, resource_id=None, tags=None, threshold=None, webhook_action=None, __name__=None, __opts__=None):
        """
        Manages a [metric-based alert rule](https://docs.microsoft.com/en-us/azure/monitoring-and-diagnostics/monitor-quick-resource-metric-alert-portal) in Azure Monitor.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aggregation: Defines how the metric data is combined over time. Possible values are `Average`, `Minimum`, `Maximum`, `Total`, and `Last`.
        :param pulumi.Input[str] description: A verbose description of the alert rule that will be included in the alert email.
        :param pulumi.Input[dict] email_action: A `email_action` block as defined below.
        :param pulumi.Input[bool] enabled: If `true`, the alert rule is enabled. Defaults to `true`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metric_name: The metric that defines what the rule monitors.
        :param pulumi.Input[str] name: Specifies the name of the alert rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] operator: The operator used to compare the metric data and the threshold. Possible values are `GreaterThan`, `GreaterThanOrEqual`, `LessThan`, and `LessThanOrEqual`.
        :param pulumi.Input[str] period: The period of time formatted in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations) that is used to monitor the alert activity based on the threshold. The period must be between 5 minutes and 1 day.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the alert rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_id: The ID of the resource monitored by the alert rule.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        :param pulumi.Input[float] threshold: The threshold value that activates the alert.
        :param pulumi.Input[dict] webhook_action: A `webhook_action` block as defined below.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if aggregation is None:
            raise TypeError("Missing required property 'aggregation'")
        __props__['aggregation'] = aggregation

        __props__['description'] = description

        __props__['email_action'] = email_action

        __props__['enabled'] = enabled

        __props__['location'] = location

        if metric_name is None:
            raise TypeError("Missing required property 'metric_name'")
        __props__['metric_name'] = metric_name

        __props__['name'] = name

        if operator is None:
            raise TypeError("Missing required property 'operator'")
        __props__['operator'] = operator

        if period is None:
            raise TypeError("Missing required property 'period'")
        __props__['period'] = period

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        if resource_id is None:
            raise TypeError("Missing required property 'resource_id'")
        __props__['resource_id'] = resource_id

        __props__['tags'] = tags

        if threshold is None:
            raise TypeError("Missing required property 'threshold'")
        __props__['threshold'] = threshold

        __props__['webhook_action'] = webhook_action

        super(MetricAlertRule, __self__).__init__(
            'azure:monitoring/metricAlertRule:MetricAlertRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

