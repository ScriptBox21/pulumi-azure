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

__all__ = ['DiagnosticSetting']


class DiagnosticSetting(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eventhub_authorization_rule_id: Optional[pulumi.Input[str]] = None,
                 eventhub_name: Optional[pulumi.Input[str]] = None,
                 log_analytics_destination_type: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 logs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingLogArgs']]]]] = None,
                 metrics: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingMetricArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Diagnostic Setting for an existing Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = example_resource_group.name.apply(lambda name: azure.storage.get_account(name="examplestoracc",
            resource_group_name=name))
        example_key_vault = example_resource_group.name.apply(lambda name: azure.keyvault.get_key_vault(name="example-vault",
            resource_group_name=name))
        example_diagnostic_setting = azure.monitoring.DiagnosticSetting("exampleDiagnosticSetting",
            target_resource_id=example_key_vault.id,
            storage_account_id=example_account.id,
            logs=[azure.monitoring.DiagnosticSettingLogArgs(
                category="AuditEvent",
                enabled=False,
                retention_policy={
                    "enabled": False,
                },
            )],
            metrics=[azure.monitoring.DiagnosticSettingMetricArgs(
                category="AllMetrics",
                retention_policy={
                    "enabled": False,
                },
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_authorization_rule_id: Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_destination_type: When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingLogArgs']]]] logs: One or more `log` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingMetricArgs']]]] metrics: One or more `metric` blocks as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
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

            __props__['eventhub_authorization_rule_id'] = eventhub_authorization_rule_id
            __props__['eventhub_name'] = eventhub_name
            __props__['log_analytics_destination_type'] = log_analytics_destination_type
            __props__['log_analytics_workspace_id'] = log_analytics_workspace_id
            __props__['logs'] = logs
            __props__['metrics'] = metrics
            __props__['name'] = name
            __props__['storage_account_id'] = storage_account_id
            if target_resource_id is None:
                raise TypeError("Missing required property 'target_resource_id'")
            __props__['target_resource_id'] = target_resource_id
        super(DiagnosticSetting, __self__).__init__(
            'azure:monitoring/diagnosticSetting:DiagnosticSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eventhub_authorization_rule_id: Optional[pulumi.Input[str]] = None,
            eventhub_name: Optional[pulumi.Input[str]] = None,
            log_analytics_destination_type: Optional[pulumi.Input[str]] = None,
            log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
            logs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingLogArgs']]]]] = None,
            metrics: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingMetricArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None) -> 'DiagnosticSetting':
        """
        Get an existing DiagnosticSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_authorization_rule_id: Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[str] log_analytics_destination_type: When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingLogArgs']]]] logs: One or more `log` blocks as defined below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DiagnosticSettingMetricArgs']]]] metrics: One or more `metric` blocks as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["eventhub_authorization_rule_id"] = eventhub_authorization_rule_id
        __props__["eventhub_name"] = eventhub_name
        __props__["log_analytics_destination_type"] = log_analytics_destination_type
        __props__["log_analytics_workspace_id"] = log_analytics_workspace_id
        __props__["logs"] = logs
        __props__["metrics"] = metrics
        __props__["name"] = name
        __props__["storage_account_id"] = storage_account_id
        __props__["target_resource_id"] = target_resource_id
        return DiagnosticSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eventhubAuthorizationRuleId")
    def eventhub_authorization_rule_id(self) -> Optional[str]:
        """
        Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_authorization_rule_id")

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> Optional[str]:
        """
        Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "eventhub_name")

    @property
    @pulumi.getter(name="logAnalyticsDestinationType")
    def log_analytics_destination_type(self) -> Optional[str]:
        """
        When set to 'Dedicated' logs sent to a Log Analytics workspace will go into resource specific tables, instead of the legacy AzureDiagnostics table.
        """
        return pulumi.get(self, "log_analytics_destination_type")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> Optional[str]:
        """
        Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @property
    @pulumi.getter
    def logs(self) -> Optional[List['outputs.DiagnosticSettingLog']]:
        """
        One or more `log` blocks as defined below.
        """
        return pulumi.get(self, "logs")

    @property
    @pulumi.getter
    def metrics(self) -> Optional[List['outputs.DiagnosticSettingMetric']]:
        """
        One or more `metric` blocks as defined below.
        """
        return pulumi.get(self, "metrics")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[str]:
        """
        The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> str:
        """
        The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "target_resource_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

