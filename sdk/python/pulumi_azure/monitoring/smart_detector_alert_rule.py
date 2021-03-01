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

__all__ = ['SmartDetectorAlertRule']


class SmartDetectorAlertRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_group: Optional[pulumi.Input[pulumi.InputType['SmartDetectorAlertRuleActionGroupArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detector_type: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scope_resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 throttling_duration: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Monitor Smart Detector Alert Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_insights = azure.appinsights.Insights("exampleInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            application_type="web")
        example_action_group = azure.monitoring.ActionGroup("exampleActionGroup",
            resource_group_name=example_resource_group.name,
            short_name="exampleactiongroup")
        example_smart_detector_alert_rule = azure.monitoring.SmartDetectorAlertRule("exampleSmartDetectorAlertRule",
            resource_group_name=example_resource_group.name,
            severity="Sev0",
            scope_resource_ids=[example_insights.id],
            frequency="PT1M",
            detector_type="FailureAnomaliesDetector",
            action_group=azure.monitoring.SmartDetectorAlertRuleActionGroupArgs(
                ids=[azurerm_monitor_action_group["test"]["id"]],
            ))
        ```

        ## Import

        Monitor Smart Detector Alert Rule can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:monitoring/smartDetectorAlertRule:SmartDetectorAlertRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AlertsManagement/smartdetectoralertrules/rule1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SmartDetectorAlertRuleActionGroupArgs']] action_group: An `action_group` block as defined below.
        :param pulumi.Input[str] description: Specifies a description for the Smart Detector Alert Rule.
        :param pulumi.Input[str] detector_type: Specifies the Built-In Smart Detector type that this alert rule will use. Currently the only possible value is `FailureAnomaliesDetector`.
        :param pulumi.Input[bool] enabled: Is the Smart Detector Alert Rule enabled? Defaults to `true`.
        :param pulumi.Input[str] frequency: Specifies the frequency of this Smart Detector Alert Rule in ISO8601 format.
        :param pulumi.Input[str] name: Specifies the name of the Monitor Smart Detector Alert Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which the Monitor Smart Detector Alert Rule should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scope_resource_ids: Specifies the scopes of this Smart Detector Alert Rule.
        :param pulumi.Input[str] severity: Specifies the severity of this Smart Detector Alert Rule. Possible values are `Sev0`, `Sev1`, `Sev2`, `Sev3` or `Sev4`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] throttling_duration: Specifies the duration (in ISO8601 format) to wait before notifying on the alert rule again.
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

            if action_group is None and not opts.urn:
                raise TypeError("Missing required property 'action_group'")
            __props__['action_group'] = action_group
            __props__['description'] = description
            if detector_type is None and not opts.urn:
                raise TypeError("Missing required property 'detector_type'")
            __props__['detector_type'] = detector_type
            __props__['enabled'] = enabled
            if frequency is None and not opts.urn:
                raise TypeError("Missing required property 'frequency'")
            __props__['frequency'] = frequency
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if scope_resource_ids is None and not opts.urn:
                raise TypeError("Missing required property 'scope_resource_ids'")
            __props__['scope_resource_ids'] = scope_resource_ids
            if severity is None and not opts.urn:
                raise TypeError("Missing required property 'severity'")
            __props__['severity'] = severity
            __props__['tags'] = tags
            __props__['throttling_duration'] = throttling_duration
        super(SmartDetectorAlertRule, __self__).__init__(
            'azure:monitoring/smartDetectorAlertRule:SmartDetectorAlertRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action_group: Optional[pulumi.Input[pulumi.InputType['SmartDetectorAlertRuleActionGroupArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            detector_type: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            frequency: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            scope_resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            throttling_duration: Optional[pulumi.Input[str]] = None) -> 'SmartDetectorAlertRule':
        """
        Get an existing SmartDetectorAlertRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SmartDetectorAlertRuleActionGroupArgs']] action_group: An `action_group` block as defined below.
        :param pulumi.Input[str] description: Specifies a description for the Smart Detector Alert Rule.
        :param pulumi.Input[str] detector_type: Specifies the Built-In Smart Detector type that this alert rule will use. Currently the only possible value is `FailureAnomaliesDetector`.
        :param pulumi.Input[bool] enabled: Is the Smart Detector Alert Rule enabled? Defaults to `true`.
        :param pulumi.Input[str] frequency: Specifies the frequency of this Smart Detector Alert Rule in ISO8601 format.
        :param pulumi.Input[str] name: Specifies the name of the Monitor Smart Detector Alert Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which the Monitor Smart Detector Alert Rule should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scope_resource_ids: Specifies the scopes of this Smart Detector Alert Rule.
        :param pulumi.Input[str] severity: Specifies the severity of this Smart Detector Alert Rule. Possible values are `Sev0`, `Sev1`, `Sev2`, `Sev3` or `Sev4`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] throttling_duration: Specifies the duration (in ISO8601 format) to wait before notifying on the alert rule again.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action_group"] = action_group
        __props__["description"] = description
        __props__["detector_type"] = detector_type
        __props__["enabled"] = enabled
        __props__["frequency"] = frequency
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["scope_resource_ids"] = scope_resource_ids
        __props__["severity"] = severity
        __props__["tags"] = tags
        __props__["throttling_duration"] = throttling_duration
        return SmartDetectorAlertRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionGroup")
    def action_group(self) -> pulumi.Output['outputs.SmartDetectorAlertRuleActionGroup']:
        """
        An `action_group` block as defined below.
        """
        return pulumi.get(self, "action_group")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a description for the Smart Detector Alert Rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="detectorType")
    def detector_type(self) -> pulumi.Output[str]:
        """
        Specifies the Built-In Smart Detector type that this alert rule will use. Currently the only possible value is `FailureAnomaliesDetector`.
        """
        return pulumi.get(self, "detector_type")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the Smart Detector Alert Rule enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Output[str]:
        """
        Specifies the frequency of this Smart Detector Alert Rule in ISO8601 format.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Monitor Smart Detector Alert Rule. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the resource group in which the Monitor Smart Detector Alert Rule should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="scopeResourceIds")
    def scope_resource_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the scopes of this Smart Detector Alert Rule.
        """
        return pulumi.get(self, "scope_resource_ids")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        Specifies the severity of this Smart Detector Alert Rule. Possible values are `Sev0`, `Sev1`, `Sev2`, `Sev3` or `Sev4`.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="throttlingDuration")
    def throttling_duration(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the duration (in ISO8601 format) to wait before notifying on the alert rule again.
        """
        return pulumi.get(self, "throttling_duration")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

