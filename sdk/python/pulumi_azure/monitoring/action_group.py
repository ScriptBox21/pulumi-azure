# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ActionGroup(pulumi.CustomResource):
    arm_role_receivers: pulumi.Output[list]
    """
    One or more `arm_role_receiver` blocks as defined below.

      * `name` (`str`) - The name of the ARM role receiver.
      * `roleId` (`str`) - The arm role id.
      * `useCommonAlertSchema` (`bool`) - Enables or disables the common alert schema.
    """
    automation_runbook_receivers: pulumi.Output[list]
    """
    One or more `automation_runbook_receiver` blocks as defined below.

      * `automationAccountId` (`str`) - The automation account ID which holds this runbook and authenticates to Azure resources.
      * `isGlobalRunbook` (`bool`) - Indicates whether this instance is global runbook.
      * `name` (`str`) - The name of the automation runbook receiver.
      * `runbook_name` (`str`) - The name for this runbook.
      * `service_uri` (`str`) - The URI where webhooks should be sent.
      * `useCommonAlertSchema` (`bool`) - Enables or disables the common alert schema.
      * `webhookResourceId` (`str`) - The resource id for webhook linked to this runbook.
    """
    azure_app_push_receivers: pulumi.Output[list]
    """
    One or more `azure_app_push_receiver` blocks as defined below.

      * `email_address` (`str`) - The email address of the user signed into the mobile app who will receive push notifications from this receiver.
      * `name` (`str`) - The name of the Azure app push receiver.
    """
    azure_function_receivers: pulumi.Output[list]
    """
    One or more `azure_function_receiver` blocks as defined below.

      * `functionAppResourceId` (`str`)
      * `functionName` (`str`) - The function name in the function app.
      * `httpTriggerUrl` (`str`) - The http trigger url where http request sent to.
      * `name` (`str`) - The name of the Azure Function receiver.
      * `useCommonAlertSchema` (`bool`) - Enables or disables the common alert schema.
    """
    email_receivers: pulumi.Output[list]
    """
    One or more `email_receiver` blocks as defined below.

      * `email_address` (`str`) - The email address of this receiver.
      * `name` (`str`) - The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group.
      * `useCommonAlertSchema` (`bool`) - Enables or disables the common alert schema.
    """
    enabled: pulumi.Output[bool]
    """
    Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
    """
    itsm_receivers: pulumi.Output[list]
    """
    One or more `itsm_receiver` blocks as defined below.

      * `connectionId` (`str`) - The unique connection identifier of the ITSM connection.
      * `name` (`str`) - The name of the ITSM receiver.
      * `region` (`str`) - The region of the workspace.
      * `ticketConfiguration` (`str`) - A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
      * `workspace_id` (`str`) - The Azure Log Analytics workspace ID where this connection is defined.
    """
    logic_app_receivers: pulumi.Output[list]
    """
    One or more `logic_app_receiver` blocks as defined below.

      * `callbackUrl` (`str`) - The callback url where http request sent to.
      * `name` (`str`) - The name of the logic app receiver.
      * `resource_id` (`str`) - The Azure resource ID of the logic app.
      * `useCommonAlertSchema` (`bool`) - Enables or disables the common alert schema.
    """
    name: pulumi.Output[str]
    """
    The name of the Action Group. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Action Group instance.
    """
    short_name: pulumi.Output[str]
    """
    The short name of the action group. This will be used in SMS messages.
    """
    sms_receivers: pulumi.Output[list]
    """
    One or more `sms_receiver` blocks as defined below.

      * `countryCode` (`str`) - The country code of the SMS receiver.
      * `name` (`str`) - The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group.
      * `phoneNumber` (`str`) - The phone number of the SMS receiver.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    voice_receivers: pulumi.Output[list]
    """
    One or more `voice_receiver` blocks as defined below.

      * `countryCode` (`str`) - The country code of the voice receiver.
      * `name` (`str`) - The name of the voice receiver.
      * `phoneNumber` (`str`) - The phone number of the voice receiver.
    """
    webhook_receivers: pulumi.Output[list]
    """
    One or more `webhook_receiver` blocks as defined below.

      * `name` (`str`) - The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
      * `service_uri` (`str`) - The URI where webhooks should be sent.
      * `useCommonAlertSchema` (`bool`) - Enables or disables the common alert schema.
    """
    def __init__(__self__, resource_name, opts=None, arm_role_receivers=None, automation_runbook_receivers=None, azure_app_push_receivers=None, azure_function_receivers=None, email_receivers=None, enabled=None, itsm_receivers=None, logic_app_receivers=None, name=None, resource_group_name=None, short_name=None, sms_receivers=None, tags=None, voice_receivers=None, webhook_receivers=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Action Group within Azure Monitor.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] arm_role_receivers: One or more `arm_role_receiver` blocks as defined below.
        :param pulumi.Input[list] automation_runbook_receivers: One or more `automation_runbook_receiver` blocks as defined below.
        :param pulumi.Input[list] azure_app_push_receivers: One or more `azure_app_push_receiver` blocks as defined below.
        :param pulumi.Input[list] azure_function_receivers: One or more `azure_function_receiver` blocks as defined below.
        :param pulumi.Input[list] email_receivers: One or more `email_receiver` blocks as defined below.
        :param pulumi.Input[bool] enabled: Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        :param pulumi.Input[list] itsm_receivers: One or more `itsm_receiver` blocks as defined below.
        :param pulumi.Input[list] logic_app_receivers: One or more `logic_app_receiver` blocks as defined below.
        :param pulumi.Input[str] name: The name of the Action Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Action Group instance.
        :param pulumi.Input[str] short_name: The short name of the action group. This will be used in SMS messages.
        :param pulumi.Input[list] sms_receivers: One or more `sms_receiver` blocks as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[list] voice_receivers: One or more `voice_receiver` blocks as defined below.
        :param pulumi.Input[list] webhook_receivers: One or more `webhook_receiver` blocks as defined below.

        The **arm_role_receivers** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the ARM role receiver.
          * `roleId` (`pulumi.Input[str]`) - The arm role id.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **automation_runbook_receivers** object supports the following:

          * `automationAccountId` (`pulumi.Input[str]`) - The automation account ID which holds this runbook and authenticates to Azure resources.
          * `isGlobalRunbook` (`pulumi.Input[bool]`) - Indicates whether this instance is global runbook.
          * `name` (`pulumi.Input[str]`) - The name of the automation runbook receiver.
          * `runbook_name` (`pulumi.Input[str]`) - The name for this runbook.
          * `service_uri` (`pulumi.Input[str]`) - The URI where webhooks should be sent.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.
          * `webhookResourceId` (`pulumi.Input[str]`) - The resource id for webhook linked to this runbook.

        The **azure_app_push_receivers** object supports the following:

          * `email_address` (`pulumi.Input[str]`) - The email address of the user signed into the mobile app who will receive push notifications from this receiver.
          * `name` (`pulumi.Input[str]`) - The name of the Azure app push receiver.

        The **azure_function_receivers** object supports the following:

          * `functionAppResourceId` (`pulumi.Input[str]`)
          * `functionName` (`pulumi.Input[str]`) - The function name in the function app.
          * `httpTriggerUrl` (`pulumi.Input[str]`) - The http trigger url where http request sent to.
          * `name` (`pulumi.Input[str]`) - The name of the Azure Function receiver.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **email_receivers** object supports the following:

          * `email_address` (`pulumi.Input[str]`) - The email address of this receiver.
          * `name` (`pulumi.Input[str]`) - The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **itsm_receivers** object supports the following:

          * `connectionId` (`pulumi.Input[str]`) - The unique connection identifier of the ITSM connection.
          * `name` (`pulumi.Input[str]`) - The name of the ITSM receiver.
          * `region` (`pulumi.Input[str]`) - The region of the workspace.
          * `ticketConfiguration` (`pulumi.Input[str]`) - A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
          * `workspace_id` (`pulumi.Input[str]`) - The Azure Log Analytics workspace ID where this connection is defined.

        The **logic_app_receivers** object supports the following:

          * `callbackUrl` (`pulumi.Input[str]`) - The callback url where http request sent to.
          * `name` (`pulumi.Input[str]`) - The name of the logic app receiver.
          * `resource_id` (`pulumi.Input[str]`) - The Azure resource ID of the logic app.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **sms_receivers** object supports the following:

          * `countryCode` (`pulumi.Input[str]`) - The country code of the SMS receiver.
          * `name` (`pulumi.Input[str]`) - The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group.
          * `phoneNumber` (`pulumi.Input[str]`) - The phone number of the SMS receiver.

        The **voice_receivers** object supports the following:

          * `countryCode` (`pulumi.Input[str]`) - The country code of the voice receiver.
          * `name` (`pulumi.Input[str]`) - The name of the voice receiver.
          * `phoneNumber` (`pulumi.Input[str]`) - The phone number of the voice receiver.

        The **webhook_receivers** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
          * `service_uri` (`pulumi.Input[str]`) - The URI where webhooks should be sent.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['arm_role_receivers'] = arm_role_receivers
            __props__['automation_runbook_receivers'] = automation_runbook_receivers
            __props__['azure_app_push_receivers'] = azure_app_push_receivers
            __props__['azure_function_receivers'] = azure_function_receivers
            __props__['email_receivers'] = email_receivers
            __props__['enabled'] = enabled
            __props__['itsm_receivers'] = itsm_receivers
            __props__['logic_app_receivers'] = logic_app_receivers
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if short_name is None:
                raise TypeError("Missing required property 'short_name'")
            __props__['short_name'] = short_name
            __props__['sms_receivers'] = sms_receivers
            __props__['tags'] = tags
            __props__['voice_receivers'] = voice_receivers
            __props__['webhook_receivers'] = webhook_receivers
        super(ActionGroup, __self__).__init__(
            'azure:monitoring/actionGroup:ActionGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arm_role_receivers=None, automation_runbook_receivers=None, azure_app_push_receivers=None, azure_function_receivers=None, email_receivers=None, enabled=None, itsm_receivers=None, logic_app_receivers=None, name=None, resource_group_name=None, short_name=None, sms_receivers=None, tags=None, voice_receivers=None, webhook_receivers=None):
        """
        Get an existing ActionGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] arm_role_receivers: One or more `arm_role_receiver` blocks as defined below.
        :param pulumi.Input[list] automation_runbook_receivers: One or more `automation_runbook_receiver` blocks as defined below.
        :param pulumi.Input[list] azure_app_push_receivers: One or more `azure_app_push_receiver` blocks as defined below.
        :param pulumi.Input[list] azure_function_receivers: One or more `azure_function_receiver` blocks as defined below.
        :param pulumi.Input[list] email_receivers: One or more `email_receiver` blocks as defined below.
        :param pulumi.Input[bool] enabled: Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
        :param pulumi.Input[list] itsm_receivers: One or more `itsm_receiver` blocks as defined below.
        :param pulumi.Input[list] logic_app_receivers: One or more `logic_app_receiver` blocks as defined below.
        :param pulumi.Input[str] name: The name of the Action Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Action Group instance.
        :param pulumi.Input[str] short_name: The short name of the action group. This will be used in SMS messages.
        :param pulumi.Input[list] sms_receivers: One or more `sms_receiver` blocks as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[list] voice_receivers: One or more `voice_receiver` blocks as defined below.
        :param pulumi.Input[list] webhook_receivers: One or more `webhook_receiver` blocks as defined below.

        The **arm_role_receivers** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the ARM role receiver.
          * `roleId` (`pulumi.Input[str]`) - The arm role id.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **automation_runbook_receivers** object supports the following:

          * `automationAccountId` (`pulumi.Input[str]`) - The automation account ID which holds this runbook and authenticates to Azure resources.
          * `isGlobalRunbook` (`pulumi.Input[bool]`) - Indicates whether this instance is global runbook.
          * `name` (`pulumi.Input[str]`) - The name of the automation runbook receiver.
          * `runbook_name` (`pulumi.Input[str]`) - The name for this runbook.
          * `service_uri` (`pulumi.Input[str]`) - The URI where webhooks should be sent.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.
          * `webhookResourceId` (`pulumi.Input[str]`) - The resource id for webhook linked to this runbook.

        The **azure_app_push_receivers** object supports the following:

          * `email_address` (`pulumi.Input[str]`) - The email address of the user signed into the mobile app who will receive push notifications from this receiver.
          * `name` (`pulumi.Input[str]`) - The name of the Azure app push receiver.

        The **azure_function_receivers** object supports the following:

          * `functionAppResourceId` (`pulumi.Input[str]`)
          * `functionName` (`pulumi.Input[str]`) - The function name in the function app.
          * `httpTriggerUrl` (`pulumi.Input[str]`) - The http trigger url where http request sent to.
          * `name` (`pulumi.Input[str]`) - The name of the Azure Function receiver.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **email_receivers** object supports the following:

          * `email_address` (`pulumi.Input[str]`) - The email address of this receiver.
          * `name` (`pulumi.Input[str]`) - The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **itsm_receivers** object supports the following:

          * `connectionId` (`pulumi.Input[str]`) - The unique connection identifier of the ITSM connection.
          * `name` (`pulumi.Input[str]`) - The name of the ITSM receiver.
          * `region` (`pulumi.Input[str]`) - The region of the workspace.
          * `ticketConfiguration` (`pulumi.Input[str]`) - A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
          * `workspace_id` (`pulumi.Input[str]`) - The Azure Log Analytics workspace ID where this connection is defined.

        The **logic_app_receivers** object supports the following:

          * `callbackUrl` (`pulumi.Input[str]`) - The callback url where http request sent to.
          * `name` (`pulumi.Input[str]`) - The name of the logic app receiver.
          * `resource_id` (`pulumi.Input[str]`) - The Azure resource ID of the logic app.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.

        The **sms_receivers** object supports the following:

          * `countryCode` (`pulumi.Input[str]`) - The country code of the SMS receiver.
          * `name` (`pulumi.Input[str]`) - The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group.
          * `phoneNumber` (`pulumi.Input[str]`) - The phone number of the SMS receiver.

        The **voice_receivers** object supports the following:

          * `countryCode` (`pulumi.Input[str]`) - The country code of the voice receiver.
          * `name` (`pulumi.Input[str]`) - The name of the voice receiver.
          * `phoneNumber` (`pulumi.Input[str]`) - The phone number of the voice receiver.

        The **webhook_receivers** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group.
          * `service_uri` (`pulumi.Input[str]`) - The URI where webhooks should be sent.
          * `useCommonAlertSchema` (`pulumi.Input[bool]`) - Enables or disables the common alert schema.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arm_role_receivers"] = arm_role_receivers
        __props__["automation_runbook_receivers"] = automation_runbook_receivers
        __props__["azure_app_push_receivers"] = azure_app_push_receivers
        __props__["azure_function_receivers"] = azure_function_receivers
        __props__["email_receivers"] = email_receivers
        __props__["enabled"] = enabled
        __props__["itsm_receivers"] = itsm_receivers
        __props__["logic_app_receivers"] = logic_app_receivers
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["short_name"] = short_name
        __props__["sms_receivers"] = sms_receivers
        __props__["tags"] = tags
        __props__["voice_receivers"] = voice_receivers
        __props__["webhook_receivers"] = webhook_receivers
        return ActionGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

