# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PolicyVM(pulumi.CustomResource):
    backup: pulumi.Output[dict]
    """
    Configures the Policy backup frequency, times & days as documented in the `backup` block below. 
    
      * `frequency` (`str`)
      * `time` (`str`)
      * `weekdays` (`list`)
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
    """
    recovery_vault_name: pulumi.Output[str]
    """
    Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
    """
    retention_daily: pulumi.Output[dict]
    """
    Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
    
      * `count` (`float`)
    """
    retention_monthly: pulumi.Output[dict]
    """
    Configures the policy monthly retention as documented in the `retention_monthly` block below.
    
      * `count` (`float`)
      * `weekdays` (`list`)
      * `weeks` (`list`)
    """
    retention_weekly: pulumi.Output[dict]
    """
    Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
    
      * `count` (`float`)
      * `weekdays` (`list`)
    """
    retention_yearly: pulumi.Output[dict]
    """
    Configures the policy yearly retention as documented in the `retention_yearly` block below.
    
      * `count` (`float`)
      * `months` (`list`)
      * `weekdays` (`list`)
      * `weeks` (`list`)
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    timezone: pulumi.Output[str]
    """
    Specifies the timezone. Defaults to `UTC`
    """
    def __init__(__self__, resource_name, opts=None, backup=None, name=None, recovery_vault_name=None, resource_group_name=None, retention_daily=None, retention_monthly=None, retention_weekly=None, retention_yearly=None, tags=None, timezone=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Backup VM Backup Policy.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] backup: Configures the Policy backup frequency, times & days as documented in the `backup` block below. 
        :param pulumi.Input[str] name: Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        :param pulumi.Input[dict] retention_monthly: Configures the policy monthly retention as documented in the `retention_monthly` block below.
        :param pulumi.Input[dict] retention_weekly: Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        :param pulumi.Input[dict] retention_yearly: Configures the policy yearly retention as documented in the `retention_yearly` block below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] timezone: Specifies the timezone. Defaults to `UTC`
        
        The **backup** object supports the following:
        
          * `frequency` (`pulumi.Input[str]`)
          * `time` (`pulumi.Input[str]`)
          * `weekdays` (`pulumi.Input[list]`)
        
        The **retention_daily** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
        
        The **retention_monthly** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
          * `weekdays` (`pulumi.Input[list]`)
          * `weeks` (`pulumi.Input[list]`)
        
        The **retention_weekly** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
          * `weekdays` (`pulumi.Input[list]`)
        
        The **retention_yearly** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
          * `months` (`pulumi.Input[list]`)
          * `weekdays` (`pulumi.Input[list]`)
          * `weeks` (`pulumi.Input[list]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/backup_policy_vm.html.markdown.
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

            if backup is None:
                raise TypeError("Missing required property 'backup'")
            __props__['backup'] = backup
            __props__['name'] = name
            if recovery_vault_name is None:
                raise TypeError("Missing required property 'recovery_vault_name'")
            __props__['recovery_vault_name'] = recovery_vault_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retention_daily'] = retention_daily
            __props__['retention_monthly'] = retention_monthly
            __props__['retention_weekly'] = retention_weekly
            __props__['retention_yearly'] = retention_yearly
            __props__['tags'] = tags
            __props__['timezone'] = timezone
        super(PolicyVM, __self__).__init__(
            'azure:backup/policyVM:PolicyVM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backup=None, name=None, recovery_vault_name=None, resource_group_name=None, retention_daily=None, retention_monthly=None, retention_weekly=None, retention_yearly=None, tags=None, timezone=None):
        """
        Get an existing PolicyVM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] backup: Configures the Policy backup frequency, times & days as documented in the `backup` block below. 
        :param pulumi.Input[str] name: Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        :param pulumi.Input[dict] retention_monthly: Configures the policy monthly retention as documented in the `retention_monthly` block below.
        :param pulumi.Input[dict] retention_weekly: Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        :param pulumi.Input[dict] retention_yearly: Configures the policy yearly retention as documented in the `retention_yearly` block below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] timezone: Specifies the timezone. Defaults to `UTC`
        
        The **backup** object supports the following:
        
          * `frequency` (`pulumi.Input[str]`)
          * `time` (`pulumi.Input[str]`)
          * `weekdays` (`pulumi.Input[list]`)
        
        The **retention_daily** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
        
        The **retention_monthly** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
          * `weekdays` (`pulumi.Input[list]`)
          * `weeks` (`pulumi.Input[list]`)
        
        The **retention_weekly** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
          * `weekdays` (`pulumi.Input[list]`)
        
        The **retention_yearly** object supports the following:
        
          * `count` (`pulumi.Input[float]`)
          * `months` (`pulumi.Input[list]`)
          * `weekdays` (`pulumi.Input[list]`)
          * `weeks` (`pulumi.Input[list]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/backup_policy_vm.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["backup"] = backup
        __props__["name"] = name
        __props__["recovery_vault_name"] = recovery_vault_name
        __props__["resource_group_name"] = resource_group_name
        __props__["retention_daily"] = retention_daily
        __props__["retention_monthly"] = retention_monthly
        __props__["retention_weekly"] = retention_weekly
        __props__["retention_yearly"] = retention_yearly
        __props__["tags"] = tags
        __props__["timezone"] = timezone
        return PolicyVM(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
